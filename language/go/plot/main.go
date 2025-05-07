package main

//
//import (
//	"fmt"
//	"gonum.org/v1/plot"
//	"gonum.org/v1/plot/plotter"
//	"gonum.org/v1/plot/vg"
//	"gonum.org/v1/plot/vg/draw"
//	"image/color"
//	"log"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//// Point 数据结构
//type Point struct {
//	Time   float64
//	Field1 float64
//}
//
//// queryBatch 模拟分页查询
//func queryBatch(lastOffset int, batchSize int, totalPoints int) ([]Point, int) {
//	points := make([]Point, 0, batchSize)
//	for i := 0; i < batchSize; i++ {
//		currentIndex := lastOffset + i
//		if currentIndex >= totalPoints {
//			break
//		}
//		pt := Point{
//			Time:   float64(currentIndex),
//			Field1: 100 + 10*rand.NormFloat64(),
//		}
//		points = append(points, pt)
//	}
//	newOffset := lastOffset + batchSize
//	return points, newOffset
//}
//
//// MaxMinDownSample 按组压缩
//func MaxMinDownSample(points []Point) []Point {
//	if len(points) == 0 {
//		return nil
//	}
//	maxP, minP := points[0], points[0]
//	for _, p := range points {
//		if p.Field1 > maxP.Field1 {
//			maxP = p
//		}
//		if p.Field1 < minP.Field1 {
//			minP = p
//		}
//	}
//	return []Point{minP, maxP}
//}
//
//// makeSparseTicks 生成稀疏的x轴刻度
//func makeSparseTicks(start, end float64, desiredTicks int) []plot.Tick {
//	ticks := []plot.Tick{}
//	if desiredTicks <= 0 || end <= start {
//		return ticks
//	}
//	interval := (end - start) / float64(desiredTicks)
//	for i := 0; i <= desiredTicks; i++ {
//		value := start + interval*float64(i)
//		ticks = append(ticks, plot.Tick{
//			Value: value,
//			Label: fmt.Sprintf("%.0f", value),
//		})
//	}
//	return ticks
//}
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	startTime := time.Now()
//	// 配置
//	totalPoints := 10_000_000_000 // 100亿条
//	targetDrawPoints := 2000
//
//	fmt.Printf("模拟总数据量: %d\n", totalPoints)
//
//	// 动态计算压缩比例
//	compressionRatio := totalPoints / (targetDrawPoints / 2)
//	fmt.Printf("动态计算 compressionRatio: %d\n", compressionRatio)
//
//	// 初始化Plot
//	p := plot.New()
//	p.Title.Text = "Concurrent Streaming Downsampled Curve"
//	p.X.Label.Text = "Time (timestamp)"
//	p.Y.Label.Text = "Field1 Value"
//
//	line := &plotter.Line{
//		XYs: plotter.XYs{},
//		LineStyle: draw.LineStyle{
//			Color:  color.RGBA{R: 255, G: 99, B: 71, A: 255},
//			Width:  vg.Points(0.5),
//			Dashes: []vg.Length{},
//		},
//		FillColor: nil,
//	}
//
//	// 并发流式压缩配置
//	batchSize := 500_000          // 每次小批量读取
//	groupSize := compressionRatio // 每组多少条压成2点
//	numWorkers := 8               // 开8个worker
//
//	workChan := make(chan []Point, numWorkers)
//	resultChan := make(chan []Point, numWorkers)
//
//	var wg sync.WaitGroup
//
//	// worker池，处理group压缩
//	for i := 0; i < numWorkers; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			for group := range workChan {
//				sampled := MaxMinDownSample(group)
//				resultChan <- sampled
//			}
//		}()
//	}
//
//	// 收集压缩后的结果
//	collectedLine := make(plotter.XYs, 0, targetDrawPoints*2)
//	var collectWg sync.WaitGroup
//	collectWg.Add(1)
//
//	go func() {
//		defer collectWg.Done()
//		for sampled := range resultChan {
//			for _, pt := range sampled {
//				collectedLine = append(collectedLine, plotter.XY{X: pt.Time, Y: pt.Field1})
//			}
//		}
//	}()
//
//	var offset int
//	var groupBuffer []Point
//	var minX, maxX float64
//	firstPoint := true
//
//	// 主线程流式读取小批次
//	for {
//		batch, newOffset := queryBatch(offset, batchSize, totalPoints)
//		if len(batch) == 0 {
//			break
//		}
//		offset = newOffset
//
//		for _, pt := range batch {
//			groupBuffer = append(groupBuffer, pt)
//			if len(groupBuffer) >= int(groupSize) {
//				tmp := make([]Point, len(groupBuffer))
//				copy(tmp, groupBuffer)
//				workChan <- tmp
//				groupBuffer = groupBuffer[:0]
//			}
//		}
//
//		fmt.Printf("读取到偏移量: %d\n", offset)
//	}
//
//	// 最后处理不足一组的数据
//	if len(groupBuffer) > 0 {
//		workChan <- groupBuffer
//		groupBuffer = nil
//	}
//
//	close(workChan) // 发完了任务
//
//	wg.Wait()         // 等所有worker结束
//	close(resultChan) // 关闭收集通道
//	collectWg.Wait()  // 等收集线程结束
//
//	// 设置X轴范围
//	for _, pt := range collectedLine {
//		if firstPoint {
//			minX = pt.X
//			firstPoint = false
//		}
//		maxX = pt.X
//	}
//
//	line.XYs = collectedLine
//
//	fmt.Printf("最终绘制点数: %d\n", len(line.XYs))
//
//	p.Add(line)
//	p.X.Tick.Marker = plot.ConstantTicks(makeSparseTicks(minX, maxX, 12))
//
//	// 保存PNG
//	sizeL := 10 * vg.Inch
//	sizeH := 5 * vg.Inch
//
//	if err := p.Save(sizeL, sizeH, "./final_concurrent_stable_curve.png"); err != nil {
//		log.Fatalf("保存图片失败: %v", err)
//	}
//
//	defer func() {
//		fmt.Printf("程序运行时间: %s\n", time.Since(startTime))
//	}()
//	fmt.Println("绘制完成！保存为 final_concurrent_stable_curve.png")
//}
