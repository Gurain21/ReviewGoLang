package main

import (
	"errors"
	"fmt"
)

/**
变量 （variable） 程序在运行过程中可能会发生变化的量，可以先声明，程序运行过程中再赋值。
	关键字 var
	语法规则：
		完整声明定义写法：
			var 变量名 数据类型 = 值
		类型自动推断写法：
			var 变量名 = 值
		简短变量定义写法：
			变量名 := 值
		多变量同时定义写法：



常量 （constant） 程序运行过程中恒定不变的量，声明时就必须为其赋值。
	语法规则：
			const 常量名 数据类型 = 值


*/
func main() {
	//	var 变量名 数据类型 = 变量值
	var number int = 10
	//在go语言中，变量一经声明就必须使用，否则无法通过编译，因为go是一门强类型静态编程语言，这里我们打印输出一下
	fmt.Println("number 的值为：", number) //number 的值为： 10
	
	//类型自动推断，GO语言能够自动根据变量的value值去推断它应该是哪个类型
	var str = "Hello,Golang!"
	fmt.Printf("str输出信息： %v\n", str)
	
	//简短方式赋值，只能用于大括号 { } 内部，用于声明局部变量
	number1, number2 := 1, 2
	fmt.Println("number1,number = ", number1, ",", number2)
	
	var (
		a, b byte
		c, d rune
		e    error
	)
	a, b, c, d = 'a', 'b', 'c', 'd'
	e = errors.New("这是一个错误提示信息，用于帮助程序员排错滴")
	fmt.Println(a, b, c, d, e)
	
	//const 常量名 数据类型 = 值
	//可以看到常量和变量的对比，常量定义后是可以不用的，编译器也不会报错。
	const π = 3.14
	
}
