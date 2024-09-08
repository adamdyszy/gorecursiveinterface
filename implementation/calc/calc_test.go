package calc

import (
	"testing"

	"github.com/adamdyszy/gorecursiveinterface/abstract"
)

func BenchmarkDirectChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = c.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
	}
}

func BenchmarkWrapperChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
	}
}

func BenchmarkManyOperationsDirectCalc(b *testing.B) {
	c := Calc{Value: 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			c = c.WithAdd(j).WithMult(j + 1)
		}
	}
}

func BenchmarkManyOperationsWrapperCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			wrapper = wrapper.WithAdd(j).WithMult(j + 1)
		}
	}
}

func BenchmarkDirectCalc(b *testing.B) {
	c := Calc{Value: 1}.WithMult(425354)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.LightCalc(5)
		c.HeavyCalc(10)
		c.LightCalc(3)
		c.HeavyCalc(7)
	}
}

func BenchmarkWrapperCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c).WithMult(425354)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper.LightCalc(5)
		wrapper.HeavyCalc(10)
		wrapper.LightCalc(3)
		wrapper.HeavyCalc(7)
	}
}

func BenchmarkMixedDirectChainCalc(b *testing.B) {
	c := Calc{Value: 1}.WithMult(2).WithAdd(3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = c.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		c.HeavyCalc(5234)
	}
}

func BenchmarkMixedWrapperChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c).WithMult(2).WithAdd(3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		wrapper.HeavyCalc(5234)
	}
}

func BenchmarkHeavyInitMixedDirectChainCalc(b *testing.B) {
	c := Calc{Value: 1}.WithMult(2).WithAdd(3)
	for _ = range 10000000 {
		c = c.WithAdd(72345)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = c.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		c.HeavyCalc(5234)
	}
}

func BenchmarkHeavyInitMixedWrapperChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c).WithMult(2).WithAdd(3)
	for _ = range 10000000 {
		wrapper = wrapper.WithAdd(72345)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		wrapper.HeavyCalc(5234)
	}
}
