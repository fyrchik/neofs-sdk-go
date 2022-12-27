// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Sorter

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseSorterListener is a complete listener for a parse tree produced by Sorter.
type BaseSorterListener struct{}

var _ SorterListener = &BaseSorterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSorterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSorterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSorterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSorterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSorter is called when production sorter is entered.
func (s *BaseSorterListener) EnterSorter(ctx *SorterContext) {}

// ExitSorter is called when production sorter is exited.
func (s *BaseSorterListener) ExitSorter(ctx *SorterContext) {}

// EnterWhenStmt is called when production whenStmt is entered.
func (s *BaseSorterListener) EnterWhenStmt(ctx *WhenStmtContext) {}

// ExitWhenStmt is called when production whenStmt is exited.
func (s *BaseSorterListener) ExitWhenStmt(ctx *WhenStmtContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseSorterListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseSorterListener) ExitAttribute(ctx *AttributeContext) {}

// EnterPrio is called when production prio is entered.
func (s *BaseSorterListener) EnterPrio(ctx *PrioContext) {}

// ExitPrio is called when production prio is exited.
func (s *BaseSorterListener) ExitPrio(ctx *PrioContext) {}
