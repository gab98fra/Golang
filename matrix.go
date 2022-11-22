package main
import "fmt"

func main(){
	//Matrix
	var data[2][3] int

	for x:=0; x<2; x++{
		for y:=0;y<3;y++{
			fmt.Println("Give me a integer number")
			fmt.Scan(&data[x][y])
		}
	}

	fmt.Println(data)
	
}