package main
import "fmt"
//my first program in Go 
func main(){
	/**
	*  escape sequence
	*/
	fmt.Println("Hello\tWorld!")
	fmt.Println("I am in Go programming World!")
	var x float64=64
	y:=42//infer
	var a,b,c=3,4,"foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("a is of Type %T\n",a)
	fmt.Printf("b is of Type %T\n",b)
	fmt.Printf("c is of Type %T\n",c)
	fmt.Printf("x is of type %T\n",x)
	fmt.Printf("y is of type %T\n",y)

	//Constant
	const LENGTH int=10
	const WIDTH int =5
	var area int

	area=LENGTH * WIDTH
	fmt.Printf("value of area: %d",area)

	// Array
	var n [10]int
	var i, j int

	for i=0;i<10;i++{
		n[i]=i+100
	}

	for j=0;j<10;j++{
		fmt.Printf("Element[%d]=%d\n",j,n[j])
	}

	//for loop
	sum :=0

	for i:=1;i<5;i++{
		sum++
	}
	fmt.Println("sum=",sum)

	su :=1
	for su<5{
		su *=2
	}
	fmt.Println("su=",su)

	//For loop

	strs := []string{"hello", "world"}

	for i,str:=range strs {

		fmt.Println(i,str)
	}

	//Strings are immutable in Go lang can be read but not modified

	s1:="hello"
	s2:="world"

	greeting:=s1+" " +s2
	fmt.Println(greeting)
	fmt.Println(s1[0])

	// character Encoding
	s:="世界"
	for i,value:=range s{
		fmt.Printf("index: %d, char: %c, unicode: %U\n",i,value,value)
	}

	fmt.Println(&s)

	// map
	m :=map[string]float64{}

	m["pi"]=3.1416

	fmt.Println(m["pi"])

	m1:=map[string]float64{
		"pi":3.146,
		"e":2.71,
	}
	fmt.Println(m1["e"])

	// copy struct
	// create a new Person object
	john := person{
		name: "John",
		age: 23,
	}
	// make a copy
	john2 := john
	// before change
	fmt.Println(john)
	john2.age = 25
	// after change
	fmt.Println(john2)
	fmt.Println(john)
	
	// pass By value
	x1 := 42
	fmt.Println("Before:", x1)
	passByReference(x1)
	fmt.Println("After:", x1)     
	
	// pass By reference
	x2 := 42
	fmt.Println("Before:", x2)
	passByReference1(&x2)
	fmt.Println("After:", x2)
	                          
	// interface In Go, an Interface is a type that defines a set of methods
	r := rectangle{width: 10, height: 5}
	fmt.Println("Area:", calArea(r))
	
	// Composition (Embedding) While there is no inheritance in Go, you can use composition (or embedding) to reuse
	// fields and methods of an existing struct in a new struct
	e := employee{
		person1: person1{
			name: "John",
			dob: 1980,
		},
		company: "NUS",
	}
	e.intro()
}
	func passByReference(x1 int) {
		// de-reference the pointer
		// to access the value
		x1 = 43
	}
	func passByReference1(x *int) {
	// de-reference the pointer
	// to access the value
	*x = 43
	}
	type person struct {
	name string
	age int
	}
	type shape interface {
	area() float64
	}
	type rectangle struct {
	width, height float64
	}
	func (r rectangle) area() float64 {
	return r.width * r.height
	}
	func calArea(s shape) float64 {
		return s.area()
	}
	type person1 struct {
		name string
		dob int
	}
	type employee struct {
	person1
	company string
	}
	func (e employee) intro() {
		fmt.Printf("I am %s, born in %d and work at %s.\n",e.name, e.dob, e.company)
	}