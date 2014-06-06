package helpers

import (
	"strconv"
	"strings"
)

func TriangleMaxSum(triangle []string) int {
	maxList := make(map[int][]int, 0)
	i := 0

	for _, str := range triangle {
		maxList[i] = make([]int, 0)
		lines := strings.Split(str, " ")
		for _, ele := range lines {
			j, _ := strconv.Atoi(ele)
			maxList[i] = append(maxList[i], j)
		}

		var leftParent int
		var rightParent int
		for j, _ := range maxList[i] {
			if _, ok := maxList[i-1]; ok {
				leftParent = 0
				rightParent = 0
				if j != 0 {
					leftParent = maxList[i-1][j-1]
				}
				if j <= i-1 {
					rightParent = maxList[i-1][j]
				}
				if leftParent > rightParent {
					maxList[i][j] += leftParent
				} else {
					maxList[i][j] += rightParent
				}
			}
		}
		i++
	}

	max := 0
	for _, ele := range maxList[i-1] {
		if ele > max {
			max = ele
		}
	}
	return max
}
