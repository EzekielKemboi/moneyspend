package logic

import (
	"log"
	"sort"
	"time"

	"github.com/spf13/cast"
)

type spendData struct {
	spendMap     map[string]int
	dayCostMap   map[int]map[int]map[int]int
	monthCostMap map[int]map[int]int
	yearCostMap  map[int]int
	spendList    spendMapItems
}

func newSpendData() *spendData {
	return &spendData{
		spendMap: map[string]int{},
	}
}

func (spendData *spendData) genSpendItems() {
	spendData.spendList = generateSpendItems(spendData.spendMap)

	sort.Sort(sort.Reverse(spendData.spendList))
}

func (spendData *spendData) setCostMap(year int, month int, day int, num int) {
	if spendData.dayCostMap == nil {
		spendData.dayCostMap = make(map[int]map[int]map[int]int)
	}
	if spendData.dayCostMap[year] == nil {
		spendData.dayCostMap[year] = make(map[int]map[int]int)
	}
	if spendData.dayCostMap[year][month] == nil {
		spendData.dayCostMap[year][month] = make(map[int]int)
	}
	if spendData.monthCostMap == nil {
		spendData.monthCostMap = make(map[int]map[int]int)
	}
	if spendData.monthCostMap[year] == nil {
		spendData.monthCostMap[year] = make(map[int]int)
	}
	if spendData.yearCostMap == nil {
		spendData.yearCostMap = make(map[int]int)
	}

	spendData.dayCostMap[year][month][day] += num
	spendData.monthCostMap[year][month] += num
	spendData.yearCostMap[year] += num
}

func (spendData *spendData) getYearCost(year int) int {
	if spendData.yearCostMap == nil {
		log.Printf("yearCostMap is nil %v", year)
		return 0
	}
	return spendData.yearCostMap[year]
}

func (spendData *spendData) getMonthCost(year int, month int) int {
	if spendData.monthCostMap == nil || spendData.monthCostMap[year] == nil {
		log.Printf("monthCostMap is nil %v.%v", year, month)
		return 0
	}
	return spendData.monthCostMap[year][month]
}

func (spendData *spendData) getDayCost(year int, month int, day int) int {
	if spendData.dayCostMap == nil || spendData.dayCostMap[year] == nil || spendData.dayCostMap[year][month] == nil {
		log.Printf("dayCostMap is nil %v.%v.%v", year, month, day)
		return 0
	}
	return spendData.dayCostMap[year][month][day]
}

func (spendData *spendData) getYearCosts() ([]float64, []string) {
	start := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)
	end := time.Now()

	var yearCosts []float64
	var yearCostLabels []string

	current := start
	for current.Before(end) || current.Equal(end) {
		cost := spendData.getYearCost(current.Year())
		yearCosts = append(yearCosts, cast.ToFloat64(cost))
		yearCostLabels = append(yearCostLabels,
			cast.ToString(current.Year()))
		current = current.AddDate(1, 0, 0)
	}

	return yearCosts, yearCostLabels
}

func (spendData *spendData) getMonthCosts() ([]float64, []string) {
	start := time.Date(2020, time.November, 1, 0, 0, 0, 0, time.Local)
	end := time.Now()

	var monthCosts []float64
	var monthCostLabels []string

	current := start
	for current.Before(end) || current.Equal(end) {
		cost := spendData.getMonthCost(current.Year(), int(current.Month()))
		monthCosts = append(monthCosts, cast.ToFloat64(cost))
		monthCostLabels = append(monthCostLabels,
			cast.ToString(current.Year())+"."+cast.ToString(int(current.Month())))
		current = current.AddDate(0, 1, 0)
	}

	return monthCosts, monthCostLabels
}

func (spendData *spendData) getWeekCosts() ([]float64, []string) {
	start := time.Date(2020, time.November, 2, 0, 0, 0, 0, time.Local)
	end := time.Now()

	var weekCosts []float64
	var weekCostLabels []string

	current := start
	for current.Before(end) || current.Equal(end) {
		weekCostLabels = append(weekCostLabels,
			cast.ToString(current.Year())+"."+cast.ToString(int(current.Month()))+"."+cast.ToString(current.Day()))

		weekCost := 0
		for i := 1; i <= 50; i++ {
			cost := spendData.getDayCost(current.Year(), int(current.Month()), current.Day())
			weekCost += cost
			current = current.AddDate(0, 0, 1)
		}

		weekCosts = append(weekCosts, cast.ToFloat64(weekCost))
	}

	return weekCosts, weekCostLabels
}

func (spendData *spendData) getDayCosts() ([]float64, []string) {
	start := time.Date(2020, time.November, 1, 0, 0, 0, 0, time.Local)
	end := time.Now()

	var dayCosts []float64
	var dayCostLabels []string

	current := start
	for current.Before(end) || current.Equal(end) {
		cost := spendData.getDayCost(current.Year(), int(current.Month()), current.Day())
		dayCosts = append(dayCosts, cast.ToFloat64(cost))
		dayCostLabels = append(dayCostLabels,
			cast.ToString(current.Year())+"."+cast.ToString(int(current.Month()))+"."+cast.ToString(current.Day()))
		current = current.AddDate(0, 0, 1)
	}

	return dayCosts, dayCostLabels
}
