package main

import (
	"reflect"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		index     int
		want      []int
		expectErr bool
	}{
		{
			name:      "remove middle element",
			input:     []int{1, 2, 3, 4, 5},
			index:     2,
			want:      []int{1, 2, 4, 5},
			expectErr: false,
		},
		{
			name:      "remove first element",
			input:     []int{10, 20, 30},
			index:     0,
			want:      []int{20, 30},
			expectErr: false,
		},
		{
			name:      "remove last element",
			input:     []int{1, 2, 3},
			index:     2,
			want:      []int{1, 2},
			expectErr: false,
		},
		{
			name:      "invalid negative index",
			input:     []int{1, 2, 3},
			index:     -1,
			want:      []int{1, 2, 3},
			expectErr: true,
		},
		{
			name:      "invalid out-of-range index",
			input:     []int{1, 2, 3},
			index:     5,
			want:      []int{1, 2, 3},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := remove(tt.input, tt.index)

			if tt.expectErr {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
