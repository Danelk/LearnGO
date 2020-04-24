package slice_demo

import "fmt"

func MakeSlice()  {
	s := make([]string, 3)
	s[0] = "123"
	s[1] = "234"
	s[2] = "345"
	fmt.Println(s)
}

func ExtentSlice(s []int, f int, l int) []int {
	return s[f:l]
}

func AppendSlice(s []int, i int) []int {
	return append(s, i)
}

func CopySlice(s1 []int, s2 []int) []int {
	copy(s1, s2)
	return s1
}

func DeleteSlice(s []int, x int, op string) ([]int, error) {
	switch op {
	case "f":
		return s[x:], nil //前删除
	case "b":
		return s[:x], nil //后删除
	default:
		return nil, fmt.Errorf("Can not find opreation:" + op)

	}
}

// 最长不重复子串 start   lastOccurred[x]  对于每个字母x  lastOccurred[x]不存在或者<start 不操作  lastOccurred[x]>=start 更新start  更新lastOccurred[x]
func LengthOfNonRepeatingSubStr(s string) int  {
	lastOccurred := make(map[byte]int)  //创建字节map
	start := 0
	maxLength := 0
	for i, ch := range []byte(s){
		lastI, ok := lastOccurred[ch]
		fmt.Println("lastI=", lastI)
		fmt.Println("ok=", ok)
		if  ok && lastI >= start{
			start = lastI + 1
			fmt.Println("start=", start)
		}
		if i - start+1 > maxLength{
			maxLength = i-start+1
			fmt.Println("maxLength=", maxLength)
		}
		lastOccurred[ch] = i
		fmt.Printf("ch= %X", ch)
		fmt.Println()
	}
	return maxLength
}