package main

// import (
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/faiface/beep"
// 	"github.com/faiface/beep/mp3"
// 	"github.com/faiface/beep/speaker"
// )

// func create() *beep.StreamSeekCloser {
// 	f, err := os.Open("../Drunktesting/airhorn.mp3")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	streamer, format, err := mp3.Decode(f)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer streamer.Close()

// 	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
// 	return &streamer
// }

// func play(streamer beep.StreamSeekCloser) {
// 	defer streamer.Close()
// 	done := make(chan bool)
// 	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
// 		done <- true
// 	})))

// 	<-done
// }
