package api

import (
	"net/http"

	"github.com/Marvalero/cryptogo/app/exchange_calculator"
	chart "github.com/wcharczuk/go-chart"
)

type ExchangeApi struct {
	Exchannel chan [100]exchange_calculator.Serie
}

func (exc ExchangeApi) ShowExchange(res http.ResponseWriter, r *http.Request) {
	serie := <-exc.Exchannel
	xValue := make([]float64, len(serie))
	yValue := make([]float64, len(serie))
	for i, point := range serie {
		xValue[i] = float64(point.Timestamp)
		yValue[i] = point.Value
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Time",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "GBP",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValue,
				YValues: yValue,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	res.WriteHeader(http.StatusOK)
	graph.Render(chart.PNG, res)
	// currentExchange := <-exc.Exchannel
	// message := fmt.Sprintf("exchange: %.6f", currentExchange)
	// res.Write([]byte(message))
}
