package z3

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/** @name AST vectors */
/**@{*/
/**
  \brief Return an empty AST vector.

  \remark Reference counting must be used to manage AST vectors, even when the Z3_context was
  created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_mk_ast_vector', AST_VECTOR, (_in(CONTEXT),))
*/
//go:linkname MkAstVector C.Z3_mk_ast_vector
func MkAstVector(c Context) AstVector

/**
  \brief Increment the reference counter of the given AST vector.

  def_API('Z3_ast_vector_inc_ref', VOID, (_in(CONTEXT), _in(AST_VECTOR)))
*/
//go:linkname AstVectorIncRef C.Z3_ast_vector_inc_ref
func AstVectorIncRef(c Context, v AstVector)

/**
  \brief Decrement the reference counter of the given AST vector.

  def_API('Z3_ast_vector_dec_ref', VOID, (_in(CONTEXT), _in(AST_VECTOR)))
*/
//go:linkname AstVectorDecRef C.Z3_ast_vector_dec_ref
func AstVectorDecRef(c Context, v AstVector)

/**
  \brief Return the size of the given AST vector.

  def_API('Z3_ast_vector_size', UINT, (_in(CONTEXT), _in(AST_VECTOR)))
*/
//go:linkname AstVectorSize C.Z3_ast_vector_size
func AstVectorSize(c Context, v AstVector) c.Uint

/**
  \brief Return the AST at position \c i in the AST vector \c v.

  \pre i < Z3_ast_vector_size(c, v)

  def_API('Z3_ast_vector_get', AST, (_in(CONTEXT), _in(AST_VECTOR), _in(UINT)))
*/
//go:linkname AstVectorGet C.Z3_ast_vector_get
func AstVectorGet(c Context, v AstVector, i c.Uint) Ast

/**
  \brief Update position \c i of the AST vector \c v with the AST \c a.

  \pre i < Z3_ast_vector_size(c, v)

  def_API('Z3_ast_vector_set', VOID, (_in(CONTEXT), _in(AST_VECTOR), _in(UINT), _in(AST)))
*/
//go:linkname AstVectorSet C.Z3_ast_vector_set
func AstVectorSet(c Context, v AstVector, i c.Uint, a Ast)

/**
  \brief Resize the AST vector \c v.

  def_API('Z3_ast_vector_resize', VOID, (_in(CONTEXT), _in(AST_VECTOR), _in(UINT)))
*/
//go:linkname AstVectorResize C.Z3_ast_vector_resize
func AstVectorResize(c Context, v AstVector, n c.Uint)

/**
  \brief Add the AST \c a in the end of the AST vector \c v. The size of \c v is increased by one.

  def_API('Z3_ast_vector_push', VOID, (_in(CONTEXT), _in(AST_VECTOR), _in(AST)))
*/
//go:linkname AstVectorPush C.Z3_ast_vector_push
func AstVectorPush(c Context, v AstVector, a Ast)

/**
  \brief Translate the AST vector \c v from context \c s into an AST vector in context \c t.

  def_API('Z3_ast_vector_translate', AST_VECTOR, (_in(CONTEXT), _in(AST_VECTOR), _in(CONTEXT)))
*/
//go:linkname AstVectorTranslate C.Z3_ast_vector_translate
func AstVectorTranslate(s Context, v AstVector, t Context) AstVector

/**
  \brief Convert AST vector into a string.

  def_API('Z3_ast_vector_to_string', STRING, (_in(CONTEXT), _in(AST_VECTOR)))
*/
//go:linkname AstVectorToString C.Z3_ast_vector_to_string
func AstVectorToString(c Context, v AstVector) String

/** @name AST maps */
/**@{*/
/**
  \brief Return an empty mapping from AST to AST

  \remark Reference counting must be used to manage AST maps, even when the Z3_context was
  created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_mk_ast_map', AST_MAP, (_in(CONTEXT),) )
*/
//go:linkname MkAstMap C.Z3_mk_ast_map
func MkAstMap(c Context) AstMap

/**
  \brief Increment the reference counter of the given AST map.

  def_API('Z3_ast_map_inc_ref', VOID, (_in(CONTEXT), _in(AST_MAP)))
*/
//go:linkname AstMapIncRef C.Z3_ast_map_inc_ref
func AstMapIncRef(c Context, m AstMap)

/**
  \brief Decrement the reference counter of the given AST map.

  def_API('Z3_ast_map_dec_ref', VOID, (_in(CONTEXT), _in(AST_MAP)))
*/
//go:linkname AstMapDecRef C.Z3_ast_map_dec_ref
func AstMapDecRef(c Context, m AstMap)

/**
  \brief Return true if the map \c m contains the AST key \c k.

  def_API('Z3_ast_map_contains', BOOL, (_in(CONTEXT), _in(AST_MAP), _in(AST)))
*/
//go:linkname AstMapContains C.Z3_ast_map_contains
func AstMapContains(c Context, m AstMap, k Ast) bool

/**
  \brief Return the value associated with the key \c k.

  The procedure invokes the error handler if \c k is not in the map.

  def_API('Z3_ast_map_find', AST, (_in(CONTEXT), _in(AST_MAP), _in(AST)))
*/
//go:linkname AstMapFind C.Z3_ast_map_find
func AstMapFind(c Context, m AstMap, k Ast) Ast

/**
  \brief Store/Replace a new key, value pair in the given map.

  def_API('Z3_ast_map_insert', VOID, (_in(CONTEXT), _in(AST_MAP), _in(AST), _in(AST)))
*/
//go:linkname AstMapInsert C.Z3_ast_map_insert
func AstMapInsert(c Context, m AstMap, k Ast, v Ast)

/**
  \brief Erase a key from the map.

  def_API('Z3_ast_map_erase', VOID, (_in(CONTEXT), _in(AST_MAP), _in(AST)))
*/
//go:linkname AstMapErase C.Z3_ast_map_erase
func AstMapErase(c Context, m AstMap, k Ast)

/**
  \brief Remove all keys from the given map.

  def_API('Z3_ast_map_reset', VOID, (_in(CONTEXT), _in(AST_MAP)))
*/
//go:linkname AstMapReset C.Z3_ast_map_reset
func AstMapReset(c Context, m AstMap)

/**
  \brief Return the size of the given map.

  def_API('Z3_ast_map_size', UINT, (_in(CONTEXT), _in(AST_MAP)))
*/
//go:linkname AstMapSize C.Z3_ast_map_size
func AstMapSize(c Context, m AstMap) c.Uint

/**
  \brief Return the keys stored in the given map.

  def_API('Z3_ast_map_keys', AST_VECTOR, (_in(CONTEXT), _in(AST_MAP)))
*/
//go:linkname AstMapKeys C.Z3_ast_map_keys
func AstMapKeys(c Context, m AstMap) AstVector

/**
  \brief Convert the given map into a string.

  def_API('Z3_ast_map_to_string', STRING, (_in(CONTEXT), _in(AST_MAP)))
*/
//go:linkname AstMapToString C.Z3_ast_map_to_string
func AstMapToString(c Context, m AstMap) String
