package egu

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	fmt.Println(GetUnixTime())
	fmt.Println(GetMicroTimeStampStr())
	fmt.Println(GetTimeStr())
	fmt.Println(GetTime())
	fmt.Println(GetDateStr())
	fmt.Println(GetTimeStamp())
	fmt.Println(GetLocalTimeZone())
}
