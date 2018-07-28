package exchange_calculator

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Exchange struct {
	ReadChan     chan float64
	WriteChan    chan float64
	currentValue float64
	upLimit      float64
	downLimit    float64
	Currency     string
}

func (exc Exchange) beepIfLimitReached() {

	for {
		value := <-exc.ReadChan
		if value >= exc.upLimit {
			fmt.Println("---- UP", exc.Currency, " >= ", exc.upLimit)
			playSound()
		} else if value <= exc.downLimit {
			fmt.Println("---- DOWN", exc.Currency, " <= ", exc.downLimit)
			playSound()
		}
		time.Sleep(10 * time.Second)
	}
}

func playSound() {
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

func (exc Exchange) readWriteExchange() {
	for {
		select {
		case exc.ReadChan <- exc.currentValue:
			continue
		case exc.currentValue = <-exc.WriteChan:
			continue
		}
	}
}

func NewExchange(currency string, upLimit float64, downLimit float64) Exchange {
	writeChan := make(chan float64)
	readChan := make(chan float64)
	exc := Exchange{readChan, writeChan, downLimit + 0.01, upLimit, downLimit, currency}

	go exc.readWriteExchange()
	go exc.beepIfLimitReached()
	return exc
}
