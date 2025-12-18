package homework01

import (
	"math/big"
	"sort"
	"strconv"
	"strings"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// 升序
	sort.Ints(nums)

	length := len(nums)
	// 一旦一个元素与左右出现不相等，即为只出现一次的数字
	var once int

	if length == 1 {
		once = nums[0]
	} else {
		for i := 0; i < length; i++ {
			if i == 0 {
				if nums[i] != nums[i+1] {
					once = nums[i]
					break
				}
			} else if i == length-1 {
				once = nums[i]
				break

			} else {
				if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
					once = nums[i]
					break
				}
			}
		}
	}
	return once
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// 负数必不可能是回文数
	if x < 0 {
		return false
	}
	// 转成string
	str := strconv.Itoa(x)
	// 转成切片
	runes := []rune(str)
	// 定义Builder
	var builder strings.Builder
	// 从后往前加
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	// 比较
	return str == builder.String()
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// 字符串转为切片
	runes := []rune(s)

	length := len(runes)
	// 字符数为奇数一定无效
	if length%2 != 0 {
		return false
	}
	// 循环剔除 () [] {}
	var result3 string
	for i := 1; i <= length; i++ {
		result1 := strings.ReplaceAll(s, "()", "")
		result2 := strings.ReplaceAll(result1, "[]", "")
		result3 = strings.ReplaceAll(result2, "{}", "")
	}
	return result3 == ""
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// 假设第一个就是最长公共前缀
	first := strs[0]
	// 从索引1开始遍历
	for i := 1; i < len(strs); i++ {
		// 用第一个元素跟后面去匹配，未匹配上去除最后一个字符，直到全部匹配为止或为空
		for !strings.HasPrefix(strs[i], first) {
			if len(first) == 0 {
				return ""
			}
			length := len([]rune(first))
			first = string([]rune(first)[0 : length-1])
		}
	}
	return first
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 转成string后拼接在转为int
	var str string
	for i := 0; i < len(digits); i++ {
		str = str + strconv.Itoa(digits[i])
	}
	// 加1(要使用bigint，题干中说了digits长度最大为100)
	num := new(big.Int)
	num.SetString(str, 10)
	num = new(big.Int).Add(num, big.NewInt(1))

	newStr := num.String()
	//初始化
	newDigits := make([]int, len(newStr))
	runes := []rune(newStr)
	for i := 0; i < len(runes); i++ {
		newDigits[i], _ = strconv.Atoi(string(runes[i]))
	}

	return newDigits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// nums 整体成递增 说明后一个元素 >= 前一个元素
	// 切片长度位1，直接返回1
	if len(nums) == 1 {
		return 1
	}
	// 切片长度大于1 [0,1,2,3,3,3,3,3,4] -> [0,1,2,3,4]
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] > nums[fast-1] {
			if slow != fast {
				nums[slow] = nums[fast]
			}
			slow++
		}
	}

	return slow
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 按照第一个元素进行升序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// [[1,5],[1,6],[1,4],[7,8],[9,10]] ->[[1,6],[7,8],[9,10]]
	mergeIntervals := [][]int{}
	j := 0
	for i := 1; i < len(intervals); i++ {
		if i == 1 {
			mergeIntervals = append(mergeIntervals, intervals[i-1])
		}

		if mergeIntervals[j][1] >= intervals[i][0] {
			// 存在交集
			if mergeIntervals[j][1] < intervals[i][1] {
				mergeIntervals[j][1] = intervals[i][1]
			}
		} else {
			// 不存在交集
			mergeIntervals = append(mergeIntervals, intervals[i])
			j++
		}
	}

	return mergeIntervals
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// 两数之和索引切片
	indexs := []int{}
	map1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			map1[target-nums[i]] = i
		} else {
			value, exists := map1[nums[i]]
			if exists {
				indexs = append(indexs, value, i)
				break
			} else {
				// 假设A+B = target，map中key为B，value为A在nums的索引位
				map1[target-nums[i]] = i
			}
		}
	}

	return indexs
}
