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

func StartExchangeCalculator(client HttpClient) chan [100]Serie {
	ethExc := NewExchange("ETH", 380.0, 320.0)
	btcExc := NewExchange("BTC", 6300.0, 6100.0)

	go calculateExchange(ethExc, client)
	go calculateExchange(btcExc, client)
	return newExchangeSeries(ethExc.ReadNewValue)
}

func newExchangeSeries(newValueObserver chan float64) chan [100]Serie {
	readSeriesChan := make(chan [100]Serie)
	go exchangeSeries(newValueObserver, readSeriesChan)
	return readSeriesChan
}

func exchangeSeries(newValueObserver chan float64, seriesChan chan [100]Serie) {
	var series [100]Serie
	for i, _ := range series {
		series[i] = Serie{0.0, makeTimestamp()}
	}
	for {
		select {
		case value := <-newValueObserver:
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
