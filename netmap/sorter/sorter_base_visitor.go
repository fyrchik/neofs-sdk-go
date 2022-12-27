// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Sorter

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseSorterVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSorterVisitor) VisitSorter(ctx *SorterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSorterVisitor) VisitWhenStmt(ctx *WhenStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSorterVisitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSorterVisitor) VisitPrio(ctx *PrioContext) interface{} {
	return v.VisitChildren(ctx)
}
