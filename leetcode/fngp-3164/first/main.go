package first

import "math"

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	m := map[int]int{}
	for _, v := range nums1 {
		if v%k == 0 {
			m[v/k] += 1
		}
	}
	m2 := map[int]int{}
	for _, v := range nums2 {
		m2[v] += 1
	}
	n := int64(0)
	for k := range m {
		for i := 1; i <= int(math.Sqrt(float64(k))); i++ {
			if (k % i) == 0 {
				ans := k / i
				if v, ok := m2[i]; ok {
					n += int64(m[k] * v)
				}
				if v, ok := m2[ans]; ok && i != ans {
					n += int64(m[k] * v)
				}
			}
		}
	}
	return n
}
