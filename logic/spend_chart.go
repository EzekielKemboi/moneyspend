package logic

import (
	"fmt"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/wcharczuk/go-chart/v2"
)

func generateCostCharts(spendData *spendData) {
	generateYearCostChart(spendData)
	generateMonthCostChart(spendData)
	// generateWeekCostChart(spendData)
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

func generateChartOld(name string, costs []float64, labels []string) {
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

// generateChart generates a bar chart using go-echarts library and saves it as a PNG
func generateChart(name string, costs []float64, labels []string) {
	// Create a new bar instance
	bar := charts.NewBar()

	// Set chart title and some global settings
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    name,
			Subtitle: "Cost Analysis",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Cost",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "Labels",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "3000px",
			Height: "2000px",
		}),
	)

	// Prepare data
	var barData []opts.BarData
	for _, cost := range costs {
		barData = append(barData, opts.BarData{Value: cost})
	}

	// Set data to the bar chart
	bar.SetXAxis(labels).
		AddSeries("Cost", barData).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Formatter: "{c}",
			}),
		)

	// Create HTML file
	htmlFile := name + ".html"
	f, err := os.Create(htmlFile)
	if err != nil {
		fmt.Printf("Create file %v err: %v", name, err)
		return
	}
	defer f.Close()

	// Render the chart to the HTML file
	err = bar.Render(f)
	if err != nil {
		fmt.Printf("Render chart %v err: %v", name, err)
		return
	}
}
