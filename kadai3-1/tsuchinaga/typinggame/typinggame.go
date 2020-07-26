package typinggame

import (
	"math/rand"
	"sync"
)

// New - 新しいタイピンゲームの生成
func New(seed int64) Game {
	return &game{
		r: rand.New(rand.NewSource(seed)),
	}
}

// Game - ゲームのinterface
type Game interface {
	Next() string
	Answer(ans string)
	Result() int
}

// game - ゲーム
type game struct {
	r     *rand.Rand // 乱数
	q     string     // 問題
	pass  int        // 正答数
	mutex sync.Mutex
}

// Next - 次の問題
func (g *game) Next() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.q = ""
	n := g.r.Intn(4) + 3
	for i := 0; i < n; i++ {
		g.q += string('a' + g.r.Intn(24))
	}
	return g.q
}

// Answer - 回答入力
func (g *game) Answer(ans string) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.q == ans {
		g.pass++
	}
}

func (g *game) Result() int {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.pass
}
