package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math"
)

func main() {
	// Sg_gravity、level_gravity、juge_gravity
	// time_list 时间(time)
	// lon_error_list 重力经度误差(lon_error)
	// slope_lon 重力经度误差增长斜率
	time_list := []float64{29, 97, 86, 20, 11, 40, 58, 99, 62, 95}
	lon_error_list := []float64{10, 79, 77, 97, 91, 70, 70, 89, 19, 66}

	// 执行线性回归 重力经度误差增长斜率
	_, slope_lon := stat.LinearRegression(time_list, lon_error_list, nil, false)
	fmt.Printf("重力经度误差增长斜率: %v\n", slope_lon)

	// 计算重力纬度误差增长斜率
	// slope_lat 重力纬度误差增长斜率
	lat_error_list := []float64{10, 79, 77, 97, 91, 70, 70, 89, 19, 66}

	// 执行线性回归 重力纬度误差增长斜率
	_, slope_lat := stat.LinearRegression(time_list, lat_error_list, nil, false)
	fmt.Printf("重力纬度误差增长斜率: %v\n", slope_lat)

	// 计算 (slope_lon + slope_lat) / 2->average_slope
	average_slope := (slope_lon + slope_lat) / 2
	fmt.Printf("平均斜率: %v\n", average_slope)

	// 提取重力异常模值->gravity_magnitude
	// 重力异常为[3]float64{0.05, 0.05, 0.05}
	// 平方和开根号使用math.Sqrt()
	Gravity := []float64{3, 4, 5}
	gravity_magnitude := math.Sqrt(Gravity[0]*Gravity[0] + Gravity[1]*Gravity[1] + Gravity[2]*Gravity[2])
	fmt.Printf("重力异常模值: %v\n", gravity_magnitude)

	// 计算Sg_gravity->Sg_gravity = average_slope÷gravity_magnitude。gravity_magnitud = "gravity"的模
	Sg_gravity := average_slope / gravity_magnitude
	fmt.Printf("Sg_gravity: %v\n", Sg_gravity)

	// 根据 Sg_gravity 评估等级和适应性
	// Sg_gravity 是敏感度
	/*
		敏感度范围	等级	适应性评价
		≤0.1m/s²/mGal	优	重力扰动影响可忽略，适合高精度任务
		0.1∼0.3m/s²/mGal	良	需辅助重力补偿，适合常规任务
		>0.3m/s²/mGal	差	需升级重力建模或硬件抗扰能力
	*/
	if Sg_gravity <= 0.1 {
		fmt.Println("优")
		fmt.Println("重力扰动影响可忽略，适合高精度任务")
	} else if Sg_gravity > 0.1 && Sg_gravity <= 0.3 {
		fmt.Println("良")
		fmt.Println("需辅助重力补偿，适合常规任务")
	} else {
		fmt.Println("差")
		fmt.Println("需升级重力建模或硬件抗扰能力")
	}

}
