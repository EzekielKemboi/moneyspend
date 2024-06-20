package main

import (
	"fmt"
)

func main() {
	var x, m, n float64
	var z int

	x = 11 // 每年基准消费
	m = 3  // 前些年消费减少
	n = 0  // 后些年消费增加
	z = 10 // 从第几年开始消费增加

	A := 50.0 // 初始资产
	S := 45.0 // 年薪
	i := 0.02 // 通货膨胀率
	r := 0.08 // 理财收益率

	startAge := 25   // 开始年龄
	endAge := 90     // 结束年龄
	workYear := 3    // 工作年限
	inflation := 1.0 // 通货膨胀

	for t := startAge; t < endAge; t++ {
		var Ct, It float64

		// 第几年
		numYear := t - startAge + 1

		// 计算消费
		if numYear <= z {
			Ct = (x - m) * inflation
		} else {
			Ct = (x + n) * inflation
		}

		// 计算收入
		if numYear <= workYear {
			It = S
		} else {
			It = 0
		}

		// 更新资产
		A = (A + It - Ct) * (1 + r)

		// 更新通货膨胀率
		inflation = inflation * (1 + i)

		// fmt.Printf("第 %d 年,%v 岁：消费 %.2f 万元，收入 %.2f 万元，资产 %.2f 万元\n", numYear, startAge+numYear-1, Ct, It, A)
	}

	fmt.Printf("在90岁时的资产为：%.2f 万元, 去通胀后：%.2f 万元\n", A, A/inflation)
}
