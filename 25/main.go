/* 
	选取 receiver 参数的原则
	1. 如果 go 方法要把 receiver 参数代表的类型实例的修改，反映到原类型实例上，那么就应该选择 *T 作为 receiver 参数的类型
	2. 如果 receiver 参数类型的 size 较大，以值传递形式传入会有较大性能开销，这时我们也应该选择 *T 作为 receiver 参数的类型
*/
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

func main2()  {
	/*
	这是 go 的语法糖
	编译器自动进行类型转换
	无论是 T 类型实例，还是 *T 类型实例，
	都既可以调用 receiver 为 T 类型的方法，
	也可以调用 receiver 为 *T 类型的方法
	*/

	var t1 T
	fmt.Println(t1.a)
	t1.M1()
	fmt.Println(t1.a)
	t1.M2()
	fmt.Println(t1.a)

	var t2 = &T{}
	fmt.Println(t2.a)
	t2.M1()
	fmt.Println(t2.a)
	t2.M2()
	fmt.Println(t2.a)
}


type Interface interface {
	M1()
	M2()
}

type T2 struct{}

func (t T2) M1() {
}

func (t *T2) M2()  {
}

func main()  {
	var t T
	var pt *T
	var i Interface

	i = pt
	i = t
}