package utils

import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// GetUnixTime
func GetUnixTime() time.Time {
	return time.Now()
}

// GetTime return current local time
func GetTime() time.Time {
	return GetUnixTime().In(GetLocalTimeZone())
}

// GetLocalTimeZone
func GetLocalTimeZone() *time.Location {
	return time.FixedZone("CST", 8*3600) // UTC+8
}

// GetDateStr return current date string,eg: 2019-12-30
func GetDateStr() string {
	return GetTime().Format(DateFormat)
}

// GetTimeStr return current time string,eg:2019-12-30 22:00:00
func GetTimeStr() string {
	return GetTime().Format(TimeFormat)
}

// GetTimeStamp return current timestamp
func GetTimeStamp() int64 {
	return GetTime().Unix()
}

// GetMicroTimeStampStr return micro timestamp string
func GetMicroTimeStampStr() string {
	return fmt.Sprintf("%.6f", float64(GetTime().UnixNano())/1e9)
}

// DatetimeToUnixTimestamp return unix timestamp
func DatetimeToUnixTimestamp(toBeCharge string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, toBeCharge, loc)
	sr := theTime.Unix()
	return sr
}

// UnixTimestampToDatetime return datetime string
func UnixTimestampToDatetime(timestamp int64) string {
	dataTimeStr := time.Unix(timestamp, 0).Format(TimeFormat)
	return dataTimeStr
}

// ParseDateTime return year, month, day, hour, minute, second
func ParseDateTime(toBeCharge string) (year, month, day, hour, minute, second int) {
	timeObj, err := time.Parse(TimeFormat, toBeCharge)
	if err != nil {
		return
	}
	year = timeObj.Year()
	month = int(timeObj.Month())
	day = timeObj.Day()
	hour = timeObj.Hour()
	minute = timeObj.Minute()
	second = timeObj.Second()
	return
}

// ParseDate return year, month, day
func ParseDate(toBeCharge string) (year, month, day int) {
	timeObj, err := time.Parse(DateFormat, toBeCharge)
	if err != nil {
		return
	}
	year = timeObj.Year()
	month = int(timeObj.Month())
	day = timeObj.Day()
	return
}

func leapYears(date time.Time) (leaps int) {
	// returns year, month,
	// date of a time object
	y, m, _ := date.Date()
	if m <= 2 {
		y--
	}
	leaps = y/4 + y/400 - y/100
	return leaps
}

// GetDifference 获取两个时间相差的天数、小时、分、秒
func GetDifference(sTime, eTime string) (days, hours, minutes, seconds int) {

	startTime, err := time.Parse(TimeFormat, sTime)
	if err != nil {
		startTime = time.Now()
	}
	endTime, err := time.Parse(TimeFormat, eTime)
	if err != nil {
		endTime = time.Now()
	}
	if startTime.After(endTime) {
		startTime, endTime = endTime, startTime
	}
	// month-wise days
	monthDays := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// extracting years, months,
	// days of two dates
	y1, m1, d1 := startTime.Date()
	y2, m2, d2 := endTime.Date()

	// extracting hours, minutes,
	// seconds of two times
	h1, min1, s1 := startTime.Clock()
	h2, min2, s2 := endTime.Clock()

	// totalDays since the
	// beginning = year*365 + number_of_days
	totalDays1 := y1*365 + d1

	// adding days of the months
	// before the current month
	for i := 0; i < (int)(m1)-1; i++ {
		totalDays1 += monthDays[i]
	}

	// counting leap years since
	// beginning to the year "a"
	// and adding that many extra
	// days to the totaldays
	totalDays1 += leapYears(startTime)

	// Similar procedure for second date
	totalDays2 := y2*365 + d2

	for i := 0; i < (int)(m2)-1; i++ {
		totalDays2 += monthDays[i]
	}

	totalDays2 += leapYears(startTime)

	// Number of days between two days
	days = totalDays2 - totalDays1

	// calculating hour, minutes,
	// seconds differences
	hours = h2 - h1
	minutes = min2 - min1
	seconds = s2 - s1

	// if seconds difference goes below 0,
	// add 60 and decrement number of minutes
	if seconds < 0 {
		seconds += 60
		minutes--
	}

	// performing similar operations
	// on minutes and hours
	if minutes < 0 {
		minutes += 60
		hours--
	}

	// performing similar operations
	// on hours and days
	if hours < 0 {
		hours += 24
		days--
	}

	return days, hours, minutes, seconds

}
