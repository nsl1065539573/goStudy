package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func wordCount(s string)  {
	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}

func main() {
	// var m = map[string]Vertex{
	// 	"Bell Labs": Vertex{
	// 		40.68433, -74.39967,
	// 	},
	// 	"Google": Vertex{
	// 		37.42202, -122.08408,
	// 	},
	// }
	// m["FireFox"] = Vertex{11, 12}
	// fmt.Println(m)
	var s string = "hello world"
	wordCount(s)
}
