package quickjs

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const JS_PROP_HAS_SHIFT = 8
const JS_ATOM_NULL = 0
const JS_DEF_CFUNC = 0
const JS_DEF_CGETSET = 1
const JS_DEF_CGETSET_MAGIC = 2
const JS_DEF_PROP_STRING = 3
const JS_DEF_PROP_INT32 = 4
const JS_DEF_PROP_INT64 = 5
const JS_DEF_PROP_DOUBLE = 6
const JS_DEF_PROP_UNDEFINED = 7
const JS_DEF_OBJECT = 8
const JS_DEF_ALIAS = 9

type JSRuntime struct {
	Unused [8]uint8
}

type JSContext struct {
	Unused [8]uint8
}

type JSObject struct {
	Unused [8]uint8
}

type JSClass struct {
	Unused [8]uint8
}
type JSClassID uint32
type JSAtom uint32

const (
	JSTAGFIRST            c.Int = -11
	JSTAGBIGDECIMAL       c.Int = -11
	JSTAGBIGINT           c.Int = -10
	JSTAGBIGFLOAT         c.Int = -9
	JSTAGSYMBOL           c.Int = -8
	JSTAGSTRING           c.Int = -7
	JSTAGMODULE           c.Int = -3
	JSTAGFUNCTIONBYTECODE c.Int = -2
	JSTAGOBJECT           c.Int = -1
	JSTAGINT              c.Int = 0
	JSTAGBOOL             c.Int = 1
	JSTAGNULL             c.Int = 2
	JSTAGUNDEFINED        c.Int = 3
	JSTAGUNINITIALIZED    c.Int = 4
	JSTAGCATCHOFFSET      c.Int = 5
	JSTAGEXCEPTION        c.Int = 6
	JSTAGFLOAT64          c.Int = 7
)

type JSRefCountHeader struct {
	RefCount c.Int
}

type JSValueUnion struct {
	Float64 float64
}

type JSValue struct {
	U   JSValueUnion
	Tag int64
}

// llgo:type C
type JSCFunction func(*JSContext, JSValue, c.Int, *JSValue) JSValue

// llgo:type C
type JSCFunctionMagic func(*JSContext, JSValue, c.Int, *JSValue, c.Int) JSValue

// llgo:type C
type JSCFunctionData func(*JSContext, JSValue, c.Int, *JSValue, c.Int, *JSValue) JSValue

type JSMallocState struct {
	MallocCount uintptr
	MallocSize  uintptr
	MallocLimit uintptr
	Opaque      unsafe.Pointer
}

type JSMallocFunctions struct {
	JsMalloc           unsafe.Pointer
	JsFree             unsafe.Pointer
	JsRealloc          unsafe.Pointer
	JsMallocUsableSize unsafe.Pointer
}

type JSGCObjectHeader struct {
	Unused [8]uint8
}

//go:linkname JSNewRuntime C.JS_NewRuntime
func JSNewRuntime() *JSRuntime

/* info lifetime must exceed that of rt */
// llgo:link (*JSRuntime).JSSetRuntimeInfo C.JS_SetRuntimeInfo
func (recv_ *JSRuntime) JSSetRuntimeInfo(info *int8) {
}

// llgo:link (*JSRuntime).JSSetMemoryLimit C.JS_SetMemoryLimit
func (recv_ *JSRuntime) JSSetMemoryLimit(limit uintptr) {
}

// llgo:link (*JSRuntime).JSSetGCThreshold C.JS_SetGCThreshold
func (recv_ *JSRuntime) JSSetGCThreshold(gc_threshold uintptr) {
}

/* use 0 to disable maximum stack size check */
// llgo:link (*JSRuntime).JSSetMaxStackSize C.JS_SetMaxStackSize
func (recv_ *JSRuntime) JSSetMaxStackSize(stack_size uintptr) {
}

/* should be called when changing thread to update the stack top value
   used to check stack overflow. */
// llgo:link (*JSRuntime).JSUpdateStackTop C.JS_UpdateStackTop
func (recv_ *JSRuntime) JSUpdateStackTop() {
}

// llgo:link (*JSMallocFunctions).JSNewRuntime2 C.JS_NewRuntime2
func (recv_ *JSMallocFunctions) JSNewRuntime2(opaque unsafe.Pointer) *JSRuntime {
	return nil
}

// llgo:link (*JSRuntime).JSFreeRuntime C.JS_FreeRuntime
func (recv_ *JSRuntime) JSFreeRuntime() {
}

// llgo:link (*JSRuntime).JSGetRuntimeOpaque C.JS_GetRuntimeOpaque
func (recv_ *JSRuntime) JSGetRuntimeOpaque() unsafe.Pointer {
	return nil
}

// llgo:link (*JSRuntime).JSSetRuntimeOpaque C.JS_SetRuntimeOpaque
func (recv_ *JSRuntime) JSSetRuntimeOpaque(opaque unsafe.Pointer) {
}

// llgo:type C
type JSMarkFunc func(*JSRuntime, *JSGCObjectHeader)

// llgo:link (*JSRuntime).JSMarkValue C.JS_MarkValue
func (recv_ *JSRuntime) JSMarkValue(val JSValue, mark_func JSMarkFunc) {
}

// llgo:link (*JSRuntime).JSRunGC C.JS_RunGC
func (recv_ *JSRuntime) JSRunGC() {
}

// llgo:link (*JSRuntime).JSIsLiveObject C.JS_IsLiveObject
func (recv_ *JSRuntime) JSIsLiveObject(obj JSValue) c.Int {
	return 0
}

// llgo:link (*JSRuntime).JSNewContext C.JS_NewContext
func (recv_ *JSRuntime) JSNewContext() *JSContext {
	return nil
}

// llgo:link (*JSContext).JSFreeContext C.JS_FreeContext
func (recv_ *JSContext) JSFreeContext() {
}

// llgo:link (*JSContext).JSDupContext C.JS_DupContext
func (recv_ *JSContext) JSDupContext() *JSContext {
	return nil
}

// llgo:link (*JSContext).JSGetContextOpaque C.JS_GetContextOpaque
func (recv_ *JSContext) JSGetContextOpaque() unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JSSetContextOpaque C.JS_SetContextOpaque
func (recv_ *JSContext) JSSetContextOpaque(opaque unsafe.Pointer) {
}

// llgo:link (*JSContext).JSGetRuntime C.JS_GetRuntime
func (recv_ *JSContext) JSGetRuntime() *JSRuntime {
	return nil
}

// llgo:link (*JSContext).JSSetClassProto C.JS_SetClassProto
func (recv_ *JSContext) JSSetClassProto(class_id JSClassID, obj JSValue) {
}

// llgo:link (*JSContext).JSGetClassProto C.JS_GetClassProto
func (recv_ *JSContext) JSGetClassProto(class_id JSClassID) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

/* the following functions are used to select the intrinsic object to
   save memory */
// llgo:link (*JSRuntime).JSNewContextRaw C.JS_NewContextRaw
func (recv_ *JSRuntime) JSNewContextRaw() *JSContext {
	return nil
}

// llgo:link (*JSContext).JSAddIntrinsicBaseObjects C.JS_AddIntrinsicBaseObjects
func (recv_ *JSContext) JSAddIntrinsicBaseObjects() {
}

// llgo:link (*JSContext).JSAddIntrinsicDate C.JS_AddIntrinsicDate
func (recv_ *JSContext) JSAddIntrinsicDate() {
}

// llgo:link (*JSContext).JSAddIntrinsicEval C.JS_AddIntrinsicEval
func (recv_ *JSContext) JSAddIntrinsicEval() {
}

// llgo:link (*JSContext).JSAddIntrinsicStringNormalize C.JS_AddIntrinsicStringNormalize
func (recv_ *JSContext) JSAddIntrinsicStringNormalize() {
}

// llgo:link (*JSContext).JSAddIntrinsicRegExpCompiler C.JS_AddIntrinsicRegExpCompiler
func (recv_ *JSContext) JSAddIntrinsicRegExpCompiler() {
}

// llgo:link (*JSContext).JSAddIntrinsicRegExp C.JS_AddIntrinsicRegExp
func (recv_ *JSContext) JSAddIntrinsicRegExp() {
}

// llgo:link (*JSContext).JSAddIntrinsicJSON C.JS_AddIntrinsicJSON
func (recv_ *JSContext) JSAddIntrinsicJSON() {
}

// llgo:link (*JSContext).JSAddIntrinsicProxy C.JS_AddIntrinsicProxy
func (recv_ *JSContext) JSAddIntrinsicProxy() {
}

// llgo:link (*JSContext).JSAddIntrinsicMapSet C.JS_AddIntrinsicMapSet
func (recv_ *JSContext) JSAddIntrinsicMapSet() {
}

// llgo:link (*JSContext).JSAddIntrinsicTypedArrays C.JS_AddIntrinsicTypedArrays
func (recv_ *JSContext) JSAddIntrinsicTypedArrays() {
}

// llgo:link (*JSContext).JSAddIntrinsicPromise C.JS_AddIntrinsicPromise
func (recv_ *JSContext) JSAddIntrinsicPromise() {
}

// llgo:link (*JSContext).JSAddIntrinsicBigInt C.JS_AddIntrinsicBigInt
func (recv_ *JSContext) JSAddIntrinsicBigInt() {
}

// llgo:link (*JSContext).JSAddIntrinsicBigFloat C.JS_AddIntrinsicBigFloat
func (recv_ *JSContext) JSAddIntrinsicBigFloat() {
}

// llgo:link (*JSContext).JSAddIntrinsicBigDecimal C.JS_AddIntrinsicBigDecimal
func (recv_ *JSContext) JSAddIntrinsicBigDecimal() {
}

/* enable operator overloading */
// llgo:link (*JSContext).JSAddIntrinsicOperators C.JS_AddIntrinsicOperators
func (recv_ *JSContext) JSAddIntrinsicOperators() {
}

/* enable "use math" */
// llgo:link (*JSContext).JSEnableBignumExt C.JS_EnableBignumExt
func (recv_ *JSContext) JSEnableBignumExt(enable c.Int) {
}

// llgo:link (*JSContext).JsStringCodePointRange C.js_string_codePointRange
func (recv_ *JSContext) JsStringCodePointRange(this_val JSValue, argc c.Int, argv *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSRuntime).JsMallocRt C.js_malloc_rt
func (recv_ *JSRuntime) JsMallocRt(size uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSRuntime).JsFreeRt C.js_free_rt
func (recv_ *JSRuntime) JsFreeRt(ptr unsafe.Pointer) {
}

// llgo:link (*JSRuntime).JsReallocRt C.js_realloc_rt
func (recv_ *JSRuntime) JsReallocRt(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSRuntime).JsMallocUsableSizeRt C.js_malloc_usable_size_rt
func (recv_ *JSRuntime) JsMallocUsableSizeRt(ptr unsafe.Pointer) uintptr {
	return 0
}

// llgo:link (*JSRuntime).JsMalloczRt C.js_mallocz_rt
func (recv_ *JSRuntime) JsMalloczRt(size uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JsMalloc C.js_malloc
func (recv_ *JSContext) JsMalloc(size uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JsFree C.js_free
func (recv_ *JSContext) JsFree(ptr unsafe.Pointer) {
}

// llgo:link (*JSContext).JsRealloc C.js_realloc
func (recv_ *JSContext) JsRealloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JsMallocUsableSize C.js_malloc_usable_size
func (recv_ *JSContext) JsMallocUsableSize(ptr unsafe.Pointer) uintptr {
	return 0
}

// llgo:link (*JSContext).JsRealloc2 C.js_realloc2
func (recv_ *JSContext) JsRealloc2(ptr unsafe.Pointer, size uintptr, pslack *uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JsMallocz C.js_mallocz
func (recv_ *JSContext) JsMallocz(size uintptr) unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JsStrdup C.js_strdup
func (recv_ *JSContext) JsStrdup(str *int8) *int8 {
	return nil
}

// llgo:link (*JSContext).JsStrndup C.js_strndup
func (recv_ *JSContext) JsStrndup(s *int8, n uintptr) *int8 {
	return nil
}

type JSMemoryUsage struct {
	MallocSize         int64
	MallocLimit        int64
	MemoryUsedSize     int64
	MallocCount        int64
	MemoryUsedCount    int64
	AtomCount          int64
	AtomSize           int64
	StrCount           int64
	StrSize            int64
	ObjCount           int64
	ObjSize            int64
	PropCount          int64
	PropSize           int64
	ShapeCount         int64
	ShapeSize          int64
	JsFuncCount        int64
	JsFuncSize         int64
	JsFuncCodeSize     int64
	JsFuncPc2lineCount int64
	JsFuncPc2lineSize  int64
	CFuncCount         int64
	ArrayCount         int64
	FastArrayCount     int64
	FastArrayElements  int64
	BinaryObjectCount  int64
	BinaryObjectSize   int64
}

// llgo:link (*JSRuntime).JSComputeMemoryUsage C.JS_ComputeMemoryUsage
func (recv_ *JSRuntime) JSComputeMemoryUsage(s *JSMemoryUsage) {
}

//go:linkname JSDumpMemoryUsage C.JS_DumpMemoryUsage
func JSDumpMemoryUsage(fp *c.FILE, s *JSMemoryUsage, rt *JSRuntime)

// llgo:link (*JSContext).JSNewAtomLen C.JS_NewAtomLen
func (recv_ *JSContext) JSNewAtomLen(str *int8, len uintptr) JSAtom {
	return 0
}

// llgo:link (*JSContext).JSNewAtom C.JS_NewAtom
func (recv_ *JSContext) JSNewAtom(str *int8) JSAtom {
	return 0
}

// llgo:link (*JSContext).JSNewAtomUInt32 C.JS_NewAtomUInt32
func (recv_ *JSContext) JSNewAtomUInt32(n uint32) JSAtom {
	return 0
}

// llgo:link (*JSContext).JSDupAtom C.JS_DupAtom
func (recv_ *JSContext) JSDupAtom(v JSAtom) JSAtom {
	return 0
}

// llgo:link (*JSContext).JSFreeAtom C.JS_FreeAtom
func (recv_ *JSContext) JSFreeAtom(v JSAtom) {
}

// llgo:link (*JSRuntime).JSFreeAtomRT C.JS_FreeAtomRT
func (recv_ *JSRuntime) JSFreeAtomRT(v JSAtom) {
}

// llgo:link (*JSContext).JSAtomToValue C.JS_AtomToValue
func (recv_ *JSContext) JSAtomToValue(atom JSAtom) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSAtomToString C.JS_AtomToString
func (recv_ *JSContext) JSAtomToString(atom JSAtom) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSAtomToCString C.JS_AtomToCString
func (recv_ *JSContext) JSAtomToCString(atom JSAtom) *int8 {
	return nil
}

// llgo:link (*JSContext).JSValueToAtom C.JS_ValueToAtom
func (recv_ *JSContext) JSValueToAtom(val JSValue) JSAtom {
	return 0
}

/* object class support */

type JSPropertyEnum struct {
	IsEnumerable c.Int
	Atom         JSAtom
}

type JSPropertyDescriptor struct {
	Flags  c.Int
	Value  JSValue
	Getter JSValue
	Setter JSValue
}

type JSClassExoticMethods struct {
	GetOwnProperty      unsafe.Pointer
	GetOwnPropertyNames unsafe.Pointer
	DeleteProperty      unsafe.Pointer
	DefineOwnProperty   unsafe.Pointer
	HasProperty         unsafe.Pointer
	GetProperty         unsafe.Pointer
	SetProperty         unsafe.Pointer
}

// llgo:type C
type JSClassFinalizer func(*JSRuntime, JSValue)

// llgo:type C
type JSClassGCMark func(*JSRuntime, JSValue, JSMarkFunc)

// llgo:type C
type JSClassCall func(*JSContext, JSValue, JSValue, c.Int, *JSValue, c.Int) JSValue

type JSClassDef struct {
	ClassName *int8
	Finalizer *unsafe.Pointer
	GcMark    *unsafe.Pointer
	Call      *unsafe.Pointer
	Exotic    *JSClassExoticMethods
}

// llgo:link (*JSClassID).JSNewClassID C.JS_NewClassID
func (recv_ *JSClassID) JSNewClassID() JSClassID {
	return 0
}

// llgo:link (*JSRuntime).JSNewClass C.JS_NewClass
func (recv_ *JSRuntime) JSNewClass(class_id JSClassID, class_def *JSClassDef) c.Int {
	return 0
}

// llgo:link (*JSRuntime).JSIsRegisteredClass C.JS_IsRegisteredClass
func (recv_ *JSRuntime) JSIsRegisteredClass(class_id JSClassID) c.Int {
	return 0
}

// llgo:link (*JSContext).JSNewBigInt64 C.JS_NewBigInt64
func (recv_ *JSContext) JSNewBigInt64(v int64) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewBigUint64 C.JS_NewBigUint64
func (recv_ *JSContext) JSNewBigUint64(v uint64) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrow C.JS_Throw
func (recv_ *JSContext) JSThrow(obj JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSGetException C.JS_GetException
func (recv_ *JSContext) JSGetException() JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSIsError C.JS_IsError
func (recv_ *JSContext) JSIsError(val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSResetUncatchableError C.JS_ResetUncatchableError
func (recv_ *JSContext) JSResetUncatchableError() {
}

// llgo:link (*JSContext).JSNewError C.JS_NewError
func (recv_ *JSContext) JSNewError() JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrowSyntaxError C.JS_ThrowSyntaxError
func (recv_ *JSContext) JSThrowSyntaxError(fmt *int8, __llgo_va_list ...interface{}) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrowTypeError C.JS_ThrowTypeError
func (recv_ *JSContext) JSThrowTypeError(fmt *int8, __llgo_va_list ...interface{}) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrowReferenceError C.JS_ThrowReferenceError
func (recv_ *JSContext) JSThrowReferenceError(fmt *int8, __llgo_va_list ...interface{}) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrowRangeError C.JS_ThrowRangeError
func (recv_ *JSContext) JSThrowRangeError(fmt *int8, __llgo_va_list ...interface{}) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrowInternalError C.JS_ThrowInternalError
func (recv_ *JSContext) JSThrowInternalError(fmt *int8, __llgo_va_list ...interface{}) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSThrowOutOfMemory C.JS_ThrowOutOfMemory
func (recv_ *JSContext) JSThrowOutOfMemory() JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).X__JSFreeValue C.__JS_FreeValue
func (recv_ *JSContext) X__JSFreeValue(v JSValue) {
}

// llgo:link (*JSRuntime).X__JSFreeValueRT C.__JS_FreeValueRT
func (recv_ *JSRuntime) X__JSFreeValueRT(v JSValue) {
}

// llgo:link (*JSContext).JSToBool C.JS_ToBool
func (recv_ *JSContext) JSToBool(val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSToInt32 C.JS_ToInt32
func (recv_ *JSContext) JSToInt32(pres *int32, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSToInt64 C.JS_ToInt64
func (recv_ *JSContext) JSToInt64(pres *int64, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSToIndex C.JS_ToIndex
func (recv_ *JSContext) JSToIndex(plen *uint64, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSToFloat64 C.JS_ToFloat64
func (recv_ *JSContext) JSToFloat64(pres *float64, val JSValue) c.Int {
	return 0
}

/* return an exception if 'val' is a Number */
// llgo:link (*JSContext).JSToBigInt64 C.JS_ToBigInt64
func (recv_ *JSContext) JSToBigInt64(pres *int64, val JSValue) c.Int {
	return 0
}

/* same as JS_ToInt64() but allow BigInt */
// llgo:link (*JSContext).JSToInt64Ext C.JS_ToInt64Ext
func (recv_ *JSContext) JSToInt64Ext(pres *int64, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSNewStringLen C.JS_NewStringLen
func (recv_ *JSContext) JSNewStringLen(str1 *int8, len1 uintptr) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewString C.JS_NewString
func (recv_ *JSContext) JSNewString(str *int8) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewAtomString C.JS_NewAtomString
func (recv_ *JSContext) JSNewAtomString(str *int8) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSToString C.JS_ToString
func (recv_ *JSContext) JSToString(val JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSToPropertyKey C.JS_ToPropertyKey
func (recv_ *JSContext) JSToPropertyKey(val JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSToCStringLen2 C.JS_ToCStringLen2
func (recv_ *JSContext) JSToCStringLen2(plen *uintptr, val1 JSValue, cesu8 c.Int) *int8 {
	return nil
}

// llgo:link (*JSContext).JSFreeCString C.JS_FreeCString
func (recv_ *JSContext) JSFreeCString(ptr *int8) {
}

// llgo:link (*JSContext).JSNewObjectProtoClass C.JS_NewObjectProtoClass
func (recv_ *JSContext) JSNewObjectProtoClass(proto JSValue, class_id JSClassID) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewObjectClass C.JS_NewObjectClass
func (recv_ *JSContext) JSNewObjectClass(class_id c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewObjectProto C.JS_NewObjectProto
func (recv_ *JSContext) JSNewObjectProto(proto JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewObject C.JS_NewObject
func (recv_ *JSContext) JSNewObject() JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSIsFunction C.JS_IsFunction
func (recv_ *JSContext) JSIsFunction(val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSIsConstructor C.JS_IsConstructor
func (recv_ *JSContext) JSIsConstructor(val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSSetConstructorBit C.JS_SetConstructorBit
func (recv_ *JSContext) JSSetConstructorBit(func_obj JSValue, val c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSNewArray C.JS_NewArray
func (recv_ *JSContext) JSNewArray() JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSIsArray C.JS_IsArray
func (recv_ *JSContext) JSIsArray(val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSGetPropertyInternal C.JS_GetPropertyInternal
func (recv_ *JSContext) JSGetPropertyInternal(obj JSValue, prop JSAtom, receiver JSValue, throw_ref_error c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSGetPropertyStr C.JS_GetPropertyStr
func (recv_ *JSContext) JSGetPropertyStr(this_obj JSValue, prop *int8) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSGetPropertyUint32 C.JS_GetPropertyUint32
func (recv_ *JSContext) JSGetPropertyUint32(this_obj JSValue, idx uint32) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSSetPropertyInternal C.JS_SetPropertyInternal
func (recv_ *JSContext) JSSetPropertyInternal(obj JSValue, prop JSAtom, val JSValue, this_obj JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSSetPropertyUint32 C.JS_SetPropertyUint32
func (recv_ *JSContext) JSSetPropertyUint32(this_obj JSValue, idx uint32, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSSetPropertyInt64 C.JS_SetPropertyInt64
func (recv_ *JSContext) JSSetPropertyInt64(this_obj JSValue, idx int64, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSSetPropertyStr C.JS_SetPropertyStr
func (recv_ *JSContext) JSSetPropertyStr(this_obj JSValue, prop *int8, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSHasProperty C.JS_HasProperty
func (recv_ *JSContext) JSHasProperty(this_obj JSValue, prop JSAtom) c.Int {
	return 0
}

// llgo:link (*JSContext).JSIsExtensible C.JS_IsExtensible
func (recv_ *JSContext) JSIsExtensible(obj JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSPreventExtensions C.JS_PreventExtensions
func (recv_ *JSContext) JSPreventExtensions(obj JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSDeleteProperty C.JS_DeleteProperty
func (recv_ *JSContext) JSDeleteProperty(obj JSValue, prop JSAtom, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSSetPrototype C.JS_SetPrototype
func (recv_ *JSContext) JSSetPrototype(obj JSValue, proto_val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSGetPrototype C.JS_GetPrototype
func (recv_ *JSContext) JSGetPrototype(val JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSGetOwnPropertyNames C.JS_GetOwnPropertyNames
func (recv_ *JSContext) JSGetOwnPropertyNames(ptab **JSPropertyEnum, plen *uint32, obj JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSGetOwnProperty C.JS_GetOwnProperty
func (recv_ *JSContext) JSGetOwnProperty(desc *JSPropertyDescriptor, obj JSValue, prop JSAtom) c.Int {
	return 0
}

// llgo:link (*JSContext).JSCall C.JS_Call
func (recv_ *JSContext) JSCall(func_obj JSValue, this_obj JSValue, argc c.Int, argv *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSInvoke C.JS_Invoke
func (recv_ *JSContext) JSInvoke(this_val JSValue, atom JSAtom, argc c.Int, argv *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSCallConstructor C.JS_CallConstructor
func (recv_ *JSContext) JSCallConstructor(func_obj JSValue, argc c.Int, argv *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSCallConstructor2 C.JS_CallConstructor2
func (recv_ *JSContext) JSCallConstructor2(func_obj JSValue, new_target JSValue, argc c.Int, argv *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

//go:linkname JSDetectModule C.JS_DetectModule
func JSDetectModule(input *int8, input_len uintptr) c.Int

/* 'input' must be zero terminated i.e. input[input_len] = '\0'. */
// llgo:link (*JSContext).JSEval C.JS_Eval
func (recv_ *JSContext) JSEval(input *int8, input_len uintptr, filename *int8, eval_flags c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

/* same as JS_Eval() but with an explicit 'this_obj' parameter */
// llgo:link (*JSContext).JSEvalThis C.JS_EvalThis
func (recv_ *JSContext) JSEvalThis(this_obj JSValue, input *int8, input_len uintptr, filename *int8, eval_flags c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSGetGlobalObject C.JS_GetGlobalObject
func (recv_ *JSContext) JSGetGlobalObject() JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSIsInstanceOf C.JS_IsInstanceOf
func (recv_ *JSContext) JSIsInstanceOf(val JSValue, obj JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSDefineProperty C.JS_DefineProperty
func (recv_ *JSContext) JSDefineProperty(this_obj JSValue, prop JSAtom, val JSValue, getter JSValue, setter JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSDefinePropertyValue C.JS_DefinePropertyValue
func (recv_ *JSContext) JSDefinePropertyValue(this_obj JSValue, prop JSAtom, val JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSDefinePropertyValueUint32 C.JS_DefinePropertyValueUint32
func (recv_ *JSContext) JSDefinePropertyValueUint32(this_obj JSValue, idx uint32, val JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSDefinePropertyValueStr C.JS_DefinePropertyValueStr
func (recv_ *JSContext) JSDefinePropertyValueStr(this_obj JSValue, prop *int8, val JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link (*JSContext).JSDefinePropertyGetSet C.JS_DefinePropertyGetSet
func (recv_ *JSContext) JSDefinePropertyGetSet(this_obj JSValue, prop JSAtom, getter JSValue, setter JSValue, flags c.Int) c.Int {
	return 0
}

// llgo:link JSValue.JSSetOpaque C.JS_SetOpaque
func (recv_ JSValue) JSSetOpaque(opaque unsafe.Pointer) {
}

// llgo:link JSValue.JSGetOpaque C.JS_GetOpaque
func (recv_ JSValue) JSGetOpaque(class_id JSClassID) unsafe.Pointer {
	return nil
}

// llgo:link (*JSContext).JSGetOpaque2 C.JS_GetOpaque2
func (recv_ *JSContext) JSGetOpaque2(obj JSValue, class_id JSClassID) unsafe.Pointer {
	return nil
}

/* 'buf' must be zero terminated i.e. buf[buf_len] = '\0'. */
// llgo:link (*JSContext).JSParseJSON C.JS_ParseJSON
func (recv_ *JSContext) JSParseJSON(buf *int8, buf_len uintptr, filename *int8) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSParseJSON2 C.JS_ParseJSON2
func (recv_ *JSContext) JSParseJSON2(buf *int8, buf_len uintptr, filename *int8, flags c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSJSONStringify C.JS_JSONStringify
func (recv_ *JSContext) JSJSONStringify(obj JSValue, replacer JSValue, space0 JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:type C
type JSFreeArrayBufferDataFunc func(*JSRuntime, unsafe.Pointer, unsafe.Pointer)

// llgo:link (*JSContext).JSNewArrayBuffer C.JS_NewArrayBuffer
func (recv_ *JSContext) JSNewArrayBuffer(buf *uint8, len uintptr, free_func JSFreeArrayBufferDataFunc, opaque unsafe.Pointer, is_shared c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewArrayBufferCopy C.JS_NewArrayBufferCopy
func (recv_ *JSContext) JSNewArrayBufferCopy(buf *uint8, len uintptr) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSDetachArrayBuffer C.JS_DetachArrayBuffer
func (recv_ *JSContext) JSDetachArrayBuffer(obj JSValue) {
}

// llgo:link (*JSContext).JSGetArrayBuffer C.JS_GetArrayBuffer
func (recv_ *JSContext) JSGetArrayBuffer(psize *uintptr, obj JSValue) *uint8 {
	return nil
}

// llgo:link (*JSContext).JSGetTypedArrayBuffer C.JS_GetTypedArrayBuffer
func (recv_ *JSContext) JSGetTypedArrayBuffer(obj JSValue, pbyte_offset *uintptr, pbyte_length *uintptr, pbytes_per_element *uintptr) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

type JSSharedArrayBufferFunctions struct {
	SabAlloc  unsafe.Pointer
	SabFree   unsafe.Pointer
	SabDup    unsafe.Pointer
	SabOpaque unsafe.Pointer
}

// llgo:link (*JSRuntime).JSSetSharedArrayBufferFunctions C.JS_SetSharedArrayBufferFunctions
func (recv_ *JSRuntime) JSSetSharedArrayBufferFunctions(sf *JSSharedArrayBufferFunctions) {
}

type JSPromiseStateEnum c.Int

const (
	JSPROMISEPENDING   JSPromiseStateEnum = 0
	JSPROMISEFULFILLED JSPromiseStateEnum = 1
	JSPROMISEREJECTED  JSPromiseStateEnum = 2
)

// llgo:link (*JSContext).JSNewPromiseCapability C.JS_NewPromiseCapability
func (recv_ *JSContext) JSNewPromiseCapability(resolving_funcs *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSPromiseState C.JS_PromiseState
func (recv_ *JSContext) JSPromiseState(promise JSValue) JSPromiseStateEnum {
	return 0
}

// llgo:link (*JSContext).JSPromiseResult C.JS_PromiseResult
func (recv_ *JSContext) JSPromiseResult(promise JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:type C
type JSHostPromiseRejectionTracker func(*JSContext, JSValue, JSValue, c.Int, unsafe.Pointer)

// llgo:link (*JSRuntime).JSSetHostPromiseRejectionTracker C.JS_SetHostPromiseRejectionTracker
func (recv_ *JSRuntime) JSSetHostPromiseRejectionTracker(cb JSHostPromiseRejectionTracker, opaque unsafe.Pointer) {
}

// llgo:type C
type JSInterruptHandler func(*JSRuntime, unsafe.Pointer) c.Int

// llgo:link (*JSRuntime).JSSetInterruptHandler C.JS_SetInterruptHandler
func (recv_ *JSRuntime) JSSetInterruptHandler(cb JSInterruptHandler, opaque unsafe.Pointer) {
}

/* if can_block is TRUE, Atomics.wait() can be used */
// llgo:link (*JSRuntime).JSSetCanBlock C.JS_SetCanBlock
func (recv_ *JSRuntime) JSSetCanBlock(can_block c.Int) {
}

/* set the [IsHTMLDDA] internal slot */
// llgo:link (*JSContext).JSSetIsHTMLDDA C.JS_SetIsHTMLDDA
func (recv_ *JSContext) JSSetIsHTMLDDA(obj JSValue) {
}

type JSModuleDef struct {
	Unused [8]uint8
}

// llgo:type C
type JSModuleNormalizeFunc func(*JSContext, *int8, *int8, unsafe.Pointer) *int8

// llgo:type C
type JSModuleLoaderFunc func(*JSContext, *int8, unsafe.Pointer) *JSModuleDef

/* module_normalize = NULL is allowed and invokes the default module
   filename normalizer */
// llgo:link (*JSRuntime).JSSetModuleLoaderFunc C.JS_SetModuleLoaderFunc
func (recv_ *JSRuntime) JSSetModuleLoaderFunc(module_normalize JSModuleNormalizeFunc, module_loader JSModuleLoaderFunc, opaque unsafe.Pointer) {
}

/* return the import.meta object of a module */
// llgo:link (*JSContext).JSGetImportMeta C.JS_GetImportMeta
func (recv_ *JSContext) JSGetImportMeta(m *JSModuleDef) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSGetModuleName C.JS_GetModuleName
func (recv_ *JSContext) JSGetModuleName(m *JSModuleDef) JSAtom {
	return 0
}

// llgo:type C
type JSJobFunc func(*JSContext, c.Int, *JSValue) JSValue

// llgo:link (*JSContext).JSEnqueueJob C.JS_EnqueueJob
func (recv_ *JSContext) JSEnqueueJob(job_func JSJobFunc, argc c.Int, argv *JSValue) c.Int {
	return 0
}

// llgo:link (*JSRuntime).JSIsJobPending C.JS_IsJobPending
func (recv_ *JSRuntime) JSIsJobPending() c.Int {
	return 0
}

// llgo:link (*JSRuntime).JSExecutePendingJob C.JS_ExecutePendingJob
func (recv_ *JSRuntime) JSExecutePendingJob(pctx **JSContext) c.Int {
	return 0
}

// llgo:link (*JSContext).JSWriteObject C.JS_WriteObject
func (recv_ *JSContext) JSWriteObject(psize *uintptr, obj JSValue, flags c.Int) *uint8 {
	return nil
}

// llgo:link (*JSContext).JSWriteObject2 C.JS_WriteObject2
func (recv_ *JSContext) JSWriteObject2(psize *uintptr, obj JSValue, flags c.Int, psab_tab ***uint8, psab_tab_len *uintptr) *uint8 {
	return nil
}

// llgo:link (*JSContext).JSReadObject C.JS_ReadObject
func (recv_ *JSContext) JSReadObject(buf *uint8, buf_len uintptr, flags c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

/* instantiate and evaluate a bytecode function. Only used when
   reading a script or module with JS_ReadObject() */
// llgo:link (*JSContext).JSEvalFunction C.JS_EvalFunction
func (recv_ *JSContext) JSEvalFunction(fun_obj JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

/* load the dependencies of the module 'obj'. Useful when JS_ReadObject()
   returns a module. */
// llgo:link (*JSContext).JSResolveModule C.JS_ResolveModule
func (recv_ *JSContext) JSResolveModule(obj JSValue) c.Int {
	return 0
}

/* only exported for os.Worker() */
// llgo:link (*JSContext).JSGetScriptOrModuleName C.JS_GetScriptOrModuleName
func (recv_ *JSContext) JSGetScriptOrModuleName(n_stack_levels c.Int) JSAtom {
	return 0
}

/* only exported for os.Worker() */
// llgo:link (*JSContext).JSLoadModule C.JS_LoadModule
func (recv_ *JSContext) JSLoadModule(basename *int8, filename *int8) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

type JSCFunctionEnum c.Int

const (
	JSCFUNCGeneric                JSCFunctionEnum = 0
	JSCFUNCGenericMagic           JSCFunctionEnum = 1
	JSCFUNCConstructor            JSCFunctionEnum = 2
	JSCFUNCConstructorMagic       JSCFunctionEnum = 3
	JSCFUNCConstructorOrFunc      JSCFunctionEnum = 4
	JSCFUNCConstructorOrFuncMagic JSCFunctionEnum = 5
	JSCFUNCFF                     JSCFunctionEnum = 6
	JSCFUNCFFF                    JSCFunctionEnum = 7
	JSCFUNCGetter                 JSCFunctionEnum = 8
	JSCFUNCSetter                 JSCFunctionEnum = 9
	JSCFUNCGetterMagic            JSCFunctionEnum = 10
	JSCFUNCSetterMagic            JSCFunctionEnum = 11
	JSCFUNCIteratorNext           JSCFunctionEnum = 12
)

type JSCFunctionType struct {
	Generic *unsafe.Pointer
}

// llgo:link (*JSContext).JSNewCFunction2 C.JS_NewCFunction2
func (recv_ *JSContext) JSNewCFunction2(func_ JSCFunction, name *int8, length c.Int, cproto JSCFunctionEnum, magic c.Int) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSNewCFunctionData C.JS_NewCFunctionData
func (recv_ *JSContext) JSNewCFunctionData(func_ JSCFunctionData, length c.Int, magic c.Int, data_len c.Int, data *JSValue) JSValue {
	return struct {
		U   JSValueUnion
		Tag int64
	}{}
}

// llgo:link (*JSContext).JSSetConstructor C.JS_SetConstructor
func (recv_ *JSContext) JSSetConstructor(func_obj JSValue, proto JSValue) {
}

/* C property definition */

type JSCFunctionListEntry struct {
	Name      *int8
	PropFlags uint8
	DefType   uint8
	Magic     int16
	U         struct {
		Func struct {
			Length uint8
			Cproto uint8
			Cfunc  JSCFunctionType
		}
	}
}

// llgo:link (*JSContext).JSSetPropertyFunctionList C.JS_SetPropertyFunctionList
func (recv_ *JSContext) JSSetPropertyFunctionList(obj JSValue, tab *JSCFunctionListEntry, len c.Int) {
}

// llgo:type C
type JSModuleInitFunc func(*JSContext, *JSModuleDef) c.Int

// llgo:link (*JSContext).JSNewCModule C.JS_NewCModule
func (recv_ *JSContext) JSNewCModule(name_str *int8, func_ JSModuleInitFunc) *JSModuleDef {
	return nil
}

/* can only be called before the module is instantiated */
// llgo:link (*JSContext).JSAddModuleExport C.JS_AddModuleExport
func (recv_ *JSContext) JSAddModuleExport(m *JSModuleDef, name_str *int8) c.Int {
	return 0
}

// llgo:link (*JSContext).JSAddModuleExportList C.JS_AddModuleExportList
func (recv_ *JSContext) JSAddModuleExportList(m *JSModuleDef, tab *JSCFunctionListEntry, len c.Int) c.Int {
	return 0
}

/* can only be called after the module is instantiated */
// llgo:link (*JSContext).JSSetModuleExport C.JS_SetModuleExport
func (recv_ *JSContext) JSSetModuleExport(m *JSModuleDef, export_name *int8, val JSValue) c.Int {
	return 0
}

// llgo:link (*JSContext).JSSetModuleExportList C.JS_SetModuleExportList
func (recv_ *JSContext) JSSetModuleExportList(m *JSModuleDef, tab *JSCFunctionListEntry, len c.Int) c.Int {
	return 0
}
