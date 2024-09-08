package main

import (
	"github.com/adamdyszy/gorecursiveinterface/abstract"
	"github.com/adamdyszy/gorecursiveinterface/implementation"
)

func main() {
	l := implementation.NewLogger(0, 5)
	abstract.SetGlobalLogger(l)
	abstract.GlobalLogger.Info("test0")
	l1 := abstract.GlobalLogger.WithKeyValue("key1", "val1")
	l1.Info("test1")
	abstract.GlobalLogger.Info("testGlobalChange")
	l1_2 := abstract.GlobalLogger.WithKeyValue("key1", "val2")
	l1_2.Info("test1_2")
	l2 := l1.WithKeyValue("key2", "val1").V(2)
	l2.Info("test2")
	l3 := l2.WithKeyValue("key3", "val1").V(1)
	l3.Info("test3")
}
