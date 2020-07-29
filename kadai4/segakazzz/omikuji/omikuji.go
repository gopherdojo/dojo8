package omikuji

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Omikuji struct {
	DateTime time.Time `json:"time"`
	Dice     int       `json:"dice"`
	Result   string    `json:"result"`
}

type StdLibProvider struct {
	JsonMarshal func(v interface{}) ([]byte, error)
	RandIntn func(int) int
	TimeNow func() time.Time
	HttpListenAndServe func (addr string, handler http.Handler) error
}

var StdMethods = StdLibProvider{
	JsonMarshal: json.Marshal,
	RandIntn: rand.Intn,
	TimeNow: time.Now,
	HttpListenAndServe: http.ListenAndServe,
}

func newOmikuji() *Omikuji {
	return &Omikuji{}
}

func (o *Omikuji) omikujiHandler(w http.ResponseWriter, r *http.Request) {
	err := o.tryOmikuji()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := o.genJson()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Run(port int) error {
	o := newOmikuji()
	http.HandleFunc("/", o.omikujiHandler)
	fmt.Println("Server is starting with port " +strconv.Itoa(port), "👍")
	err := StdMethods.HttpListenAndServe(":" + strconv.Itoa(port), nil)
	if err != nil {
		return errors.Wrapf(err, "Error in Run()\n")
	}
	return nil
}

func (o *Omikuji)throwOneToSix() int {
	rand.Seed(o.DateTime.UnixNano())
	i := StdMethods.RandIntn(6)
	return i + 1
}

func (o *Omikuji)intToStr(n int) (string, error) {
	switch n {
	case 1:
		return "凶", nil
	case 2, 3:
		return "吉", nil
	case 4, 5:
		return "中吉", nil
	case 6:
		return "大吉", nil
	default:
		return "", fmt.Errorf("invalid number %d", n)
	}
}

func (o *Omikuji) tryOmikuji() error {
	var err error
	o.DateTime = StdMethods.TimeNow()
	if !o.isNewYearHoliday(){
		o.Dice = o.throwOneToSix()
	} else {
		o.Dice = 6
	}
	o.Result, err = o.intToStr(o.Dice)
	return err
}

func (o *Omikuji) genJson() ([]byte, error) {
	js, err := StdMethods.JsonMarshal(o)
	if err != nil {
		return nil, err
	}
	return js, nil
}

func (o *Omikuji) isNewYearHoliday () bool {
	_, m, d := o.DateTime.Date()
	if int(m) == 1 && d <= 3{
		return true
	}
	return false
}
