package main

import (
	"fmt"
	"sort"
)

var sticks []int
var n, total int
var used []bool

func init() {
	n = 9
	sticks = []int{1, 1, 1, 2, 2, 2, 5, 5, 5}
	total = 24
	used = make([]bool, 9)
}

func main() {
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		used = make([]bool, n)
		sticks = make([]int, n)
		total = 0
		for i := 0; i < n; i++ {
			fmt.Scan(&sticks[i])
			total += sticks[i]
		}
		var st sort.IntSlice
		st = sticks
		st.Sort()
		poj_1011()
	}
}
func sum(index []int) int {
	sum := 0
	for _, i := range index {
		sum += sticks[i]
	}
	return sum
}
func poj_1011() {
	for i := 1; i < n; i++ {
		Combine(n, i, func(got []int) bool {
			if total-total/sum(got)*sum(got) == 0 {
				for j := 0; j < n; j++ {
					used[j] = false
				}
				return search(sum(got))
			}
			return false
		})
	}
}

func search(sum int) bool {
	if left_sum() == sum {
		return true
	}
	lastI := last_unused()
	if lastI < 0 {
		panic("last is -1")
	}
	used[lastI] = true
	need := sum - sticks[lastI]
	if !choose(need, func(got []int) bool {
		set_used(got, true)
		if !search(sum) {
			set_used(got, false)
			return false
		} else {
			return true
		}
	}) {
		return false
	}
	return true
}

func set_used(st []int, use bool) {
	for _, i := range st {
		used[i] = use
	}
}

func choose(need int, fc func([]int) bool) bool {
	max := max_index(need)
	if max < 0 {
		return false
	}
	if sticks[max] == need {
		return fc(append([]int{}, max))
	}
	return false
}

func combine_unused(need int, last int, ff func([]int) bool) bool {
	n := unused_until(last)
	for m := 2; m < n; m++ {
		if last_sum(last, m) < need {
			continue
		}else{

		}
	}
	return false
}

func last_sum(start, n int) int {
	ret := 0
	k := 0
	for i := start; i >= 0; i-- {
		if !used[i] {
			ret += sticks[i]
			k++
		}
		if k == n {
			break
		}
	}
	return ret
}

func unused_until(last int) (ret int) {
	for i := 0; i <= last; i++ {
		if !used[i] {
			ret++
		}
	}
	return
}

func max_index(nd int) int {
	for i := n - 1; i >= 0; i-- {
		if !used[i] && nd >= sticks[i] {
			return nd
		}
	}
	return -1
}

func last_unused() int {
	for j := n - 1; j >= 0; j-- {
		if !used[j] {
			return j
		}
	}
	return -1
}

func left_sum() int {
	return 0
}

func Combine(n, m int, check func([]int) bool) {
	i := 0
	a := make([]int, m)
	for {
		//到达终点
		if i == m-1 {
			if check(a) {
				break
			}
		} else {
			//fast next
			i++
			if a[i-1]+1 < n {
				a[i] = a[i-1] + 1
				continue
			}
		}

		more := false
		for j := a[i]; j < n; j++ {
			if sticks[j] != sticks[a[i]] {
				more = true
				break
			}
		}

		if !more { //no more, move bck
			for i >= 0 {
				pMore := false
				for j := a[i]; j < n; j++ {
					if sticks[j] != sticks[a[i]] {
						pMore = true
						break
					}
				}
				if !pMore {
					i--
				} else {
					break
				}
			}
			if i < 0 {
				break
			}
		}

		if i >= 0 {
			//move next
			nextI := 0
			for j := a[i]; j < n; j++ {
				if sticks[j] != sticks[a[i]] {
					nextI = j
					break
				}
			}
			a[i] = nextI
		}
	}
}
