package main

import "fmt"

func testMap() {
	var a map[string]string = map[string]string{
		"key": "value",
	}
	//a := make(map[string]string, 10)
	a["abc"] = "efg"
	fmt.Println(a)
}

func testMap2() {
	a := make(map[string]map[string]string)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "a"
	a["key1"]["key3"] = "b"
	a["key1"]["key4"] = "c"
	fmt.Println(a)

	val, ok := a["key1"]
	if ok {
		fmt.Println(val, ok)
	} else {
		fmt.Println(val, ok)
	}
}

func testMap3() {
	a := make(map[string]map[string]string)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "a"
	a["key1"]["key3"] = "b"
	a["key1"]["key4"] = "c"
	a["key2"] = make(map[string]string)
	a["key2"]["key2"] = "e"
	a["key2"]["key3"] = "f"
	a["key2"]["key4"] = "g"
	fmt.Println(a)

	for k, v := range a {
		fmt.Println(k, v)
	}  //key1 map[key3:b key4:c key2:a]

	delete(a, "key1")
	fmt.Println(a)
}

//func testMap4() {
//	a := make([]map[int]int, 5)
//
//	a[0] = 10
//
//}

func main() {
	testMap()
	testMap2()
	testMap3()
}
