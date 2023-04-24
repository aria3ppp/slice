package slice_test

import (
	"reflect"
	"testing"

	"github.com/aria3ppp/slice"
)

func TestPop(t *testing.T) {
	type args[T any] struct {
		slc   []T
		index int
	}
	tests := []struct {
		name     string
		args     args[int]
		wantItem int
		wantOk   bool
		wantSlc  []int
	}{
		{
			name: "pop from slice with length 1",
			args: args[int]{
				slc:   []int{1},
				index: 0,
			},
			wantItem: 1,
			wantOk:   true,
			wantSlc:  []int{},
		},
		{
			name: "pop index 0 from slice with length 2",
			args: args[int]{
				slc:   []int{1, 2},
				index: 0,
			},
			wantItem: 1,
			wantOk:   true,
			wantSlc:  []int{2},
		},
		{
			name: "pop index 1 from slice with length 2",
			args: args[int]{
				slc:   []int{1, 2},
				index: 1,
			},
			wantItem: 2,
			wantOk:   true,
			wantSlc:  []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slc := slice.From(tt.args.slc)
			gotItem := slc.Pop(tt.args.index)
			if !reflect.DeepEqual(gotItem, tt.wantItem) {
				t.Errorf("Pop() gotItem = %v, want %v", gotItem, tt.wantItem)
			}
			if !reflect.DeepEqual(slc.Into(), tt.wantSlc) {
				t.Errorf("Pop() slic = %v, want %v", slc.Into(), tt.wantSlc)
			}
		})
	}
}
