package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("int size = ", unsafe.Sizeof(int(0)))
	fmt.Println("int8 size = ", unsafe.Sizeof(int8(0)))
	fmt.Println("int16 size = ", unsafe.Sizeof(int16(0)))
	fmt.Println("int32 size = ", unsafe.Sizeof(int32(0)))
	fmt.Println("int64 size = ", unsafe.Sizeof(int64(0)))
	fmt.Println("uint size = ", unsafe.Sizeof(uint(0)))
	fmt.Println("uint8 size = ", unsafe.Sizeof(uint8(0)))
	fmt.Println("uint16 size = ", unsafe.Sizeof(uint16(0)))
	fmt.Println("uint32 size = ", unsafe.Sizeof(uint32(0)))
	fmt.Println("uint64 size = ", unsafe.Sizeof(uint64(0)))
	fmt.Println("uintptr size = ", unsafe.Sizeof(uintptr(0)))
	fmt.Println("byte size = ", unsafe.Sizeof(byte(0)))
	fmt.Println("rune size = ", unsafe.Sizeof(rune(0)))
	fmt.Println("float32 size = ", unsafe.Sizeof(float32(0)))
	fmt.Println("float64 size = ", unsafe.Sizeof(float64(0)))
	fmt.Println("true size = ", unsafe.Sizeof(true))
	fmt.Println("false size = ", unsafe.Sizeof(false))

}
