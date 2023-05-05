package logic

import (
	"fmt"
	"log"
	"os"

	"github.com/EzekielKemboi/moneyspend/bolt"
)

func ClassifyItems(spendData *spendData, classify bool) {
	notClassedKeyNum := 0
	allKeyNum := 0
	for _, item := range spendData.spendList {
		allKeyNum++
		if bolt.BoltGet(item.name) == "" {
			notClassedKeyNum++
		}
	}
	log.Printf("allKeyNum: %v,not classed key num: %v", allKeyNum, notClassedKeyNum)

	classSpendMap := make(map[string]int)
	for _, item := range spendData.spendList {
		if bolt.BoltGet(item.name) == "" {
			if !classify {
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
			bolt.BoltSet(item.name, class)
			log.Printf("item.name: %v,class: %v saved!", item.name, class)
		}
		class := bolt.BoltGet(item.name)
		classSpendMap[class] += item.price
	}

	total := 0
	for _, money := range classSpendMap {
		total += money
	}
	log.Printf("classSpendMap: %v,total: %v", classSpendMap, total)
}
