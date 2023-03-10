package main

import "fmt"

/**
数组：用于存储一系列相同类型元素的集合，数组的长度一旦定义就无法改变，数组的长度是类型的一部分。
	 数组中的元素可以通过下标访问，下标index 从 0 开始 。
	 数组是值类型，在进行参数的传递时会复制一份新数据，不会修改到原来的数据元素。
	语法： [长度]数据类型{元素1,元素2, ...}
*/
func main() {
	array := [5]int{}
	array = [5]int{1, 2, 3, 4}
	array2 := [...]int{2, 3, 4}                       // 不指定长度，该为长度推断，根据赋值的元素个数自行给长度
	array3 := [10]string{2: "第3个数据元素", 9: "第10个数据元素"} // 通过index给指定的下标进行赋值
	fmt.Println(array, "\n", array2)
	fmt.Println(array3[2])
	
	//数组的遍历 两种方式 第一种 for循环
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
		
	}
	
	//第二种 for range
	for index, value := range array {
		fmt.Printf("下标 为 %v 的元素 值是 %v \n", index, value)
	}
	
	//数组扩容和复制通常建议使用slice切片，这样可以提高代码性能
	
}
