package logic

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
	pricei := s[i].price
	pricej := s[j].price
	if pricei < 0 {
		pricei = -pricei
	}
	if pricej < 0 {
		pricej = -pricej
	}
	return pricei < pricej
}
