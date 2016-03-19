package datastructure

import "fmt"

type person struct {
	name string
	age  int
}

type teacher struct {
	person
	speciality string
	string //任何内置类型都可以作为匿名函数
}

func (p *person) isYoung() bool {
	if p.age < 15 {
		return true
	}

	return false
}

func (p person) setAge(a int) {
	p.age = a
}

func (t teacher) isHistoryTeacher() bool {
	if t.speciality == "history" {
		return true
	}else {
		return false
	}
}


type numberList []int

func (s numberList) countSum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func compareAge(personA *person, personB *person) int {
	if personA.age > personB.age {
		return personA.age - personB.age
	}else {
		return personB.age - personA.age
	}
}

func DemoStruct() {
	fmt.Println(person{"me", 1})

	fmt.Println(person{name:"namedV", age:12})

	fmt.Println(&person{name:"namedP", age:12})

	s := &person{name:"alice", age:26}
	fmt.Println(s.age)

	diffAge := compareAge(s, &person{age:1})
	fmt.Println(diffAge)

	p := new(person)
	p.age = 2
	p.name = "pointer"
}

func DemoMethod() {
	l := numberList{1, 2, 3}
	fmt.Println("sum:", l.countSum())

	teacherP := teacher{person:person{name:"x", age:19}, speciality:"history", string:"oh"}
	teacherP.string = "oh,my"
	teacherP.name = "y"
	teacherP.setAge(10) //note, if the receiver is not pointer, the content can not be changed.
	fmt.Println(teacherP)

	fmt.Println("Is history:", teacherP.isHistoryTeacher())
	fmt.Println("Is Young:", teacherP.isYoung())
}
