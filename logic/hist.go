package logic

import (
	"fmt"
	"math/rand"
)

type HistoryFact struct {
	Year  int
	Event string
}

var history = []HistoryFact{
	{2007, "Go development began at Google as an internal experiment led by Robert Griesemer, Rob Pike, and Ken Thompson."},
	{2009, "Go was publicly announced as an open-source language."},
	{2010, "The first official release of Go (Go 1.0) launched, introducing the standard library and garbage collector."},
	{2012, "The Go logo — the Gopher — became the official mascot."},
	{2015, "Go 1.5 was released, removing all C code from the compiler and runtime, making Go self-hosting."},
	{2016, "Go 1.7 introduced significant performance improvements and better context propagation."},
	{2018, "Go modules were introduced for dependency management, replacing GOPATH."},
	{2021, "Go 1.17 added generics preview — one of the most awaited features."},
	{2022, "Go 1.18 officially introduced generics and fuzz testing."},
	{2024, "Go remains a top language for cloud-native development, powering Docker, Kubernetes, and modern backend systems."},
}

func HistoryRandom() string {
	s := history[rand.Intn(len(history))]
	return fmt.Sprintf("%d. %s", s.Year, s.Event)
}
