/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2021/1/5 13:54
 */
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type RuleMe []string

func (p RuleMe) Len() int           { return len(p) }
func (p RuleMe) Less(i, j int) bool { return p[i]+p[j] < p[j]+p[i] }
func (p RuleMe) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func minNumber(nums []int) string {
	arr := RuleMe{}
	for _, item := range nums {
		arr = append(arr, strconv.Itoa(item))
	}
	sort.Sort(arr)
	return strings.Join(arr, "")
}

func main() {
	number := minNumber([]int{12, 3})
	fmt.Println(number)
}
