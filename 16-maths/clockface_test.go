package clockface

import (
	"math"
	"testing"
	"time"
)

// func TestSecondHanfMidnight(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 - 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}

// }

// func TestSecondHAdAt30Seconds(t *testing.T) {
// 	tm := time.Date(1377, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(tm)

// 	if got != want {
// 		t.Errorf("Got %v, wanted %v", got, want)
// 	}
// }

// func TestSecondsInRadians(t *testing.T) {
// 	thirtySeconds := time.Date(1377, time.January, 1, 0, 0, 30, 0, time.UTC)
// 	want := math.Pi
// 	got := secondsInRadians(thirtySeconds)

// 	if want != got {
// 		t.Fatalf("Wanted %v radians, but got %v", want, got)
// 	}

// }

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), math.Pi},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 30), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}

}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1492, time.October, 12, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
