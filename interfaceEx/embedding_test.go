package interfaceex

import (
	"fmt"
	"testing"
)

type People struct {
	Name string
	Age  int
}

type People2 struct {
	Name string
	Age2 int
}

type Student struct {
	People
	People2
}

func (p People) PrintAll() {
	p.PrintName()
	p.PrintAge()
}

func (p People) PrintName() {
	fmt.Printf("People Name %+v\n", p.Name)
}

func (p People) PrintAge() {
	fmt.Printf("People Age %+v\n", p.Age)
}

func (p People2) PrintName() {
	fmt.Printf("People2 Name %+v\n", p.Name)
}

func (p People2) PrintAge2() {
	fmt.Printf("People2 Age %+v\n", p.Age2)
}

func (s Student) PrintAge() {
	fmt.Printf("Student Age %+v\n", s.Age)
}

type IPeople interface {
	PrintAge()
	// PrintName() //error
}

func Print(a IPeople) {
	a.PrintAge()
}

func Test_embedding1(t *testing.T) {
	s := Student{People{Name: "name1", Age: 1}, People2{Name: "name2", Age2: 2}}

	//fmt.Println(s.Name) //衝突
	fmt.Println(s.People.Name)
	fmt.Println(s.People2.Name)

	fmt.Println(s.Age)
}

func Test_embedding2(t *testing.T) {
	s := Student{People{Name: "name1", Age: 1}, People2{Name: "name2", Age2: 2}}

	// s.PrintName() //衝突, 找不到
	s.People.PrintName()
	s.People2.PrintName()

	s.PrintAge()
	s.PrintAge2()

	Print(s)
}

func Test_embedding3(t *testing.T) {
	s := Student{People{Name: "name1", Age: 1}, People2{Name: "name2", Age2: 2}}
	s.PrintAge()
	s.PrintAll() //還是撈People的Age
}
