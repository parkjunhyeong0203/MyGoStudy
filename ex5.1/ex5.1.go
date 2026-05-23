package main

import "fmt"

func main() {
	var a int = 10
	var b = 20
	var f = 327991448743.8297
	var f64 float64 = 327991448743.8297
	var f32 float32 = 327991448743.8297
	fmt.Printf("float32 1e20 with %%f: %f\n", f32)
	fmt.Printf("float32 1e20 with %%g: %g\n", f32)

	fmt.Println("float64 with Println:", f64)
	fmt.Println("float32 with Println:", f32)
	fmt.Printf("float64 with %%f: %f\n", f64)
	fmt.Printf("float32 with %%f: %f\n", f32)
	fmt.Printf("float64 with %%g: %g\n", f64)
	fmt.Printf("float32 with %%g: %g\n", f32)

	fmt.Print("a: ", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b) // 입력값들 사이사이 빈칸 넣어줌
	fmt.Printf("a: %d b: %d f: %f\n", a, b, f)
}
