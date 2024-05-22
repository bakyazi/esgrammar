grammar Query;

// Parser rules
query: searchClause whereClause? orderByClause? limitClause? offsetClause?  # SearchQuery
       | countClause whereClause?                                           # CountQuery
       ;

searchClause: SEARCH FROM indexName ;
countClause: COUNT FROM indexName;
whereClause: WHERE condition;

condition: (andCondition | orCondition);
andCondition: conditionPart (AND conditionPart)*?;
orCondition: conditionPart (OR conditionPart)*?;

orderByClause: ORDER BY orderCondition ( COMMA orderCondition )* ;
limitClause: LIMIT NUMBER ;
offsetClause: OFFSET NUMBER ;

conditionPart: fieldName comparator value                           # ComparisonCondition
         |  LPAR condition RPAR                                     # GroupCondition
         | fieldName NOT? LIKE STRING                               # LikeCondition
         | fieldName NOT? ISNULL                                    # IsNullCondition
         | fieldName NOT? IN LPAR value ( COMMA value )* RPAR       # InCondition
         ;

orderCondition: fieldName (ASC | DESC)? ;

indexName: IDENTIFIER ;
fieldName:  IDENTIFIER (DOT IDENTIFIER)* ;
value: NUMBER | STRING ;

// Lexer rules
SEARCH: 'SEARCH';
COUNT: 'COUNT';
FROM: 'FROM';
WHERE: 'WHERE';
ORDER: 'ORDER';
BY: 'BY';
LIMIT: 'LIMIT';
OFFSET: 'OFFSET';
AND: 'AND';
OR: 'OR';
NOT: 'NOT';
LIKE: 'LIKE';
IN: 'IN';
ISNULL: 'IS' 'NULL';
ASC: 'ASC';
DESC: 'DESC';

IDENTIFIER: [a-zA-Z_][a-zA-Z_0-9]* ;
NUMBER: [0-9]+ ;
STRING: '"' (~["\r\n] | '""')* '"' ;
WS: [ \t\r\n]+ -> skip ;
LPAR: '(';
RPAR: ')';
DOT: '.';
COMMA: ',';


// Comparison operators
comparator: EQUAL
 | NOT_EQUAL
 | GREATER
 | GREATER_EQ
 | LESSER
 | LESSER_EQ
 ;

EQUAL: '=';
NOT_EQUAL: '!=';
GREATER: '>';
LESSER: '<';
GREATER_EQ: '>=';
LESSER_EQ: '<=';