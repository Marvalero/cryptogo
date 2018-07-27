package exchange_calculator

type Exchange struct {
	ReadChan     chan float64
	WriteChan    chan float64
	currentValue float64
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

func NewExchange() Exchange {
	writeChan := make(chan float64)
	readChan := make(chan float64)
	exc := Exchange{readChan, writeChan, 0.0}

	go exc.readWriteExchange()
	return exc
}
