package timeUtil

import (
	"math"
	"time"
)

type PickTime struct {
	StartTime time.Time
	EndTime   time.Time
}

func ParseTimeRangeToPickTimes(sTime, eTime time.Time, interval int) []*PickTime {
	intervalDuring := time.Duration(interval) * time.Minute
	timeRange := eTime.Sub(sTime)
	pickTimeNo := int(math.Ceil(timeRange.Minutes() / intervalDuring.Minutes()))

	var pickTimes []*PickTime
	for idx := 0; idx < pickTimeNo; idx++ {
		tmpPickTime := PickTime{
			StartTime: sTime.Add(time.Duration(interval*idx) * time.Minute),
			EndTime:   sTime.Add(time.Duration(interval*(idx+1)) * time.Minute),
		}
		pickTimes = append(pickTimes, &tmpPickTime)
	}
	return pickTimes
}

func GetMonthSlice(sTime, eTime time.Time) []*PickTime {
	firstDateOfMonth := GetFirstDateOfMonth(eTime)
	if firstDateOfMonth.Equal(eTime) {
		firstDateOfMonth = firstDateOfMonth.AddDate(0, -1, 0)
	}
	var pickTimes []*PickTime
	if firstDateOfMonth.After(sTime) {
		tmpPickTime := PickTime{
			StartTime: firstDateOfMonth,
			EndTime:   eTime,
		}
		pickTimes = append(pickTimes, GetMonthSlice(sTime, firstDateOfMonth)...)
		pickTimes = append(pickTimes, &tmpPickTime)
	} else {
		tmpPickTime := PickTime{
			StartTime: sTime,
			EndTime:   eTime,
		}
		pickTimes = append(pickTimes, &tmpPickTime)
	}

	return pickTimes
}

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
