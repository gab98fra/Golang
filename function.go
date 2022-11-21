package main
import "fmt"

func main(){
	/*
		Area of a square: side * side
	*/
	var side float32
	var area float32
	fmt.Println("Give me the side of the square")
	fmt.Scan(&side)
	area=get_area(side)
	fmt.Println("area: ", area)
}

func get_area(side float32) (area_of_square float32){
	area_of_square = side*side
	return 
}
