package calc

type Calc struct {
	Value int
}

func (c Calc) LightCalc(num int) {
	c.Value += num
}

func (c Calc) HeavyCalc(num int) {
	for i := 0; i < num; i++ {
		c.Value += i
	}
}

func (c Calc) WithAdd(num int) Calc {
	c.Value += num
	return c
}

func (c Calc) WithMult(num int) Calc {
	c.Value *= num
	return c
}

func (c Calc) GetValue() int {
	return c.Value
}
