package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	version            = "1.0.2"
	outputDataFileName = "output_data.txt"
	showItemsPerLine   = 1
	exitCode           = 99
)

var spendMap map[string]int
var spendList spendMapItems
var classMap = map[int]string{1: "吃喝", 2: "生活", 3: "医疗", 4: "出行", 8: "爬虫", 10: "家里", 11: "强制支出", 99: "退出"}

var isVersion = flag.Bool("v", false, "show the version.")
var reClass = flag.Bool("r", false, "reclass all the items.")
var passClass = flag.Bool("p", false, "pass the classify procedure.")
var moneyCsvFileName = flag.String("c", "冯敬翔的记账本 - Sheet1.csv", "the name of money csv file.")

func main() {
	preProcess()

	processMoneyCsv()

	formatPrint()

	classifyItems()
}

func classifyItems() {
	notClassedKeyNum := 0
	for _, item := range spendList {
		if boltGet(item.name) == "" {
			notClassedKeyNum++
		}
	}
	log.Printf("not classed key num: %v", notClassedKeyNum)

	classSpendMap := make(map[string]int)
	for _, item := range spendList {
		if boltGet(item.name) == "" {
			if *passClass {
				continue
			}
			log.Printf("classMap: %v", classMap)
			var classNum int
			fmt.Printf("classify your item: %v, money: %v---->  ", item.name, item.price)
			fmt.Scanf("%d", &classNum)
			if classNum == exitCode {
				os.Exit(0)
			}
			class, ok := classMap[classNum]
			if !ok {
				log.Fatalf("the class: %v not supported", classNum)
			}
			boltSet(item.name, class)
			log.Printf("item.name: %v,class: %v saved!", item.name, class)
		}
		class := boltGet(item.name)
		classSpendMap[class] += item.price
	}

	total := 0
	for _, money := range classSpendMap {
		total += money
	}
	log.Printf("classSpendMap: %v,total: %v", classSpendMap, total)
}

func preProcess() {
	flag.Parse()
	if *isVersion {
		log.Printf("version: %v", version)
		os.Exit(0)
	}

	spendMap = make(map[string]int)

	initBolt()
}

func formatPrint() {
	spendList = generateSpendItems(spendMap)

	sort.Sort(sort.Reverse(spendList))

	f, err := os.Create(outputDataFileName)
	if err != nil {
		log.Fatalf("open outputdata err: %v", err)
	}
	defer f.Close()

	for i, spendItem := range spendList {
		str := fmt.Sprintf("%v:%v", spendItem.name, spendItem.price)
		io.WriteString(f, str)
		if (i+1)%showItemsPerLine == 0 {
			io.WriteString(f, "\n")
		} else {
			io.WriteString(f, "     ")
		}
	}

	log.Printf("generate output_data.txt done!")
}

func generateSpendItems(spendMap map[string]int) spendMapItems {
	var spendItems spendMapItems
	for key, val := range spendMap {
		spendItems = append(spendItems, &spendMapItem{
			name:  key,
			price: val,
		})
	}
	return spendItems
}

func processMoneyCsv() {
	f, err := os.Open(*moneyCsvFileName)
	if err != nil {
		log.Fatalf("open csv err: %v", err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Read()

	for {
		line, _ := reader.Read()
		if len(line) == 0 {
			break
		}
		processOneDay(line)
	}
}

func processOneDay(line []string) {
	date, detail, daySumStr := line[0], line[1], line[2]
	if date == "" || detail == "" || detail == "/" || date == "2020.11.1" {
		return
	}
	if daySumStr == "" {
		log.Fatalf("line total spend is null,date: %v", date)
	}

	daySum := stringToInt(daySumStr)

	var detailSum int
	items := strings.Split(detail, "，")
	for _, item := range items {
		detailSum += processItem(item, spendMap)
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

	itemPrice := stringToInt(string(runeItem[i+1:]))
	if addMoney {
		itemPrice = -itemPrice
	}

	spendMap[itemName] += itemPrice

	return itemPrice
}

func checkEqual(detailSum int, daySum int, date string) {
	if detailSum != daySum {
		log.Fatalf("day sum check failed!,date: %v", date)
	}
}

func stringToInt(str string) int {
	itemPrice, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("item price atoi error: %v,str: %v", err, str)
	}
	return itemPrice
}

func decr(i *int) {
	if *i == 0 {
		log.Fatalf("only number in item!")
	}
	(*i)--
}

func isNum(c int32) bool {
	return c >= '0' && c <= '9'
}
