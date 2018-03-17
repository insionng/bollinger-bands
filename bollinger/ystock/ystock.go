package ystock

// wrapper for Yahoo Finance rest API

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	//"github.com/insionng/bollinger-bands/ext/timeext"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type ClosingValue struct {
	Value float64
	Date  time.Time
}

func HistoricalClosingValues(symbol string, end time.Time, days int) (historical []ClosingValue) {
	// HACK: holidays make it hard to know the days when the market is closed,
	// query again with a longer range if not enough days were retrieved the
	// last time
	qdays := days
	historical = []ClosingValue{}
	for len(historical) < days {
		// on average 10 week days adds 4 more weekend days plus 1 occasional holiday
		qdays := qdays + (qdays / 2)
		start := end.AddDate(0, 0, (-1 * qdays))
		//start, end = timeext.FixWeekdaysInterval(start, end)

		historical = query(symbol, start, end)
	}

	// remove extra days
	historical = historical[0:days]
	return
}

func query(symbol string, start time.Time, end time.Time) (historical []ClosingValue) {
	//fmt.Println(start, end)
	/*
		v := url.Values{}

		v.Set("s", symbol)
		v.Set("g", "d")
		v.Set("ignore", ".csv")

		v.Set("a", strconv.Itoa(int(start.Month()-1)))
		v.Set("b", strconv.Itoa(int(start.Day())))
		v.Set("c", strconv.Itoa(int(start.Year())))

		v.Set("d", strconv.Itoa(int(end.Month()-1)))
		v.Set("e", strconv.Itoa(int(end.Day())))
		v.Set("f", strconv.Itoa(int(end.Year())))

		query := fmt.Sprintf("http://ichart.yahoo.com/table.csv?%s", v.Encode())
	*/
	//fmt.Println(query)

	//resp, err := http.Get(query)

	// Open file and create GZ reader
	f, err := os.Open("../bollinger/ystock/hitbtcUSD.csv.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()

	// Read CSV from GZ file
	cr := csv.NewReader(gr)

	var n int = 0
	var curDate string
	for {
		n++
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("CSV Read Error:", err)
			return
		}

		var closingValue = ClosingValue{}
		for i, v := range record {
			if i == 0 { //时间
				ordertimeInt64, _ := strconv.ParseInt(v, 10, 64)
				if d := time.Unix(ordertimeInt64, 0).Format("2006-01-02"); curDate != d {
					curDate = d
					//ordertime, _ := timeext.ParseDate(v)
					ordertime := time.Unix(ordertimeInt64, 0) //.Format("2006-01-02 15:04:05")

					closingValue.Date = ordertime
					fmt.Printf("时间：%v ", ordertime)
				} else {
					break
				}
			}
			if i == 1 { //价格
				price, _ := strconv.ParseFloat(v, 64)
				fmt.Printf("价格：%v ", price)
				closingValue.Value = price
			}
			if i == 2 { //数量
				volume, _ := strconv.ParseFloat(v, 64)
				fmt.Printf("数量：%v\n", volume)
			}
		}

		historical = append(historical, closingValue)
	}

	return
}
