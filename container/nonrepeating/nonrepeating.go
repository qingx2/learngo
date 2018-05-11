package main

import "fmt"

// stores last occurred pos + 1.
// 0 means not seen
var lastOccurred = make(map[rune]int)

func lengthOfNonRepeatingSubStr(s string) int {
	// fmt.Println([]byte(s))

	// lastOccurred := make(map[rune]int)

	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	start := 0
	maxLength := 0

	// range 字符串时需要转换类型, 强制转换为 byte 不是英文时会出现问题
	for i, ch := range []rune(s) {
		lastI := lastOccurred[ch]
		// lastOccurred[ch]不存在, 或者 < start , 不会执行
		// lastOccurred[ch] >= start , 更新 start
		if lastI > start {
			start = lastI
		}
		// 更新 lastOccurred[ch], 更新maxLength
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1

		// fmt.Printf("last=%d\n", lastOccurred[ch])
		// fmt.Printf("start=%d, maxLength=%d\n", start, maxLength)
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdddd"))
	fmt.Println(lengthOfNonRepeatingSubStr("lllllll"))
	fmt.Println(lengthOfNonRepeatingSubStr("aisjdqoskjsdaq"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdddd"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("我离歌曲"))
	fmt.Println(lengthOfNonRepeatingSubStr("吧里bs防化服好"))
}
