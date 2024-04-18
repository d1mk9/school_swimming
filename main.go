package main

import (
	"fmt"
)

// Brunch Общая информация филиала
type Brunch struct {
	Name       string  // название филиала
	IncomeG    float64 // приход с групп
	IncomeMg   float64 // приход с мини-групп
	OutlayArG  float64 // расход аренды групп
	OutlayZpG  float64 // расход зарплаты тренерам групп
	OutlayArMg float64 // расход аренды мини-групп
	OutlayZpMg float64 // расход зарплаты тренерам мини-групп
}

// CalcIncome Метод для подсчета прихода филиала
func (b Brunch) CalcIncome() float64 {
	return b.IncomeG + b.IncomeMg
}

// CalcOutlay Метод для подсчета расходов филиала
func (b Brunch) CalcOutlay() float64 {
	return b.OutlayArG + b.OutlayArMg + b.OutlayZpG + b.OutlayZpMg
}

func main() {
	br1 := Brunch{
		Name:       "СТЦ",
		IncomeG:    20000,
		IncomeMg:   21000,
		OutlayArG:  22000,
		OutlayZpG:  10000,
		OutlayArMg: 5100,
		OutlayZpMg: 1200,
	}
	fmt.Println(br1)
}
