package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
  切片slice ： 能够存储一系列相同类型数据的的数据集合，是对数组的引用，容量不足时slice会自动扩容
*/
func main() {
	array := [5]int{1, 3, 2, 4, 5}
	
	//三种切出切片的方式   [start : end)
	slice1 := array[1:3]
	slice2 := array[:3]
	slice3 := array[2:]
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	
	// make 函数可用于创建 切片、映射、通道，与new不同的是，make返回类型本身，而不是对应的指针。
	slice4 := make([]int, 8, 16)
	fmt.Printf("slice4 的长度为 %v,容量的初始化值是：%v\n", len(slice4), cap(slice4))
	fmt.Println(slice4)
	
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	for i := 0; i < 10; i++ {
		n := r.Intn(10)
		
		slice4 = append(slice4, n)
		fmt.Printf("slice4 的长度为 %v,容量的值是：%v\n", len(slice4), cap(slice4))
		
	}
	/**
	slice在扩容时的：
	
	当前需要容量如果小于当前的容量，无需扩容；
	
	当前需要容量大于当前容量2倍，直接扩容至当前容量
	
	如果需要的容量小于256，直接就是2倍扩容
	
	如果需要的容量大于256，则1.25x扩容直到容量足够
	
	实际扩容的容量会根据内存分配页优化，奇数容量会分配偶数容量
	
	*/
	
	// append 向切片中追加一个切片
	slice4 = append(slice4, slice4...)
	fmt.Printf("slice4 的长度为 %v,容量的值是：%v\n", len(slice4), cap(slice4))
	
	slice5 := make([]int, 0, 10)
	i := copy(slice5, slice4[:])
	
	fmt.Println(i)
	fmt.Println("slice5 的值：", slice5)
	fmt.Printf("slice5 的长度为 %v,容量的值是：%v\n", len(slice5), cap(slice5))
	
	slice6 := make([]int, 0)
	for _, v := range slice5 {
		slice6 = append(slice6, v)
	}
	
	fmt.Println("slice6 ： ", slice6)
}
