package netmap

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	sorter "github.com/nspcc-dev/neofs-sdk-go/netmap/sorter"
)

type Sorter struct {
	attribute string
	values    map[string]uint64
}

func (s *Sorter) WeightFunc() WeightFunc {
	return func(n NodeInfo) float64 {
		attr := n.Attribute(s.attribute)
		return float64(s.values[attr])
	}
}

type sorterVisitor struct {
	errors []error
	sorter.BaseSorterVisitor
	antlr.DefaultErrorListener
}

var _ sorter.SorterVisitor = (*sorterVisitor)(nil)

func (p *sorterVisitor) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	p.reportError(fmt.Errorf("%w: line %d:%d %s", errSyntaxError, line, column, msg))
}

func (p *sorterVisitor) reportError(err error) interface{} {
	p.errors = append(p.errors, err)
	return nil
}

type keyValue struct {
	name     string
	priority uint64
}

func (p *sorterVisitor) VisitSorter(ctx *sorter.SorterContext) interface{} {
	if len(p.errors) != 0 {
		return nil
	}

	pl := &Sorter{values: make(map[string]uint64)}
	pl.attribute = ctx.IDENT().GetText()

	whenStmts := ctx.AllWhenStmt()
	for _, r := range whenStmts {
		res, ok := r.Accept(p).(keyValue)
		if !ok {
			return nil
		}

		pl.values[res.name] = res.priority
	}

	return pl
}

func (p *sorterVisitor) VisitWhenStmt(ctx *sorter.WhenStmtContext) interface{} {
	prio, err := strconv.ParseUint(ctx.GetPriority().GetText(), 10, 32)
	if err != nil {
		return p.reportError(errInvalidNumber)
	}

	return keyValue{
		name:     ctx.GetName().GetText(),
		priority: prio,
	}
}

func (s *Sorter) DecodeString(str string) error {
	var v sorterVisitor

	input := antlr.NewInputStream(str)
	lexer := sorter.NewSorterLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&v)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	pp := sorter.NewSorter(stream)
	pp.BuildParseTrees = true

	pp.RemoveErrorListeners()
	pp.AddErrorListener(&v)
	pl := pp.Sorter().Accept(&v)

	if len(v.errors) != 0 {
		return v.errors[0]
	}

	parsed, ok := pl.(*Sorter)
	if !ok {
		return fmt.Errorf("unexpected parsed instance type %T", pl)
	} else if parsed == nil {
		return errors.New("parsed nil value")
	}

	*s = *parsed
	return nil
}
