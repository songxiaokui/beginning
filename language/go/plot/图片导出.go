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

const (
	defaultFontFile      = "NotoSansSC-VariableFont_wght.ttf"
	defaultFontName      = "SimHei"
	defaultTargetPoints  = 2000
	defaultBatchSize     = 200_000
	defaultNumWorkers    = 8
	defaultPlotWidth     = 20 * vg.Inch
	defaultPlotHeight    = 10 * vg.Inch
	defaultDesiredXTicks = 12
	minGroupSize         = 1
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
	switch fieldName {
	case "Field1":
		return m.Field1
	case "Field2":
		return m.Field2
	case "Field3":
		return m.Field3
	case "Field4":
		return m.Field4
	case "Field5":
		return m.Field5
	case "Time":
		return m.Time
	default:
		return 0
	}
}

var (
	_font        *opentype.Font
	fontLoaded   bool
	loadFontOnce sync.Once
)

func initFont() error {
	var initErr error
	loadFontOnce.Do(func() {
		ttfBytes, err := os.ReadFile(defaultFontFile)
		if err != nil {
			initErr = fmt.Errorf("failed to read font file: %w", err)
			return
		}

		_font, err = opentype.Parse(ttfBytes)
		if err != nil {
			initErr = fmt.Errorf("failed to parse font: %w", err)
			return
		}

		arial := font.Font{
			Typeface: font.Typeface(defaultFontName),
		}
		font.DefaultCache.Add([]font.Face{
			{
				Font: arial,
				Face: _font,
			},
		})
		plot.DefaultFont = arial
		fontLoaded = true
	})
	return initErr
}

// DownSampleAndPlot processes data once and generates multiple plots based on configs
// DownSampleAndPlot processes data once and generates multiple plots based on configs
func DownSampleAndPlot(dataSource DataSource2, targetDrawPoints int, batchSize int, numWorkers int, configs []GraphConfig, totalCount int) error {
	if err := initFont(); err != nil {
		return fmt.Errorf("font initialization failed: %w", err)
	}

	if targetDrawPoints <= 0 {
		targetDrawPoints = defaultTargetPoints
	}
	if batchSize <= 0 {
		batchSize = defaultBatchSize
	}
	if numWorkers <= 0 {
		numWorkers = defaultNumWorkers
	}

	// Initialize drawDataList to store points for each configuration
	drawDataList := make([][]plotter.XY, len(configs))
	for i := range drawDataList {
		drawDataList[i] = make([]plotter.XY, 0, targetDrawPoints)
	}

	// Worker pool for concurrent downsampling
	workChan := make(chan []Point, numWorkers)
	resultChans := make([]chan []Point, len(configs))
	for i := range resultChans {
		resultChans[i] = make(chan []Point, numWorkers)
	}

	var wg sync.WaitGroup
	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for group := range workChan {
				for cfgIdx, cfg := range configs {
					sampled := group
					if totalCount > targetDrawPoints {
						groupSize := calculateGroupSize(totalCount, targetDrawPoints)
						sampled = GroupDownSample(group, cfg.Compare, cfg.Sort, groupSize)
					}
					resultChans[cfgIdx] <- sampled
				}
			}
		}()
	}

	// Collect results
	var collectWg sync.WaitGroup
	for cfgIdx := range configs {
		collectWg.Add(1)
		go func(idx int) {
			defer collectWg.Done()

			var cacheDrawData []Point
			for sampled := range resultChans[idx] {
				cacheDrawData = append(cacheDrawData, sampled...)
			}

			latestData := cacheDrawData
			if len(cacheDrawData) > targetDrawPoints {
				groupSize := calculateGroupSize(len(cacheDrawData), targetDrawPoints)
				latestData = GroupDownSample(cacheDrawData, configs[idx].Compare, configs[idx].Sort, groupSize)
			}

			for _, pt := range latestData {
				drawDataList[idx] = append(drawDataList[idx], plotter.XY{
					X: pt.GetValueByField(configs[idx].X),
					Y: pt.GetValueByField(configs[idx].Y),
				})
			}
		}(cfgIdx)
	}

	// Stream data from source
	var offset int
	var groupBuffer []Point
	for {
		batch, err := dataSource.GetData(offset, batchSize)
		if err != nil {
			close(workChan)
			wg.Wait()
			for _, ch := range resultChans {
				close(ch)
			}
			collectWg.Wait()
			return fmt.Errorf("error getting data: %w", err)
		}
		if len(batch) == 0 {
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
		offset += len(batch)
	}

	// Process remaining data
	if len(groupBuffer) > 0 {
		workChan <- groupBuffer
	}

	close(workChan)
	wg.Wait()
	for _, ch := range resultChans {
		close(ch)
	}
	collectWg.Wait()

	// Generate plots
	var errs []error
	for cfgIdx, cfg := range configs {
		if err := plotData(drawDataList[cfgIdx], cfg); err != nil {
			errs = append(errs, fmt.Errorf("failed to plot %s vs %s: %w", cfg.X, cfg.Y, err))
			continue
		}
		log.Printf("Plot %s vs %s completed! Saved as %s, points: %d",
			cfg.X, cfg.Y, cfg.Path, len(drawDataList[cfgIdx]))
	}

	if len(errs) > 0 {
		return fmt.Errorf("encountered %d errors during plotting", len(errs))
	}
	return nil
}
func calculateGroupSize(totalPoints, targetPoints int) int {
	group := targetPoints / 2
	if group < 1 {
		return minGroupSize
	}
	groupSize := totalPoints / group
	if groupSize < minGroupSize {
		return minGroupSize
	}
	return groupSize
}

func plotData(data plotter.XYs, config GraphConfig) error {
	p := plot.New()
	p.Title.Text = fmt.Sprintf("Concurrent Streaming Downsampled Curve - %s vs %s", config.X, config.Y)
	p.X.Label.Text = config.X
	p.Y.Label.Text = config.Y

	line, err := plotter.NewLine(data)
	if err != nil {
		return fmt.Errorf("failed to create line: %w", err)
	}
	line.LineStyle.Color = color.RGBA{R: 255, G: 99, B: 71, A: 255}
	line.LineStyle.Width = vg.Points(0.5)

	minX, maxX := calculateDataRange(data)
	p.X.Tick.Marker = plot.ConstantTicks(makeSparseTicks(minX, maxX, defaultDesiredXTicks))
	p.Add(line)

	if err := p.Save(defaultPlotWidth, defaultPlotHeight, config.Path); err != nil {
		return fmt.Errorf("failed to save plot: %w", err)
	}
	return nil
}

func calculateDataRange(data plotter.XYs) (minX, maxX float64) {
	if len(data) == 0 {
		return 0, 0
	}
	minX, maxX = data[0].X, data[0].X
	for _, pt := range data[1:] {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.X > maxX {
			maxX = pt.X
		}
	}
	return minX, maxX
}

type DataSource2 interface {
	GetData(offset, batchSize int) ([]Point, error)
}

type CompareFunc func(p1, p2 Point) bool
type SortFunc func(p1, p2 Point) bool

type GraphConfig struct {
	X       string
	Y       string
	Path    string
	Compare CompareFunc
	Sort    SortFunc
}

func GroupDownSample(points []Point, compare CompareFunc, sortFunc SortFunc, blockCount int) []Point {
	if len(points) == 0 {
		return nil
	}
	if blockCount <= 0 {
		blockCount = minGroupSize
	}

	result := make([]Point, 0, len(points)/blockCount*2)
	for i := 0; i < len(points); i += blockCount {
		end := i + blockCount
		if end > len(points) {
			end = len(points)
		}
		segment := points[i:end]
		if len(segment) == 0 {
			continue
		}

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

func makeSparseTicks(start, end float64, desiredTicks int) []plot.Tick {
	if desiredTicks <= 0 || end <= start {
		return nil
	}

	ticks := make([]plot.Tick, 0, desiredTicks+1)
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

type MockDataSource2 struct {
	TotalPoints int
	mu          sync.Mutex
	rand        *rand.Rand
}

func NewMockDataSource2(totalPoints int) *MockDataSource2 {
	return &MockDataSource2{
		TotalPoints: totalPoints,
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (ds *MockDataSource2) GetData(offset, batchSize int) ([]Point, error) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if offset >= ds.TotalPoints {
		return nil, nil
	}

	if offset+batchSize > ds.TotalPoints {
		batchSize = ds.TotalPoints - offset
	}

	points := make([]Point, 0, batchSize)
	for i := 0; i < batchSize; i++ {
		currentIndex := offset + i
		points = append(points, &moveData{
			Time:   float64(currentIndex),
			Field1: 200 + 5*ds.rand.NormFloat64(),
			Field2: 100 + 3*ds.rand.NormFloat64(),
			Field3: 50 + 1*ds.rand.NormFloat64(),
			Field4: 130 + 3*ds.rand.NormFloat64(),
			Field5: 30 + 5*ds.rand.NormFloat64(),
		})
	}
	return points, nil
}

func main() {
	startTime := time.Now()
	defer func() {
		log.Printf("Time consumed: %.2f seconds", time.Since(startTime).Seconds())
	}()

	// Configuration
	totalPoints := 100
	targetDrawPoints := defaultTargetPoints
	batchSize := defaultBatchSize
	numWorkers := defaultNumWorkers

	// Create data source
	dataSource := NewMockDataSource2(totalPoints)

	// Define graph configurations
	graphConfigs := []GraphConfig{
		{
			X:    "Time",
			Y:    "Field5",
			Path: "./Time_vs_Field5.png",
			Compare: func(p1, p2 Point) bool {
				return p1.GetValueByField("Field5") > p2.GetValueByField("Field5")
			},
			Sort: func(p1, p2 Point) bool {
				return p1.GetValueByField("Time") < p2.GetValueByField("Time")
			},
		},
		{
			X:    "Time",
			Y:    "Field2",
			Path: "./Time_vs_Field2.png",
			Compare: func(p1, p2 Point) bool {
				return p1.GetValueByField("Field2") > p2.GetValueByField("Field2")
			},
			Sort: func(p1, p2 Point) bool {
				return p1.GetValueByField("Time") < p2.GetValueByField("Time")
			},
		},
		{
			X:    "Time",
			Y:    "Field4",
			Path: "./Time_vs_Field4.png",
			Compare: func(p1, p2 Point) bool {
				return p1.GetValueByField("Field4") > p2.GetValueByField("Field4")
			},
			Sort: func(p1, p2 Point) bool {
				return p1.GetValueByField("Time") < p2.GetValueByField("Time")
			},
		},
	}

	if err := DownSampleAndPlot(dataSource, targetDrawPoints, batchSize, numWorkers, graphConfigs, totalPoints); err != nil {
		log.Fatalf("Data processing and plotting failed: %v", err)
	}
}
