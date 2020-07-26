package omikuji

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Omikuji struct {
	DateTime time.Time `json:"time"`
	Dice     int       `json:"dice"`
	Result   string    `json:"result"`
}

func newOmikuji() *Omikuji {
	return &Omikuji{DateTime:time.Now()}
}

func (o *Omikuji) omikujiHandler(w http.ResponseWriter, r *http.Request) {
	err := o.tryOmikuji()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Run() {
	o := newOmikuji()
	http.HandleFunc("/", o.omikujiHandler)
	http.ListenAndServe(":8080", nil)
}

func (o *Omikuji)throwOneToSix() int {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(6)
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
	if !o.isNewYearHoliday(){
		o.Dice = o.throwOneToSix()
	} else {
		o.Dice = 6
	}
	o.Result, err = o.intToStr(o.Dice)
	return err
}

func (o *Omikuji) isNewYearHoliday () bool {
	_, m, d := o.DateTime.Date()
	if int(m) == 1 && d <= 3{
		return true
	}
	return false
}
