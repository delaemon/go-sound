package main

import (
	"flag"
	"time"

	"github.com/delaemon/go-sound"
	"fmt"
)

var filename = flag.String("file", "bazinga.wav", "play file")
var times = flag.Int("times", 1, "play times")

func main() {
	flag.Parse()
	fmt.Println(*times)
	for i := 0; i < *times; i++ {
		sound.Play(*filename)
		time.Sleep(1 * time.Second)
	}
}
