package plot

import (
	"fmt"
	"image/color"
	"log"

	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"github.com/insionng/bollinger-bands/bollinger/bands"
)

//PlotBands 绘制带宽
func PlotBands(symbol string, all []bands.Band) {
	n := len(all)
	dclose := make(plotter.XYs, n)
	dsma := make(plotter.XYs, n)
	dup := make(plotter.XYs, n)
	ddown := make(plotter.XYs, n)

	for i, b := range all {
		dclose[i].X = float64(-1 * i)
		dclose[i].Y = b.Close

		dsma[i].X = float64(-1 * i)
		dsma[i].Y = b.SMA

		dup[i].X = float64(-1 * i)
		dup[i].Y = b.Up

		ddown[i].X = float64(-1 * i)
		ddown[i].Y = b.Down
	}

	p, err := plot.New()
	if err != nil {
		log.Fatalln("plot error:", err)
	}

	p.Title.Text = fmt.Sprintf("Bollinger Bands: %s", symbol)
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Value"

	p.Add(plotter.NewGrid())

	lclose, _ := plotter.NewLine(dclose)

	lsma, _ := plotter.NewLine(dsma)
	lsma.LineStyle.Color = color.RGBA{B: 255, A: 255}

	lup, _ := plotter.NewLine(dup)
	lup.LineStyle.Color = color.RGBA{R: 255, A: 255}

	ldown, _ := plotter.NewLine(ddown)
	ldown.LineStyle.Color = color.RGBA{G: 255, A: 255}

	p.Add(lclose, lsma, lup, ldown)
	p.Legend.Add("Close", lclose)
	p.Legend.Add("SMA", lsma)
	p.Legend.Add("Up", lup)
	p.Legend.Add("Down", ldown)

	p.Save(16, 9, fmt.Sprintf("%s.png", symbol))
}
