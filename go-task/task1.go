package go_task

func main() {
	return
}

func SingleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

func singleNumber2(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func singleNumber3(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	sum := 0
	for _, num := range nums {
		sum += num * 2
	}
	return sum - res
}

func singleNumber4(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return (sum*2 - res) / 2
}

func isValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			if (ch == ')' && top != '(') || (ch == '}' && top != '{') || (ch == ']' && top != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isValid2(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, ')')
		} else if ch == '{' {
			stack = append(stack, '}')
		} else if ch == '[' {
			stack = append(stack, ']')
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isValid3(s string) bool {
	return false
}

func longestCommonPreFix(strs []string) string {
	s0 := strs[0]
	for j, ch := range s0 {
		for _, str := range strs {
			if j == len(str) || str[j] != byte(ch) {
				return s0[:j]
			}
		}
	}
	return s0
}

func plusOne(digits []int) []int {
	return []int{}
}

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func remoDuplicates2(nums []int) int {
	k := 0
	for _, num := range nums {
		if k == 0 || num != nums[k-1] {
			nums[k] = num
			k++
		}
	}
	return k
}

func twoSum(nums []int, target int) []int {
	precNums := make(map[int]int)

	for i, num := range nums {
		if j, ok := precNums[target-num]; ok {
			return []int{j, i}
		}
		precNums[num] = i
	}
	return []int{}
}
