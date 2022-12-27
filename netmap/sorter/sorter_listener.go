// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Sorter

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SorterListener is a complete listener for a parse tree produced by Sorter.
type SorterListener interface {
	antlr.ParseTreeListener

	// EnterSorter is called when entering the sorter production.
	EnterSorter(c *SorterContext)

	// EnterWhenStmt is called when entering the whenStmt production.
	EnterWhenStmt(c *WhenStmtContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterPrio is called when entering the prio production.
	EnterPrio(c *PrioContext)

	// ExitSorter is called when exiting the sorter production.
	ExitSorter(c *SorterContext)

	// ExitWhenStmt is called when exiting the whenStmt production.
	ExitWhenStmt(c *WhenStmtContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitPrio is called when exiting the prio production.
	ExitPrio(c *PrioContext)
}
