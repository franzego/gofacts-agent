package logic

import "math/rand"

var syntaxSnippets = []string{
	"💻 Basic for-loop:\n```go\nfor i := 0; i < 5; i++ {\n    fmt.Println(i)\n}\n```",
	"🔤 Declare multiple variables:\n```go\nvar a, b = 3, 7\n```",
	"📦 Short variable declaration:\n```go\nx := 42\n```",
	"🧱 Anonymous function:\n```go\nfunc(x, y int) int { return x + y }\n```",
	"🔁 Range loop over a slice:\n```go\nfor index, value := range nums {\n    fmt.Println(index, value)\n}\n```",
	"🧩 Struct definition and usage:\n```go\ntype Person struct {\n    Name string\n    Age  int\n}\n\np := Person{Name: \"T\", Age: 30}\nfmt.Println(p.Name)\n```",
	"📜 Goroutine example:\n```go\ngo func() {\n    fmt.Println(\"Running concurrently!\")\n}()\n```",
	"📬 Channel example:\n```go\nch := make(chan string)\ngo func() { ch <- \"ping\" }()\nmsg := <-ch\nfmt.Println(msg)\n```",
	"⚙️ Defer keyword:\n```go\nfunc readFile() {\n    file, _ := os.Open(\"data.txt\")\n    defer file.Close()\n}\n```",
	"🧠 Interface implementation:\n```go\ntype Shape interface {\n    Area() float64\n}\n```",
}

func SyntaxRandom() string {
	return syntaxSnippets[rand.Intn(len(syntaxSnippets))]
}
