%{
package main
%}

%union{
    annotations []Annotation
    annotation  Annotation
    fields      []Field
    field       Field
    exprs       []Expr
    expr        Expr
    token       Token
}

%type<annotations> annotations
%type<annotation>  annotation
%type<fields>      fields
%type<field>       field
%type<exprs>       exprs
%type<token>       expr
%token<token> IDENT
%token<token> INT FLOAT STRING
%token<token> TRUE FALSE
%token<token> LPAREN RPAREN LBRACE RBRACE
%token<token> EQUAL COMMA COLON CR

%%

annotations
    : annotation
    {
        $$ = []Annotation{$1}
        yylex.(*Lexer).result = $$
    }
    | annotations annotation
    {
        $$ = append($1, $2)
        yylex.(*Lexer).result = $$
    }

annotation
    : IDENT COLON IDENT
    {
        $$ = Annotation{Namespace: $1, Name: $3}
    }
    | IDENT COLON IDENT LPAREN fields RPAREN
    {
        $$ = Annotation{Namespace: $1, Name: $3, Fields: $5}
    }

fields
    : field
    {
        $$ = []Field{$1}
    }
    | fields COMMA field
    {
        $$ = append($1, $3)
    }

field
    : IDENT EQUAL LBRACE exprs RBRACE
    {
        $$ = Field{Name: $1, Expr: $4}
    }
    | IDENT EQUAL expr
    {
        $$ = Field{Name: $1, Expr: $3}
    }

expr
    : INT    { $$ = $1 }
    | FLOAT  { $$ = $1 }
    | STRING { $$ = $1 }
    | TRUE   { $$ = $1 }
    | FALSE  { $$ = $1 }
    | IDENT  { $$ = $1 }

exprs
    : expr
    {
        $$ = []Expr{$1}
    }
    | exprs COMMA expr
    {
        $$ = append($1, $3)
    }

%%
