package main
import "fmt"

func main() {
    //Variables
    var num1 int
    var num2, param float32
    num1=232
    num2=23.2
    /*
        Param
    */
    param=float32(num1) * num2
    fmt.Println("param: ", param)

    //Execute function: get_salary
    get_salary(param)
}


func get_salary(param float32){
    var hours_worked int
    var salary float32
    fmt.Println("Give me hours worked")
    fmt.Scan(&hours_worked)
    salary=float32(hours_worked)*param
    fmt.Println("Salary: ", salary)

}

