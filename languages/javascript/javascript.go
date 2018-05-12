package languages

import (
	"fmt"
)

func main()  {
	fmt.Println("js main")
}

type Javascript struct {
	length, width int
}

func (t *Javascript) Foo() {
	fmt.Println("hel")
	fmt.Println(t.width)
}

//func Flags() {
//
//}
//
//func New () {
//	l := Javascript{
//		length: 2,
//		width: 4,
//	}
//
//	reflect.ValueOf(&l).MethodByName("Foo").Call([]reflect.Value{})
//	//l.Foo()
//
//	//return &l
//	fmt.Println(l)
//}