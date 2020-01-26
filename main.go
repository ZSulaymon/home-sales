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
	// make - чтобы сразу указать размер
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		// TODO:
		// if tariffs[i].pricePeriod == tariffs[j].pricePeriod ?
		return result[i].price < result[j].price
	})
	return result
}

func sortByPriceDesc(houses []house) []house {
	// make - чтобы сразу указать размер
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		// TODO:
		// if tariffs[i].pricePeriod == tariffs[j].pricePeriod ?
		return result[i].price > result[j].price
	})
	return result
}

func FarFromCenter(houses []house) []house {
	// make - чтобы сразу указать размер
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		// TODO:
		// if tariffs[i].pricePeriod == tariffs[j].pricePeriod ?
		return result[i].farFromCenter > result[j].farFromCenter
	})
	return result
}
func NearToCenter(houses []house) []house {
	// make - чтобы сразу указать размер
	result := make([]house, len(houses))
	copy(result, houses)
	sort.Slice(result, func(i, j int) bool {
		// TODO:
		// if tariffs[i].pricePeriod == tariffs[j].pricePeriod ?
		return result[i].farFromCenter < result[j].farFromCenter
	})
	return result
}

func MaxPrice(houses []house, limit int64) []house {
	// make - чтобы сразу указать размер
	result := make([]house, 0)
	for _, house := range houses {
		if house.price <= limit {
			// append
			result = append(result, house)
		}
	}
	return result
}
func SearchWithInCost(houses []house, startLimit int64, endLimit int64) []house {
	// make - чтобы сразу указать размер
	result := make([]house, 0)
	for _, house := range houses {
		if house.price >= startLimit && house.price <= endLimit {
			// append
			result = append(result, house)
		}
	}
	return result
}
func SearchByArea(houses []house, districts string ) []house {
	// make - чтобы сразу указать размер
	result := make([]house, 0)
	for _, house := range houses {
		if house.district == districts {
			// append
			result = append(result, house)
		}
	}
	return result
}

func main() {

}




