package first

import (
	"crypto/md5"
	"fmt"
)

type hashMap map[[16]byte][]int

var u = hashMap{}
var result = [][]int{}

func convertIntToByte(a []int) []byte {
	r := []byte{}
	for _, v := range a {
		r = append(r, byte(v))
	}
	return r
}

func backtrack(nums []int, k, start int, current []int) {
	if len(current) >= 0 {
		hash := md5.Sum(convertIntToByte(current))
		// fmt.Println(hash)
		if _, ok := u[hash]; !ok {
			u[hash] = append([]int{}, current...)
		}
	}
	if len(current) == k {
		return
	}
	for i := start; i < len(nums); i++ {
		current = append(current, i)
		backtrack(nums, k, i+1, current)
		current = append([]int{}, current[1:]...)
	}
}

func subsetXORSum(nums []int) int {
	fmt.Println("xoring")
	u = hashMap{}
	backtrack(nums, len(nums), 0, []int{})
	sum := 0
	for _, v := range u {
		fmt.Println(v)
		x := 0
		for i := 0; i < len(v); i++ {
			x ^= nums[v[i]]
		}
		fmt.Println(x)
		sum += x
	}
	return sum
}
