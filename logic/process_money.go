package logic

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

const (
	startDate        = "2020.11.1"
	showItemsPerLine = 1
	exitCode         = 99
)

var classMap = map[int]string{1: "吃喝", 2: "住房", 3: "生活", 4: "医疗", 5: "出行", 8: "女人", 10: "家里", 11: "强制支出", exitCode: "退出"}

func ProcessMoneyCsv(fileName string) *spendData {
	f, err := os.Open("data/" + fileName)
	if err != nil {
		log.Fatalf("open csv err: %v", err)
	}
	defer f.Close()

	spendData := newSpendData()

	reader := csv.NewReader(f)
	reader.Read()

	for {
		line, _ := reader.Read()
		if len(line) == 0 {
			break
		}
		processOneDay(line, spendData)
	}

	spendData.genSpendItems()

	return spendData
}

func processOneDay(line []string, spendData *spendData) {
	date, detail, daySumStr := line[0], line[1], line[2]
	if date == "" || detail == "" || date == startDate {
		return
	}

	if daySumStr == "" {
		log.Fatalf("line total spend is null,date: %v", date)
	}

	daySum := stringToInt(daySumStr)

	if detail == "/" {
		if daySum != 0 {
			log.Fatalf("detail is /,but sum not 0,date: %v", date)
		}
		return
	}

	year, month, day := parseDate(date)
	spendData.setCostMap(year, month, day, daySum)

	var detailSum int
	items := strings.Split(detail, "，")
	for _, item := range items {
		detailSum += processItem(item, spendData.spendMap)
	}

	checkEqual(detailSum, daySum, date)
}

func processItem(item string, spendMap map[string]int) int {
	runeItem := []rune(item)
	i := len(runeItem) - 1
	for {
		if isNum(runeItem[i]) {
			decr(&i)
		} else {
			break
		}
	}

	var itemName string
	addMoney := false

	switch runeItem[i] {
	case '+':
		addMoney = true
		itemName = string(runeItem[:i])
	case '：':
		itemName = string(runeItem[:i])
	default:
		itemName = string(runeItem[:i+1])
	}

	itemPriceStr := string(runeItem[i+1:])
	if itemPriceStr == "" {
		log.Fatalf("%v no price", itemName)
	}
	itemPrice := stringToInt(itemPriceStr)
	if addMoney {
		itemPrice = -itemPrice
	}

	spendMap[itemName] += itemPrice

	return itemPrice
}
