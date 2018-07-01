package strategies

// several investment strategies

import (
	"github.com/insionng/bollinger-bands/bollinger/bands"
)

//MoreDown 在最近过去的特定时间段内，收盘价在下轨通过的次数大于收盘价超过上轨的次数，返回真值
func MoreDown(all []bands.Band) bool {
	up := 0
	down := 0

	for _, b := range all {
		if b.Close >= b.Up {
			up++
		} else if b.Close <= b.Down {
			down++
		}
	}

	if down > up {
		return true
	} else {
		return false
	}
}

//MoreUp 在最近过去的特定时间段内，收盘价超过上轨的次数大于收盘价低于下轨的次数，返回真值
func MoreUp(all []bands.Band) bool {
	up := 0
	down := 0

	for _, b := range all {
		if b.Close >= b.Up {
			up++
		} else if b.Close <= b.Down {
			down++
		}
	}

	if up > down {
		return true
	} else {
		return false
	}
}

//UpOnce 在最近过去的特定时间段内，收盘价至少突破一次上轨，返回真值
func UpOnce(all []bands.Band) bool {
	for _, b := range all {
		if b.Close >= b.Up {
			return true
		}
	}
	return false
}

//DownOnce 在最近过去的特定时间段内，收盘价至少突破一次下轨，返回真值
func DownOnce(all []bands.Band) bool {
	for _, b := range all {
		if b.Close <= b.Down {
			return true
		}
	}
	return false
}
