package main

import "fmt"

func singlIn(in int) int {
	return in
}

func multIn(a, b int, c int) int {
	return a + b + c
}

func namedReturn() (out int) {
	out = 2
	return
}

func multipleReturn(in int) (int, error) {
	if in > 2 {
		return 0, fmt.Errorf("some error happend")
	}
	return in, nil
}

func multipleNamedReturn(ok bool) (rez int, err error) {
	rez = 1
	if ok {
		err = fmt.Errorf("some error happend")
		return
	}
	rez = 2
	return
}

func sum(in ...int) (result int) {
	fmt.Printf("in := %#v \n", in)
	for _, val := range in {
		result += val
	}
	return
}

func main() {

	arr3 := make([]int, 2, 3)  //[]

	sl1 := arr3 //sl1 изменяется и arr3
	sl1[0] = 9 //arr

	sl1 = append(sl1, 1, 3)//добавил 1 элемент

	sl1[1] = 5

	fmt.Println(arr3)
	fmt.Println(sl1, m1)
	return
}

