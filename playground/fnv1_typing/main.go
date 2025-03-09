// package main
//
// import "fmt"
//
// const (
// 	FNVOffset uint32 = 2166136261
// 	FNVPrime  uint32 = 16777619
// )
//
// func fnv1(a []byte) uint32 {
// 	hash := FNVOffset
// 	for _, v := range a {
// 		hash = hash * FNVPrime
// 		hash ^= uint32(v)
// 	}
// 	return hash
// }
//
// func main() {
// 	fmt.Println(fnv1([]byte("kua")))
// 	fmt.Println(fnv1([]byte("kub")))
// 	fmt.Println(fnv1([]byte("kuc")))
// }

package main

import "fmt"

const (
	FNVOffset uint32 = 2166136261
	FNVPrime  uint32 = 16777619
)

func fnv1(a []byte) uint32 {
	hash := FNVOffset
	for _, v := range a {
		hash = hash * FNVPrime
		hash ^= uint32(v)
	}
	return hash
}

func main() {
	fmt.Println(fnv1([]byte("kua")))
	fmt.Println(fnv1([]byte("kub")))
	fmt.Println(fnv1([]byte("kuc")))
}
