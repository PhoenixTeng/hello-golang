package main

import (
	"1/a"
	"fmt"
)

type guo struct {
	renkou int
	jia    *a.Jia
}

type sheng struct {
	renkou int
	a.Jia
}

type her interface {
	He()
}

func (g *guo) j() *a.Jia {
	g.jia = newJia
	return g.jia
}

func (g *guo) He() {
	g.j().He()
	println("富家大吉")
}

func (j *sheng) He() {
	j.Jia.He()
	println("富家大吉")
}

var newJia = new(a.Jia)

var abc = 1
var abd = "akkg"
var test = "hello"

func main() {
	fmt.Printf("呵；人\n")
	a.F2("554566")
	a.F3()
	println(abc)

	/*
		fptr := flag.String("fpath", "test.txt", "file path to read from")
		fprt := flag.String("fprint", "print", "hello world")
		flag.Parse()
		fmt.Println(*fprt)
		data, err := ioutil.ReadFile(*fptr)
		if err != nil {
			fmt.Println("file err:", err)
			return
			}
			fmt.Printf("Contents of file:%s", data)
	*/

	var g1, g2 guo
	var s1 sheng
	g1.renkou = 1
	g1.j().Fu = "x"

	fmt.Println(g1, g2.j().Fu)

	s1.Fu = "富家大吉"
	s1.He()

	println("\n\n数组，切片，结构体，结构体指针，接口，接口地址:")
	x := [2]int{0, 1}
	y := make([]int, 5)
	println("array address:", &x, "slice println:", y)
	x = [2]int{2, 3}
	y = []int{1, 1, 1, 1, 1}
	println("array address:", &x, "slice println:", y)
	x[1] = 4
	y = []int{1, 1, 1, 1, 1}
	println("array address:", &x, "slice println:", y, "\n")

	//结构体，结构体指针
	g := guo{renkou: 2}
	println("struct g address(&g) println:", &g)
	g = guo{2, nil}
	println("the same:", &g, "\n")

	gp := &guo{renkou: 2}
	println("struct pointer println:", gp)
	fmt.Println(gp)
	gp = &guo{renkou: 2, jia: nil}
	println("diffent:", gp)
	fmt.Println(gp)
	println()

	ftest()

	var herI interface{}
	var herI2 her
	herP := &herI
	println("nil interface println:", herI, *herP)
	println(herI, herI2)
	fmt.Println("fmt.Println:", herI, herI2)
	println("nill interface address:", &herI, &herI2)
	//init(&g)
	println("把字面量变量取地址赋给接口：")
	herI, herI2 = &g, &g
	println("init\"&g\" interface (herI herI2) println:", herI, herI2)
	fmt.Println("init\"&g\" interface fmt.Println:", herI, herI2)
	println("init interface (herI herI2) address(same):", &herI, &herI2)
	println("stack struct g address(&g):", &g, "\n")
	//intt(gp)
	println("把堆里变量地址赋给接口：")
	herI, herI2 = gp, gp
	println("init\"gp\" interface (herI herI2) println:", herI, herI2)
	fmt.Println("init\"gp\" interface fmt.Println:", herI, herI2)
	println("init interface	(herI herI2) address(same):", &herI, &herI2)
	println("heap struct pointer gp value:", gp, "\n")

	var as []string
	println("nil slice as println:", as)
	println("as address:", &as)
	fmt.Println("as fmt.Println:", as)
	as = make([]string, 5)
	as[0] = "kekaldf"
	as[1] = "iwoela"
	println("init slice println:", as)
	fmt.Println("init fmt.Println:", as)

	herI = []int{1, 2, 3}
	fmt.Println(herI.([]int))

}

func ftest() {
	test := "接口赋值："
	println(test)

}
