package MethodSet

import "fmt"

type T struct{}

func (t T) M1() {
}

func (t *T) M2()  {
}

func (t *T) M3() {
}

func (t *T) M4()  {
}


func dumpMethodSet(i interface{})  {
	dynType := reflect.TypeOf(i)
	if dynType == nil {
		fmt.Println("there is no dynamic type")
		return
	}

	n := dynType.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method is empty", dynType)
	}

	fmt.Printf("%s's method set:\n", dynType)
	for i := 0; i < n; i++ {
		fmt.Println("-", dynType.Method(i).Name)
	}

	fmt.Println("")
}

func main()  {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)
}