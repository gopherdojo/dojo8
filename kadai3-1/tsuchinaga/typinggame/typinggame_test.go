package typinggame

import (
	"math/rand"
	"testing"
)

func Test_game_Next(t *testing.T) {
	t.Parallel()
	type fields struct {
		r    *rand.Rand
		q    string
		pass int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "randの値が98の場合は3文字", fields: fields{r: rand.New(rand.NewSource(98))}, want: "nnn"},
		{name: "randの値が99の場合は4文字", fields: fields{r: rand.New(rand.NewSource(99))}, want: "hsso"},
		{name: "randの値が83の場合は5文字", fields: fields{r: rand.New(rand.NewSource(83))}, want: "jdree"},
		{name: "randの値が79の場合は6文字", fields: fields{r: rand.New(rand.NewSource(79))}, want: "wqtjme"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := &game{}

			// ロックのテスト
			g.mutex.Lock()
			go func() {
				defer g.mutex.Unlock()
				g.r = tt.fields.r
			}()

			if got := g.Next(); got != tt.want {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_game_Answer(t *testing.T) {
	t.Parallel()

	type fields struct {
		q    string
		pass int
	}
	type args struct {
		ans string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{name: "qとansが一致している場合にpassが加算される", fields: fields{q: "foo", pass: 3}, args: args{ans: "foo"}, want: 4},
		{name: "qとansが一致していない場合はpassは変わらない", fields: fields{q: "foo", pass: 3}, args: args{ans: "foo"}, want: 4},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			g := &game{}

			// ロックのテスト
			g.mutex.Lock()
			go func() {
				defer g.mutex.Unlock()
				g.q = tt.fields.q
				g.pass = tt.fields.pass
			}()

			g.Answer(tt.args.ans)
			if g.pass != tt.want {
				t.Errorf("game.pass = %v, want %v", g.pass, tt.want)
			}
		})
	}
}

func Test_game_Result(t *testing.T) {
	t.Parallel()
	type fields struct {
		pass int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{name: "passが0なら0が返される", fields: fields{pass: 0}, want: 0},
		{name: "passが3なら3が返される", fields: fields{pass: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &game{}

			// ロックのテスト
			g.mutex.Lock()
			go func() {
				defer g.mutex.Unlock()
				g.pass = tt.fields.pass
			}()
			if got := g.Result(); got != tt.want {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}
