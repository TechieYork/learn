package main

import "fmt"

func testInitMap() {
	fmt.Println("====== Test init map ======")

	//Init empty map
	var mapEmpty map[string]string  //Now this is a nil map, so do remember using make to init map
	mapEmpty = make(map[string]string)
	fmt.Println(mapEmpty)

	mapEmpty1 := make(map[string]string)
	fmt.Println(mapEmpty1)

	//Init map with value
	mapWithValue := map[string]string{"1": "One", "2": "Two"}
	fmt.Println(mapWithValue)
}

func testSetMap() {
	fmt.Println("====== Test set to map ======")

	//Test set new key and value
	mapToSet := make(map[string]string)

	mapToSet["1"] = "One"
	mapToSet["2"] = "Two"
	mapToSet["3"] = "Three"
	mapToSet["4"] = "Four"
	mapToSet["5"] = "Five"

	fmt.Println("Set new key and value to map: ", mapToSet)

	//Test set old key with new value
	mapToSet ["1"] = "ONE"
	fmt.Println("After set old key with new value to map: ", mapToSet)
}

func testGetMap() {
	fmt.Println("====== Test get from map ======")

	//Test get exist key and value
	mapToGet := map[string]string{"1": "One", "2": "Two", "3": "Three", "4": "Four", "5": "Five"}
	fmt.Println("Map to get: ", mapToGet)

	value, ok := mapToGet["1"]

	if !ok {
		fmt.Println("Map key not found! key: 1")
	} else {
		fmt.Println("Map key found! key: 1, value: ", value)
	}

	//Test get non-exist key and value
	value, ok = mapToGet["10"]

	if !ok {
		fmt.Println("Map key not found! key: 10")
	} else {
		fmt.Println("Map key found! key: 10, value: ", value)
	}
}

func testDeleteMap() {
	fmt.Println("====== Test delete from map ======")

	//Test delete exist key and value
	mapToDel := map[string]string{"1": "One", "2": "Two", "3": "Three", "4": "Four", "5": "Five"}
	fmt.Println("Map to delete: ", mapToDel)

	delete(mapToDel, "1")
	fmt.Println("After delete 1 in map: ", mapToDel)

	//Test delete non-exist key, nothing happened because the key is not in the map
	delete(mapToDel, "10")
}

func testCopy() {
	fmt.Println("====== Test copy map ======")

	//Map is complicated data structure
	//When copy, please note that, you have to loop all the key and value to assigned to the new map

	//Test copy without loop copy
	mapToCopy := map[string]string{"1": "One", "2": "Two", "3": "Three", "4": "Four", "5": "Five"}
	mapToCopy1 := mapToCopy
	fmt.Println("Original map: ", mapToCopy)
	fmt.Println("New map from original: ", mapToCopy1)

	mapToCopy["1"] = "ONE"
	fmt.Println("After set mapToCopy[\"1\"] = \"ONE\", Original map: ", mapToCopy)
	fmt.Println("After set mapToCopy[\"1\"] = \"ONE\", New map: ", mapToCopy1)

	//Test copy with loop copy
	mapToCopy2 := map[string]string{"1": "One", "2": "Two", "3": "Three", "4": "Four", "5": "Five"}
	fmt.Println("Original map: ", mapToCopy2)

	mapToCopy3 := make(map[string]string)

	for k, v := range mapToCopy2 {
		mapToCopy3[k] = v
	}

	fmt.Println("New map from original: ", mapToCopy3)

	mapToCopy2["1"] = "ONE"
	fmt.Println("After set mapToCopy2[\"1\"] = \"ONE\", Original map: ", mapToCopy2)
	fmt.Println("After set mapToCopy2[\"1\"] = \"ONE\", New map: ", mapToCopy3)

}

func testLen() {
	fmt.Println("====== Test len ======")

	//Test len
	mapToTest := map[string]string{"1": "One", "2": "Two", "3": "Three", "4": "Four", "5": "Five"}
	fmt.Printf("Map: %v, Len:%v\r\n", mapToTest, len(mapToTest))
}

func main() {
	testInitMap()
	testSetMap()
	testGetMap()
	testDeleteMap()
	testCopy()
	testLen()
}
