package first

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	numsK1 := []int{}
	for _, v := range nums1 {
		if v%k == 0 {
			numsK1 = append(numsK1, v)
		}
	}
	n := 0
	for i := 0; i < len(numsK1); i++ {
		for j := 0; j < len(nums2); j++ {
			if numsK1[i]%(nums2[j]*k) == 0 {
				n++
			}
		}
	}
	return n
}
