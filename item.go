package main

type spendMapItem struct {
	name  string
	price int
}

type spendMapItems []*spendMapItem

func (s spendMapItems) Len() int {
	return len(s)
}

func (s spendMapItems) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s spendMapItems) Less(i, j int) bool {
	return s[i].price < s[j].price
}
