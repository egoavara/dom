//
// This package is XPath Tokenizer(or Lexical scanner)
//
package compiler

import (
	. "github.com/iamGreedy/dom/grammer"
)

var GrammerXPath *Grammer

// shortfnname
var (
	base     = NewExpressionBaseToken
	refer    = NewExpressionRefer
	prefix   = NewExpressionPrefix
	regex    = MustExpressionRegexp
	and      = NewExpressionAnd
	or       = NewExpressionOr
	except   = NewExpressionExcept
	multiple = NewExpressionMultiple
	possible = NewExpressionPossible
)

//
// Fallow by github.com/iamGreedy/dom/example/XPath-Next/grammer/xpath-3.1.ebnf
// Find by github.com/XPath-Next/XPath-Next
//
// 'GrammerXPath' is Source code compiled version by parsing 'example/XPath-Next/grammer/xpath-3.1.ebnf'
//
func init() {
	gb := NewGrammerBuilder()
	xpathExpr(gb)
	advanceType(gb)
	basicTypes(gb)
	// Buildup
	var err error
	GrammerXPath, err = gb.Build()
	if err != nil {
		panic(err)
	}
}
func xpathExpr(gb *GrammerBuilder) {
	// XPath    		::= Expr
	gb.AddExpression("XPath", refer("Expr"))
	// ParamList 		::= Param ("," Param)*
	gb.AddExpression("ParamList", and(
		refer("Param"),
		multiple(and(
			prefix(","),
			refer("Param"),
		)),
	))
	// Param    		::= "$" EQName TypeDeclaration?
	gb.AddExpression("Param", and(
		prefix("$"),
		refer("EQName"),
		possible(refer("TypeDeclaration")),
	))
	// FunctionBody 	::= EnclosedExpr
	gb.AddExpression("FunctionBody", refer("EnclosedExpr"))
	// EnclosedExpr 	::= "{" Expr? "}"
	gb.AddExpression("EnclosedExpr", and(
		prefix("{"),
		possible(refer("EnclosedExpr")),
		prefix("}"),
	))

	// Expr 			::= ExprSingle ("," ExprSingle)*
	gb.AddExpression("Expr", and(
		refer("ExprSingle"),
		multiple(
			and(
				prefix(","),
				refer("ExprSingle"),
			),
		),
	))

	// ExprSingle 		::= ForExpr | LetExpr | QuantifiedExpr | IfExpr | OrExpr
	gb.AddExpression("ExprSingle", or(
		refer("ForExpr"),
		refer("LetExpr"),
		refer("QuantifiedExpr"),
		refer("IfExpr"),
		refer("OrExpr"),
	))
	// ForExpr
	forExpr(gb)
	// LetExpr
	letExpr(gb)
	// QuantifiedExpr
	quantifiedExpr(gb)
	// IfExpr
	ifExpr(gb)
	// OrExpr
	orExpr(gb)
	// VarName				::= EQName
	gb.AddExpression("VarName", refer("EQName"))
	// EQName				::= QName | URIQualifiedName
	gb.AddExpression("EQName", base(or(
		refer("QName"),
		refer("URIQualifiedName"),
	)))
}
func forExpr(gb *GrammerBuilder) {
	// ForExpr			::= SimpleForClause "return" ExprSingle
	// SimpleForClause ::= "for" SimpleForBinding ("," SimpleForBinding)*
	// SimpleForBinding ::= "$" VarName "in" ExprSingle
	gb.AddExpression("ForExpr", and(
		refer("SimpleForClause"),
		prefix("return"),
		refer("ExprSingle"),
	))
	gb.AddExpression("SimpleForClause", and(
		prefix("for"),
		refer("SimpleForBinding"),
		multiple(and(
			prefix(","),
			refer("SimpleForBinding"),
		)),
	))
	gb.AddExpression("SimpleForBinding", and(
		prefix("$"),
		refer("VarName"),
		prefix("in"),
		refer("ExprSingle"),
	))
}
func letExpr(gb *GrammerBuilder) {
	// LetExpr  ::= SimpleLetClause "return" ExprSingle
	// SimpleLetClause ::= "let" SimpleLetBinding ("," SimpleLetBinding)*
	// SimpleLetBinding ::= "$" VarName ":=" ExprSingle
	gb.AddExpression("LetExpr", and(
		refer("SimpleLetClause"),
		prefix("return"),
		refer("ExprSingle"),
	))
	gb.AddExpression("SimpleLetClause", and(
		prefix("let"),
		refer("SimpleLetBinding"),
		multiple(and(
			prefix(","),
			refer("SimpleLetBinding"),
		)),
	))
	gb.AddExpression("SimpleLetBinding", and(
		prefix("$"),
		refer("VarName"),
		prefix(":="),
		refer("ExprSingle"),
	))
}
func quantifiedExpr(gb *GrammerBuilder) {
	// QuantifiedExpr ::= ("some" | "every") "$" VarName "in" ExprSingle ("," "$" VarName "in" ExprSingle)* "satisfies" ExprSingle
	gb.AddExpression("QuantifiedExpr", and(
		or(
			prefix("some"),
			prefix("every"),
		),
		prefix("$"),
		refer("VarName"),
		prefix("in"),
		refer("ExprSingle"),
		multiple(and(
			prefix(","),
			prefix("$"),
			refer("VarName"),
			prefix("in"),
			refer("ExprSingle"),
		)),
		prefix("satisfies"),
		refer("ExprSingle"),
	))
}
func ifExpr(gb *GrammerBuilder) {
	// IfExpr   ::= "if" "(" Expr ")" "then" ExprSingle "else" ExprSingle
	gb.AddExpression("IfExpr", and(
		prefix("if"),
		prefix("("),
		refer("Expr"),
		prefix(")"),
		prefix("then"),
		refer("ExprSingle"),
		prefix("else"),
		refer("ExprSingle"),
	))
}
func orExpr(gb *GrammerBuilder) {
	// OrExpr 			::= AndExpr ( "or" AndExpr )*
	gb.AddExpression("OrExpr", and(
		refer("AndExpr"),
		multiple(and(
			prefix("or"),
			refer("AndExpr"),
		)),
	))
	// AndExpr 			::= ComparisonExpr ( "and" ComparisonExpr )*
	gb.AddExpression("AndExpr", and(
		refer("ComparisonExpr"),
		multiple(and(
			prefix("and"),
			refer("ComparisonExpr"),
		)),
	))
	// ComparisonExpr ::= StringConcatExpr ( (ValueComp | GeneralComp | NodeComp) StringConcatExpr )?
	gb.AddExpression("ComparisonExpr", and(
		refer("StringConcatExpr"),
		possible(and(
			or(
				refer("ValueComp"),
				refer("GeneralComp"),
				refer("NodeComp"),
			),
			refer("StringConcatExpr"),
		)),
	))
	// StringConcatExpr ::= RangeExpr ( "||" RangeExpr )*
	gb.AddExpression("StringConcatExpr", and(
		refer("RangeExpr"),
		multiple(and(
			prefix("||"),
			refer("RangeExpr"),
		)),
	))
	// RangeExpr ::= AdditiveExpr ( "to" AdditiveExpr )?
	gb.AddExpression("RangeExpr", and(
		refer("AdditiveExpr"),
		possible(and(
			prefix("to"),
			refer("AdditiveExpr"),
		)),
	))
	// AdditiveExpr ::= MultiplicativeExpr ( ("+" | "-") MultiplicativeExpr )*
	gb.AddExpression("AdditiveExpr", and(
		refer("MultiplicativeExpr"),
		multiple(and(
			or(
				prefix("+"),
				prefix("-"),
			),
			refer("MultiplicativeExpr"),
		)),
	))
	// MultiplicativeExpr ::= UnionExpr ( ("*" | "div" | "idiv" | "mod") UnionExpr )*
	gb.AddExpression("MultiplicativeExpr", and(
		refer("UnionExpr"),
		multiple(and(
			or(
				prefix("*"),
				prefix("div"),
				prefix("idiv"),
				prefix("mod"),
			),
			refer("UnionExpr"),
		)),
	))
	// UnionExpr ::= IntersectExceptExpr ( ("union" | "|") IntersectExceptExpr )*
	gb.AddExpression("UnionExpr", and(
		refer("IntersectExceptExpr"),
		multiple(and(
			or(
				prefix("union"),
				prefix("|"),
			),
			refer("IntersectExceptExpr"),
		)),
	))
	// IntersectExceptExpr ::= InstanceofExpr ( ("intersect" | "except") InstanceofExpr )*
	gb.AddExpression("IntersectExceptExpr", and(
		refer("InstanceofExpr"),
		multiple(and(
			or(
				prefix("intersect"),
				prefix("except"),
			),
			refer("InstanceofExpr"),
		)),
	))
	// InstanceofExpr ::= TreatExpr ( "instance" "of" SequenceType )?
	gb.AddExpression("InstanceofExpr", and(
		refer("TreatExpr"),
		possible(and(
			prefix("instance"),
			prefix("of"),
			refer("SequenceType"),
		)),
	))
	// TreatExpr ::= CastableExpr ( "treat" "as" SequenceType )?
	gb.AddExpression("TreatExpr", and(
		refer("CastableExpr"),
		possible(and(
			prefix("treat"),
			prefix("as"),
			refer("SequenceType"),
		)),
	))
	// CastableExpr ::= CastExpr ( "castable" "as" SingleType )?
	gb.AddExpression("CastableExpr", and(
		refer("CastExpr"),
		possible(and(
			prefix("castable"),
			prefix("as"),
			refer("SingleType"),
		)),
	))
	// CastExpr ::= ArrowExpr ( "cast" "as" SingleType )?
	gb.AddExpression("CastExpr", and(
		refer("ArrowExpr"),
		possible(and(
			prefix("cast"),
			prefix("as"),
			refer("SingleType"),
		)),
	))
	// ArrowExpr ::= UnaryExpr ( "=>" ArrowFunctionSpecifier ArgumentList )*
	gb.AddExpression("ArrowExpr", and(
		refer("UnaryExpr"),
		multiple(and(
			prefix("=>"),
			refer("ArrowFunctionSpecifier"),
			refer("ArgumentList"),
		)),
	))
	// UnaryExpr ::= ("-" | "+")* ValueExpr
	gb.AddExpression("UnaryExpr", and(
		multiple(or(
			prefix("-"),
			prefix("+"),
		)),
		refer("ValueExpr"),
	))
	// ValueExpr ::= SimpleMapExpr
	gb.AddExpression("ValueExpr", refer("SimpleMapExpr"))
	// GeneralComp ::= "=" | "!=" | "<" | "<=" | ">" | ">="
	gb.AddExpression("GeneralComp", or(
		prefix("="),
		prefix("!="),
		prefix("<"),
		prefix("<="),
		prefix(">"),
		prefix(">="),
	))
	// ValueComp ::= "eq" | "ne" | "lt" | "le" | "gt" | "ge"
	gb.AddExpression("ValueComp", or(
		prefix("eq"),
		prefix("ne"),
		prefix("lt"),
		prefix("le"),
		prefix("gt"),
		prefix("ge"),
	))
	// NodeComp ::= "is" | "<<" | ">>"
	gb.AddExpression("NodeComp", or(
		prefix("is"),
		prefix("<<"),
		prefix(">>"),
	))
	// SimpleMapExpr ::= PathExpr ("!" PathExpr)*
	gb.AddExpression("SimpleMapExpr", and(
		refer("PathExpr"),
		multiple(and(
			prefix("!"),
			refer("PathExpr"),
		)),
	))
	pathExpr(gb)
}
func pathExpr(gb *GrammerBuilder) {
	// PathExpr ::= ("/" RelativePathExpr?) | ("//" RelativePathExpr) | RelativePathExpr /* xgc: leading-lone-slash */
	gb.AddExpression("PathExpr", or(
		and(
			prefix("/"),
			possible(refer("RelativePathExpr")),
		),
		and(
			prefix("//"),
			refer("RelativePathExpr"),
		),
		refer("RelativePathExpr"),
	))
	// RelativePathExpr ::= StepExpr (("/" | "//") StepExpr)*
	gb.AddExpression("RelativePathExpr", and(
		refer("StepExpr"),
		multiple(and(
			or(
				prefix("/"),
				prefix("//"),
			),
			refer("RelativePathExpr"),
		)),
	))
	// StepExpr ::= PostfixExpr | AxisStep
	gb.AddExpression("StepExpr", or(
		refer("PostfixExpr"),
		refer("AxisStep"),
	))
	// AxisStep ::= (ReverseStep | ForwardStep) PredicateList
	gb.AddExpression("AxisStep", and(
		or(
			refer("ReverseStep"),
			refer("ForwardStep"),
		),
		refer("PredicateList"),
	))
	// ForwardStep ::= (ForwardAxis NodeTest) | AbbrevForwardStep
	gb.AddExpression("ForwardStep", or(
		and(
			refer("ForwardAxis"),
			refer("NodeTest"),
		),
		refer("AbbrevForwardStep"),
	))
	// ForwardAxis ::= ("child" "::") | ("descendant" "::") | ("attribute" "::") | ("self" "::") | ("descendant-or-self" "::") | ("following-sibling" "::") | ("following" "::") | ("namespace" "::")
	gb.AddExpression("ForwardAxis", or(
		and(
			prefix("child"),
			prefix("::"),
		),
		and(
			prefix("descendant"),
			prefix("::"),
		),
		and(
			prefix("attribute"),
			prefix("::"),
		),
		and(
			prefix("self"),
			prefix("::"),
		),
		and(
			prefix("descendant-or-self"),
			prefix("::"),
		),
		and(
			prefix("following-sibling"),
			prefix("::"),
		),
		and(
			prefix("following"),
			prefix("::"),
		),
		and(
			prefix("namespace"),
			prefix("::"),
		),
	))
	// AbbrevForwardStep ::= "@"? NodeTest
	gb.AddExpression("AbbrevForwardStep", and(
		possible(prefix("@")),
		refer("NodeTest"),
	))
	// ReverseStep ::= (ReverseAxis NodeTest) | AbbrevReverseStep
	gb.AddExpression("ReverseStep", or(
		and(
			refer("ReverseAxis"),
			refer("NodeTest"),
		),
		refer("AbbrevReverseStep"),
	))
	// ReverseAxis ::= ("parent" "::") | ("ancestor" "::") | ("preceding-sibling" "::") | ("preceding" "::") | ("ancestor-or-self" "::")
	gb.AddExpression("ReverseAxis", or(
		and(
			prefix("parent"),
			prefix("::"),
		),
		and(
			prefix("ancestor"),
			prefix("::"),
		),
		and(
			prefix("preceding-sibling"),
			prefix("::"),
		),
		and(
			prefix("preceding"),
			prefix("::"),
		),
		and(
			prefix("ancestor-or-self"),
			prefix("::"),
		),
	))
	// AbbrevReverseStep ::= ".."
	gb.AddExpression("AbbrevReverseStep", prefix(".."))
	// NodeTest ::= KindTest | NameTest
	gb.AddExpression("NodeTest", or(
		refer("KindTest"),
		refer("NameTest"),
	))
	// NameTest ::= EQName | Wildcard
	gb.AddExpression("NameTest", or(
		refer("EQName"),
		refer("Wildcard"),
	))
	// Wildcard ::= "*" | (NCName ":*") | ("*:" NCName) | (BracedURILiteral "*") /* ws: explicit */
	gb.AddExpression("Wildcard", or(
		prefix("*"),
		and(
			refer("NCName"),
			prefix(":*"),
		),
		and(
			prefix("*:"),
			refer("NCName"),
		),
		and(
			refer("BracedURILiteral"),
			prefix("*"),
		),
	))
	// PostfixExpr ::= PrimaryExpr (Predicate | ArgumentList | Lookup)*
	gb.AddExpression("PostfixExpr", and(
		refer("PrimaryExpr"),
		multiple(or(
			refer("Predicate"),
			refer("ArgumentList"),
			refer("Lookup"),
		)),
	))
	// Argument ::= ExprSingle | ArgumentPlaceholder
	gb.AddExpression("Argument", or(
		refer("ExprSingle"),
		refer("ArgumentPlaceholder"),
	))
	// ArgumentPlaceholder ::= "?"
	gb.AddExpression("ArgumentPlaceholder", prefix("?"))
	// ArgumentList ::= "(" (Argument ("," Argument)*)? ")"
	gb.AddExpression("ArgumentList", and(
		prefix("("),
		possible(and(
			refer("Argument"),
			multiple(and(
				prefix(","),
				refer("Argument"),
			)),
		)),
		prefix(")"),
	))
	// PredicateList ::= Predicate*
	gb.AddExpression("PredicateList", multiple(refer("Predicate")))
	// Predicate ::= "[" Expr "]"
	gb.AddExpression("Predicate", and(
		prefix("["),
		refer("Expr"),
		prefix("]"),
	))
	// Lookup   ::= "?" KeySpecifier
	gb.AddExpression("Lookup", and(
		prefix("?"),
		refer("KeySpecifier"),
	))
	// KeySpecifier ::= NCName | IntegerLiteral | ParenthesizedExpr | "*"
	gb.AddExpression("KeySpecifier", or(
		refer("NCName"),
		refer("IntegerLiteral"),
		refer("ParenthesizedExpr"),
		prefix("*"),
	))
	// ArrowFunctionSpecifier ::= EQName | VarRef | ParenthesizedExpr
	gb.AddExpression("ArrowFunctionSpecifier", or(
		refer("EQName"),
		refer("VarRef"),
		refer("ParenthesizedExpr"),
	))
	primaryExpr(gb)
}
func primaryExpr(gb *GrammerBuilder) {
	// PrimaryExpr ::= Literal | VarRef | ParenthesizedExpr | ContextItemExpr | FunctionCall | FunctionItemExpr | MapConstructor | ArrayConstructor | UnaryLookup
	gb.AddExpression("PrimaryExpr", or(
		refer("Literal"),
		refer("VarRef"),
		refer("ParenthesizedExpr"),
		refer("ContextItemExpr"),
		refer("FunctionCall"),
		refer("FunctionItemExpr"),
		refer("MapConstructor"),
		refer("ArrayConstructor"),
		refer("UnaryLookup"),
	))

	// Literal  ::= NumericLiteral | StringLiteral
	gb.AddExpression("Literal", or(
		refer("NumericLiteral"),
		refer("StringLiteral"),
	))

	// NumericLiteral ::= IntegerLiteral | DecimalLiteral | DoubleLiteral
	gb.AddExpression("NumericLiteral", or(
		refer("IntegerLiteral"),
		refer("DecimalLiteral"),
		refer("DoubleLiteral"),
	))

	// VarRef   ::= "$" VarName
	gb.AddExpression("VarRef", and(
		prefix("$"),
		refer("VarName"),
	))

	// VarName  ::= EQName
	gb.AddExpression("VarName", refer("EQName"))

	// ParenthesizedExpr ::= "(" Expr? ")"
	gb.AddExpression("ParenthesizedExpr", and(
		prefix("("),
		possible(refer("Expr")),
		prefix(")"),
	))

	// ContextItemExpr ::= "."
	gb.AddExpression("ContextItemExpr", prefix("."))

	// FunctionCall ::= EQName ArgumentList /* xgc: reserved-function-names */ /* gn: parens */
	gb.AddExpression("FunctionCall", and(
		refer("EQName"),
		refer("ArgumentList"),
	))

	// FunctionItemExpr ::= NamedFunctionRef | InlineFunctionExpr
	gb.AddExpression("FunctionItemExpr", or(
		refer("NamedFunctionRef"),
		refer("InlineFunctionExpr"),
	))

	// NamedFunctionRef ::= EQName "#" IntegerLiteral /* xgc: reserved-function-names */
	gb.AddExpression("NamedFunctionRef", and(
		refer("EQName"),
		prefix("#"),
		refer("IntegerLiteral"),
	))
	// InlineFunctionExpr ::= "function" "(" ParamList? ")" ("as" SequenceType)? FunctionBody
	gb.AddExpression("InlineFunctionExpr", and(
		prefix("function"),
		prefix("("),
		possible(refer("ParamList")),
		prefix(")"),
		possible(and(
			prefix("as"),
			refer("SequenceType"),
		)),
		refer("FunctionBody"),
	))
	// MapConstructor ::= "map" "{" (MapConstructorEntry ("," MapConstructorEntry)*)? "}"
	gb.AddExpression("MapConstructor", and(
		prefix("map"),
		prefix("{"),
		possible(and(
			refer("MapConstructorEntry"),
			multiple(and(
				prefix(","),
				refer("MapConstructorEntry"),
			)),
		)),
		prefix("}"),
	))
	// MapConstructorEntry ::= MapKeyExpr ":" MapValueExpr
	gb.AddExpression("MapConstructorEntry", and(
		refer("MapKeyExpr"),
		prefix(":"),
		refer("MapValueExpr"),
	))
	// MapKeyExpr ::= ExprSingle
	gb.AddExpression("MapKeyExpr", refer("ExprSingle"))
	// MapValueExpr ::= ExprSingle
	gb.AddExpression("MapValueExpr", refer("ExprSingle"))
	// ArrayConstructor ::= SquareArrayConstructor | CurlyArrayConstructor
	gb.AddExpression("ArrayConstructor", or(
		refer("SquareArrayConstructor"),
		refer("CurlyArrayConstructor"),
	))
	// SquareArrayConstructor ::= "[" (ExprSingle ("," ExprSingle)*)? "]"
	gb.AddExpression("SquareArrayConstructor", and(
		prefix("["),
		possible(and(
			refer("ExprSingle"),
			multiple(and(
				prefix(","),
				refer("ExprSingle"),
			)),
		)),
		prefix("]"),
	))
	// CurlyArrayConstructor ::= "array" EnclosedExpr
	gb.AddExpression("CurlyArrayConstructor", and(
		prefix("array"),
		refer("EnclosedExpr"),
	))
	// UnaryLookup ::= "?" KeySpecifier
	gb.AddExpression("UnaryLookup", and(
		prefix("?"),
		refer("KeySpecifier"),
	))
	// SingleType ::= SimpleTypeName "?"?
	gb.AddExpression("SingleType", and(
		refer("SimpleTypeName"),
		possible(prefix("?")),
	))
	// TypeDeclaration ::= "as" SequenceType
	gb.AddExpression("TypeDeclaration", and(
		prefix("as"),
		refer("SequenceType"),
	))
}
func advanceType(gb *GrammerBuilder) {
	// SequenceType 			::= ("empty-sequence" "(" ")") | (ItemType OccurrenceIndicator?)
	gb.AddExpression("SequenceType", or(
		and(
			prefix("empty-sequence"),
			prefix("("),
			prefix(")"),
		),
		and(
			refer("ItemType"),
			possible(refer("OccurrenceIndicator")),
		),
	))
	// OccurrenceIndicator 		::= "?" | "*" | "+"  /* xgc: occurrence-indicators */
	gb.AddExpression("OccurrenceIndicator", or(
		prefix("?"),
		prefix("*"),
		prefix("+"),
	))
	// ItemType ::= KindTest | ("item" "(" ")") | FunctionTest | MapTest | ArrayTest | AtomicOrUnionType | ParenthesizedItemType
	gb.AddExpression("ItemType", or(
		refer("KindTest"),
		and(
			prefix("item"),
			prefix("("),
			prefix(")"),
		),
		refer("FunctionTest"),
		refer("MapTest"),
		refer("ArrayTest"),
		refer("AtomicOrUnionType"),
		refer("ParenthesizedItemType"),
	))
	// AtomicOrUnionType ::= EQName
	gb.AddExpression("AtomicOrUnionType", refer("EQName"))
	// KindTest ::= DocumentTest | ElementTest | AttributeTest | SchemaElementTest | SchemaAttributeTest | PITest | CommentTest | TextTest | NamespaceNodeTest | AnyKindTest
	gb.AddExpression("KindTest", or(
		refer("DocumentTest"),
		refer("ElementTest"),
		refer("AttributeTest"),
		refer("SchemaElementTest"),
		refer("SchemaAttributeTest"),
		refer("PITest"),
		refer("CommentTest"),
		refer("TextTest"),
		refer("NamespaceNodeTest"),
		refer("AnyKindTest"),
	))
	// AnyKindTest ::= "node" "(" ")"
	gb.AddExpression("AnyKindTest", and(
		prefix("node"),
		prefix("("),
		prefix(")"),
	))
	// DocumentTest ::= "document-node" "(" (ElementTest | SchemaElementTest)? ")"
	gb.AddExpression("DocumentTest", and(
		prefix("document-node"),
		prefix("("),
		possible(
			or(
				refer("ElementTest"),
				refer("SchemaElementTest"),
			),
		),
		prefix(")"),
	))
	// TextTest ::= "text" "(" ")"
	gb.AddExpression("TextTest", and(
		prefix("text"),
		prefix("("),
		prefix(")"),
	))
	// CommentTest ::= "comment" "(" ")"
	gb.AddExpression("CommentTest", and(
		prefix("comment"),
		prefix("("),
		prefix(")"),
	))
	// NamespaceNodeTest ::= "namespace-node" "(" ")"
	gb.AddExpression("NamespaceNodeTest", and(
		prefix("namespace-node"),
		prefix("("),
		prefix(")"),
	))
	// PITest   ::= "processing-instruction" "(" (NCName | StringLiteral)? ")"
	gb.AddExpression("PITest", and(
		prefix("processing-instruction"),
		prefix("("),
		possible(
			or(
				refer("NCName"),
				refer("StringLiteral"),
			),
		),
		prefix(")"),
	))
	// AttributeTest ::= "attribute" "(" (AttribNameOrWildcard ("," TypeName)?)? ")"
	gb.AddExpression("AttributeTest", and(
		prefix("attribute"),
		prefix("("),
		possible(and(
			refer("AttribNameOrWildcard"),
			possible(and(prefix(","),
				refer("AttribNameOrWildcard"),
			)),
		)),
		prefix(")"),
	))
	// AttribNameOrWildcard ::= AttributeName | "*"
	gb.AddExpression("AttribNameOrWildcard", or(
		refer("AttributeName"),
		prefix("*"),
	))
	// SchemaAttributeTest ::= "schema-attribute" "(" AttributeDeclaration ")"
	gb.AddExpression("SchemaAttributeTest", and(
		prefix("schema-attribute"),
		prefix("("),
		refer("AttributeDeclaration"),
		prefix(")"),
	))
	// AttributeDeclaration ::= AttributeName
	gb.AddExpression("AttributeDeclaration", refer("AttributeName"))
	// ElementTest ::= "element" "(" (ElementNameOrWildcard ("," TypeName "?"?)?)? ")"
	gb.AddExpression("ElementTest", and(
		prefix("element"),
		prefix("("),
		possible(and(
			refer("ElementNameOrWildcard"),
			possible(and(
				prefix(","),
				refer("TypeName"),
				possible(prefix("?")),
			)),
		)),
		prefix(")"),
	))
	// ElementNameOrWildcard ::= ElementName | "*"
	gb.AddExpression("ElementNameOrWildcard", or(
		refer("ElementName"),
		prefix("*"),
	))
	// SchemaElementTest ::= "schema-element" "(" ElementDeclaration ")"
	gb.AddExpression("SchemaElementTest", and(
		prefix("schema-element"),
		prefix("("),
		refer("ElementDeclaration"),
		prefix(")"),
	))
	// ElementDeclaration ::= ElementName
	gb.AddExpression("ElementDeclaration", refer("ElementName"))
	// AttributeName ::= EQName
	gb.AddExpression("AttributeName", refer("EQName"))
	// ElementName ::= EQName
	gb.AddExpression("ElementName", refer("EQName"))
	// SimpleTypeName ::= TypeName
	gb.AddExpression("SimpleTypeName", refer("TypeName"))
	// TypeName ::= EQName
	gb.AddExpression("TypeName", refer("EQName"))
	// FunctionTest ::= AnyFunctionTest | TypedFunctionTest
	gb.AddExpression("FunctionTest", or(
		refer("AnyFunctionTest"),
		refer("TypedFunctionTest"),
	))
	// AnyFunctionTest ::= "function" "(" "*" ")"
	gb.AddExpression("AnyFunctionTest", and(
		prefix("function"),
		prefix("("),
		prefix("*"),
		prefix(")"),
	))
	// TypedFunctionTest ::= "function" "(" (SequenceType ("," SequenceType)*)? ")" "as" SequenceType
	gb.AddExpression("TypedFunctionTest", and(
		prefix("function"),
		prefix("("),
		possible(and(
			refer("SequenceType"),
			multiple(and(
				prefix(","),
				refer("SequenceType"),
			)),
		)),
		prefix(")"),
		prefix("as"),
		refer("SequenceType"),
	))
	// MapTest  ::= AnyMapTest | TypedMapTest
	gb.AddExpression("MapTest", or(
		refer("AnyMapTest"),
		refer("TypedMapTest"),
	))
	// AnyMapTest ::= "map" "(" "*" ")"
	gb.AddExpression("AnyMapTest", and(
		prefix("map"),
		prefix("("),
		prefix("*"),
		prefix(")"),
	))
	// TypedMapTest ::= "map" "(" AtomicOrUnionType "," SequenceType ")"
	gb.AddExpression("TypedMapTest", and(
		prefix("map"),
		prefix("("),
		refer("AtomicOrUnionType"),
		prefix(","),
		refer("SequenceType"),
		prefix(")"),
	))
	// ArrayTest ::= AnyArrayTest | TypedArrayTest
	gb.AddExpression("ArrayTest", or(
		refer("AnyArrayTest"),
		refer("TypedArrayTest"),
	))
	// AnyArrayTest ::= "array" "(" "*" ")"
	gb.AddExpression("AnyArrayTest", and(
		prefix("array"),
		prefix("("),
		prefix("*"),
		prefix(")"),
	))
	// TypedArrayTest ::= "array" "(" SequenceType ")"
	gb.AddExpression("TypedArrayTest", and(
		prefix("map"),
		prefix("("),
		refer("SequenceType"),
		prefix(")"),
	))
	// ParenthesizedItemType ::= "(" ItemType ")"
	gb.AddExpression("ParenthesizedItemType", and(
		prefix("("),
		refer("ItemType"),
		prefix(")"),
	))
}
func basicTypes(gb *GrammerBuilder) {
	// IntegerLiteral 		::= Digits
	gb.AddExpression("IntegerLiteral", refer("Digits"))
	// DecimalLiteral 		::= ("." Digits) | (Digits "." [0-9]*) /* ws: explicit */
	gb.AddExpression("DecimalLiteral",
		or(
			and(prefix("."), refer("Digits")),
			and(refer("Digits"), prefix("."), regex(`[0-9]*`)),
		),
	)
	// DoubleLiteral 		::= (("." Digits) | (Digits ("." [0-9]*)?)) [eE] [+-]? Digits /* ws: explicit */
	gb.AddExpression("DoubleLiteral",
		and(
			or(
				and(
					prefix("."),
					refer("Digits")),
				and(
					refer("Digits"),
					possible(
						and(
							prefix("."),
							regex(`[0-9]*`)),
					),
				),
			),
			regex(`[eE]`),
			regex(`[+-]?`),
			refer("Digits"),
		),
	)
	// StringLiteral 		::= ('"' (EscapeQuot | [^"])* '"') | ("'" (EscapeApos | [^'])* "'") /* ws: explicit */
	gb.AddExpression("StringLiteral",
		or(
			and(
				prefix(`"`),
				multiple(or(refer("EscapeQuot"), regex(`[^"]`))),
				prefix(`"`),
			),
			and(
				prefix(`'`),
				multiple(or(refer("EscapeApos"), regex(`[^']`))),
				prefix(`'`),
			),
		),
	)
	// URIQualifiedName  	::= BracedURILiteral NCName /* ws: explicit */
	gb.AddExpression("URIQualifiedName",
		and(
			refer("BracedURILiteral"),
			refer("NCName"),
		),
	)
	// BracedURILiteral 	::= "Q" "{" [^{}]* "}" /* ws: explicit */
	gb.AddExpression("BracedURILiteral",
		and(
			prefix(`Q`),
			prefix(`{`),
			regex(`[^{}]*`),
			prefix(`}`),
		),
	)
	// EscapeQuot 			::= '""'
	gb.AddExpression("EscapeQuot", prefix(`""`))
	// EscapeApos 			::= "''"
	gb.AddExpression("EscapeApos", prefix(`''`))
	// Comment  			::= "(:" (CommentContents | Comment)* ":)" /* ws: explicit */ /* gn: comments */
	gb.AddExpression("BracedURILiteral",
		and(
			prefix(`(:`),
			multiple(or(
				refer("CommentContents"),
				refer("Comment"),
			)),
			prefix(`:)`),
		),
	)
	// QName    			::= [http://www.w3.org/TR/REC-xml-names/#NT-QName] /* xgc: xml-version */
	//
	// [http://www.w3.org/TR/REC-xml-names/#NT-QName]
	// QName				::= PrefixedName | UnprefixedName
	// PrefixedName			::= Prefix ':' LocalPart
	// UnprefixedName		::= LocalPart
	// Prefix				::= NCName
	// LocalPart			::= NCName
	gb.AddExpression("QName",
		or(
			refer("PrefixedName"),
			refer("UnprefixedName"),
		),
	)
	gb.AddExpression("PrefixedName",
		and(
			refer("Prefix"),
			prefix(`:`),
			refer("LocalPart"),
		),
	)
	gb.AddExpression("UnprefixedName", refer("LocalPart"))
	gb.AddExpression("Prefix", refer("NCName"))
	gb.AddExpression("LocalPart", refer("NCName"))

	// NCName   			::= [http://www.w3.org/TR/REC-xml-names/#NT-NCName] /* xgc: xml-version */
	//
	// [http://www.w3.org/TR/REC-xml-names/#NT-NCName]
	// NCName 				::= Name - (Char* ':' Char*)
	// NameStartChar	   	::= ":" | [A-Z] | "_" | [a-z] | [\xC0-\xD6] | [\xD8-\xF6] | [\xF8-\x2FF] | [\x370-\x37D] | [\x37F-\x1FFF] | [\x200C-\x200D] | [\x2070-\x218F] | [\x2C00-\x2FEF] | [\x3001-\xD7FF] | [\xF900-\xFDCF] | [\xFDF0-\xFFFD] | [\x10000-\xEFFFF]
	// NameChar	   			::= NameStartChar | "-" | "." | [0-9] | \xB7 | [\x0300-\x036F] | [\x203F-\x2040]
	// Name					::= NameStartChar (NameChar)*
	gb.AddExpression("NCName", except(
		refer("Name"),
		and(
			multiple(refer("Char")),
			prefix(`:`),
			multiple(refer("Char")),
		),
	))
	gb.AddExpression("NameStartChar", or(
		prefix(":"),
		regex(`[A-Z]`),
		prefix("_"),
		regex(`[a-z]`),
		regex("[\u00C0-\u00D6]"),
		regex("[\u00D8-\u00F6]"),
		regex("[\u00F8-\u02FF]"),
		regex("[\u0370-\u037D]"),
		regex("[\u037F-\u1FFF]"),
		regex("[\u200C-\u200D]"),
		regex("[\u2070-\u218F]"),
		regex("[\u2C00-\u2FEF]"),
		regex("[\u3001-\uD7FF]"),
		regex("[\uF900-\uFDCF]"),
		regex("[\uFDF0-\uFFFD]"),
		regex("["+string(rune(0x010000))+"-"+string(rune(0x0EFFFF))+"]"),
	))
	gb.AddExpression("NameChar", or(
		refer("NameStartChar"),
		prefix("-"),
		prefix("."),
		regex(`[0-9]`),
		regex("\u00B7"),
		regex("[\u0300-\u036F]"),
		regex("[\u203F-\u2040]"),
	))
	gb.AddExpression("Name", and(
		refer("NameStartChar"),
		multiple(refer("NameChar")),
	))
	// Char     			::= [http://www.w3.org/TR/REC-xml#NT-Char] /* xgc: xml-version */
	//
	// [http://www.w3.org/TR/REC-xml#NT-Char]
	// Char					::= \x9 | \xA | \xD | [\x20-\xD7FF] | [\xE000-\xFFFD] | [\x10000-\x10FFFF]
	gb.AddExpression("Char", or(
		regex("\u0009"),
		regex("\u000A"),
		regex("[\u0020-\uD7FF]"),
		regex("[\uE000-\uFFFD]"),
		regex("["+string(rune(0x010000))+"-"+string(rune(0x10FFFF))+"]"),
	))
	// Digits   			::= [0-9]+
	gb.AddExpression("Digits", regex(`[0-9]+`))
	// CommentContents 		::= (Char+ - (Char* ('(:' | ':)') Char*))
}
