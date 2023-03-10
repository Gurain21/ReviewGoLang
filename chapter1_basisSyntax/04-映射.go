package main

import "fmt"

func main() {
	/**
	映射是go中lang中的内置数据类型，其数据结构表现为 key -->value 一一对应。
		语法声明为  map[Type]Type
	
			map中的keys的排序是无规律的，每次遍历出来的值都有可能不一样。
	*/
	
	//
	var map1 map[string]int //panic: assignment to entry in nil map
	
	//,map必须经由make函数初始化后才能够使用  panic: assignment to entry in nil map
	map1 = make(map[string]int)
	map1["小明"] = 18
	map1["小华"] = 19
	map1["小龙"] = 16
	map1["小光"] = 16
	fmt.Println(map1)
	
	//取值的语法2
	v, ok := map1["小明"]
	fmt.Println(v, ok)
	//如果map中没有我们取值的key，会返回其对应类型的零值
	v2, ok2 := map1["小红"]
	fmt.Println(v2, ok2)
	
	// len 函数能够输出map的长度
	fmt.Println(len(map1))
	
	//for range 遍历 map 类型
	for key, value := range map1 {
		fmt.Println(key, "---->", value)
	}
	
	map2 := make(map[int]int, 512)
	for i := 0; i < 100; i++ {
		map2[i] = i
	}
	
	fmt.Println(map2)
	
}
