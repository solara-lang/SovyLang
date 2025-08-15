package formatter

import (
	"bytes"
	"sovylang/internal/ast"
	"strings"
)

type Formatter struct {
	indentLevel int
	indentSize  int
}

func Format(program *ast.Program) string {
	f := &Formatter{
		indentLevel: 0,
		indentSize:  4,
	}
	return f.formatProgram(program)
}

func (f *Formatter) formatProgram(program *ast.Program) string {
	var out bytes.Buffer
	
	for i, stmt := range program.Statements {
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString(f.formatStatement(stmt))
	}
	
	return out.String()
}

func (f *Formatter) formatStatement(stmt ast.Statement) string {
	switch s := stmt.(type) {
	case *ast.VarStatement:
		return f.formatVarStatement(s)
	case *ast.ReturnStatement:
		return f.formatReturnStatement(s)
	case *ast.ExpressionStatement:
		return f.formatExpressionStatement(s)
	case *ast.ForStatement:
		return f.formatForStatement(s)
	case *ast.BlockStatement:
		return f.formatBlockStatement(s)
	default:
		return s.String()
	}
}

func (f *Formatter) formatVarStatement(vs *ast.VarStatement) string {
	var out bytes.Buffer
	out.WriteString(f.indent())
	out.WriteString(vs.Type + " " + vs.Name.Value + " = ")
	out.WriteString(f.formatExpression(vs.Value))
	return out.String()
}

func (f *Formatter) formatReturnStatement(rs *ast.ReturnStatement) string {
	var out bytes.Buffer
	out.WriteString(f.indent())
	out.WriteString("retorne ")
	if rs.ReturnValue != nil {
		out.WriteString(f.formatExpression(rs.ReturnValue))
	}
	return out.String()
}

func (f *Formatter) formatExpressionStatement(es *ast.ExpressionStatement) string {
	var out bytes.Buffer
	out.WriteString(f.indent())
	if es.Expression != nil {
		out.WriteString(f.formatExpression(es.Expression))
	}
	return out.String()
}

func (f *Formatter) formatForStatement(fs *ast.ForStatement) string {
	var out bytes.Buffer
	out.WriteString(f.indent())
	out.WriteString("para numero " + fs.Variable.Value + " = ")
	out.WriteString(f.formatExpression(fs.Start))
	out.WriteString(" até ")
	out.WriteString(f.formatExpression(fs.End))
	out.WriteString("\n")
	
	f.indentLevel++
	out.WriteString(f.formatBlockStatement(fs.Body))
	f.indentLevel--
	
	out.WriteString("\n" + f.indent() + "fim")
	return out.String()
}

func (f *Formatter) formatBlockStatement(bs *ast.BlockStatement) string {
	var out bytes.Buffer
	
	for i, stmt := range bs.Statements {
		if i > 0 {
			out.WriteString("\n")
		}
		out.WriteString(f.formatStatement(stmt))
	}
	
	return out.String()
}

func (f *Formatter) formatExpression(exp ast.Expression) string {
	switch e := exp.(type) {
	case *ast.Identifier:
		return e.Value
	case *ast.IntegerLiteral:
		return e.String()
	case *ast.FloatLiteral:
		return e.String()
	case *ast.StringLiteral:
		return "\"" + e.Value + "\""
	case *ast.Boolean:
		return e.String()
	case *ast.PrefixExpression:
		return f.formatPrefixExpression(e)
	case *ast.InfixExpression:
		return f.formatInfixExpression(e)
	case *ast.IfExpression:
		return f.formatIfExpression(e)
	case *ast.FunctionLiteral:
		return f.formatFunctionLiteral(e)
	case *ast.CallExpression:
		return f.formatCallExpression(e)
	case *ast.ArrayLiteral:
		return f.formatArrayLiteral(e)
	case *ast.HashLiteral:
		return f.formatHashLiteral(e)
	case *ast.IndexExpression:
		return f.formatIndexExpression(e)
	default:
		return e.String()
	}
}

func (f *Formatter) formatPrefixExpression(pe *ast.PrefixExpression) string {
	return pe.Operator + f.formatExpression(pe.Right)
}

func (f *Formatter) formatInfixExpression(ie *ast.InfixExpression) string {
	return f.formatExpression(ie.Left) + " " + ie.Operator + " " + f.formatExpression(ie.Right)
}

func (f *Formatter) formatIfExpression(ie *ast.IfExpression) string {
	var out bytes.Buffer
	out.WriteString("se " + f.formatExpression(ie.Condition))
	out.WriteString("\n")
	
	f.indentLevel++
	out.WriteString(f.formatBlockStatement(ie.Consequence))
	f.indentLevel--
	
	if ie.Alternative != nil {
		out.WriteString("\n" + f.indent() + "senão")
		out.WriteString("\n")
		
		f.indentLevel++
		out.WriteString(f.formatBlockStatement(ie.Alternative))
		f.indentLevel--
	}
	
	out.WriteString("\n" + f.indent() + "fim")
	return out.String()
}

func (f *Formatter) formatFunctionLiteral(fl *ast.FunctionLiteral) string {
	var out bytes.Buffer
	
	out.WriteString("função")
	if fl.Name != nil {
		out.WriteString(" " + fl.Name.Value)
	}
	out.WriteString("(")
	
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString("\n")
	
	f.indentLevel++
	out.WriteString(f.formatBlockStatement(fl.Body))
	f.indentLevel--
	
	out.WriteString("\n" + f.indent() + "fim")
	return out.String()
}

func (f *Formatter) formatCallExpression(ce *ast.CallExpression) string {
	var out bytes.Buffer
	
	out.WriteString(f.formatExpression(ce.Function))
	out.WriteString("(")
	
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, f.formatExpression(a))
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	
	return out.String()
}

func (f *Formatter) formatArrayLiteral(al *ast.ArrayLiteral) string {
	var out bytes.Buffer
	
	elements := []string{}
	for _, e := range al.Elements {
		elements = append(elements, f.formatExpression(e))
	}
	
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	
	return out.String()
}

func (f *Formatter) formatHashLiteral(hl *ast.HashLiteral) string {
	var out bytes.Buffer
	
	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, f.formatExpression(key)+": "+f.formatExpression(value))
	}
	
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	
	return out.String()
}

func (f *Formatter) formatIndexExpression(ie *ast.IndexExpression) string {
	return f.formatExpression(ie.Left) + "[" + f.formatExpression(ie.Index) + "]"
}

func (f *Formatter) indent() string {
	return strings.Repeat(" ", f.indentLevel*f.indentSize)
}
