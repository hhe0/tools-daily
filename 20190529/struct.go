package main

import "fmt"

type author struct {
	firstName string
	lastName  string
	bio       string
}

func (a author) getFullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post struct {
	title   string
	content string
	author
}

func (p post) getDetail() {
	fmt.Println("Title: ", p.title)
	fmt.Println("Content: ", p.content)
	//fmt.Println("Author: ", p.author.getFullName())
	fmt.Println("Author: ", p.getFullName())
	//fmt.Println("Bio: ", p.author.bio)
	fmt.Println("Bio: ", p.bio)
}

type website struct {
	p []post
}

func (w website) contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.p {
		v.getDetail()
		fmt.Println()
	}
}

func main() {
	a := author{
		"he",
		"hong",
		"Minc",
	}
	fmt.Println(a.getFullName())

	p := post{
		"新闻",
		"内容",
		a,
	}
	p.getDetail()

	w := website{
		p: []post{p},
	}
	w.contents()

	for _, post := range w.p {
		post.getDetail()
	}
}
