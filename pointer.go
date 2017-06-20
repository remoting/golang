package main

import "fmt"

func TTT() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

func TTT2() {
	var ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr)
	if ptr == nil {
		fmt.Printf("ptr is null")
	} else {
		fmt.Printf("ptr not is null")
	}
}

func TTT3() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap2(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func swap1(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func swap2(x *int, y *int) {
	var temp *int
	temp = x
	x = y
	y = temp
}

type Person struct {
	Name    string
	Gender  string
	Age     uint8
	Address string
}

func (person *Person) Grow() {
	person.Age++
}

func (person *Person) Move(newAddres string) string {
	old := person.Address
	person.Address = newAddres
	return old
}
