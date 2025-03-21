package z3

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/** @name Real Closed Fields */
/**@{*/
/**
  \brief Delete a RCF numeral created using the RCF API.

  def_API('Z3_rcf_del', VOID, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfDel C.Z3_rcf_del
func RcfDel(c Context, a RcfNum)

/**
  \brief Return a RCF rational using the given string.

  def_API('Z3_rcf_mk_rational', RCF_NUM, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname RcfMkRational C.Z3_rcf_mk_rational
func RcfMkRational(c Context, val String) RcfNum

/**
  \brief Return a RCF small integer.

  def_API('Z3_rcf_mk_small_int', RCF_NUM, (_in(CONTEXT), _in(INT)))
*/
//go:linkname RcfMkSmallInt C.Z3_rcf_mk_small_int
func RcfMkSmallInt(c Context, val c.Int) RcfNum

/**
  \brief Return Pi

  def_API('Z3_rcf_mk_pi', RCF_NUM, (_in(CONTEXT),))
*/
//go:linkname RcfMkPi C.Z3_rcf_mk_pi
func RcfMkPi(c Context) RcfNum

/**
  \brief Return e (Euler's constant)

  def_API('Z3_rcf_mk_e', RCF_NUM, (_in(CONTEXT),))
*/
//go:linkname RcfMkE C.Z3_rcf_mk_e
func RcfMkE(c Context) RcfNum

/**
  \brief Return a new infinitesimal that is smaller than all elements in the Z3 field.

  def_API('Z3_rcf_mk_infinitesimal', RCF_NUM, (_in(CONTEXT),))
*/
//go:linkname RcfMkInfinitesimal C.Z3_rcf_mk_infinitesimal
func RcfMkInfinitesimal(c Context) RcfNum

/**
  \brief Store in roots the roots of the polynomial \ccode{a[n-1]*x^{n-1} + ... + a[0]}.
  The output vector \c roots must have size \c n.
  It returns the number of roots of the polynomial.

  \pre The input polynomial is not the zero polynomial.

  def_API('Z3_rcf_mk_roots', UINT, (_in(CONTEXT), _in(UINT), _in_array(1, RCF_NUM), _out_array(1, RCF_NUM)))
*/
//go:linkname RcfMkRoots C.Z3_rcf_mk_roots
func RcfMkRoots(c Context, n c.Uint, a *RcfNum, roots *RcfNum) c.Uint

/**
  \brief Return the value \ccode{a + b}.

  def_API('Z3_rcf_add', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfAdd C.Z3_rcf_add
func RcfAdd(c Context, a RcfNum, b RcfNum) RcfNum

/**
  \brief Return the value \ccode{a - b}.

  def_API('Z3_rcf_sub', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfSub C.Z3_rcf_sub
func RcfSub(c Context, a RcfNum, b RcfNum) RcfNum

/**
  \brief Return the value \ccode{a * b}.

  def_API('Z3_rcf_mul', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfMul C.Z3_rcf_mul
func RcfMul(c Context, a RcfNum, b RcfNum) RcfNum

/**
  \brief Return the value \ccode{a / b}.

  def_API('Z3_rcf_div', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfDiv C.Z3_rcf_div
func RcfDiv(c Context, a RcfNum, b RcfNum) RcfNum

/**
  \brief Return the value \ccode{-a}.

  def_API('Z3_rcf_neg', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfNeg C.Z3_rcf_neg
func RcfNeg(c Context, a RcfNum) RcfNum

/**
  \brief Return the value \ccode{1/a}.

  def_API('Z3_rcf_inv', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfInv C.Z3_rcf_inv
func RcfInv(c Context, a RcfNum) RcfNum

/**
  \brief Return the value \ccode{a^k}.

  def_API('Z3_rcf_power', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(UINT)))
*/
//go:linkname RcfPower C.Z3_rcf_power
func RcfPower(c Context, a RcfNum, k c.Uint) RcfNum

/**
  \brief Return \c true if \ccode{a < b}.

  def_API('Z3_rcf_lt', BOOL, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfLt C.Z3_rcf_lt
func RcfLt(c Context, a RcfNum, b RcfNum) bool

/**
  \brief Return \c true if \ccode{a > b}.

  def_API('Z3_rcf_gt', BOOL, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfGt C.Z3_rcf_gt
func RcfGt(c Context, a RcfNum, b RcfNum) bool

/**
  \brief Return \c true if \ccode{a <= b}.

  def_API('Z3_rcf_le', BOOL, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfLe C.Z3_rcf_le
func RcfLe(c Context, a RcfNum, b RcfNum) bool

/**
  \brief Return \c true if \ccode{a >= b}.

  def_API('Z3_rcf_ge', BOOL, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfGe C.Z3_rcf_ge
func RcfGe(c Context, a RcfNum, b RcfNum) bool

/**
  \brief Return \c true if \ccode{a == b}.

  def_API('Z3_rcf_eq', BOOL, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfEq C.Z3_rcf_eq
func RcfEq(c Context, a RcfNum, b RcfNum) bool

/**
  \brief Return \c true if \ccode{a != b}.

  def_API('Z3_rcf_neq', BOOL, (_in(CONTEXT), _in(RCF_NUM), _in(RCF_NUM)))
*/
//go:linkname RcfNeq C.Z3_rcf_neq
func RcfNeq(c Context, a RcfNum, b RcfNum) bool

/**
  \brief Convert the RCF numeral into a string.

  def_API('Z3_rcf_num_to_string', STRING, (_in(CONTEXT), _in(RCF_NUM), _in(BOOL), _in(BOOL)))
*/
//go:linkname RcfNumToString C.Z3_rcf_num_to_string
func RcfNumToString(c Context, a RcfNum, compact bool, html bool) String

/**
  \brief Convert the RCF numeral into a string in decimal notation.

  def_API('Z3_rcf_num_to_decimal_string', STRING, (_in(CONTEXT), _in(RCF_NUM), _in(UINT)))
*/
//go:linkname RcfNumToDecimalString C.Z3_rcf_num_to_decimal_string
func RcfNumToDecimalString(c Context, a RcfNum, prec c.Uint) String

/**
  \brief Extract the "numerator" and "denominator" of the given RCF numeral.
  We have that \ccode{a = n/d}, moreover \c n and \c d are not represented using rational functions.

  def_API('Z3_rcf_get_numerator_denominator', VOID, (_in(CONTEXT), _in(RCF_NUM), _out(RCF_NUM), _out(RCF_NUM)))
*/
//go:linkname RcfGetNumeratorDenominator C.Z3_rcf_get_numerator_denominator
func RcfGetNumeratorDenominator(c Context, a RcfNum, n *RcfNum, d *RcfNum)

/**
  \brief Return \c true if \c a represents a rational number.

  def_API('Z3_rcf_is_rational', BOOL, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfIsRational C.Z3_rcf_is_rational
func RcfIsRational(c Context, a RcfNum) bool

/**
  \brief Return \c true if \c a represents an algebraic number.

  def_API('Z3_rcf_is_algebraic', BOOL, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfIsAlgebraic C.Z3_rcf_is_algebraic
func RcfIsAlgebraic(c Context, a RcfNum) bool

/**
  \brief Return \c true if \c a represents an infinitesimal.

  def_API('Z3_rcf_is_infinitesimal', BOOL, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfIsInfinitesimal C.Z3_rcf_is_infinitesimal
func RcfIsInfinitesimal(c Context, a RcfNum) bool

/**
  \brief Return \c true if \c a represents a transcendental number.

  def_API('Z3_rcf_is_transcendental', BOOL, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfIsTranscendental C.Z3_rcf_is_transcendental
func RcfIsTranscendental(c Context, a RcfNum) bool

/**
   \brief Return the index of a field extension.

  def_API('Z3_rcf_extension_index', UINT, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfExtensionIndex C.Z3_rcf_extension_index
func RcfExtensionIndex(c Context, a RcfNum) c.Uint

/**
  \brief Return the name of a transcendental.

  \pre Z3_rcf_is_transcendtal(ctx, a);

  def_API('Z3_rcf_transcendental_name', SYMBOL, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfTranscendentalName C.Z3_rcf_transcendental_name
func RcfTranscendentalName(c Context, a RcfNum) Symbol

/**
  \brief Return the name of an infinitesimal.

  \pre Z3_rcf_is_infinitesimal(ctx, a);

  def_API('Z3_rcf_infinitesimal_name', SYMBOL, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfInfinitesimalName C.Z3_rcf_infinitesimal_name
func RcfInfinitesimalName(c Context, a RcfNum) Symbol

/**
  \brief Return the number of coefficients in an algebraic number.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_num_coefficients', UINT, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfNumCoefficients C.Z3_rcf_num_coefficients
func RcfNumCoefficients(c Context, a RcfNum) c.Uint

/**
  \brief Extract a coefficient from an algebraic number.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_coefficient', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(UINT)))
*/
//go:linkname RcfCoefficient C.Z3_rcf_coefficient
func RcfCoefficient(c Context, a RcfNum, i c.Uint) RcfNum

/**
  \brief Extract an interval from an algebraic number.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_interval', INT, (_in(CONTEXT), _in(RCF_NUM), _out(INT), _out(INT), _out(RCF_NUM), _out(INT), _out(INT), _out(RCF_NUM)))
*/
//go:linkname RcfInterval C.Z3_rcf_interval
func RcfInterval(c Context, a RcfNum, lower_is_inf *c.Int, lower_is_open *c.Int, lower *RcfNum, upper_is_inf *c.Int, upper_is_open *c.Int, upper *RcfNum) c.Int

/**
  \brief Return the number of sign conditions of an algebraic number.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_num_sign_conditions', UINT, (_in(CONTEXT), _in(RCF_NUM)))
*/
//go:linkname RcfNumSignConditions C.Z3_rcf_num_sign_conditions
func RcfNumSignConditions(c Context, a RcfNum) c.Uint

/**
  \brief Extract the sign of a sign condition from an algebraic number.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_sign_condition_sign', INT, (_in(CONTEXT), _in(RCF_NUM), _in(UINT)))
*/
//go:linkname RcfSignConditionSign C.Z3_rcf_sign_condition_sign
func RcfSignConditionSign(c Context, a RcfNum, i c.Uint) c.Int

/**
  \brief Return the number of sign condition polynomial coefficients of an algebraic number.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_num_sign_condition_coefficients', UINT, (_in(CONTEXT), _in(RCF_NUM), _in(UINT)))
*/
//go:linkname RcfNumSignConditionCoefficients C.Z3_rcf_num_sign_condition_coefficients
func RcfNumSignConditionCoefficients(c Context, a RcfNum, i c.Uint) c.Uint

/**
  \brief Extract the j-th polynomial coefficient of the i-th sign condition.

  \pre Z3_rcf_is_algebraic(ctx, a);

  def_API('Z3_rcf_sign_condition_coefficient', RCF_NUM, (_in(CONTEXT), _in(RCF_NUM), _in(UINT), _in(UINT)))
*/
//go:linkname RcfSignConditionCoefficient C.Z3_rcf_sign_condition_coefficient
func RcfSignConditionCoefficient(c Context, a RcfNum, i c.Uint, j c.Uint) RcfNum
