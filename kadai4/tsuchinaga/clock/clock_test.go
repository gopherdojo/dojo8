package clock

import (
	"reflect"
	"testing"
	"time"
)

func Test_clock_Now(t *testing.T) {
	c := &clock{}
	got := c.Now()
	now := time.Now()
	if got.After(now) {
		t.Errorf("Now() = %v, want before %v", got, now)
	}
}

func TestNew(t *testing.T) {
	want := &clock{}
	got := New()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}
