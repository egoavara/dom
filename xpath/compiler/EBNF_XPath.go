//
// This package is XPath Tokenizer(or Lexical scanner)
//
package compiler

import . "github.com/iamGreedy/dom/grammer"

var GrammerXPath *Grammer

//
// Fallow by github.com/iamGreedy/dom/example/XPath-Next/grammer/xpath-3.1.ebnf
// Find by github.com/XPath-Next/XPath-Next
//
// 'GrammerXPath' is Source code compiled version by parsing 'example/XPath-Next/grammer/xpath-3.1.ebnf'
//
func init() {
	gb := NewGrammerBuilder()
	basicTypes(gb)
	// Buildup
	var err error
	GrammerXPath, err = gb.Build()
	if err != nil {
		panic(err)
	}
}


func basicTypes(gb *GrammerBuilder)  {
	// IntegerLiteral 		::= Digits
	gb.AddExpression("IntegerLiteral", NewExpressionRefer("Digits"))
	// DecimalLiteral 		::= ("." Digits) | (Digits "." [0-9]*) /* ws: explicit */
	gb.AddExpression("DecimalLiteral",
		NewExpressionOr(
			NewExpressionAnd(NewExpressionPrefix("."), NewExpressionRefer("Digits")),
			NewExpressionAnd(NewExpressionRefer("Digits"), NewExpressionPrefix("."), MustExpressionRegexp(`[0-9]*`)),
		),
	)
	// DoubleLiteral 		::= (("." Digits) | (Digits ("." [0-9]*)?)) [eE] [+-]? Digits /* ws: explicit */
	gb.AddExpression("DoubleLiteral",
		NewExpressionAnd(
			NewExpressionOr(
				NewExpressionAnd(
					NewExpressionPrefix("."),
					NewExpressionRefer("Digits")),
				NewExpressionAnd(
					NewExpressionRefer("Digits"),
					NewExpressionPossible(
						NewExpressionAnd(
							NewExpressionPrefix("."),
							MustExpressionRegexp(`[0-9]*`)),
					),
				),
			),
			MustExpressionRegexp(`[eE]`),
			MustExpressionRegexp(`[+-]?`),
			NewExpressionRefer("Digits"),
		),
	)
	// StringLiteral 		::= ('"' (EscapeQuot | [^"])* '"') | ("'" (EscapeApos | [^'])* "'") /* ws: explicit */
	gb.AddExpression("StringLiteral",
		NewExpressionOr(
			NewExpressionAnd(
				NewExpressionPrefix(`"`),
				NewExpressionMultiple(NewExpressionOr(NewExpressionRefer("EscapeQuot"), MustExpressionRegexp(`[^"]`)),),
				NewExpressionPrefix(`"`),
			),
			NewExpressionAnd(
				NewExpressionPrefix(`'`),
				NewExpressionMultiple(NewExpressionOr(NewExpressionRefer("EscapeApos"), MustExpressionRegexp(`[^']`)),),
				NewExpressionPrefix(`'`),
			),
		),
	)
	// URIQualifiedName  	::= BracedURILiteral NCName /* ws: explicit */
	gb.AddExpression("StringLiteral",
		NewExpressionAnd(
			NewExpressionRefer("BracedURILiteral"),
			NewExpressionRefer("NCName"),
		),
	)
	// BracedURILiteral 	::= "Q" "{" [^{}]* "}" /* ws: explicit */
	gb.AddExpression("BracedURILiteral",
		NewExpressionAnd(
			NewExpressionPrefix(`Q`),
			NewExpressionPrefix(`{`),
			MustExpressionRegexp(`[^{}]*`),
			NewExpressionPrefix(`}`),
		),
	)
	// EscapeQuot 			::= '""'
	gb.AddExpression("EscapeQuot", NewExpressionPrefix(`""`), )
	// EscapeApos 			::= "''"
	gb.AddExpression("EscapeQuot", NewExpressionPrefix(`''`), )
	// Comment  			::= "(:" (CommentContents | Comment)* ":)" /* ws: explicit */ /* gn: comments */
	gb.AddExpression("BracedURILiteral",
		NewExpressionAnd(
			NewExpressionPrefix(`(:`),
			NewExpressionMultiple(NewExpressionOr(
				NewExpressionRefer("CommentContents"),
				NewExpressionRefer("Comment"),
			)),
			NewExpressionPrefix(`:)`),
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
		NewExpressionOr(
			NewExpressionRefer("PrefixedName"),
			NewExpressionRefer("UnprefixedName"),
		),
	)
	gb.AddExpression("PrefixedName",
		NewExpressionAnd(
			NewExpressionRefer("Prefix"),
			NewExpressionPrefix(`:`),
			NewExpressionRefer("LocalPart"),
		),
	)
	gb.AddExpression("UnprefixedName", NewExpressionRefer("LocalPart"), )
	gb.AddExpression("Prefix", NewExpressionRefer("NCName"), )
	gb.AddExpression("LocalPart", NewExpressionRefer("NCName"), )

	// NCName   			::= [http://www.w3.org/TR/REC-xml-names/#NT-NCName] /* xgc: xml-version */
	//
	// [http://www.w3.org/TR/REC-xml-names/#NT-NCName]
	// NCName 				::= Name - (Char* ':' Char*)
	// NameStartChar	   	::= ":" | [A-Z] | "_" | [a-z] | [#xC0-#xD6] | [#xD8-#xF6] | [#xF8-#x2FF] | [#x370-#x37D] | [#x37F-#x1FFF] | [#x200C-#x200D] | [#x2070-#x218F] | [#x2C00-#x2FEF] | [#x3001-#xD7FF] | [#xF900-#xFDCF] | [#xFDF0-#xFFFD] | [#x10000-#xEFFFF]
	// NameChar	   			::= NameStartChar | "-" | "." | [0-9] | #xB7 | [#x0300-#x036F] | [#x203F-#x2040]
	// Name					::= NameStartChar (NameChar)*
	gb.AddExpression("NCName", NewExpressionExcept(
		NewExpressionRefer("Name"),
		NewExpressionAnd(
			NewExpressionMultiple(NewExpressionRefer("Char")),
			NewExpressionPrefix(`:`),
			NewExpressionMultiple(NewExpressionRefer("Char")),
		),
	))
	gb.AddExpression("NameStartChar", NewExpressionOr(
		NewExpressionPrefix(":"),
		MustExpressionRegexp(`[A-Z]`),
		NewExpressionPrefix("_"),
		MustExpressionRegexp(`[a-z]`),
		MustExpressionRegexp(`[#xC0-#xD6]`),
		MustExpressionRegexp(`[#xD8-#xF6]`),
		MustExpressionRegexp(`[#xF8-#x2FF]`),
		MustExpressionRegexp(`[#x370-#x37D]`),
		MustExpressionRegexp(`[#x37F-#x1FFF]`),
		MustExpressionRegexp(`[#x200C-#x200D]`),
		MustExpressionRegexp(`[#x2070-#x218F]`),
		MustExpressionRegexp(`[#x2C00-#x2FEF]`),
		MustExpressionRegexp(`[#x3001-#xD7FF]`),
		MustExpressionRegexp(`[#xF900-#xFDCF]`),
		MustExpressionRegexp(`[#xFDF0-#xFFFD]`),
		MustExpressionRegexp(`[#x10000-#xEFFFF]`),
	))
	gb.AddExpression("NameChar", NewExpressionOr(
		NewExpressionRefer("NameStartChar"),
		NewExpressionPrefix("-"),
		NewExpressionPrefix("."),
		MustExpressionRegexp(`[0-9]`),
		MustExpressionRegexp("#xB7"),
		MustExpressionRegexp(`[#x0300-#x036F]`),
		MustExpressionRegexp(`[#x203F-#x2040]`),
	))
	gb.AddExpression("Name", NewExpressionAnd(
		NewExpressionRefer("NameStartChar"),
		NewExpressionMultiple(NewExpressionRefer("NameChar")),
	))
	// Char     			::= [http://www.w3.org/TR/REC-xml#NT-Char] /* xgc: xml-version */
	//
	// [http://www.w3.org/TR/REC-xml#NT-Char]
	// Char					::= #x9 | #xA | #xD | [#x20-#xD7FF] | [#xE000-#xFFFD] | [#x10000-#x10FFFF]
	gb.AddExpression("Char", NewExpressionOr(
		MustExpressionRegexp(`#x9`),
		MustExpressionRegexp(`#xA`),
		MustExpressionRegexp(`[#x20-#xD7FF]`),
		MustExpressionRegexp(`[#xE000-#xFFFD]`),
		MustExpressionRegexp(`[#x10000-#x10FFFF]`),
	))
	// Digits   			::= [0-9]+
	gb.AddExpression("Digits", MustExpressionRegexp(`[0-9]+`))
	// CommentContents 		::= (Char+ - (Char* ('(:' | ':)') Char*))
}