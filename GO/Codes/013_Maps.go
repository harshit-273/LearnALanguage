package main

import "fmt"

func main() {
	// initializing and declaration of map
	var map1 map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(map1) // accessing the full map

	map1["four"] = 4 // adding the new value to the map

	fmt.Println(map1["four"]) // accessing one of the item in map
	fmt.Println(map1)

	delete(map1, "four") // deleting the item from the map if it exists
	fmt.Println(map1)

	map2 := make(map[string]string) // another way of making a map(empty map)
	fmt.Println(map2)

	// checking if the key exists or not
	map2["a"] = "A"
	map2["b"] = "B"
	map2["c"] = "C"
	valueOfKey, ok := map2["d"]
	fmt.Printf("value of map2[\"d\"] - %q and it exists - %t\n", valueOfKey, ok)
	valueOfKey, ok = map2["b"]
	fmt.Printf("value of map2[\"b\"] - %q and it exists - %t\n", valueOfKey, ok)

	fmt.Println(len(map1), len(map2)) // for getting the number of elements of the map

	for key, val := range map1 {
		fmt.Printf("key - %q : value - %d\n", key, val)
	}

}
