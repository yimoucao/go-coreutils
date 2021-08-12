package ranges

import (
	"math"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		range_ string
	}
	tests := []struct {
		name    string
		args    args
		want    *Range
		wantErr bool
	}{
		{
			"123-",
			args{range_: "123-"},
			&Range{range_: "123-", segs: [][2]int{{123, math.MaxInt64}}},
			false,
		},
		{
			"123",
			args{range_: "123"},
			&Range{range_: "123", segs: [][2]int{{123, 123}}},
			false,
		},
		{
			"-123",
			args{range_: "-123"},
			&Range{range_: "-123", segs: [][2]int{{math.MinInt64, 123}}},
			false,
		},
		{
			"123-456",
			args{range_: "123-456"},
			&Range{range_: "123-456", segs: [][2]int{{123, 456}}},
			false,
		},
		{
			"empty range",
			args{range_: ""},
			&Range{range_: "", segs: nil},
			false,
		},
		{
			"a",
			args{range_: "a"},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.range_)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Logf("Parse() error = %v", err)
			} else {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Parse() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
