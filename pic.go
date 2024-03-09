package main

// import "golang.org/x/tour/pic"
//import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var res [][]uint8

	for i := 0; i < dy; i++ {
		var tem []uint8
		for j := 0; j < dx; j++ {
			tem = append(tem, uint8(i^2+j^2))
		}
		res = append(res, tem)
	}
	return res
}

// func main() {
// 	//fmt.Println(Pic(10,10))
// 	pic.Show(Pic)
// }
