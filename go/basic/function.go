package main

import "fmt"

func calArea(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

func calCapital(country string) string {
	countryCapitalMap := map[string]string{"Việt Nam": "hanoi", "France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	for i := range countryCapitalMap {
		if i == country {
			return countryCapitalMap[i]
		}
	}
	return ""
}
func main() {
	countryCapitalMap := map[string]string{"Việt Nam": "hanoi", "France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}
	// var tên map[k]v
	//var tên = make(map[k]v)
	// add a key-value
	countryCapitalMap["đất nước x"] = "thủ đô x"
	w, h, area := calArea(10, 20)
	fmt.Println(w, h, area)
	var capital string = calCapital("Việt Nam")
	fmt.Println(capital)
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	// delete a key: delete(tên map, k)
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("Updated map---------------------------------")
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
	// truy cap 1 phan tu trong map
	value, found := countryCapitalMap["not found"]
	if found {
		fmt.Println(value)
	} else {
		fmt.Println("không tìm thấy")
	}
}
