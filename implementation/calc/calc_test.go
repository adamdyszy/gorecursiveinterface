package calc

import (
	"testing"

	"github.com/adamdyszy/gorecursiveinterface/abstract"
	"github.com/stretchr/testify/assert"
)

func BenchmarkDirectChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	val := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = c.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, c.GetValue())
	}
}

func BenchmarkWrapperChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c)
	val := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, wrapper.GetValue())
	}
}

func BenchmarkWrapperGlobChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.SetGlobalCalc(c)
	val := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, wrapper.GetValue())
	}
}

func BenchmarkManyOperationsDirectCalc(b *testing.B) {
	c := Calc{Value: 1}
	val := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			c = c.WithAdd(j).WithMult(j + 1)
			val += j
			val *= (j + 1)
		}
		assert.Equal(b, val, c.GetValue())
	}
}

func BenchmarkManyOperationsWrapperCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c)
	val := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			wrapper = wrapper.WithAdd(j).WithMult(j + 1)
			val += j
			val *= (j + 1)
		}
		assert.Equal(b, val, wrapper.GetValue())
	}
}

func BenchmarkManyOperationsWrapperGlobCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.SetGlobalCalc(c)
	val := 1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			wrapper = wrapper.WithAdd(j).WithMult(j + 1)
			val += j
			val *= (j + 1)
		}
		assert.Equal(b, val, wrapper.GetValue())
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

func BenchmarkWrapperGlobCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.SetGlobalCalc(c).WithMult(425354)
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
	val := 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = c.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, c.GetValue())
		c.HeavyCalc(5234)
	}
}

func BenchmarkMixedWrapperChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c).WithMult(2).WithAdd(3)
	val := 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, wrapper.GetValue())
		wrapper.HeavyCalc(5234)
	}
}

func BenchmarkMixedWrapperGlobChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.SetGlobalCalc(c).WithMult(2).WithAdd(3)
	val := 5
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, wrapper.GetValue())
		wrapper.HeavyCalc(5234)
	}
}

func BenchmarkHeavyInitMixedDirectChainCalc(b *testing.B) {
	c := Calc{Value: 1}.WithMult(2).WithAdd(3)
	val := 5
	for range 10000000 {
		c = c.WithAdd(72345)
		val += 72345
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c = c.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, c.GetValue())
		c.HeavyCalc(5234)
	}
}

func BenchmarkHeavyInitMixedWrapperChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.GetCalc(c).WithMult(2).WithAdd(3)
	val := 5
	for range 10000000 {
		wrapper = wrapper.WithAdd(72345)
		val += 72345
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, wrapper.GetValue())
		wrapper.HeavyCalc(5234)
	}
}

func BenchmarkHeavyInitMixedWrapperGlobChainCalc(b *testing.B) {
	c := Calc{Value: 1}
	wrapper := abstract.SetGlobalCalc(c).WithMult(2).WithAdd(3)
	val := 5
	for range 10000000 {
		wrapper = wrapper.WithAdd(72345)
		val += 72345
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wrapper = wrapper.WithAdd(1).WithMult(2).WithAdd(3).WithMult(4)
		val += 1
		val *= 2
		val += 3
		val *= 4
		assert.Equal(b, val, wrapper.GetValue())
		wrapper.HeavyCalc(5234)
	}
}
