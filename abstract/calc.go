package abstract

type Calc[T any] interface {
	NoRecursionCalc
	WithAdd(num int) T
	WithMult(num int) T
}

type NoRecursionCalc interface {
	LightCalc(num int)
	HeavyCalc(num int)
}

type WrapperCalc struct {
	NoRecursionCalc
	WithAdd  func(num int) WrapperCalc
	WithMult func(num int) WrapperCalc
}

func GetCalc[T NoRecursionCalc](calc Calc[T]) WrapperCalc {
	NewWrapperCalc := make([]func(calc Calc[T]) WrapperCalc, 1)
	NewWrapperCalc[0] = func(calc Calc[T]) WrapperCalc {
		return WrapperCalc{
			NoRecursionCalc: calc,
			WithAdd: func(num int) WrapperCalc {
				return NewWrapperCalc[0](any(calc.WithAdd(num)).(Calc[T]))
			},
			WithMult: func(num int) WrapperCalc {
				return NewWrapperCalc[0](any(calc.WithMult(num)).(Calc[T]))
			},
		}
	}
	return NewWrapperCalc[0](calc)
}
