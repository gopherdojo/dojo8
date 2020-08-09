package omikuji_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/gopherdojo/dojo8/kadai4/hokita/omikuji/omikuji"
)

func TestMain(m *testing.M) {
	m.Run()
	rand.Seed(time.Now().UnixNano())
}

func TestOmikuji_Draw(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		inTime time.Time
		want1  *omikuji.Result
		want2  *omikuji.Result
	}{
		"standard day": {
			inTime: time.Date(2000, 1, 4, 0, 0, 0, 0, time.Local),
			want1:  &omikuji.Result{"大凶"},
			want2:  &omikuji.Result{"末吉"},
		},
		"new year": {
			inTime: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			want1:  &omikuji.Result{"大吉"},
			want2:  &omikuji.Result{"大吉"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			rand.Seed(1)

			o := &omikuji.Omikuji{
				test.inTime,
			}
			got1 := o.Draw()
			if !reflect.DeepEqual(got1, test.want1) {
				t.Fatalf(`want1: "%v" actual1: "%v"`, test.want1, got1)
			}
			got2 := o.Draw()
			if !reflect.DeepEqual(got2, test.want2) {
				t.Fatalf(`want2: "%v" actual2: "%v"`, test.want2, got2)
			}
		})
	}
}
