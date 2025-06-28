package go_base

import (
	"fmt"
	"math"
)

type PayMethod interface {
	Account //匿名字段
	Pay(amount int) bool
}

type Account interface {
	Getbalance() int
}

// 定义一个接口 任何类型都可以赋值给空接口
type Fccount interface {
}

type CCard struct {
	balance int
	limit   int
}
type DCard struct {
	balance int
}

func (D DCard) Pay(amount int) bool {
	//TODO implement me
	panic("implement me")
}

func (D DCard) Getbalance() int {
	//TODO implement me
	panic("implement me")
}

func (c *CCard) Getbalance() int {

	return c.balance
}

func (c *CCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Println("支付成功: %d\n", amount)
		return true
	}
	fmt.Println("支付失败: 余额不足")
	return false
}

func per(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Println("购买成功")
	} else {
		fmt.Println("购买失败")
	}
}

func Interface_test1() {
	cread1 := &CCard{
		balance: 0,
		limit:   2000,
	}
	dcread := DCard{
		balance: 1000,
	}
	per(cread1, 100)

	//只能将指针赋值接口类型的变量
	var cr Account = cread1
	fmt.Println(cr)
	var dc Account = dcread
	fmt.Println(dc)

	var b PayMethod = cread1
	var c PayMethod = dcread

	//fmt.Println(b.Getbalance())
}

type notifer interface {
	notify()
}
type user struct {
	name  string
	email string
}

type admin struct {
	name string
	age  int
}

type adm struct {
	name string
	age  int
}

// 使用指针接收者实现了notofy接口,方法会共享接收者所指向的值user
func (u *user) notify() {
	fmt.Println("sendNotify to user", u.name)
}

// 使用值接收者实现了notofy接口,方法使用a值的副本,对a的修改不会影响原值
func (a admin) notify() {
	fmt.Println("sendNotify to admin", a.name)
}

// 多态的含义就是不需要修改函数，只需要修改外部实现
// 同一个接口有不同的表现形式
// 接收一个notifer接口类型的值，如果一个实体类型实现了该接口，
// sendNotify函数会根据实体类型的值类执行notifer接口的notify行为，这个函数具有多态的能力。
func SendNotify(n notifer) {
	n.notify()
}

func Interface_test2() {
	u1 := user{
		name:  "jack",
		email: "qq.com",
	}
	SendNotify(&u1)

	a := admin{
		name: "rose",
		age:  1,
	}

	var ad notifer = a

	SendNotify(ad)

	var n notifer
	n = &u1
	n.notify()

	var n2 notifer
	n2 = a
	n2.notify()
}

type person struct {
	name string
	age  int
}

// 如果一个struct嵌入另一个匿名结构体，就可以直接访问匿名结构体的字段或方法，从而实现继承
type student struct {
	person //匿名字段
	number string
}

// 如果一个struct嵌套了另一个【有名】的结构体，叫做组合
type teacher struct {
	p      person //有名字段
	mobile string
}

func (p *person) run() {
	fmt.Println(p.name, " person run")
}

func (s *student) reading() {
	fmt.Println(s.name, " student reading")
}

func Interface_test3() {
	s := student{person{"lisi", 20}, "000"}
	fmt.Println(s.name) //访问结构体的【匿名】字段，student对象没有name字段，访问的是从person继承过来的name字段
	s.run()             //访问结构体的【匿名】字段实现的方法，虽然student对象没有实现run方法，但student继承了person，person类型实现的方法能被直接访问

	t := teacher{person{"wangwu", 25}, "111"}
	fmt.Println(t.p.name) //访问【有名】结构体的字段。不是继承，不能被直接访问，需要指定结构体
	t.p.run()
}

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func Interface_test4() {
	r := Rectangle{Width: 4, Height: 5}
	c := Circle{Radius: 3}

	shapes := []Shape{r, c}

	for _, shape := range shapes {
		fmt.Printf("Area: %f\n", shape.Area())
	}
}
