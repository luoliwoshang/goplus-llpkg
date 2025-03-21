package z3

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/** @name Algebraic Numbers */
/**@{*/
/**
  \brief Return \c true if \c a can be used as value in the Z3 real algebraic
  number package.

  def_API('Z3_algebraic_is_value', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicIsValue C.Z3_algebraic_is_value
func AlgebraicIsValue(c Context, a Ast) bool

/**
  \brief Return \c true if \c a is positive, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)

  def_API('Z3_algebraic_is_pos', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicIsPos C.Z3_algebraic_is_pos
func AlgebraicIsPos(c Context, a Ast) bool

/**
  \brief Return \c true if \c a is negative, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)

  def_API('Z3_algebraic_is_neg', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicIsNeg C.Z3_algebraic_is_neg
func AlgebraicIsNeg(c Context, a Ast) bool

/**
  \brief Return \c true if \c a is zero, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)

  def_API('Z3_algebraic_is_zero', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicIsZero C.Z3_algebraic_is_zero
func AlgebraicIsZero(c Context, a Ast) bool

/**
  \brief Return 1 if \c a is positive, 0 if \c a is zero, and -1 if \c a is negative.

  \pre Z3_algebraic_is_value(c, a)

  def_API('Z3_algebraic_sign', INT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicSign C.Z3_algebraic_sign
func AlgebraicSign(c Context, a Ast) c.Int

/**
  \brief Return the value a + b.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)
  \post Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_add', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicAdd C.Z3_algebraic_add
func AlgebraicAdd(c Context, a Ast, b Ast) Ast

/**
  \brief Return the value a - b.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)
  \post Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_sub', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicSub C.Z3_algebraic_sub
func AlgebraicSub(c Context, a Ast, b Ast) Ast

/**
  \brief Return the value a * b.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)
  \post Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_mul', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicMul C.Z3_algebraic_mul
func AlgebraicMul(c Context, a Ast, b Ast) Ast

/**
  \brief Return the value a / b.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)
  \pre !Z3_algebraic_is_zero(c, b)
  \post Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_div', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicDiv C.Z3_algebraic_div
func AlgebraicDiv(c Context, a Ast, b Ast) Ast

/**
  \brief Return the a^(1/k)

  \pre Z3_algebraic_is_value(c, a)
  \pre k is even => !Z3_algebraic_is_neg(c, a)
  \post Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_root', AST, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname AlgebraicRoot C.Z3_algebraic_root
func AlgebraicRoot(c Context, a Ast, k c.Uint) Ast

/**
  \brief Return the a^k

  \pre Z3_algebraic_is_value(c, a)
  \post Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_power', AST, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname AlgebraicPower C.Z3_algebraic_power
func AlgebraicPower(c Context, a Ast, k c.Uint) Ast

/**
  \brief Return \c true if a < b, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)

  def_API('Z3_algebraic_lt', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicLt C.Z3_algebraic_lt
func AlgebraicLt(c Context, a Ast, b Ast) bool

/**
  \brief Return \c true if a > b, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)

  def_API('Z3_algebraic_gt', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicGt C.Z3_algebraic_gt
func AlgebraicGt(c Context, a Ast, b Ast) bool

/**
  \brief Return \c true if a <= b, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)

  def_API('Z3_algebraic_le', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicLe C.Z3_algebraic_le
func AlgebraicLe(c Context, a Ast, b Ast) bool

/**
  \brief Return \c true if a >= b, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)

  def_API('Z3_algebraic_ge', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicGe C.Z3_algebraic_ge
func AlgebraicGe(c Context, a Ast, b Ast) bool

/**
  \brief Return \c true if a == b, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)

  def_API('Z3_algebraic_eq', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicEq C.Z3_algebraic_eq
func AlgebraicEq(c Context, a Ast, b Ast) bool

/**
  \brief Return \c true if a != b, and \c false otherwise.

  \pre Z3_algebraic_is_value(c, a)
  \pre Z3_algebraic_is_value(c, b)

  def_API('Z3_algebraic_neq', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname AlgebraicNeq C.Z3_algebraic_neq
func AlgebraicNeq(c Context, a Ast, b Ast) bool

/**
  \brief Given a multivariate polynomial p(x_0, ..., x_{n-1}, x_n), returns the
  roots of the univariate polynomial p(a[0], ..., a[n-1], x_n).

  \pre p is a Z3 expression that contains only arithmetic terms and free variables.
  \pre forall i in [0, n) Z3_algebraic_is_value(c, a[i])
  \post forall r in result Z3_algebraic_is_value(c, result)

  def_API('Z3_algebraic_roots', AST_VECTOR, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST)))
*/
//go:linkname AlgebraicRoots C.Z3_algebraic_roots
func AlgebraicRoots(c Context, p Ast, n c.Uint, a *Ast) AstVector

/**
  \brief Given a multivariate polynomial p(x_0, ..., x_{n-1}), return the
  sign of p(a[0], ..., a[n-1]).

  \pre p is a Z3 expression that contains only arithmetic terms and free variables.
  \pre forall i in [0, n) Z3_algebraic_is_value(c, a[i])

  def_API('Z3_algebraic_eval', INT, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST)))
*/
//go:linkname AlgebraicEval C.Z3_algebraic_eval
func AlgebraicEval(c Context, p Ast, n c.Uint, a *Ast) c.Int

/**
  \brief Return the coefficients of the defining polynomial.

  \pre Z3_algebraic_is_value(c, a)

  def_API('Z3_algebraic_get_poly', AST_VECTOR, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicGetPoly C.Z3_algebraic_get_poly
func AlgebraicGetPoly(c Context, a Ast) AstVector

/**
  \brief Return which root of the polynomial the algebraic number represents.

  \pre Z3_algebraic_is_value(c, a)

  def_API('Z3_algebraic_get_i', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AlgebraicGetI C.Z3_algebraic_get_i
func AlgebraicGetI(c Context, a Ast) c.Uint
