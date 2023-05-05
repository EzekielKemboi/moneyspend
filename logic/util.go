package logic

import (
	"log"
	"strconv"
	"strings"
)

func checkEqual(detailSum int, daySum int, date string) {
	if detailSum != daySum {
		log.Fatalf("day sum check failed!,date: %v", date)
	}
}

func stringToInt(str string) int {
	itemPrice, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("stringToInt atoi error: %v,str: %v", err, str)
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

func parseDate(date string) (int, int, int) {
	splitedDate := strings.Split(date, ".")
	if len(splitedDate) != 3 {
		log.Fatalf("date error: %v", date)
	}
	year := stringToInt(splitedDate[0])
	month := stringToInt(splitedDate[1])
	day := stringToInt(splitedDate[2])
	return year, month, day
}
