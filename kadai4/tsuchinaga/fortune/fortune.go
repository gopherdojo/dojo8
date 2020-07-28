package fortune

import (
	"math/rand"
	"time"
)

// Foredoom - 吉凶
type Foredoom string

const (
	ForedoomDaikichi Foredoom = "大吉"
	ForedoomKichi    Foredoom = "吉"
	ForedoomChukichi Foredoom = "中吉"
	ForedoomShokichi Foredoom = "小吉"
	ForedoomKyou     Foredoom = "凶"
)

var foredooms = []Foredoom{ForedoomDaikichi, ForedoomKichi, ForedoomChukichi, ForedoomShokichi, ForedoomKyou}

// Paper - 結果の載ってる紙
type Paper struct {
	No       int      `json:"no"`       // おみくじ番号 [1,100]
	Foredoom Foredoom `json:"foredoom"` // 吉凶
	Wish     Foredoom `json:"wish"`     // 願望
	Health   Foredoom `json:"health"`   // 健康
	Business Foredoom `json:"business"` // 仕事
	Love     Foredoom `json:"love"`     // 恋愛
	Study    Foredoom `json:"study"`    // 勉強
}

func Get(now time.Time) Paper {
	if now.Month() == 1 && now.Day() <= 3 {
		return Paper{
			No:       now.Day(),
			Foredoom: ForedoomDaikichi,
			Wish:     ForedoomDaikichi,
			Health:   ForedoomDaikichi,
			Business: ForedoomDaikichi,
			Love:     ForedoomDaikichi,
			Study:    ForedoomDaikichi,
		}
	}

	r := rand.New(rand.NewSource(now.UnixNano()))
	return Paper{
		No:       r.Intn(100) + 1,
		Foredoom: foredooms[r.Intn(5)],
		Wish:     foredooms[r.Intn(5)],
		Health:   foredooms[r.Intn(5)],
		Business: foredooms[r.Intn(5)],
		Love:     foredooms[r.Intn(5)],
		Study:    foredooms[r.Intn(5)],
	}
}
