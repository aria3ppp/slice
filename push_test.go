package slice_test

import (
	"reflect"
	"testing"

	"github.com/aria3ppp/slice"
)

func TestPush(t *testing.T) {
	type args[T any] struct {
		slc   []T
		index int
		item  T
	}
	tests := []struct {
		name    string
		args    args[int]
		wantSlc []int
	}{
		{
			name: "push to slice with length 1",
			args: args[int]{
				slc:   []int{1},
				index: 0,
				item:  0,
			},
			wantSlc: []int{0, 1},
		},
		{
			name: "push to index 0 of slice with length 2",
			args: args[int]{
				slc:   []int{1, 2},
				index: 0,
				item:  0,
			},
			wantSlc: []int{0, 1, 2},
		},
		{
			name: "push to index 1 of slice with length 2",
			args: args[int]{
				slc:   []int{1, 2},
				index: 1,
				item:  3,
			},
			wantSlc: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slc := slice.From(tt.args.slc)
			slc.Push(tt.args.item, tt.args.index)
			if !reflect.DeepEqual(slc.Into(), tt.wantSlc) {
				t.Errorf("Pop() slic = %v, want %v", slc.Into(), tt.wantSlc)
			}
		})
	}
}
