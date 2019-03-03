package main

import "fmt"

func greetz(channel chan string) {
	channel <- "howdy, junior"
}

func printrnd(channel chan string, item string) {
	for i := 1; i <= 10; i++ {
		channel <- item
	}
}

func main() {
	channel := make(chan string)
	go greetz(channel)
	var text string = <-channel
	fmt.Println(text)

	go printrnd(channel, "a")
	go printrnd(channel, "b")

	for x := 1; x < 21; x++ {
		fmt.Println(<-channel)
	}

}
