package main

import "fmt"

type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

func main1()  {
	var t T = T	{a: 20}
	fmt.Println(t.a) // 20

	// T 作为参数类型，传递的是 t 的副本，所以在 M1 中修改 t.a，原 t 的值并没有变化
	t.M1()
	fmt.Println(t.a) // 20

	// *T 作为参数类型，传递的是 t 的指针
	pt := &t
	pt.M2()
	fmt.Println(t.a) // 10
}

func main()  {
	/*
	无论是 T 类型实例，还是 *T 类型实例，
	都既可以调用 receiver 为 T 类型的方法，
	也可以调用 receiver 为 *T 类型的方法
	*/

	var t1 T
	fmt.Println(t1.a)
	t1.M1()
	fmt.Println(t1.a)
	t1.M2() // 编译报错，不存在此函数
	fmt.Println(t1.a)

	var t2 = &T{}
	fmt.Println(t2.a)
	t2.M1()
	fmt.Println(t2.a)
	t2.M2()
	fmt.Println(t2.a)
}