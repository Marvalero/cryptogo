package beeper

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func Beep() {
	file, err := os.Open("mix-2.mp3")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	defer file.Close()
	sound, format, _ := mp3.Decode(file)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	playing := make(chan struct{})
	speaker.Play(beep.Seq(sound, beep.Callback(func() {
		close(playing)
	})))
	<-playing
}
