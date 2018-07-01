package settings

import (
	"time"
)

// number of days to calculate SMA
const SMA_TIME = 20

// number of days to display on graph
const GRAPH_TIME = 20

// standard deviation multiplier, default 2
// more info: http://www.great-trades.com/Help/bollinger%20bands%20calculation.htm
const STANDARD_DEVIATIONS = 2

//StartTime start date
var StartTime = time.Now().Add(-(time.Hour * 24 * 365))
