package main

import "fmt"

func testString() {
	var str = "hello"
	fmt.Printf("str[0]=%c\n", str[0])

	for index, val := range str {
		fmt.Printf("str[%d]=%c\n", index, val)
	}

	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = '0'
	str = string(byteSlice)
	fmt.Println(str)

	fmt.Printf("len(str)=%d\n", len(str))

	str = "hello, 鲁班七号"
	fmt.Printf("len(str)=%d\n", len(str))

	var b rune = '中'
	fmt.Printf("b=%c\n", b)

	var runeSlice []rune

	runeSlice = []rune(str)
	fmt.Printf("str长度：%d\n", len(runeSlice))
}

// 字符反转
func testReverseStringV1() {
	var str = "hello"
	var bytes []byte = []byte(str)

	for i := 0; i < len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}
	str = string(bytes)
	fmt.Println(str)
}

// 中文字符反转
func testReverseStringV2() {
	var str = "hello中国"
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	str = string(r)
	fmt.Println(str)
}

// 检查是否是回文
func testHuiWen() {
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}
	str2 := string(r)
	if str2 == str {
		fmt.Println(str, "is huiwen")
	} else {
		fmt.Println(str, "is not huiwen")
	}
}

func main() {
	// testString()
	// testReverseStringV1()
	// testReverseStringV2()
	testHuiWen()
}
