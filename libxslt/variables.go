package libxslt

import (
	"github.com/goplus/llgo/c"
	"github.com/luoliwoshang/goplus-llpkg/libxml2"
	"unsafe"
)

/*
 * Interfaces for the variable module.
 */
//go:linkname EvalGlobalVariables C.xsltEvalGlobalVariables
func EvalGlobalVariables(ctxt TransformContextPtr) c.Int

//go:linkname EvalUserParams C.xsltEvalUserParams
func EvalUserParams(ctxt TransformContextPtr, params **int8) c.Int

//go:linkname QuoteUserParams C.xsltQuoteUserParams
func QuoteUserParams(ctxt TransformContextPtr, params **int8) c.Int

//go:linkname EvalOneUserParam C.xsltEvalOneUserParam
func EvalOneUserParam(ctxt TransformContextPtr, name *libxml2.Char, value *libxml2.Char) c.Int

//go:linkname QuoteOneUserParam C.xsltQuoteOneUserParam
func QuoteOneUserParam(ctxt TransformContextPtr, name *libxml2.Char, value *libxml2.Char) c.Int

//go:linkname ParseGlobalVariable C.xsltParseGlobalVariable
func ParseGlobalVariable(style StylesheetPtr, cur libxml2.NodePtr)

//go:linkname ParseGlobalParam C.xsltParseGlobalParam
func ParseGlobalParam(style StylesheetPtr, cur libxml2.NodePtr)

//go:linkname ParseStylesheetVariable C.xsltParseStylesheetVariable
func ParseStylesheetVariable(ctxt TransformContextPtr, cur libxml2.NodePtr)

//go:linkname ParseStylesheetParam C.xsltParseStylesheetParam
func ParseStylesheetParam(ctxt TransformContextPtr, cur libxml2.NodePtr)

//go:linkname ParseStylesheetCallerParam C.xsltParseStylesheetCallerParam
func ParseStylesheetCallerParam(ctxt TransformContextPtr, cur libxml2.NodePtr) StackElemPtr

//go:linkname AddStackElemList C.xsltAddStackElemList
func AddStackElemList(ctxt TransformContextPtr, elems StackElemPtr) c.Int

//go:linkname FreeGlobalVariables C.xsltFreeGlobalVariables
func FreeGlobalVariables(ctxt TransformContextPtr)

//go:linkname VariableLookup C.xsltVariableLookup
func VariableLookup(ctxt TransformContextPtr, name *libxml2.Char, ns_uri *libxml2.Char) libxml2.XPathObjectPtr

//go:linkname XPathVariableLookup C.xsltXPathVariableLookup
func XPathVariableLookup(ctxt unsafe.Pointer, name *libxml2.Char, ns_uri *libxml2.Char) libxml2.XPathObjectPtr
