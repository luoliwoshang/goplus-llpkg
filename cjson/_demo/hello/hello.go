package main

import (
	"unsafe"

	"github.com/luoliwoshang/goplus-llpkg/cjson"

	"github.com/goplus/llgo/c"
)

func main() {
	mod := cjson.Object()
	mod.SetItem(c.Str("name"), cjson.String(c.Str("math")))

	syms := cjson.Array()

	fn := cjson.Object()
	fn.SetItem(c.Str("name"), cjson.String(c.Str("sqrt")))
	fn.SetItem(c.Str("sig"), cjson.String(c.Str("(x, /)")))
	syms.AddItem(fn)

	v := cjson.Object()
	v.SetItem(c.Str("name"), cjson.String(c.Str("pi")))
	syms.AddItem(v)

	mod.SetItem(c.Str("items"), syms)

	cstr := mod.CStr()
	str := c.GoString(cstr)
	c.Printf(c.Str("%s\n"), cstr)
	cjson.FreeCStr(unsafe.Pointer(cstr))

	mod.Delete()

	cjsonLoad(str)
}

func cjsonLoad(str string) {
	mod := ParseString(str)

	cstr := mod.Print()
	c.Printf(c.Str("%s\n"), cstr)
	cjson.FreeCStr(unsafe.Pointer(cstr))

	mod.Delete()
}

func ParseString(value string) *cjson.JSON {
	return cjson.ParseWithLength((*c.Char)(unsafe.Pointer(unsafe.StringData(value))), uintptr(len(value)))
}
