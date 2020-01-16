package main

import "fmt"

var houses = []house{
	{
		id:            1,
		name:          "дом 4 этажном , 2 этаж, И. Сомони",
		price:         500000,
		discription:   "Продаю 4-квартиру в 5-комнатной квартире со всеми условиями",
		farFromCenter: 10,
		district:      "Сино",
	},
	{
		id:            2,
		name:          "Дом в 5 этажном , 3 этаж, И. Сомони",
		price:         600000,
		discription:   "Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам! ",
		farFromCenter: 15,
		district:      "Сомони",
	},
}

// Ctrl + Alt + L - выровнять код
func ExampleSortByPriceAsc() {
	ascSorted := sortByPriceAsc(houses)
	fmt.Println(ascSorted)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино} {2 Дом в 5 этажном , 3 этаж, И. Сомони 600000 Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам!  15 Сомони}]
}
func ExampleSortByPriceDesc() {
	descSorted := sortByPriceDesc(houses)
	fmt.Println(descSorted)
	// Output: [{2 Дом в 5 этажном , 3 этаж, И. Сомони 600000 Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам!  15 Сомони} {1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино}]
}

func ExampleFarFromCenter() {
	farFrom := FarFromCenter(houses)
	fmt.Println(farFrom)
	// Output: [{2 Дом в 5 этажном , 3 этаж, И. Сомони 600000 Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам!  15 Сомони} {1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино}]
}
func ExampleNearToCenter() {
	nearTo := NearToCenter(houses)
	fmt.Println(nearTo)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино} {2 Дом в 5 этажном , 3 этаж, И. Сомони 600000 Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам!  15 Сомони}]
}
func ExampleMaxPriceNo() {
	mPrice := MaxPrice(houses, 30)
	fmt.Println(mPrice)
	// Output: []
}
func ExampleMaxPriceOne() {
	mPrice := MaxPrice(houses, 550000)
	fmt.Println(mPrice)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино}]
}
func ExampleMaxPriceTwo() {
	mPrice := MaxPrice(houses, 650000)
	fmt.Println(mPrice)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино} {2 Дом в 5 этажном , 3 этаж, И. Сомони 600000 Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам!  15 Сомони}]
}

func ExampleSearchWithInCostNoR() {
	inCost := SearchWithInCost (houses, 300_000, 400_000)
	fmt.Println(inCost)
	// Output: []
}
func ExampleSearchWithInCostOneR() {
	inCost := SearchWithInCost(houses, 400_000, 550_000)
	fmt.Println(inCost)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино}]
}
func ExampleSearchWithInCostWtoR() {
	inCost := SearchWithInCost(houses, 400_000, 800_000)
	fmt.Println(inCost)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино} {2 Дом в 5 этажном , 3 этаж, И. Сомони 600000 Продаю 1-квартиру в  со всеми условиями.Только девушкам и женщинам!  15 Сомони}]
}
func ExampleSearchSearchByArea() {
	searchDis := SearchByArea(houses, "Сино")
	fmt.Println(searchDis)
	// Output: [{1 дом 4 этажном , 2 этаж, И. Сомони 500000 Продаю 4-квартиру в 5-комнатной квартире со всеми условиями 10 Сино}]
}
func ExampleSearchSearchByArea_NoResults() {
	searchDis := SearchByArea(houses, "Шохмансур")
	fmt.Println(searchDis)
	// Output: []
}

//SearchByAreas(houses, []string{"Шохмансур", "Сино"})