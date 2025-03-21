package z3

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/** @name Floating-Point Arithmetic */
/**@{*/
/**
  \brief Create the RoundingMode sort.

  \param c logical context

  \sa Z3_mk_fpa_round_nearest_ties_to_away or Z3_mk_fpa_rna
  \sa Z3_mk_fpa_round_nearest_ties_to_even or Z3_mk_fpa_rne
  \sa Z3_mk_fpa_round_toward_negative or Z3_mk_fpa_rtn
  \sa Z3_mk_fpa_round_toward_positive or Z3_mk_fpa_rtp
  \sa Z3_mk_fpa_round_toward_zero or Z3_mk_fpa_rtz

  def_API('Z3_mk_fpa_rounding_mode_sort', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaRoundingModeSort C.Z3_mk_fpa_rounding_mode_sort
func MkFpaRoundingModeSort(c Context) Sort

/**
  \brief Create a numeral of RoundingMode sort which represents the NearestTiesToEven rounding mode.

  This is the same as #Z3_mk_fpa_rne.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_round_nearest_ties_to_away
  \sa Z3_mk_fpa_round_toward_negative
  \sa Z3_mk_fpa_round_toward_positive
  \sa Z3_mk_fpa_round_toward_zero

  def_API('Z3_mk_fpa_round_nearest_ties_to_even', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRoundNearestTiesToEven C.Z3_mk_fpa_round_nearest_ties_to_even
func MkFpaRoundNearestTiesToEven(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the NearestTiesToEven rounding mode.

  This is the same as #Z3_mk_fpa_round_nearest_ties_to_even.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_rna
  \sa Z3_mk_fpa_rtn
  \sa Z3_mk_fpa_rtp
  \sa Z3_mk_fpa_rtz

  def_API('Z3_mk_fpa_rne', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRne C.Z3_mk_fpa_rne
func MkFpaRne(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the NearestTiesToAway rounding mode.

  This is the same as #Z3_mk_fpa_rna.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_round_nearest_ties_to_even
  \sa Z3_mk_fpa_round_toward_negative
  \sa Z3_mk_fpa_round_toward_positive
  \sa Z3_mk_fpa_round_toward_zero

  def_API('Z3_mk_fpa_round_nearest_ties_to_away', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRoundNearestTiesToAway C.Z3_mk_fpa_round_nearest_ties_to_away
func MkFpaRoundNearestTiesToAway(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the NearestTiesToAway rounding mode.

  This is the same as #Z3_mk_fpa_round_nearest_ties_to_away.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_rne
  \sa Z3_mk_fpa_rtn
  \sa Z3_mk_fpa_rtp
  \sa Z3_mk_fpa_rtz

  def_API('Z3_mk_fpa_rna', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRna C.Z3_mk_fpa_rna
func MkFpaRna(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the TowardPositive rounding mode.

  This is the same as #Z3_mk_fpa_rtp.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_round_nearest_ties_to_away
  \sa Z3_mk_fpa_round_nearest_ties_to_even
  \sa Z3_mk_fpa_round_toward_negative
  \sa Z3_mk_fpa_round_toward_zero

  def_API('Z3_mk_fpa_round_toward_positive', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRoundTowardPositive C.Z3_mk_fpa_round_toward_positive
func MkFpaRoundTowardPositive(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the TowardPositive rounding mode.

  This is the same as #Z3_mk_fpa_round_toward_positive.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_rna
  \sa Z3_mk_fpa_rne
  \sa Z3_mk_fpa_rtn
  \sa Z3_mk_fpa_rtz

  def_API('Z3_mk_fpa_rtp', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRtp C.Z3_mk_fpa_rtp
func MkFpaRtp(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the TowardNegative rounding mode.

  This is the same as #Z3_mk_fpa_rtn.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_round_nearest_ties_to_away
  \sa Z3_mk_fpa_round_nearest_ties_to_even
  \sa Z3_mk_fpa_round_toward_positive
  \sa Z3_mk_fpa_round_toward_zero

  def_API('Z3_mk_fpa_round_toward_negative', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRoundTowardNegative C.Z3_mk_fpa_round_toward_negative
func MkFpaRoundTowardNegative(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the TowardNegative rounding mode.

  This is the same as #Z3_mk_fpa_round_toward_negative.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_rna
  \sa Z3_mk_fpa_rne
  \sa Z3_mk_fpa_rtp
  \sa Z3_mk_fpa_rtz

  def_API('Z3_mk_fpa_rtn', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRtn C.Z3_mk_fpa_rtn
func MkFpaRtn(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the TowardZero rounding mode.

  This is the same as #Z3_mk_fpa_rtz.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_round_nearest_ties_to_away
  \sa Z3_mk_fpa_round_nearest_ties_to_even
  \sa Z3_mk_fpa_round_toward_negative
  \sa Z3_mk_fpa_round_toward_positive

  def_API('Z3_mk_fpa_round_toward_zero', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRoundTowardZero C.Z3_mk_fpa_round_toward_zero
func MkFpaRoundTowardZero(c Context) Ast

/**
  \brief Create a numeral of RoundingMode sort which represents the TowardZero rounding mode.

  This is the same as #Z3_mk_fpa_round_toward_zero.

  \param c logical context

  \sa Z3_mk_fpa_rounding_mode_sort
  \sa Z3_mk_fpa_rna
  \sa Z3_mk_fpa_rne
  \sa Z3_mk_fpa_rtn
  \sa Z3_mk_fpa_rtp

  def_API('Z3_mk_fpa_rtz', AST, (_in(CONTEXT),))
*/
//go:linkname MkFpaRtz C.Z3_mk_fpa_rtz
func MkFpaRtz(c Context) Ast

/**
  \brief Create a FloatingPoint sort.

  \param c logical context
  \param ebits number of exponent bits
  \param sbits number of significand bits

  \remark \c ebits must be larger than 1 and \c sbits must be larger than 2.

  \sa Z3_mk_fpa_sort_half or Z3_mk_fpa_sort_16
  \sa Z3_mk_fpa_sort_single or Z3_mk_fpa_sort_32
  \sa Z3_mk_fpa_sort_double or Z3_mk_fpa_sort_64
  \sa Z3_mk_fpa_sort_quadruple or Z3_mk_fpa_sort_128

  def_API('Z3_mk_fpa_sort', SORT, (_in(CONTEXT), _in(UINT), _in(UINT)))
*/
//go:linkname MkFpaSort C.Z3_mk_fpa_sort
func MkFpaSort(c Context, ebits c.Uint, sbits c.Uint) Sort

/**
  \brief Create the half-precision (16-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_16.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_single
  \sa Z3_mk_fpa_sort_double
  \sa Z3_mk_fpa_sort_quadruple

  def_API('Z3_mk_fpa_sort_half', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSortHalf C.Z3_mk_fpa_sort_half
func MkFpaSortHalf(c Context) Sort

/**
  \brief Create the half-precision (16-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_half.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_32
  \sa Z3_mk_fpa_sort_64
  \sa Z3_mk_fpa_sort_128

  def_API('Z3_mk_fpa_sort_16', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSort16 C.Z3_mk_fpa_sort_16
func MkFpaSort16(c Context) Sort

/**
  \brief Create the single-precision (32-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_32.

  \param c logical context.

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_half
  \sa Z3_mk_fpa_sort_double
  \sa Z3_mk_fpa_sort_quadruple

  def_API('Z3_mk_fpa_sort_single', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSortSingle C.Z3_mk_fpa_sort_single
func MkFpaSortSingle(c Context) Sort

/**
  \brief Create the single-precision (32-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_single.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_16
  \sa Z3_mk_fpa_sort_64
  \sa Z3_mk_fpa_sort_128

  def_API('Z3_mk_fpa_sort_32', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSort32 C.Z3_mk_fpa_sort_32
func MkFpaSort32(c Context) Sort

/**
  \brief Create the double-precision (64-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_64.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_half
  \sa Z3_mk_fpa_sort_single
  \sa Z3_mk_fpa_sort_quadruple

  def_API('Z3_mk_fpa_sort_double', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSortDouble C.Z3_mk_fpa_sort_double
func MkFpaSortDouble(c Context) Sort

/**
  \brief Create the double-precision (64-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_double.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_16
  \sa Z3_mk_fpa_sort_32
  \sa Z3_mk_fpa_sort_128

  def_API('Z3_mk_fpa_sort_64', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSort64 C.Z3_mk_fpa_sort_64
func MkFpaSort64(c Context) Sort

/**
  \brief Create the quadruple-precision (128-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_128.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_half
  \sa Z3_mk_fpa_sort_single
  \sa Z3_mk_fpa_sort_double

  def_API('Z3_mk_fpa_sort_quadruple', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSortQuadruple C.Z3_mk_fpa_sort_quadruple
func MkFpaSortQuadruple(c Context) Sort

/**
  \brief Create the quadruple-precision (128-bit) FloatingPoint sort.

  This is the same as #Z3_mk_fpa_sort_quadruple.

  \param c logical context

  \sa Z3_mk_fpa_sort
  \sa Z3_mk_fpa_sort_16
  \sa Z3_mk_fpa_sort_32
  \sa Z3_mk_fpa_sort_64

  def_API('Z3_mk_fpa_sort_128', SORT, (_in(CONTEXT),))
*/
//go:linkname MkFpaSort128 C.Z3_mk_fpa_sort_128
func MkFpaSort128(c Context) Sort

/**
  \brief Create a floating-point NaN of sort \c s.

  \param c logical context
  \param s target sort

  \sa Z3_mk_fpa_inf
  \sa Z3_mk_fpa_is_nan
  \sa Z3_mk_fpa_zero

  def_API('Z3_mk_fpa_nan', AST, (_in(CONTEXT),_in(SORT)))
*/
//go:linkname MkFpaNan C.Z3_mk_fpa_nan
func MkFpaNan(c Context, s Sort) Ast

/**
  \brief Create a floating-point infinity of sort \c s.

  \param c logical context
  \param s target sort
  \param negative indicates whether the result should be negative

  When \c negative is \c true, -oo will be generated instead of +oo.

  \sa Z3_mk_fpa_is_infinite
  \sa Z3_mk_fpa_nan
  \sa Z3_mk_fpa_zero

  def_API('Z3_mk_fpa_inf', AST, (_in(CONTEXT),_in(SORT),_in(BOOL)))
*/
//go:linkname MkFpaInf C.Z3_mk_fpa_inf
func MkFpaInf(c Context, s Sort, negative bool) Ast

/**
  \brief Create a floating-point zero of sort \c s.

  \param c logical context
  \param s target sort
  \param negative indicates whether the result should be negative

  When \c negative is \c true, -zero will be generated instead of +zero.

  \sa Z3_mk_fpa_inf
  \sa Z3_mk_fpa_is_zero
  \sa Z3_mk_fpa_nan

  def_API('Z3_mk_fpa_zero', AST, (_in(CONTEXT),_in(SORT),_in(BOOL)))
*/
//go:linkname MkFpaZero C.Z3_mk_fpa_zero
func MkFpaZero(c Context, s Sort, negative bool) Ast

/**
  \brief Create an expression of FloatingPoint sort from three bit-vector expressions.

  This is the operator named `fp' in the SMT FP theory definition.
  Note that \c sgn is required to be a bit-vector of size 1. Significand and exponent
  are required to be longer than 1 and 2 respectively. The FloatingPoint sort
  of the resulting expression is automatically determined from the bit-vector sizes
  of the arguments. The exponent is assumed to be in IEEE-754 biased representation.

  \param c logical context
  \param sgn sign
  \param exp exponent
  \param sig significand

  \sa Z3_mk_fpa_numeral_double
  \sa Z3_mk_fpa_numeral_float
  \sa Z3_mk_fpa_numeral_int
  \sa Z3_mk_fpa_numeral_int_uint
  \sa Z3_mk_fpa_numeral_int64_uint64
  \sa Z3_mk_numeral

  def_API('Z3_mk_fpa_fp', AST, (_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname MkFpaFp C.Z3_mk_fpa_fp
func MkFpaFp(c Context, sgn Ast, exp Ast, sig Ast) Ast

/**
  \brief Create a numeral of FloatingPoint sort from a float.

  This function is used to create numerals that fit in a float value.
  It is slightly faster than #Z3_mk_numeral since it is not necessary to parse a string.

  \param c logical context
  \param v value
  \param ty sort

  \c ty must be a FloatingPoint sort

  \sa Z3_mk_fpa_fp
  \sa Z3_mk_fpa_numeral_double
  \sa Z3_mk_fpa_numeral_int
  \sa Z3_mk_fpa_numeral_int_uint
  \sa Z3_mk_fpa_numeral_int64_uint64
  \sa Z3_mk_numeral

  def_API('Z3_mk_fpa_numeral_float', AST, (_in(CONTEXT), _in(FLOAT), _in(SORT)))
*/
//go:linkname MkFpaNumeralFloat C.Z3_mk_fpa_numeral_float
func MkFpaNumeralFloat(c Context, v float32, ty Sort) Ast

/**
  \brief Create a numeral of FloatingPoint sort from a double.

  This function is used to create numerals that fit in a double value.
  It is slightly faster than #Z3_mk_numeral since it is not necessary to parse a string.

  \param c logical context
  \param v value
  \param ty sort

  \c ty must be a FloatingPoint sort

  \sa Z3_mk_fpa_fp
  \sa Z3_mk_fpa_numeral_float
  \sa Z3_mk_fpa_numeral_int
  \sa Z3_mk_fpa_numeral_int_uint
  \sa Z3_mk_fpa_numeral_int64_uint64
  \sa Z3_mk_numeral

  def_API('Z3_mk_fpa_numeral_double', AST, (_in(CONTEXT), _in(DOUBLE), _in(SORT)))
*/
//go:linkname MkFpaNumeralDouble C.Z3_mk_fpa_numeral_double
func MkFpaNumeralDouble(c Context, v float64, ty Sort) Ast

/**
  \brief Create a numeral of FloatingPoint sort from a signed integer.

  \param c logical context
  \param v value
  \param ty result sort

  \c ty must be a FloatingPoint sort

  \sa Z3_mk_fpa_fp
  \sa Z3_mk_fpa_numeral_double
  \sa Z3_mk_fpa_numeral_float
  \sa Z3_mk_fpa_numeral_int_uint
  \sa Z3_mk_fpa_numeral_int64_uint64
  \sa Z3_mk_numeral

  def_API('Z3_mk_fpa_numeral_int', AST, (_in(CONTEXT), _in(INT), _in(SORT)))
*/
//go:linkname MkFpaNumeralInt C.Z3_mk_fpa_numeral_int
func MkFpaNumeralInt(c Context, v c.Int, ty Sort) Ast

/**
  \brief Create a numeral of FloatingPoint sort from a sign bit and two integers.

  \param c logical context
  \param sgn sign bit (true == negative)
  \param sig significand
  \param exp exponent
  \param ty result sort

  \c ty must be a FloatingPoint sort

  \sa Z3_mk_fpa_fp
  \sa Z3_mk_fpa_numeral_double
  \sa Z3_mk_fpa_numeral_float
  \sa Z3_mk_fpa_numeral_int
  \sa Z3_mk_fpa_numeral_int64_uint64
  \sa Z3_mk_numeral

  def_API('Z3_mk_fpa_numeral_int_uint', AST, (_in(CONTEXT), _in(BOOL), _in(INT), _in(UINT), _in(SORT)))
*/
//go:linkname MkFpaNumeralIntUint C.Z3_mk_fpa_numeral_int_uint
func MkFpaNumeralIntUint(c Context, sgn bool, exp c.Int, sig c.Uint, ty Sort) Ast

/**
  \brief Create a numeral of FloatingPoint sort from a sign bit and two 64-bit integers.

  \param c logical context
  \param sgn sign bit (true == negative)
  \param sig significand
  \param exp exponent
  \param ty result sort

  \c ty must be a FloatingPoint sort

  \sa Z3_mk_fpa_fp
  \sa Z3_mk_fpa_numeral_double
  \sa Z3_mk_fpa_numeral_float
  \sa Z3_mk_fpa_numeral_int
  \sa Z3_mk_fpa_numeral_int_uint
  \sa Z3_mk_numeral

  def_API('Z3_mk_fpa_numeral_int64_uint64', AST, (_in(CONTEXT), _in(BOOL), _in(INT64), _in(UINT64), _in(SORT)))
*/
//go:linkname MkFpaNumeralInt64Uint64 C.Z3_mk_fpa_numeral_int64_uint64
func MkFpaNumeralInt64Uint64(c Context, sgn bool, exp int64, sig uint64, ty Sort) Ast

/**
  \brief Floating-point absolute value

  \param c logical context
  \param t term of FloatingPoint sort

  \sa Z3_mk_fpa_is_negative
  \sa Z3_mk_fpa_is_positive
  \sa Z3_mk_fpa_neg

  def_API('Z3_mk_fpa_abs', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaAbs C.Z3_mk_fpa_abs
func MkFpaAbs(c Context, t Ast) Ast

/**
  \brief Floating-point negation

  \param c logical context
  \param t term of FloatingPoint sort

  \sa Z3_mk_fpa_abs
  \sa Z3_mk_fpa_is_negative
  \sa Z3_mk_fpa_is_positive

  def_API('Z3_mk_fpa_neg', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaNeg C.Z3_mk_fpa_neg
func MkFpaNeg(c Context, t Ast) Ast

/**
  \brief Floating-point addition

  \param c logical context
  \param rm term of RoundingMode sort
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c rm must be of RoundingMode sort, \c t1 and \c t2 must have the same FloatingPoint sort.

  def_API('Z3_mk_fpa_add', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(AST)))
*/
//go:linkname MkFpaAdd C.Z3_mk_fpa_add
func MkFpaAdd(c Context, rm Ast, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point subtraction

  \param c logical context
  \param rm term of RoundingMode sort
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c rm must be of RoundingMode sort, \c t1 and \c t2 must have the same FloatingPoint sort.

  def_API('Z3_mk_fpa_sub', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(AST)))
*/
//go:linkname MkFpaSub C.Z3_mk_fpa_sub
func MkFpaSub(c Context, rm Ast, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point multiplication

  \param c logical context
  \param rm term of RoundingMode sort
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c rm must be of RoundingMode sort, \c t1 and \c t2 must have the same FloatingPoint sort.

  def_API('Z3_mk_fpa_mul', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(AST)))
*/
//go:linkname MkFpaMul C.Z3_mk_fpa_mul
func MkFpaMul(c Context, rm Ast, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point division

  \param c logical context
  \param rm term of RoundingMode sort
  \param t1 term of FloatingPoint sort.
  \param t2 term of FloatingPoint sort

  The nodes \c rm must be of RoundingMode sort, \c t1 and \c t2 must have the same FloatingPoint sort.

  def_API('Z3_mk_fpa_div', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(AST)))
*/
//go:linkname MkFpaDiv C.Z3_mk_fpa_div
func MkFpaDiv(c Context, rm Ast, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point fused multiply-add.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort
  \param t3 term of FloatingPoint sort

  The result is \ccode{round((t1 * t2) + t3)}.

  \c rm must be of RoundingMode sort, \c t1, \c t2, and \c t3 must have the same FloatingPoint sort.

  def_API('Z3_mk_fpa_fma', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(AST),_in(AST)))
*/
//go:linkname MkFpaFma C.Z3_mk_fpa_fma
func MkFpaFma(c Context, rm Ast, t1 Ast, t2 Ast, t3 Ast) Ast

/**
  \brief Floating-point square root

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of FloatingPoint sort

  \c rm must be of RoundingMode sort, \c t must have FloatingPoint sort.

  def_API('Z3_mk_fpa_sqrt', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaSqrt C.Z3_mk_fpa_sqrt
func MkFpaSqrt(c Context, rm Ast, t Ast) Ast

/**
  \brief Floating-point remainder

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1 and \c t2 must have the same FloatingPoint sort.

  def_API('Z3_mk_fpa_rem', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaRem C.Z3_mk_fpa_rem
func MkFpaRem(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point roundToIntegral. Rounds a floating-point number to
  the closest integer, again represented as a floating-point number.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of FloatingPoint sort

  \c t must be of FloatingPoint sort.

  def_API('Z3_mk_fpa_round_to_integral', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaRoundToIntegral C.Z3_mk_fpa_round_to_integral
func MkFpaRoundToIntegral(c Context, rm Ast, t Ast) Ast

/**
  \brief Minimum of floating-point numbers.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1, \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_max

  def_API('Z3_mk_fpa_min', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaMin C.Z3_mk_fpa_min
func MkFpaMin(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Maximum of floating-point numbers.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1, \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_min

  def_API('Z3_mk_fpa_max', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaMax C.Z3_mk_fpa_max
func MkFpaMax(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point less than or equal.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1 and \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_eq
  \sa Z3_mk_fpa_geq
  \sa Z3_mk_fpa_gt
  \sa Z3_mk_fpa_lt

  def_API('Z3_mk_fpa_leq', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaLeq C.Z3_mk_fpa_leq
func MkFpaLeq(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point less than.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1 and \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_eq
  \sa Z3_mk_fpa_geq
  \sa Z3_mk_fpa_gt
  \sa Z3_mk_fpa_leq

  def_API('Z3_mk_fpa_lt', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaLt C.Z3_mk_fpa_lt
func MkFpaLt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point greater than or equal.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1 and \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_eq
  \sa Z3_mk_fpa_gt
  \sa Z3_mk_fpa_leq
  \sa Z3_mk_fpa_lt

  def_API('Z3_mk_fpa_geq', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaGeq C.Z3_mk_fpa_geq
func MkFpaGeq(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point greater than.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  \c t1 and \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_eq
  \sa Z3_mk_fpa_geq
  \sa Z3_mk_fpa_leq
  \sa Z3_mk_fpa_lt

  def_API('Z3_mk_fpa_gt', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaGt C.Z3_mk_fpa_gt
func MkFpaGt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Floating-point equality.

  \param c logical context
  \param t1 term of FloatingPoint sort
  \param t2 term of FloatingPoint sort

  Note that this is IEEE 754 equality (as opposed to SMT-LIB \ccode{=}).

  \c t1 and \c t2 must have the same FloatingPoint sort.

  \sa Z3_mk_fpa_geq
  \sa Z3_mk_fpa_gt
  \sa Z3_mk_fpa_leq
  \sa Z3_mk_fpa_lt

  def_API('Z3_mk_fpa_eq', AST, (_in(CONTEXT),_in(AST),_in(AST)))
*/
//go:linkname MkFpaEq C.Z3_mk_fpa_eq
func MkFpaEq(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Predicate indicating whether \c t is a normal floating-point number.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_is_infinite
  \sa Z3_mk_fpa_is_nan
  \sa Z3_mk_fpa_is_subnormal
  \sa Z3_mk_fpa_is_zero

  def_API('Z3_mk_fpa_is_normal', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsNormal C.Z3_mk_fpa_is_normal
func MkFpaIsNormal(c Context, t Ast) Ast

/**
  \brief Predicate indicating whether \c t is a subnormal floating-point number.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_is_infinite
  \sa Z3_mk_fpa_is_nan
  \sa Z3_mk_fpa_is_normal
  \sa Z3_mk_fpa_is_zero

  def_API('Z3_mk_fpa_is_subnormal', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsSubnormal C.Z3_mk_fpa_is_subnormal
func MkFpaIsSubnormal(c Context, t Ast) Ast

/**
  \brief Predicate indicating whether \c t is a floating-point number with zero value, i.e., +zero or -zero.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_is_infinite
  \sa Z3_mk_fpa_is_nan
  \sa Z3_mk_fpa_is_normal
  \sa Z3_mk_fpa_is_subnormal
  \sa Z3_mk_fpa_zero

  def_API('Z3_mk_fpa_is_zero', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsZero C.Z3_mk_fpa_is_zero
func MkFpaIsZero(c Context, t Ast) Ast

/**
  \brief Predicate indicating whether \c t is a floating-point number representing +oo or -oo.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_inf
  \sa Z3_mk_fpa_is_nan
  \sa Z3_mk_fpa_is_normal
  \sa Z3_mk_fpa_is_subnormal
  \sa Z3_mk_fpa_is_zero

  def_API('Z3_mk_fpa_is_infinite', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsInfinite C.Z3_mk_fpa_is_infinite
func MkFpaIsInfinite(c Context, t Ast) Ast

/**
  \brief Predicate indicating whether \c t is a NaN.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_is_infinite
  \sa Z3_mk_fpa_is_normal
  \sa Z3_mk_fpa_is_subnormal
  \sa Z3_mk_fpa_is_zero
  \sa Z3_mk_fpa_nan

  def_API('Z3_mk_fpa_is_nan', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsNan C.Z3_mk_fpa_is_nan
func MkFpaIsNan(c Context, t Ast) Ast

/**
  \brief Predicate indicating whether \c t is a negative floating-point number.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_abs
  \sa Z3_mk_fpa_is_positive
  \sa Z3_mk_fpa_neg

  def_API('Z3_mk_fpa_is_negative', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsNegative C.Z3_mk_fpa_is_negative
func MkFpaIsNegative(c Context, t Ast) Ast

/**
  \brief Predicate indicating whether \c t is a positive floating-point number.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort.

  \sa Z3_mk_fpa_abs
  \sa Z3_mk_fpa_is_negative
  \sa Z3_mk_fpa_neg

  def_API('Z3_mk_fpa_is_positive', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaIsPositive C.Z3_mk_fpa_is_positive
func MkFpaIsPositive(c Context, t Ast) Ast

/**
  \brief Conversion of a single IEEE 754-2008 bit-vector into a floating-point number.

  Produces a term that represents the conversion of a bit-vector term \c bv to a
  floating-point term of sort \c s.

  \param c logical context
  \param bv a bit-vector term
  \param s floating-point sort

  \c s must be a FloatingPoint sort, \c t must be of bit-vector sort, and the bit-vector
  size of \c bv must be equal to \ccode{ebits+sbits} of \c s. The format of the bit-vector is
  as defined by the IEEE 754-2008 interchange format.

  def_API('Z3_mk_fpa_to_fp_bv', AST, (_in(CONTEXT),_in(AST),_in(SORT)))
*/
//go:linkname MkFpaToFpBv C.Z3_mk_fpa_to_fp_bv
func MkFpaToFpBv(c Context, bv Ast, s Sort) Ast

/**
  \brief Conversion of a FloatingPoint term into another term of different FloatingPoint sort.

  Produces a term that represents the conversion of a floating-point term \c t to a
  floating-point term of sort \c s. If necessary, the result will be rounded according
  to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of FloatingPoint sort
  \param s floating-point sort

  \c s must be a FloatingPoint sort, \c rm must be of RoundingMode sort, \c t must be of floating-point sort.

  def_API('Z3_mk_fpa_to_fp_float', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(SORT)))
*/
//go:linkname MkFpaToFpFloat C.Z3_mk_fpa_to_fp_float
func MkFpaToFpFloat(c Context, rm Ast, t Ast, s Sort) Ast

/**
  \brief Conversion of a term of real sort into a term of FloatingPoint sort.

  Produces a term that represents the conversion of term \c t of real sort into a
  floating-point term of sort \c s. If necessary, the result will be rounded according
  to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of Real sort
  \param s floating-point sort

  \c s must be a FloatingPoint sort, \c rm must be of RoundingMode sort, \c t must be of real sort.

  def_API('Z3_mk_fpa_to_fp_real', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(SORT)))
*/
//go:linkname MkFpaToFpReal C.Z3_mk_fpa_to_fp_real
func MkFpaToFpReal(c Context, rm Ast, t Ast, s Sort) Ast

/**
  \brief Conversion of a 2's complement signed bit-vector term into a term of FloatingPoint sort.

  Produces a term that represents the conversion of the bit-vector term \c t into a
  floating-point term of sort \c s. The bit-vector \c t is taken to be in signed
  2's complement format. If necessary, the result will be rounded according
  to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of bit-vector sort
  \param s floating-point sort

  \c s must be a FloatingPoint sort, \c rm must be of RoundingMode sort, \c t must be of bit-vector sort.

  def_API('Z3_mk_fpa_to_fp_signed', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(SORT)))
*/
//go:linkname MkFpaToFpSigned C.Z3_mk_fpa_to_fp_signed
func MkFpaToFpSigned(c Context, rm Ast, t Ast, s Sort) Ast

/**
  \brief Conversion of a 2's complement unsigned bit-vector term into a term of FloatingPoint sort.

  Produces a term that represents the conversion of the bit-vector term \c t into a
  floating-point term of sort \c s. The bit-vector \c t is taken to be in unsigned
  2's complement format. If necessary, the result will be rounded according
  to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of bit-vector sort
  \param s floating-point sort

  \c s must be a FloatingPoint sort, \c rm must be of RoundingMode sort, \c t must be of bit-vector sort.

  def_API('Z3_mk_fpa_to_fp_unsigned', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(SORT)))
*/
//go:linkname MkFpaToFpUnsigned C.Z3_mk_fpa_to_fp_unsigned
func MkFpaToFpUnsigned(c Context, rm Ast, t Ast, s Sort) Ast

/**
  \brief Conversion of a floating-point term into an unsigned bit-vector.

  Produces a term that represents the conversion of the floating-point term \c t into a
  bit-vector term of size \c sz in unsigned 2's complement format. If necessary, the result
  will be rounded according to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of FloatingPoint sort
  \param sz size of the resulting bit-vector

  def_API('Z3_mk_fpa_to_ubv', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(UINT)))
*/
//go:linkname MkFpaToUbv C.Z3_mk_fpa_to_ubv
func MkFpaToUbv(c Context, rm Ast, t Ast, sz c.Uint) Ast

/**
  \brief Conversion of a floating-point term into a signed bit-vector.

  Produces a term that represents the conversion of the floating-point term \c t into a
  bit-vector term of size \c sz in signed 2's complement format. If necessary, the result
  will be rounded according to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param t term of FloatingPoint sort
  \param sz size of the resulting bit-vector

  def_API('Z3_mk_fpa_to_sbv', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(UINT)))
*/
//go:linkname MkFpaToSbv C.Z3_mk_fpa_to_sbv
func MkFpaToSbv(c Context, rm Ast, t Ast, sz c.Uint) Ast

/**
  \brief Conversion of a floating-point term into a real-numbered term.

  Produces a term that represents the conversion of the floating-point term \c t into a
  real number. Note that this type of conversion will often result in non-linear
  constraints over real terms.

  \param c logical context
  \param t term of FloatingPoint sort

  def_API('Z3_mk_fpa_to_real', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaToReal C.Z3_mk_fpa_to_real
func MkFpaToReal(c Context, t Ast) Ast

/** @name Z3-specific floating-point extensions */
/**@{*/
/**
  \brief Retrieves the number of bits reserved for the exponent in a FloatingPoint sort.

  \param c logical context
  \param s FloatingPoint sort

  \sa Z3_fpa_get_sbits

  def_API('Z3_fpa_get_ebits', UINT, (_in(CONTEXT),_in(SORT)))
*/
//go:linkname FpaGetEbits C.Z3_fpa_get_ebits
func FpaGetEbits(c Context, s Sort) c.Uint

/**
  \brief Retrieves the number of bits reserved for the significand in a FloatingPoint sort.

  \param c logical context
  \param s FloatingPoint sort

  \sa Z3_fpa_get_ebits

  def_API('Z3_fpa_get_sbits', UINT, (_in(CONTEXT),_in(SORT)))
*/
//go:linkname FpaGetSbits C.Z3_fpa_get_sbits
func FpaGetSbits(c Context, s Sort) c.Uint

/**
  \brief Checks whether a given floating-point numeral is a NaN.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_inf
  \sa Z3_fpa_is_numeral_normal
  \sa Z3_fpa_is_numeral_subnormal
  \sa Z3_fpa_is_numeral_zero

  def_API('Z3_fpa_is_numeral_nan', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralNan C.Z3_fpa_is_numeral_nan
func FpaIsNumeralNan(c Context, t Ast) bool

/**
  \brief Checks whether a given floating-point numeral is a +oo or -oo.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_nan
  \sa Z3_fpa_is_numeral_normal
  \sa Z3_fpa_is_numeral_subnormal
  \sa Z3_fpa_is_numeral_zero

  def_API('Z3_fpa_is_numeral_inf', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralInf C.Z3_fpa_is_numeral_inf
func FpaIsNumeralInf(c Context, t Ast) bool

/**
  \brief Checks whether a given floating-point numeral is +zero or -zero.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_inf
  \sa Z3_fpa_is_numeral_nan
  \sa Z3_fpa_is_numeral_normal
  \sa Z3_fpa_is_numeral_subnormal

  def_API('Z3_fpa_is_numeral_zero', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralZero C.Z3_fpa_is_numeral_zero
func FpaIsNumeralZero(c Context, t Ast) bool

/**
  \brief Checks whether a given floating-point numeral is normal.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_inf
  \sa Z3_fpa_is_numeral_nan
  \sa Z3_fpa_is_numeral_subnormal
  \sa Z3_fpa_is_numeral_zero

  def_API('Z3_fpa_is_numeral_normal', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralNormal C.Z3_fpa_is_numeral_normal
func FpaIsNumeralNormal(c Context, t Ast) bool

/**
  \brief Checks whether a given floating-point numeral is subnormal.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_inf
  \sa Z3_fpa_is_numeral_nan
  \sa Z3_fpa_is_numeral_normal
  \sa Z3_fpa_is_numeral_zero

  def_API('Z3_fpa_is_numeral_subnormal', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralSubnormal C.Z3_fpa_is_numeral_subnormal
func FpaIsNumeralSubnormal(c Context, t Ast) bool

/**
  \brief Checks whether a given floating-point numeral is positive.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_negative

  def_API('Z3_fpa_is_numeral_positive', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralPositive C.Z3_fpa_is_numeral_positive
func FpaIsNumeralPositive(c Context, t Ast) bool

/**
  \brief Checks whether a given floating-point numeral is negative.

  \param c logical context
  \param t a floating-point numeral

  \sa Z3_fpa_is_numeral_positive

  def_API('Z3_fpa_is_numeral_negative', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaIsNumeralNegative C.Z3_fpa_is_numeral_negative
func FpaIsNumeralNegative(c Context, t Ast) bool

/**
  \brief Retrieves the sign of a floating-point literal as a bit-vector expression.

  \param c logical context
  \param t a floating-point numeral

  Remarks: NaN is an invalid argument.

  def_API('Z3_fpa_get_numeral_sign_bv', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaGetNumeralSignBv C.Z3_fpa_get_numeral_sign_bv
func FpaGetNumeralSignBv(c Context, t Ast) Ast

/**
  \brief Retrieves the significand of a floating-point literal as a bit-vector expression.

  \param c logical context
  \param t a floating-point numeral

  Remarks: NaN is an invalid argument.

  def_API('Z3_fpa_get_numeral_significand_bv', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaGetNumeralSignificandBv C.Z3_fpa_get_numeral_significand_bv
func FpaGetNumeralSignificandBv(c Context, t Ast) Ast

/**
  \brief Retrieves the sign of a floating-point literal.

  \param c logical context
  \param t a floating-point numeral
  \param sgn the retrieved sign
  \returns true if \c t corresponds to a floating point numeral, otherwise invokes exception handler or returns false

  Remarks: sets \c sgn to 0 if `t' is positive and to 1 otherwise, except for
  NaN, which is an invalid argument.

  def_API('Z3_fpa_get_numeral_sign', BOOL, (_in(CONTEXT), _in(AST), _out(INT)))
*/
//go:linkname FpaGetNumeralSign C.Z3_fpa_get_numeral_sign
func FpaGetNumeralSign(c Context, t Ast, sgn *c.Int) bool

/**
  \brief Return the significand value of a floating-point numeral as a string.

  \param c logical context
  \param t a floating-point numeral
  \returns true if \c t corresponds to a floating point numeral, otherwise invokes exception handler or returns false

  Remarks: The significand \c s is always \ccode{0.0 <= s < 2.0}; the resulting string is long
  enough to represent the real significand precisely.

  def_API('Z3_fpa_get_numeral_significand_string', STRING, (_in(CONTEXT), _in(AST)))
*/
//go:linkname FpaGetNumeralSignificandString C.Z3_fpa_get_numeral_significand_string
func FpaGetNumeralSignificandString(c Context, t Ast) String

/**
  \brief Return the significand value of a floating-point numeral as a uint64.

  \param c logical context
  \param t a floating-point numeral
  \param n pointer to output uint64

  Remarks: This function extracts the significand bits in `t`, without the
  hidden bit or normalization. Sets the \c Z3_INVALID_ARG error code if the
  significand does not fit into a \c uint64. NaN is an invalid argument.

  def_API('Z3_fpa_get_numeral_significand_uint64', BOOL, (_in(CONTEXT), _in(AST), _out(UINT64)))
*/
//go:linkname FpaGetNumeralSignificandUint64 C.Z3_fpa_get_numeral_significand_uint64
func FpaGetNumeralSignificandUint64(c Context, t Ast, n *uint64) bool

/**
      \brief Return the exponent value of a floating-point numeral as a string.

      \param c logical context
      \param t a floating-point numeral
      \param biased flag to indicate whether the result is in biased representation
      \returns true if \c t corresponds to a floating point numeral, otherwise invokes exception handler or returns false

      Remarks: This function extracts the exponent in `t`, without normalization.
      NaN is an invalid argument.

  def_API('Z3_fpa_get_numeral_exponent_string', STRING, (_in(CONTEXT), _in(AST), _in(BOOL)))
*/
//go:linkname FpaGetNumeralExponentString C.Z3_fpa_get_numeral_exponent_string
func FpaGetNumeralExponentString(c Context, t Ast, biased bool) String

/**
  \brief Return the exponent value of a floating-point numeral as a signed 64-bit integer

  \param c logical context
  \param t a floating-point numeral
  \param n exponent
  \param biased flag to indicate whether the result is in biased representation
  \returns true if \c t corresponds to a floating point numeral, otherwise invokes exception handler or returns false

  Remarks: This function extracts the exponent in `t`, without normalization.
  NaN is an invalid argument.

  def_API('Z3_fpa_get_numeral_exponent_int64', BOOL, (_in(CONTEXT), _in(AST), _out(INT64), _in(BOOL)))
*/
//go:linkname FpaGetNumeralExponentInt64 C.Z3_fpa_get_numeral_exponent_int64
func FpaGetNumeralExponentInt64(c Context, t Ast, n *int64, biased bool) bool

/**
  \brief Retrieves the exponent of a floating-point literal as a bit-vector expression.

  \param c logical context
  \param t a floating-point numeral
  \param biased flag to indicate whether the result is in biased representation

  Remarks: This function extracts the exponent in `t`, without normalization.
  NaN is an invalid arguments.

  def_API('Z3_fpa_get_numeral_exponent_bv', AST, (_in(CONTEXT), _in(AST), _in(BOOL)))
*/
//go:linkname FpaGetNumeralExponentBv C.Z3_fpa_get_numeral_exponent_bv
func FpaGetNumeralExponentBv(c Context, t Ast, biased bool) Ast

/**
  \brief Conversion of a floating-point term into a bit-vector term in IEEE 754-2008 format.

  \param c logical context
  \param t term of FloatingPoint sort

  \c t must have FloatingPoint sort. The size of the resulting bit-vector is automatically
  determined.

  Note that IEEE 754-2008 allows multiple different representations of NaN. This conversion
  knows only one NaN and it will always produce the same bit-vector representation of
  that NaN.

  def_API('Z3_mk_fpa_to_ieee_bv', AST, (_in(CONTEXT),_in(AST)))
*/
//go:linkname MkFpaToIeeeBv C.Z3_mk_fpa_to_ieee_bv
func MkFpaToIeeeBv(c Context, t Ast) Ast

/**
  \brief Conversion of a real-sorted significand and an integer-sorted exponent into a term of FloatingPoint sort.

  Produces a term that represents the conversion of \ccode{sig * 2^exp} into a
  floating-point term of sort \c s. If necessary, the result will be rounded
  according to rounding mode \c rm.

  \param c logical context
  \param rm term of RoundingMode sort
  \param exp exponent term of Int sort
  \param sig significand term of Real sort
  \param s FloatingPoint sort

  \c s must be a FloatingPoint sort, \c rm must be of RoundingMode sort, \c exp must be of int sort, \c sig must be of real sort.

  def_API('Z3_mk_fpa_to_fp_int_real', AST, (_in(CONTEXT),_in(AST),_in(AST),_in(AST),_in(SORT)))
*/
//go:linkname MkFpaToFpIntReal C.Z3_mk_fpa_to_fp_int_real
func MkFpaToFpIntReal(c Context, rm Ast, exp Ast, sig Ast, s Sort) Ast
