package calc

type CalcPointer struct {
	Value int
}

func (c *CalcPointer) LightCalc(num int) {
	c.Value += num
}

func (c *CalcPointer) HeavyCalc(num int) {
	for i := 0; i < num; i++ {
		c.Value += i
	}
}

func (c *CalcPointer) WithAdd(num int) *CalcPointer {
	c.Value += num
	return c
}

func (c *CalcPointer) WithMult(num int) *CalcPointer {
	c.Value *= num
	return c
}
