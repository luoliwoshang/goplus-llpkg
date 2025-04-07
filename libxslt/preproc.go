package libxslt

import (
	"github.com/goplus/llpkg/libxml2"
	_ "unsafe"
)

//go:linkname DocumentComp C.xsltDocumentComp
func DocumentComp(style StylesheetPtr, inst libxml2.NodePtr, function TransformFunction) ElemPreCompPtr

//go:linkname StylePreCompute C.xsltStylePreCompute
func StylePreCompute(style StylesheetPtr, inst libxml2.NodePtr)

//go:linkname FreeStylePreComps C.xsltFreeStylePreComps
func FreeStylePreComps(style StylesheetPtr)
