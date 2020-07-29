package omikuji

import "fmt"

var NewOmikuji = newOmikuji
var OmikujiHandler = (*Omikuji).omikujiHandler
var ThrowOneToSix = (*Omikuji).throwOneToSix
var IntToStr = (*Omikuji).intToStr
var IsNewYearHoliday = (*Omikuji).isNewYearHoliday
var GenJson = (*Omikuji).genJson
var TryOmikuji = (*Omikuji).tryOmikuji

var StdStdMethods = StdMethods

var ErrStdMethods = stdLibProvider{
	jsonMarshal: func(v interface{}) ([]byte, error){
		return nil, fmt.Errorf("json.marshal dummy error...")
	},
}
