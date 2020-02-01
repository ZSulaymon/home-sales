package main

import (
	"sort"
)

type house struct {
	id            int64
	name          string
	price         int64
	discription   string
	farFromCenter int
	district      string
}


func SortBy(houses []house, less func(a, b house) bool) []house{
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return less(result[i], result[j])
	})
	return result
}


func sortByPriceAsc(houses []house) []house {
	return SortBy(houses, func(a, b house) bool{
		return a.price < b.price
	})

}

func sortByPriceDesc(houses []house) []house {
	return SortBy(houses, func(a, b house) bool {
		return a.price > b.price
	})
}

func FarFromCenterAsc(houses []house) []house {
	return SortBy(houses, func(a, b house) bool {
		return a.farFromCenter > b.farFromCenter

		})
}


func FarFromCenterDesc(houses []house) []house {
	return SortBy(houses, func(a, b house) bool {
		return a.farFromCenter < b.farFromCenter

	})
}
func FilterBy(houses []house, less func(a house) bool) []house{
	result := make([]house, 0)
	for _, house := range houses{
		if less(house){
			result = append(result, house)
		}
	}
	return result
}



func SearchByMaxPrice(houses []house, MaxPrice int64) []house {
	return FilterBy(houses, func(a house) bool {
		return a.price <= MaxPrice

	})
}

func SearchWithInCost(houses []house, startLimit int64, endLimit int64) []house {
	return FilterBy(houses, func (a  house) bool {
	return a.price >= startLimit && a.price <= endLimit
	})
}

func SearchByArea(houses []house, districts string ) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.district == districts {
			result = append(result, house)
		}
	}
	return result
}

func main() {

}