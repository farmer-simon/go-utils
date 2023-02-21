package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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
func ParseDateTime(toBeCharge string) (year, month, day, hour, minute, second, weekday int) {
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
	weekday = int(timeObj.Weekday())
	return
}

// ParseDate return year, month, day
func ParseDate(toBeCharge string) (year, month, day, weekday int) {
	timeObj, err := time.Parse(DateFormat, toBeCharge)
	if err != nil {
		return
	}
	year = timeObj.Year()
	month = int(timeObj.Month())
	day = timeObj.Day()
	weekday = int(timeObj.Weekday())
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

// GetDateTimeDifference 获取两个时间相差的天数、小时、分、秒, 开始时间是否大于结束时间
func GetDateTimeDifference(sTime, eTime string) (days, hours, minutes, seconds int, exchange bool) {
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
		exchange = true
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

	return days, hours, minutes, seconds, exchange

}

// FormatDateTimeStringToDateTime 格式化日期/时间字符串到标准时间格式
func FormatDateTimeStringToDateTime(oriTime string) (time.Time, error) {
	var y, m, d, h, i, s = "0", "0", "0", "0", "0", "0"
	if !strings.Contains(oriTime, ":") {

		reg := regexp.MustCompile(`(\d+)[/|-](\d+)[/|-](\d+)`) // 查找连续的小写字母
		match := reg.FindStringSubmatch(oriTime)
		if len(match) != 4 {
			return time.Now(), errors.New("暂不支持的日期格式：" + oriTime)
		}
		y, m, d = match[1], match[2], match[3]
	} else {
		reg := regexp.MustCompile(`(\d+)[/|-](\d+)[/|-](\d+)[T|\s](\d+):(\d+):?(\d+)?(\+)?(\d+)?:?(\d+)?`) // 查找连续的小写字母
		match := reg.FindStringSubmatch(oriTime)
		if len(match) < 7 {
			return time.Now(), errors.New("暂不支持的日期格式：" + oriTime)
		}
		y, m, d, h, i, s = match[1], match[2], match[3], match[4], match[5], match[6]
	}
	datetime := fmt.Sprintf("%04s-%02s-%02s %02s:%02s:%02s", y, m, d, h, i, s)
	parseTime, err := time.Parse(TimeFormat, datetime)

	return parseTime, err
}

// GetThisWeekFirstTime 获取当前周第一天的时间
func GetThisWeekFirstTime() (thisMonday int64) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	toBeCharge := weekStartDate.Format(TimeFormat)

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(TimeFormat, toBeCharge, loc)
	thisMonday = theTime.Unix()
	return
}
