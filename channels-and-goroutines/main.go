package main

// import "fmt"

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("john")
	println(id)
	ch := make(chan *Message, 2)

	wh := &sync.WaitGroup{}

	wh.Add(2)
	go getUserChats(id, ch, wh)
	go getUserFriends(id, ch, wh)
	wh.Wait()
	close(ch)

	log.Println(time.Since(now))
	for msg := range ch {
		fmt.Println(msg)
	}
	// fmt.Println(chats)
	// fmt.Println(friends)
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	ch <- &Message{
		chats: []string{
			"john", "jane", "joe", "bob", "james",
		},
	}
	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s-2", name)
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	ch <- &Message{
		friends: []string{
			"john", "jane", "joe",
		},
	}
	wg.Done()
}

// func main() {
// 	userID := 44
// 	anotheruserID := &userID

// 	println(anotheruserID)
// 	var age *int
// 	println(age)

// 	age = &userID
// 	println(age)
// 	println(*age)

// 	age = new(int)
// 	update(age, 42)

// 	newUpdate(*age, 69)
// 	fmt.Println(*age)
// 	fmt.Println(age)

// }
// func update(val *int, to int) {
// 	*val = to
// }

// func newUpdate(vv int, to int) {
// 	vv = to
// }
