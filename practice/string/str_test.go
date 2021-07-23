package string

import (
	"fmt"
	"strconv"
)

func Example_strCat(){
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_strconv(){
	var i int
	var k int64
	var f float64
	var s string
	var err error
	i, err = strconv.Atoi("350")
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	f, err = strconv.ParseFloat("3.14", 64)
	s = strconv.Itoa(340)
	s = strconv.FormatInt(13402077, 16)

	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(err)

	// Output:
	// 350
	// 13402077
	// 3.14
	// cc7fdd
	// <nil>
}

func Example_fmt(){
	var num int
	fmt.Sscanf("57", "%d", &num)

	var s string
	s = fmt.Sprint(3.14)
	s = fmt.Sprintf("%x", 13402077)

	fmt.Println(num)
	fmt.Println(s)

	// Output:
	// 57
	// cc7fdd
}