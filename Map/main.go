package main

import "fmt"

func main() {
	//color := map[string]string{
	//	"a": "1",
	//	"b": "2",
	//}
	color := make(map[string]int)
	color["a"] = 1
	color["b"] = 2
	delete(color, "a")
	//fmt.Println(color)
	for key, val := range color {
		fmt.Printf("%v: %v\n", key, val)
	}
}
