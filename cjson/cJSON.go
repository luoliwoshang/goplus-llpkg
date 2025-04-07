package cjson

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const VERSION_MAJOR = 1
const VERSION_MINOR = 7
const VERSION_PATCH = 18
const IsReference = 256
const StringIsConst = 512
const NESTING_LIMIT = 1000

/* The cJSON structure: */

type JSON struct {
	Next        *JSON
	Prev        *JSON
	Child       *JSON
	Type        c.Int
	Valuestring *int8
	Valueint    c.Int
	Valuedouble float64
	String      *int8
}

type Hooks struct {
	MallocFn unsafe.Pointer
	FreeFn   unsafe.Pointer
}
type Bool c.Int

/* returns the version of cJSON as a string */
//go:linkname Version C.cJSON_Version
func Version() *int8

/* Supply malloc, realloc and free functions to cJSON */
// llgo:link (*Hooks).InitHooks C.cJSON_InitHooks
func (recv_ *Hooks) InitHooks() {
}

/* Memory Management: the caller is always responsible to free the results from all variants of cJSON_Parse (with cJSON_Delete) and cJSON_Print (with stdlib free, cJSON_Hooks.free_fn, or cJSON_free as appropriate). The exception is cJSON_PrintPreallocated, where the caller has full responsibility of the buffer. */
/* Supply a block of JSON, and this returns a cJSON object you can interrogate. */
//go:linkname Parse C.cJSON_Parse
func Parse(value *int8) *JSON

//go:linkname ParseWithLength C.cJSON_ParseWithLength
func ParseWithLength(value *int8, buffer_length uintptr) *JSON

/* ParseWithOpts allows you to require (and check) that the JSON is null terminated, and to retrieve the pointer to the final byte parsed. */
/* If you supply a ptr in return_parse_end and parsing fails, then return_parse_end will contain a pointer to the error so will match cJSON_GetErrorPtr(). */
//go:linkname ParseWithOpts C.cJSON_ParseWithOpts
func ParseWithOpts(value *int8, return_parse_end **int8, require_null_terminated Bool) *JSON

//go:linkname ParseWithLengthOpts C.cJSON_ParseWithLengthOpts
func ParseWithLengthOpts(value *int8, buffer_length uintptr, return_parse_end **int8, require_null_terminated Bool) *JSON

/* Render a cJSON entity to text for transfer/storage. */
// llgo:link (*JSON).Print C.cJSON_Print
func (recv_ *JSON) Print() *int8 {
	return nil
}

/* Render a cJSON entity to text for transfer/storage without any formatting. */
// llgo:link (*JSON).CStr C.cJSON_PrintUnformatted
func (recv_ *JSON) CStr() *int8 {
	return nil
}

/* Render a cJSON entity to text using a buffered strategy. prebuffer is a guess at the final size. guessing well reduces reallocation. fmt=0 gives unformatted, =1 gives formatted */
// llgo:link (*JSON).PrintBuffered C.cJSON_PrintBuffered
func (recv_ *JSON) PrintBuffered(prebuffer c.Int, fmt Bool) *int8 {
	return nil
}

/* Render a cJSON entity to text using a buffer already allocated in memory with given length. Returns 1 on success and 0 on failure. */
/* NOTE: cJSON is not always 100% accurate in estimating how much memory it will use, so to be safe allocate 5 bytes more than you actually need */
// llgo:link (*JSON).PrintPreallocated C.cJSON_PrintPreallocated
func (recv_ *JSON) PrintPreallocated(buffer *int8, length c.Int, format Bool) Bool {
	return 0
}

/* Delete a cJSON entity and all subentities. */
// llgo:link (*JSON).Delete C.cJSON_Delete
func (recv_ *JSON) Delete() {
}

/* Returns the number of items in an array (or object). */
// llgo:link (*JSON).GetArraySize C.cJSON_GetArraySize
func (recv_ *JSON) GetArraySize() c.Int {
	return 0
}

/* Retrieve item number "index" from array "array". Returns NULL if unsuccessful. */
// llgo:link (*JSON).GetArrayItem C.cJSON_GetArrayItem
func (recv_ *JSON) GetArrayItem(index c.Int) *JSON {
	return nil
}

/* Get item "string" from object. Case insensitive. */
// llgo:link (*JSON).GetObjectItem C.cJSON_GetObjectItem
func (recv_ *JSON) GetObjectItem(string *int8) *JSON {
	return nil
}

// llgo:link (*JSON).GetObjectItemCaseSensitive C.cJSON_GetObjectItemCaseSensitive
func (recv_ *JSON) GetObjectItemCaseSensitive(string *int8) *JSON {
	return nil
}

// llgo:link (*JSON).HasObjectItem C.cJSON_HasObjectItem
func (recv_ *JSON) HasObjectItem(string *int8) Bool {
	return 0
}

/* For analysing failed parses. This returns a pointer to the parse error. You'll probably need to look a few chars back to make sense of it. Defined when cJSON_Parse() returns 0. 0 when cJSON_Parse() succeeds. */
//go:linkname GetErrorPtr C.cJSON_GetErrorPtr
func GetErrorPtr() *int8

/* Check item type and return its value */
// llgo:link (*JSON).GetStringValue C.cJSON_GetStringValue
func (recv_ *JSON) GetStringValue() *int8 {
	return nil
}

// llgo:link (*JSON).GetNumberValue C.cJSON_GetNumberValue
func (recv_ *JSON) GetNumberValue() float64 {
	return 0
}

/* These functions check the type of an item */
// llgo:link (*JSON).IsInvalid C.cJSON_IsInvalid
func (recv_ *JSON) IsInvalid() Bool {
	return 0
}

// llgo:link (*JSON).IsFalse C.cJSON_IsFalse
func (recv_ *JSON) IsFalse() Bool {
	return 0
}

// llgo:link (*JSON).IsTrue C.cJSON_IsTrue
func (recv_ *JSON) IsTrue() Bool {
	return 0
}

// llgo:link (*JSON).IsBool C.cJSON_IsBool
func (recv_ *JSON) IsBool() Bool {
	return 0
}

// llgo:link (*JSON).IsNull C.cJSON_IsNull
func (recv_ *JSON) IsNull() Bool {
	return 0
}

// llgo:link (*JSON).IsNumber C.cJSON_IsNumber
func (recv_ *JSON) IsNumber() Bool {
	return 0
}

// llgo:link (*JSON).IsString C.cJSON_IsString
func (recv_ *JSON) IsString() Bool {
	return 0
}

// llgo:link (*JSON).IsArray C.cJSON_IsArray
func (recv_ *JSON) IsArray() Bool {
	return 0
}

// llgo:link (*JSON).IsObject C.cJSON_IsObject
func (recv_ *JSON) IsObject() Bool {
	return 0
}

// llgo:link (*JSON).IsRaw C.cJSON_IsRaw
func (recv_ *JSON) IsRaw() Bool {
	return 0
}

/* These calls create a cJSON item of the appropriate type. */
//go:linkname CreateNull C.cJSON_CreateNull
func CreateNull() *JSON

//go:linkname CreateTrue C.cJSON_CreateTrue
func CreateTrue() *JSON

//go:linkname CreateFalse C.cJSON_CreateFalse
func CreateFalse() *JSON

// llgo:link Bool.CreateBool C.cJSON_CreateBool
func (recv_ Bool) CreateBool() *JSON {
	return nil
}

//go:linkname CreateNumber C.cJSON_CreateNumber
func CreateNumber(num float64) *JSON

//go:linkname String C.cJSON_CreateString
func String(string *int8) *JSON

/* raw json */
//go:linkname CreateRaw C.cJSON_CreateRaw
func CreateRaw(raw *int8) *JSON

//go:linkname Array C.cJSON_CreateArray
func Array() *JSON

//go:linkname Object C.cJSON_CreateObject
func Object() *JSON

/* Create a string where valuestring references a string so
 * it will not be freed by cJSON_Delete */
//go:linkname CreateStringReference C.cJSON_CreateStringReference
func CreateStringReference(string *int8) *JSON

/* Create an object/array that only references it's elements so
 * they will not be freed by cJSON_Delete */
// llgo:link (*JSON).CreateObjectReference C.cJSON_CreateObjectReference
func (recv_ *JSON) CreateObjectReference() *JSON {
	return nil
}

// llgo:link (*JSON).CreateArrayReference C.cJSON_CreateArrayReference
func (recv_ *JSON) CreateArrayReference() *JSON {
	return nil
}

/* These utilities create an Array of count items.
 * The parameter count cannot be greater than the number of elements in the number array, otherwise array access will be out of bounds.*/
//go:linkname CreateIntArray C.cJSON_CreateIntArray
func CreateIntArray(numbers *c.Int, count c.Int) *JSON

//go:linkname CreateFloatArray C.cJSON_CreateFloatArray
func CreateFloatArray(numbers *float32, count c.Int) *JSON

//go:linkname CreateDoubleArray C.cJSON_CreateDoubleArray
func CreateDoubleArray(numbers *float64, count c.Int) *JSON

//go:linkname CreateStringArray C.cJSON_CreateStringArray
func CreateStringArray(strings **int8, count c.Int) *JSON

/* Append item to the specified array/object. */
// llgo:link (*JSON).AddItem C.cJSON_AddItemToArray
func (recv_ *JSON) AddItem(item *JSON) Bool {
	return 0
}

// llgo:link (*JSON).SetItem C.cJSON_AddItemToObject
func (recv_ *JSON) SetItem(string *int8, item *JSON) Bool {
	return 0
}

/* Use this when string is definitely const (i.e. a literal, or as good as), and will definitely survive the cJSON object.
 * WARNING: When this function was used, make sure to always check that (item->type & cJSON_StringIsConst) is zero before
 * writing to `item->string` */
// llgo:link (*JSON).AddItemToObjectCS C.cJSON_AddItemToObjectCS
func (recv_ *JSON) AddItemToObjectCS(string *int8, item *JSON) Bool {
	return 0
}

/* Append reference to item to the specified array/object. Use this when you want to add an existing cJSON to a new cJSON, but don't want to corrupt your existing cJSON. */
// llgo:link (*JSON).AddItemReferenceToArray C.cJSON_AddItemReferenceToArray
func (recv_ *JSON) AddItemReferenceToArray(item *JSON) Bool {
	return 0
}

// llgo:link (*JSON).AddItemReferenceToObject C.cJSON_AddItemReferenceToObject
func (recv_ *JSON) AddItemReferenceToObject(string *int8, item *JSON) Bool {
	return 0
}

/* Remove/Detach items from Arrays/Objects. */
// llgo:link (*JSON).DetachItemViaPointer C.cJSON_DetachItemViaPointer
func (recv_ *JSON) DetachItemViaPointer(item *JSON) *JSON {
	return nil
}

// llgo:link (*JSON).DetachItemFromArray C.cJSON_DetachItemFromArray
func (recv_ *JSON) DetachItemFromArray(which c.Int) *JSON {
	return nil
}

// llgo:link (*JSON).DeleteItemFromArray C.cJSON_DeleteItemFromArray
func (recv_ *JSON) DeleteItemFromArray(which c.Int) {
}

// llgo:link (*JSON).DetachItemFromObject C.cJSON_DetachItemFromObject
func (recv_ *JSON) DetachItemFromObject(string *int8) *JSON {
	return nil
}

// llgo:link (*JSON).DetachItemFromObjectCaseSensitive C.cJSON_DetachItemFromObjectCaseSensitive
func (recv_ *JSON) DetachItemFromObjectCaseSensitive(string *int8) *JSON {
	return nil
}

// llgo:link (*JSON).DeleteItemFromObject C.cJSON_DeleteItemFromObject
func (recv_ *JSON) DeleteItemFromObject(string *int8) {
}

// llgo:link (*JSON).DeleteItemFromObjectCaseSensitive C.cJSON_DeleteItemFromObjectCaseSensitive
func (recv_ *JSON) DeleteItemFromObjectCaseSensitive(string *int8) {
}

/* Update array items. */
// llgo:link (*JSON).InsertItemInArray C.cJSON_InsertItemInArray
func (recv_ *JSON) InsertItemInArray(which c.Int, newitem *JSON) Bool {
	return 0
}

// llgo:link (*JSON).ReplaceItemViaPointer C.cJSON_ReplaceItemViaPointer
func (recv_ *JSON) ReplaceItemViaPointer(item *JSON, replacement *JSON) Bool {
	return 0
}

// llgo:link (*JSON).ReplaceItemInArray C.cJSON_ReplaceItemInArray
func (recv_ *JSON) ReplaceItemInArray(which c.Int, newitem *JSON) Bool {
	return 0
}

// llgo:link (*JSON).ReplaceItemInObject C.cJSON_ReplaceItemInObject
func (recv_ *JSON) ReplaceItemInObject(string *int8, newitem *JSON) Bool {
	return 0
}

// llgo:link (*JSON).ReplaceItemInObjectCaseSensitive C.cJSON_ReplaceItemInObjectCaseSensitive
func (recv_ *JSON) ReplaceItemInObjectCaseSensitive(string *int8, newitem *JSON) Bool {
	return 0
}

/* Duplicate a cJSON item */
// llgo:link (*JSON).Duplicate C.cJSON_Duplicate
func (recv_ *JSON) Duplicate(recurse Bool) *JSON {
	return nil
}

/* Duplicate will create a new, identical cJSON item to the one you pass, in new memory that will
 * need to be released. With recurse!=0, it will duplicate any children connected to the item.
 * The item->next and ->prev pointers are always zero on return from Duplicate. */
/* Recursively compare two cJSON items for equality. If either a or b is NULL or invalid, they will be considered unequal.
 * case_sensitive determines if object keys are treated case sensitive (1) or case insensitive (0) */
// llgo:link (*JSON).Compare C.cJSON_Compare
func (recv_ *JSON) Compare(b *JSON, case_sensitive Bool) Bool {
	return 0
}

/* Minify a strings, remove blank characters(such as ' ', '\t', '\r', '\n') from strings.
 * The input pointer json cannot point to a read-only address area, such as a string constant,
 * but should point to a readable and writable address area. */
//go:linkname Minify C.cJSON_Minify
func Minify(json *int8)

/* Helper functions for creating and adding items to an object at the same time.
 * They return the added item or NULL on failure. */
// llgo:link (*JSON).AddNullToObject C.cJSON_AddNullToObject
func (recv_ *JSON) AddNullToObject(name *int8) *JSON {
	return nil
}

// llgo:link (*JSON).AddTrueToObject C.cJSON_AddTrueToObject
func (recv_ *JSON) AddTrueToObject(name *int8) *JSON {
	return nil
}

// llgo:link (*JSON).AddFalseToObject C.cJSON_AddFalseToObject
func (recv_ *JSON) AddFalseToObject(name *int8) *JSON {
	return nil
}

// llgo:link (*JSON).AddBoolToObject C.cJSON_AddBoolToObject
func (recv_ *JSON) AddBoolToObject(name *int8, boolean Bool) *JSON {
	return nil
}

// llgo:link (*JSON).AddNumberToObject C.cJSON_AddNumberToObject
func (recv_ *JSON) AddNumberToObject(name *int8, number float64) *JSON {
	return nil
}

// llgo:link (*JSON).AddStringToObject C.cJSON_AddStringToObject
func (recv_ *JSON) AddStringToObject(name *int8, string *int8) *JSON {
	return nil
}

// llgo:link (*JSON).AddRawToObject C.cJSON_AddRawToObject
func (recv_ *JSON) AddRawToObject(name *int8, raw *int8) *JSON {
	return nil
}

// llgo:link (*JSON).AddObjectToObject C.cJSON_AddObjectToObject
func (recv_ *JSON) AddObjectToObject(name *int8) *JSON {
	return nil
}

// llgo:link (*JSON).AddArrayToObject C.cJSON_AddArrayToObject
func (recv_ *JSON) AddArrayToObject(name *int8) *JSON {
	return nil
}

/* helper for the cJSON_SetNumberValue macro */
// llgo:link (*JSON).SetNumberHelper C.cJSON_SetNumberHelper
func (recv_ *JSON) SetNumberHelper(number float64) float64 {
	return 0
}

/* Change the valuestring of a cJSON_String object, only takes effect when type of object is cJSON_String */
// llgo:link (*JSON).SetValuestring C.cJSON_SetValuestring
func (recv_ *JSON) SetValuestring(valuestring *int8) *int8 {
	return nil
}

/* malloc/free objects using the malloc/free functions that have been set with cJSON_InitHooks */
//go:linkname Malloc C.cJSON_malloc
func Malloc(size uintptr) unsafe.Pointer

//go:linkname FreeCStr C.cJSON_free
func FreeCStr(object unsafe.Pointer)
