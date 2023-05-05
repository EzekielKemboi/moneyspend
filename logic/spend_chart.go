package logic

import (
	"log"
	"os"

	"github.com/wcharczuk/go-chart/v2"
)

func generateCostCharts(spendData *spendData) {
	generateYearCostChart(spendData)
	generateMonthCostChart(spendData)
	generateWeekCostChart(spendData)
	// generateDayCostChart(spendData)
}

func generateYearCostChart(spendData *spendData) {
	yearCosts, yearCostLabels := spendData.getYearCosts()

	generateChart("year_spend", yearCosts, yearCostLabels)
}

func generateMonthCostChart(spendData *spendData) {
	monthCosts, monthCostLabels := spendData.getMonthCosts()

	generateChart("month_spend", monthCosts, monthCostLabels)
}

func generateWeekCostChart(spendData *spendData) {
	weekCosts, weekCostLabels := spendData.getWeekCosts()

	generateChart("week_spend", weekCosts, weekCostLabels)
}

func generateDayCostChart(spendData *spendData) {
	dayCosts, dayCostLabels := spendData.getDayCosts()

	generateChart("day_spend", dayCosts, dayCostLabels)
}

func generateChart(name string, costs []float64, labels []string) {
	var data []chart.Value
	for i, n := range costs {
		data = append(data,
			chart.Value{
				Value: n,
				Label: labels[i],
			},
		)
	}

	graph := chart.BarChart{
		Title:      name,
		Background: chart.Style{Padding: chart.Box{}},
		Width:      2096,
		Height:     1440,
		Bars:       data,
	}

	file, err := os.Create(name + ".png")
	if err != nil {
		log.Fatalf("Create pic %v err: %v", name, err)
	}
	defer file.Close()

	err = graph.Render(chart.PNG, file)
	if err != nil {
		log.Fatalf("Render pic %v err: %v", name, err)
	}
}
