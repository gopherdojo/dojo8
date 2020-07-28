package fortune

import (
	"reflect"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		arg  time.Time
		want Paper
	}{
		{name: "2020/01/01 00:00:00は1番で全部大吉",
			arg: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local),
			want: Paper{
				No:       1,
				Foredoom: ForedoomDaikichi,
				Wish:     ForedoomDaikichi,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomDaikichi,
			}},
		{name: "2020/01/01 23:59:59は1番で全部大吉",
			arg: time.Date(2020, 1, 1, 23, 59, 59, 0, time.Local),
			want: Paper{
				No:       1,
				Foredoom: ForedoomDaikichi,
				Wish:     ForedoomDaikichi,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomDaikichi,
			}},
		{name: "2020/01/02は2番で全部大吉",
			arg: time.Date(2020, 1, 2, 0, 0, 0, 0, time.Local),
			want: Paper{
				No:       2,
				Foredoom: ForedoomDaikichi,
				Wish:     ForedoomDaikichi,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomDaikichi,
			}},
		{name: "2020/01/03は3番で全部大吉",
			arg: time.Date(2020, 1, 3, 0, 0, 0, 0, time.Local),
			want: Paper{
				No:       3,
				Foredoom: ForedoomDaikichi,
				Wish:     ForedoomDaikichi,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomDaikichi,
			}},
		{name: "2030/01/01は1番で全部大吉",
			arg: time.Date(2030, 1, 1, 0, 0, 0, 0, time.Local),
			want: Paper{
				No:       1,
				Foredoom: ForedoomDaikichi,
				Wish:     ForedoomDaikichi,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomDaikichi,
			}},
		{name: "2040/01/01は1番で全部大吉",
			arg: time.Date(2040, 1, 1, 0, 0, 0, 0, time.Local),
			want: Paper{
				No:       1,
				Foredoom: ForedoomDaikichi,
				Wish:     ForedoomDaikichi,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomDaikichi,
			}},
		{name: "2020/01/04 00:00:00はランダム",
			arg: time.Date(2020, 1, 4, 0, 0, 0, 0, time.Local),
			want: Paper{
				No:       30,
				Foredoom: ForedoomChukichi,
				Wish:     ForedoomKyou,
				Health:   ForedoomDaikichi,
				Business: ForedoomDaikichi,
				Love:     ForedoomDaikichi,
				Study:    ForedoomKyou,
			}},
		{name: "2020/01/04 00:00:01はランダム",
			arg: time.Date(2020, 1, 4, 0, 0, 1, 0, time.Local),
			want: Paper{
				No:       53,
				Foredoom: ForedoomKichi,
				Wish:     ForedoomShokichi,
				Health:   ForedoomShokichi,
				Business: ForedoomKyou,
				Love:     ForedoomKyou,
				Study:    ForedoomKichi,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
