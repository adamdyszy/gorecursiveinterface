package abstract

// func SetGlobalCalcGlob[T CalcGlob[T]](calc T) WrapperCalcGlob[T] {
// 	? = WrapperCalcGlob[T]{calc: calc}
// }

var GlobalCalc WrapperCalcGlob

type WrapperCalcGlob struct {
	NoRecursionCalc
	WithAdd  func(num int) WrapperCalcGlob
	WithMult func(num int) WrapperCalcGlob
}

//	func (w WrapperLogger) V(level int) WrapperLogger {
//		w.NoRecursionLogger = ????????.V(level).NoRecursionLogger
//		return w
//	}

func NewWrapperCalc[T any](calc Calc[T]) WrapperCalcGlob {
	return WrapperCalcGlob{
		NoRecursionCalc: calc,
		WithAdd: func(num int) WrapperCalcGlob {
			return NewWrapperCalc(any(calc.WithAdd(num)).(Calc[T]))
		},
		WithMult: func(num int) WrapperCalcGlob {
			return NewWrapperCalc(any(calc.WithMult(num)).(Calc[T]))
		},
	}
}

func SetGlobalCalc[T any](calc Calc[T]) WrapperCalcGlob {
	GlobalCalc = NewWrapperCalc(calc)
	return GlobalCalc
}
