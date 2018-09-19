package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHello = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		// data di masukan ke channel messages
		messages <- data
	}

	go sayHello("Teten")
	go sayHello("Yasna")
	go sayHello("Abbu")

	// data dari channel messages di masukan ke var message1
	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}
