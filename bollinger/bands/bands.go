package bands

import (
	"math"
	"time"

	"github.com/insionng/bollinger-bands/bollinger/data"
	"github.com/insionng/bollinger-bands/bollinger/settings"
	//"fmt"
)

type Band struct {
	Date  time.Time
	Close float64
	SMA   float64
	Up    float64
	Down  float64
}

const (
	DAYS    = 60 * 60 * 24
	HOURS   = 60 * 60
	MINUTES = 60
)

//CalculatesBands Calculates bands for all times
func CalculatesBands(symbol string) (bands []Band) {

	historical := data.HistoricalClosingValues(symbol, settings.StartTime, HOURS, (settings.SMA_TIME + settings.GRAPH_TIME))

	start := 0
	end := settings.SMA_TIME

	for i := 0; i < settings.GRAPH_TIME; i++ {
		//fmt.Println(historical[start].Date.String(), historical[end-1].Date.String())
		if len(historical) >= end {
			bands = append(bands, CalculatesBand(historical[start:end]))
			start++
			end++
		}
	}

	return
}

//CalculatesBand Calculates bands for one time
func CalculatesBand(historical []data.ClosingValue) (result Band) {
	size := len(historical)
	//fmt.Println(size, historical[0].Date.String(), historical[size-1].Date.String())

	sum := float64(0)
	for _, h := range historical {
		sum += h.Value
	}

	// simple moving average
	sma := sum / float64(size)

	squares := float64(0)
	for i := 0; i < size; i++ {
		squares += math.Pow((historical[i].Value - sma), 2)
	}

	// standard deviation
	dev := math.Sqrt(squares / float64(size))

	// upper band
	up := sma + (settings.STANDARD_DEVIATIONS * dev)

	// lower band
	down := sma - (settings.STANDARD_DEVIATIONS * dev)

	return Band{historical[0].Date, historical[0].Value, sma, up, down}
}
