// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Sorter

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by Sorter.
type SorterVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Sorter#sorter.
	VisitSorter(ctx *SorterContext) interface{}

	// Visit a parse tree produced by Sorter#whenStmt.
	VisitWhenStmt(ctx *WhenStmtContext) interface{}

	// Visit a parse tree produced by Sorter#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by Sorter#prio.
	VisitPrio(ctx *PrioContext) interface{}
}
