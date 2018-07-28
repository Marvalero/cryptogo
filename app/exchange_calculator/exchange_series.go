package exchange_calculator

import (
	"sort"
	"strconv"
	"time"
)

type Serie struct {
	Value     float64
	Timestamp int64
}

func StartExchangeCalculator(client HttpClient, currency string, uplimit float64, downlimit float64) chan [100]Serie {
	exc := NewExchange(currency, uplimit, downlimit)
	go calculateExchange(exc, client)
	return newExchangeSeries(exc)
}

func newExchangeSeries(exchange Exchange) chan [100]Serie {
	readSeriesChan := make(chan [100]Serie)
	go exchangeSeries(exchange, readSeriesChan)
	return readSeriesChan
}

func exchangeSeries(exchange Exchange, seriesChan chan [100]Serie) {
	var series [100]Serie
	currentVal := <-exchange.ReadCurrentValue
	for i, _ := range series {
		series[i] = Serie{currentVal, makeTimestamp()}
	}
	for {
		select {
		case value := <-exchange.ObserveNewValue:
			newPoint := Serie{value, makeTimestamp()}
			series[0] = newPoint
			sort.Slice(series[:], func(i, j int) bool { return series[i].Timestamp < series[j].Timestamp })
			continue
		case seriesChan <- series:
			continue
		}
	}
}

func makeTimestamp() int64 {
	timestamp, _ := strconv.Atoi(time.Now().Format("150405"))
	return int64(timestamp)
}
