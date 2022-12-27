parser grammar Sorter;

options {
    tokenVocab = SorterLexer;
}

sorter: CASE IDENT whenStmt* END;
whenStmt: WHEN Name = attribute Priority = prio;

attribute : IDENT | prio | STRING;
prio : ZERO | NUMBER1;
