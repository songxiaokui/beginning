package main

import (
	"fmt"
	"math/rand"
)

func calRollsErrorNumber() int64 {
	return int64(rand.Int())
}
func main() {
	// 提取rolls_error_ud【rolls_error 横摇误差】
	// 从数据库中晒横摇误差>5的个数 rolls_error_ud
	// TODO 借助clickhouse的数据筛选功能实现即可
	rolls_error_ud := calRollsErrorNumber()

	//找出 rolls_error_ud中所有>5的个数，个数为limit_ud。
	limit_ud := rolls_error_ud

	/*
			等级，根据limit_ud选择对应等级->level_ud。
			step4：适应性评价，根据limit_ud选择对应适应性评价->jude_ud。
			limit_ud对应敏感度范围，根据敏感度范围选取“等级”&“适应性评价”。

		limit_ud = 敏感度范围
		敏感度范围	等级	适应性评价
		≥1.0kHz	优	抗升沉能力极强，适合恶劣海况任务
		0.5∼1.0kHz	良	需动态阻尼补偿，适合常规航行
		<0.5kHz	差	需升级姿态解算算法或传感器带宽
	*/
	var (
		level_ud string
		jude_ud  string
	)
	switch {
	case limit_ud >= 1000:
		level_ud = "优"
		jude_ud = "抗升沉能力极强，适合恶劣海况任务"
	case limit_ud >= 500:
		level_ud = "良"
		jude_ud = "需动态阻尼补偿，适合常规航行"
	default:
		level_ud = "差"
		jude_ud = "需升级姿态解算算法或传感器带宽"
	}

	fmt.Printf("等级: %s\n", level_ud)
	fmt.Printf("适应性评价: %s\n", jude_ud)

}
