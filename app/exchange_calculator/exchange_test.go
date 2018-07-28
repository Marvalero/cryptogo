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
		exchange := NewExchange("BGP", 20.0, 1.0)
		exchange.WriteCurrentValue <- testCase.write
		read := <-exchange.ReadCurrentValue
		if read != testCase.read {
			t.Errorf("Exchange returned %f but expected %f", read, testCase.read)
		}
	}
}
