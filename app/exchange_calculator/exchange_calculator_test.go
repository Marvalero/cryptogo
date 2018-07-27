package exchange_calculator

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var exchangeCalculatorTestCases = []struct {
	requestBody   string
	expectedValue float64
}{
	{"{\"USD\":33.33,\"GBP\":10.55}", 10.55},
	{"{\"USD\":33.33,\"GBP\":9.999}", 9.999},
	{"{\"GBP\":\"potato\"}", 0.0},
	{"{\"AAA\":33.33}", 0.0},
	{"watermelon", 0.0},
}

type ClientMock struct {
	returnMessage string
}

func (c *ClientMock) Get(url string) (*http.Response, error) {
	t := http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(c.returnMessage)),
	}
	return &t, nil
}

func TestExchangeCalculator(t *testing.T) {
	for _, testCase := range exchangeCalculatorTestCases {
		readChan := StartExchangeCalculator(&ClientMock{testCase.requestBody})
		time.Sleep(1 * time.Second)
		gbpValue := <-readChan
		if gbpValue != testCase.expectedValue {
			t.Errorf("handler returned wrong status code: got %v want %v",
				gbpValue, testCase.expectedValue)
		}
	}
}
