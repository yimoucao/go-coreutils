package ranges

import (
	"math"
	"strconv"
	"strings"
)

// TODO: this is a simplist implementation. To improve: segments sort needed

type Range struct {
	range_ string   // str representation
	segs   [][2]int // both end inclusive
}

func (r *Range) Has(i int) bool {
	for _, seg := range r.segs {
		if i >= seg[0] && i <= seg[1] {
			return true
		}
	}
	return false
}

// LIST: one or more ranges separated by commas
// N
// N-
// N-M
// -M
func Parse(ranges string) (*Range, error) {
	var err error
	result := &Range{range_: ranges, segs: nil} // nil slice can be appended
	for _, seg := range strings.Split(ranges, ",") {
		points := strings.Split(seg, "-")
		if len(points) < 2 {
			if points[0] == "" { // we got empty string
				continue
			}
			// we got single num
			a, err := strconv.Atoi(points[0])
			if err != nil {
				return nil, err
			}
			result.segs = append(result.segs, [2]int{a, a})
		} else {
			var a, b int
			if points[0] == "" {
				a = math.MinInt64
			} else {
				a, err = strconv.Atoi(points[0])
				if err != nil {
					return nil, err
				}
			}
			if points[1] == "" {
				b = math.MaxInt64
			} else {
				b, err = strconv.Atoi(points[1])
				if err != nil {
					return nil, err
				}
			}
			result.segs = append(result.segs, [2]int{a, b})
		}
	}
	return result, nil
}
