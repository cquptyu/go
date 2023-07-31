package main

type alive interface {
	walk()
}
type people struct {
	name string
	age  int
}

func (people) walk() {
	println("走路")

}

func main() {

	var peple = new(people)
	println(peple)
}
