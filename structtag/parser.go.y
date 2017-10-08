%{

package main

%}

%union{
    token Token
    tags  []Tag
    tag   Tag
}

%type<tags> tags
%type<tag> tag
%type<token> literal
%token<token> KEY
%token<token> INT FLOAT STRING
%token<token> TRUE FALSE
%token<token> COMMA COLON

%%

tags
    : tag
    {
        $$ = []Tag{$1}
        yylex.(*Lexer).result = $$
    }
    | tags COMMA tag
    {
        $$ = append($1, $3)
        yylex.(*Lexer).result = $$
    }

tag
    : KEY COLON literal
    {
        $$ = Tag{key: $1.literal, value: $3}
    }

literal
    : INT    { $$ = $1 }
    | FLOAT  { $$ = $1 }
    | STRING { $$ = $1 }
    | TRUE   { $$ = $1 }
    | FALSE  { $$ = $1 }
%%
