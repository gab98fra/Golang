package main
import "fmt"

func main(){
	var option int
	fmt.Println("selected option")
	fmt.Scan(&option)
	if option==1{
		fmt.Println("selected option: 1 ")
	}else{
		if option < 1 {
			fmt.Println("selected option: less than 1")
		}
	}
}