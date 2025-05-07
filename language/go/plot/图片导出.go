package main

import (
	"fmt"
	"golang.org/x/image/font/opentype"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

// Point 数据结构
type Point interface {
	GetValueByField(fieldName string) float64
}

type moveData struct {
	Time   float64
	Field1 float64
	Field2 float64
	Field3 float64
	Field4 float64
	Field5 float64
}

func (m *moveData) GetValueByField(fieldName string) float64 {
	var value float64
	switch fieldName {
	case "Field1":
		value = m.Field1
	case "Field2":
		value = m.Field2
	case "Field3":
		value = m.Field3
	case "Field4":
		value = m.Field4
	case "Field5":
		value = m.Field5
	case "Time":
		value = m.Time
	}
	return value
}

var _font *opentype.Font
var fontName = "SimHei"

func init() {
	ttfBytes, err := os.ReadFile("NotoSansSC-VariableFont_wght.ttf")
	if err != nil {
		panic(fmt.Errorf("读取字体文件失败：%v", err.Error()))
	}
	_font, _ = opentype.Parse(ttfBytes)
	arial := font.Font{
		Typeface: font.Typeface(fontName),
	}
	font.DefaultCache.Add([]font.Face{
		{
			Font: arial,
			Face: _font,
		},
	})
	plot.DefaultFont = arial
}

// DownSampleAndPlot processes data once and generates multiple plots based on configs
func DownSampleAndPlot(dataSource Data2Source, targetDrawPoints int, batchSize int, numWorkers int, configs []GraphConfig, totalCount int) error {
	// Initialize drawDataList to store points for each configuration
	drawDataList := make([][]plotter.XY, len(configs))

	// Initialize result channels for each configuration
	resultChans := make([]chan []Point, len(configs))
	for i := range resultChans {
		resultChans[i] = make(chan []Point, numWorkers)
		drawDataList[i] = make([]plotter.XY, 0, targetDrawPoints)
	}

	// Worker pool for concurrent downsampling
	workChan := make(chan []Point, numWorkers)
	var wg sync.WaitGroup

	// Start workers for each configuration
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for group := range workChan {
				for cfgIdx, cfg := range configs {
					sampled := group

					// 计算几个点进行采样
					_group := targetDrawPoints / 2
					// 计算几个点为一个组
					groupSize := totalCount / _group

					fmt.Printf("总共 %d 个组, 每 %d 个数据为一个组 取最大值和最小值\n", _group, groupSize)
					fmt.Println("分块数据长度: ", len(sampled))

					if totalCount > targetDrawPoints {
						sampled = GroupDownSample(group, cfg.Compare, cfg.Sort, groupSize)
						fmt.Printf("Batch sampled points: %d\n", len(sampled))
					}

					resultChans[cfgIdx] <- sampled

				}
			}
		}()
	}

	// Collect results for each configuration
	var collectWg sync.WaitGroup
	for cfgIdx := range configs {
		collectWg.Add(1)
		go func(idx int) {
			defer collectWg.Done()

			cacheDrawData := make([]Point, 0)
			for sampled := range resultChans[idx] {
				cacheDrawData = append(cacheDrawData, sampled...)
			}

			latestData := cacheDrawData
			if len(cacheDrawData) > targetDrawPoints {
				fmt.Println("数据大于目标绘制 需要进行二次压缩")
				// 计算几个点进行采样
				_group := targetDrawPoints / 2
				// 计算几个点为一个组
				groupSize := len(cacheDrawData) / _group
				fmt.Println("---groupSize---", groupSize)
				latestData = GroupDownSample(cacheDrawData, configs[idx].Compare, configs[idx].Sort, groupSize)
				fmt.Println("压缩后的数据长度: ", len(latestData))
			}

			for _, pt := range latestData {
				var xValue, yValue float64
				xValue = pt.GetValueByField(configs[idx].X)
				yValue = pt.GetValueByField(configs[idx].Y)
				drawDataList[idx] = append(drawDataList[idx], plotter.XY{X: xValue, Y: yValue})
			}

			fmt.Printf("Config %d drawDataList size: %d\n", idx, len(drawDataList[idx]))

		}(cfgIdx)
	}

	// Stream data from source and count total points
	var offset int
	var groupBuffer []Point
	for {
		batch, err := dataSource.GetData(offset, batchSize)
		if err != nil || len(batch) == 0 {
			break
		}
		for _, pt := range batch {
			groupBuffer = append(groupBuffer, pt)
			if len(groupBuffer) >= batchSize {
				tmp := make([]Point, len(groupBuffer))
				copy(tmp, groupBuffer)
				workChan <- tmp
				groupBuffer = groupBuffer[:0]
			}
		}
		offset += batchSize
	}

	// Process remaining data
	if len(groupBuffer) > 0 {
		workChan <- groupBuffer
	}

	// Close channels and wait for processing
	close(workChan)
	wg.Wait()
	for _, ch := range resultChans {
		close(ch)
	}
	collectWg.Wait()

	// Ensure drawDataList does not exceed targetDrawPoints for large datasets

	// Generate plots for each configuration
	for cfgIdx, cfg := range configs {
		fmt.Printf("index: %d, count: %d \n", cfgIdx, len(drawDataList[cfgIdx]))
		if err := plotData(drawDataList[cfgIdx], cfg); err != nil {
			log.Printf("绘制 %s vs %s 失败: %v", cfg.X, cfg.Y, err)
		}
	}

	return nil
}

// plotData creates and saves a plot for the given data and configuration
func plotData(data plotter.XYs, config GraphConfig) error {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Concurrent Streaming Downsampled Curve - %s vs %s", config.X, config.Y)
	p.X.Label.Text = config.X
	p.Y.Label.Text = config.Y

	// Create line plot
	line, err := plotter.NewLine(data)
	if err != nil {
		return fmt.Errorf("创建线条失败: %v", err)
	}
	line.LineStyle.Color = color.RGBA{R: 255, G: 99, B: 71, A: 255}
	line.LineStyle.Width = vg.Points(0.5)

	// Determine X-axis range for ticks
	var minX, maxX float64
	firstPoint := true
	for _, pt := range data {
		if firstPoint {
			minX = pt.X
			maxX = pt.X
			firstPoint = false
		}
		if pt.X < minX {
			minX = pt.X
		}
		if pt.X > maxX {
			maxX = pt.X
		}
	}

	// Add sparse ticks
	p.X.Tick.Marker = plot.ConstantTicks(makeSparseTicks(minX, maxX, 12))
	p.Add(line)

	// Save plot
	sizeL := 20 * vg.Inch
	sizeH := 10 * vg.Inch
	if err := p.Save(sizeL, sizeH, config.Path); err != nil {
		return fmt.Errorf("保存绘图失败: %v", err)
	}

	fmt.Printf("%s vs %s 图表完成！保存为 %s，绘制点数: %d\n", config.X, config.Y, config.Path, len(data))
	return nil
}

// Data2Source 定义数据源接口
type Data2Source interface {
	GetData(offset, batchSize int) ([]Point, error)
}

// compare 函数类型，定义比较逻辑
type compare func(p1, p2 Point) bool
type sortFunc func(p1, p2 Point) bool

// GraphConfig 配置图形的结构体
type GraphConfig struct {
	X       string
	Y       string
	Path    string
	Compare compare
	Sort    sortFunc
}

// GroupDownSample compresses points by grouping and taking max/min
func GroupDownSample(points []Point, compare compare, sortFunc sortFunc, blockCount int) []Point {
	var result []Point
	if len(points) == 0 {
		return result
	}
	if blockCount <= 0 {
		blockCount = 1
	}

	for i := 0; i < len(points); i += blockCount {
		end := i + blockCount
		if end > len(points) {
			end = len(points)
		}
		segment := points[i:end]
		maxP, minP := segment[0], segment[0]
		for _, p := range segment[1:] {
			if compare(p, maxP) {
				maxP = p
			}
			if !compare(p, minP) {
				minP = p
			}
		}
		result = append(result, minP, maxP)
	}

	sort.Slice(result, func(i, j int) bool {
		return sortFunc(result[i], result[j])
	})

	return result
}

// makeSparseTicks 生成稀疏x轴刻度
func makeSparseTicks(start, end float64, desiredTicks int) []plot.Tick {
	var ticks []plot.Tick
	if desiredTicks <= 0 || end <= start {
		return ticks
	}
	interval := (end - start) / float64(desiredTicks)
	for i := 0; i <= desiredTicks; i++ {
		value := start + interval*float64(i)
		ticks = append(ticks, plot.Tick{
			Value: value,
			Label: fmt.Sprintf("%.0f", value),
		})
	}
	return ticks
}

// MockData2Source 数据源示例实现（模拟数据）
type MockData2Source struct {
	TotalPoints int
}

func (ds *MockData2Source) GetData(offset, batchSize int) ([]Point, error) {
	rand.Seed(time.Now().UnixNano())
	points := make([]Point, 0, batchSize)
	for i := 0; i < batchSize; i++ {
		currentIndex := offset + i
		if currentIndex >= ds.TotalPoints {
			break
		}
		pt := &moveData{
			Time:   float64(currentIndex),
			Field1: 200 + 5*rand.NormFloat64(),
			Field2: 100 + 3*rand.NormFloat64(),
			Field3: 50 + 1*rand.NormFloat64(),
			Field4: 130 + 3*rand.NormFloat64(),
			Field5: 30 + 5*rand.NormFloat64(),
		}
		points = append(points, pt)
	}
	return points, nil
}

func main() {
	// 配置
	totalPoints := 30_0000_0000 // 测试介于1000-2000
	//totalPoints := 4_000_0000 // 测试介于1000-2000
	//totalPoints := 4000 // 测试介于1000-2000
	//totalPoints := 1500 // 测试介于1000-2000
	targetDrawPoints := 2000
	batchSize := 200_000
	numWorkers := 8

	startTime := time.Now()
	defer func() {
		fmt.Printf("time counsume: %.2f\n", time.Now().Sub(startTime).Seconds())
	}()
	// 模拟数据源
	dataSource := &MockData2Source{TotalPoints: totalPoints}

	// 定义要生成的图的配置列表
	graphConfigs := []GraphConfig{
		{
			X:       "Time",
			Y:       "Field5",
			Path:    "./Time_vs_Field5.png",
			Compare: func(p1, p2 Point) bool { return p1.GetValueByField("Field5") > p1.GetValueByField("Field5") },
			Sort: func(p1, p2 Point) bool {
				return p1.GetValueByField("Time") < p2.GetValueByField("Time")
			},
		},
		{
			X:       "Time",
			Y:       "Field2",
			Path:    "./Time_vs_Field2.png",
			Compare: func(p1, p2 Point) bool { return p1.GetValueByField("Field2") > p1.GetValueByField("Field2") },
			Sort: func(p1, p2 Point) bool {
				return p1.GetValueByField("Time") < p2.GetValueByField("Time")
			},
		},
		{
			X:       "Time",
			Y:       "Field4",
			Path:    "./Time_vs_Field4.png",
			Compare: func(p1, p2 Point) bool { return p1.GetValueByField("Field4") > p1.GetValueByField("Field4") },
			Sort: func(p1, p2 Point) bool {
				return p1.GetValueByField("Time") < p2.GetValueByField("Time")
			},
		},
	}

	if err := DownSampleAndPlot(dataSource, targetDrawPoints, batchSize, numWorkers, graphConfigs, totalPoints); err != nil {
		log.Fatalf("数据压缩和绘图失败: %v", err)
	}
}
