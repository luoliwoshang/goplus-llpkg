package z3

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

/** @name Fixedpoint facilities */
/**@{*/
/**
  \brief Create a new fixedpoint context.

  \remark User must use #Z3_fixedpoint_inc_ref and #Z3_fixedpoint_dec_ref to manage fixedpoint objects.
  Even if the context was created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_mk_fixedpoint', FIXEDPOINT, (_in(CONTEXT), ))
*/
//go:linkname MkFixedpoint C.Z3_mk_fixedpoint
func MkFixedpoint(c Context) Fixedpoint

/**
  \brief Increment the reference counter of the given fixedpoint context

  def_API('Z3_fixedpoint_inc_ref', VOID, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointIncRef C.Z3_fixedpoint_inc_ref
func FixedpointIncRef(c Context, d Fixedpoint)

/**
  \brief Decrement the reference counter of the given fixedpoint context.

  def_API('Z3_fixedpoint_dec_ref', VOID, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointDecRef C.Z3_fixedpoint_dec_ref
func FixedpointDecRef(c Context, d Fixedpoint)

/**
  \brief Add a universal Horn clause as a named rule.
  The \c horn_rule should be of the form:

  \code
      horn_rule ::= (forall (bound-vars) horn_rule)
                 |  (=> atoms horn_rule)
                 |  atom
  \endcode

  def_API('Z3_fixedpoint_add_rule', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(AST), _in(SYMBOL)))
*/
//go:linkname FixedpointAddRule C.Z3_fixedpoint_add_rule
func FixedpointAddRule(c Context, d Fixedpoint, rule Ast, name Symbol)

/**
   \brief Add a Database fact.

   \param c - context
   \param d - fixed point context
   \param r - relation signature for the row.
   \param num_args - number of columns for the given row.
   \param args - array of the row elements.

   The number of arguments \c num_args should be equal to the number
   of sorts in the domain of \c r. Each sort in the domain should be an integral
  (bit-vector, Boolean or or finite domain sort).

   The call has the same effect as adding a rule where \c r is applied to the arguments.

   def_API('Z3_fixedpoint_add_fact', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(FUNC_DECL), _in(UINT), _in_array(3, UINT)))
*/
//go:linkname FixedpointAddFact C.Z3_fixedpoint_add_fact
func FixedpointAddFact(c Context, d Fixedpoint, r FuncDecl, num_args c.Uint, args *c.Uint)

/**
  \brief Assert a constraint to the fixedpoint context.

  The constraints are used as background axioms when the fixedpoint engine uses the PDR mode.
  They are ignored for standard Datalog mode.

  def_API('Z3_fixedpoint_assert', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(AST)))
*/
//go:linkname FixedpointAssert C.Z3_fixedpoint_assert
func FixedpointAssert(c Context, d Fixedpoint, axiom Ast)

/**
  \brief Pose a query against the asserted rules.

  \code
     query ::= (exists (bound-vars) query)
           |  literals
  \endcode

  query returns
  - \c Z3_L_FALSE if the query is unsatisfiable.
  - \c Z3_L_TRUE if the query is satisfiable. Obtain the answer by calling #Z3_fixedpoint_get_answer.
  - \c Z3_L_UNDEF if the query was interrupted, timed out or otherwise failed.

  def_API('Z3_fixedpoint_query', LBOOL, (_in(CONTEXT), _in(FIXEDPOINT), _in(AST)))
*/
//go:linkname FixedpointQuery C.Z3_fixedpoint_query
func FixedpointQuery(c Context, d Fixedpoint, query Ast) Lbool

/**
  \brief Pose multiple queries against the asserted rules.

  The queries are encoded as relations (function declarations).

  query returns
  - \c Z3_L_FALSE if the query is unsatisfiable.
  - \c Z3_L_TRUE if the query is satisfiable. Obtain the answer by calling #Z3_fixedpoint_get_answer.
  - \c Z3_L_UNDEF if the query was interrupted, timed out or otherwise failed.

  def_API('Z3_fixedpoint_query_relations', LBOOL, (_in(CONTEXT), _in(FIXEDPOINT), _in(UINT), _in_array(2, FUNC_DECL)))
*/
//go:linkname FixedpointQueryRelations C.Z3_fixedpoint_query_relations
func FixedpointQueryRelations(c Context, d Fixedpoint, num_relations c.Uint, relations *FuncDecl) Lbool

/**
  \brief Retrieve a formula that encodes satisfying answers to the query.


  When used in Datalog mode, the returned answer is a disjunction of conjuncts.
  Each conjunct encodes values of the bound variables of the query that are satisfied.
  In PDR mode, the returned answer is a single conjunction.

  When used in Datalog mode the previous call to #Z3_fixedpoint_query must have returned \c Z3_L_TRUE.
  When used with the PDR engine, the previous call must have been either \c Z3_L_TRUE or \c Z3_L_FALSE.

  def_API('Z3_fixedpoint_get_answer', AST, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetAnswer C.Z3_fixedpoint_get_answer
func FixedpointGetAnswer(c Context, d Fixedpoint) Ast

/**
  \brief Retrieve a string that describes the last status returned by #Z3_fixedpoint_query.

  Use this method when #Z3_fixedpoint_query returns \c Z3_L_UNDEF.

  def_API('Z3_fixedpoint_get_reason_unknown', STRING, (_in(CONTEXT), _in(FIXEDPOINT) ))
*/
//go:linkname FixedpointGetReasonUnknown C.Z3_fixedpoint_get_reason_unknown
func FixedpointGetReasonUnknown(c Context, d Fixedpoint) String

/**
  \brief Update a named rule.
  A rule with the same name must have been previously created.

  def_API('Z3_fixedpoint_update_rule', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(AST), _in(SYMBOL)))
*/
//go:linkname FixedpointUpdateRule C.Z3_fixedpoint_update_rule
func FixedpointUpdateRule(c Context, d Fixedpoint, a Ast, name Symbol)

/**
  \brief Query the PDR engine for the maximal levels properties are known about predicate.

  This call retrieves the maximal number of relevant unfoldings
  of \c pred with respect to the current exploration state.
  Note: this functionality is PDR specific.

  def_API('Z3_fixedpoint_get_num_levels', UINT, (_in(CONTEXT), _in(FIXEDPOINT), _in(FUNC_DECL)))
*/
//go:linkname FixedpointGetNumLevels C.Z3_fixedpoint_get_num_levels
func FixedpointGetNumLevels(c Context, d Fixedpoint, pred FuncDecl) c.Uint

/**
  Retrieve the current cover of \c pred up to \c level unfoldings.
  Return just the delta that is known at \c level. To
  obtain the full set of properties of \c pred one should query
  at \c level+1 , \c level+2 etc, and include \c level=-1.

  Note: this functionality is PDR specific.

  def_API('Z3_fixedpoint_get_cover_delta', AST, (_in(CONTEXT), _in(FIXEDPOINT), _in(INT), _in(FUNC_DECL)))
*/
//go:linkname FixedpointGetCoverDelta C.Z3_fixedpoint_get_cover_delta
func FixedpointGetCoverDelta(c Context, d Fixedpoint, level c.Int, pred FuncDecl) Ast

/**
  \brief Add property about the predicate \c pred.
  Add a property of predicate \c pred at \c level.
  It gets pushed forward when possible.

  Note: level = -1 is treated as the fixedpoint. So passing -1 for the \c level
  means that the property is true of the fixed-point unfolding with respect to \c pred.

  Note: this functionality is PDR specific.

  def_API('Z3_fixedpoint_add_cover', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(INT), _in(FUNC_DECL), _in(AST)))
*/
//go:linkname FixedpointAddCover C.Z3_fixedpoint_add_cover
func FixedpointAddCover(c Context, d Fixedpoint, level c.Int, pred FuncDecl, property Ast)

/**
  \brief Retrieve statistics information from the last call to #Z3_fixedpoint_query.

  def_API('Z3_fixedpoint_get_statistics', STATS, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetStatistics C.Z3_fixedpoint_get_statistics
func FixedpointGetStatistics(c Context, d Fixedpoint) Stats

/**
  \brief Register relation as Fixedpoint defined.
  Fixedpoint defined relations have least-fixedpoint semantics.
  For example, the relation is empty if it does not occur
  in a head or a fact.

  def_API('Z3_fixedpoint_register_relation', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(FUNC_DECL)))
*/
//go:linkname FixedpointRegisterRelation C.Z3_fixedpoint_register_relation
func FixedpointRegisterRelation(c Context, d Fixedpoint, f FuncDecl)

/**
  \brief Configure the predicate representation.

  It sets the predicate to use a set of domains given by the list of symbols.
  The domains given by the list of symbols must belong to a set
  of built-in domains.

  def_API('Z3_fixedpoint_set_predicate_representation', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(FUNC_DECL), _in(UINT), _in_array(3, SYMBOL)))
*/
//go:linkname FixedpointSetPredicateRepresentation C.Z3_fixedpoint_set_predicate_representation
func FixedpointSetPredicateRepresentation(c Context, d Fixedpoint, f FuncDecl, num_relations c.Uint, relation_kinds *Symbol)

/**
  \brief Retrieve set of rules from fixedpoint context.

  def_API('Z3_fixedpoint_get_rules', AST_VECTOR, (_in(CONTEXT),_in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetRules C.Z3_fixedpoint_get_rules
func FixedpointGetRules(c Context, f Fixedpoint) AstVector

/**
  \brief Retrieve set of background assertions from fixedpoint context.

  def_API('Z3_fixedpoint_get_assertions', AST_VECTOR, (_in(CONTEXT),_in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetAssertions C.Z3_fixedpoint_get_assertions
func FixedpointGetAssertions(c Context, f Fixedpoint) AstVector

/**
  \brief Set parameters on fixedpoint context.

  \sa Z3_fixedpoint_get_help
  \sa Z3_fixedpoint_get_param_descrs

  def_API('Z3_fixedpoint_set_params', VOID, (_in(CONTEXT), _in(FIXEDPOINT), _in(PARAMS)))
*/
//go:linkname FixedpointSetParams C.Z3_fixedpoint_set_params
func FixedpointSetParams(c Context, f Fixedpoint, p Params)

/**
  \brief Return a string describing all fixedpoint available parameters.

  \sa Z3_fixedpoint_get_param_descrs
  \sa Z3_fixedpoint_set_params

  def_API('Z3_fixedpoint_get_help', STRING, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetHelp C.Z3_fixedpoint_get_help
func FixedpointGetHelp(c Context, f Fixedpoint) String

/**
  \brief Return the parameter description set for the given fixedpoint object.

  \sa Z3_fixedpoint_get_help
  \sa Z3_fixedpoint_set_params

  def_API('Z3_fixedpoint_get_param_descrs', PARAM_DESCRS, (_in(CONTEXT), _in(FIXEDPOINT)))
*/
//go:linkname FixedpointGetParamDescrs C.Z3_fixedpoint_get_param_descrs
func FixedpointGetParamDescrs(c Context, f Fixedpoint) ParamDescrs

/**
  \brief Print the current rules and background axioms as a string.
  \param c - context.
  \param f - fixedpoint context.
  \param num_queries - number of additional queries to print.
  \param queries - additional queries.

  \sa Z3_fixedpoint_from_file
  \sa Z3_fixedpoint_from_string

  def_API('Z3_fixedpoint_to_string', STRING, (_in(CONTEXT), _in(FIXEDPOINT), _in(UINT), _in_array(2, AST)))
*/
//go:linkname FixedpointToString C.Z3_fixedpoint_to_string
func FixedpointToString(c Context, f Fixedpoint, num_queries c.Uint, queries *Ast) String

/**
  \brief Parse an SMT-LIB2 string with fixedpoint rules.
  Add the rules to the current fixedpoint context.
  Return the set of queries in the string.

  \param c - context.
  \param f - fixedpoint context.
  \param s - string containing SMT2 specification.

  \sa Z3_fixedpoint_from_file
  \sa Z3_fixedpoint_to_string

  def_API('Z3_fixedpoint_from_string', AST_VECTOR, (_in(CONTEXT), _in(FIXEDPOINT), _in(STRING)))
*/
//go:linkname FixedpointFromString C.Z3_fixedpoint_from_string
func FixedpointFromString(c Context, f Fixedpoint, s String) AstVector

/**
  \brief Parse an SMT-LIB2 file with fixedpoint rules.
  Add the rules to the current fixedpoint context.
  Return the set of queries in the file.

  \param c - context.
  \param f - fixedpoint context.
  \param s - path to file containing SMT2 specification.

  \sa Z3_fixedpoint_from_string
  \sa Z3_fixedpoint_to_string

  def_API('Z3_fixedpoint_from_file', AST_VECTOR, (_in(CONTEXT), _in(FIXEDPOINT), _in(STRING)))
*/
//go:linkname FixedpointFromFile C.Z3_fixedpoint_from_file
func FixedpointFromFile(c Context, f Fixedpoint, s String) AstVector

// llgo:type C
type FixedpointReduceAssignCallbackFptr func(unsafe.Pointer, FuncDecl, c.Uint, *Ast, c.Uint, *Ast)

// llgo:type C
type FixedpointReduceAppCallbackFptr func(unsafe.Pointer, FuncDecl, c.Uint, *Ast, *Ast)

/** \brief Initialize the context with a user-defined state. */
//go:linkname FixedpointInit C.Z3_fixedpoint_init
func FixedpointInit(c Context, d Fixedpoint, state unsafe.Pointer)

/**
  \brief Register a callback to destructive updates.

  Registers are identified with terms encoded as fresh constants,
*/
//go:linkname FixedpointSetReduceAssignCallback C.Z3_fixedpoint_set_reduce_assign_callback
func FixedpointSetReduceAssignCallback(c Context, d Fixedpoint, cb FixedpointReduceAssignCallbackFptr)

/** \brief Register a callback for building terms based on the relational operators. */
//go:linkname FixedpointSetReduceAppCallback C.Z3_fixedpoint_set_reduce_app_callback
func FixedpointSetReduceAppCallback(c Context, d Fixedpoint, cb FixedpointReduceAppCallbackFptr)

// llgo:type C
type FixedpointNewLemmaEh func(unsafe.Pointer, Ast, c.Uint)

// llgo:type C
type FixedpointPredecessorEh func(unsafe.Pointer)

// llgo:type C
type FixedpointUnfoldEh func(unsafe.Pointer)

/** \brief set export callback for lemmas */
//go:linkname FixedpointAddCallback C.Z3_fixedpoint_add_callback
func FixedpointAddCallback(ctx Context, f Fixedpoint, state unsafe.Pointer, new_lemma_eh FixedpointNewLemmaEh, predecessor_eh FixedpointPredecessorEh, unfold_eh FixedpointUnfoldEh)

//go:linkname FixedpointAddConstraint C.Z3_fixedpoint_add_constraint
func FixedpointAddConstraint(c Context, d Fixedpoint, e Ast, lvl c.Uint)
