package testunsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))
	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}

func TestPointer(t *testing.T) {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}
}

func TestChangeValue(t *testing.T) {
	var a = 3
	a2 := (*float64)(unsafe.Pointer(&a))
	fmt.Println(*a2)
	*a2 = 3.5
	fmt.Printf("a value %v", a)
	fmt.Println(a2)
	fmt.Println(&a)
}

func TestOffset(t *testing.T) {
	array := []int{1, 2, 3, 4}
	//获取array第一个元素的地址,加上偏移量,就能得到第三个元素的地址,然后再转换成Pointer
	a1 := unsafe.Pointer(uintptr(unsafe.Pointer(&array[0])) + 2*unsafe.Sizeof(array[1]))
	fmt.Println(*(*int)(a1))
	//重新赋值
	*(*int)(a1) = 5
	fmt.Println("===========")
	for _, v := range array {
		fmt.Println(v)
	}
}

//结构体中第一个成员的地址就是这个结构体的地址
func TestStruct(t *testing.T) {
	person := Person{name: "Nicholas", age: 18}
	pa := unsafe.Pointer(&person)
	personName := (*string)(unsafe.Pointer(uintptr(pa) + unsafe.Offsetof(person.name)))
	personAge := (*int)(unsafe.Pointer(uintptr(pa) + unsafe.Offsetof(person.age)))
	*personName = "Jasper2"
	*personAge = 28
	fmt.Println(person) // {Jasper2 28}
}

type Person struct {
	name string
	age  int
}
