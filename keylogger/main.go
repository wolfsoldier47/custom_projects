package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func keyCloser() {
	_ = keyboard.Close()
}

func KeyLoggerRunner(keysEvents <-chan keyboard.KeyEvent) {
	for {
		event := <-keysEvents
		if event.Err != nil {
			fmt.Println(event.Err)
			panic(event.Err)
		}
		fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)
		if event.Key == keyboard.KeyEsc {
			break
		}
	}
}

func main() {

	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer keyCloser()

	fmt.Println("Starting key checker")
	fmt.Println("Press ESC to quit")
	KeyLoggerRunner(keysEvents)

}
