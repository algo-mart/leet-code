package toptal

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func Solution(message string, K int) string {
	var result string
	var i int
	temp := ""
	arr := strings.Split(message, "")
	myMap := make(map[string]bool)
	for _, s := range arr {
		myMap[s] = true
	}
	for i = 0; i < len(message) && i < K; i++ {
		temp += string(message[i])
		if string(message[i]) == (" ") || myMap[temp] {
			result += temp
			temp = ""
		}
	}
	return strings.Trim(result, " ")
}

var result = float64(0)

func Solution4(A []int) int {
	sum := float64(0)
	newArray := make([]float64, len(A))
	for _, i2 := range A {
		newArray = append(newArray, float64(i2))
		result += float64(i2)
	}
	return helper(newArray, sum, 0)
}
func helper(A []float64, sum float64, count int) int {
	sort.Float64s(A)
	A[len(A)-1] = A[len(A)-1] / 2
	sum += A[len(A)-1]
	count++
	if result-sum <= result/2 {
		return count
	}
	return helper(A, sum, count)
}

func Solution2(P []int, S []int) int {
	// write your code in Go 1.4
	var result int
	sumP := 0
	sumS := 0
	for i, i2 := range P {
		sumP += i2
		sumS += S[i]
	}
	if sumS <= sumP {
		return len(S)
	}

	sort.Ints(S)
	result = 1
	for i := len(S) - 1; i >= 0; i-- {
		if sumP <= S[i] {
			return result
		}
		sumP -= S[i]
		result++
	}
	return result
}

func Solution3(S string) int {
	next := ' '
	for i, s := range S {
		temp := i
		for rune(S[temp]) == '?' && temp < len(S)-1 {
			temp++
		}

		for j := i; j < temp; j++ {
			S = S[:j] + string(next) + S[j+1:]
			if next == 'a' && rune(S[j]+1) != 'b' {
				next = 'b'
			} else {
				next = 'a'
			}
		}
		if s == 'a' {
			next = 'b'
		} else {
			next = 'a'
		}
	}
	return 2
}

func searchSuggestions(repository []string, customerQuery string) [][]string {
	fmt.Println(repository, customerQuery)
	// Write your code here
	sort.Sort(sort.StringSlice(repository))
	result := make([][]string, 0)
	for i, char := range customerQuery {
		start := 0
		temp := make([]string, 0)
		for start < len(repository) {
			suggestion := repository[start]
			if i < len(suggestion) && char == rune(suggestion[i]) {
				fmt.Println(i, len(temp))
				if i > 0 && len(temp) < 3 {
					temp = append(temp, suggestion)
				}
			} else {
				repository = append(repository[:start], repository[start+1:]...)
			}
			start++
		}
		fmt.Println(repository, temp)
		if len(temp) > 0 {
			result = append(result, temp)
		}
	}
	return result
}

func lengthOfLongestSubstring(s string) int {
	myMap := map[rune]int{}
	i := 0
	max := 0
	for j, v := range s {
		if idx, _ := myMap[v]; idx > i {
			i = idx
		}
		myMap[v] = j + 1
		max = int(math.Max(float64(j+1-i), float64(max)))
	}
	return max
}

//func main() {
//	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
//	//fmt.Println(Solution3("aa?b"))
//	//fmt.Println(Solution([]int{3, 0, 5}))
//	//fmt.Println(Solution2([]int{3,1,2,1,2}, []int{3, 1, 1, 3, 3}))
//	//fmt.Println(Solution1("The quick brown fox jumps over the lazy dog", 39))
//}
