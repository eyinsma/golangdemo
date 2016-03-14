package datastructure
import (
	"math"
	"fmt"
	"strconv"
	"reflect"
)

type geometry interface{
	area() float64
	perim() float64
}

type shape interface{
	geometry
	another()
}

type circle struct{
	radius float64
}

type rect struct{
	width, height float64
}

func (r rect) String() string{
	return "rect:" + strconv.FormatFloat(r.height, 'f', 6, 64) + "," + strconv.FormatFloat(r.width, 'f', 6, 64)
}

func (c circle) String()string{
	return "circle:" + strconv.FormatFloat(c.radius, 'f', 6, 64)
}

func (r *rect) area() float64{
	return r.width * r.height
}

func (r *rect) perim() float64{
	return 2 * (r.width + r.height)
}

func (c *circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64{
	return 2* math.Pi*c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func DemoInterface(){
	r:= rect{3,4}
	c:= circle{2}

	var g geometry = &r
	measure(g) //measure(r) is not correct
	measure(&c)

	var e interface{} = g
	if value, ok := e.(int); ok{
		fmt.Println("is int type", value)
	}else if value,ok := e.(*rect); ok{
		fmt.Println("is rect type", value)
	}

	//special switch
	switch value:= e.(type){
	case int:
		fmt.Println("case int type", value)
	case *rect:
		fmt.Println("case rect type",value)
	default:
		fmt.Println("case not matched.")
	}
}


func DemoReflect(){
	var f float64 = 3.4
	v := reflect.ValueOf(f)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())

	var x float64 = 3.2
	p := reflect.ValueOf(&x)
	fmt.Println("type:", p.Type())
	fmt.Println("kind:", p.Kind()) //reflect.Ptr
	v1 := p.Elem()
	v1.SetFloat(3.4)

}