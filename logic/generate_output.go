package logic

import (
	"fmt"
	"io"
	"log"
	"os"
)

func GenerateOutput(spendData *spendData) {
	var err error
	err = os.MkdirAll("./output", 0755)
	if err != nil {
		log.Fatalf("os.Mkdir err: %v", err)
		return
	}
	err = os.Chdir("./output")
	if err != nil {
		log.Fatalf("os.Chdir err: %v", err)
		return
	}
	f, err := os.Create("output_data.txt")
	if err != nil {
		log.Fatalf("os.Create err: %v", err)
	}
	defer f.Close()

	for i, spendItem := range spendData.spendList {
		str := fmt.Sprintf("%v:%v", spendItem.name, spendItem.price)
		io.WriteString(f, str)
		if (i+1)%showItemsPerLine == 0 {
			io.WriteString(f, "\n")
		} else {
			io.WriteString(f, "     ")
		}
	}

	log.Printf("generate output_data.txt done!")

	generateCostCharts(spendData)

	log.Printf("generate MonthSpend.png done!")
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
