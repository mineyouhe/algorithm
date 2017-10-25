package main

import "fmt"

type RInterval struct {
	left  int
	right int
}
type RangeModule struct {
	Include map[int]int
}

func Constructor() RangeModule {
	r := RangeModule{}
	r.Include = make(map[int]int)
	return r
}

//两个记录六种关系，add,remove就是动态调整map中记录的范围
func (this *RangeModule) AddRange(left int, right int) {
	exclude := make([]int, 0)
	min, max := left, right
	for k, v := range this.Include {
		if !(k > max || v < min) {
			if v <= max && k >= min {
				exclude = append(exclude, k)
				continue
			}
			if max >= k && max <= v {
				max = v
				exclude = append(exclude, k)
			}
			if min >= k && min <= v {
				min = k
				exclude = append(exclude, k)
			}
			if min <= k && max >= v {
				exclude = append(exclude, k)
			}
		}

	}
	for k, _ := range exclude {
		delete(this.Include, exclude[k])
	}
	this.Include[min] = max
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	for k, v := range this.Include {
		if left >= k && right <= v {
			return true
		}
	}
	return false
}

func (this *RangeModule) RemoveRange(left int, right int) {
	exclude := make([]int, 0)
	intervals := make([]RInterval, 0)
	for k, v := range this.Include {
		if left >= v || right <= k {
			continue
		}
		if left >= k && right <= v {
			this.Include[right] = this.Include[k]
			this.Include[k] = left
			break
		}
		if left <= k && right >= v {
			exclude = append(exclude, k)
			continue
		}
		if left >= k && left <= v {
			this.Include[k] = left
		}
		if left <= k && right >= k {
			exclude = append(exclude, k)
			interval := RInterval{}
			interval.left = right
			interval.right = this.Include[k]
			intervals = append(intervals, interval)
			continue
		}
	}
	for k, _ := range exclude {
		delete(this.Include, exclude[k])
		fmt.Println(this.Include, exclude[k])
	}
	for _, v := range intervals {
		this.Include[v.left] = v.right
	}
	fmt.Println(this.Include)
}

func main() {
	obj := Constructor()
	qus := [67]string{"queryRange", "queryRange", "addRange", "addRange", "queryRange", "queryRange", "queryRange", "removeRange", "addRange", "removeRange", "addRange", "removeRange", "removeRange", "queryRange", "queryRange", "queryRange", "queryRange", "removeRange", "addRange", "removeRange", "queryRange", "addRange", "addRange", "removeRange", "queryRange", "removeRange", "removeRange", "removeRange", "addRange", "removeRange", "addRange", "queryRange", "queryRange", "queryRange", "queryRange", "queryRange", "addRange", "removeRange", "addRange", "addRange", "addRange", "queryRange", "removeRange", "addRange", "queryRange", "addRange", "queryRange", "removeRange", "removeRange", "addRange", "addRange", "queryRange", "queryRange", "addRange", "addRange", "removeRange", "removeRange", "removeRange", "queryRange", "removeRange", "removeRange", "addRange", "queryRange", "removeRange", "addRange", "addRange", "queryRange"} //, "removeRange", "queryRange", "addRange", "addRange", "addRange", "addRange", "addRange", "removeRange", "removeRange", "addRange", "queryRange", "queryRange", "removeRange", "removeRange", "removeRange", "addRange", "queryRange", "removeRange", "queryRange", "addRange", "removeRange", "removeRange", "queryRange"}
	data := [67][2]int{{21, 34}, {27, 87}, {44, 53}, {69, 89}, {23, 26}, {80, 84}, {11, 12}, {86, 91}, {87, 94}, {34, 52}, {1, 59}, {62, 96}, {34, 83}, {11, 59}, {59, 79}, {1, 13}, {21, 23}, {9, 61}, {17, 21}, {4, 8}, {19, 25}, {71, 98}, {23, 94}, {58, 95}, {12, 78}, {46, 47}, {50, 70}, {84, 91}, {51, 63}, {26, 64}, {38, 87}, {41, 84}, {19, 21}, {18, 56}, {23, 39}, {29, 95}, {79, 100}, {76, 82}, {37, 55}, {30, 97}, {1, 36}, {18, 96}, {45, 86}, {74, 92}, {27, 78}, {31, 35}, {87, 91}, {37, 84}, {26, 57}, {65, 87}, {76, 91}, {37, 77}, {18, 66}, {22, 97}, {2, 91}, {82, 98}, {41, 46}, {6, 78}, {44, 80}, {90, 94}, {37, 88}, {75, 90}, {23, 37}, {18, 80}, {92, 93}, {3, 80}, {68, 86}}                                                                                                                                                                                                                                                             //, {68, 92}, {52, 91}, {43, 53}, {36, 37}, {60, 74}, {4, 9}, {44, 80}, {85, 93}, {56, 83}, {9, 26}, {59, 64}, {16, 66}, {29, 36}, {51, 96}, {56, 80}, {13, 87}, {42, 72}, {6, 56}, {24, 53}, {43, 71}, {36, 83}, {15, 45}, {10, 48}}
	for k, v := range qus {
		switch v {
		case "addRange":
			obj.AddRange(data[k][0], data[k][1])
		case "removeRange":
			obj.RemoveRange(data[k][0], data[k][1])
		case "queryRange":
			fmt.Println(k, qus[k], data[k], obj.QueryRange(data[k][0], data[k][1]))
		}
	}
	fmt.Println(obj.Include)
}
