package omikuji

import (
	"fmt"
	"time"
)

var NewOmikuji = newOmikuji
var OmikujiHandler = (*Omikuji).omikujiHandler
var ThrowOneToSix = (*Omikuji).throwOneToSix
var IntToStr = (*Omikuji).intToStr
var IsNewYearHoliday = (*Omikuji).isNewYearHoliday
var GenJson = (*Omikuji).genJson
var TryOmikuji = (*Omikuji).tryOmikuji

var StdStdMethods = StdMethods

var ErrStdMethods = StdLibProvider{
	JsonMarshal: func(v interface{}) ([]byte, error){
		return nil, fmt.Errorf("json.marshal dummy error...")
	},
}

var MockStdMethods = StdLibProvider{
	RandIntn:func(i int) int {
		return 5
	},
}

var MockHoliday = StdLibProvider{
	TimeNow: func () time.Time {
		return time.Date(2020, time.January, 1, 10,23,30,9, time.UTC)
	},
}

var MockNotHoliday = StdLibProvider{
	TimeNow: func () time.Time {
		return time.Date(2020, time.March, 1, 10,23,30,9, time.UTC)
	},
	RandIntn:func(i int) int {
		return 5
	},
}

