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

func sortByPriceAsc(houses []house) []house {

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}

func sortByPriceDesc(houses []house) []house {

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}

func FarFromCenter(houses []house) []house {

	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].farFromCenter > result[j].farFromCenter
	})
	return result
}
func NearToCenter(houses []house) []house {
		result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		return result[i].farFromCenter < result[j].farFromCenter
	})
	return result
}

func MaxPrice(houses []house, limit int64) []house {
		result := make([]house, 0)
	for _, house := range houses {
		if house.price <= limit {
			result = append(result, house)
		}
	}
	return result
}
func SearchWithInCost(houses []house, startLimit int64, endLimit int64) []house {
	result := make([]house, 0)
	for _, house := range houses {
		if house.price >= startLimit && house.price <= endLimit {
			result = append(result, house)
		}
	}
	return result
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




