package z3

import _ "unsafe"

/**
  \brief Return the nonzero subresultants of \c p and \c q with respect to the "variable" \c x.

  \pre \c p, \c q and \c x are Z3 expressions where \c p and \c q are arithmetic terms.
  Note that, any subterm that cannot be viewed as a polynomial is assumed to be a variable.
  Example: \ccode{f(a)} is a considered to be a variable in the polynomial \ccode{
  f(a)*f(a) + 2*f(a) + 1}

  def_API('Z3_polynomial_subresultants', AST_VECTOR, (_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname PolynomialSubresultants C.Z3_polynomial_subresultants
func PolynomialSubresultants(c Context, p Ast, q Ast, x Ast) AstVector
