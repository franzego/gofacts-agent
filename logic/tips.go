package logic

import "math/rand"

var tips = []string{
	"✅ Keep your functions short and focused — one function, one responsibility.",
	"💡 Always run 'go fmt ./...' before committing — consistency builds clarity.",
	"🧠 Handle errors explicitly. Avoid ignoring them with '_'.",
	"🚀 Use goroutines for concurrency, but use channels or sync.WaitGroup to control them.",
	"🔍 Use 'go vet' and 'golangci-lint' to catch subtle bugs early.",
	"📦 Organize your code by domain, not type — keep related logic close together.",
	"⚙️ Prefer composition over inheritance — Go’s interfaces are about behavior, not hierarchy.",
	"🧱 Avoid global state; pass dependencies explicitly via structs or function parameters.",
	"⏱ Benchmark critical code paths with 'go test -bench .' to find real performance gains.",
	"🧩 Use 'context.Context' in all long-running functions to enable graceful cancellation.",
}

func TipsRandom() string {
	// num := rand.Intn(len(tips))
	return tips[rand.Intn(len(tips))]
}
