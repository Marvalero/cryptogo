package exchange_calculator

import (
	"fmt"
	"time"

	"github.com/Marvalero/cryptogo/app/beeper"
)

type Exchange struct {
	ReadCurrentValue  chan float64
	ReadNewValue      chan float64
	WriteCurrentValue chan float64
	currentValue      float64
	upLimit           float64
	downLimit         float64
	Currency          string
}

func (exc Exchange) beepIfLimitReached() {

	for {
		value := <-exc.ReadCurrentValue
		if value >= exc.upLimit {
			fmt.Println("---- UP", exc.Currency, " >= ", exc.upLimit)
			beeper.Beep()
		} else if value <= exc.downLimit {
			fmt.Println("---- DOWN", exc.Currency, " <= ", exc.downLimit)
			beeper.Beep()
		}
		time.Sleep(10 * time.Second)
	}
}

func (exc Exchange) handleExchange() {
	for {
		select {
		case exc.ReadCurrentValue <- exc.currentValue:
			continue
		case exc.currentValue = <-exc.WriteCurrentValue:
			exc.ReadNewValue <- exc.currentValue
			continue
		}
	}
}

func NewExchange(currency string, upLimit float64, downLimit float64) Exchange {
	writeCurrentValue := make(chan float64)
	readCurrentValue := make(chan float64)
	readNewValue := make(chan float64, 10)
	exc := Exchange{readCurrentValue, writeCurrentValue, readNewValue, downLimit + 0.01, upLimit, downLimit, currency}

	go exc.handleExchange()
	go exc.beepIfLimitReached()
	return exc
}
