package main

//refered to leetcode solution for the pointer approach which is the fastest
func merge(nums1 []int, m int, nums2 []int, n int) {
	//pointer for nums1, nums2,and combined array
	var p1, p2, pc int
	//comparison variables
	p1, p2, pc = m-1, n-1, m+n-1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] >= nums2[p2] {
			nums1[pc] = nums1[p1]
			p1--
		} else {
			nums1[pc] = nums2[p2]
			p2--
		}
		pc--
	}
	for p1 >= 0 {
		nums1[pc] = nums1[p1]
		p1--
		pc--
	}
	for p2 >= 0 {
		nums1[pc] = nums2[p2]
		p2--
		pc--
	}
}

//proper test scripts to be written
func main() {
	var l1, l2 []int
	// fmt.Println("test")
	// l1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	// l2 = []int{1, 2, 2}
	// merge(l1, 6, l2, 3)
	// fmt.Println("test")
	// l1 = []int{1, 2, 3, 0, 0, 0}
	// l2 = []int{2, 5, 6}
	// merge(l1, 3, l2, 3)
	// l1 = []int{1, 0}
	// l2 = []int{2}
	// merge(l1, 1, l2, 1)
	// l1 = []int{2, 0}
	// l2 = []int{1}
	// merge(l1, 1, l2, 1)
	// l1 = []int{1, 5, 7, 0, 0, 0, 0, 0}
	// l2 = []int{3, 16, 500, 5001, 5001}
	// merge(l1, 3, l2, 5)
	l1 = []int{4, 5, 6, 0, 0, 0}
	l2 = []int{1, 2, 3}
	merge(l1, 3, l2, 3)
}
