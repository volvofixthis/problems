package first

import "sort"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	numsK1 := []int{}
	m := map[int]int{}
	for _, v := range nums1 {
		if v%k == 0 {
			m[v] += 1
		}
	}
	for k := range m {
		numsK1 = append(numsK1, k)
	}
	sort.Ints(numsK1)
	m2 := map[int]int{}
	for _, v := range nums2 {
		m2[v] += 1
	}
	n := int64(0)
	for c, v := range m2 {
		d := c * k
		for i := len(numsK1) - 1; i >= 0; i-- {
			if numsK1[i] < d {
				break
			} else if numsK1[i]%d == 0 {
				n += int64(m[numsK1[i]] * v)
			}
		}
	}
	return n
}
