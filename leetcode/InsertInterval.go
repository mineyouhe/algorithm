/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
 //解题思路跟RangeModule中addrange函数类似，只不过要考虑下排序问题
 func insert(intervals []Interval, newInterval Interval) []Interval {
	nIntervals := make([]Interval, 0)
	afterIntervals := make([]Interval, 0)
	if len(intervals) > 0 {
			t := intervals[0]
			if newInterval.End < t.Start {
					nIntervals = append(nIntervals, newInterval)
					nIntervals = append(nIntervals, intervals...)
					return nIntervals
			}
			if newInterval.End == t.Start {
					intervals[0].Start = newInterval.Start
					return intervals
			}
	} else {
			nIntervals = append(nIntervals, newInterval)
			return nIntervals
	}
   
	   
	min, max := newInterval.Start, newInterval.End
	exist := false
	for _, v := range intervals {
			if !(v.Start > max || v.End < min) {
					if min <= v.Start {
							exist = true
					}
					if max >= v.Start && max <= v.End {
							max = v.End
							exist = true
					}
					if min >= v.Start && min <= v.End {
							min = v.Start
							exist = true
					}
   
			} else {
					if v.Start >= min {
							exist = true
					}
					t := Interval{}
					t.Start = v.Start
					t.End = v.End
					if exist {
							afterIntervals = append(afterIntervals, t)
					} else {
							nIntervals = append(nIntervals, t)
					}
   
			}
	}
	t := Interval{}
	t.Start = min
	t.End = max
	nIntervals = append(nIntervals, t)
	nIntervals = append(nIntervals, afterIntervals...)
	return nIntervals

}    