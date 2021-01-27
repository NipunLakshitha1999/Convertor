package main

import "fmt"

func main() {
	var number int
	fmt.Print("Input Your Number :")


	//get User input
	_,isErro := fmt.Scanf("%d",&number)

	//Aply condition - chek it is a string or number
	if isErro!= nil{
		fmt.Println("PlZ input Number")
	}else {
		number = number*180
	}


	//print Out put
	fmt.Println("your Value is:",number)
}
