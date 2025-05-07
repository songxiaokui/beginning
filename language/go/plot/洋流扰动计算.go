package main

import (
	"fmt"
	"math"
	"sync"
)

// calculateRMS 计算均方根平均值
func calculateRMS(data []float64) float64 {
	var sumOfSquares float64
	for _, value := range data {
		sumOfSquares += value * value
	}
	rms := math.Sqrt(sumOfSquares / float64(len(data)))
	return rms
}

// 数据源接口
type DataSource interface {
	// 获取数据，接收 offset 和 batchSize，返回相应的数据片段
	GetData(offset, batchSize int) ([]float64, error)
}

// 计算均方根（RMS）
func calculateRMSFromSource(ds DataSource, numWorkers, batchSize int) (float64, error) {
	var totalSumOfSquares float64
	var totalDataCount int
	var closeOnce sync.Once // 确保关闭 resultChan 只执行一次

	// 用于存储每个goroutine计算的平方和
	resultChan := make(chan float64, numWorkers)
	var wg sync.WaitGroup

	offset := 0
	for {
		// 从数据源获取数据
		data, err := ds.GetData(offset, batchSize)
		if err != nil {
			return 0, fmt.Errorf("Error getting data: %v", err)
		}

		// 如果没有更多数据，结束循环
		if len(data) == 0 {
			break
		}

		// 启动并行计算
		wg.Add(1)
		go func(data []float64) {
			defer wg.Done()

			// 计算该批次的平方和
			var sumOfSquares float64
			for _, value := range data {
				sumOfSquares += value * value
			}

			// 将每个批次的计算结果传递到 resultChan
			resultChan <- sumOfSquares
		}(data)

		// 更新偏移量
		offset += batchSize
	}

	// 等待所有 goroutine 完成
	go func() {
		wg.Wait()
		closeOnce.Do(func() {
			close(resultChan)
		})
	}()

	// 汇总所有批次的平方和
	for result := range resultChan {
		totalSumOfSquares += result
		totalDataCount += batchSize
	}

	// 计算最终的 RMS
	if totalDataCount == 0 {
		return 0, fmt.Errorf("No data to calculate RMS")
	}

	rms := math.Sqrt(totalSumOfSquares / float64(totalDataCount))
	return rms, nil
}

// 模拟的数据源实现
type MockDataSource struct {
	Data []float64
}

// GetData 模拟从数据库或其他源获取数据
func (ds *MockDataSource) GetData(offset, batchSize int) ([]float64, error) {
	// 模拟数据
	mockAllData := []float64{10, 79, 77, 97, 91, 70, 70, 89, 19, 66}

	if offset+batchSize > len(mockAllData) {
		return mockAllData[offset:], nil
	}
	if offset >= len(mockAllData) {
		return []float64{}, nil
	}
	return mockAllData[offset : offset+batchSize], nil
}

func main() {
	// 提取洋流干扰度下纬度误差参数lat_error_turbulence
	// 暂时和 lat_error_list := []float64{10, 79, 77, 97, 91, 70, 70, 89, 19, 66}
	lat_error_turbulence := []float64{10, 79, 77, 97, 91, 70, 70, 89, 19, 66}

	// 计算纬度误差参数的均方根平均值rms_turbulence
	// 此处应该改成分批次读取数据 然后进行计算
	// rms_turbulence := calculateRMS(lat_error_turbulence)
	rms_turbulence, _ := calculateRMSFromSource(&MockDataSource{Data: lat_error_turbulence}, 4, 1)
	fmt.Printf("纬度误差参数的均方根： %v", rms_turbulence)

	/*
			step3：等级，根据rms_turbulence选择对应等级->level_turbulence。
			step4：适应性评价，根据level_turbulence选择对应适应性评价->jude_turbulence。
		rms_turbulence 就是敏感度
		敏感度范围	等级	适应性评价
		≤500m	优	洋流扰动影响可忽略，适合隐蔽任务
		500∼1500m	良	需轨迹修正补偿，适合常规航行
		>1500m	差	需升级组合导航或动态洋流补偿算法
	*/
	if rms_turbulence <= 500 {
		fmt.Println("优")
		fmt.Println("洋流扰动影响可忽略，适合隐蔽任务")
	} else if rms_turbulence > 500 && rms_turbulence <= 1500 {
		fmt.Println("良")
		fmt.Println("需轨迹修正补偿，适合常规航行")
	} else {
		fmt.Println("差")
		fmt.Println("需升级组合导航或动态洋流补偿算法")
	}

	/*
		纬度误差参数的均方根： 72.38646282282345优
		洋流扰动影响可忽略，适合隐蔽任务
	*/
}
