package main

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cast"
	"github.com/wcharczuk/go-chart/v2"
)

var dateMap map[int]map[int]map[int]int

func setDateMap(year int, month int, day int, num int) {
	if dateMap == nil {
		dateMap = make(map[int]map[int]map[int]int)
	}
	if dateMap[year] == nil {
		dateMap[year] = make(map[int]map[int]int)
	}
	if dateMap[year][month] == nil {
		dateMap[year][month] = make(map[int]int)
	}
	dateMap[year][month][day] += num
}

func getMonthCostDateMap(year int, month int) int {
	if dateMap == nil || dateMap[year] == nil {
		log.Printf("dateMap is nil %v.%v", year, month)
	}
	monthCost := 0
	for _, val := range dateMap[year][month] {
		monthCost += val
	}
	return monthCost
}

func getDayCostDateMap(year int, month int, day int) int {
	if dateMap == nil || dateMap[year] == nil || dateMap[year][month] == nil {
		log.Printf("dateMap is nil %v.%v.%v", year, month, day)
	}
	return dateMap[year][month][day]
}

func generateMonthCostPic() {
	start := time.Date(2020, time.November, 1, 0, 0, 0, 0, time.Local)
	end := time.Now()

	var monthCosts []int
	var monthCostLabels []string

	current := start
	for current.Before(end) || current.Equal(end) {
		cost := getMonthCostDateMap(current.Year(), int(current.Month()))
		monthCosts = append(monthCosts, cost)
		monthCostLabels = append(monthCostLabels, cast.ToString(current.Year())+"."+cast.ToString(int(current.Month())))
		current = current.AddDate(0, 1, 0)
	}

	var data []chart.Value
	for i, n := range monthCosts {
		data = append(data,
			chart.Value{
				Value: float64(n),
				Label: monthCostLabels[i],
			},
		)
	}

	graph := chart.BarChart{
		Title:      "MonthSpend",
		Background: chart.Style{Padding: chart.Box{}},
		Width:      2096,
		Height:     1440,
		Bars:       data,
	}

	file, err := os.Create("MonthSpend.png")
	if err != nil {
		log.Fatalf("Create pic err: %v", err)
	}
	defer file.Close()

	err = graph.Render(chart.PNG, file)
	if err != nil {
		log.Fatalf("Render pic err: %v", err)
	}
}
