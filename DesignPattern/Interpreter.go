package main

import (
	"fmt"
	"strings"
)

/*
	解释器模式（Interpreter Pattern）提供了评估语言的语法或表达式的方式，它属于行为型模式。
	这种模式实现了一个表达式接口，该接口解释一个特定的上下文。
	这种模式被用在 SQL 解析、符号处理引擎等。
*/

type Expression interface {
	Interpret(context string) bool
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (a *AndExpression) AndExpression(expr1 Expression,expr2 Expression)  {
	a.expr1 = expr1
	a.expr2 = expr2
}

func (a *AndExpression) Interpret(context string) bool {
	return a.expr1.Interpret(context) && a.expr2.Interpret(context)
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (o *OrExpression) OrExpression(expr1 Expression,expr2 Expression)  {
	o.expr1 = expr1
	o.expr2 = expr2
}

func (o *OrExpression) Interpret(context string) bool {
	return o.expr1.Interpret(context) || o.expr2.Interpret(context)
}

type TerminalExpression struct {
	Data string
}

func (t *TerminalExpression) TerminalExpression(data string)  {
	t.Data = data
}

func (t *TerminalExpression) Interpret(context string) bool {
	// 这个接口注意 后面的才是子串 表示context里包含t.Data
	return strings.Contains(context,t.Data)
}

// 规则1构建 Robert和John是男性
func GetMaleExpression() Expression {
	robert := new(TerminalExpression)
	robert.TerminalExpression("Robert")
	john := new(TerminalExpression)
	john.TerminalExpression("John")
	orExpression := new(OrExpression)
	orExpression.OrExpression(robert,john)
	return orExpression
}

// 规则2 构建 Julie已婚
func GetMarriedWomenExpression() Expression {
	julie := new(TerminalExpression)
	julie.TerminalExpression("Julie")
	married := new(TerminalExpression)
	married.TerminalExpression("married")
	andExpression := new(AndExpression)
	andExpression.AndExpression(julie,married)
	return andExpression
}

func TestInterpreter() {

	isMale := GetMaleExpression()
	isMarriedWoman := GetMarriedWomenExpression()
	fmt.Println("John is male? ", isMale.Interpret("John"))
	fmt.Println("Julie is a married women? ", isMarriedWoman.Interpret("Married Julie"))
}

func main() {
	TestInterpreter()
}