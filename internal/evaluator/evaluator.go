package evaluator

import (
	"fmt"
	"sovylang/internal/ast"
	"sovylang/internal/library"
	"sovylang/internal/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func New() *Evaluator {
	return &Evaluator{
		libraryManager: library.NewLibraryManager(),
	}
}

type Evaluator struct{
	libraryManager *library.LibraryManager
}

func (e *Evaluator) Eval(node ast.Node) object.Object {
	env := object.NewEnvironment()
	return e.EvalWithEnv(node, env)
}

func (e *Evaluator) EvalWithEnv(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {


	case *ast.Program:
		return e.evalProgram(node, env)

	case *ast.BlockStatement:
		return e.evalBlockStatement(node, env)

	case *ast.ExpressionStatement:
		return e.EvalWithEnv(node.Expression, env)

	case *ast.VarStatement:
		val := e.EvalWithEnv(node.Value, env)
		if isError(val) {
			return val
		}
		env.Set(node.Name.Value, val)
		return val

	case *ast.ReturnStatement:
		val := e.EvalWithEnv(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.ForStatement:
		return e.evalForStatement(node, env)

	case *ast.IncludeStatement:
		return e.evalIncludeStatement(node, env)


	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToPyObject(node.Value)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.PrefixExpression:
		right := e.EvalWithEnv(node.Right, env)
		if isError(right) {
			return right
		}
		return e.evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:

		if e.needsMathLibrary(node.Operator) && !e.libraryManager.IsLibraryLoaded("smath") {
			return newError("operações matemáticas avançadas requerem a biblioteca 'smath'. Execute: sovy install smath")
		}

		left := e.EvalWithEnv(node.Left, env)
		if isError(left) {
			return left
		}

		right := e.EvalWithEnv(node.Right, env)
		if isError(right) {
			return right
		}

		return e.evalInfixExpression(node.Operator, left, right)

	case *ast.IfExpression:
		return e.evalIfExpression(node, env)

	case *ast.Identifier:
		return e.evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		params := node.Parameters
		body := node.Body
		fn := &object.Function{Parameters: params, Env: env, Body: body}


		if node.Name != nil {
			env.Set(node.Name.Value, fn)
		}

		return fn

	case *ast.CallExpression:
		function := e.EvalWithEnv(node.Function, env)
		if isError(function) {
			return function
		}

		args := e.evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return e.applyFunction(function, args)

	case *ast.ArrayLiteral:
		elements := e.evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *ast.IndexExpression:
		left := e.EvalWithEnv(node.Left, env)
		if isError(left) {
			return left
		}
		index := e.EvalWithEnv(node.Index, env)
		if isError(index) {
			return index
		}
		return e.evalIndexExpression(left, index)

	case *ast.HashLiteral:
		return e.evalHashLiteral(node, env)

	default:
		return newError("nó desconhecido: %T (%+v)", node, node)
	}
}

func (e *Evaluator) evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = e.EvalWithEnv(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func (e *Evaluator) evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = e.EvalWithEnv(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func (e *Evaluator) evalForStatement(node *ast.ForStatement, env *object.Environment) object.Object {
	start := e.EvalWithEnv(node.Start, env)
	if isError(start) {
		return start
	}

	end := e.EvalWithEnv(node.End, env)
	if isError(end) {
		return end
	}

	startInt, ok := start.(*object.Integer)
	if !ok {
		return newError("valor inicial do loop deve ser inteiro, recebido=%T", start)
	}

	endInt, ok := end.(*object.Integer)
	if !ok {
		return newError("valor final do loop deve ser inteiro, recebido=%T", end)
	}

	var result object.Object

	for i := startInt.Value; i <= endInt.Value; i++ {

		env.Set(node.Variable.Value, &object.Integer{Value: i})

		result = e.EvalWithEnv(node.Body, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func (e *Evaluator) evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!", "não", "nao":
		return e.evalBangOperatorExpression(right)
	case "-":
		return e.evalMinusPrefixOperatorExpression(right)
	default:
		return newError("operador desconhecido: %s%s", operator, right.Type())
	}
}

func (e *Evaluator) evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return e.evalIntegerInfixExpression(operator, left, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return e.evalFloatInfixExpression(operator, left, right)
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		leftFloat := &object.Float{Value: float64(left.(*object.Integer).Value)}
		return e.evalFloatInfixExpression(operator, leftFloat, right)
	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		rightFloat := &object.Float{Value: float64(right.(*object.Integer).Value)}
		return e.evalFloatInfixExpression(operator, left, rightFloat)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return e.evalStringInfixExpression(operator, left, right)
	case operator == "==":
		return nativeBoolToPyObject(left == right)
	case operator == "!=":
		return nativeBoolToPyObject(left != right)
	case operator == "e":
		return nativeBoolToPyObject(isTruthy(left) && isTruthy(right))
	case operator == "ou":
		return nativeBoolToPyObject(isTruthy(left) || isTruthy(right))
	default:
		return newError("operador desconhecido: %s %s %s", left.Type(), operator, right.Type())
	}
}

func (e *Evaluator) evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("divisão por zero")
		}
		return &object.Float{Value: float64(leftVal) / float64(rightVal)}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToPyObject(leftVal < rightVal)
	case ">":
		return nativeBoolToPyObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToPyObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToPyObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToPyObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToPyObject(leftVal != rightVal)
	default:
		return newError("operador desconhecido: %s", operator)
	}
}

func (e *Evaluator) evalFloatInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "/":
		if rightVal == 0 {
			return newError("divisão por zero")
		}
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToPyObject(leftVal < rightVal)
	case ">":
		return nativeBoolToPyObject(leftVal > rightVal)
	case "<=":
		return nativeBoolToPyObject(leftVal <= rightVal)
	case ">=":
		return nativeBoolToPyObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToPyObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToPyObject(leftVal != rightVal)
	default:
		return newError("operador desconhecido: %s", operator)
	}
}

func (e *Evaluator) evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToPyObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToPyObject(leftVal != rightVal)
	default:
		return newError("operador desconhecido: %s", operator)
	}
}

func (e *Evaluator) evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func (e *Evaluator) evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	switch right := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -right.Value}
	case *object.Float:
		return &object.Float{Value: -right.Value}
	default:
		return newError("operador desconhecido: -%s", right.Type())
	}
}

func (e *Evaluator) evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := e.EvalWithEnv(ie.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return e.EvalWithEnv(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return e.EvalWithEnv(ie.Alternative, env)
	} else {
		return NULL
	}
}

func (e *Evaluator) evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identificador não encontrado: " + node.Value)
	}

	return val
}

func (e *Evaluator) evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, exp := range exps {
		evaluated := e.EvalWithEnv(exp, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func (e *Evaluator) applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := e.extendFunctionEnv(fn, args)
		evaluated := e.EvalWithEnv(fn.Body, extendedEnv)
		return e.unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("não é uma função: %T", fn)
	}
}

func (e *Evaluator) extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func (e *Evaluator) unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func (e *Evaluator) evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return e.evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return e.evalHashIndexExpression(left, index)
	default:
		return newError("operador de índice não suportado: %s", left.Type())
	}
}

func (e *Evaluator) evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func (e *Evaluator) evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := e.EvalWithEnv(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("chave de hash inválida: %T", key)
		}

		value := e.EvalWithEnv(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func (e *Evaluator) evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("chave de hash inválida: %T", index)
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func nativeBoolToPyObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func (e *Evaluator) evalIncludeStatement(node *ast.IncludeStatement, env *object.Environment) object.Object {
	libraryName := node.Library.Value

	err := e.libraryManager.LoadLibrary(libraryName)
	if err != nil {
		return newError(err.Error())
	}


	if libraryName == "smath" {
		smathLib := library.NewSMathLibrary()
		builtins := smathLib.GetBuiltins()


		for name, fn := range builtins {
			env.Set(name, fn)
		}
	}

	return NULL
}

func (e *Evaluator) needsMathLibrary(operator string) bool {


	mathOperators := []string{"/", "%"}
	for _, op := range mathOperators {
		if operator == op {
			return true
		}
	}
	return false
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}
