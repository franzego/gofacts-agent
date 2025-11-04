package logic

import (
	"fmt"
	"math/rand"
)

type Facts struct {
	ID          int
	Fact        string
	Explanation string
}

var fact = []Facts{
	{
		ID:          1,
		Fact:        "Go functions can return multiple values — perfect for (value, error) pairs.",
		Explanation: "This design avoids exceptions and encourages explicit error handling.",
	},
	{
		ID:          2,
		Fact:        "The Go mascot is called Gopher, created by Renee French in 2009.",
		Explanation: "It became one of the most recognized mascots in programming communities.",
	},
	{
		ID:          3,
		Fact:        "Go compiles directly to machine code using a static linking model.",
		Explanation: "This makes Go binaries fast to start and easy to deploy — no dependencies needed.",
	},
	{
		ID:          4,
		Fact:        "Go was created at Google by Robert Griesemer, Rob Pike, and Ken Thompson.",
		Explanation: "They designed it to improve developer productivity at scale.",
	},
	{
		ID:          5,
		Fact:        "Go has garbage collection, but its designed for low latency.",
		Explanation: "Its concurrent garbage collector minimizes pauses and supports large-scale systems.",
	},
	{
		ID:          6,
		Fact:        "The keyword 'defer' in Go executes a function after the surrounding function returns.",
		Explanation: "It's useful for resource cleanup like closing files or unlocking mutexes.",
	},
	{
		ID:          7,
		Fact:        "Go doesn't support traditional inheritance.",
		Explanation: "Instead, it promotes composition and interfaces for clean abstraction.",
	},
	{
		ID:          8,
		Fact:        "Go binaries include everything needed to run — no runtime dependencies.",
		Explanation: "This makes Go perfect for serverless and containerized environments.",
	},
	{
		ID:          9,
		Fact:        "The Go scheduler automatically balances goroutines across CPU cores.",
		Explanation: "This allows massive concurrency with minimal developer effort.",
	},
	{
		ID:          10,
		Fact:        "The Go standard library includes a built-in HTTP server — 'net/http'.",
		Explanation: "It is powerful enough to build production-grade APIs out of the box.",
	},
}

func FactRandom() string {
	s := fact[rand.Intn(len(fact))]
	return fmt.Sprintf("%s. %s", s.Fact, s.Explanation)
}
