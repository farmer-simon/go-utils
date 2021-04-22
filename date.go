package egu


import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	DateFormat = "2006-01-02"
)

// GetUnixTime
func GetUnixTime() time.Time  {
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