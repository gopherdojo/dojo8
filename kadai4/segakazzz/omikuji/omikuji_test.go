package omikuji

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

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
			o := &Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			got, err := o.intToStr(tt.args.n)
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
			o := &Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if got := o.isNewYearHoliday(); got != tt.want {
				t.Errorf("isNewYearHoliday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOmikuji_omikujiHandler(t *testing.T) {
	type fields struct {
		DateTime time.Time
		Dice     int
		Result   string
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
		})
	}
}

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
			o := &Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if got := o.throwOneToSix(); got != tt.want {
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
			o := &Omikuji{
				DateTime: tt.fields.DateTime,
				Dice:     tt.fields.Dice,
				Result:   tt.fields.Result,
			}
			if err := o.tryOmikuji(); (err != nil) != tt.wantErr {
				t.Errorf("tryOmikuji() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_newOmikuji(t *testing.T) {
	tests := []struct {
		name string
		want *Omikuji
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newOmikuji(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOmikuji() = %v, want %v", got, tt.want)
			}
		})
	}
}