package omikuji_test

import (
	"github.com/gopherdojo/dojo8/kadai4/segakazzz/omikuji"
	"reflect"
	"testing"
	"time"
)

func TestOmikuji_genJson(t *testing.T) {
	type fields struct {
		DateTime time.Time
		Dice     int
		Result   string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "Expects Success 1",
			fields: fields{
				DateTime: time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
				Dice:     1,
				Result:   "凶",
			},
			want: "{\"time\":\"2009-11-10T23:00:00Z\",\"dice\":1,\"result\":\"凶\"}",
			wantErr: false,
		},
		{
			name: "Expects Error",
			fields: fields{
				DateTime: time.Date(2020, time.July, 21, 5, 9, 23, 3424, time.UTC),
				Dice:     6,
				Result:   "大吉",
			},
			want: "",
			wantErr: true,
		},
		{
			name: "Expects Success 2",
			fields: fields{
				DateTime: time.Date(2020, time.July, 21, 5, 9, 23, 3424, time.UTC),
				Dice:     6,
				Result:   "大吉",
			},
			want: "{\"time\":\"2020-07-21T05:09:23.000003424Z\",\"dice\":6,\"result\":\"大吉\"}",
			wantErr: false,
		},

	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			o := &omikuji.Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if tt.wantErr == true {
				omikuji.StdMethods = omikuji.ErrStdMethods
			} else {
				omikuji.StdMethods = omikuji.StdStdMethods
			}

			got, err := omikuji.GenJson(o)
			if (err != nil) != tt.wantErr {
				t.Errorf("genJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("genJson() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestOmikuji_intToStr(t *testing.T) {
	type fields struct {
		DateTime time.Time
		Dice     int
		Result   string
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &omikuji.Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			got, err := omikuji.IntToStr(o, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("intToStr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("intToStr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOmikuji_isNewYearHoliday(t *testing.T) {
	type fields struct {
		DateTime time.Time
		Dice     int
		Result   string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &omikuji.Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if got := omikuji.IsNewYearHoliday(o); got != tt.want {
				t.Errorf("isNewYearHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestOmikuji_omikujiHandler(t *testing.T) {
//	type fields struct {
//		DateTime time.Time
//		Dice     int
//		Result   string
//	}
//	type args struct {
//		w http.ResponseWriter
//		r *http.Request
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			o := &omikuji.Omikuji{
//				DateTime: tt.fields.DateTime,
//				Dice:     tt.fields.Dice,
//				Result:   tt.fields.Result,
//			}
//		})
//	}
//}

func TestOmikuji_throwOneToSix(t *testing.T) {
	type fields struct {
		DateTime time.Time
		Dice     int
		Result   string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &omikuji.Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if got := omikuji.ThrowOneToSix(o); got != tt.want {
				t.Errorf("throwOneToSix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOmikuji_tryOmikuji(t *testing.T) {
	type fields struct {
		DateTime time.Time
		Dice     int
		Result   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &omikuji.Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if err := omikuji.TryOmikuji(o); (err != nil) != tt.wantErr {
				t.Errorf("tryOmikuji() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRun(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := omikuji.Run(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newOmikuji(t *testing.T) {
	tests := []struct {
		name string
		want *omikuji.Omikuji
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := omikuji.NewOmikuji(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOmikuji() = %v, want %v", got, tt.want)
			}
		})
	}
}
