package main

import (
	"fmt"
	"sync"

	"github.com/insionng/bollinger-bands/bollinger/bands"
	"github.com/insionng/bollinger-bands/bollinger/plot"
	"github.com/insionng/bollinger-bands/bollinger/strategies"
)

//Suggest 利用每个策略代号计算投资建议
func Suggest(symbols []string, strategy string) (suggest bool) {
	// default strategy
	fn := strategies.MoreDown

	switch strategy {
	case "moredown":
		fn = strategies.MoreDown
	case "moreup":
		fn = strategies.MoreUp
	case "uponce":
		fn = strategies.UpOnce
	case "downonce":
		fn = strategies.DownOnce
	default:
		fmt.Println("Invalid strategy: ", strategy, "use any of: moredown, moreup, uponce, downonce")
		return
	}

	// process each symbol in parallel
	var wg sync.WaitGroup
	wg.Add(len(symbols))

	for _, symbol := range symbols {
		go func(symbol string) {
			all := bands.CalculatesBands(symbol)
			if fn(all) {
				suggest = true
			} else {
				suggest = false
			}
			wg.Done()
		}(symbol)
	}

	// wait for execution of all goroutines
	wg.Wait()
	return

}

//Plot 根据代号执行绘图
func Plot(symbols []string) {

	// process each symbol in parallel
	var wg sync.WaitGroup
	wg.Add(len(symbols))

	for _, symbol := range symbols {
		go func(symbol string) {
			all := bands.CalculatesBands(symbol)
			plot.PlotBands(symbol, all)
			fmt.Println("Plot [", symbol, "] Okay")
			wg.Done()
		}(symbol)
	}

	// wait for execution of all goroutines
	wg.Wait()
}

func main() {

	var symbols = []string{"Bitcoin"}

	Plot(symbols)
	fmt.Println(Suggest(symbols, "moreup"))
}
