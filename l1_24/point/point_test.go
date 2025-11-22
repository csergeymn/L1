package point

import "testing"

func TestDistance(t *testing.T) {
	tests := []struct {
		p1, p2   Point
		expected float64
	}{
		{NewPoint(0, 0), NewPoint(3, 4), 5},
		{NewPoint(1, 1), NewPoint(1, 1), 0},
		{NewPoint(-1, -1), NewPoint(2, 3), 5},
		{NewPoint(10, 1), NewPoint(11, 2), 1.4142135623730951},
	}

	for _, tt := range tests {
		got := tt.p1.Distance(tt.p2)

		if got != tt.expected {
			t.Errorf("Distance(%v, %v) = %v; want %v",
				tt.p1, tt.p2, got, tt.expected)
		}
	}
}
