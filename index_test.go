package slice_test

import (
	"testing"

	"github.com/aria3ppp/slice"
)

func TestIndex(t *testing.T) {
	type args[T any] struct {
		slc  []T
		item T
	}
	tests := []struct {
		name      string
		args      args[int]
		wantIndex int
	}{
		{
			name: "empty slice",
			args: args[int]{
				slc:  []int{},
				item: 1,
			},
			wantIndex: -1,
		},
		{
			name: "slice with length 1",
			args: args[int]{
				slc:  []int{1},
				item: 1,
			},
			wantIndex: 0,
		},
		{
			name: "slice with length 1 but not found",
			args: args[int]{
				slc:  []int{0},
				item: 1,
			},
			wantIndex: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slc := slice.From(tt.args.slc)
			if gotIndex := slc.Index(tt.args.item); gotIndex != tt.wantIndex {
				t.Errorf("Index() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
