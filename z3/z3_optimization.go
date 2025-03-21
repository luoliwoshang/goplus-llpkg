package z3

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

// llgo:type C
type ModelEh func(unsafe.Pointer)

/** @name Optimization facilities */
/**@{*/
/**
  \brief Create a new optimize context.

  \remark User must use #Z3_optimize_inc_ref and #Z3_optimize_dec_ref to manage optimize objects.
  Even if the context was created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_mk_optimize', OPTIMIZE, (_in(CONTEXT), ))
*/
//go:linkname MkOptimize C.Z3_mk_optimize
func MkOptimize(c Context) Optimize

/**
  \brief Increment the reference counter of the given optimize context

  def_API('Z3_optimize_inc_ref', VOID, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeIncRef C.Z3_optimize_inc_ref
func OptimizeIncRef(c Context, d Optimize)

/**
  \brief Decrement the reference counter of the given optimize context.

  def_API('Z3_optimize_dec_ref', VOID, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeDecRef C.Z3_optimize_dec_ref
func OptimizeDecRef(c Context, d Optimize)

/**
  \brief Assert hard constraint to the optimization context.

  \sa Z3_optimize_assert_soft
  \sa Z3_optimize_assert_and_track

  def_API('Z3_optimize_assert', VOID, (_in(CONTEXT), _in(OPTIMIZE), _in(AST)))
*/
//go:linkname OptimizeAssert C.Z3_optimize_assert
func OptimizeAssert(c Context, o Optimize, a Ast)

/**
  \brief Assert tracked hard constraint to the optimization context.

  \sa Z3_optimize_assert
  \sa Z3_optimize_assert_soft

  def_API('Z3_optimize_assert_and_track', VOID, (_in(CONTEXT), _in(OPTIMIZE), _in(AST), _in(AST)))
*/
//go:linkname OptimizeAssertAndTrack C.Z3_optimize_assert_and_track
func OptimizeAssertAndTrack(c Context, o Optimize, a Ast, t Ast)

/**
  \brief Assert soft constraint to the optimization context.
  \param c - context
  \param o - optimization context
  \param a - formula
  \param weight - a penalty for violating soft constraint. Negative weights convert into rewards.
  \param id - optional identifier to group soft constraints

  \sa Z3_optimize_assert
  \sa Z3_optimize_assert_and_track

  def_API('Z3_optimize_assert_soft', UINT, (_in(CONTEXT), _in(OPTIMIZE), _in(AST), _in(STRING), _in(SYMBOL)))
*/
//go:linkname OptimizeAssertSoft C.Z3_optimize_assert_soft
func OptimizeAssertSoft(c Context, o Optimize, a Ast, weight String, id Symbol) c.Uint

/**
  \brief Add a maximization constraint.
  \param c - context
  \param o - optimization context
  \param t - arithmetical term

  \sa Z3_optimize_minimize

  def_API('Z3_optimize_maximize', UINT, (_in(CONTEXT), _in(OPTIMIZE), _in(AST)))
*/
//go:linkname OptimizeMaximize C.Z3_optimize_maximize
func OptimizeMaximize(c Context, o Optimize, t Ast) c.Uint

/**
  \brief Add a minimization constraint.
  \param c - context
  \param o - optimization context
  \param t - arithmetical term

  \sa Z3_optimize_maximize

  def_API('Z3_optimize_minimize', UINT, (_in(CONTEXT), _in(OPTIMIZE), _in(AST)))
*/
//go:linkname OptimizeMinimize C.Z3_optimize_minimize
func OptimizeMinimize(c Context, o Optimize, t Ast) c.Uint

/**
  \brief Create a backtracking point.

  The optimize solver contains a set of rules, added facts and assertions.
  The set of rules, facts and assertions are restored upon calling #Z3_optimize_pop.

  \sa Z3_optimize_pop

  def_API('Z3_optimize_push', VOID, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizePush C.Z3_optimize_push
func OptimizePush(c Context, d Optimize)

/**
  \brief Backtrack one level.

  \sa Z3_optimize_push

  \pre The number of calls to pop cannot exceed calls to push.

  def_API('Z3_optimize_pop', VOID, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizePop C.Z3_optimize_pop
func OptimizePop(c Context, d Optimize)

/**
  \brief Check consistency and produce optimal values.
  \param c - context
  \param o - optimization context
  \param num_assumptions - number of additional assumptions
  \param assumptions - the additional assumptions

  \sa Z3_optimize_get_reason_unknown
  \sa Z3_optimize_get_model
  \sa Z3_optimize_get_statistics
  \sa Z3_optimize_get_unsat_core

  def_API('Z3_optimize_check', LBOOL, (_in(CONTEXT), _in(OPTIMIZE), _in(UINT), _in_array(2, AST)))
*/
//go:linkname OptimizeCheck C.Z3_optimize_check
func OptimizeCheck(c Context, o Optimize, num_assumptions c.Uint, assumptions *Ast) Lbool

/**
  \brief Retrieve a string that describes the last status returned by #Z3_optimize_check.

  Use this method when #Z3_optimize_check returns \c Z3_L_UNDEF.

  def_API('Z3_optimize_get_reason_unknown', STRING, (_in(CONTEXT), _in(OPTIMIZE) ))
*/
//go:linkname OptimizeGetReasonUnknown C.Z3_optimize_get_reason_unknown
func OptimizeGetReasonUnknown(c Context, d Optimize) String

/**
  \brief Retrieve the model for the last #Z3_optimize_check

  The error handler is invoked if a model is not available because
  the commands above were not invoked for the given optimization
  solver, or if the result was \c Z3_L_FALSE.

  def_API('Z3_optimize_get_model', MODEL, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetModel C.Z3_optimize_get_model
func OptimizeGetModel(c Context, o Optimize) Model

/**
  \brief Retrieve the unsat core for the last #Z3_optimize_check
  The unsat core is a subset of the assumptions \c a.

  def_API('Z3_optimize_get_unsat_core', AST_VECTOR, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetUnsatCore C.Z3_optimize_get_unsat_core
func OptimizeGetUnsatCore(c Context, o Optimize) AstVector

/**
  \brief Set parameters on optimization context.

  \param c - context
  \param o - optimization context
  \param p - parameters

  \sa Z3_optimize_get_help
  \sa Z3_optimize_get_param_descrs

  def_API('Z3_optimize_set_params', VOID, (_in(CONTEXT), _in(OPTIMIZE), _in(PARAMS)))
*/
//go:linkname OptimizeSetParams C.Z3_optimize_set_params
func OptimizeSetParams(c Context, o Optimize, p Params)

/**
  \brief Return the parameter description set for the given optimize object.

  \param c - context
  \param o - optimization context

  \sa Z3_optimize_get_help
  \sa Z3_optimize_set_params

  def_API('Z3_optimize_get_param_descrs', PARAM_DESCRS, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetParamDescrs C.Z3_optimize_get_param_descrs
func OptimizeGetParamDescrs(c Context, o Optimize) ParamDescrs

/**
  \brief Retrieve lower bound value or approximation for the i'th optimization objective.

  \param c - context
  \param o - optimization context
  \param idx - index of optimization objective

  \sa Z3_optimize_get_upper
  \sa Z3_optimize_get_lower_as_vector
  \sa Z3_optimize_get_upper_as_vector

  def_API('Z3_optimize_get_lower', AST, (_in(CONTEXT), _in(OPTIMIZE), _in(UINT)))
*/
//go:linkname OptimizeGetLower C.Z3_optimize_get_lower
func OptimizeGetLower(c Context, o Optimize, idx c.Uint) Ast

/**
  \brief Retrieve upper bound value or approximation for the i'th optimization objective.

  \param c - context
  \param o - optimization context
  \param idx - index of optimization objective

  \sa Z3_optimize_get_lower
  \sa Z3_optimize_get_lower_as_vector
  \sa Z3_optimize_get_upper_as_vector

  def_API('Z3_optimize_get_upper', AST, (_in(CONTEXT), _in(OPTIMIZE), _in(UINT)))
*/
//go:linkname OptimizeGetUpper C.Z3_optimize_get_upper
func OptimizeGetUpper(c Context, o Optimize, idx c.Uint) Ast

/**
  \brief Retrieve lower bound value or approximation for the i'th optimization objective.
         The returned vector is of length 3. It always contains numerals.
         The three numerals are coefficients \c a, \c b, \c c and encode the result of
         #Z3_optimize_get_lower \ccode{a * infinity + b + c * epsilon}.

  \param c - context
  \param o - optimization context
  \param idx - index of optimization objective

  \sa Z3_optimize_get_lower
  \sa Z3_optimize_get_upper
  \sa Z3_optimize_get_upper_as_vector

  def_API('Z3_optimize_get_lower_as_vector', AST_VECTOR, (_in(CONTEXT), _in(OPTIMIZE), _in(UINT)))
*/
//go:linkname OptimizeGetLowerAsVector C.Z3_optimize_get_lower_as_vector
func OptimizeGetLowerAsVector(c Context, o Optimize, idx c.Uint) AstVector

/**
  \brief Retrieve upper bound value or approximation for the i'th optimization objective.

  \param c - context
  \param o - optimization context
  \param idx - index of optimization objective

  \sa Z3_optimize_get_lower
  \sa Z3_optimize_get_upper
  \sa Z3_optimize_get_lower_as_vector

  def_API('Z3_optimize_get_upper_as_vector', AST_VECTOR, (_in(CONTEXT), _in(OPTIMIZE), _in(UINT)))
*/
//go:linkname OptimizeGetUpperAsVector C.Z3_optimize_get_upper_as_vector
func OptimizeGetUpperAsVector(c Context, o Optimize, idx c.Uint) AstVector

/**
  \brief Print the current context as a string.
  \param c - context.
  \param o - optimization context.

  \sa Z3_optimize_from_file
  \sa Z3_optimize_from_string

  def_API('Z3_optimize_to_string', STRING, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeToString C.Z3_optimize_to_string
func OptimizeToString(c Context, o Optimize) String

/**
  \brief Parse an SMT-LIB2 string with assertions,
  soft constraints and optimization objectives.
  Add the parsed constraints and objectives to the optimization context.

  \param c - context.
  \param o - optimize context.
  \param s - string containing SMT2 specification.

  \sa Z3_optimize_from_file
  \sa Z3_optimize_to_string

  def_API('Z3_optimize_from_string', VOID, (_in(CONTEXT), _in(OPTIMIZE), _in(STRING)))
*/
//go:linkname OptimizeFromString C.Z3_optimize_from_string
func OptimizeFromString(c Context, o Optimize, s String)

/**
  \brief Parse an SMT-LIB2 file with assertions,
  soft constraints and optimization objectives.
  Add the parsed constraints and objectives to the optimization context.

  \param c - context.
  \param o - optimize context.
  \param s - path to file containing SMT2 specification.

  \sa Z3_optimize_from_string
  \sa Z3_optimize_to_string

  def_API('Z3_optimize_from_file', VOID, (_in(CONTEXT), _in(OPTIMIZE), _in(STRING)))
*/
//go:linkname OptimizeFromFile C.Z3_optimize_from_file
func OptimizeFromFile(c Context, o Optimize, s String)

/**
  \brief Return a string containing a description of parameters accepted by optimize.

  \sa Z3_optimize_get_param_descrs
  \sa Z3_optimize_set_params

  def_API('Z3_optimize_get_help', STRING, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetHelp C.Z3_optimize_get_help
func OptimizeGetHelp(c Context, t Optimize) String

/**
  \brief Retrieve statistics information from the last call to #Z3_optimize_check

  def_API('Z3_optimize_get_statistics', STATS, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetStatistics C.Z3_optimize_get_statistics
func OptimizeGetStatistics(c Context, d Optimize) Stats

/**
  \brief Return the set of asserted formulas on the optimization context.

  def_API('Z3_optimize_get_assertions', AST_VECTOR, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetAssertions C.Z3_optimize_get_assertions
func OptimizeGetAssertions(c Context, o Optimize) AstVector

/**
  \brief Return objectives on the optimization context.
  If the objective function is a max-sat objective it is returned
  as a Pseudo-Boolean (minimization) sum of the form \ccode{(+ (if f1 w1 0) (if f2 w2 0) ...)}
  If the objective function is entered as a maximization objective, then return
  the corresponding minimization objective. In this way the resulting objective
  function is always returned as a minimization objective.

  def_API('Z3_optimize_get_objectives', AST_VECTOR, (_in(CONTEXT), _in(OPTIMIZE)))
*/
//go:linkname OptimizeGetObjectives C.Z3_optimize_get_objectives
func OptimizeGetObjectives(c Context, o Optimize) AstVector

/**
  \brief register a model event handler for new models.
*/
//go:linkname OptimizeRegisterModelEh C.Z3_optimize_register_model_eh
func OptimizeRegisterModelEh(c Context, o Optimize, m Model, ctx unsafe.Pointer, model_eh ModelEh)
