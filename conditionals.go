package main
import "fmt"

func main(){
	var option int
	//Another way to declare variables
	msg :="selected option: "
	fmt.Println("please, select an option")
	fmt.Scan(&option)
	if option==1{
		fmt.Println(msg, "1")
	}else{
		if option < 1 {
			fmt.Println(msg, "less than 1")
		}else{
			fmt.Println("Another option")
		}
	}
}