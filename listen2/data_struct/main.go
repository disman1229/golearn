package main

import (
	"fmt"
	"strings"
)

func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	a = !a
	fmt.Println(a)

	var b bool
	b = true
	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}

	b = true
	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}

	// 格式化输出
	fmt.Printf("%t %t\n", a, b)
}

func testInt() {
	var a int8
	a = 18
	fmt.Println("a=", a)
	a = -12
	fmt.Println("a=", a)

	var b int
	b = 138134115664
	fmt.Println("b=", b)

	var c float64
	c = 5.6
	fmt.Println("c=", c)

	fmt.Printf("a=%d a=%x c=%f\n", a, a, c)
}

func testStr() {
	var a string = "a:hello"
	fmt.Println(a)

	b := a
	fmt.Println(b)

	c := "c:\nhello"
	fmt.Println(c)

	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)

	c = `c:\nhello`
	fmt.Println(c)

	d := `
		今夕何夕兮，搴舟中流。
		今日何日兮，得与王子同舟。
		蒙羞被好兮，不訾诟耻。
		心几烦而不绝兮，得知王子。
		山有木兮木有枝，心悦君兮君不知。
	`

	// 字符串长度len
	clen := len(d)
	fmt.Println(clen)

	// 字符串拼接
	d = d + d
	clen = len(d)
	fmt.Println(clen)

	d = fmt.Sprintf("%s%s", d, d)
	clen = len(d)
	fmt.Println(clen)

	// 字符串分割
	ips := "192.168.1.1;192.168.1.2"
	ipSlice := strings.Split(ips, ";")
	fmt.Printf("first ip :%s\n", ipSlice[0])
	fmt.Printf("second ip :%s\n", ipSlice[1])

	// 字符串包含
	result := strings.Contains(ips, "192.168.1.1")
	fmt.Println(result)

	// 字符串前缀后或者后缀判断
	str := "https://baidu.com"
	if strings.HasPrefix(str, "https") {
		fmt.Println("str is http url")
	} else {
		fmt.Println("str is not http url")
	}

	if strings.HasSuffix(str, ".com") {
		fmt.Println("str is .com url")
	} else {
		fmt.Println("str is not .com url")
	}

	index := strings.Index(str, "bai")
	fmt.Printf("baidu is index:%d\n", index)

	lastindex := strings.LastIndex(str, "du")
	fmt.Printf("du last index:%d\n", lastindex)

	var str1 []string = []string{"192.168.2.0", "192.168.2.1", "192.168.2.4"}
	result1 := strings.Join(str1, ";")
	fmt.Printf("result=%s\n", result1)
}

func main() {
	// testBool()
	// testInt()
	testStr()
}
