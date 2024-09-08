package abstract

type Calc[T any] interface {
	NoRecursionCalc
	WithAdd(num int) T
	WithMult(num int) T
}

type NoRecursionCalc interface {
	LightCalc(num int)
	HeavyCalc(num int)
	GetValue() int
}

type WrapperCalc[T Calc[T]] struct {
	calc T
}

func (w WrapperCalc[T]) LightCalc(num int) {
	w.calc.LightCalc(num)
}

func (w WrapperCalc[T]) HeavyCalc(num int) {
	w.calc.HeavyCalc(num)
}

func (w WrapperCalc[T]) GetValue() int {
	return w.calc.GetValue()
}

func (w WrapperCalc[T]) WithAdd(num int) WrapperCalc[T] {
	w.calc = w.calc.WithAdd(num)
	return w
}

func (w WrapperCalc[T]) WithMult(num int) WrapperCalc[T] {
	w.calc = w.calc.WithMult(num)
	return w
}

func GetCalc[T Calc[T]](calc T) WrapperCalc[T] {
	return WrapperCalc[T]{calc: calc}
}

// func SetGlobalCalc[T Calc[T]](calc T) WrapperCalc[T] {
// 	? = WrapperCalc[T]{calc: calc}
// }
