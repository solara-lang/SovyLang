package evaluator

import (
	"fmt"
	"sovylang/internal/object"
)

var builtins = map[string]*object.Builtin{
	"imprimir": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"tamanho": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("número errado de argumentos. esperado=1, recebido=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argumento para `tamanho` não suportado, recebido %s", args[0].Type())
			}
		},
	},
	"primeiro": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("número errado de argumentos. esperado=1, recebido=%d", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argumento para `primeiro` deve ser ARRAY, recebido %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"ultimo": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("número errado de argumentos. esperado=1, recebido=%d", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argumento para `ultimo` deve ser ARRAY, recebido %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"resto": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("número errado de argumentos. esperado=1, recebido=%d", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argumento para `resto` deve ser ARRAY, recebido %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"adicionar": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("número errado de argumentos. esperado=2, recebido=%d", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argumento para `adicionar` deve ser ARRAY, recebido %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
}
