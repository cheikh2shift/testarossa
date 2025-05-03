
package pkg2

import (
	"reflect"
	"testing"

	"github.com/cheikh2shift/testarossa/test_src/pkg1"
)

func TestGenTestWalks(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []pkg1.Tracker
	}{
		{
			name: "count is zero",
			args: args{
				count: 0,
			},
			want: []pkg1.Tracker{{Walks: 2}},
		},
		{
			name: "count is two",
			args: args{
				count: 2,
			},
			want: []pkg1.Tracker{{Walks: 2}, {Walks: 3}, {Walks: 4}},
		},
		{
			name: "count is positive",
			args: args{
				count: 3,
			},
			want: []pkg1.Tracker{{Walks: 2}, {Walks: 3}, {Walks: 4}, {Walks: 5}},
		},
		{
			name: "count is large",
			args: args{
				count: 10,
			},
			want: []pkg1.Tracker{{Walks: 2}, {Walks: 3}, {Walks: 4}, {Walks: 5}, {Walks: 6}, {Walks: 7}, {Walks: 8}, {Walks: 9}, {Walks: 10}, {Walks: 11}, {Walks: 12}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenTestWalks(tt.args.count)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenTestWalks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenTestWalksInverted(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want []pkg1.Tracker
	}{
		{
			name: "count is zero",
			args: args{
				count: 0,
			},
			want: []pkg1.Tracker{{Walks: -2}},
		},
		{
			name: "count is two",
			args: args{
				count: 2,
			},
			want: []pkg1.Tracker{{Walks: -2}, {Walks: -1}, {Walks: 0}},
		},
		{
			name: "count is positive",
			args: args{
				count: 3,
			},
			want: []pkg1.Tracker{{Walks: -2}, {Walks: -1}, {Walks: 0}, {Walks: 1}},
		},
		{
			name: "count is large",
			args: args{
				count: 10,
			},
			want: []pkg1.Tracker{{Walks: -2}, {Walks: -1}, {Walks: 0}, {Walks: 1}, {Walks: 2}, {Walks: 3}, {Walks: 4}, {Walks: 5}, {Walks: 6}, {Walks: 7}, {Walks: 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenTestWalksInverted(tt.args.count)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenTestWalksInverted() = %v, want %v", got, tt.want)
			}
		})
	}
}
