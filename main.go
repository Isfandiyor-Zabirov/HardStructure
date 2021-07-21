package main

import "fmt"

type PhysicChar struct {
	Height int
	Weight int
	Color  string
}

type Phone struct {
	ID         int
	Model      string
	Name       string
	CameraPxl  int
	PhysicChar PhysicChar
	Memory     int
}

type Rectangle struct {
	Height int
	Length int
}

func Area(rec Rectangle) (area int) {
	area = rec.Length * rec.Height
	return area
}

func Perimeter(rec Rectangle) (perimeter int) {
	perimeter = 2 * (rec.Length + rec.Height)
	return perimeter
}

func (receiver *Phone) AddMemory() {
	receiver.Memory++

}

func (receiver *Phone) ChangeColor() {
	receiver.PhysicChar.Color = "Black"
}

func SumofTwoDigits(a int, b int) (sum int) {
	sum = a + b
	return sum

}

func main() {
	phone := Phone{
		ID:        12345,
		Model:     "Samsung",
		Name:      "Galaxy S5",
		CameraPxl: 12,
		PhysicChar: PhysicChar{
			1,
			2,
			"Red",
		},
		Memory: 1,
	}

	fmt.Println(phone)
	phone.AddMemory()
	fmt.Println(phone)
	phone.ChangeColor()
	fmt.Println(phone)

	a := 2
	b := 3
	/*SumofTwoDigits(a, b)
	SumofTwoDigits(5, 6)*/

	c := SumofTwoDigits(a, b)
	fmt.Println(c)

	rectangle := Rectangle{
		Height: 5,
		Length: 10,
	}

	area := Area(rectangle)
	perimeter := Perimeter(rectangle)
	fmt.Println(area)
	fmt.Println(perimeter)
}

//TODO: Rectangle sides a and b ---- make structure of the rectangle, make function to find Perimeter and Area
// Area = S = a * b, Perimeter = P = 2 * (a + b)
