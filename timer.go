package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"flag"

	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {

	h:=flag.Int("h",0,"Hour")
	m:=flag.Int("m",0,"Minute")
	s:=flag.Int("s",0,"Second")
	flag.Parse()
	durationFromArgument := (*h*60+*m)*60+*s
	d := time.Duration(time.Second * time.Duration(durationFromArgument))

	fmt.Printf("Timer is start at %d:%d:%d\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	t := time.NewTimer(d)
	defer t.Stop()

	<-t.C

	t.Stop()
	fmt.Printf("Timer is stop at %d:%d:%d\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	// for i := 0; i < 30; i++ {

		beepPlayer()
	
	// }

}

func getExePath() (exPath string) {
	//get application's path
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath = filepath.Dir(ex)
	return exPath
}

func beepPlayer() {
	
	file, err := os.Open(getExePath() +"/"+"beep.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	loop := beep.Loop(30, streamer)

	done := make(chan bool)

	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		fmt.Println("Done")
		done <- true
	})))
	<-done

}