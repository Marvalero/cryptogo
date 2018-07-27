package exchange_calculator

import "testing"

var exchangeTestCases = []struct {
	write float64
	read  float64
}{
	{9.323, 9.323},
	{10.0, 10.0},
}

func TestExchange(t *testing.T) {
	for _, testCase := range exchangeTestCases {
		exchange := NewExchange()
		exchange.WriteChan <- testCase.write
		read := <-exchange.ReadChan
		if read != testCase.read {
			t.Errorf("Exchange returned %f but expected %f", read, testCase.read)
		}
	}
}
