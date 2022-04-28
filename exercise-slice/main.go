package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		idx, ok := m[target-v]
		if ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := twoSum(nums, 9)
	for _, val := range res {
		fmt.Println(val)
	}
}
