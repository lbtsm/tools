package slicehelper

import (
	"reflect"
	"testing"
)

func TestSplitStringSlice(t *testing.T) {
	type args struct {
		object []string
		size   int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "standard",
			args: args{
				object: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
				size:   10,
			},
			want: [][]string{{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitSlice(tt.args.object, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
