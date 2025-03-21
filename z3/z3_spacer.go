package z3

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

/** @name Spacer facilities */
/**@{*/
/**
  \brief Pose a query against the asserted rules at the given level.

  \code
     query ::= (exists (bound-vars) query)
           |  literals
  \endcode

  query returns
  - \c Z3_L_FALSE if the query is unsatisfiable.
  - \c Z3_L_TRUE if the query is satisfiable. Obtain the answer by calling #Z3_fixedpoint_get_answer.
  - \c Z3_L_UNDEF if the query was interrupted, timed out or otherwise failed.

  def_API('Z3_fixedpoint_query_from_lvl', LBOOL, (_in(CONTEXT), _in(FIXEDPOINT), _in(AST), _in(UINT)))
*/
//go:linkname FixedpointQueryFromLvl C.Z3_fixedpoint_query_from_lvl
func FixedpointQueryFromLvl(c Context, d Fixedpoint, query Ast, lvl c.Uint) Lbool

/**
  \brief Retrieve a bottom-up (from query) sequence of ground facts

  The previous call to #Z3_fixedpoint_query must have returned \c Z3_L_TRUE.

  def_API('Z3_fixedpoint_get_ground_sat_answer', AST, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetGroundSatAnswer C.Z3_fixedpoint_get_ground_sat_answer
func FixedpointGetGroundSatAnswer(c Context, d Fixedpoint) Ast

/**
  \brief Obtain the list of rules along the counterexample trace.

  def_API('Z3_fixedpoint_get_rules_along_trace', AST_VECTOR, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetRulesAlongTrace C.Z3_fixedpoint_get_rules_along_trace
func FixedpointGetRulesAlongTrace(c Context, d Fixedpoint) AstVector

/**
  \brief Obtain the list of rules along the counterexample trace.

  def_API('Z3_fixedpoint_get_rule_names_along_trace', SYMBOL, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetRuleNamesAlongTrace C.Z3_fixedpoint_get_rule_names_along_trace
func FixedpointGetRuleNamesAlongTrace(c Context, d Fixedpoint) Symbol

/**
  \brief Add an invariant for the predicate \c pred.
  Add an assumed invariant of predicate \c pred.

  Note: this functionality is Spacer specific.

  def_API('Z3_fixedpoint_add_invariant', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(FUNC_DECL), _in(AST)))
*/
//go:linkname FixedpointAddInvariant C.Z3_fixedpoint_add_invariant
func FixedpointAddInvariant(c Context, d Fixedpoint, pred FuncDecl, property Ast)

/**
  Retrieve reachable states of a predicate.
  Note: this functionality is Spacer specific.

  def_API('Z3_fixedpoint_get_reachable', AST, (_in(CONTEXT), _in(FIXEDPOINT), _in(FUNC_DECL)))
*/
//go:linkname FixedpointGetReachable C.Z3_fixedpoint_get_reachable
func FixedpointGetReachable(c Context, d Fixedpoint, pred FuncDecl) Ast

/**
  \brief Project variables given a model

  def_API('Z3_qe_model_project', AST, (_in(CONTEXT), _in(MODEL), _in(UINT), _in_array(2, APP), _in(AST)))
*/
//go:linkname QeModelProject C.Z3_qe_model_project
func QeModelProject(c Context, m Model, num_bounds c.Uint, bound *App, body Ast) Ast

/**
  \brief Project variables given a model

  def_API('Z3_qe_model_project_skolem', AST, (_in(CONTEXT), _in(MODEL), _in(UINT), _in_array(2, APP), _in(AST), _in(AST_MAP)))
*/
//go:linkname QeModelProjectSkolem C.Z3_qe_model_project_skolem
func QeModelProjectSkolem(c Context, m Model, num_bounds c.Uint, bound *App, body Ast, map_ AstMap) Ast

/**
  \brief Extrapolates a model of a formula

  def_API('Z3_model_extrapolate', AST, (_in(CONTEXT), _in(MODEL), _in(AST)))
*/
//go:linkname ModelExtrapolate C.Z3_model_extrapolate
func ModelExtrapolate(c Context, m Model, fml Ast) Ast

/**
  \brief Best-effort quantifier elimination

  def_API ('Z3_qe_lite', AST, (_in(CONTEXT), _in(AST_VECTOR), _in(AST)))
*/
//go:linkname QeLite C.Z3_qe_lite
func QeLite(c Context, vars AstVector, body Ast) Ast
