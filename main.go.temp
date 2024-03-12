package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var age = 20
	const one = 1
	const two = 2
	const three = 3
	var empty string
	var zero int8
	zero = 0
	// var done bool
	var first, second = "first", "second"
	fmt.Println("this", "that", "there")
	fmt.Printf("Hello %d\n", age)
	fmt.Printf("%d, %d, %d, %d\n", zero, one, two, three)
	fmt.Printf("Empty is \"%s\" and is %d long\n", empty, len(empty))
	fmt.Println(first, second)
	fmt.Printf("%s has prefix \"fi\"? %t\n", first, strings.HasPrefix(first, "fi"))
	var arr = [...]string{"what", "the", "fuck"}
	fmt.Println(arr[2])
	arrCopy := arr
	arrCopy[2] = "hell"
	var arrSlice []string
	arrSlice = arr[:2]
	fmt.Println(arrSlice)
	arrSlice2 := make([]string, len(arrSlice))
	copy(arrSlice2, arrSlice)
	arrSlice2[len(arrSlice2)-1] = "in"
	arrSlice3 := make([]string, 0, len(arrSlice2))
	arrSlice3 = append(arrSlice3, arrSlice2...)
	arrSlice3[len(arrSlice3)-1] = "on"
	arrSlice3 = append(arrSlice3, "earth")
	fmt.Println(arrSlice2)
	fmt.Println(arrSlice3)
	fmt.Printf("%v.\n", arrCopy)
	for i := 0; i < len(arrCopy); i++ {
		fmt.Println(arrCopy[i])
	}
	numMap := make(map[string]int)
	numMap["one"] = 1
	fmt.Println(numMap)
	numMap1 := map[string]int{}
	numMap1["two"] = 2
	fmt.Printf("Two is %d.\n", numMap1["two"])
	{
		i := 0
		var arr []string
		arr = append(arr, arrSlice3...)
		for i < len(arr) {
			fmt.Print(arr[i])
			i++
			if i < len(arr) {
				fmt.Print(" ")
			} else {
				fmt.Print(".")
				break
			}
		}
		fmt.Println()
	}
	{
		var huh string
		for _, v := range arrCopy {
			huh += v + "."
		}
		fmt.Println(huh)
	}
	{
		type fullname struct {
			first string
			last  string
		}
		type person struct {
			name fullname
			age  int
		}
		bob := person{}
		bob.name = fullname{"Bob", "The Builder"}
		bob.age = 69
		fmt.Printf("Person is %s and is %d years old.\n", bob.name, bob.age)
	}
	{
		getHUH := func(num ...int) (string, string) {
			var builder strings.Builder
			huh := "HUH"
			if len(num) > 0 {
				for i := 0; i < num[0]; i++ {
					if i == 0 {
						builder.WriteString(fmt.Sprintf("%s[%d]", huh, i+1))
					} else {
						builder.WriteString(fmt.Sprintf(" %s[%d]", huh, i+1))
					}
				}
				return huh, builder.String()
			} else {
				return huh, builder.String()
			}
		}
		amount := 3
		amount2 := 5
		huh, str := getHUH(5, amount)
		fmt.Printf("%s %d times is %s.\n", huh, amount2, str)
	}
	{
		one := 1
		onePtr := &one
		oneVal := *onePtr
		fmt.Printf("One is %d.\n", oneVal)
		increment := func(num *int) {
			*num = *num + 1
		}
		increment(&one)
		fmt.Printf("One is %d.\n", one)
	}
	{
		str, len, err := addExclaimation[StringWithExclaimation, StringLen]("Hello")
		if err != nil {
			fmt.Printf("\"%s\" has %d characters (%s).\n", str, len, err)
		} else {
			fmt.Printf("\"%s\" has %d characters.\n", str, len)
		}
	}
	{
		blue := cat{speech: "Meow", name: "Blue"}
		SaySomething(blue)
	}
	{
		var bulldog Animal = &Dog{
			Fullname: Fullname{First: "Bull", Last: "Dog"},
			Color:    Color{Red: 0, Green: 0, Blue: 255},
			Age:      420,
		}
		fmt.Printf("%s is %v and is %d years old.\n", bulldog.GetName(), bulldog.GetColor(), bulldog.GetAge())
		bulldog.SetName(Fullname{"Golden", "Retriever"})
		fmt.Printf("%s is %v and is %d years old.\n", bulldog.GetName(), bulldog.GetColor(), bulldog.GetAge())
		dog, ok := bulldog.(*Dog)
		fmt.Printf("%s is %t.\n", dog.GetName(), ok)
	}
}

func (a cat) Say() {
	fmt.Printf("%s says %s.\n", a.name, a.speech)
}

type cat struct {
	name   string
	speech string
}

type Say interface {
	Say()
}

func SaySomething(o Say) {
	o.Say()
}

type StringWithExclaimation interface {
	GetString() string
}

type StringLen interface {
	GetLength() int
}

type myStringWithExclaimation string

func (s myStringWithExclaimation) GetString() string {
	return string(s)
}

type myStringLength int

func (l myStringLength) GetLength() int {
	return int(l)
}

func addExclaimation[stringWithExclaimation StringWithExclaimation, strLen StringLen](str string) (StringWithExclaimation, StringLen, error) {
	var err error
	if len(str) < 3 {
		err = errors.New("why so short?")
	}
	str = str + "!"
	return myStringWithExclaimation(str), myStringLength(len(str)), err
}

type Animal interface {
	GetName() string
	GetColor() Color
	GetAge() int
	SetName(fullname Fullname) string
}

type Fullname struct {
	First string
	Last  string
}

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

type Dog struct {
	Color
	Fullname
	Age int
}

func (d Dog) GetName() string {
	return fmt.Sprintf("%s %s", d.First, d.Last)
}

func (d *Dog) SetName(fullname Fullname) string {
	d.First = fullname.First
	d.Last = fullname.Last
	return fmt.Sprintf("%s %s", d.First, d.Last)
}

func (d Dog) GetColor() Color {
	return d.Color
}

func (d Dog) GetAge() int {
	return d.Age
}
