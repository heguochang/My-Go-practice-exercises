package main

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
func main() {
	rectangle := &Rectangle{}
	rectangle.Area()
	rectangle.Perimeter()

	c := &Circle{}
	c.Area()
	c.Perimeter()
}

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r *Rectangle) Area() {

}

func (r *Rectangle) Perimeter() {

}

type Circle struct {
}

func (c *Circle) Area() {

}

func (c *Circle) Perimeter() {

}
