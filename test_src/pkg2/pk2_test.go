
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
			name: "count is positive",
			args: args{
				count: 2,
			},
			want: []pkg1.Tracker{{Walks: 2}, {Walks: 3}, {Walks: 4}},
		},
		{
			name: "count is large",
			args: args{
				count: 10,
			},
			want: []pkg1.Tracker{{Walks: 2}, {Walks: 3}, {Walks: 4}, {Walks: 5}, {Walks: 6}, {Walks: 7}, {Walks: 8}, {Walks: 9}, {Walks: 10}, {Walks: 11}, {Walks: 12}},
		},
		{
			name: "count is negative",
			args: args{
				count: -1,
			},
			want: []pkg1.Tracker{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenTestWalks(tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenTestWalks() = %v, want %v", got, tt.want)
			}
		})
	}
}
