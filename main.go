package main 

import "fmt"



func myFunc(ch chan int, age int) {
	ch <- age 
}


func myFunc2(ch chan string, name string){
	ch <- name
}



func main() {
	ch1:= make(chan int)
	ch2:= make(chan string) 
     
	age:= 21
	name:= "munashe"
	go myFunc(ch1, age)
	go myFunc2(ch2, name)

	select {
	case server1:= <-ch1: 
	fmt.Println(server1)
	case server2:= <-ch2: 
	fmt.Println(server2)
	default: 
	panic("nothing found")
	}
}