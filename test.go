package main

import "fmt"

func main() {

    //Default value
    var num0 int

    //Initialized with value
    var num1 int = 1

    //Without strict showing type
    var num2 = 20

    fmt.Println(num0, num1, num2)
    
    //Short initialization of variable
    num3 := 30

    num3 += 1
    fmt.Println("+=", num3)

    num3++
    fmt.Println("++", num3)

    //camelCase accepted standart
    userIndex := 10

    //underscore is not acceptable
    user_index := 10

    fmt.Println(userIndex, user_index)

    //Several var init
    var weight, height int = 10, 20

    //Assign values
    weight, height = 11, 21

    //Shortcut
    weight, age := 12, 22

    fmt.Println(weight, height, age)
}
