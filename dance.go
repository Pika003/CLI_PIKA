package main

import (
	"fmt"
	"time"
)

var pikaFrames = []string{
	`
(\__/)
(•ㅅ•)
/ 　 づ   ♪
`,
	`
(\__/)
( •ㅅ•)
づ　　/   ♫
`,
	`
(\__/)
(•ㅅ•)ノ
/ 　     ♪♪
`,
	`
(\__/)
(ノ•ㅅ•)ノ
/   　   ♫
`,
	`
(\__/)
( •ㅅ•)づ
/   　   ♪
`,
	`
(\__/)
(づ｡•ㅅ•｡)づ
/   　   ♫♪
`,
}

func PikaDance() {
	// fmt.Println("Dancing Pikachu!")
	// fmt.Println("Get ready to dance...")
	// time.Sleep(1 * time.Second)

	soundDone := make(chan bool)

	// Start playing sound immediately
	go func() {
		playSound("pika_pika.mp3")
		soundDone <- true
	}()

	// Reduced from 53 to 30 frames for shorter duration
	totalFrames := 24

	for i := 0; i < totalFrames; i++ {
		fmt.Print("\033[H\033[2J")
		fmt.Println(pikaFrames[i%len(pikaFrames)])

		// Slightly reduced timing
		if i%4 == 0 {
			time.Sleep(250 * time.Millisecond)
		} else {
			time.Sleep(150 * time.Millisecond)
		}
	}

	fmt.Println("Pika pika! (That was fun!)")
	fmt.Println("⚡️ Thanks for dancing with me! ⚡️")

	// Wait for sound to finish before exiting
	<-soundDone
}
