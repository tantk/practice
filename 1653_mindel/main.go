package main

import "fmt"

//complexity of 2n result in 112 ms which is only 7.32% faster than go submission
//memmory usage 6.5mb scores a 87%
//had to rely on youtuber Programming Live with Larry for the approach to solve this problem

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minimumDeletions(s string) int {
	total := make(map[string]int, 0)
	for _, v := range s {
		total[string(v)]++
	}

	inLoop := false
	aLeft, bLeft := 0, 0
	aRight, bRight := total["a"], total["b"]
	best := bLeft + aRight

	for _, v := range s {
		inLoop = true
		if string(v) == "a" {
			aLeft++
			aRight--
		} else {
			bLeft++
			bRight--
		}
		best = min(best, bLeft+aRight)
	}
	if inLoop && best >= 1 {
		return best
	}

	return 0
}

func main() {
	fmt.Println(minimumDeletions("babbbbbbbababaababbbbbabaaabaabaabbbbbabbbaababababaaaabbbbaababbaabbbbbbaabbabababbbbbbaaaaaaaababababaababaaabaabbabbaaa"))
}
