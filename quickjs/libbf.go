package quickjs

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const LIMB_LOG2_BITS = 6
const LIMB_DIGITS = 19
const BF_EXP_BITS_MIN = 3
const BF_PREC_MIN = 2
const BF_RND_MASK = 0x7
const BF_EXP_BITS_SHIFT = 5
const BF_EXP_BITS_MASK = 0x3f
const BF_RADIX_MAX = 36

type Int128T struct {
	Unused [8]uint8
}
type Uint128T struct {
	Unused [8]uint8
}
type SlimbT int64
type LimbT uint64
type DlimbT Uint128T

/* +/-zero is represented with expn = BF_EXP_ZERO and len = 0,
   +/-infinity is represented with expn = BF_EXP_INF and len = 0,
   NaN is represented with expn = BF_EXP_NAN and len = 0 (sign is ignored)
*/

type BfT struct {
	Ctx  *BfContextT
	Sign c.Int
	Expn SlimbT
	Len  LimbT
	Tab  *LimbT
}

type BfContextT struct {
	ReallocOpaque unsafe.Pointer
	ReallocFunc   *unsafe.Pointer
	Log2Cache     BFConstCache
	PiCache       BFConstCache
	NttState      *BFNTTState
}

type BfdecT struct {
	Ctx  *BfContextT
	Sign c.Int
	Expn SlimbT
	Len  LimbT
	Tab  *LimbT
}
type BfRndT c.Int

const (
	BFRNDN  BfRndT = 0
	BFRNDZ  BfRndT = 1
	BFRNDD  BfRndT = 2
	BFRNDU  BfRndT = 3
	BFRNDNA BfRndT = 4
	BFRNDA  BfRndT = 5
	BFRNDF  BfRndT = 6
)

type BfFlagsT uint32

// llgo:type C
type BfReallocFuncT func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer

type BFConstCache struct {
	Val  BfT
	Prec LimbT
}

type BFNTTState struct {
	Unused [8]uint8
}

// llgo:link (*BfContextT).BfContextInit C.bf_context_init
func (recv_ *BfContextT) BfContextInit(realloc_func BfReallocFuncT, realloc_opaque unsafe.Pointer) {
}

// llgo:link (*BfContextT).BfContextEnd C.bf_context_end
func (recv_ *BfContextT) BfContextEnd() {
}

/* free memory allocated for the bf cache data */
// llgo:link (*BfContextT).BfClearCache C.bf_clear_cache
func (recv_ *BfContextT) BfClearCache() {
}

// llgo:link (*BfContextT).BfInit C.bf_init
func (recv_ *BfContextT) BfInit(r *BfT) {
}

// llgo:link (*BfT).BfSetUi C.bf_set_ui
func (recv_ *BfT) BfSetUi(a uint64) c.Int {
	return 0
}

// llgo:link (*BfT).BfSetSi C.bf_set_si
func (recv_ *BfT) BfSetSi(a int64) c.Int {
	return 0
}

// llgo:link (*BfT).BfSetNan C.bf_set_nan
func (recv_ *BfT) BfSetNan() {
}

// llgo:link (*BfT).BfSetZero C.bf_set_zero
func (recv_ *BfT) BfSetZero(is_neg c.Int) {
}

// llgo:link (*BfT).BfSetInf C.bf_set_inf
func (recv_ *BfT) BfSetInf(is_neg c.Int) {
}

// llgo:link (*BfT).BfSet C.bf_set
func (recv_ *BfT) BfSet(a *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfMove C.bf_move
func (recv_ *BfT) BfMove(a *BfT) {
}

// llgo:link (*BfT).BfGetFloat64 C.bf_get_float64
func (recv_ *BfT) BfGetFloat64(pres *float64, rnd_mode BfRndT) c.Int {
	return 0
}

// llgo:link (*BfT).BfSetFloat64 C.bf_set_float64
func (recv_ *BfT) BfSetFloat64(d float64) c.Int {
	return 0
}

// llgo:link (*BfT).BfCmpu C.bf_cmpu
func (recv_ *BfT) BfCmpu(b *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfCmpFull C.bf_cmp_full
func (recv_ *BfT) BfCmpFull(b *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfCmp C.bf_cmp
func (recv_ *BfT) BfCmp(b *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAdd C.bf_add
func (recv_ *BfT) BfAdd(a *BfT, b *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfSub C.bf_sub
func (recv_ *BfT) BfSub(a *BfT, b *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAddSi C.bf_add_si
func (recv_ *BfT) BfAddSi(a *BfT, b1 int64, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfMul C.bf_mul
func (recv_ *BfT) BfMul(a *BfT, b *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfMulUi C.bf_mul_ui
func (recv_ *BfT) BfMulUi(a *BfT, b1 uint64, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfMulSi C.bf_mul_si
func (recv_ *BfT) BfMulSi(a *BfT, b1 int64, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfMul2exp C.bf_mul_2exp
func (recv_ *BfT) BfMul2exp(e SlimbT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfDiv C.bf_div
func (recv_ *BfT) BfDiv(a *BfT, b *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfDivrem C.bf_divrem
func (recv_ *BfT) BfDivrem(r *BfT, a *BfT, b *BfT, prec LimbT, flags BfFlagsT, rnd_mode c.Int) c.Int {
	return 0
}

// llgo:link (*BfT).BfRem C.bf_rem
func (recv_ *BfT) BfRem(a *BfT, b *BfT, prec LimbT, flags BfFlagsT, rnd_mode c.Int) c.Int {
	return 0
}

// llgo:link (*SlimbT).BfRemquo C.bf_remquo
func (recv_ *SlimbT) BfRemquo(r *BfT, a *BfT, b *BfT, prec LimbT, flags BfFlagsT, rnd_mode c.Int) c.Int {
	return 0
}

/* round to integer with infinite precision */
// llgo:link (*BfT).BfRint C.bf_rint
func (recv_ *BfT) BfRint(rnd_mode c.Int) c.Int {
	return 0
}

// llgo:link (*BfT).BfRound C.bf_round
func (recv_ *BfT) BfRound(prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfSqrtrem C.bf_sqrtrem
func (recv_ *BfT) BfSqrtrem(rem1 *BfT, a *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfSqrt C.bf_sqrt
func (recv_ *BfT) BfSqrt(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfGetExpMin C.bf_get_exp_min
func (recv_ *BfT) BfGetExpMin() SlimbT {
	return 0
}

// llgo:link (*BfT).BfLogicOr C.bf_logic_or
func (recv_ *BfT) BfLogicOr(a *BfT, b *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfLogicXor C.bf_logic_xor
func (recv_ *BfT) BfLogicXor(a *BfT, b *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfLogicAnd C.bf_logic_and
func (recv_ *BfT) BfLogicAnd(a *BfT, b *BfT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAtof C.bf_atof
func (recv_ *BfT) BfAtof(str *int8, pnext **int8, radix c.Int, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

/* this version accepts prec = BF_PREC_INF and returns the radix
   exponent */
// llgo:link (*BfT).BfAtof2 C.bf_atof2
func (recv_ *BfT) BfAtof2(pexponent *SlimbT, str *int8, pnext **int8, radix c.Int, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfMulPowRadix C.bf_mul_pow_radix
func (recv_ *BfT) BfMulPowRadix(T *BfT, radix LimbT, expn SlimbT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

//go:linkname BfFtoa C.bf_ftoa
func BfFtoa(plen *uintptr, a *BfT, radix c.Int, prec LimbT, flags BfFlagsT) *int8

//go:linkname BfGetInt32 C.bf_get_int32
func BfGetInt32(pres *c.Int, a *BfT, flags c.Int) c.Int

//go:linkname BfGetInt64 C.bf_get_int64
func BfGetInt64(pres *int64, a *BfT, flags c.Int) c.Int

//go:linkname BfGetUint64 C.bf_get_uint64
func BfGetUint64(pres *uint64, a *BfT) c.Int

/* the following functions are exported for testing only. */
//go:linkname MpPrintStr C.mp_print_str
func MpPrintStr(str *int8, tab *LimbT, n LimbT)

//go:linkname BfPrintStr C.bf_print_str
func BfPrintStr(str *int8, a *BfT)

// llgo:link (*BfT).BfResize C.bf_resize
func (recv_ *BfT) BfResize(len LimbT) c.Int {
	return 0
}

//go:linkname BfGetFftSize C.bf_get_fft_size
func BfGetFftSize(pdpl *c.Int, pnb_mods *c.Int, len LimbT) c.Int

// llgo:link (*BfT).BfNormalizeAndRound C.bf_normalize_and_round
func (recv_ *BfT) BfNormalizeAndRound(prec1 LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfCanRound C.bf_can_round
func (recv_ *BfT) BfCanRound(prec SlimbT, rnd_mode BfRndT, k SlimbT) c.Int {
	return 0
}

// llgo:link SlimbT.BfMulLog2Radix C.bf_mul_log2_radix
func (recv_ SlimbT) BfMulLog2Radix(radix c.Uint, is_inv c.Int, is_ceil1 c.Int) SlimbT {
	return 0
}

// llgo:link (*BfContextT).MpMul C.mp_mul
func (recv_ *BfContextT) MpMul(result *LimbT, op1 *LimbT, op1_size LimbT, op2 *LimbT, op2_size LimbT) c.Int {
	return 0
}

// llgo:link (*LimbT).MpAdd C.mp_add
func (recv_ *LimbT) MpAdd(op1 *LimbT, op2 *LimbT, n LimbT, carry LimbT) LimbT {
	return 0
}

// llgo:link (*LimbT).MpAddUi C.mp_add_ui
func (recv_ *LimbT) MpAddUi(b LimbT, n uintptr) LimbT {
	return 0
}

// llgo:link (*BfContextT).MpSqrtrem C.mp_sqrtrem
func (recv_ *BfContextT) MpSqrtrem(tabs *LimbT, taba *LimbT, n LimbT) c.Int {
	return 0
}

// llgo:link (*BfContextT).MpRecip C.mp_recip
func (recv_ *BfContextT) MpRecip(tabr *LimbT, taba *LimbT, n LimbT) c.Int {
	return 0
}

// llgo:link LimbT.BfIsqrt C.bf_isqrt
func (recv_ LimbT) BfIsqrt() LimbT {
	return 0
}

/* transcendental functions */
// llgo:link (*BfT).BfConstLog2 C.bf_const_log2
func (recv_ *BfT) BfConstLog2(prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfConstPi C.bf_const_pi
func (recv_ *BfT) BfConstPi(prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfExp C.bf_exp
func (recv_ *BfT) BfExp(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfLog C.bf_log
func (recv_ *BfT) BfLog(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfPow C.bf_pow
func (recv_ *BfT) BfPow(x *BfT, y *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfCos C.bf_cos
func (recv_ *BfT) BfCos(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfSin C.bf_sin
func (recv_ *BfT) BfSin(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfTan C.bf_tan
func (recv_ *BfT) BfTan(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAtan C.bf_atan
func (recv_ *BfT) BfAtan(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAtan2 C.bf_atan2
func (recv_ *BfT) BfAtan2(y *BfT, x *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAsin C.bf_asin
func (recv_ *BfT) BfAsin(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfT).BfAcos C.bf_acos
func (recv_ *BfT) BfAcos(a *BfT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecSetUi C.bfdec_set_ui
func (recv_ *BfdecT) BfdecSetUi(a uint64) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecSetSi C.bfdec_set_si
func (recv_ *BfdecT) BfdecSetSi(a int64) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecCmpu C.bfdec_cmpu
func (recv_ *BfdecT) BfdecCmpu(b *BfdecT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecAdd C.bfdec_add
func (recv_ *BfdecT) BfdecAdd(a *BfdecT, b *BfdecT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecSub C.bfdec_sub
func (recv_ *BfdecT) BfdecSub(a *BfdecT, b *BfdecT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecAddSi C.bfdec_add_si
func (recv_ *BfdecT) BfdecAddSi(a *BfdecT, b1 int64, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecMul C.bfdec_mul
func (recv_ *BfdecT) BfdecMul(a *BfdecT, b *BfdecT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecMulSi C.bfdec_mul_si
func (recv_ *BfdecT) BfdecMulSi(a *BfdecT, b1 int64, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecDiv C.bfdec_div
func (recv_ *BfdecT) BfdecDiv(a *BfdecT, b *BfdecT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecDivrem C.bfdec_divrem
func (recv_ *BfdecT) BfdecDivrem(r *BfdecT, a *BfdecT, b *BfdecT, prec LimbT, flags BfFlagsT, rnd_mode c.Int) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecRem C.bfdec_rem
func (recv_ *BfdecT) BfdecRem(a *BfdecT, b *BfdecT, prec LimbT, flags BfFlagsT, rnd_mode c.Int) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecRint C.bfdec_rint
func (recv_ *BfdecT) BfdecRint(rnd_mode c.Int) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecSqrt C.bfdec_sqrt
func (recv_ *BfdecT) BfdecSqrt(a *BfdecT, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecRound C.bfdec_round
func (recv_ *BfdecT) BfdecRound(prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

//go:linkname BfdecGetInt32 C.bfdec_get_int32
func BfdecGetInt32(pres *c.Int, a *BfdecT) c.Int

// llgo:link (*BfdecT).BfdecPowUi C.bfdec_pow_ui
func (recv_ *BfdecT) BfdecPowUi(a *BfdecT, b LimbT) c.Int {
	return 0
}

//go:linkname BfdecFtoa C.bfdec_ftoa
func BfdecFtoa(plen *uintptr, a *BfdecT, prec LimbT, flags BfFlagsT) *int8

// llgo:link (*BfdecT).BfdecAtof C.bfdec_atof
func (recv_ *BfdecT) BfdecAtof(str *int8, pnext **int8, prec LimbT, flags BfFlagsT) c.Int {
	return 0
}

//go:linkname BfdecPrintStr C.bfdec_print_str
func BfdecPrintStr(str *int8, a *BfdecT)

// llgo:link (*BfdecT).BfdecResize C.bfdec_resize
func (recv_ *BfdecT) BfdecResize(len LimbT) c.Int {
	return 0
}

// llgo:link (*BfdecT).BfdecNormalizeAndRound C.bfdec_normalize_and_round
func (recv_ *BfdecT) BfdecNormalizeAndRound(prec1 LimbT, flags BfFlagsT) c.Int {
	return 0
}
