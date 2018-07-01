package data

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
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

//HistoricalClosingValues start 起始时间， interval 时间间隔 单位是秒，limit 窗口限制
func HistoricalClosingValues(symbol string, start time.Time, interval, limit int64) (historical []ClosingValue) {

	// Open file and create GZ reader
	f, err := os.Open("../bollinger/data/hitbtcUSD.csv.gz")
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

	var curTime, nTimes int64

	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("CSV Read Error:", err)
			return
		}

		if nTimes > limit {
			break
		}

		var has bool
		var closingValue = ClosingValue{}
		if len(record) == 3 {
			for i, v := range record {
				if i == 0 { //时间
					ordertimeInt64, _ := strconv.ParseInt(v, 10, 64)
					if start.Unix() < ordertimeInt64 {
						if curTime < ordertimeInt64 {
							curTime = ordertimeInt64 + interval

							ordertime := time.Unix(ordertimeInt64, 0)
							closingValue.Date = ordertime
							fmt.Printf("[%v] 时间：%v ", symbol, ordertime)
							nTimes++
							has = true
						}
					}
				}
				if i == 1 && has { //价格
					price, _ := strconv.ParseFloat(v, 64)
					fmt.Printf("价格：%v\n", price)
					closingValue.Value = price
				}
			}
		}
		if has {
			historical = append(historical, closingValue)
		}
	}

	return
}
