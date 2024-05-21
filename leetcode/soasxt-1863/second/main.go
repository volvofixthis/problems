package first

import (
	"fmt"
	"sort"
)

type hashKey [12]int

type hashMap map[hashKey][]int

var u = hashMap{}

func hashResult(a []int) hashKey {
	hash := hashKey{}
	for i := 0; i < len(hash); i++ {
		hash[i] = -1
	}
	copy(hash[:], a)
	sort.Ints(hash[:])
	return hash
}

func backtrack(nums []int, k, start int, current []int) {
	if len(current) > 0 {
		hash := hashResult(current)
		if _, ok := u[hash]; !ok {
			u[hash] = current
		}
	}
	if k == 0 {
		return
	}
	for i := start; i < len(nums); i++ {
		backtrack(nums, k-1, i+1, append(current, i))
	}
}

func subsetXORSum(nums []int) int {
	u = hashMap{}
	backtrack(nums, len(nums), 0, []int{})
	sum := 0
	for _, v := range u {
		x := 0
		for i := 0; i < len(v); i++ {
			fmt.Print(v[i], " ")
			x ^= nums[v[i]]
		}
		fmt.Println()
		fmt.Println("xor", x)
		sum += x
	}
	fmt.Println()
	return sum
}
