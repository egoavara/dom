// Follow by [https://github.com/XPath-Next/XPath-Next/blob/master/grammars/xpath-3.1.ebnf]
//
package compiler

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
	"errors"
	"bytes"
)

var ErrorUnmatching = errors.New("No matching syntax")

type Token interface{
	Take(rs io.ReadSeeker) error
}
type (
	BPath struct {Data IPath}
	// PathExpr ::= ("/" RelativePathExpr?) | ("//" RelativePathExpr) | RelativePathExpr /* xgc: leading-lone-slash */
	IPath interface {

	}
	// "/"
	IPath0 struct {}
	// "/" RelativePathExpr
	IPath1 struct {
		BRelativePath
	}
	// "//" RelativePathExpr
	IPath2 struct {
		BRelativePath
	}
	// RelativePathExpr
	IPath3 struct {
		BRelativePath
	}
)

func (s *BPath)Take(rs io.ReadSeeker) (error) {
	var buf [2]byte
	n, err := rs.Read(buf[:])
	if err != io.EOF && err != nil {
		return err
	}
	if bytes.HasPrefix(buf[:], []byte{'/', '/'}){
		// IPath2
		s.Data = IPath2{}
		return s.Data.(Token).Take(rs)
	}
	if bytes.HasPrefix(buf[:], []byte{'/'}){
		if n == 1{
			// IPath0
			s.Data = IPath0{}
			return nil
		}
		// IPath1
		rs.Seek(-1, io.SeekCurrent)
		s.Data = IPath1{}
		return s.Data.(Token).Take(rs)
	}
	// IPath3
	s.Data = IPath3{}
	return s.Data.(Token).Take(rs)
}

type (
	BRelativePath struct {
		Data IRelativePath
	}
	// RelativePathExpr ::= StepExpr (("/" | "//") StepExpr)*
	IRelativePath struct {

		Data0   IStep
		Spliter []string
		Data    []IStep
	}
)

func (s *BRelativePath) Take(rs io.ReadSeeker) error {
}

type (
	BStep struct {
		Data IStep
	}
	// PostfixExpr | AxisStep
	IStep interface{}
)

func (s *BStep) Take(rs io.ReadSeeker) error {
	return nil
}

type (
	IPostfix struct {
		Data0 IPrimary
		Data1 []IIPostfix0
	}
	// Predicate | ArgumentList | Lookup
	IIPostfix0 interface{}
)

type (
	IPrimary struct {
		Data IIPrimary0
	}
	// Literal | VarRef | ParenthesizedExpr | ContextItemExpr | FunctionCall | FunctionItemExpr | MapConstructor | ArrayConstructor | UnaryLookup
	IIPrimary0 interface{}
)

type (
	// NumericLiteral | StringLiteral
	ILiteral interface{}
	// ('"' (EscapeQuot | [^"])* '"') | ("'" (EscapeApos | [^'])* "'")
	// EscapeQuot ::= '""'
	// EscapeApos ::= "''"
	StringLiteral string

	//
	// IntegerLiteral | DecimalLiteral | DoubleLiteral
	// Digits   		::= [0-9]+
	// IntegerLiteral 	::= Digits
	// DecimalLiteral 	::= ("." Digits) | (Digits "." [0-9]*)
	// DoubleLiteral  	::= (("." Digits) | (Digits ("." [0-9]*)?)) [eE] [+-]? Digits /* ws: explicit */
	//
	NumericLiteral float64
)

type (
	// "$" VarName
	// VarName  ::= EQName
	IVarRef IEQName
)

type (
	// TODO : current, save as text
	//  "(" Expr? ")"
	IParenthesizedExpr struct {
		Data string
	}
)
type (
	// "."
	IContextItemExpr struct{}
)

type (
	// TODO : current, save as text
	// EQName ArgumentList /* xgc: reserved-function-names */ /* gn: parens */
	IFunctionCall struct {
	}
)

type (
	// TODO : current, save as text
	// EQName ArgumentList /* xgc: reserved-function-names */ /* gn: parens */
	// NamedFunctionRef | InlineFunctionExpr
	// NamedFunctionRef ::= EQName "#" IntegerLiteral
	// InlineFunctionExpr ::= "function" "(" ParamList? ")" ("as" SequenceType)? FunctionBody
	IFunctionItemExpr struct {
	}
)

type (
	// TODO : current, save as text
	// MapConstructor ::= "map" "{" (MapConstructorEntry ("," MapConstructorEntry)*)? "}"
	// MapConstructorEntry ::= MapKeyExpr ":" MapValueExpr
	// MapKeyExpr ::= ExprSingle
	// MapValueExpr ::= ExprSingle
	IMapConstructor struct {
	}
)

type (
	// TODO :
	// ArrayConstructor ::= SquareArrayConstructor | CurlyArrayConstructor
	// SquareArrayConstructor ::= "[" (ExprSingle ("," ExprSingle)*)? "]"
	// CurlyArrayConstructor ::= "array" EnclosedExpr
	IArrayConstructor struct {
	}
)


type (
	// TODO :
	// UnaryLookup ::= "?" KeySpecifier
	IUnaryLookup struct {
	}
)

type (
	// (ReverseStep | ForwardStep) PredicateList
	IAxisStep struct {
		Prefix IIAxisStep
		Surfix PredicateList
	}
	// ReverseStep | ForwardStep
	IIAxisStep interface {
	}
)
type (
	// (ReverseAxis NodeTest) | AbbrevReverseStep
	IReverseStep interface {}
	IReverseStep0 struct {
		Prefix IReverseAxis
		Surfix INodeTest
	}
	// ReverseAxis ::= ("parent" "::") | ("ancestor" "::") | ("preceding-sibling" "::") | ("preceding" "::") | ("ancestor-or-self" "::")
	IReverseAxis struct {

	}
	// NodeTest ::= KindTest | NameTest
	INodeTest struct {

	}
	// AbbrevReverseStep ::= ".."
	IAbbrevReverseStep struct{}
)

type (
	// KindTest ::= DocumentTest | ElementTest | AttributeTest | SchemaElementTest | SchemaAttributeTest | PITest | CommentTest | TextTest | NamespaceNodeTest | AnyKindTest
	IKindTest struct {

	}
	// NameTest ::= EQName | Wildcard
	INameTest interface {}
)

type (
	// Predicate*
	PredicateList []Predicate
	// Predicate ::= "[" Expr "]"
	Predicate struct {}
)

type (
	// KeySpecifier ::= NCName | IntegerLiteral | ParenthesizedExpr | "*"
)
type (
	// QName | URIQualifiedName
	IEQName interface{}

	// [http://www.w3.org/TR/REC-xml-names/#NT-QName]
	IQName xml.Name

	// URIQualifiedName ::= BracedURILiteral NCName
	IURIQualifiedName struct {
		URI  IBracedURILiteral
		Name string
	}
	// BracedURILiteral ::= "Q" "{" [^{}]* "}"
	IBracedURILiteral url.URL

	// NCName   ::= [http://www.w3.org/TR/REC-xml-names/#NT-NCName] /* xgc: xml-version */
	INCName string

	// Wildcard ::= "*" | (NCName ":*") | ("*:" NCName) | (BracedURILiteral "*")  /* ws: explicit */
	IWildcard interface {}
	// "*"
	IIWildcard0 struct {}
	// (NCName ":*")
	IIWildcard1 INCName
	// ("*:" NCName)
	IIWildcard2 INCName
	// BracedURILiteral "*"
	IIWildcard3 IBracedURILiteral
)

func Tokenize(src io.Reader) []Token {
	bts, err := ioutil.ReadAll(src)
	if err != io.EOF && err != nil {
		return nil
	}
	srcs := strings.TrimSpace(string(bts))
	switch srcs[:2] {
	case "fo": // for
	case "le": // let
	case "so": // some
	case "ev": // every
	case "if": // if
	default:

	}
	return res
}
