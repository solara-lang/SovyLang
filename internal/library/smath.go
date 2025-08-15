package library

import (
	"math"
	"sovylang/internal/object"
)


type SMathLibrary struct{}

func NewSMathLibrary() *SMathLibrary {
	return &SMathLibrary{}
}

func (s *SMathLibrary) GetBuiltins() map[string]*object.Builtin {
	return map[string]*object.Builtin{
		"potencia": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return &object.Error{Message: "potencia() requer 2 argumentos (base, expoente)"}
				}

				base, ok := args[0].(*object.Float)
				if !ok {
					if intBase, ok := args[0].(*object.Integer); ok {
						base = &object.Float{Value: float64(intBase.Value)}
					} else {
						return &object.Error{Message: "primeiro argumento deve ser um número"}
					}
				}

				exponent, ok := args[1].(*object.Float)
				if !ok {
					if intExp, ok := args[1].(*object.Integer); ok {
						exponent = &object.Float{Value: float64(intExp.Value)}
					} else {
						return &object.Error{Message: "segundo argumento deve ser um número"}
					}
				}

				result := math.Pow(base.Value, exponent.Value)
				return &object.Float{Value: result}
			},
		},
		"raiz": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return &object.Error{Message: "raiz() requer 1 argumento"}
				}

				num, ok := args[0].(*object.Float)
				if !ok {
					if intNum, ok := args[0].(*object.Integer); ok {
						num = &object.Float{Value: float64(intNum.Value)}
					} else {
						return &object.Error{Message: "argumento deve ser um número"}
					}
				}

				if num.Value < 0 {
					return &object.Error{Message: "não é possível calcular raiz de número negativo"}
				}

				result := math.Sqrt(num.Value)
				return &object.Float{Value: result}
			},
		},
		"sin": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return &object.Error{Message: "sin() requer 1 argumento"}
				}

				num, ok := args[0].(*object.Float)
				if !ok {
					if intNum, ok := args[0].(*object.Integer); ok {
						num = &object.Float{Value: float64(intNum.Value)}
					} else {
						return &object.Error{Message: "argumento deve ser um número"}
					}
				}

				result := math.Sin(num.Value)
				return &object.Float{Value: result}
			},
		},
		"cos": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return &object.Error{Message: "cos() requer 1 argumento"}
				}

				num, ok := args[0].(*object.Float)
				if !ok {
					if intNum, ok := args[0].(*object.Integer); ok {
						num = &object.Float{Value: float64(intNum.Value)}
					} else {
						return &object.Error{Message: "argumento deve ser um número"}
					}
				}

				result := math.Cos(num.Value)
				return &object.Float{Value: result}
			},
		},
		"abs": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 1 {
					return &object.Error{Message: "abs() requer 1 argumento"}
				}

				switch arg := args[0].(type) {
				case *object.Integer:
					if arg.Value < 0 {
						return &object.Integer{Value: -arg.Value}
					}
					return arg
				case *object.Float:
					return &object.Float{Value: math.Abs(arg.Value)}
				default:
					return &object.Error{Message: "argumento deve ser um número"}
				}
			},
		},
		"max": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return &object.Error{Message: "max() requer 2 argumentos"}
				}


				var a, b float64

				switch arg := args[0].(type) {
				case *object.Integer:
					a = float64(arg.Value)
				case *object.Float:
					a = arg.Value
				default:
					return &object.Error{Message: "primeiro argumento deve ser um número"}
				}

				switch arg := args[1].(type) {
				case *object.Integer:
					b = float64(arg.Value)
				case *object.Float:
					b = arg.Value
				default:
					return &object.Error{Message: "segundo argumento deve ser um número"}
				}

				result := math.Max(a, b)
				return &object.Float{Value: result}
			},
		},
		"min": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 2 {
					return &object.Error{Message: "min() requer 2 argumentos"}
				}


				var a, b float64

				switch arg := args[0].(type) {
				case *object.Integer:
					a = float64(arg.Value)
				case *object.Float:
					a = arg.Value
				default:
					return &object.Error{Message: "primeiro argumento deve ser um número"}
				}

				switch arg := args[1].(type) {
				case *object.Integer:
					b = float64(arg.Value)
				case *object.Float:
					b = arg.Value
				default:
					return &object.Error{Message: "segundo argumento deve ser um número"}
				}

				result := math.Min(a, b)
				return &object.Float{Value: result}
			},
		},
		"pi": {
			Fn: func(args ...object.Object) object.Object {
				if len(args) != 0 {
					return &object.Error{Message: "pi() não aceita argumentos"}
				}
				return &object.Float{Value: math.Pi}
			},
		},
	}
}
