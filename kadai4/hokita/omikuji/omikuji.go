package omikuji

import (
	"math/rand"
	"time"
)

type Omikuji struct {
	Time time.Time
}

type Result struct {
	Result string `json:"result"`
}

func New(t time.Time) *Omikuji {
	return &Omikuji{t}
}

func (o *Omikuji) Draw() *Result {
	results := []string{"大吉", "中吉", "吉", "末吉", "凶", "大凶"}

	if o.isNewYear(o.Time) {
		return &Result{results[0]}
	}
	return &Result{results[rand.Intn(6)]}
}

func (o *Omikuji) isNewYear(time time.Time) bool {
	newYearDays := []string{
		"Jan 1",
		"Jan 2",
		"Jan 3",
	}

	for _, day := range newYearDays {
		if day == time.Format("Jan 2") {
			return true
		}
	}
	return false
}
