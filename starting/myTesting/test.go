package test

import (
	"fmt"
	"golang.org/x/net/html"
	"reflect"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var stack []string = make([]string, 0)

func (u User) String() string {

	return fmt.Sprintf("id: %d, name: %s", u.Id, u.Name)
}

func search(ls *[]string, temp *html.Node) {
	if temp.Data == "a" {
		for _, a := range temp.Attr {
			if a.Key == "href" {
				*ls = append(*ls, a.Val)
			}
		}
	}
	for c := temp.FirstChild; c != nil; c = c.NextSibling {
		stack = append(stack, c.Data)
		search(ls, c)
	}
	return
}

func solve(operate func(int) int, i int) int {
	return operate(i)
}
func sum(x int) int {
	defer func() {
		fmt.Println("worked")
	}()
	return x * x
}

func Maining() {

	fmt.Println(solve(sum, 10))
	s := "123"
	fmt.Println(reflect.TypeOf(rune(s[0])))
}
