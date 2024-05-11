package first

func mergeTwoArrays(a, b []int) []int {
	cSize := len(a) + len(b)
	c := make([]int, cSize)
	ap := 0
	bp := 0
	for i := 0; i < cSize; i++ {
		if ap >= len(a) {
			c[i] = a[bp]
			bp++
		} else if bp >= len(b) {
			c[i] = a[ap]
			ap++
		} else {
			if a[ap] < b[bp] {
				c[i] = a[ap]
				ap++
			} else {
				c[i] = b[bp]
				bp++
			}
		}
	}
	return c
}

func median(a []int) float64 {
	if len(a)%2 == 1 {
		return float64(a[len(a)/2])
	} else {
		return float64((a[len(a)/2-1] + a[len(a)/2])) / 2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	c := mergeTwoArrays(nums1, nums2)
	return median(c)
}
