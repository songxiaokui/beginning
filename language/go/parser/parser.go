package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sbinet/npyio"
)

func GetResuIndex(path string) (int, error) {
	base := filepath.Base(path)           // resu_123.npy
	ext := filepath.Ext(base)             // .npy
	name := strings.TrimSuffix(base, ext) // resu_123

	if !strings.HasPrefix(name, "resu_") {
		return 0, fmt.Errorf("invalid filename: %s", base)
	}

	idStr := strings.TrimPrefix(name, "resu_")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

type Output struct {
	Step      int     `json:"step"`       // 步数索引
	Time      float64 `json:"time"`       // 时间秒数（或步计数）
	LonTrue   float64 `json:"lon_true"`   // 真值经度
	LatTrue   float64 `json:"lat_true"`   // 真值纬度
	PhiTrue   float64 `json:"phi_true"`   // 真值横摇
	ThetaTrue float64 `json:"theta_true"` // 真值纵摇
	PsiTrue   float64 `json:"psi_true"`   // 真值航向
	VnX       float64 `json:"v_n_x"`      // 真值北向速度
	VnY       float64 `json:"v_n_y"`      // 真值东向速度
	VnZ       float64 `json:"v_n_z"`      // 真值地向速度
	HTrue     float64 `json:"h_true"`     // 真值深度

	LonCal   float64 `json:"lon_cal"`   // 计算经度
	LatCal   float64 `json:"lat_cal"`   // 计算纬度
	HCal     float64 `json:"h_cal"`     // 计算深度
	VCalX    float64 `json:"v_cal_x"`   // 计算北向速度
	VCalY    float64 `json:"v_cal_y"`   // 计算东向速度
	VCalZ    float64 `json:"v_cal_z"`   // 计算地向速度
	PhiCal   float64 `json:"phi_cal"`   // 计算横摇
	ThetaCal float64 `json:"theta_cal"` // 计算纵摇
	PsiCal   float64 `json:"psi_cal"`   // 计算航向

	Rudder float64 `json:"rudder"` // 舵角（可选）
	Speed  float64 `json:"speed"`  // 船速（可选）
}

const Stride = 20
const RowsPerHour = 3600

func ParseResu(path string) ([]Output, error) {
	f, _ := os.Open(path)
	defer f.Close()

	// 读python的二进制文件npy
	var data []float64
	if err := npyio.Read(f, &data); err != nil {
		return nil, err
	}

	fileIndex, err := GetResuIndex(path)
	if err != nil {
		return nil, err
	}

	// 判读数据是否合法
	if len(data)%20 != 0 {
		return nil, fmt.Errorf("invalid data length: %d", len(data))
	}

	// 对数据进行分组处理
	// 每20个一组
	rows := len(data) / Stride
	result := make([]Output, rows)
	for row := 0; row < rows; row++ {
		out := buildOutput(data, row, fileIndex)

		// 示例：只打印最后一条
		if row == rows-1 {
			fmt.Printf("last output:\n%+v\n", out)
		}

		result[row] = out
	}

	return result, nil
}

func buildOutput(data []float64, rowIndex int, fileIndex int) Output {
	base := rowIndex * Stride

	step := (fileIndex-1)*RowsPerHour + rowIndex + 1

	return Output{
		Step: step,
		Time: float64(step),

		LonTrue:   data[base+0],
		LatTrue:   data[base+1],
		PhiTrue:   data[base+2],
		ThetaTrue: data[base+3],
		PsiTrue:   data[base+4],

		VnX: data[base+5],
		VnY: data[base+6],
		VnZ: data[base+7],

		HTrue: data[base+8],

		LonCal: data[base+9],
		LatCal: data[base+10],
		HCal:   data[base+11],

		VCalX: data[base+12],
		VCalY: data[base+13],
		VCalZ: data[base+14],

		PhiCal:   data[base+15],
		ThetaCal: data[base+16],
		PsiCal:   data[base+17],

		Rudder: data[base+18],
		Speed:  data[base+19],
	}
}

func main() {
	f, _ := os.Open("resu_1.npy")
	defer f.Close()

	var data []float64
	if err := npyio.Read(f, &data); err != nil {
		panic(err)
	}
	fmt.Println(len(data)%20 == 00)
	fmt.Println(len(data))
	for _, v := range data[40:60] {
		fmt.Println(v)
	}

	res, err := ParseResu("resu_1.npy")
	if err != nil {
		panic(err)
	}
	fmt.Println(res[2])
}
