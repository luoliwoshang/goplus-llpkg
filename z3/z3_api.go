package z3

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_Z3Symbol struct {
	Unused [8]uint8
}
type Symbol *X_Z3Symbol

type X_Z3Config struct {
	Unused [8]uint8
}
type Config *X_Z3Config

type X_Z3Context struct {
	Unused [8]uint8
}
type Context *X_Z3Context

type X_Z3Sort struct {
	Unused [8]uint8
}
type Sort *X_Z3Sort

type X_Z3FuncDecl struct {
	Unused [8]uint8
}
type FuncDecl *X_Z3FuncDecl

type X_Z3Ast struct {
	Unused [8]uint8
}
type Ast *X_Z3Ast

type X_Z3App struct {
	Unused [8]uint8
}
type App *X_Z3App

type X_Z3Pattern struct {
	Unused [8]uint8
}
type Pattern *X_Z3Pattern

type X_Z3Model struct {
	Unused [8]uint8
}
type Model *X_Z3Model

type X_Z3Constructor struct {
	Unused [8]uint8
}
type Constructor *X_Z3Constructor

type X_Z3ConstructorList struct {
	Unused [8]uint8
}
type ConstructorList *X_Z3ConstructorList

type X_Z3Params struct {
	Unused [8]uint8
}
type Params *X_Z3Params

type X_Z3ParamDescrs struct {
	Unused [8]uint8
}
type ParamDescrs *X_Z3ParamDescrs

type X_Z3ParserContext struct {
	Unused [8]uint8
}
type ParserContext *X_Z3ParserContext

type X_Z3Goal struct {
	Unused [8]uint8
}
type Goal *X_Z3Goal

type X_Z3Tactic struct {
	Unused [8]uint8
}
type Tactic *X_Z3Tactic

type X_Z3Simplifier struct {
	Unused [8]uint8
}
type Simplifier *X_Z3Simplifier

type X_Z3Probe struct {
	Unused [8]uint8
}
type Probe *X_Z3Probe

type X_Z3Stats struct {
	Unused [8]uint8
}
type Stats *X_Z3Stats

type X_Z3Solver struct {
	Unused [8]uint8
}
type Solver *X_Z3Solver

type X_Z3SolverCallback struct {
	Unused [8]uint8
}
type SolverCallback *X_Z3SolverCallback

type X_Z3AstVector struct {
	Unused [8]uint8
}
type AstVector *X_Z3AstVector

type X_Z3AstMap struct {
	Unused [8]uint8
}
type AstMap *X_Z3AstMap

type X_Z3ApplyResult struct {
	Unused [8]uint8
}
type ApplyResult *X_Z3ApplyResult

type X_Z3FuncInterp struct {
	Unused [8]uint8
}
type FuncInterp *X_Z3FuncInterp

type X_Z3FuncEntry struct {
	Unused [8]uint8
}
type FuncEntry *X_Z3FuncEntry

type X_Z3Fixedpoint struct {
	Unused [8]uint8
}
type Fixedpoint *X_Z3Fixedpoint

type X_Z3Optimize struct {
	Unused [8]uint8
}
type Optimize *X_Z3Optimize

type X_Z3RcfNum struct {
	Unused [8]uint8
}
type RcfNum *X_Z3RcfNum
type String *int8
type CharPtr *int8
type StringPtr *String
type Lbool c.Int

const (
	LFALSE Lbool = -1
	LUNDEF Lbool = 0
	LTRUE  Lbool = 1
)

type SymbolKind c.Int

const (
	INTSYMBOL    SymbolKind = 0
	STRINGSYMBOL SymbolKind = 1
)

type ParameterKind c.Int

const (
	PARAMETERINT      ParameterKind = 0
	PARAMETERDOUBLE   ParameterKind = 1
	PARAMETERRATIONAL ParameterKind = 2
	PARAMETERSYMBOL   ParameterKind = 3
	PARAMETERSORT     ParameterKind = 4
	PARAMETERAST      ParameterKind = 5
	PARAMETERFUNCDECL ParameterKind = 6
)

type SortKind c.Int

const (
	UNINTERPRETEDSORT SortKind = 0
	BOOLSORT          SortKind = 1
	INTSORT           SortKind = 2
	REALSORT          SortKind = 3
	BVSORT            SortKind = 4
	ARRAYSORT         SortKind = 5
	DATATYPESORT      SortKind = 6
	RELATIONSORT      SortKind = 7
	FINITEDOMAINSORT  SortKind = 8
	FLOATINGPOINTSORT SortKind = 9
	ROUNDINGMODESORT  SortKind = 10
	SEQSORT           SortKind = 11
	RESORT            SortKind = 12
	CHARSORT          SortKind = 13
	TYPEVAR           SortKind = 14
	UNKNOWNSORT       SortKind = 1000
)

type AstKind c.Int

const (
	NUMERALAST    AstKind = 0
	APPAST        AstKind = 1
	VARAST        AstKind = 2
	QUANTIFIERAST AstKind = 3
	SORTAST       AstKind = 4
	FUNCDECLAST   AstKind = 5
	UNKNOWNAST    AstKind = 1000
)

type DeclKind c.Int

const (
	OPTRUE                   DeclKind = 256
	OPFALSE                  DeclKind = 257
	OPEQ                     DeclKind = 258
	OPDISTINCT               DeclKind = 259
	OPITE                    DeclKind = 260
	OPAND                    DeclKind = 261
	OPOR                     DeclKind = 262
	OPIFF                    DeclKind = 263
	OPXOR                    DeclKind = 264
	OPNOT                    DeclKind = 265
	OPIMPLIES                DeclKind = 266
	OPOEQ                    DeclKind = 267
	OPANUM                   DeclKind = 512
	OPAGNUM                  DeclKind = 513
	OPLE                     DeclKind = 514
	OPGE                     DeclKind = 515
	OPLT                     DeclKind = 516
	OPGT                     DeclKind = 517
	OPADD                    DeclKind = 518
	OPSUB                    DeclKind = 519
	OPUMINUS                 DeclKind = 520
	OPMUL                    DeclKind = 521
	OPDIV                    DeclKind = 522
	OPIDIV                   DeclKind = 523
	OPREM                    DeclKind = 524
	OPMOD                    DeclKind = 525
	OPTOREAL                 DeclKind = 526
	OPTOINT                  DeclKind = 527
	OPISINT                  DeclKind = 528
	OPPOWER                  DeclKind = 529
	OPSTORE                  DeclKind = 768
	OPSELECT                 DeclKind = 769
	OPCONSTARRAY             DeclKind = 770
	OPARRAYMAP               DeclKind = 771
	OPARRAYDEFAULT           DeclKind = 772
	OPSETUNION               DeclKind = 773
	OPSETINTERSECT           DeclKind = 774
	OPSETDIFFERENCE          DeclKind = 775
	OPSETCOMPLEMENT          DeclKind = 776
	OPSETSUBSET              DeclKind = 777
	OPASARRAY                DeclKind = 778
	OPARRAYEXT               DeclKind = 779
	OPSETHASSIZE             DeclKind = 780
	OPSETCARD                DeclKind = 781
	OPBNUM                   DeclKind = 1024
	OPBIT1                   DeclKind = 1025
	OPBIT0                   DeclKind = 1026
	OPBNEG                   DeclKind = 1027
	OPBADD                   DeclKind = 1028
	OPBSUB                   DeclKind = 1029
	OPBMUL                   DeclKind = 1030
	OPBSDIV                  DeclKind = 1031
	OPBUDIV                  DeclKind = 1032
	OPBSREM                  DeclKind = 1033
	OPBUREM                  DeclKind = 1034
	OPBSMOD                  DeclKind = 1035
	OPBSDIV0                 DeclKind = 1036
	OPBUDIV0                 DeclKind = 1037
	OPBSREM0                 DeclKind = 1038
	OPBUREM0                 DeclKind = 1039
	OPBSMOD0                 DeclKind = 1040
	OPULEQ                   DeclKind = 1041
	OPSLEQ                   DeclKind = 1042
	OPUGEQ                   DeclKind = 1043
	OPSGEQ                   DeclKind = 1044
	OPULT                    DeclKind = 1045
	OPSLT                    DeclKind = 1046
	OPUGT                    DeclKind = 1047
	OPSGT                    DeclKind = 1048
	OPBAND                   DeclKind = 1049
	OPBOR                    DeclKind = 1050
	OPBNOT                   DeclKind = 1051
	OPBXOR                   DeclKind = 1052
	OPBNAND                  DeclKind = 1053
	OPBNOR                   DeclKind = 1054
	OPBXNOR                  DeclKind = 1055
	OPCONCAT                 DeclKind = 1056
	OPSIGNEXT                DeclKind = 1057
	OPZEROEXT                DeclKind = 1058
	OPEXTRACT                DeclKind = 1059
	OPREPEAT                 DeclKind = 1060
	OPBREDOR                 DeclKind = 1061
	OPBREDAND                DeclKind = 1062
	OPBCOMP                  DeclKind = 1063
	OPBSHL                   DeclKind = 1064
	OPBLSHR                  DeclKind = 1065
	OPBASHR                  DeclKind = 1066
	OPROTATELEFT             DeclKind = 1067
	OPROTATERIGHT            DeclKind = 1068
	OPEXTROTATELEFT          DeclKind = 1069
	OPEXTROTATERIGHT         DeclKind = 1070
	OPBIT2BOOL               DeclKind = 1071
	OPINT2BV                 DeclKind = 1072
	OPBV2INT                 DeclKind = 1073
	OPCARRY                  DeclKind = 1074
	OPXOR3                   DeclKind = 1075
	OPBSMULNOOVFL            DeclKind = 1076
	OPBUMULNOOVFL            DeclKind = 1077
	OPBSMULNOUDFL            DeclKind = 1078
	OPBSDIVI                 DeclKind = 1079
	OPBUDIVI                 DeclKind = 1080
	OPBSREMI                 DeclKind = 1081
	OPBUREMI                 DeclKind = 1082
	OPBSMODI                 DeclKind = 1083
	OPPRUNDEF                DeclKind = 1280
	OPPRTRUE                 DeclKind = 1281
	OPPRASSERTED             DeclKind = 1282
	OPPRGOAL                 DeclKind = 1283
	OPPRMODUSPONENS          DeclKind = 1284
	OPPRREFLEXIVITY          DeclKind = 1285
	OPPRSYMMETRY             DeclKind = 1286
	OPPRTRANSITIVITY         DeclKind = 1287
	OPPRTRANSITIVITYSTAR     DeclKind = 1288
	OPPRMONOTONICITY         DeclKind = 1289
	OPPRQUANTINTRO           DeclKind = 1290
	OPPRBIND                 DeclKind = 1291
	OPPRDISTRIBUTIVITY       DeclKind = 1292
	OPPRANDELIM              DeclKind = 1293
	OPPRNOTORELIM            DeclKind = 1294
	OPPRREWRITE              DeclKind = 1295
	OPPRREWRITESTAR          DeclKind = 1296
	OPPRPULLQUANT            DeclKind = 1297
	OPPRPUSHQUANT            DeclKind = 1298
	OPPRELIMUNUSEDVARS       DeclKind = 1299
	OPPRDER                  DeclKind = 1300
	OPPRQUANTINST            DeclKind = 1301
	OPPRHYPOTHESIS           DeclKind = 1302
	OPPRLEMMA                DeclKind = 1303
	OPPRUNITRESOLUTION       DeclKind = 1304
	OPPRIFFTRUE              DeclKind = 1305
	OPPRIFFFALSE             DeclKind = 1306
	OPPRCOMMUTATIVITY        DeclKind = 1307
	OPPRDEFAXIOM             DeclKind = 1308
	OPPRASSUMPTIONADD        DeclKind = 1309
	OPPRLEMMAADD             DeclKind = 1310
	OPPRREDUNDANTDEL         DeclKind = 1311
	OPPRCLAUSETRAIL          DeclKind = 1312
	OPPRDEFINTRO             DeclKind = 1313
	OPPRAPPLYDEF             DeclKind = 1314
	OPPRIFFOEQ               DeclKind = 1315
	OPPRNNFPOS               DeclKind = 1316
	OPPRNNFNEG               DeclKind = 1317
	OPPRSKOLEMIZE            DeclKind = 1318
	OPPRMODUSPONENSOEQ       DeclKind = 1319
	OPPRTHLEMMA              DeclKind = 1320
	OPPRHYPERRESOLVE         DeclKind = 1321
	OPRASTORE                DeclKind = 1536
	OPRAEMPTY                DeclKind = 1537
	OPRAISEMPTY              DeclKind = 1538
	OPRAJOIN                 DeclKind = 1539
	OPRAUNION                DeclKind = 1540
	OPRAWIDEN                DeclKind = 1541
	OPRAPROJECT              DeclKind = 1542
	OPRAFILTER               DeclKind = 1543
	OPRANEGATIONFILTER       DeclKind = 1544
	OPRARENAME               DeclKind = 1545
	OPRACOMPLEMENT           DeclKind = 1546
	OPRASELECT               DeclKind = 1547
	OPRACLONE                DeclKind = 1548
	OPFDCONSTANT             DeclKind = 1549
	OPFDLT                   DeclKind = 1550
	OPSEQUNIT                DeclKind = 1551
	OPSEQEMPTY               DeclKind = 1552
	OPSEQCONCAT              DeclKind = 1553
	OPSEQPREFIX              DeclKind = 1554
	OPSEQSUFFIX              DeclKind = 1555
	OPSEQCONTAINS            DeclKind = 1556
	OPSEQEXTRACT             DeclKind = 1557
	OPSEQREPLACE             DeclKind = 1558
	OPSEQREPLACERE           DeclKind = 1559
	OPSEQREPLACEREALL        DeclKind = 1560
	OPSEQREPLACEALL          DeclKind = 1561
	OPSEQAT                  DeclKind = 1562
	OPSEQNTH                 DeclKind = 1563
	OPSEQLENGTH              DeclKind = 1564
	OPSEQINDEX               DeclKind = 1565
	OPSEQLASTINDEX           DeclKind = 1566
	OPSEQTORE                DeclKind = 1567
	OPSEQINRE                DeclKind = 1568
	OPSTRTOINT               DeclKind = 1569
	OPINTTOSTR               DeclKind = 1570
	OPUBVTOSTR               DeclKind = 1571
	OPSBVTOSTR               DeclKind = 1572
	OPSTRTOCODE              DeclKind = 1573
	OPSTRFROMCODE            DeclKind = 1574
	OPSTRINGLT               DeclKind = 1575
	OPSTRINGLE               DeclKind = 1576
	OPREPLUS                 DeclKind = 1577
	OPRESTAR                 DeclKind = 1578
	OPREOPTION               DeclKind = 1579
	OPRECONCAT               DeclKind = 1580
	OPREUNION                DeclKind = 1581
	OPRERANGE                DeclKind = 1582
	OPREDIFF                 DeclKind = 1583
	OPREINTERSECT            DeclKind = 1584
	OPRELOOP                 DeclKind = 1585
	OPREPOWER                DeclKind = 1586
	OPRECOMPLEMENT           DeclKind = 1587
	OPREEMPTYSET             DeclKind = 1588
	OPREFULLSET              DeclKind = 1589
	OPREFULLCHARSET          DeclKind = 1590
	OPREOFPRED               DeclKind = 1591
	OPREREVERSE              DeclKind = 1592
	OPREDERIVATIVE           DeclKind = 1593
	OPCHARCONST              DeclKind = 1594
	OPCHARLE                 DeclKind = 1595
	OPCHARTOINT              DeclKind = 1596
	OPCHARTOBV               DeclKind = 1597
	OPCHARFROMBV             DeclKind = 1598
	OPCHARISDIGIT            DeclKind = 1599
	OPLABEL                  DeclKind = 1792
	OPLABELLIT               DeclKind = 1793
	OPDTCONSTRUCTOR          DeclKind = 2048
	OPDTRECOGNISER           DeclKind = 2049
	OPDTIS                   DeclKind = 2050
	OPDTACCESSOR             DeclKind = 2051
	OPDTUPDATEFIELD          DeclKind = 2052
	OPPBATMOST               DeclKind = 2304
	OPPBATLEAST              DeclKind = 2305
	OPPBLE                   DeclKind = 2306
	OPPBGE                   DeclKind = 2307
	OPPBEQ                   DeclKind = 2308
	OPSPECIALRELATIONLO      DeclKind = 40960
	OPSPECIALRELATIONPO      DeclKind = 40961
	OPSPECIALRELATIONPLO     DeclKind = 40962
	OPSPECIALRELATIONTO      DeclKind = 40963
	OPSPECIALRELATIONTC      DeclKind = 40964
	OPSPECIALRELATIONTRC     DeclKind = 40965
	OPFPARMNEARESTTIESTOEVEN DeclKind = 45056
	OPFPARMNEARESTTIESTOAWAY DeclKind = 45057
	OPFPARMTOWARDPOSITIVE    DeclKind = 45058
	OPFPARMTOWARDNEGATIVE    DeclKind = 45059
	OPFPARMTOWARDZERO        DeclKind = 45060
	OPFPANUM                 DeclKind = 45061
	OPFPAPLUSINF             DeclKind = 45062
	OPFPAMINUSINF            DeclKind = 45063
	OPFPANAN                 DeclKind = 45064
	OPFPAPLUSZERO            DeclKind = 45065
	OPFPAMINUSZERO           DeclKind = 45066
	OPFPAADD                 DeclKind = 45067
	OPFPASUB                 DeclKind = 45068
	OPFPANEG                 DeclKind = 45069
	OPFPAMUL                 DeclKind = 45070
	OPFPADIV                 DeclKind = 45071
	OPFPAREM                 DeclKind = 45072
	OPFPAABS                 DeclKind = 45073
	OPFPAMIN                 DeclKind = 45074
	OPFPAMAX                 DeclKind = 45075
	OPFPAFMA                 DeclKind = 45076
	OPFPASQRT                DeclKind = 45077
	OPFPAROUNDTOINTEGRAL     DeclKind = 45078
	OPFPAEQ                  DeclKind = 45079
	OPFPALT                  DeclKind = 45080
	OPFPAGT                  DeclKind = 45081
	OPFPALE                  DeclKind = 45082
	OPFPAGE                  DeclKind = 45083
	OPFPAISNAN               DeclKind = 45084
	OPFPAISINF               DeclKind = 45085
	OPFPAISZERO              DeclKind = 45086
	OPFPAISNORMAL            DeclKind = 45087
	OPFPAISSUBNORMAL         DeclKind = 45088
	OPFPAISNEGATIVE          DeclKind = 45089
	OPFPAISPOSITIVE          DeclKind = 45090
	OPFPAFP                  DeclKind = 45091
	OPFPATOFP                DeclKind = 45092
	OPFPATOFPUNSIGNED        DeclKind = 45093
	OPFPATOUBV               DeclKind = 45094
	OPFPATOSBV               DeclKind = 45095
	OPFPATOREAL              DeclKind = 45096
	OPFPATOIEEEBV            DeclKind = 45097
	OPFPABVWRAP              DeclKind = 45098
	OPFPABV2RM               DeclKind = 45099
	OPINTERNAL               DeclKind = 45100
	OPRECURSIVE              DeclKind = 45101
	OPUNINTERPRETED          DeclKind = 45102
)

type ParamKind c.Int

const (
	PKUINT    ParamKind = 0
	PKBOOL    ParamKind = 1
	PKDOUBLE  ParamKind = 2
	PKSYMBOL  ParamKind = 3
	PKSTRING  ParamKind = 4
	PKOTHER   ParamKind = 5
	PKINVALID ParamKind = 6
)

type AstPrintMode c.Int

const (
	PRINTSMTLIBFULL       AstPrintMode = 0
	PRINTLOWLEVEL         AstPrintMode = 1
	PRINTSMTLIB2COMPLIANT AstPrintMode = 2
)

type ErrorCode c.Int

const (
	OK              ErrorCode = 0
	SORTERROR       ErrorCode = 1
	IOB             ErrorCode = 2
	INVALIDARG      ErrorCode = 3
	PARSERERROR     ErrorCode = 4
	NOPARSER        ErrorCode = 5
	INVALIDPATTERN  ErrorCode = 6
	MEMOUTFAIL      ErrorCode = 7
	FILEACCESSERROR ErrorCode = 8
	INTERNALFATAL   ErrorCode = 9
	INVALIDUSAGE    ErrorCode = 10
	DECREFERROR     ErrorCode = 11
	EXCEPTION       ErrorCode = 12
)

// llgo:type C
type ErrorHandler func(Context, ErrorCode)

// llgo:type C
type PushEh func(unsafe.Pointer, SolverCallback)

// llgo:type C
type PopEh func(unsafe.Pointer, SolverCallback, c.Uint)

// llgo:type C
type FreshEh func(unsafe.Pointer, Context) unsafe.Pointer

// llgo:type C
type FixedEh func(unsafe.Pointer, SolverCallback, Ast, Ast)

// llgo:type C
type EqEh func(unsafe.Pointer, SolverCallback, Ast, Ast)

// llgo:type C
type FinalEh func(unsafe.Pointer, SolverCallback)

// llgo:type C
type CreatedEh func(unsafe.Pointer, SolverCallback, Ast)

// llgo:type C
type DecideEh func(unsafe.Pointer, SolverCallback, Ast, c.Uint, bool)

// llgo:type C
type OnClauseEh func(unsafe.Pointer, Ast, c.Uint, *c.Uint, AstVector)
type GoalPrec c.Int

const (
	GOALPRECISE   GoalPrec = 0
	GOALUNDER     GoalPrec = 1
	GOALOVER      GoalPrec = 2
	GOALUNDEROVER GoalPrec = 3
)

/**@{*/
/**
  \brief Set a global (or module) parameter.
  This setting is shared by all Z3 contexts.

  When a Z3 module is initialized it will use the value of these parameters
  when Z3_params objects are not provided.

  The name of parameter can be composed of characters [a-z][A-Z], digits [0-9], '-' and '_'.
  The character '.' is a delimiter (more later).

  The parameter names are case-insensitive. The character '-' should be viewed as an "alias" for '_'.
  Thus, the following parameter names are considered equivalent: "pp.decimal-precision"  and "PP.DECIMAL_PRECISION".

  This function can be used to set parameters for a specific Z3 module.
  This can be done by using <module-name>.<parameter-name>.
  For example:
  Z3_global_param_set('pp.decimal', 'true')
  will set the parameter "decimal" in the module "pp" to true.

  \sa Z3_global_param_get
  \sa Z3_global_param_reset_all

  def_API('Z3_global_param_set', VOID, (_in(STRING), _in(STRING)))
*/
//go:linkname GlobalParamSet C.Z3_global_param_set
func GlobalParamSet(param_id String, param_value String)

/**
  \brief Restore the value of all global (and module) parameters.
  This command will not affect already created objects (such as tactics and solvers).

  \sa Z3_global_param_get
  \sa Z3_global_param_set

  def_API('Z3_global_param_reset_all', VOID, ())
*/
//go:linkname GlobalParamResetAll C.Z3_global_param_reset_all
func GlobalParamResetAll()

/**
  \brief Get a global (or module) parameter.

  Returns \c false if the parameter value does not exist.

  \sa Z3_global_param_reset_all
  \sa Z3_global_param_set

  \remark This function cannot be invoked simultaneously from different threads without synchronization.
  The result string stored in param_value is stored in shared location.

  def_API('Z3_global_param_get', BOOL, (_in(STRING), _out(STRING)))
*/
//go:linkname GlobalParamGet C.Z3_global_param_get
func GlobalParamGet(param_id String, param_value StringPtr) bool

/**
  \brief Create a configuration object for the Z3 context object.

  Configurations are created in order to assign parameters prior to creating
  contexts for Z3 interaction. For example, if the users wishes to use proof
  generation, then call:

  \ccode{Z3_set_param_value(cfg\, "proof"\, "true")}

  \remark In previous versions of Z3, the \c Z3_config was used to store
  global and module configurations. Now, we should use \c Z3_global_param_set.

  The following parameters can be set:

      - proof  (Boolean)           Enable proof generation
      - debug_ref_count (Boolean)  Enable debug support for Z3_ast reference counting
      - trace  (Boolean)           Tracing support for VCC
      - trace_file_name (String)   Trace out file for VCC traces
      - timeout (unsigned)         default timeout (in milliseconds) used for solvers
      - well_sorted_check          type checker
      - auto_config                use heuristics to automatically select solver and configure it
      - model                      model generation for solvers, this parameter can be overwritten when creating a solver
      - model_validate             validate models produced by solvers
      - unsat_core                 unsat-core generation for solvers, this parameter can be overwritten when creating a solver
      - encoding                   the string encoding used internally (must be either "unicode" - 18 bit, "bmp" - 16 bit or "ascii" - 8 bit)

  \sa Z3_set_param_value
  \sa Z3_del_config

  def_API('Z3_mk_config', CONFIG, ())
*/
//go:linkname MkConfig C.Z3_mk_config
func MkConfig() Config

/**
  \brief Delete the given configuration object.

  \sa Z3_mk_config

  def_API('Z3_del_config', VOID, (_in(CONFIG),))
*/
//go:linkname DelConfig C.Z3_del_config
func DelConfig(c Config)

/**
  \brief Set a configuration parameter.

  The following parameters can be set for

  \sa Z3_mk_config

  def_API('Z3_set_param_value', VOID, (_in(CONFIG), _in(STRING), _in(STRING)))
*/
//go:linkname SetParamValue C.Z3_set_param_value
func SetParamValue(c Config, param_id String, param_value String)

/**
  \brief Create a context using the given configuration.

  After a context is created, the configuration cannot be changed,
  although some parameters can be changed using #Z3_update_param_value.
  All main interaction with Z3 happens in the context of a \c Z3_context.

  In contrast to #Z3_mk_context_rc, the life time of \c Z3_ast objects
  are determined by the scope level of #Z3_solver_push and #Z3_solver_pop.
  In other words, a \c Z3_ast object remains valid until there is a
  call to #Z3_solver_pop that takes the current scope below the level where
  the object was created.

  Note that all other reference counted objects, including \c Z3_model,
  \c Z3_solver, \c Z3_func_interp have to be managed by the caller.
  Their reference counts are not handled by the context.

  \remark Thread safety: objects created using a given context should not be
  accessed from different threads without synchronization. In other words,
  operations on a context are not thread safe. To use Z3 from different threads
  create separate context objects. The \c Z3_translate, \c Z3_solver_translate,
  \c Z3_model_translate, \c Z3_goal_translate
  methods are exposed to allow copying state from one context to another.

  \remark
  - \c Z3_sort, \c Z3_func_decl, \c Z3_app, \c Z3_pattern are \c Z3_ast's.
  - Z3 uses hash-consing, i.e., when the same \c Z3_ast is created twice,
    Z3 will return the same pointer twice.

  \sa Z3_del_context

  def_API('Z3_mk_context', CONTEXT, (_in(CONFIG),))
*/
//go:linkname MkContext C.Z3_mk_context
func MkContext(c Config) Context

/**
  \brief Create a context using the given configuration.
  This function is similar to #Z3_mk_context. However,
  in the context returned by this function, the user
  is responsible for managing \c Z3_ast reference counters.
  Managing reference counters is a burden and error-prone,
  but allows the user to use the memory more efficiently.
  The user must invoke #Z3_inc_ref for any \c Z3_ast returned
  by Z3, and #Z3_dec_ref whenever the \c Z3_ast is not needed
  anymore. This idiom is similar to the one used in
  BDD (binary decision diagrams) packages such as CUDD.

  Remarks:

  - \c Z3_sort, \c Z3_func_decl, \c Z3_app, \c Z3_pattern are \c Z3_ast's.
  - After a context is created, the configuration cannot be changed.
  - All main interaction with Z3 happens in the context of a \c Z3_context.
  - Z3 uses hash-consing, i.e., when the same \c Z3_ast is created twice,
    Z3 will return the same pointer twice.

  def_API('Z3_mk_context_rc', CONTEXT, (_in(CONFIG),))
*/
//go:linkname MkContextRc C.Z3_mk_context_rc
func MkContextRc(c Config) Context

/**
  \brief Delete the given logical context.

  \sa Z3_mk_context

  def_API('Z3_del_context', VOID, (_in(CONTEXT),))
*/
//go:linkname DelContext C.Z3_del_context
func DelContext(c Context)

/**
  \brief Increment the reference counter of the given AST.
  The context \c c should have been created using #Z3_mk_context_rc.
  This function is a NOOP if \c c was created using #Z3_mk_context.

  def_API('Z3_inc_ref', VOID, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IncRef C.Z3_inc_ref
func IncRef(c Context, a Ast)

/**
  \brief Decrement the reference counter of the given AST.
  The context \c c should have been created using #Z3_mk_context_rc.
  This function is a NOOP if \c c was created using #Z3_mk_context.

  def_API('Z3_dec_ref', VOID, (_in(CONTEXT), _in(AST)))
*/
//go:linkname DecRef C.Z3_dec_ref
func DecRef(c Context, a Ast)

/**
  \brief Set a value of a context parameter.

  \sa Z3_global_param_set

  def_API('Z3_update_param_value', VOID, (_in(CONTEXT), _in(STRING), _in(STRING)))
*/
//go:linkname UpdateParamValue C.Z3_update_param_value
func UpdateParamValue(c Context, param_id String, param_value String)

/**
  \brief Retrieve description of global parameters.

  def_API('Z3_get_global_param_descrs', PARAM_DESCRS, (_in(CONTEXT),))
*/
//go:linkname GetGlobalParamDescrs C.Z3_get_global_param_descrs
func GetGlobalParamDescrs(c Context) ParamDescrs

/**
  \brief Interrupt the execution of a Z3 procedure.
  This procedure can be used to interrupt: solvers, simplifiers and tactics.

  def_API('Z3_interrupt', VOID, (_in(CONTEXT),))
*/
//go:linkname Interrupt C.Z3_interrupt
func Interrupt(c Context)

/**
  \brief use concurrency control for dec-ref.
  Reference counting decrements are allowed in separate threads from the context.
  If this setting is not invoked, reference counting decrements are not going to be thread safe.

  def_API('Z3_enable_concurrent_dec_ref', VOID, (_in(CONTEXT),))
*/
//go:linkname EnableConcurrentDecRef C.Z3_enable_concurrent_dec_ref
func EnableConcurrentDecRef(c Context)

/**
  \brief Create a Z3 (empty) parameter set.
  Starting at Z3 4.0, parameter sets are used to configure many components such as:
  simplifiers, tactics, solvers, etc.

  \remark Reference counting must be used to manage parameter sets, even when the \c Z3_context was
  created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_mk_params', PARAMS, (_in(CONTEXT),))
*/
//go:linkname MkParams C.Z3_mk_params
func MkParams(c Context) Params

/**
  \brief Increment the reference counter of the given parameter set.

  def_API('Z3_params_inc_ref', VOID, (_in(CONTEXT), _in(PARAMS)))
*/
//go:linkname ParamsIncRef C.Z3_params_inc_ref
func ParamsIncRef(c Context, p Params)

/**
  \brief Decrement the reference counter of the given parameter set.

  def_API('Z3_params_dec_ref', VOID, (_in(CONTEXT), _in(PARAMS)))
*/
//go:linkname ParamsDecRef C.Z3_params_dec_ref
func ParamsDecRef(c Context, p Params)

/**
  \brief Add a Boolean parameter \c k with value \c v to the parameter set \c p.

  def_API('Z3_params_set_bool', VOID, (_in(CONTEXT), _in(PARAMS), _in(SYMBOL), _in(BOOL)))
*/
//go:linkname ParamsSetBool C.Z3_params_set_bool
func ParamsSetBool(c Context, p Params, k Symbol, v bool)

/**
  \brief Add a unsigned parameter \c k with value \c v to the parameter set \c p.

  def_API('Z3_params_set_uint', VOID, (_in(CONTEXT), _in(PARAMS), _in(SYMBOL), _in(UINT)))
*/
//go:linkname ParamsSetUint C.Z3_params_set_uint
func ParamsSetUint(c Context, p Params, k Symbol, v c.Uint)

/**
  \brief Add a double parameter \c k with value \c v to the parameter set \c p.

  def_API('Z3_params_set_double', VOID, (_in(CONTEXT), _in(PARAMS), _in(SYMBOL), _in(DOUBLE)))
*/
//go:linkname ParamsSetDouble C.Z3_params_set_double
func ParamsSetDouble(c Context, p Params, k Symbol, v float64)

/**
  \brief Add a symbol parameter \c k with value \c v to the parameter set \c p.

  def_API('Z3_params_set_symbol', VOID, (_in(CONTEXT), _in(PARAMS), _in(SYMBOL), _in(SYMBOL)))
*/
//go:linkname ParamsSetSymbol C.Z3_params_set_symbol
func ParamsSetSymbol(c Context, p Params, k Symbol, v Symbol)

/**
  \brief Convert a parameter set into a string. This function is mainly used for printing the
  contents of a parameter set.

  def_API('Z3_params_to_string', STRING, (_in(CONTEXT), _in(PARAMS)))
*/
//go:linkname ParamsToString C.Z3_params_to_string
func ParamsToString(c Context, p Params) String

/**
  \brief Validate the parameter set \c p against the parameter description set \c d.

  The procedure invokes the error handler if \c p is invalid.

  def_API('Z3_params_validate', VOID, (_in(CONTEXT), _in(PARAMS), _in(PARAM_DESCRS)))
*/
//go:linkname ParamsValidate C.Z3_params_validate
func ParamsValidate(c Context, p Params, d ParamDescrs)

/**
  \brief Increment the reference counter of the given parameter description set.

  def_API('Z3_param_descrs_inc_ref', VOID, (_in(CONTEXT), _in(PARAM_DESCRS)))
*/
//go:linkname ParamDescrsIncRef C.Z3_param_descrs_inc_ref
func ParamDescrsIncRef(c Context, p ParamDescrs)

/**
  \brief Decrement the reference counter of the given parameter description set.

  def_API('Z3_param_descrs_dec_ref', VOID, (_in(CONTEXT), _in(PARAM_DESCRS)))
*/
//go:linkname ParamDescrsDecRef C.Z3_param_descrs_dec_ref
func ParamDescrsDecRef(c Context, p ParamDescrs)

/**
  \brief Return the kind associated with the given parameter name \c n.

  def_API('Z3_param_descrs_get_kind', UINT, (_in(CONTEXT), _in(PARAM_DESCRS), _in(SYMBOL)))
*/
//go:linkname ParamDescrsGetKind C.Z3_param_descrs_get_kind
func ParamDescrsGetKind(c Context, p ParamDescrs, n Symbol) ParamKind

/**
  \brief Return the number of parameters in the given parameter description set.

  def_API('Z3_param_descrs_size', UINT, (_in(CONTEXT), _in(PARAM_DESCRS)))
*/
//go:linkname ParamDescrsSize C.Z3_param_descrs_size
func ParamDescrsSize(c Context, p ParamDescrs) c.Uint

/**
  \brief Return the name of the parameter at given index \c i.

  \pre i < Z3_param_descrs_size(c, p)

  def_API('Z3_param_descrs_get_name', SYMBOL, (_in(CONTEXT), _in(PARAM_DESCRS), _in(UINT)))
*/
//go:linkname ParamDescrsGetName C.Z3_param_descrs_get_name
func ParamDescrsGetName(c Context, p ParamDescrs, i c.Uint) Symbol

/**
  \brief Retrieve documentation string corresponding to parameter name \c s.

  def_API('Z3_param_descrs_get_documentation', STRING, (_in(CONTEXT), _in(PARAM_DESCRS), _in(SYMBOL)))
*/
//go:linkname ParamDescrsGetDocumentation C.Z3_param_descrs_get_documentation
func ParamDescrsGetDocumentation(c Context, p ParamDescrs, s Symbol) String

/**
  \brief Convert a parameter description set into a string. This function is mainly used for printing the
  contents of a parameter description set.

  def_API('Z3_param_descrs_to_string', STRING, (_in(CONTEXT), _in(PARAM_DESCRS)))
*/
//go:linkname ParamDescrsToString C.Z3_param_descrs_to_string
func ParamDescrsToString(c Context, p ParamDescrs) String

/**
  \brief Create a Z3 symbol using an integer.

  Symbols are used to name several term and type constructors.

  NB. Not all integers can be passed to this function.
  The legal range of unsigned integers is 0 to 2^30-1.

  \sa Z3_get_symbol_int
  \sa Z3_mk_string_symbol

  def_API('Z3_mk_int_symbol', SYMBOL, (_in(CONTEXT), _in(INT)))
*/
//go:linkname MkIntSymbol C.Z3_mk_int_symbol
func MkIntSymbol(c Context, i c.Int) Symbol

/**
  \brief Create a Z3 symbol using a C string.

  Symbols are used to name several term and type constructors.

  \sa Z3_get_symbol_string
  \sa Z3_mk_int_symbol

  def_API('Z3_mk_string_symbol', SYMBOL, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname MkStringSymbol C.Z3_mk_string_symbol
func MkStringSymbol(c Context, s String) Symbol

/**
  \brief Create a free (uninterpreted) type using the given name (symbol).

  Two free types are considered the same iff the have the same name.

  def_API('Z3_mk_uninterpreted_sort', SORT, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname MkUninterpretedSort C.Z3_mk_uninterpreted_sort
func MkUninterpretedSort(c Context, s Symbol) Sort

/**
  \brief Create a type variable.

  Functions using type variables can be applied to instantiations that match the signature
  of the function. Assertions using type variables correspond to assertions over all possible
  instantiations.

  def_API('Z3_mk_type_variable', SORT, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname MkTypeVariable C.Z3_mk_type_variable
func MkTypeVariable(c Context, s Symbol) Sort

/**
  \brief Create the Boolean type.

  This type is used to create propositional variables and predicates.

  def_API('Z3_mk_bool_sort', SORT, (_in(CONTEXT), ))
*/
//go:linkname MkBoolSort C.Z3_mk_bool_sort
func MkBoolSort(c Context) Sort

/**
  \brief Create the integer type.

  This type is not the int type found in programming languages.
  A machine integer can be represented using bit-vectors. The function
  #Z3_mk_bv_sort creates a bit-vector type.

  \sa Z3_mk_bv_sort

  def_API('Z3_mk_int_sort', SORT, (_in(CONTEXT), ))
*/
//go:linkname MkIntSort C.Z3_mk_int_sort
func MkIntSort(c Context) Sort

/**
  \brief Create the real type.

  Note that this type is not a floating point number.

  def_API('Z3_mk_real_sort', SORT, (_in(CONTEXT), ))
*/
//go:linkname MkRealSort C.Z3_mk_real_sort
func MkRealSort(c Context) Sort

/**
  \brief Create a bit-vector type of the given size.

  This type can also be seen as a machine integer.

  \remark The size of the bit-vector type must be greater than zero.

  def_API('Z3_mk_bv_sort', SORT, (_in(CONTEXT), _in(UINT)))
*/
//go:linkname MkBvSort C.Z3_mk_bv_sort
func MkBvSort(c Context, sz c.Uint) Sort

/**
  \brief Create a named finite domain sort.

  To create constants that belong to the finite domain,
  use the APIs for creating numerals and pass a numeric
  constant together with the sort returned by this call.
  The numeric constant should be between 0 and the less
  than the size of the domain.

  \sa Z3_get_finite_domain_sort_size

  def_API('Z3_mk_finite_domain_sort', SORT, (_in(CONTEXT), _in(SYMBOL), _in(UINT64)))
*/
//go:linkname MkFiniteDomainSort C.Z3_mk_finite_domain_sort
func MkFiniteDomainSort(c Context, name Symbol, size uint64) Sort

/**
  \brief Create an array type.

  We usually represent the array type as: \ccode{[domain -> range]}.
  Arrays are usually used to model the heap/memory in software verification.

  \sa Z3_mk_select
  \sa Z3_mk_store

  def_API('Z3_mk_array_sort', SORT, (_in(CONTEXT), _in(SORT), _in(SORT)))
*/
//go:linkname MkArraySort C.Z3_mk_array_sort
func MkArraySort(c Context, domain Sort, range_ Sort) Sort

/**
  \brief Create an array type with N arguments

  \sa Z3_mk_select_n
  \sa Z3_mk_store_n

  def_API('Z3_mk_array_sort_n', SORT, (_in(CONTEXT), _in(UINT), _in_array(1, SORT), _in(SORT)))
*/
//go:linkname MkArraySortN C.Z3_mk_array_sort_n
func MkArraySortN(c Context, n c.Uint, domain *Sort, range_ Sort) Sort

/**
  \brief Create a tuple type.

  A tuple with \c n fields has a constructor and \c n projections.
  This function will also declare the constructor and projection functions.

  \param c logical context
  \param mk_tuple_name name of the constructor function associated with the tuple type.
  \param num_fields number of fields in the tuple type.
  \param field_names name of the projection functions.
  \param field_sorts type of the tuple fields.
  \param mk_tuple_decl output parameter that will contain the constructor declaration.
  \param proj_decl output parameter that will contain the projection function declarations. This field must be a buffer of size \c num_fields allocated by the user.

  def_API('Z3_mk_tuple_sort', SORT, (_in(CONTEXT), _in(SYMBOL), _in(UINT), _in_array(2, SYMBOL), _in_array(2, SORT), _out(FUNC_DECL), _out_array(2, FUNC_DECL)))
*/
//go:linkname MkTupleSort C.Z3_mk_tuple_sort
func MkTupleSort(c Context, mk_tuple_name Symbol, num_fields c.Uint, field_names *Symbol, field_sorts *Sort, mk_tuple_decl *FuncDecl, proj_decl *FuncDecl) Sort

/**
  \brief Create a enumeration sort.

  An enumeration sort with \c n elements.
  This function will also declare the functions corresponding to the enumerations.

  \param c logical context
  \param name name of the enumeration sort.
  \param n number of elements in enumeration sort.
  \param enum_names names of the enumerated elements.
  \param enum_consts constants corresponding to the enumerated elements.
  \param enum_testers predicates testing if terms of the enumeration sort correspond to an enumeration.

  For example, if this function is called with three symbols A, B, C and the name S, then
  \c s is a sort whose name is S, and the function returns three terms corresponding to A, B, C in
  \c enum_consts. The array \c enum_testers has three predicates of type \ccode{(s -> Bool)}.
  The first predicate (corresponding to A) is true when applied to A, and false otherwise.
  Similarly for the other predicates.

  def_API('Z3_mk_enumeration_sort', SORT, (_in(CONTEXT), _in(SYMBOL), _in(UINT), _in_array(2, SYMBOL), _out_array(2, FUNC_DECL), _out_array(2, FUNC_DECL)))
*/
//go:linkname MkEnumerationSort C.Z3_mk_enumeration_sort
func MkEnumerationSort(c Context, name Symbol, n c.Uint, enum_names *Symbol, enum_consts *FuncDecl, enum_testers *FuncDecl) Sort

/**
  \brief Create a list sort

  A list sort over \c elem_sort
  This function declares the corresponding constructors and testers for lists.

  \param c logical context
  \param name name of the list sort.
  \param elem_sort sort of list elements.
  \param nil_decl declaration for the empty list.
  \param is_nil_decl test for the empty list.
  \param cons_decl declaration for a cons cell.
  \param is_cons_decl cons cell test.
  \param head_decl list head.
  \param tail_decl list tail.

  def_API('Z3_mk_list_sort', SORT, (_in(CONTEXT), _in(SYMBOL), _in(SORT), _out(FUNC_DECL), _out(FUNC_DECL), _out(FUNC_DECL), _out(FUNC_DECL), _out(FUNC_DECL), _out(FUNC_DECL)))
*/
//go:linkname MkListSort C.Z3_mk_list_sort
func MkListSort(c Context, name Symbol, elem_sort Sort, nil_decl *FuncDecl, is_nil_decl *FuncDecl, cons_decl *FuncDecl, is_cons_decl *FuncDecl, head_decl *FuncDecl, tail_decl *FuncDecl) Sort

/**
  \brief Create a constructor.

  \param c logical context.
  \param name constructor name.
  \param recognizer name of recognizer function.
  \param num_fields number of fields in constructor.
  \param field_names names of the constructor fields.
  \param sorts field sorts, 0 if the field sort refers to a recursive sort.
  \param sort_refs reference to datatype sort that is an argument to the constructor; if the corresponding
                   sort reference is 0, then the value in sort_refs should be an index referring to
                   one of the recursive datatypes that is declared.

  \sa Z3_del_constructor
  \sa Z3_mk_constructor_list
  \sa Z3_query_constructor

  def_API('Z3_mk_constructor', CONSTRUCTOR, (_in(CONTEXT), _in(SYMBOL), _in(SYMBOL), _in(UINT), _in_array(3, SYMBOL), _in_array(3, SORT), _in_array(3, UINT)))
*/
//go:linkname MkConstructor C.Z3_mk_constructor
func MkConstructor(c Context, name Symbol, recognizer Symbol, num_fields c.Uint, field_names *Symbol, sorts *Sort, sort_refs *c.Uint) Constructor

/**
  \brief Retrieve the number of fields of a constructor

  \param c logical context.
  \param constr constructor.

  def_API('Z3_constructor_num_fields', UINT, (_in(CONTEXT), _in(CONSTRUCTOR)))
*/
//go:linkname ConstructorNumFields C.Z3_constructor_num_fields
func ConstructorNumFields(c Context, constr Constructor) c.Uint

/**
  \brief Reclaim memory allocated to constructor.

  \param c logical context.
  \param constr constructor.

  \sa Z3_mk_constructor

  def_API('Z3_del_constructor', VOID, (_in(CONTEXT), _in(CONSTRUCTOR)))
*/
//go:linkname DelConstructor C.Z3_del_constructor
func DelConstructor(c Context, constr Constructor)

/**
  \brief Create datatype, such as lists, trees, records, enumerations or unions of records.
  The datatype may be recursive. Return the datatype sort.

  \param c logical context.
  \param name name of datatype.
  \param num_constructors number of constructors passed in.
  \param constructors array of constructor containers.

  \sa Z3_mk_constructor
  \sa Z3_mk_constructor_list
  \sa Z3_mk_datatypes

  def_API('Z3_mk_datatype', SORT, (_in(CONTEXT), _in(SYMBOL), _in(UINT), _inout_array(2, CONSTRUCTOR)))
*/
//go:linkname MkDatatype C.Z3_mk_datatype
func MkDatatype(c Context, name Symbol, num_constructors c.Uint, constructors *Constructor) Sort

/**
  \brief create a forward reference to a recursive datatype being declared.
  The forward reference can be used in a nested occurrence: the range of an array
  or as element sort of a sequence. The forward reference should only be used when
  used in an accessor for a recursive datatype that gets declared.

  Forward references can replace the use sort references, that are unsigned integers
  in the \c Z3_mk_constructor call

  def_API('Z3_mk_datatype_sort', SORT, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname MkDatatypeSort C.Z3_mk_datatype_sort
func MkDatatypeSort(c Context, name Symbol) Sort

/**
  \brief Create list of constructors.

  \param c logical context.
  \param num_constructors number of constructors in list.
  \param constructors list of constructors.

  \sa Z3_del_constructor_list
  \sa Z3_mk_constructor

  def_API('Z3_mk_constructor_list', CONSTRUCTOR_LIST, (_in(CONTEXT), _in(UINT), _in_array(1, CONSTRUCTOR)))
*/
//go:linkname MkConstructorList C.Z3_mk_constructor_list
func MkConstructorList(c Context, num_constructors c.Uint, constructors *Constructor) ConstructorList

/**
  \brief Reclaim memory allocated for constructor list.

  Each constructor inside the constructor list must be independently reclaimed using #Z3_del_constructor.

  \param c logical context.
  \param clist constructor list container.

  \sa Z3_mk_constructor_list

  def_API('Z3_del_constructor_list', VOID, (_in(CONTEXT), _in(CONSTRUCTOR_LIST)))
*/
//go:linkname DelConstructorList C.Z3_del_constructor_list
func DelConstructorList(c Context, clist ConstructorList)

/**
  \brief Create mutually recursive datatypes.

  \param c logical context.
  \param num_sorts number of datatype sorts.
  \param sort_names names of datatype sorts.
  \param sorts array of datatype sorts.
  \param constructor_lists list of constructors, one list per sort.

  \sa Z3_mk_constructor
  \sa Z3_mk_constructor_list
  \sa Z3_mk_datatype

  def_API('Z3_mk_datatypes', VOID, (_in(CONTEXT), _in(UINT), _in_array(1, SYMBOL), _out_array(1, SORT), _inout_array(1, CONSTRUCTOR_LIST)))
*/
//go:linkname MkDatatypes C.Z3_mk_datatypes
func MkDatatypes(c Context, num_sorts c.Uint, sort_names *Symbol, sorts *Sort, constructor_lists *ConstructorList)

/**
  \brief Query constructor for declared functions.

  \param c logical context.
  \param constr constructor container. The container must have been passed into a #Z3_mk_datatype call.
  \param num_fields number of accessor fields in the constructor.
  \param constructor constructor function declaration, allocated by user.
  \param tester constructor test function declaration, allocated by user.
  \param accessors array of accessor function declarations allocated by user. The array must contain num_fields elements.

  \sa Z3_mk_constructor

  def_API('Z3_query_constructor', VOID, (_in(CONTEXT), _in(CONSTRUCTOR), _in(UINT), _out(FUNC_DECL), _out(FUNC_DECL), _out_array(2, FUNC_DECL)))
*/
//go:linkname QueryConstructor C.Z3_query_constructor
func QueryConstructor(c Context, constr Constructor, num_fields c.Uint, constructor *FuncDecl, tester *FuncDecl, accessors *FuncDecl)

/**
  \brief Declare a constant or function.

  \param c logical context.
  \param s name of the constant or function.
  \param domain_size number of arguments. It is 0 when declaring a constant.
  \param domain array containing the sort of each argument. The array must contain domain_size elements. It is 0 when declaring a constant.
  \param range sort of the constant or the return sort of the function.

  After declaring a constant or function, the function
  #Z3_mk_app can be used to create a constant or function
  application.

  \sa Z3_mk_app
  \sa Z3_mk_fresh_func_decl
  \sa Z3_mk_rec_func_decl

  def_API('Z3_mk_func_decl', FUNC_DECL, (_in(CONTEXT), _in(SYMBOL), _in(UINT), _in_array(2, SORT), _in(SORT)))
*/
//go:linkname MkFuncDecl C.Z3_mk_func_decl
func MkFuncDecl(c Context, s Symbol, domain_size c.Uint, domain *Sort, range_ Sort) FuncDecl

/**
  \brief Create a constant or function application.

  \sa Z3_mk_fresh_func_decl
  \sa Z3_mk_func_decl
  \sa Z3_mk_rec_func_decl

  def_API('Z3_mk_app', AST, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT), _in_array(2, AST)))
*/
//go:linkname MkApp C.Z3_mk_app
func MkApp(c Context, d FuncDecl, num_args c.Uint, args *Ast) Ast

/**
  \brief Declare and create a constant.

  This function is a shorthand for:
  \code
  Z3_func_decl d = Z3_mk_func_decl(c, s, 0, 0, ty);
  Z3_ast n            = Z3_mk_app(c, d, 0, 0);
  \endcode

  \sa Z3_mk_app
  \sa Z3_mk_fresh_const
  \sa Z3_mk_func_decl

  def_API('Z3_mk_const', AST, (_in(CONTEXT), _in(SYMBOL), _in(SORT)))
*/
//go:linkname MkConst C.Z3_mk_const
func MkConst(c Context, s Symbol, ty Sort) Ast

/**
  \brief Declare a fresh constant or function.

  Z3 will generate an unique name for this function declaration.
  If prefix is different from \c NULL, then the name generate by Z3 will start with \c prefix.

  \remark If \c prefix is \c NULL, then it is assumed to be the empty string.

  \sa Z3_mk_func_decl

  def_API('Z3_mk_fresh_func_decl', FUNC_DECL, (_in(CONTEXT), _in(STRING), _in(UINT), _in_array(2, SORT), _in(SORT)))
*/
//go:linkname MkFreshFuncDecl C.Z3_mk_fresh_func_decl
func MkFreshFuncDecl(c Context, prefix String, domain_size c.Uint, domain *Sort, range_ Sort) FuncDecl

/**
  \brief Declare and create a fresh constant.

  This function is a shorthand for:
  \code Z3_func_decl d = Z3_mk_fresh_func_decl(c, prefix, 0, 0, ty); Z3_ast n = Z3_mk_app(c, d, 0, 0); \endcode

  \remark If \c prefix is \c NULL, then it is assumed to be the empty string.

  \sa Z3_mk_app
  \sa Z3_mk_const
  \sa Z3_mk_fresh_func_decl
  \sa Z3_mk_func_decl

  def_API('Z3_mk_fresh_const', AST, (_in(CONTEXT), _in(STRING), _in(SORT)))
*/
//go:linkname MkFreshConst C.Z3_mk_fresh_const
func MkFreshConst(c Context, prefix String, ty Sort) Ast

/**
  \brief Declare a recursive function

  \param c logical context.
  \param s name of the function.
  \param domain_size number of arguments. It should be greater than 0.
  \param domain array containing the sort of each argument. The array must contain domain_size elements.
  \param range sort of the constant or the return sort of the function.

  After declaring recursive function, it should be associated with a recursive definition #Z3_add_rec_def.
  The function #Z3_mk_app can be used to create a constant or function
  application.

  \sa Z3_add_rec_def
  \sa Z3_mk_app
  \sa Z3_mk_func_decl

  def_API('Z3_mk_rec_func_decl', FUNC_DECL, (_in(CONTEXT), _in(SYMBOL), _in(UINT), _in_array(2, SORT), _in(SORT)))
*/
//go:linkname MkRecFuncDecl C.Z3_mk_rec_func_decl
func MkRecFuncDecl(c Context, s Symbol, domain_size c.Uint, domain *Sort, range_ Sort) FuncDecl

/**
  \brief Define the body of a recursive function.

  \param c logical context.
  \param f function declaration.
  \param n number of arguments to the function
  \param args constants that are used as arguments to the recursive function in the definition.
  \param body body of the recursive function

  After declaring a recursive function or a collection of mutually recursive functions, use
  this function to provide the definition for the recursive function.

  \sa Z3_mk_rec_func_decl

  def_API('Z3_add_rec_def', VOID, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT), _in_array(2, AST), _in(AST)))
*/
//go:linkname AddRecDef C.Z3_add_rec_def
func AddRecDef(c Context, f FuncDecl, n c.Uint, args *Ast, body Ast)

/** @name Propositional Logic and Equality */
/**@{*/
/**
  \brief Create an AST node representing \c true.

  def_API('Z3_mk_true', AST, (_in(CONTEXT), ))
*/
//go:linkname MkTrue C.Z3_mk_true
func MkTrue(c Context) Ast

/**
  \brief Create an AST node representing \c false.

  def_API('Z3_mk_false', AST, (_in(CONTEXT), ))
*/
//go:linkname MkFalse C.Z3_mk_false
func MkFalse(c Context) Ast

/**
  \brief Create an AST node representing \ccode{l = r}.

  The nodes \c l and \c r must have the same type.

  def_API('Z3_mk_eq', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkEq C.Z3_mk_eq
func MkEq(c Context, l Ast, r Ast) Ast

/**
  \brief Create an AST node representing \ccode{distinct(args[0], ..., args[num_args-1])}.

  The \c distinct construct is used for declaring the arguments pairwise distinct.
  That is, \ccode{Forall 0 <= i < j < num_args. not args[i] = args[j]}.

  All arguments must have the same sort.

  \remark The number of arguments of a distinct construct must be greater than one.

  def_API('Z3_mk_distinct', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkDistinct C.Z3_mk_distinct
func MkDistinct(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Create an AST node representing \ccode{not(a)}.

  The node \c a must have Boolean sort.

  def_API('Z3_mk_not', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkNot C.Z3_mk_not
func MkNot(c Context, a Ast) Ast

/**
  \brief Create an AST node representing an if-then-else: \ccode{ite(t1, t2, t3)}.

  The node \c t1 must have Boolean sort, \c t2 and \c t3 must have the same sort.
  The sort of the new node is equal to the sort of \c t2 and \c t3.

  def_API('Z3_mk_ite', AST, (_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname MkIte C.Z3_mk_ite
func MkIte(c Context, t1 Ast, t2 Ast, t3 Ast) Ast

/**
  \brief Create an AST node representing \ccode{t1 iff t2}.

  The nodes \c t1 and \c t2 must have Boolean sort.

  def_API('Z3_mk_iff', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkIff C.Z3_mk_iff
func MkIff(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create an AST node representing \ccode{t1 implies t2}.

  The nodes \c t1 and \c t2 must have Boolean sort.

  def_API('Z3_mk_implies', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkImplies C.Z3_mk_implies
func MkImplies(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create an AST node representing \ccode{t1 xor t2}.

  The nodes \c t1 and \c t2 must have Boolean sort.

  def_API('Z3_mk_xor', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkXor C.Z3_mk_xor
func MkXor(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create an AST node representing \ccode{args[0] and ... and args[num_args-1]}.

  The array \c args must have \c num_args elements.
  All arguments must have Boolean sort.

  \remark The number of arguments must be greater than zero.

  def_API('Z3_mk_and', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkAnd C.Z3_mk_and
func MkAnd(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Create an AST node representing \ccode{args[0] or ... or args[num_args-1]}.

  The array \c args must have \c num_args elements.
  All arguments must have Boolean sort.

  \remark The number of arguments must be greater than zero.

  def_API('Z3_mk_or', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkOr C.Z3_mk_or
func MkOr(c Context, num_args c.Uint, args *Ast) Ast

/** @name Integers and Reals */
/**@{*/
/**
  \brief Create an AST node representing \ccode{args[0] + ... + args[num_args-1]}.

  The array \c args must have \c num_args elements.
  All arguments must have int or real sort.

  \remark The number of arguments must be greater than zero.

  def_API('Z3_mk_add', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkAdd C.Z3_mk_add
func MkAdd(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Create an AST node representing \ccode{args[0] * ... * args[num_args-1]}.

  The array \c args must have \c num_args elements.
  All arguments must have int or real sort.

  \remark Z3 has limited support for non-linear arithmetic.
  \remark The number of arguments must be greater than zero.

  def_API('Z3_mk_mul', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkMul C.Z3_mk_mul
func MkMul(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Create an AST node representing \ccode{args[0] - ... - args[num_args - 1]}.

  The array \c args must have \c num_args elements.
  All arguments must have int or real sort.

  \remark The number of arguments must be greater than zero.

  def_API('Z3_mk_sub', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkSub C.Z3_mk_sub
func MkSub(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Create an AST node representing \ccode{- arg}.

  The arguments must have int or real type.

  def_API('Z3_mk_unary_minus', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkUnaryMinus C.Z3_mk_unary_minus
func MkUnaryMinus(c Context, arg Ast) Ast

/**
  \brief Create an AST node representing \ccode{arg1 div arg2}.

  The arguments must either both have int type or both have real type.
  If the arguments have int type, then the result type is an int type, otherwise the
  the result type is real.

  def_API('Z3_mk_div', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkDiv C.Z3_mk_div
func MkDiv(c Context, arg1 Ast, arg2 Ast) Ast

/**
  \brief Create an AST node representing \ccode{arg1 mod arg2}.

  The arguments must have int type.

  def_API('Z3_mk_mod', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkMod C.Z3_mk_mod
func MkMod(c Context, arg1 Ast, arg2 Ast) Ast

/**
  \brief Create an AST node representing \ccode{arg1 rem arg2}.

  The arguments must have int type.

  def_API('Z3_mk_rem', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkRem C.Z3_mk_rem
func MkRem(c Context, arg1 Ast, arg2 Ast) Ast

/**
  \brief Create an AST node representing \ccode{arg1 ^ arg2}.

  The arguments must have int or real type.

  def_API('Z3_mk_power', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkPower C.Z3_mk_power
func MkPower(c Context, arg1 Ast, arg2 Ast) Ast

/**
  \brief Create less than.

  The nodes \c t1 and \c t2 must have the same sort, and must be int or real.

  def_API('Z3_mk_lt', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkLt C.Z3_mk_lt
func MkLt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create less than or equal to.

  The nodes \c t1 and \c t2 must have the same sort, and must be int or real.

  def_API('Z3_mk_le', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkLe C.Z3_mk_le
func MkLe(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create greater than.

  The nodes \c t1 and \c t2 must have the same sort, and must be int or real.

  def_API('Z3_mk_gt', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkGt C.Z3_mk_gt
func MkGt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create greater than or equal to.

  The nodes \c t1 and \c t2 must have the same sort, and must be int or real.

  def_API('Z3_mk_ge', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkGe C.Z3_mk_ge
func MkGe(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create division predicate.

  The nodes \c t1 and \c t2 must be of integer sort.
  The predicate is true when \c t1 divides \c t2. For the predicate to be part of
  linear integer arithmetic, the first argument \c t1 must be a non-zero integer.

  def_API('Z3_mk_divides', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkDivides C.Z3_mk_divides
func MkDivides(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Coerce an integer to a real.

  There is also a converse operation exposed.
  It follows the semantics prescribed by the SMT-LIB standard.

  You can take the floor of a real by
  creating an auxiliary integer constant \c k and
  and asserting \ccode{mk_int2real(k) <= t1 < mk_int2real(k)+1}.

  The node \c t1 must have sort integer.

  \sa Z3_mk_real2int
  \sa Z3_mk_is_int

  def_API('Z3_mk_int2real', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkInt2real C.Z3_mk_int2real
func MkInt2real(c Context, t1 Ast) Ast

/**
  \brief Coerce a real to an integer.

  The semantics of this function follows the SMT-LIB standard
  for the function to_int

  \sa Z3_mk_int2real
  \sa Z3_mk_is_int

  def_API('Z3_mk_real2int', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkReal2int C.Z3_mk_real2int
func MkReal2int(c Context, t1 Ast) Ast

/**
  \brief Check if a real number is an integer.

  \sa Z3_mk_int2real
  \sa Z3_mk_real2int

  def_API('Z3_mk_is_int', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkIsInt C.Z3_mk_is_int
func MkIsInt(c Context, t1 Ast) Ast

/** @name Bit-vectors */
/**@{*/
/**
  \brief Bitwise negation.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_bvnot', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkBvnot C.Z3_mk_bvnot
func MkBvnot(c Context, t1 Ast) Ast

/**
  \brief Take conjunction of bits in vector, return vector of length 1.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_bvredand', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkBvredand C.Z3_mk_bvredand
func MkBvredand(c Context, t1 Ast) Ast

/**
  \brief Take disjunction of bits in vector, return vector of length 1.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_bvredor', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkBvredor C.Z3_mk_bvredor
func MkBvredor(c Context, t1 Ast) Ast

/**
  \brief Bitwise and.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvand', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvand C.Z3_mk_bvand
func MkBvand(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Bitwise or.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvor', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvor C.Z3_mk_bvor
func MkBvor(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Bitwise exclusive-or.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvxor', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvxor C.Z3_mk_bvxor
func MkBvxor(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Bitwise nand.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvnand', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvnand C.Z3_mk_bvnand
func MkBvnand(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Bitwise nor.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvnor', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvnor C.Z3_mk_bvnor
func MkBvnor(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Bitwise xnor.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvxnor', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvxnor C.Z3_mk_bvxnor
func MkBvxnor(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Standard two's complement unary minus.

  The node \c t1 must have bit-vector sort.

  def_API('Z3_mk_bvneg', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkBvneg C.Z3_mk_bvneg
func MkBvneg(c Context, t1 Ast) Ast

/**
  \brief Standard two's complement addition.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvadd', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvadd C.Z3_mk_bvadd
func MkBvadd(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Standard two's complement subtraction.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvsub', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsub C.Z3_mk_bvsub
func MkBvsub(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Standard two's complement multiplication.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvmul', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvmul C.Z3_mk_bvmul
func MkBvmul(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Unsigned division.

  It is defined as the \c floor of \ccode{t1/t2} if \c t2 is
  different from zero. If \ccode{t2} is zero, then the result
  is undefined.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvudiv', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvudiv C.Z3_mk_bvudiv
func MkBvudiv(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed division.

  It is defined in the following way:

  - The \c floor of \ccode{t1/t2} if \c t2 is different from zero, and \ccode{t1*t2 >= 0}.

  - The \c ceiling of \ccode{t1/t2} if \c t2 is different from zero, and \ccode{t1*t2 < 0}.

  If \ccode{t2} is zero, then the result is undefined.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvsdiv', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsdiv C.Z3_mk_bvsdiv
func MkBvsdiv(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Unsigned remainder.

  It is defined as \ccode{t1 - (t1 /u t2) * t2}, where \ccode{/u} represents unsigned division.

  If \ccode{t2} is zero, then the result is undefined.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvurem', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvurem C.Z3_mk_bvurem
func MkBvurem(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed remainder (sign follows dividend).

  It is defined as \ccode{t1 - (t1 /s t2) * t2}, where \ccode{/s} represents signed division.
  The most significant bit (sign) of the result is equal to the most significant bit of \c t1.

  If \ccode{t2} is zero, then the result is undefined.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  \sa Z3_mk_bvsmod

  def_API('Z3_mk_bvsrem', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsrem C.Z3_mk_bvsrem
func MkBvsrem(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed remainder (sign follows divisor).

  If \ccode{t2} is zero, then the result is undefined.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  \sa Z3_mk_bvsrem

  def_API('Z3_mk_bvsmod', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsmod C.Z3_mk_bvsmod
func MkBvsmod(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Unsigned less than.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvult', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvult C.Z3_mk_bvult
func MkBvult(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed less than.

  It abbreviates:
  \code
   (or (and (= (extract[|m-1|:|m-1|] t1) bit1)
           (= (extract[|m-1|:|m-1|] t2) bit0))
       (and (= (extract[|m-1|:|m-1|] t1) (extract[|m-1|:|m-1|] t2))
           (bvult t1 t2)))
  \endcode

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvslt', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvslt C.Z3_mk_bvslt
func MkBvslt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Unsigned less than or equal to.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvule', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvule C.Z3_mk_bvule
func MkBvule(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed less than or equal to.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvsle', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsle C.Z3_mk_bvsle
func MkBvsle(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Unsigned greater than or equal to.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvuge', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvuge C.Z3_mk_bvuge
func MkBvuge(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed greater than or equal to.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvsge', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsge C.Z3_mk_bvsge
func MkBvsge(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Unsigned greater than.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvugt', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvugt C.Z3_mk_bvugt
func MkBvugt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Two's complement signed greater than.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvsgt', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsgt C.Z3_mk_bvsgt
func MkBvsgt(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Concatenate the given bit-vectors.

  The nodes \c t1 and \c t2 must have (possibly different) bit-vector sorts

  The result is a bit-vector of size \ccode{n1+n2}, where \c n1 (\c n2) is the size
  of \c t1 (\c t2).

  def_API('Z3_mk_concat', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkConcat C.Z3_mk_concat
func MkConcat(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Extract the bits \c high down to \c low from a bit-vector of
  size \c m to yield a new bit-vector of size \c n, where \ccode{n = high - low + 1}.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_extract', AST, (_in(CONTEXT), _in(UINT), _in(UINT), _in(AST)))
*/
//go:linkname MkExtract C.Z3_mk_extract
func MkExtract(c Context, high c.Uint, low c.Uint, t1 Ast) Ast

/**
  \brief Sign-extend of the given bit-vector to the (signed) equivalent bit-vector of
  size \ccode{m+i}, where \c m is the size of the given
  bit-vector.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_sign_ext', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkSignExt C.Z3_mk_sign_ext
func MkSignExt(c Context, i c.Uint, t1 Ast) Ast

/**
  \brief Extend the given bit-vector with zeros to the (unsigned) equivalent
  bit-vector of size \ccode{m+i}, where \c m is the size of the
  given bit-vector.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_zero_ext', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkZeroExt C.Z3_mk_zero_ext
func MkZeroExt(c Context, i c.Uint, t1 Ast) Ast

/**
  \brief Repeat the given bit-vector up length \ccode{i}.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_repeat', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkRepeat C.Z3_mk_repeat
func MkRepeat(c Context, i c.Uint, t1 Ast) Ast

/**
  \brief Extracts the bit at position \ccode{i} of a bit-vector and
  yields a boolean.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_bit2bool', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkBit2bool C.Z3_mk_bit2bool
func MkBit2bool(c Context, i c.Uint, t1 Ast) Ast

/**
  \brief Shift left.

  It is equivalent to multiplication by \ccode{2^x} where \c x is the value of the
  third argument.

  NB. The semantics of shift operations varies between environments. This
  definition does not necessarily capture directly the semantics of the
  programming language or assembly architecture you are modeling.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvshl', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvshl C.Z3_mk_bvshl
func MkBvshl(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Logical shift right.

  It is equivalent to unsigned division by \ccode{2^x} where \c x is the
  value of the third argument.

  NB. The semantics of shift operations varies between environments. This
  definition does not necessarily capture directly the semantics of the
  programming language or assembly architecture you are modeling.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvlshr', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvlshr C.Z3_mk_bvlshr
func MkBvlshr(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Arithmetic shift right.

  It is like logical shift right except that the most significant
  bits of the result always copy the most significant bit of the
  second argument.

  The semantics of shift operations varies between environments. This
  definition does not necessarily capture directly the semantics of the
  programming language or assembly architecture you are modeling.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_bvashr', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvashr C.Z3_mk_bvashr
func MkBvashr(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Rotate bits of \c t1 to the left \c i times.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_rotate_left', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkRotateLeft C.Z3_mk_rotate_left
func MkRotateLeft(c Context, i c.Uint, t1 Ast) Ast

/**
  \brief Rotate bits of \c t1 to the right \c i times.

  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_rotate_right', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkRotateRight C.Z3_mk_rotate_right
func MkRotateRight(c Context, i c.Uint, t1 Ast) Ast

/**
  \brief Rotate bits of \c t1 to the left \c t2 times.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_ext_rotate_left', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkExtRotateLeft C.Z3_mk_ext_rotate_left
func MkExtRotateLeft(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Rotate bits of \c t1 to the right \c t2 times.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.

  def_API('Z3_mk_ext_rotate_right', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkExtRotateRight C.Z3_mk_ext_rotate_right
func MkExtRotateRight(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create an \c n bit bit-vector from the integer argument \c t1.

  The resulting bit-vector has \c n bits, where the i'th bit (counting
  from 0 to \c n-1) is 1 if \c (t1 div 2^i) mod 2 is 1.

  The node \c t1 must have integer sort.

  def_API('Z3_mk_int2bv', AST, (_in(CONTEXT), _in(UINT), _in(AST)))
*/
//go:linkname MkInt2bv C.Z3_mk_int2bv
func MkInt2bv(c Context, n c.Uint, t1 Ast) Ast

/**
  \brief Create an integer from the bit-vector argument \c t1.
  If \c is_signed is false, then the bit-vector \c t1 is treated as unsigned.
  So the result is non-negative
  and in the range \ccode{[0..2^N-1]}, where N are the number of bits in \c t1.
  If \c is_signed is true, \c t1 is treated as a signed bit-vector.


  The node \c t1 must have a bit-vector sort.

  def_API('Z3_mk_bv2int', AST, (_in(CONTEXT), _in(AST), _in(BOOL)))
*/
//go:linkname MkBv2int C.Z3_mk_bv2int
func MkBv2int(c Context, t1 Ast, is_signed bool) Ast

/**
  \brief Create a predicate that checks that the bit-wise addition
  of \c t1 and \c t2 does not overflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvadd_no_overflow', AST, (_in(CONTEXT), _in(AST), _in(AST), _in(BOOL)))
*/
//go:linkname MkBvaddNoOverflow C.Z3_mk_bvadd_no_overflow
func MkBvaddNoOverflow(c Context, t1 Ast, t2 Ast, is_signed bool) Ast

/**
  \brief Create a predicate that checks that the bit-wise signed addition
  of \c t1 and \c t2 does not underflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvadd_no_underflow', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvaddNoUnderflow C.Z3_mk_bvadd_no_underflow
func MkBvaddNoUnderflow(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create a predicate that checks that the bit-wise signed subtraction
  of \c t1 and \c t2 does not overflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvsub_no_overflow', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsubNoOverflow C.Z3_mk_bvsub_no_overflow
func MkBvsubNoOverflow(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Create a predicate that checks that the bit-wise subtraction
  of \c t1 and \c t2 does not underflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvsub_no_underflow', AST, (_in(CONTEXT), _in(AST), _in(AST), _in(BOOL)))
*/
//go:linkname MkBvsubNoUnderflow C.Z3_mk_bvsub_no_underflow
func MkBvsubNoUnderflow(c Context, t1 Ast, t2 Ast, is_signed bool) Ast

/**
  \brief Create a predicate that checks that the bit-wise signed division
  of \c t1 and \c t2 does not overflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvsdiv_no_overflow', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvsdivNoOverflow C.Z3_mk_bvsdiv_no_overflow
func MkBvsdivNoOverflow(c Context, t1 Ast, t2 Ast) Ast

/**
  \brief Check that bit-wise negation does not overflow when
  \c t1 is interpreted as a signed bit-vector.

  The node \c t1 must have bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvneg_no_overflow', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkBvnegNoOverflow C.Z3_mk_bvneg_no_overflow
func MkBvnegNoOverflow(c Context, t1 Ast) Ast

/**
  \brief Create a predicate that checks that the bit-wise multiplication
  of \c t1 and \c t2 does not overflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvmul_no_overflow', AST, (_in(CONTEXT), _in(AST), _in(AST), _in(BOOL)))
*/
//go:linkname MkBvmulNoOverflow C.Z3_mk_bvmul_no_overflow
func MkBvmulNoOverflow(c Context, t1 Ast, t2 Ast, is_signed bool) Ast

/**
  \brief Create a predicate that checks that the bit-wise signed multiplication
  of \c t1 and \c t2 does not underflow.

  The nodes \c t1 and \c t2 must have the same bit-vector sort.
  The returned node is of sort Bool.

  def_API('Z3_mk_bvmul_no_underflow', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkBvmulNoUnderflow C.Z3_mk_bvmul_no_underflow
func MkBvmulNoUnderflow(c Context, t1 Ast, t2 Ast) Ast

/** @name Arrays */
/**@{*/
/**
  \brief Array read.
  The argument \c a is the array and \c i is the index of the array that gets read.

  The node \c a must have an array sort \ccode{[domain -> range]},
  and \c i must have the sort \c domain.
  The sort of the result is \c range.

  \sa Z3_mk_array_sort
  \sa Z3_mk_store

  def_API('Z3_mk_select', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSelect C.Z3_mk_select
func MkSelect(c Context, a Ast, i Ast) Ast

/**
  \brief n-ary Array read.
  The argument \c a is the array and \c idxs are the indices of the array that gets read.

  def_API('Z3_mk_select_n', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST)))

*/
//go:linkname MkSelectN C.Z3_mk_select_n
func MkSelectN(c Context, a Ast, n c.Uint, idxs *Ast) Ast

/**
  \brief Array update.

  The node \c a must have an array sort \ccode{[domain -> range]}, \c i must have sort \c domain,
  \c v must have sort range. The sort of the result is \ccode{[domain -> range]}.
  The semantics of this function is given by the theory of arrays described in the SMT-LIB
  standard. See http://smtlib.org for more details.
  The result of this function is an array that is equal to \c a (with respect to \c select)
  on all indices except for \c i, where it maps to \c v (and the \c select of \c a with
  respect to \c i may be a different value).

  \sa Z3_mk_array_sort
  \sa Z3_mk_select

  def_API('Z3_mk_store', AST, (_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname MkStore C.Z3_mk_store
func MkStore(c Context, a Ast, i Ast, v Ast) Ast

/**
  \brief n-ary Array update.

  def_API('Z3_mk_store_n', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST), _in(AST)))

*/
//go:linkname MkStoreN C.Z3_mk_store_n
func MkStoreN(c Context, a Ast, n c.Uint, idxs *Ast, v Ast) Ast

/**
  \brief Create the constant array.

  The resulting term is an array, such that a \c select on an arbitrary index
  produces the value \c v.

  \param c logical context.
  \param domain domain sort for the array.
  \param v value that the array maps to.

  def_API('Z3_mk_const_array', AST, (_in(CONTEXT), _in(SORT), _in(AST)))
*/
//go:linkname MkConstArray C.Z3_mk_const_array
func MkConstArray(c Context, domain Sort, v Ast) Ast

/**
  \brief Map f on the argument arrays.

  The \c n nodes \c args must be of array sorts \ccode{[domain_i -> range_i]}.
  The function declaration \c f must have type \ccode{ range_1 .. range_n -> range}.
  \c v must have sort range. The sort of the result is \ccode{[domain_i -> range]}.

  \sa Z3_mk_array_sort
  \sa Z3_mk_store
  \sa Z3_mk_select

  def_API('Z3_mk_map', AST, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT), _in_array(2, AST)))
*/
//go:linkname MkMap C.Z3_mk_map
func MkMap(c Context, f FuncDecl, n c.Uint, args *Ast) Ast

/**
  \brief Access the array default value.
  Produces the default range value, for arrays that can be represented as
  finite maps with a default range value.

  \param c logical context.
  \param array array value whose default range value is accessed.

  def_API('Z3_mk_array_default', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkArrayDefault C.Z3_mk_array_default
func MkArrayDefault(c Context, array Ast) Ast

/**
  \brief Create array with the same interpretation as a function.
  The array satisfies the property (f x) = (select (_ as-array f) x)
  for every argument x.

  def_API('Z3_mk_as_array', AST, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname MkAsArray C.Z3_mk_as_array
func MkAsArray(c Context, f FuncDecl) Ast

/**
  \brief Create predicate that holds if Boolean array \c set has \c k elements set to true.

  def_API('Z3_mk_set_has_size', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSetHasSize C.Z3_mk_set_has_size
func MkSetHasSize(c Context, set Ast, k Ast) Ast

/** @name Sets */
/**@{*/
/**
  \brief Create Set type.

  def_API('Z3_mk_set_sort', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkSetSort C.Z3_mk_set_sort
func MkSetSort(c Context, ty Sort) Sort

/**
  \brief Create the empty set.

  def_API('Z3_mk_empty_set', AST, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkEmptySet C.Z3_mk_empty_set
func MkEmptySet(c Context, domain Sort) Ast

/**
  \brief Create the full set.

  def_API('Z3_mk_full_set', AST, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkFullSet C.Z3_mk_full_set
func MkFullSet(c Context, domain Sort) Ast

/**
  \brief Add an element to a set.

  The first argument must be a set, the second an element.

  def_API('Z3_mk_set_add', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSetAdd C.Z3_mk_set_add
func MkSetAdd(c Context, set Ast, elem Ast) Ast

/**
  \brief Remove an element to a set.

  The first argument must be a set, the second an element.

  def_API('Z3_mk_set_del', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSetDel C.Z3_mk_set_del
func MkSetDel(c Context, set Ast, elem Ast) Ast

/**
  \brief Take the union of a list of sets.

  def_API('Z3_mk_set_union', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkSetUnion C.Z3_mk_set_union
func MkSetUnion(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Take the intersection of a list of sets.

  def_API('Z3_mk_set_intersect', AST, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkSetIntersect C.Z3_mk_set_intersect
func MkSetIntersect(c Context, num_args c.Uint, args *Ast) Ast

/**
  \brief Take the set difference between two sets.

  def_API('Z3_mk_set_difference', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSetDifference C.Z3_mk_set_difference
func MkSetDifference(c Context, arg1 Ast, arg2 Ast) Ast

/**
  \brief Take the complement of a set.

  def_API('Z3_mk_set_complement', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkSetComplement C.Z3_mk_set_complement
func MkSetComplement(c Context, arg Ast) Ast

/**
  \brief Check for set membership.

  The first argument should be an element type of the set.

  def_API('Z3_mk_set_member', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSetMember C.Z3_mk_set_member
func MkSetMember(c Context, elem Ast, set Ast) Ast

/**
  \brief Check for subsetness of sets.

  def_API('Z3_mk_set_subset', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSetSubset C.Z3_mk_set_subset
func MkSetSubset(c Context, arg1 Ast, arg2 Ast) Ast

/**
  \brief Create array extensionality index given two arrays with the same sort.
         The meaning is given by the axiom:
         (=> (= (select A (array-ext A B)) (select B (array-ext A B))) (= A B))

  def_API('Z3_mk_array_ext', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkArrayExt C.Z3_mk_array_ext
func MkArrayExt(c Context, arg1 Ast, arg2 Ast) Ast

/** @name Numerals */
/**@{*/
/**
  \brief Create a numeral of a given sort.

  \param c logical context.
  \param numeral A string representing the numeral value in decimal notation. The string may be of the form `[num]*[.[num]*][E[+|-][num]+]`.
                 If the given sort is a real, then the numeral can be a rational, that is, a string of the form `[num]* / [num]*` .
  \param ty The sort of the numeral. In the current implementation, the given sort can be an int, real, finite-domain, or bit-vectors of arbitrary size.

  \sa Z3_mk_int
  \sa Z3_mk_unsigned_int

  def_API('Z3_mk_numeral', AST, (_in(CONTEXT), _in(STRING), _in(SORT)))
*/
//go:linkname MkNumeral C.Z3_mk_numeral
func MkNumeral(c Context, numeral String, ty Sort) Ast

/**
  \brief Create a real from a fraction.

  \param c logical context.
  \param num numerator of rational.
  \param den denominator of rational.

  \pre den != 0

  \sa Z3_mk_numeral
  \sa Z3_mk_int
  \sa Z3_mk_real_int64
  \sa Z3_mk_unsigned_int

  def_API('Z3_mk_real', AST, (_in(CONTEXT), _in(INT), _in(INT)))
*/
//go:linkname MkReal C.Z3_mk_real
func MkReal(c Context, num c.Int, den c.Int) Ast

/**
  \brief Create a real from a fraction of int64.

  \sa Z3_mk_real
  def_API('Z3_mk_real_int64', AST, (_in(CONTEXT), _in(INT64), _in(INT64)))
*/
//go:linkname MkRealInt64 C.Z3_mk_real_int64
func MkRealInt64(c Context, num int64, den int64) Ast

/**
  \brief Create a numeral of an int, bit-vector, or finite-domain sort.

  This function can be used to create numerals that fit in a machine integer.
  It is slightly faster than #Z3_mk_numeral since it is not necessary to parse a string.

  \sa Z3_mk_numeral

  def_API('Z3_mk_int', AST, (_in(CONTEXT), _in(INT), _in(SORT)))
*/
//go:linkname MkInt C.Z3_mk_int
func MkInt(c Context, v c.Int, ty Sort) Ast

/**
  \brief Create a numeral of a int, bit-vector, or finite-domain sort.

  This function can be used to create numerals that fit in a machine unsigned integer.
  It is slightly faster than #Z3_mk_numeral since it is not necessary to parse a string.

  \sa Z3_mk_numeral

  def_API('Z3_mk_unsigned_int', AST, (_in(CONTEXT), _in(UINT), _in(SORT)))
*/
//go:linkname MkUnsignedInt C.Z3_mk_unsigned_int
func MkUnsignedInt(c Context, v c.Uint, ty Sort) Ast

/**
  \brief Create a numeral of a int, bit-vector, or finite-domain sort.

  This function can be used to create numerals that fit in a machine \c int64_t integer.
  It is slightly faster than #Z3_mk_numeral since it is not necessary to parse a string.

  \sa Z3_mk_numeral

  def_API('Z3_mk_int64', AST, (_in(CONTEXT), _in(INT64), _in(SORT)))
*/
//go:linkname MkInt64 C.Z3_mk_int64
func MkInt64(c Context, v int64, ty Sort) Ast

/**
  \brief Create a numeral of a int, bit-vector, or finite-domain sort.

  This function can be used to create numerals that fit in a machine \c uint64_t integer.
  It is slightly faster than #Z3_mk_numeral since it is not necessary to parse a string.

  \sa Z3_mk_numeral

  def_API('Z3_mk_unsigned_int64', AST, (_in(CONTEXT), _in(UINT64), _in(SORT)))
*/
//go:linkname MkUnsignedInt64 C.Z3_mk_unsigned_int64
func MkUnsignedInt64(c Context, v uint64, ty Sort) Ast

/**
  \brief create a bit-vector numeral from a vector of Booleans.

  \sa Z3_mk_numeral
  def_API('Z3_mk_bv_numeral', AST, (_in(CONTEXT), _in(UINT), _in_array(1, BOOL)))
*/
//go:linkname MkBvNumeral C.Z3_mk_bv_numeral
func MkBvNumeral(c Context, sz c.Uint, bits *bool) Ast

/**
  \brief Create a sequence sort out of the sort for the elements.

  def_API('Z3_mk_seq_sort', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkSeqSort C.Z3_mk_seq_sort
func MkSeqSort(c Context, s Sort) Sort

/**
  \brief Check if \c s is a sequence sort.

  def_API('Z3_is_seq_sort', BOOL, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname IsSeqSort C.Z3_is_seq_sort
func IsSeqSort(c Context, s Sort) bool

/**
  \brief Retrieve basis sort for sequence sort.

  def_API('Z3_get_seq_sort_basis', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetSeqSortBasis C.Z3_get_seq_sort_basis
func GetSeqSortBasis(c Context, s Sort) Sort

/**
  \brief Create a regular expression sort out of a sequence sort.

  def_API('Z3_mk_re_sort', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkReSort C.Z3_mk_re_sort
func MkReSort(c Context, seq Sort) Sort

/**
  \brief Check if \c s is a regular expression sort.

  def_API('Z3_is_re_sort', BOOL, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname IsReSort C.Z3_is_re_sort
func IsReSort(c Context, s Sort) bool

/**
  \brief Retrieve basis sort for regex sort.

  def_API('Z3_get_re_sort_basis', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetReSortBasis C.Z3_get_re_sort_basis
func GetReSortBasis(c Context, s Sort) Sort

/**
  \brief Create a sort for unicode strings.

  The sort for characters can be changed to ASCII by setting
  the global parameter \c encoding to \c ascii, or alternative
  to 16 bit characters by setting \c encoding to \c bmp.

  def_API('Z3_mk_string_sort', SORT, (_in(CONTEXT), ))
*/
//go:linkname MkStringSort C.Z3_mk_string_sort
func MkStringSort(c Context) Sort

/**
  \brief Create a sort for unicode characters.

  The sort for characters can be changed to ASCII by setting
  the global parameter \c encoding to \c ascii, or alternative
  to 16 bit characters by setting \c encoding to \c bmp.

  def_API('Z3_mk_char_sort', SORT, (_in(CONTEXT), ))
*/
//go:linkname MkCharSort C.Z3_mk_char_sort
func MkCharSort(c Context) Sort

/**
  \brief Check if \c s is a string sort.

  def_API('Z3_is_string_sort', BOOL, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname IsStringSort C.Z3_is_string_sort
func IsStringSort(c Context, s Sort) bool

/**
  \brief Check if \c s is a character sort.

  def_API('Z3_is_char_sort', BOOL, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname IsCharSort C.Z3_is_char_sort
func IsCharSort(c Context, s Sort) bool

/**
  \brief Create a string constant out of the string that is passed in
  The string may contain escape encoding for non-printable characters
  or characters outside of the basic printable ASCII range. For example,
  the escape encoding \\u{0} represents the character 0 and the encoding
  \\u{100} represents the character 256.

  def_API('Z3_mk_string', AST, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname MkString C.Z3_mk_string
func MkString(c Context, s String) Ast

/**
  \brief Create a string constant out of the string that is passed in
  It takes the length of the string as well to take into account
  0 characters. The string is treated as if it is unescaped so a sequence
  of characters \\u{0} is treated as 5 characters and not the character 0.

  def_API('Z3_mk_lstring', AST, (_in(CONTEXT), _in(UINT), _in(STRING)))
*/
//go:linkname MkLstring C.Z3_mk_lstring
func MkLstring(c Context, len c.Uint, s String) Ast

/**
  \brief Create a string constant out of the string that is passed in
  It takes the length of the string as well to take into account
  0 characters. The string is unescaped.

  def_API('Z3_mk_u32string', AST, (_in(CONTEXT), _in(UINT), _in_array(1, UINT)))
*/
//go:linkname MkU32string C.Z3_mk_u32string
func MkU32string(c Context, len c.Uint, chars *c.Uint) Ast

/**
  \brief Determine if \c s is a string constant.

  def_API('Z3_is_string', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsString C.Z3_is_string
func IsString(c Context, s Ast) bool

/**
  \brief Retrieve the string constant stored in \c s.
  Characters outside the basic printable ASCII range are escaped.

  \pre  Z3_is_string(c, s)

  def_API('Z3_get_string', STRING, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetString C.Z3_get_string
func GetString(c Context, s Ast) String

/**
  \brief Retrieve the string constant stored in \c s. The string can contain escape sequences.
  Characters in the range 1 to 255 are literal.
  Characters in the range 0, and 256 above are escaped.

  \pre  Z3_is_string(c, s)

  def_API('Z3_get_lstring', CHAR_PTR, (_in(CONTEXT), _in(AST), _out(UINT)))
*/
//go:linkname GetLstring C.Z3_get_lstring
func GetLstring(c Context, s Ast, length *c.Uint) CharPtr

/**
  \brief Retrieve the length of the unescaped string constant stored in \c s.

  \pre  Z3_is_string(c, s)

  def_API('Z3_get_string_length', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetStringLength C.Z3_get_string_length
func GetStringLength(c Context, s Ast) c.Uint

/**
  \brief Retrieve the unescaped string constant stored in \c s.

  \pre  Z3_is_string(c, s)

  \pre length contains the number of characters in s

  def_API('Z3_get_string_contents', VOID, (_in(CONTEXT), _in(AST), _in(UINT), _out_array(2, UINT)))
*/
//go:linkname GetStringContents C.Z3_get_string_contents
func GetStringContents(c Context, s Ast, length c.Uint, contents *c.Uint)

/**
  \brief Create an empty sequence of the sequence sort \c seq.

  \pre s is a sequence sort.

  def_API('Z3_mk_seq_empty', AST ,(_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkSeqEmpty C.Z3_mk_seq_empty
func MkSeqEmpty(c Context, seq Sort) Ast

/**
  \brief Create a unit sequence of \c a.

  def_API('Z3_mk_seq_unit', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkSeqUnit C.Z3_mk_seq_unit
func MkSeqUnit(c Context, a Ast) Ast

/**
  \brief Concatenate sequences.

  \pre n > 0

  def_API('Z3_mk_seq_concat', AST ,(_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkSeqConcat C.Z3_mk_seq_concat
func MkSeqConcat(c Context, n c.Uint, args *Ast) Ast

/**
  \brief Check if \c prefix is a prefix of \c s.

  \pre prefix and s are the same sequence sorts.

  def_API('Z3_mk_seq_prefix', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqPrefix C.Z3_mk_seq_prefix
func MkSeqPrefix(c Context, prefix Ast, s Ast) Ast

/**
  \brief Check if \c suffix is a suffix of \c s.

  \pre \c suffix and \c s are the same sequence sorts.

  def_API('Z3_mk_seq_suffix', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqSuffix C.Z3_mk_seq_suffix
func MkSeqSuffix(c Context, suffix Ast, s Ast) Ast

/**
  \brief Check if \c container contains \c containee.

  \pre \c container and \c containee are the same sequence sorts.

  def_API('Z3_mk_seq_contains', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqContains C.Z3_mk_seq_contains
func MkSeqContains(c Context, container Ast, containee Ast) Ast

/**
  \brief Check if \c s1 is lexicographically strictly less than \c s2.

  \pre \c s1 and \c s2 are strings

  def_API('Z3_mk_str_lt', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkStrLt C.Z3_mk_str_lt
func MkStrLt(c Context, prefix Ast, s Ast) Ast

/**
  \brief Check if \c s1 is equal or lexicographically strictly less than \c s2.

  \pre \c s1 and \c s2 are strings

  def_API('Z3_mk_str_le', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkStrLe C.Z3_mk_str_le
func MkStrLe(c Context, prefix Ast, s Ast) Ast

/**
  \brief Extract subsequence starting at \c offset of \c length.

  def_API('Z3_mk_seq_extract', AST ,(_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname MkSeqExtract C.Z3_mk_seq_extract
func MkSeqExtract(c Context, s Ast, offset Ast, length Ast) Ast

/**
  \brief Replace the first occurrence of \c src with \c dst in \c s.

  def_API('Z3_mk_seq_replace', AST ,(_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname MkSeqReplace C.Z3_mk_seq_replace
func MkSeqReplace(c Context, s Ast, src Ast, dst Ast) Ast

/**
  \brief Retrieve from \c s the unit sequence positioned at position \c index.
  The sequence is empty if the index is out of bounds.

  def_API('Z3_mk_seq_at', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqAt C.Z3_mk_seq_at
func MkSeqAt(c Context, s Ast, index Ast) Ast

/**
  \brief Retrieve from \c s the element positioned at position \c index.
  The function is under-specified if the index is out of bounds.

  def_API('Z3_mk_seq_nth', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqNth C.Z3_mk_seq_nth
func MkSeqNth(c Context, s Ast, index Ast) Ast

/**
  \brief Return the length of the sequence \c s.

  def_API('Z3_mk_seq_length', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkSeqLength C.Z3_mk_seq_length
func MkSeqLength(c Context, s Ast) Ast

/**
  \brief Return index of the first occurrence of \c substr in \c s starting from offset \c offset.
  If \c s does not contain \c substr, then the value is -1, if \c offset is the length of \c s, then the value is -1 as well.
  The value is -1 if \c offset is negative or larger than the length of \c s.

  def_API('Z3_mk_seq_index', AST ,(_in(CONTEXT), _in(AST), _in(AST), _in(AST)))
*/
//go:linkname MkSeqIndex C.Z3_mk_seq_index
func MkSeqIndex(c Context, s Ast, substr Ast, offset Ast) Ast

/**
  \brief Return index of the last occurrence of \c substr in \c s.
  If \c s does not contain \c substr, then the value is -1,
  def_API('Z3_mk_seq_last_index', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqLastIndex C.Z3_mk_seq_last_index
func MkSeqLastIndex(c Context, s Ast, substr Ast) Ast

/**
  \brief Convert string to integer.

  def_API('Z3_mk_str_to_int', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkStrToInt C.Z3_mk_str_to_int
func MkStrToInt(c Context, s Ast) Ast

/**
  \brief Integer to string conversion.

  def_API('Z3_mk_int_to_str', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkIntToStr C.Z3_mk_int_to_str
func MkIntToStr(c Context, s Ast) Ast

/**
  \brief String to code conversion.

  def_API('Z3_mk_string_to_code', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkStringToCode C.Z3_mk_string_to_code
func MkStringToCode(c Context, a Ast) Ast

/**
  \brief Code to string conversion.

  def_API('Z3_mk_string_from_code', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkStringFromCode C.Z3_mk_string_from_code
func MkStringFromCode(c Context, a Ast) Ast

/**
  \brief Unsigned bit-vector to string conversion.

  def_API('Z3_mk_ubv_to_str', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkUbvToStr C.Z3_mk_ubv_to_str
func MkUbvToStr(c Context, s Ast) Ast

/**
  \brief Signed bit-vector to string conversion.

  def_API('Z3_mk_sbv_to_str', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkSbvToStr C.Z3_mk_sbv_to_str
func MkSbvToStr(c Context, s Ast) Ast

/**
  \brief Create a regular expression that accepts the sequence \c seq.

  def_API('Z3_mk_seq_to_re', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkSeqToRe C.Z3_mk_seq_to_re
func MkSeqToRe(c Context, seq Ast) Ast

/**
  \brief Check if \c seq is in the language generated by the regular expression \c re.

  def_API('Z3_mk_seq_in_re', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkSeqInRe C.Z3_mk_seq_in_re
func MkSeqInRe(c Context, seq Ast, re Ast) Ast

/**
  \brief Create the regular language \c re+.

  def_API('Z3_mk_re_plus', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkRePlus C.Z3_mk_re_plus
func MkRePlus(c Context, re Ast) Ast

/**
  \brief Create the regular language \c re*.

  def_API('Z3_mk_re_star', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkReStar C.Z3_mk_re_star
func MkReStar(c Context, re Ast) Ast

/**
  \brief Create the regular language \c [re].

  def_API('Z3_mk_re_option', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkReOption C.Z3_mk_re_option
func MkReOption(c Context, re Ast) Ast

/**
  \brief Create the union of the regular languages.

  \pre n > 0

  def_API('Z3_mk_re_union', AST ,(_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkReUnion C.Z3_mk_re_union
func MkReUnion(c Context, n c.Uint, args *Ast) Ast

/**
  \brief Create the concatenation of the regular languages.

  \pre n > 0

  def_API('Z3_mk_re_concat', AST ,(_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkReConcat C.Z3_mk_re_concat
func MkReConcat(c Context, n c.Uint, args *Ast) Ast

/**
  \brief Create the range regular expression over two sequences of length 1.

  def_API('Z3_mk_re_range', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkReRange C.Z3_mk_re_range
func MkReRange(c Context, lo Ast, hi Ast) Ast

/**
   \brief Create a regular expression that accepts all singleton sequences of the regular expression sort

  def_API('Z3_mk_re_allchar', AST, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkReAllchar C.Z3_mk_re_allchar
func MkReAllchar(c Context, regex_sort Sort) Ast

/**
  \brief Create a regular expression loop. The supplied regular expression \c r is repeated
  between \c lo and \c hi times. The \c lo should be below \c hi with one exception: when
  supplying the value \c hi as 0, the meaning is to repeat the argument \c r at least
  \c lo number of times, and with an unbounded upper bound.

  def_API('Z3_mk_re_loop', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in(UINT)))
*/
//go:linkname MkReLoop C.Z3_mk_re_loop
func MkReLoop(c Context, r Ast, lo c.Uint, hi c.Uint) Ast

/**
  \brief Create a power regular expression.

  def_API('Z3_mk_re_power', AST, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname MkRePower C.Z3_mk_re_power
func MkRePower(c Context, re Ast, n c.Uint) Ast

/**
  \brief Create the intersection of the regular languages.

  \pre n > 0

  def_API('Z3_mk_re_intersect', AST ,(_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkReIntersect C.Z3_mk_re_intersect
func MkReIntersect(c Context, n c.Uint, args *Ast) Ast

/**
  \brief Create the complement of the regular language \c re.

  def_API('Z3_mk_re_complement', AST ,(_in(CONTEXT), _in(AST)))
*/
//go:linkname MkReComplement C.Z3_mk_re_complement
func MkReComplement(c Context, re Ast) Ast

/**
  \brief Create the difference of regular expressions.

  def_API('Z3_mk_re_diff', AST ,(_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkReDiff C.Z3_mk_re_diff
func MkReDiff(c Context, re1 Ast, re2 Ast) Ast

/**
  \brief Create an empty regular expression of sort \c re.

  \pre re is a regular expression sort.

  def_API('Z3_mk_re_empty', AST ,(_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkReEmpty C.Z3_mk_re_empty
func MkReEmpty(c Context, re Sort) Ast

/**
  \brief Create an universal regular expression of sort \c re.

  \pre re is a regular expression sort.

  def_API('Z3_mk_re_full', AST ,(_in(CONTEXT), _in(SORT)))
*/
//go:linkname MkReFull C.Z3_mk_re_full
func MkReFull(c Context, re Sort) Ast

/**
  \brief Create a character literal
  def_API('Z3_mk_char', AST, (_in(CONTEXT), _in(UINT)))
*/
//go:linkname MkChar C.Z3_mk_char
func MkChar(c Context, ch c.Uint) Ast

/**
  \brief Create less than or equal to between two characters.

  def_API('Z3_mk_char_le', AST, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname MkCharLe C.Z3_mk_char_le
func MkCharLe(c Context, ch1 Ast, ch2 Ast) Ast

/**
  \brief Create an integer (code point) from character.

  def_API('Z3_mk_char_to_int', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkCharToInt C.Z3_mk_char_to_int
func MkCharToInt(c Context, ch Ast) Ast

/**
  \brief Create a bit-vector (code point) from character.

  def_API('Z3_mk_char_to_bv', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkCharToBv C.Z3_mk_char_to_bv
func MkCharToBv(c Context, ch Ast) Ast

/**
  \brief Create a character from a bit-vector (code point).

  def_API('Z3_mk_char_from_bv', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkCharFromBv C.Z3_mk_char_from_bv
func MkCharFromBv(c Context, bv Ast) Ast

/**
  \brief Create a check if the character is a digit.

  def_API('Z3_mk_char_is_digit', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname MkCharIsDigit C.Z3_mk_char_is_digit
func MkCharIsDigit(c Context, ch Ast) Ast

/** @name Special relations */
/**@{*/
/**
  \brief create a linear ordering relation over signature \c a.
  The relation is identified by the index \c id.

  def_API('Z3_mk_linear_order', FUNC_DECL ,(_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname MkLinearOrder C.Z3_mk_linear_order
func MkLinearOrder(c Context, a Sort, id c.Uint) FuncDecl

/**
  \brief create a partial ordering relation over signature \c a and index \c id.

  def_API('Z3_mk_partial_order', FUNC_DECL ,(_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname MkPartialOrder C.Z3_mk_partial_order
func MkPartialOrder(c Context, a Sort, id c.Uint) FuncDecl

/**
  \brief create a piecewise linear ordering relation over signature \c a and index \c id.

  def_API('Z3_mk_piecewise_linear_order', FUNC_DECL ,(_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname MkPiecewiseLinearOrder C.Z3_mk_piecewise_linear_order
func MkPiecewiseLinearOrder(c Context, a Sort, id c.Uint) FuncDecl

/**
  \brief create a tree ordering relation over signature \c a identified using index \c id.

  def_API('Z3_mk_tree_order', FUNC_DECL, (_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname MkTreeOrder C.Z3_mk_tree_order
func MkTreeOrder(c Context, a Sort, id c.Uint) FuncDecl

/**
  \brief create transitive closure of binary relation.

  \pre f is a binary relation, such that the two arguments have the same sorts.

  The resulting relation f+ represents the transitive closure of f.

  def_API('Z3_mk_transitive_closure', FUNC_DECL ,(_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname MkTransitiveClosure C.Z3_mk_transitive_closure
func MkTransitiveClosure(c Context, f FuncDecl) FuncDecl

/** @name Quantifiers */
/**@{*/
/**
  \brief Create a pattern for quantifier instantiation.

  Z3 uses pattern matching to instantiate quantifiers. If a
  pattern is not provided for a quantifier, then Z3 will
  automatically compute a set of patterns for it. However, for
  optimal performance, the user should provide the patterns.

  Patterns comprise a list of terms. The list should be
  non-empty.  If the list comprises of more than one term, it is
  a called a multi-pattern.

  In general, one can pass in a list of (multi-)patterns in the
  quantifier constructor.

  \sa Z3_mk_forall
  \sa Z3_mk_exists

  def_API('Z3_mk_pattern', PATTERN, (_in(CONTEXT), _in(UINT), _in_array(1, AST)))
*/
//go:linkname MkPattern C.Z3_mk_pattern
func MkPattern(c Context, num_patterns c.Uint, terms *Ast) Pattern

/**
  \brief Create a variable.

  Variables are intended to be bound by a scope created by a quantifier. So we call them bound variables
  even if they appear as free variables in the expression produced by \c Z3_mk_bound.

  Bound variables are indexed by de-Bruijn indices. It is perhaps easiest to explain
  the meaning of de-Bruijn indices by indicating the compilation process from
  non-de-Bruijn formulas to de-Bruijn format.

  \verbatim
  abs(forall (x1) phi) = forall (x1) abs1(phi, x1, 0)
  abs(forall (x1, x2) phi) = abs(forall (x1) abs(forall (x2) phi))
  abs1(x, x, n) = b_n
  abs1(y, x, n) = y
  abs1(f(t1,...,tn), x, n) = f(abs1(t1,x,n), ..., abs1(tn,x,n))
  abs1(forall (x1) phi, x, n) = forall (x1) (abs1(phi, x, n+1))
  \endverbatim

  The last line is significant: the index of a bound variable is different depending
  on the scope in which it appears. The deeper x appears, the higher is its
  index.

  \param c logical context
  \param index de-Bruijn index
  \param ty sort of the bound variable

  \sa Z3_mk_forall
  \sa Z3_mk_exists

  def_API('Z3_mk_bound', AST, (_in(CONTEXT), _in(UINT), _in(SORT)))
*/
//go:linkname MkBound C.Z3_mk_bound
func MkBound(c Context, index c.Uint, ty Sort) Ast

/**
  \brief Create a forall formula. It takes an expression \c body that contains bound variables
  of the same sorts as the sorts listed in the array \c sorts. The bound variables are de-Bruijn indices created
  using #Z3_mk_bound. The array \c decl_names contains the names that the quantified formula uses for the
  bound variables. Z3 applies the convention that the last element in the \c decl_names and \c sorts array
  refers to the variable with index 0, the second to last element of \c decl_names and \c sorts refers
  to the variable with index 1, etc.

  \param c logical context.
  \param weight quantifiers are associated with weights indicating the importance of using the quantifier during instantiation. By default, pass the weight 0.
  \param num_patterns number of patterns.
  \param patterns array containing the patterns created using #Z3_mk_pattern.
  \param num_decls number of variables to be bound.
  \param sorts the sorts of the bound variables.
  \param decl_names names of the bound variables
  \param body the body of the quantifier.

  \sa Z3_mk_pattern
  \sa Z3_mk_bound
  \sa Z3_mk_exists

  def_API('Z3_mk_forall', AST, (_in(CONTEXT), _in(UINT), _in(UINT), _in_array(2, PATTERN), _in(UINT), _in_array(4, SORT), _in_array(4, SYMBOL), _in(AST)))
*/
//go:linkname MkForall C.Z3_mk_forall
func MkForall(c Context, weight c.Uint, num_patterns c.Uint, patterns *Pattern, num_decls c.Uint, sorts *Sort, decl_names *Symbol, body Ast) Ast

/**
  \brief Create an exists formula. Similar to #Z3_mk_forall.

  \sa Z3_mk_pattern
  \sa Z3_mk_bound
  \sa Z3_mk_forall
  \sa Z3_mk_quantifier

  def_API('Z3_mk_exists', AST, (_in(CONTEXT), _in(UINT), _in(UINT), _in_array(2, PATTERN), _in(UINT), _in_array(4, SORT), _in_array(4, SYMBOL), _in(AST)))
*/
//go:linkname MkExists C.Z3_mk_exists
func MkExists(c Context, weight c.Uint, num_patterns c.Uint, patterns *Pattern, num_decls c.Uint, sorts *Sort, decl_names *Symbol, body Ast) Ast

/**
  \brief Create a quantifier - universal or existential, with pattern hints.
  See the documentation for #Z3_mk_forall for an explanation of the parameters.

  \param c logical context.
  \param is_forall flag to indicate if this is a universal or existential quantifier.
  \param weight quantifiers are associated with weights indicating the importance of using the quantifier during instantiation. By default, pass the weight 0.
  \param num_patterns number of patterns.
  \param patterns array containing the patterns created using #Z3_mk_pattern.
  \param num_decls number of variables to be bound.
  \param sorts array of sorts of the bound variables.
  \param decl_names names of the bound variables.
  \param body the body of the quantifier.

  \sa Z3_mk_pattern
  \sa Z3_mk_bound
  \sa Z3_mk_forall
  \sa Z3_mk_exists

  def_API('Z3_mk_quantifier', AST, (_in(CONTEXT), _in(BOOL), _in(UINT), _in(UINT), _in_array(3, PATTERN), _in(UINT), _in_array(5, SORT), _in_array(5, SYMBOL), _in(AST)))
*/
//go:linkname MkQuantifier C.Z3_mk_quantifier
func MkQuantifier(c Context, is_forall bool, weight c.Uint, num_patterns c.Uint, patterns *Pattern, num_decls c.Uint, sorts *Sort, decl_names *Symbol, body Ast) Ast

/**
  \brief Create a quantifier - universal or existential, with pattern hints, no patterns, and attributes

  \param c logical context.
  \param is_forall flag to indicate if this is a universal or existential quantifier.
  \param quantifier_id identifier to identify quantifier
  \param skolem_id identifier to identify skolem constants introduced by quantifier.
  \param weight quantifiers are associated with weights indicating the importance of using the quantifier during instantiation. By default, pass the weight 0.
  \param num_patterns number of patterns.
  \param patterns array containing the patterns created using #Z3_mk_pattern.
  \param num_no_patterns number of no_patterns.
  \param no_patterns array containing subexpressions to be excluded from inferred patterns.
  \param num_decls number of variables to be bound.
  \param sorts array of sorts of the bound variables.
  \param decl_names names of the bound variables.
  \param body the body of the quantifier.

  \sa Z3_mk_pattern
  \sa Z3_mk_bound
  \sa Z3_mk_forall
  \sa Z3_mk_exists

  def_API('Z3_mk_quantifier_ex', AST, (_in(CONTEXT), _in(BOOL), _in(UINT), _in(SYMBOL), _in(SYMBOL), _in(UINT), _in_array(5, PATTERN), _in(UINT), _in_array(7, AST), _in(UINT), _in_array(9, SORT), _in_array(9, SYMBOL), _in(AST)))
*/
//go:linkname MkQuantifierEx C.Z3_mk_quantifier_ex
func MkQuantifierEx(c Context, is_forall bool, weight c.Uint, quantifier_id Symbol, skolem_id Symbol, num_patterns c.Uint, patterns *Pattern, num_no_patterns c.Uint, no_patterns *Ast, num_decls c.Uint, sorts *Sort, decl_names *Symbol, body Ast) Ast

/**
  \brief Create a universal quantifier using a list of constants that
  will form the set of bound variables.

  \param c logical context.
  \param weight quantifiers are associated with weights indicating the importance of using
         the quantifier during instantiation. By default, pass the weight 0.
  \param num_bound number of constants to be abstracted into bound variables.
  \param bound array of constants to be abstracted into bound variables.
  \param num_patterns number of patterns.
  \param patterns array containing the patterns created using #Z3_mk_pattern.
  \param body the body of the quantifier.

  \sa Z3_mk_pattern
  \sa Z3_mk_exists_const

  def_API('Z3_mk_forall_const', AST, (_in(CONTEXT), _in(UINT), _in(UINT), _in_array(2, APP), _in(UINT), _in_array(4, PATTERN), _in(AST)))
*/
//go:linkname MkForallConst C.Z3_mk_forall_const
func MkForallConst(c Context, weight c.Uint, num_bound c.Uint, bound *App, num_patterns c.Uint, patterns *Pattern, body Ast) Ast

/**
  \brief Similar to #Z3_mk_forall_const.

  \brief Create an existential quantifier using a list of constants that
  will form the set of bound variables.

  \param c logical context.
  \param weight quantifiers are associated with weights indicating the importance of using
         the quantifier during instantiation. By default, pass the weight 0.
  \param num_bound number of constants to be abstracted into bound variables.
  \param bound array of constants to be abstracted into bound variables.
  \param num_patterns number of patterns.
  \param patterns array containing the patterns created using #Z3_mk_pattern.
  \param body the body of the quantifier.

  \sa Z3_mk_pattern
  \sa Z3_mk_forall_const

  def_API('Z3_mk_exists_const', AST, (_in(CONTEXT), _in(UINT), _in(UINT), _in_array(2, APP), _in(UINT), _in_array(4, PATTERN), _in(AST)))
*/
//go:linkname MkExistsConst C.Z3_mk_exists_const
func MkExistsConst(c Context, weight c.Uint, num_bound c.Uint, bound *App, num_patterns c.Uint, patterns *Pattern, body Ast) Ast

/**
  \brief Create a universal or existential quantifier using a list of
  constants that will form the set of bound variables.

  def_API('Z3_mk_quantifier_const', AST, (_in(CONTEXT), _in(BOOL), _in(UINT), _in(UINT), _in_array(3, APP), _in(UINT), _in_array(5, PATTERN), _in(AST)))
*/
//go:linkname MkQuantifierConst C.Z3_mk_quantifier_const
func MkQuantifierConst(c Context, is_forall bool, weight c.Uint, num_bound c.Uint, bound *App, num_patterns c.Uint, patterns *Pattern, body Ast) Ast

/**
  \brief Create a universal or existential quantifier using a list of
  constants that will form the set of bound variables.

  def_API('Z3_mk_quantifier_const_ex', AST, (_in(CONTEXT), _in(BOOL), _in(UINT), _in(SYMBOL), _in(SYMBOL), _in(UINT), _in_array(5, APP), _in(UINT), _in_array(7, PATTERN), _in(UINT), _in_array(9, AST), _in(AST)))
*/
//go:linkname MkQuantifierConstEx C.Z3_mk_quantifier_const_ex
func MkQuantifierConstEx(c Context, is_forall bool, weight c.Uint, quantifier_id Symbol, skolem_id Symbol, num_bound c.Uint, bound *App, num_patterns c.Uint, patterns *Pattern, num_no_patterns c.Uint, no_patterns *Ast, body Ast) Ast

/**
  \brief Create a lambda expression. It takes an expression \c body that contains bound variables
  of the same sorts as the sorts listed in the array \c sorts. The bound variables are de-Bruijn indices created
  using #Z3_mk_bound. The array \c decl_names contains the names that the quantified formula uses for the
  bound variables. Z3 applies the convention that the last element in the \c decl_names and \c sorts array
  refers to the variable with index 0, the second to last element of \c decl_names and \c sorts refers
  to the variable with index 1, etc.
  The sort of the resulting expression is \c (Array sorts range) where \c range is the sort of \c body.
  For example, if the lambda binds two variables of sort \c Int and \c Bool, and the \c body has sort \c Real,
  the sort of the expression is \c (Array Int Bool Real).

  \param c logical context
  \param num_decls number of variables to be bound.
  \param sorts the sorts of the bound variables.
  \param decl_names names of the bound variables
  \param body the body of the lambda expression.

  \sa Z3_mk_bound
  \sa Z3_mk_forall
  \sa Z3_mk_lambda_const

  def_API('Z3_mk_lambda', AST, (_in(CONTEXT), _in(UINT), _in_array(1, SORT), _in_array(1, SYMBOL), _in(AST)))
*/
//go:linkname MkLambda C.Z3_mk_lambda
func MkLambda(c Context, num_decls c.Uint, sorts *Sort, decl_names *Symbol, body Ast) Ast

/**
  \brief Create a lambda expression using a list of constants that form the set
  of bound variables

  \param c logical context.
  \param num_bound number of constants to be abstracted into bound variables.
  \param bound array of constants to be abstracted into bound variables.
  \param body the body of the lambda expression.

  \sa Z3_mk_bound
  \sa Z3_mk_forall
  \sa Z3_mk_lambda

  def_API('Z3_mk_lambda_const', AST, (_in(CONTEXT), _in(UINT), _in_array(1, APP), _in(AST)))
*/
//go:linkname MkLambdaConst C.Z3_mk_lambda_const
func MkLambdaConst(c Context, num_bound c.Uint, bound *App, body Ast) Ast

/** @name Accessors */
/**@{*/
/**
  \brief Return \c Z3_INT_SYMBOL if the symbol was constructed
  using #Z3_mk_int_symbol, and \c Z3_STRING_SYMBOL if the symbol
  was constructed using #Z3_mk_string_symbol.

  def_API('Z3_get_symbol_kind', UINT, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname GetSymbolKind C.Z3_get_symbol_kind
func GetSymbolKind(c Context, s Symbol) SymbolKind

/**
  \brief Return the symbol int value.

  \pre Z3_get_symbol_kind(s) == Z3_INT_SYMBOL

  \sa Z3_mk_int_symbol

  def_API('Z3_get_symbol_int', INT, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname GetSymbolInt C.Z3_get_symbol_int
func GetSymbolInt(c Context, s Symbol) c.Int

/**
  \brief Return the symbol name.

  \pre Z3_get_symbol_kind(s) == Z3_STRING_SYMBOL

  \warning The returned buffer is statically allocated by Z3. It will
  be automatically deallocated when #Z3_del_context is invoked.
  So, the buffer is invalidated in the next call to \c Z3_get_symbol_string.

  \sa Z3_mk_string_symbol

  def_API('Z3_get_symbol_string', STRING, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname GetSymbolString C.Z3_get_symbol_string
func GetSymbolString(c Context, s Symbol) String

/**
  \brief Return the sort name as a symbol.

  def_API('Z3_get_sort_name', SYMBOL, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetSortName C.Z3_get_sort_name
func GetSortName(c Context, d Sort) Symbol

/**
  \brief Return a unique identifier for \c s.

  def_API('Z3_get_sort_id', UINT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetSortId C.Z3_get_sort_id
func GetSortId(c Context, s Sort) c.Uint

/**
  \brief Convert a \c Z3_sort into \c Z3_ast. This is just type casting.

  def_API('Z3_sort_to_ast', AST, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname SortToAst C.Z3_sort_to_ast
func SortToAst(c Context, s Sort) Ast

/**
  \brief compare sorts.

  def_API('Z3_is_eq_sort', BOOL, (_in(CONTEXT), _in(SORT), _in(SORT)))
*/
//go:linkname IsEqSort C.Z3_is_eq_sort
func IsEqSort(c Context, s1 Sort, s2 Sort) bool

/**
  \brief Return the sort kind (e.g., array, tuple, int, bool, etc).

  \sa Z3_sort_kind

  def_API('Z3_get_sort_kind', UINT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetSortKind C.Z3_get_sort_kind
func GetSortKind(c Context, t Sort) SortKind

/**
  \brief Return the size of the given bit-vector sort.

  \pre Z3_get_sort_kind(c, t) == Z3_BV_SORT

  \sa Z3_mk_bv_sort
  \sa Z3_get_sort_kind

  def_API('Z3_get_bv_sort_size', UINT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetBvSortSize C.Z3_get_bv_sort_size
func GetBvSortSize(c Context, t Sort) c.Uint

/**
  \brief Store the size of the sort in \c r. Return \c false if the call failed.
  That is, Z3_get_sort_kind(s) == Z3_FINITE_DOMAIN_SORT

  def_API('Z3_get_finite_domain_sort_size', BOOL, (_in(CONTEXT), _in(SORT), _out(UINT64)))
*/
//go:linkname GetFiniteDomainSortSize C.Z3_get_finite_domain_sort_size
func GetFiniteDomainSortSize(c Context, s Sort, r *uint64) bool

/**
  \brief Return the domain of the given array sort.
  In the case of a multi-dimensional array, this function returns the sort of the first dimension.

  \pre Z3_get_sort_kind(c, t) == Z3_ARRAY_SORT

  \sa Z3_mk_array_sort
  \sa Z3_get_sort_kind
  \sa Z3_get_array_sort_domain_n

  def_API('Z3_get_array_sort_domain', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetArraySortDomain C.Z3_get_array_sort_domain
func GetArraySortDomain(c Context, t Sort) Sort

/**
  \brief Return the i'th domain sort of an n-dimensional array.

  \pre Z3_get_sort_kind(c, t) == Z3_ARRAY_SORT

  \sa Z3_mk_array_sort
  \sa Z3_get_sort_kind
  \sa Z3_get_array_sort_domain

  def_API('Z3_get_array_sort_domain_n', SORT, (_in(CONTEXT), _in(SORT), _in(UINT)))

*/
//go:linkname GetArraySortDomainN C.Z3_get_array_sort_domain_n
func GetArraySortDomainN(c Context, t Sort, idx c.Uint) Sort

/**
  \brief Return the range of the given array sort.

  \pre Z3_get_sort_kind(c, t) == Z3_ARRAY_SORT

  \sa Z3_mk_array_sort
  \sa Z3_get_sort_kind

  def_API('Z3_get_array_sort_range', SORT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetArraySortRange C.Z3_get_array_sort_range
func GetArraySortRange(c Context, t Sort) Sort

/**
  \brief Return the constructor declaration of the given tuple
  sort.

  \pre Z3_get_sort_kind(c, t) == Z3_DATATYPE_SORT

  \sa Z3_mk_tuple_sort
  \sa Z3_get_sort_kind

  def_API('Z3_get_tuple_sort_mk_decl', FUNC_DECL, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetTupleSortMkDecl C.Z3_get_tuple_sort_mk_decl
func GetTupleSortMkDecl(c Context, t Sort) FuncDecl

/**
  \brief Return the number of fields of the given tuple sort.

  \pre Z3_get_sort_kind(c, t) == Z3_DATATYPE_SORT

  \sa Z3_mk_tuple_sort
  \sa Z3_get_sort_kind

  def_API('Z3_get_tuple_sort_num_fields', UINT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetTupleSortNumFields C.Z3_get_tuple_sort_num_fields
func GetTupleSortNumFields(c Context, t Sort) c.Uint

/**
  \brief Return the i-th field declaration (i.e., projection function declaration)
  of the given tuple sort.

  \pre Z3_get_sort_kind(t) == Z3_DATATYPE_SORT
  \pre i < Z3_get_tuple_sort_num_fields(c, t)

  \sa Z3_mk_tuple_sort
  \sa Z3_get_sort_kind

  def_API('Z3_get_tuple_sort_field_decl', FUNC_DECL, (_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname GetTupleSortFieldDecl C.Z3_get_tuple_sort_field_decl
func GetTupleSortFieldDecl(c Context, t Sort, i c.Uint) FuncDecl

/**
  \brief Return number of constructors for datatype.

  \pre Z3_get_sort_kind(t) == Z3_DATATYPE_SORT

  \sa Z3_get_datatype_sort_constructor
  \sa Z3_get_datatype_sort_recognizer
  \sa Z3_get_datatype_sort_constructor_accessor

  def_API('Z3_get_datatype_sort_num_constructors', UINT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetDatatypeSortNumConstructors C.Z3_get_datatype_sort_num_constructors
func GetDatatypeSortNumConstructors(c Context, t Sort) c.Uint

/**
  \brief Return idx'th constructor.

  \pre Z3_get_sort_kind(t) == Z3_DATATYPE_SORT
  \pre idx < Z3_get_datatype_sort_num_constructors(c, t)

  \sa Z3_get_datatype_sort_num_constructors
  \sa Z3_get_datatype_sort_recognizer
  \sa Z3_get_datatype_sort_constructor_accessor

  def_API('Z3_get_datatype_sort_constructor', FUNC_DECL, (_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname GetDatatypeSortConstructor C.Z3_get_datatype_sort_constructor
func GetDatatypeSortConstructor(c Context, t Sort, idx c.Uint) FuncDecl

/**
  \brief Return idx'th recognizer.

  \pre Z3_get_sort_kind(t) == Z3_DATATYPE_SORT
  \pre idx < Z3_get_datatype_sort_num_constructors(c, t)

  \sa Z3_get_datatype_sort_num_constructors
  \sa Z3_get_datatype_sort_constructor
  \sa Z3_get_datatype_sort_constructor_accessor

  def_API('Z3_get_datatype_sort_recognizer', FUNC_DECL, (_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname GetDatatypeSortRecognizer C.Z3_get_datatype_sort_recognizer
func GetDatatypeSortRecognizer(c Context, t Sort, idx c.Uint) FuncDecl

/**
  \brief Return idx_a'th accessor for the idx_c'th constructor.

  \pre Z3_get_sort_kind(t) == Z3_DATATYPE_SORT
  \pre idx_c < Z3_get_datatype_sort_num_constructors(c, t)
  \pre idx_a < Z3_get_domain_size(c, Z3_get_datatype_sort_constructor(c, idx_c))

  \sa Z3_get_datatype_sort_num_constructors
  \sa Z3_get_datatype_sort_constructor
  \sa Z3_get_datatype_sort_recognizer

  def_API('Z3_get_datatype_sort_constructor_accessor', FUNC_DECL, (_in(CONTEXT), _in(SORT), _in(UINT), _in(UINT)))
*/
//go:linkname GetDatatypeSortConstructorAccessor C.Z3_get_datatype_sort_constructor_accessor
func GetDatatypeSortConstructorAccessor(c Context, t Sort, idx_c c.Uint, idx_a c.Uint) FuncDecl

/**
  \brief Update record field with a value.

  This corresponds to the 'with' construct in OCaml.
  It has the effect of updating a record field with a given value.
  The remaining fields are left unchanged. It is the record
  equivalent of an array store (see \sa Z3_mk_store).
  If the datatype has more than one constructor, then the update function
  behaves as identity if there is a mismatch between the accessor and
  constructor. For example ((_ update-field car) nil 1) is nil,
  while ((_ update-field car) (cons 2 nil) 1) is (cons 1 nil).


  \pre Z3_get_sort_kind(Z3_get_sort(c, t)) == Z3_get_domain(c, field_access, 1) == Z3_DATATYPE_SORT
  \pre Z3_get_sort(c, value) == Z3_get_range(c, field_access)


  def_API('Z3_datatype_update_field', AST, (_in(CONTEXT), _in(FUNC_DECL), _in(AST), _in(AST)))
*/
//go:linkname DatatypeUpdateField C.Z3_datatype_update_field
func DatatypeUpdateField(c Context, field_access FuncDecl, t Ast, value Ast) Ast

/**
  \brief Return arity of relation.

  \pre Z3_get_sort_kind(s) == Z3_RELATION_SORT

  \sa Z3_get_relation_column

  def_API('Z3_get_relation_arity', UINT, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname GetRelationArity C.Z3_get_relation_arity
func GetRelationArity(c Context, s Sort) c.Uint

/**
  \brief Return sort at i'th column of relation sort.

  \pre Z3_get_sort_kind(c, s) == Z3_RELATION_SORT
  \pre col < Z3_get_relation_arity(c, s)

  \sa Z3_get_relation_arity

  def_API('Z3_get_relation_column', SORT, (_in(CONTEXT), _in(SORT), _in(UINT)))
*/
//go:linkname GetRelationColumn C.Z3_get_relation_column
func GetRelationColumn(c Context, s Sort, col c.Uint) Sort

/**
  \brief Pseudo-Boolean relations.

  Encode p1 + p2 + ... + pn <= k

  def_API('Z3_mk_atmost', AST, (_in(CONTEXT), _in(UINT), _in_array(1,AST), _in(UINT)))
*/
//go:linkname MkAtmost C.Z3_mk_atmost
func MkAtmost(c Context, num_args c.Uint, args *Ast, k c.Uint) Ast

/**
  \brief Pseudo-Boolean relations.

  Encode p1 + p2 + ... + pn >= k

  def_API('Z3_mk_atleast', AST, (_in(CONTEXT), _in(UINT), _in_array(1,AST), _in(UINT)))
*/
//go:linkname MkAtleast C.Z3_mk_atleast
func MkAtleast(c Context, num_args c.Uint, args *Ast, k c.Uint) Ast

/**
  \brief Pseudo-Boolean relations.

  Encode k1*p1 + k2*p2 + ... + kn*pn <= k

  def_API('Z3_mk_pble', AST, (_in(CONTEXT), _in(UINT), _in_array(1,AST), _in_array(1,INT), _in(INT)))
*/
//go:linkname MkPble C.Z3_mk_pble
func MkPble(c Context, num_args c.Uint, args *Ast, coeffs *c.Int, k c.Int) Ast

/**
  \brief Pseudo-Boolean relations.

  Encode k1*p1 + k2*p2 + ... + kn*pn >= k

  def_API('Z3_mk_pbge', AST, (_in(CONTEXT), _in(UINT), _in_array(1,AST), _in_array(1,INT), _in(INT)))
*/
//go:linkname MkPbge C.Z3_mk_pbge
func MkPbge(c Context, num_args c.Uint, args *Ast, coeffs *c.Int, k c.Int) Ast

/**
  \brief Pseudo-Boolean relations.

  Encode k1*p1 + k2*p2 + ... + kn*pn = k

  def_API('Z3_mk_pbeq', AST, (_in(CONTEXT), _in(UINT), _in_array(1,AST), _in_array(1,INT), _in(INT)))
*/
//go:linkname MkPbeq C.Z3_mk_pbeq
func MkPbeq(c Context, num_args c.Uint, args *Ast, coeffs *c.Int, k c.Int) Ast

/**
  \brief Convert a \c Z3_func_decl into \c Z3_ast. This is just type casting.

  def_API('Z3_func_decl_to_ast', AST, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname FuncDeclToAst C.Z3_func_decl_to_ast
func FuncDeclToAst(c Context, f FuncDecl) Ast

/**
  \brief Compare terms.

  def_API('Z3_is_eq_func_decl', BOOL, (_in(CONTEXT), _in(FUNC_DECL), _in(FUNC_DECL)))
*/
//go:linkname IsEqFuncDecl C.Z3_is_eq_func_decl
func IsEqFuncDecl(c Context, f1 FuncDecl, f2 FuncDecl) bool

/**
  \brief Return a unique identifier for \c f.

  def_API('Z3_get_func_decl_id', UINT, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetFuncDeclId C.Z3_get_func_decl_id
func GetFuncDeclId(c Context, f FuncDecl) c.Uint

/**
  \brief Return the constant declaration name as a symbol.

  def_API('Z3_get_decl_name', SYMBOL, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetDeclName C.Z3_get_decl_name
func GetDeclName(c Context, d FuncDecl) Symbol

/**
  \brief Return declaration kind corresponding to declaration.

  def_API('Z3_get_decl_kind', UINT, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetDeclKind C.Z3_get_decl_kind
func GetDeclKind(c Context, d FuncDecl) DeclKind

/**
  \brief Return the number of parameters of the given declaration.

  \sa Z3_get_arity

  def_API('Z3_get_domain_size', UINT, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetDomainSize C.Z3_get_domain_size
func GetDomainSize(c Context, d FuncDecl) c.Uint

/**
  \brief Alias for \c Z3_get_domain_size.

  \sa Z3_get_domain_size

  def_API('Z3_get_arity', UINT, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetArity C.Z3_get_arity
func GetArity(c Context, d FuncDecl) c.Uint

/**
  \brief Return the sort of the i-th parameter of the given function declaration.

  \pre i < Z3_get_domain_size(d)

  \sa Z3_get_domain_size

  def_API('Z3_get_domain', SORT, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDomain C.Z3_get_domain
func GetDomain(c Context, d FuncDecl, i c.Uint) Sort

/**
  \brief Return the range of the given declaration.

  If \c d is a constant (i.e., has zero arguments), then this
  function returns the sort of the constant.

  def_API('Z3_get_range', SORT, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetRange C.Z3_get_range
func GetRange(c Context, d FuncDecl) Sort

/**
  \brief Return the number of parameters associated with a declaration.

  def_API('Z3_get_decl_num_parameters', UINT, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname GetDeclNumParameters C.Z3_get_decl_num_parameters
func GetDeclNumParameters(c Context, d FuncDecl) c.Uint

/**
  \brief Return the parameter type associated with a declaration.

  \param c the context
  \param d the function declaration
  \param idx is the index of the named parameter it should be between 0 and the number of parameters.

  def_API('Z3_get_decl_parameter_kind', UINT, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclParameterKind C.Z3_get_decl_parameter_kind
func GetDeclParameterKind(c Context, d FuncDecl, idx c.Uint) ParameterKind

/**
  \brief Return the integer value associated with an integer parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_INT

  def_API('Z3_get_decl_int_parameter', INT, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclIntParameter C.Z3_get_decl_int_parameter
func GetDeclIntParameter(c Context, d FuncDecl, idx c.Uint) c.Int

/**
  \brief Return the double value associated with an double parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_DOUBLE

  def_API('Z3_get_decl_double_parameter', DOUBLE, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclDoubleParameter C.Z3_get_decl_double_parameter
func GetDeclDoubleParameter(c Context, d FuncDecl, idx c.Uint) float64

/**
  \brief Return the double value associated with an double parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_SYMBOL

  def_API('Z3_get_decl_symbol_parameter', SYMBOL, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclSymbolParameter C.Z3_get_decl_symbol_parameter
func GetDeclSymbolParameter(c Context, d FuncDecl, idx c.Uint) Symbol

/**
  \brief Return the sort value associated with a sort parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_SORT

  def_API('Z3_get_decl_sort_parameter', SORT, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclSortParameter C.Z3_get_decl_sort_parameter
func GetDeclSortParameter(c Context, d FuncDecl, idx c.Uint) Sort

/**
  \brief Return the expression value associated with an expression parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_AST

  def_API('Z3_get_decl_ast_parameter', AST, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclAstParameter C.Z3_get_decl_ast_parameter
func GetDeclAstParameter(c Context, d FuncDecl, idx c.Uint) Ast

/**
  \brief Return the expression value associated with an expression parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_FUNC_DECL

  def_API('Z3_get_decl_func_decl_parameter', FUNC_DECL, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclFuncDeclParameter C.Z3_get_decl_func_decl_parameter
func GetDeclFuncDeclParameter(c Context, d FuncDecl, idx c.Uint) FuncDecl

/**
  \brief Return the rational value, as a string, associated with a rational parameter.

  \pre Z3_get_decl_parameter_kind(c, d, idx) == Z3_PARAMETER_RATIONAL

  def_API('Z3_get_decl_rational_parameter', STRING, (_in(CONTEXT), _in(FUNC_DECL), _in(UINT)))
*/
//go:linkname GetDeclRationalParameter C.Z3_get_decl_rational_parameter
func GetDeclRationalParameter(c Context, d FuncDecl, idx c.Uint) String

/**
  \brief Convert a \c Z3_app into \c Z3_ast. This is just type casting.

  def_API('Z3_app_to_ast', AST, (_in(CONTEXT), _in(APP)))
*/
//go:linkname AppToAst C.Z3_app_to_ast
func AppToAst(c Context, a App) Ast

/**
  \brief Return the declaration of a constant or function application.

  def_API('Z3_get_app_decl', FUNC_DECL, (_in(CONTEXT), _in(APP)))
*/
//go:linkname GetAppDecl C.Z3_get_app_decl
func GetAppDecl(c Context, a App) FuncDecl

/**
  \brief Return the number of argument of an application. If \c t
  is an constant, then the number of arguments is 0.

  \sa Z3_get_app_arg

  def_API('Z3_get_app_num_args', UINT, (_in(CONTEXT), _in(APP)))
*/
//go:linkname GetAppNumArgs C.Z3_get_app_num_args
func GetAppNumArgs(c Context, a App) c.Uint

/**
  \brief Return the i-th argument of the given application.

  \pre i < Z3_get_app_num_args(c, a)

  \sa Z3_get_app_num_args

  def_API('Z3_get_app_arg', AST, (_in(CONTEXT), _in(APP), _in(UINT)))
*/
//go:linkname GetAppArg C.Z3_get_app_arg
func GetAppArg(c Context, a App, i c.Uint) Ast

/**
  \brief Compare terms.

  def_API('Z3_is_eq_ast', BOOL, (_in(CONTEXT), _in(AST), _in(AST)))
*/
//go:linkname IsEqAst C.Z3_is_eq_ast
func IsEqAst(c Context, t1 Ast, t2 Ast) bool

/**
  \brief Return a unique identifier for \c t.
  The identifier is unique up to structural equality. Thus, two ast nodes
  created by the same context and having the same children and same function symbols
  have the same identifiers. Ast nodes created in the same context, but having
  different children or different functions have different identifiers.
  Variables and quantifiers are also assigned different identifiers according to
  their structure.

  def_API('Z3_get_ast_id', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetAstId C.Z3_get_ast_id
func GetAstId(c Context, t Ast) c.Uint

/**
  \brief Return a hash code for the given AST.
  The hash code is structural but two different AST objects can map to the same hash.
  The result of \c Z3_get_ast_id returns an identifier that is unique over the
  set of live AST objects.

  def_API('Z3_get_ast_hash', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetAstHash C.Z3_get_ast_hash
func GetAstHash(c Context, a Ast) c.Uint

/**
  \brief Return the sort of an AST node.

  The AST node must be a constant, application, numeral, bound variable, or quantifier.

  def_API('Z3_get_sort', SORT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetSort C.Z3_get_sort
func GetSort(c Context, a Ast) Sort

/**
  \brief Return \c true if the given expression \c t is well sorted.

  def_API('Z3_is_well_sorted', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsWellSorted C.Z3_is_well_sorted
func IsWellSorted(c Context, t Ast) bool

/**
  \brief Return \c Z3_L_TRUE if \c a is true, \c Z3_L_FALSE if it is false, and \c Z3_L_UNDEF otherwise.

  def_API('Z3_get_bool_value', LBOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetBoolValue C.Z3_get_bool_value
func GetBoolValue(c Context, a Ast) Lbool

/**
  \brief Return the kind of the given AST.

  def_API('Z3_get_ast_kind', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetAstKind C.Z3_get_ast_kind
func GetAstKind(c Context, a Ast) AstKind

/**
  def_API('Z3_is_app', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsApp C.Z3_is_app
func IsApp(c Context, a Ast) bool

/**
  def_API('Z3_is_numeral_ast', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsNumeralAst C.Z3_is_numeral_ast
func IsNumeralAst(c Context, a Ast) bool

/**
  \brief Return \c true if the given AST is a real algebraic number.

  def_API('Z3_is_algebraic_number', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsAlgebraicNumber C.Z3_is_algebraic_number
func IsAlgebraicNumber(c Context, a Ast) bool

/**
  \brief Convert an \c ast into an \c APP_AST. This is just type casting.

  \pre \code Z3_get_ast_kind(c, a) == \c Z3_APP_AST \endcode

  def_API('Z3_to_app', APP, (_in(CONTEXT), _in(AST)))
*/
//go:linkname ToApp C.Z3_to_app
func ToApp(c Context, a Ast) App

/**
  \brief Convert an AST into a FUNC_DECL_AST. This is just type casting.

  \pre \code Z3_get_ast_kind(c, a) == Z3_FUNC_DECL_AST \endcode

  def_API('Z3_to_func_decl', FUNC_DECL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname ToFuncDecl C.Z3_to_func_decl
func ToFuncDecl(c Context, a Ast) FuncDecl

/**
  \brief Return numeral value, as a decimal string of a numeric constant term

  \pre Z3_get_ast_kind(c, a) == Z3_NUMERAL_AST

  def_API('Z3_get_numeral_string', STRING, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetNumeralString C.Z3_get_numeral_string
func GetNumeralString(c Context, a Ast) String

/**
  \brief Return numeral value, as a binary string of a numeric constant term

  \pre Z3_get_ast_kind(c, a) == Z3_NUMERAL_AST
  \pre a represents a non-negative integer

  def_API('Z3_get_numeral_binary_string', STRING, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetNumeralBinaryString C.Z3_get_numeral_binary_string
func GetNumeralBinaryString(c Context, a Ast) String

/**
  \brief Return numeral as a string in decimal notation.
  The result has at most \c precision decimal places.

  \pre Z3_get_ast_kind(c, a) == Z3_NUMERAL_AST || Z3_is_algebraic_number(c, a)

  def_API('Z3_get_numeral_decimal_string', STRING, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetNumeralDecimalString C.Z3_get_numeral_decimal_string
func GetNumeralDecimalString(c Context, a Ast, precision c.Uint) String

/**
  \brief Return numeral as a double.

  \pre Z3_get_ast_kind(c, a) == Z3_NUMERAL_AST || Z3_is_algebraic_number(c, a)

  def_API('Z3_get_numeral_double', DOUBLE, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetNumeralDouble C.Z3_get_numeral_double
func GetNumeralDouble(c Context, a Ast) float64

/**
  \brief Return the numerator (as a numeral AST) of a numeral AST of sort Real.

  \pre Z3_get_ast_kind(c, a) == Z3_NUMERAL_AST

  def_API('Z3_get_numerator', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetNumerator C.Z3_get_numerator
func GetNumerator(c Context, a Ast) Ast

/**
  \brief Return the denominator (as a numeral AST) of a numeral AST of sort Real.

  \pre Z3_get_ast_kind(c, a) == Z3_NUMERAL_AST

  def_API('Z3_get_denominator', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetDenominator C.Z3_get_denominator
func GetDenominator(c Context, a Ast) Ast

/**
  \brief Return numeral value, as a pair of 64 bit numbers if the representation fits.

  \param c logical context.
  \param a term.
  \param num numerator.
  \param den denominator.

  Return \c true if the numeral value fits in 64 bit numerals, \c false otherwise.

  \pre Z3_get_ast_kind(a) == Z3_NUMERAL_AST

  def_API('Z3_get_numeral_small', BOOL, (_in(CONTEXT), _in(AST), _out(INT64), _out(INT64)))
*/
//go:linkname GetNumeralSmall C.Z3_get_numeral_small
func GetNumeralSmall(c Context, a Ast, num *int64, den *int64) bool

/**
  \brief Similar to #Z3_get_numeral_string, but only succeeds if
  the value can fit in a machine int. Return \c true if the call succeeded.

  \pre Z3_get_ast_kind(c, v) == Z3_NUMERAL_AST

  \sa Z3_get_numeral_string

  def_API('Z3_get_numeral_int', BOOL, (_in(CONTEXT), _in(AST), _out(INT)))
*/
//go:linkname GetNumeralInt C.Z3_get_numeral_int
func GetNumeralInt(c Context, v Ast, i *c.Int) bool

/**
  \brief Similar to #Z3_get_numeral_string, but only succeeds if
  the value can fit in a machine unsigned int. Return \c true if the call succeeded.

  \pre Z3_get_ast_kind(c, v) == Z3_NUMERAL_AST

  \sa Z3_get_numeral_string

  def_API('Z3_get_numeral_uint', BOOL, (_in(CONTEXT), _in(AST), _out(UINT)))
*/
//go:linkname GetNumeralUint C.Z3_get_numeral_uint
func GetNumeralUint(c Context, v Ast, u *c.Uint) bool

/**
  \brief Similar to #Z3_get_numeral_string, but only succeeds if
  the value can fit in a machine \c uint64_t int. Return \c true if the call succeeded.

  \pre Z3_get_ast_kind(c, v) == Z3_NUMERAL_AST

  \sa Z3_get_numeral_string

  def_API('Z3_get_numeral_uint64', BOOL, (_in(CONTEXT), _in(AST), _out(UINT64)))
*/
//go:linkname GetNumeralUint64 C.Z3_get_numeral_uint64
func GetNumeralUint64(c Context, v Ast, u *uint64) bool

/**
  \brief Similar to #Z3_get_numeral_string, but only succeeds if
  the value can fit in a machine \c int64_t int. Return \c true if the call succeeded.

  \pre Z3_get_ast_kind(c, v) == Z3_NUMERAL_AST

  \sa Z3_get_numeral_string

  def_API('Z3_get_numeral_int64', BOOL, (_in(CONTEXT), _in(AST), _out(INT64)))
*/
//go:linkname GetNumeralInt64 C.Z3_get_numeral_int64
func GetNumeralInt64(c Context, v Ast, i *int64) bool

/**
  \brief Similar to #Z3_get_numeral_string, but only succeeds if
  the value can fit as a rational number as machine \c int64_t int. Return \c true if the call succeeded.

  \pre Z3_get_ast_kind(c, v) == Z3_NUMERAL_AST

  \sa Z3_get_numeral_string

  def_API('Z3_get_numeral_rational_int64', BOOL, (_in(CONTEXT), _in(AST), _out(INT64), _out(INT64)))
*/
//go:linkname GetNumeralRationalInt64 C.Z3_get_numeral_rational_int64
func GetNumeralRationalInt64(c Context, v Ast, num *int64, den *int64) bool

/**
  \brief Return a lower bound for the given real algebraic number.
  The interval isolating the number is smaller than 1/10^precision.
  The result is a numeral AST of sort Real.

  \pre Z3_is_algebraic_number(c, a)

  def_API('Z3_get_algebraic_number_lower', AST, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetAlgebraicNumberLower C.Z3_get_algebraic_number_lower
func GetAlgebraicNumberLower(c Context, a Ast, precision c.Uint) Ast

/**
  \brief Return a upper bound for the given real algebraic number.
  The interval isolating the number is smaller than 1/10^precision.
  The result is a numeral AST of sort Real.

  \pre Z3_is_algebraic_number(c, a)

  def_API('Z3_get_algebraic_number_upper', AST, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetAlgebraicNumberUpper C.Z3_get_algebraic_number_upper
func GetAlgebraicNumberUpper(c Context, a Ast, precision c.Uint) Ast

/**
  \brief Convert a Z3_pattern into Z3_ast. This is just type casting.

  def_API('Z3_pattern_to_ast', AST, (_in(CONTEXT), _in(PATTERN)))
*/
//go:linkname PatternToAst C.Z3_pattern_to_ast
func PatternToAst(c Context, p Pattern) Ast

/**
  \brief Return number of terms in pattern.

  def_API('Z3_get_pattern_num_terms', UINT, (_in(CONTEXT), _in(PATTERN)))
*/
//go:linkname GetPatternNumTerms C.Z3_get_pattern_num_terms
func GetPatternNumTerms(c Context, p Pattern) c.Uint

/**
  \brief Return i'th ast in pattern.

  def_API('Z3_get_pattern', AST, (_in(CONTEXT), _in(PATTERN), _in(UINT)))
*/
//go:linkname GetPattern C.Z3_get_pattern
func GetPattern(c Context, p Pattern, idx c.Uint) Ast

/**
  \brief Return index of de-Bruijn bound variable.

  \pre Z3_get_ast_kind(a) == Z3_VAR_AST

  def_API('Z3_get_index_value', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetIndexValue C.Z3_get_index_value
func GetIndexValue(c Context, a Ast) c.Uint

/**
  \brief Determine if an ast is a universal quantifier.

  def_API('Z3_is_quantifier_forall', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsQuantifierForall C.Z3_is_quantifier_forall
func IsQuantifierForall(c Context, a Ast) bool

/**
  \brief Determine if ast is an existential quantifier.


  def_API('Z3_is_quantifier_exists', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsQuantifierExists C.Z3_is_quantifier_exists
func IsQuantifierExists(c Context, a Ast) bool

/**
  \brief Determine if ast is a lambda expression.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_is_lambda', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsLambda C.Z3_is_lambda
func IsLambda(c Context, a Ast) bool

/**
  \brief Obtain weight of quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_weight', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierWeight C.Z3_get_quantifier_weight
func GetQuantifierWeight(c Context, a Ast) c.Uint

/**
  \brief Obtain skolem id of quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_skolem_id', SYMBOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierSkolemId C.Z3_get_quantifier_skolem_id
func GetQuantifierSkolemId(c Context, a Ast) Symbol

/**
  \brief Obtain id of quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_id', SYMBOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierId C.Z3_get_quantifier_id
func GetQuantifierId(c Context, a Ast) Symbol

/**
  \brief Return number of patterns used in quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_num_patterns', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierNumPatterns C.Z3_get_quantifier_num_patterns
func GetQuantifierNumPatterns(c Context, a Ast) c.Uint

/**
  \brief Return i'th pattern.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_pattern_ast', PATTERN, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetQuantifierPatternAst C.Z3_get_quantifier_pattern_ast
func GetQuantifierPatternAst(c Context, a Ast, i c.Uint) Pattern

/**
  \brief Return number of no_patterns used in quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_num_no_patterns', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierNumNoPatterns C.Z3_get_quantifier_num_no_patterns
func GetQuantifierNumNoPatterns(c Context, a Ast) c.Uint

/**
  \brief Return i'th no_pattern.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_no_pattern_ast', AST, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetQuantifierNoPatternAst C.Z3_get_quantifier_no_pattern_ast
func GetQuantifierNoPatternAst(c Context, a Ast, i c.Uint) Ast

/**
  \brief Return number of bound variables of quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_num_bound', UINT, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierNumBound C.Z3_get_quantifier_num_bound
func GetQuantifierNumBound(c Context, a Ast) c.Uint

/**
  \brief Return symbol of the i'th bound variable.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_bound_name', SYMBOL, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetQuantifierBoundName C.Z3_get_quantifier_bound_name
func GetQuantifierBoundName(c Context, a Ast, i c.Uint) Symbol

/**
  \brief Return sort of the i'th bound variable.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_bound_sort', SORT, (_in(CONTEXT), _in(AST), _in(UINT)))
*/
//go:linkname GetQuantifierBoundSort C.Z3_get_quantifier_bound_sort
func GetQuantifierBoundSort(c Context, a Ast, i c.Uint) Sort

/**
  \brief Return body of quantifier.

  \pre Z3_get_ast_kind(a) == Z3_QUANTIFIER_AST

  def_API('Z3_get_quantifier_body', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetQuantifierBody C.Z3_get_quantifier_body
func GetQuantifierBody(c Context, a Ast) Ast

/**
  \brief Interface to simplifier.

  Provides an interface to the AST simplifier used by Z3.
  It returns an AST object which is equal to the argument.
  The returned AST is simplified using algebraic simplification rules,
  such as constant propagation (propagating true/false over logical connectives).

  \sa Z3_simplify_ex

  def_API('Z3_simplify', AST, (_in(CONTEXT), _in(AST)))
*/
//go:linkname Simplify C.Z3_simplify
func Simplify(c Context, a Ast) Ast

/**
  \brief Interface to simplifier.

  Provides an interface to the AST simplifier used by Z3.
  This procedure is similar to #Z3_simplify, but the behavior of the simplifier
  can be configured using the given parameter set.

  \sa Z3_simplify
  \sa Z3_simplify_get_help
  \sa Z3_simplify_get_param_descrs

  def_API('Z3_simplify_ex', AST, (_in(CONTEXT), _in(AST), _in(PARAMS)))
*/
//go:linkname SimplifyEx C.Z3_simplify_ex
func SimplifyEx(c Context, a Ast, p Params) Ast

/**
  \brief Return a string describing all available parameters.

   \sa Z3_simplify_ex
   \sa Z3_simplify_get_param_descrs

  def_API('Z3_simplify_get_help', STRING, (_in(CONTEXT),))
*/
//go:linkname SimplifyGetHelp C.Z3_simplify_get_help
func SimplifyGetHelp(c Context) String

/**
  \brief Return the parameter description set for the simplify procedure.

   \sa Z3_simplify_ex
   \sa Z3_simplify_get_help

  def_API('Z3_simplify_get_param_descrs', PARAM_DESCRS, (_in(CONTEXT),))
*/
//go:linkname SimplifyGetParamDescrs C.Z3_simplify_get_param_descrs
func SimplifyGetParamDescrs(c Context) ParamDescrs

/** @name Modifiers */
/**@{*/
/**
  \brief Update the arguments of term \c a using the arguments \c args.
  The number of arguments \c num_args should coincide
  with the number of arguments to \c a.
  If \c a is a quantifier, then num_args has to be 1.

  def_API('Z3_update_term', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST)))
*/
//go:linkname UpdateTerm C.Z3_update_term
func UpdateTerm(c Context, a Ast, num_args c.Uint, args *Ast) Ast

/**
  \brief Substitute every occurrence of \ccode{from[i]} in \c a with \ccode{to[i]}, for \c i smaller than \c num_exprs.
  The result is the new AST. The arrays \c from and \c to must have size \c num_exprs.
  For every \c i smaller than \c num_exprs, we must have that sort of \ccode{from[i]} must be equal to sort of \ccode{to[i]}.

  def_API('Z3_substitute', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST), _in_array(2, AST)))
*/
//go:linkname Substitute C.Z3_substitute
func Substitute(c Context, a Ast, num_exprs c.Uint, from *Ast, to *Ast) Ast

/**
  \brief Substitute the variables in \c a with the expressions in \c to.
  For every \c i smaller than \c num_exprs, the variable with de-Bruijn index \c i is replaced with term \ccode{to[i]}.
  Note that a variable is created using the function \ref Z3_mk_bound.

  def_API('Z3_substitute_vars', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, AST)))
*/
//go:linkname SubstituteVars C.Z3_substitute_vars
func SubstituteVars(c Context, a Ast, num_exprs c.Uint, to *Ast) Ast

/**
  \brief Substitute functions in \c from with new expressions in \c to.

  The expressions in \c to can have free variables. The free variable in \c to at index 0
  refers to the first argument of \c from, the free variable at index 1 corresponds to the second argument.

  def_API('Z3_substitute_funs', AST, (_in(CONTEXT), _in(AST), _in(UINT), _in_array(2, FUNC_DECL), _in_array(2, AST)))
*/
//go:linkname SubstituteFuns C.Z3_substitute_funs
func SubstituteFuns(c Context, a Ast, num_funs c.Uint, from *FuncDecl, to *Ast) Ast

/**
  \brief Translate/Copy the AST \c a from context \c source to context \c target.
  AST \c a must have been created using context \c source.
  \pre source != target

  def_API('Z3_translate', AST, (_in(CONTEXT), _in(AST), _in(CONTEXT)))
*/
//go:linkname Translate C.Z3_translate
func Translate(source Context, a Ast, target Context) Ast

/**
  \brief Create a fresh model object. It has reference count 0.

  def_API('Z3_mk_model', MODEL, (_in(CONTEXT),))
*/
//go:linkname MkModel C.Z3_mk_model
func MkModel(c Context) Model

/**
  \brief Increment the reference counter of the given model.

  def_API('Z3_model_inc_ref', VOID, (_in(CONTEXT), _in(MODEL)))
*/
//go:linkname ModelIncRef C.Z3_model_inc_ref
func ModelIncRef(c Context, m Model)

/**
  \brief Decrement the reference counter of the given model.

  def_API('Z3_model_dec_ref', VOID, (_in(CONTEXT), _in(MODEL)))
*/
//go:linkname ModelDecRef C.Z3_model_dec_ref
func ModelDecRef(c Context, m Model)

/**
  \brief Evaluate the AST node \c t in the given model.
  Return \c true if succeeded, and store the result in \c v.

  If \c model_completion is \c true, then Z3 will assign an interpretation for any constant or function that does
  not have an interpretation in \c m. These constants and functions were essentially don't cares.

  If \c model_completion is \c false, then Z3 will not assign interpretations to constants for functions that do
  not have interpretations in \c m. Evaluation behaves as the identify function in this case.

  The evaluation may fail for the following reasons:

  - \c t contains a quantifier.

  - the model \c m is partial, that is, it doesn't have a complete interpretation for uninterpreted functions.
  That is, the option \ccode{MODEL_PARTIAL=true} was used.

  - \c t is type incorrect.

  - \c Z3_interrupt was invoked during evaluation.

  def_API('Z3_model_eval', BOOL, (_in(CONTEXT), _in(MODEL), _in(AST), _in(BOOL), _out(AST)))
*/
//go:linkname ModelEval C.Z3_model_eval
func ModelEval(c Context, m Model, t Ast, model_completion bool, v *Ast) bool

/**
  \brief Return the interpretation (i.e., assignment) of constant \c a in the model \c m.
  Return \c NULL, if the model does not assign an interpretation for \c a.
  That should be interpreted as: the value of \c a does not matter.

  \pre Z3_get_arity(c, a) == 0

  def_API('Z3_model_get_const_interp', AST, (_in(CONTEXT), _in(MODEL), _in(FUNC_DECL)))
*/
//go:linkname ModelGetConstInterp C.Z3_model_get_const_interp
func ModelGetConstInterp(c Context, m Model, a FuncDecl) Ast

/**
  \brief Test if there exists an interpretation (i.e., assignment) for \c a in the model \c m.

  def_API('Z3_model_has_interp', BOOL, (_in(CONTEXT), _in(MODEL), _in(FUNC_DECL)))
*/
//go:linkname ModelHasInterp C.Z3_model_has_interp
func ModelHasInterp(c Context, m Model, a FuncDecl) bool

/**
  \brief Return the interpretation of the function \c f in the model \c m.
  Return \c NULL, if the model does not assign an interpretation for \c f.
  That should be interpreted as: the \c f does not matter.

  \pre Z3_get_arity(c, f) > 0

  \remark Reference counting must be used to manage Z3_func_interp objects, even when the Z3_context was
  created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_model_get_func_interp', FUNC_INTERP, (_in(CONTEXT), _in(MODEL), _in(FUNC_DECL)))
*/
//go:linkname ModelGetFuncInterp C.Z3_model_get_func_interp
func ModelGetFuncInterp(c Context, m Model, f FuncDecl) FuncInterp

/**
  \brief Return the number of constants assigned by the given model.

  \sa Z3_model_get_const_decl

  def_API('Z3_model_get_num_consts', UINT, (_in(CONTEXT), _in(MODEL)))
*/
//go:linkname ModelGetNumConsts C.Z3_model_get_num_consts
func ModelGetNumConsts(c Context, m Model) c.Uint

/**
  \brief Return the i-th constant in the given model.

  \pre i < Z3_model_get_num_consts(c, m)

  \sa Z3_model_eval
  \sa Z3_model_get_num_consts

  def_API('Z3_model_get_const_decl', FUNC_DECL, (_in(CONTEXT), _in(MODEL), _in(UINT)))
*/
//go:linkname ModelGetConstDecl C.Z3_model_get_const_decl
func ModelGetConstDecl(c Context, m Model, i c.Uint) FuncDecl

/**
  \brief Return the number of function interpretations in the given model.

  A function interpretation is represented as a finite map and an 'else' value.
  Each entry in the finite map represents the value of a function given a set of arguments.

  \sa Z3_model_get_func_decl

  def_API('Z3_model_get_num_funcs', UINT, (_in(CONTEXT), _in(MODEL)))
*/
//go:linkname ModelGetNumFuncs C.Z3_model_get_num_funcs
func ModelGetNumFuncs(c Context, m Model) c.Uint

/**
  \brief Return the declaration of the i-th function in the given model.

  \pre i < Z3_model_get_num_funcs(c, m)

  \sa Z3_model_get_num_funcs

  def_API('Z3_model_get_func_decl', FUNC_DECL, (_in(CONTEXT), _in(MODEL), _in(UINT)))
*/
//go:linkname ModelGetFuncDecl C.Z3_model_get_func_decl
func ModelGetFuncDecl(c Context, m Model, i c.Uint) FuncDecl

/**
  \brief Return the number of uninterpreted sorts that \c m assigns an interpretation to.

  Z3 also provides an interpretation for uninterpreted sorts used in a formula.
  The interpretation for a sort \c s is a finite set of distinct values. We say this finite set is
  the "universe" of \c s.

  \sa Z3_model_get_sort
  \sa Z3_model_get_sort_universe

  def_API('Z3_model_get_num_sorts', UINT, (_in(CONTEXT), _in(MODEL)))
*/
//go:linkname ModelGetNumSorts C.Z3_model_get_num_sorts
func ModelGetNumSorts(c Context, m Model) c.Uint

/**
  \brief Return a uninterpreted sort that \c m assigns an interpretation.

  \pre i < Z3_model_get_num_sorts(c, m)

  \sa Z3_model_get_num_sorts
  \sa Z3_model_get_sort_universe

  def_API('Z3_model_get_sort', SORT, (_in(CONTEXT), _in(MODEL), _in(UINT)))
*/
//go:linkname ModelGetSort C.Z3_model_get_sort
func ModelGetSort(c Context, m Model, i c.Uint) Sort

/**
  \brief Return the finite set of distinct values that represent the interpretation for sort \c s.

  \sa Z3_model_get_num_sorts
  \sa Z3_model_get_sort

  def_API('Z3_model_get_sort_universe', AST_VECTOR, (_in(CONTEXT), _in(MODEL), _in(SORT)))
*/
//go:linkname ModelGetSortUniverse C.Z3_model_get_sort_universe
func ModelGetSortUniverse(c Context, m Model, s Sort) AstVector

/**
  \brief translate model from context \c c to context \c dst.

  \remark Use this method for cloning state between contexts. Note that
  operations on contexts are not thread safe and therefore all operations
  that related to a given context have to be synchronized (or run in the same thread).

  def_API('Z3_model_translate', MODEL, (_in(CONTEXT), _in(MODEL), _in(CONTEXT)))
*/
//go:linkname ModelTranslate C.Z3_model_translate
func ModelTranslate(c Context, m Model, dst Context) Model

/**
  \brief The \ccode{(_ as-array f)} AST node is a construct for assigning interpretations for arrays in Z3.
  It is the array such that forall indices \c i we have that \ccode{(select (_ as-array f) i)} is equal to \ccode{(f i)}.
  This procedure returns \c true if the \c a is an \c as-array AST node.

  Z3 current solvers have minimal support for \c as_array nodes.

  \sa Z3_get_as_array_func_decl

  def_API('Z3_is_as_array', BOOL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname IsAsArray C.Z3_is_as_array
func IsAsArray(c Context, a Ast) bool

/**
  \brief Return the function declaration \c f associated with a \ccode{(_ as_array f)} node.

  \sa Z3_is_as_array

  def_API('Z3_get_as_array_func_decl', FUNC_DECL, (_in(CONTEXT), _in(AST)))
*/
//go:linkname GetAsArrayFuncDecl C.Z3_get_as_array_func_decl
func GetAsArrayFuncDecl(c Context, a Ast) FuncDecl

/**
  \brief Create a fresh func_interp object, add it to a model for a specified function.
  It has reference count 0.

  \param c context
  \param m model
  \param f function declaration
  \param default_value default value for function interpretation

  def_API('Z3_add_func_interp', FUNC_INTERP, (_in(CONTEXT), _in(MODEL), _in(FUNC_DECL), _in(AST)))
*/
//go:linkname AddFuncInterp C.Z3_add_func_interp
func AddFuncInterp(c Context, m Model, f FuncDecl, default_value Ast) FuncInterp

/**
  \brief Add a constant interpretation.

  def_API('Z3_add_const_interp', VOID, (_in(CONTEXT), _in(MODEL), _in(FUNC_DECL), _in(AST)))
*/
//go:linkname AddConstInterp C.Z3_add_const_interp
func AddConstInterp(c Context, m Model, f FuncDecl, a Ast)

/**
  \brief Increment the reference counter of the given \c Z3_func_interp object.

  def_API('Z3_func_interp_inc_ref', VOID, (_in(CONTEXT), _in(FUNC_INTERP)))
*/
//go:linkname FuncInterpIncRef C.Z3_func_interp_inc_ref
func FuncInterpIncRef(c Context, f FuncInterp)

/**
  \brief Decrement the reference counter of the given \c Z3_func_interp object.

  def_API('Z3_func_interp_dec_ref', VOID, (_in(CONTEXT), _in(FUNC_INTERP)))
*/
//go:linkname FuncInterpDecRef C.Z3_func_interp_dec_ref
func FuncInterpDecRef(c Context, f FuncInterp)

/**
  \brief Return the number of entries in the given function interpretation.

  A function interpretation is represented as a finite map and an 'else' value.
  Each entry in the finite map represents the value of a function given a set of arguments.
  This procedure return the number of element in the finite map of \c f.

  \sa Z3_func_interp_get_entry

  def_API('Z3_func_interp_get_num_entries', UINT, (_in(CONTEXT), _in(FUNC_INTERP)))
*/
//go:linkname FuncInterpGetNumEntries C.Z3_func_interp_get_num_entries
func FuncInterpGetNumEntries(c Context, f FuncInterp) c.Uint

/**
  \brief Return a "point" of the given function interpretation. It represents the
  value of \c f in a particular point.

  \pre i < Z3_func_interp_get_num_entries(c, f)

  \sa Z3_func_interp_get_num_entries

  def_API('Z3_func_interp_get_entry', FUNC_ENTRY, (_in(CONTEXT), _in(FUNC_INTERP), _in(UINT)))
*/
//go:linkname FuncInterpGetEntry C.Z3_func_interp_get_entry
func FuncInterpGetEntry(c Context, f FuncInterp, i c.Uint) FuncEntry

/**
  \brief Return the 'else' value of the given function interpretation.

  A function interpretation is represented as a finite map and an 'else' value.
  This procedure returns the 'else' value.

  def_API('Z3_func_interp_get_else', AST, (_in(CONTEXT), _in(FUNC_INTERP)))
*/
//go:linkname FuncInterpGetElse C.Z3_func_interp_get_else
func FuncInterpGetElse(c Context, f FuncInterp) Ast

/**
  \brief Return the 'else' value of the given function interpretation.

  A function interpretation is represented as a finite map and an 'else' value.
  This procedure can be used to update the 'else' value.

  def_API('Z3_func_interp_set_else', VOID, (_in(CONTEXT), _in(FUNC_INTERP), _in(AST)))
*/
//go:linkname FuncInterpSetElse C.Z3_func_interp_set_else
func FuncInterpSetElse(c Context, f FuncInterp, else_value Ast)

/**
  \brief Return the arity (number of arguments) of the given function interpretation.

  def_API('Z3_func_interp_get_arity', UINT, (_in(CONTEXT), _in(FUNC_INTERP)))
*/
//go:linkname FuncInterpGetArity C.Z3_func_interp_get_arity
func FuncInterpGetArity(c Context, f FuncInterp) c.Uint

/**
  \brief add a function entry to a function interpretation.

  \param c logical context
  \param fi a function interpretation to be updated.
  \param args list of arguments. They should be constant values (such as integers) and be of the same types as the domain of the function.
  \param value value of the function when the parameters match args.

  It is assumed that entries added to a function cover disjoint arguments.
  If an two entries are added with the same arguments, only the second insertion survives and the
  first inserted entry is removed.

  def_API('Z3_func_interp_add_entry', VOID, (_in(CONTEXT), _in(FUNC_INTERP), _in(AST_VECTOR), _in(AST)))
*/
//go:linkname FuncInterpAddEntry C.Z3_func_interp_add_entry
func FuncInterpAddEntry(c Context, fi FuncInterp, args AstVector, value Ast)

/**
  \brief Increment the reference counter of the given \c Z3_func_entry object.

  def_API('Z3_func_entry_inc_ref', VOID, (_in(CONTEXT), _in(FUNC_ENTRY)))
*/
//go:linkname FuncEntryIncRef C.Z3_func_entry_inc_ref
func FuncEntryIncRef(c Context, e FuncEntry)

/**
  \brief Decrement the reference counter of the given \c Z3_func_entry object.

  def_API('Z3_func_entry_dec_ref', VOID, (_in(CONTEXT), _in(FUNC_ENTRY)))
*/
//go:linkname FuncEntryDecRef C.Z3_func_entry_dec_ref
func FuncEntryDecRef(c Context, e FuncEntry)

/**
  \brief Return the value of this point.

  A \c Z3_func_entry object represents an element in the finite map used to encode
  a function interpretation.

  \sa Z3_func_interp_get_entry

  def_API('Z3_func_entry_get_value', AST, (_in(CONTEXT), _in(FUNC_ENTRY)))
*/
//go:linkname FuncEntryGetValue C.Z3_func_entry_get_value
func FuncEntryGetValue(c Context, e FuncEntry) Ast

/**
  \brief Return the number of arguments in a \c Z3_func_entry object.

  \sa Z3_func_entry_get_arg
  \sa Z3_func_interp_get_entry

  def_API('Z3_func_entry_get_num_args', UINT, (_in(CONTEXT), _in(FUNC_ENTRY)))
*/
//go:linkname FuncEntryGetNumArgs C.Z3_func_entry_get_num_args
func FuncEntryGetNumArgs(c Context, e FuncEntry) c.Uint

/**
  \brief Return an argument of a \c Z3_func_entry object.

  \pre i < Z3_func_entry_get_num_args(c, e)

  \sa Z3_func_entry_get_num_args
  \sa Z3_func_interp_get_entry

  def_API('Z3_func_entry_get_arg', AST, (_in(CONTEXT), _in(FUNC_ENTRY), _in(UINT)))
*/
//go:linkname FuncEntryGetArg C.Z3_func_entry_get_arg
func FuncEntryGetArg(c Context, e FuncEntry, i c.Uint) Ast

/** @name Interaction logging */
/**@{*/
/**
  \brief Log interaction to a file.

  \sa Z3_append_log
  \sa Z3_close_log

  extra_API('Z3_open_log', INT, (_in(STRING),))
*/
//go:linkname OpenLog C.Z3_open_log
func OpenLog(filename String) bool

/**
  \brief Append user-defined string to interaction log.

  The interaction log is opened using #Z3_open_log.
  It contains the formulas that are checked using Z3.
  You can use this command to append comments, for instance.

  \sa Z3_open_log
  \sa Z3_close_log

  extra_API('Z3_append_log', VOID, (_in(STRING),))
*/
//go:linkname AppendLog C.Z3_append_log
func AppendLog(string String)

/**
  \brief Close interaction log.

  \sa Z3_open_log
  \sa Z3_append_log

  extra_API('Z3_close_log', VOID, ())
*/
//go:linkname CloseLog C.Z3_close_log
func CloseLog()

/**
  \brief Enable/disable printing warning messages to the console.

  Warnings are printed after passing \c true, warning messages are
  suppressed after calling this method with \c false.

  def_API('Z3_toggle_warning_messages', VOID, (_in(BOOL),))
*/
//go:linkname ToggleWarningMessages C.Z3_toggle_warning_messages
func ToggleWarningMessages(enabled bool)

/** @name String conversion */
/**@{*/
/**
  \brief Select mode for the format used for pretty-printing AST nodes.

  The default mode for pretty printing AST nodes is to produce
  SMT-LIB style output where common subexpressions are printed
  at each occurrence. The mode is called \c Z3_PRINT_SMTLIB_FULL.
  To print shared common subexpressions only once,
  use the \c Z3_PRINT_LOW_LEVEL mode.
  To print in way that conforms to SMT-LIB standards and uses let
  expressions to share common sub-expressions use \c Z3_PRINT_SMTLIB2_COMPLIANT.

  \sa Z3_ast_to_string
  \sa Z3_pattern_to_string
  \sa Z3_func_decl_to_string

  def_API('Z3_set_ast_print_mode', VOID, (_in(CONTEXT), _in(PRINT_MODE)))
*/
//go:linkname SetAstPrintMode C.Z3_set_ast_print_mode
func SetAstPrintMode(c Context, mode AstPrintMode)

/**
  \brief Convert the given AST node into a string.

  \warning The result buffer is statically allocated by Z3. It will
  be automatically deallocated when #Z3_del_context is invoked.
  So, the buffer is invalidated in the next call to \c Z3_ast_to_string.

  \sa Z3_pattern_to_string
  \sa Z3_sort_to_string

  def_API('Z3_ast_to_string', STRING, (_in(CONTEXT), _in(AST)))
*/
//go:linkname AstToString C.Z3_ast_to_string
func AstToString(c Context, a Ast) String

/**
  def_API('Z3_pattern_to_string', STRING, (_in(CONTEXT), _in(PATTERN)))
*/
//go:linkname PatternToString C.Z3_pattern_to_string
func PatternToString(c Context, p Pattern) String

/**
  def_API('Z3_sort_to_string', STRING, (_in(CONTEXT), _in(SORT)))
*/
//go:linkname SortToString C.Z3_sort_to_string
func SortToString(c Context, s Sort) String

/**
  def_API('Z3_func_decl_to_string', STRING, (_in(CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname FuncDeclToString C.Z3_func_decl_to_string
func FuncDeclToString(c Context, d FuncDecl) String

/**
  \brief Convert the given model into a string.

  \warning The result buffer is statically allocated by Z3. It will
  be automatically deallocated when #Z3_del_context is invoked.
  So, the buffer is invalidated in the next call to \c Z3_model_to_string.

  def_API('Z3_model_to_string', STRING, (_in(CONTEXT), _in(MODEL)))
*/
//go:linkname ModelToString C.Z3_model_to_string
func ModelToString(c Context, m Model) String

/**
  \brief Convert the given benchmark into SMT-LIB formatted string.

  \warning The result buffer is statically allocated by Z3. It will
  be automatically deallocated when #Z3_del_context is invoked.
  So, the buffer is invalidated in the next call to \c Z3_benchmark_to_smtlib_string.

  \param c - context.
  \param name - name of benchmark. The argument is optional.
  \param logic - the benchmark logic.
  \param status - the status string (sat, unsat, or unknown)
  \param attributes - other attributes, such as source, difficulty or category.
  \param num_assumptions - number of assumptions.
  \param assumptions - auxiliary assumptions.
  \param formula - formula to be checked for consistency in conjunction with assumptions.

  def_API('Z3_benchmark_to_smtlib_string', STRING, (_in(CONTEXT), _in(STRING), _in(STRING), _in(STRING), _in(STRING), _in(UINT), _in_array(5, AST), _in(AST)))
*/
//go:linkname BenchmarkToSmtlibString C.Z3_benchmark_to_smtlib_string
func BenchmarkToSmtlibString(c Context, name String, logic String, status String, attributes String, num_assumptions c.Uint, assumptions *Ast, formula Ast) String

/** @name Parser interface */
/**@{*/
/**
  \brief Parse the given string using the SMT-LIB2 parser.

  It returns a formula comprising of the conjunction of assertions in the scope
  (up to push/pop) at the end of the string.

  def_API('Z3_parse_smtlib2_string', AST_VECTOR, (_in(CONTEXT), _in(STRING), _in(UINT), _in_array(2, SYMBOL), _in_array(2, SORT), _in(UINT), _in_array(5, SYMBOL), _in_array(5, FUNC_DECL)))
*/
//go:linkname ParseSmtlib2String C.Z3_parse_smtlib2_string
func ParseSmtlib2String(c Context, str String, num_sorts c.Uint, sort_names *Symbol, sorts *Sort, num_decls c.Uint, decl_names *Symbol, decls *FuncDecl) AstVector

/**
  \brief Similar to #Z3_parse_smtlib2_string, but reads the benchmark from a file.

  def_API('Z3_parse_smtlib2_file', AST_VECTOR, (_in(CONTEXT), _in(STRING), _in(UINT), _in_array(2, SYMBOL), _in_array(2, SORT), _in(UINT), _in_array(5, SYMBOL), _in_array(5, FUNC_DECL)))
*/
//go:linkname ParseSmtlib2File C.Z3_parse_smtlib2_file
func ParseSmtlib2File(c Context, file_name String, num_sorts c.Uint, sort_names *Symbol, sorts *Sort, num_decls c.Uint, decl_names *Symbol, decls *FuncDecl) AstVector

/**
  \brief Parse and evaluate and SMT-LIB2 command sequence. The state from a previous call is saved so the next
         evaluation builds on top of the previous call.

  \returns output generated from processing commands.

  def_API('Z3_eval_smtlib2_string', STRING, (_in(CONTEXT), _in(STRING),))
*/
//go:linkname EvalSmtlib2String C.Z3_eval_smtlib2_string
func EvalSmtlib2String(c Context, str String) String

/**
  \brief Create a parser context.

  A parser context maintains state between calls to \c Z3_parser_context_parse_string
  where the caller can pass in a set of SMTLIB2 commands.
  It maintains all the declarations from previous calls together with
  of sorts and function declarations (including 0-ary) that are added directly to the context.

  def_API('Z3_mk_parser_context', PARSER_CONTEXT, (_in(CONTEXT),))
*/
//go:linkname MkParserContext C.Z3_mk_parser_context
func MkParserContext(c Context) ParserContext

/**
  \brief Increment the reference counter of the given \c Z3_parser_context object.

  def_API('Z3_parser_context_inc_ref', VOID, (_in(CONTEXT), _in(PARSER_CONTEXT)))
*/
//go:linkname ParserContextIncRef C.Z3_parser_context_inc_ref
func ParserContextIncRef(c Context, pc ParserContext)

/**
  \brief Decrement the reference counter of the given \c Z3_parser_context object.

  def_API('Z3_parser_context_dec_ref', VOID, (_in(CONTEXT), _in(PARSER_CONTEXT)))
*/
//go:linkname ParserContextDecRef C.Z3_parser_context_dec_ref
func ParserContextDecRef(c Context, pc ParserContext)

/**
  \brief Add a sort declaration.

  def_API('Z3_parser_context_add_sort', VOID, (_in(CONTEXT), _in(PARSER_CONTEXT), _in(SORT)))
*/
//go:linkname ParserContextAddSort C.Z3_parser_context_add_sort
func ParserContextAddSort(c Context, pc ParserContext, s Sort)

/**
  \brief Add a function declaration.

  def_API('Z3_parser_context_add_decl', VOID, (_in(CONTEXT), _in(PARSER_CONTEXT), _in(FUNC_DECL)))
*/
//go:linkname ParserContextAddDecl C.Z3_parser_context_add_decl
func ParserContextAddDecl(c Context, pc ParserContext, f FuncDecl)

/**
  \brief Parse a string of SMTLIB2 commands. Return assertions.

  def_API('Z3_parser_context_from_string', AST_VECTOR, (_in(CONTEXT), _in(PARSER_CONTEXT), _in(STRING)))
*/
//go:linkname ParserContextFromString C.Z3_parser_context_from_string
func ParserContextFromString(c Context, pc ParserContext, s String) AstVector

/**
  \brief Return the error code for the last API call.

  A call to a Z3 function may return a non Z3_OK error code,
  when it is not used correctly.

  \sa Z3_set_error_handler

  def_API('Z3_get_error_code', UINT, (_in(CONTEXT), ))
*/
//go:linkname GetErrorCode C.Z3_get_error_code
func GetErrorCode(c Context) ErrorCode

/**
  \brief Register a Z3 error handler.

  A call to a Z3 function may return a non \c Z3_OK error code, when
  it is not used correctly.  An error handler can be registered
  and will be called in this case.  To disable the use of the
  error handler, simply register with \c h=NULL.

  \warning Log files, created using #Z3_open_log, may be potentially incomplete/incorrect if error handlers are used.

  \sa Z3_get_error_code
*/
//go:linkname SetErrorHandler C.Z3_set_error_handler
func SetErrorHandler(c Context, h ErrorHandler)

/**
  \brief Set an error.

  def_API('Z3_set_error', VOID, (_in(CONTEXT), _in(ERROR_CODE)))
*/
//go:linkname SetError C.Z3_set_error
func SetError(c Context, e ErrorCode)

/**
  \brief Return a string describing the given error code.

  def_API('Z3_get_error_msg', STRING, (_in(CONTEXT), _in(ERROR_CODE)))
*/
//go:linkname GetErrorMsg C.Z3_get_error_msg
func GetErrorMsg(c Context, err ErrorCode) String

/**
  \brief Return Z3 version number information.

  \sa Z3_get_full_version

  def_API('Z3_get_version', VOID, (_out(UINT), _out(UINT), _out(UINT), _out(UINT)))
*/
//go:linkname GetVersion C.Z3_get_version
func GetVersion(major *c.Uint, minor *c.Uint, build_number *c.Uint, revision_number *c.Uint)

/**
  \brief Return a string that fully describes the version of Z3 in use.

  \sa Z3_get_version

  def_API('Z3_get_full_version', STRING, ())
*/
//go:linkname GetFullVersion C.Z3_get_full_version
func GetFullVersion() String

/**
  \brief Enable tracing messages tagged as \c tag when Z3 is compiled in debug mode.
  It is a NOOP otherwise

  \sa Z3_disable_trace

  def_API('Z3_enable_trace', VOID, (_in(STRING),))
*/
//go:linkname EnableTrace C.Z3_enable_trace
func EnableTrace(tag String)

/**
  \brief Disable tracing messages tagged as \c tag when Z3 is compiled in debug mode.
  It is a NOOP otherwise

  \sa Z3_enable_trace

  def_API('Z3_disable_trace', VOID, (_in(STRING),))
*/
//go:linkname DisableTrace C.Z3_disable_trace
func DisableTrace(tag String)

/**
  \brief Reset all allocated resources.

  Use this facility on out-of memory errors.
  It allows discharging the previous state and resuming afresh.
  Any pointers previously returned by the API
  become invalid.

  def_API('Z3_reset_memory', VOID, ())
*/
//go:linkname ResetMemory C.Z3_reset_memory
func ResetMemory()

/**
  \brief Destroy all allocated resources.

  Any pointers previously returned by the API become invalid.
  Can be used for memory leak detection.

  def_API('Z3_finalize_memory', VOID, ())
*/
//go:linkname FinalizeMemory C.Z3_finalize_memory
func FinalizeMemory()

/** @name Goals */
/**@{*/
/**
  \brief Create a goal (aka problem). A goal is essentially a set
  of formulas, that can be solved and/or transformed using
  tactics and solvers.

  If \c models is \c true, then model generation is enabled for the new goal.

  If \c unsat_cores is \c true, then unsat core generation is enabled for the new goal.

  If \c proofs is \c true, then proof generation is enabled for the new goal. Remark, the
  Z3 context \c c must have been created with proof generation support.

  \remark Reference counting must be used to manage goals, even when the \c Z3_context was
  created using #Z3_mk_context instead of #Z3_mk_context_rc.

  def_API('Z3_mk_goal', GOAL, (_in(CONTEXT), _in(BOOL), _in(BOOL), _in(BOOL)))
*/
//go:linkname MkGoal C.Z3_mk_goal
func MkGoal(c Context, models bool, unsat_cores bool, proofs bool) Goal

/**
  \brief Increment the reference counter of the given goal.

  def_API('Z3_goal_inc_ref', VOID, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalIncRef C.Z3_goal_inc_ref
func GoalIncRef(c Context, g Goal)

/**
  \brief Decrement the reference counter of the given goal.

  def_API('Z3_goal_dec_ref', VOID, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalDecRef C.Z3_goal_dec_ref
func GoalDecRef(c Context, g Goal)

/**
  \brief Return the "precision" of the given goal. Goals can be transformed using over and under approximations.
  A under approximation is applied when the objective is to find a model for a given goal.
  An over approximation is applied when the objective is to find a proof for a given goal.

  def_API('Z3_goal_precision', UINT, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalPrecision C.Z3_goal_precision
func GoalPrecision(c Context, g Goal) GoalPrec

/**
  \brief Add a new formula \c a to the given goal.
   The formula is split according to the following procedure that is applied
   until a fixed-point:
      Conjunctions are split into separate formulas.
      Negations are distributed over disjunctions, resulting in separate formulas.
   If the goal is \c false, adding new formulas is a no-op.
   If the formula \c a is \c true, then nothing is added.
   If the formula \c a is \c false, then the entire goal is replaced by the formula \c false.

  def_API('Z3_goal_assert', VOID, (_in(CONTEXT), _in(GOAL), _in(AST)))
*/
//go:linkname GoalAssert C.Z3_goal_assert
func GoalAssert(c Context, g Goal, a Ast)

/**
  \brief Return \c true if the given goal contains the formula \c false.

  def_API('Z3_goal_inconsistent', BOOL, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalInconsistent C.Z3_goal_inconsistent
func GoalInconsistent(c Context, g Goal) bool

/**
  \brief Return the depth of the given goal. It tracks how many transformations were applied to it.

  def_API('Z3_goal_depth', UINT, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalDepth C.Z3_goal_depth
func GoalDepth(c Context, g Goal) c.Uint

/**
  \brief Erase all formulas from the given goal.

  def_API('Z3_goal_reset', VOID, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalReset C.Z3_goal_reset
func GoalReset(c Context, g Goal)

/**
  \brief Return the number of formulas in the given goal.

  def_API('Z3_goal_size', UINT, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalSize C.Z3_goal_size
func GoalSize(c Context, g Goal) c.Uint

/**
  \brief Return a formula from the given goal.

  \pre idx < Z3_goal_size(c, g)

  def_API('Z3_goal_formula', AST, (_in(CONTEXT), _in(GOAL), _in(UINT)))
*/
//go:linkname GoalFormula C.Z3_goal_formula
func GoalFormula(c Context, g Goal, idx c.Uint) Ast

/**
  \brief Return the number of formulas, subformulas and terms in the given goal.

  def_API('Z3_goal_num_exprs', UINT, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalNumExprs C.Z3_goal_num_exprs
func GoalNumExprs(c Context, g Goal) c.Uint

/**
  \brief Return \c true if the goal is empty, and it is precise or the product of a under approximation.

  def_API('Z3_goal_is_decided_sat', BOOL, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalIsDecidedSat C.Z3_goal_is_decided_sat
func GoalIsDecidedSat(c Context, g Goal) bool

/**
  \brief Return \c true if the goal contains false, and it is precise or the product of an over approximation.

  def_API('Z3_goal_is_decided_unsat', BOOL, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalIsDecidedUnsat C.Z3_goal_is_decided_unsat
func GoalIsDecidedUnsat(c Context, g Goal) bool

/**
  \brief Copy a goal \c g from the context \c source to the context \c target.

  def_API('Z3_goal_translate', GOAL, (_in(CONTEXT), _in(GOAL), _in(CONTEXT)))
*/
//go:linkname GoalTranslate C.Z3_goal_translate
func GoalTranslate(source Context, g Goal, target Context) Goal

/**
  \brief Convert a model of the formulas of a goal to a model of an original goal.
  The model may be null, in which case the returned model is valid if the goal was
  established satisfiable.

  def_API('Z3_goal_convert_model', MODEL, (_in(CONTEXT), _in(GOAL), _in(MODEL)))
*/
//go:linkname GoalConvertModel C.Z3_goal_convert_model
func GoalConvertModel(c Context, g Goal, m Model) Model

/**
  \brief Convert a goal into a string.

  def_API('Z3_goal_to_string', STRING, (_in(CONTEXT), _in(GOAL)))
*/
//go:linkname GoalToString C.Z3_goal_to_string
func GoalToString(c Context, g Goal) String

/**
  \brief Convert a goal into a DIMACS formatted string.
  The goal must be in CNF. You can convert a goal to CNF
  by applying the tseitin-cnf tactic. Bit-vectors are not automatically
  converted to Booleans either, so if the caller intends to
  preserve satisfiability, it should apply bit-blasting tactics.
  Quantifiers and theory atoms will not be encoded.

  def_API('Z3_goal_to_dimacs_string', STRING, (_in(CONTEXT), _in(GOAL), _in(BOOL)))
*/
//go:linkname GoalToDimacsString C.Z3_goal_to_dimacs_string
func GoalToDimacsString(c Context, g Goal, include_names bool) String

/** @name Tactics, Simplifiers and Probes */
/**@{*/
/**
  \brief Return a tactic associated with the given name.
  The complete list of tactics may be obtained using the procedures #Z3_get_num_tactics and #Z3_get_tactic_name.
  It may also be obtained using the command \ccode{(help-tactic)} in the SMT 2.0 front-end.

  Tactics are the basic building block for creating custom solvers for specific problem domains.

  def_API('Z3_mk_tactic', TACTIC, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname MkTactic C.Z3_mk_tactic
func MkTactic(c Context, name String) Tactic

/**
  \brief Increment the reference counter of the given tactic.

  def_API('Z3_tactic_inc_ref', VOID, (_in(CONTEXT), _in(TACTIC)))
*/
//go:linkname TacticIncRef C.Z3_tactic_inc_ref
func TacticIncRef(c Context, t Tactic)

/**
  \brief Decrement the reference counter of the given tactic.

  def_API('Z3_tactic_dec_ref', VOID, (_in(CONTEXT), _in(TACTIC)))
*/
//go:linkname TacticDecRef C.Z3_tactic_dec_ref
func TacticDecRef(c Context, g Tactic)

/**
  \brief Return a probe associated with the given name.
  The complete list of probes may be obtained using the procedures #Z3_get_num_probes and #Z3_get_probe_name.
  It may also be obtained using the command \ccode{(help-tactic)} in the SMT 2.0 front-end.

  Probes are used to inspect a goal (aka problem) and collect information that may be used to decide
  which solver and/or preprocessing step will be used.

  def_API('Z3_mk_probe', PROBE, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname MkProbe C.Z3_mk_probe
func MkProbe(c Context, name String) Probe

/**
  \brief Increment the reference counter of the given probe.

  def_API('Z3_probe_inc_ref', VOID, (_in(CONTEXT), _in(PROBE)))
*/
//go:linkname ProbeIncRef C.Z3_probe_inc_ref
func ProbeIncRef(c Context, p Probe)

/**
  \brief Decrement the reference counter of the given probe.

  def_API('Z3_probe_dec_ref', VOID, (_in(CONTEXT), _in(PROBE)))
*/
//go:linkname ProbeDecRef C.Z3_probe_dec_ref
func ProbeDecRef(c Context, p Probe)

/**
  \brief Return a tactic that applies \c t1 to a given goal and \c t2
  to every subgoal produced by \c t1.

  def_API('Z3_tactic_and_then', TACTIC, (_in(CONTEXT), _in(TACTIC), _in(TACTIC)))
*/
//go:linkname TacticAndThen C.Z3_tactic_and_then
func TacticAndThen(c Context, t1 Tactic, t2 Tactic) Tactic

/**
  \brief Return a tactic that first applies \c t1 to a given goal,
  if it fails then returns the result of \c t2 applied to the given goal.

  def_API('Z3_tactic_or_else', TACTIC, (_in(CONTEXT), _in(TACTIC), _in(TACTIC)))
*/
//go:linkname TacticOrElse C.Z3_tactic_or_else
func TacticOrElse(c Context, t1 Tactic, t2 Tactic) Tactic

/**
  \brief Return a tactic that applies the given tactics in parallel.

  def_API('Z3_tactic_par_or', TACTIC, (_in(CONTEXT), _in(UINT), _in_array(1, TACTIC)))
*/
//go:linkname TacticParOr C.Z3_tactic_par_or
func TacticParOr(c Context, num c.Uint, ts *Tactic) Tactic

/**
  \brief Return a tactic that applies \c t1 to a given goal and then \c t2
  to every subgoal produced by \c t1. The subgoals are processed in parallel.

  def_API('Z3_tactic_par_and_then', TACTIC, (_in(CONTEXT), _in(TACTIC), _in(TACTIC)))
*/
//go:linkname TacticParAndThen C.Z3_tactic_par_and_then
func TacticParAndThen(c Context, t1 Tactic, t2 Tactic) Tactic

/**
  \brief Return a tactic that applies \c t to a given goal for \c ms milliseconds.
  If \c t does not terminate in \c ms milliseconds, then it fails.

  def_API('Z3_tactic_try_for', TACTIC, (_in(CONTEXT), _in(TACTIC), _in(UINT)))
*/
//go:linkname TacticTryFor C.Z3_tactic_try_for
func TacticTryFor(c Context, t Tactic, ms c.Uint) Tactic

/**
  \brief Return a tactic that applies \c t to a given goal is the probe \c p evaluates to true.
  If \c p evaluates to false, then the new tactic behaves like the skip tactic.

  def_API('Z3_tactic_when', TACTIC, (_in(CONTEXT), _in(PROBE), _in(TACTIC)))
*/
//go:linkname TacticWhen C.Z3_tactic_when
func TacticWhen(c Context, p Probe, t Tactic) Tactic

/**
  \brief Return a tactic that applies \c t1 to a given goal if the probe \c p evaluates to true,
  and \c t2 if \c p evaluates to false.

  def_API('Z3_tactic_cond', TACTIC, (_in(CONTEXT), _in(PROBE), _in(TACTIC), _in(TACTIC)))
*/
//go:linkname TacticCond C.Z3_tactic_cond
func TacticCond(c Context, p Probe, t1 Tactic, t2 Tactic) Tactic

/**
  \brief Return a tactic that keeps applying \c t until the goal is not modified anymore or the maximum
  number of iterations \c max is reached.

  def_API('Z3_tactic_repeat', TACTIC, (_in(CONTEXT), _in(TACTIC), _in(UINT)))
*/
//go:linkname TacticRepeat C.Z3_tactic_repeat
func TacticRepeat(c Context, t Tactic, max c.Uint) Tactic

/**
  \brief Return a tactic that just return the given goal.

  def_API('Z3_tactic_skip', TACTIC, (_in(CONTEXT),))
*/
//go:linkname TacticSkip C.Z3_tactic_skip
func TacticSkip(c Context) Tactic

/**
  \brief Return a tactic that always fails.

  def_API('Z3_tactic_fail', TACTIC, (_in(CONTEXT),))
*/
//go:linkname TacticFail C.Z3_tactic_fail
func TacticFail(c Context) Tactic

/**
  \brief Return a tactic that fails if the probe \c p evaluates to false.

  def_API('Z3_tactic_fail_if', TACTIC, (_in(CONTEXT), _in(PROBE)))
*/
//go:linkname TacticFailIf C.Z3_tactic_fail_if
func TacticFailIf(c Context, p Probe) Tactic

/**
  \brief Return a tactic that fails if the goal is not trivially satisfiable (i.e., empty) or
  trivially unsatisfiable (i.e., contains false).

  def_API('Z3_tactic_fail_if_not_decided', TACTIC, (_in(CONTEXT),))
*/
//go:linkname TacticFailIfNotDecided C.Z3_tactic_fail_if_not_decided
func TacticFailIfNotDecided(c Context) Tactic

/**
  \brief Return a tactic that applies \c t using the given set of parameters.

  def_API('Z3_tactic_using_params', TACTIC, (_in(CONTEXT), _in(TACTIC), _in(PARAMS)))
*/
//go:linkname TacticUsingParams C.Z3_tactic_using_params
func TacticUsingParams(c Context, t Tactic, p Params) Tactic

/**
  \brief Return a simplifier associated with the given name.
  The complete list of simplifiers may be obtained using the procedures #Z3_get_num_simplifiers and #Z3_get_simplifier_name.
  It may also be obtained using the command \ccode{(help-simplifier)} in the SMT 2.0 front-end.

  Simplifiers are the basic building block for creating custom solvers for specific problem domains.

  def_API('Z3_mk_simplifier', SIMPLIFIER, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname MkSimplifier C.Z3_mk_simplifier
func MkSimplifier(c Context, name String) Simplifier

/**
  \brief Increment the reference counter of the given simplifier.

  def_API('Z3_simplifier_inc_ref', VOID, (_in(CONTEXT), _in(SIMPLIFIER)))
*/
//go:linkname SimplifierIncRef C.Z3_simplifier_inc_ref
func SimplifierIncRef(c Context, t Simplifier)

/**
  \brief Decrement the reference counter of the given simplifier.

  def_API('Z3_simplifier_dec_ref', VOID, (_in(CONTEXT), _in(SIMPLIFIER)))
*/
//go:linkname SimplifierDecRef C.Z3_simplifier_dec_ref
func SimplifierDecRef(c Context, g Simplifier)

/**
  \brief Attach simplifier to a solver. The solver will use the simplifier for incremental pre-processing.

     def_API('Z3_solver_add_simplifier', SOLVER, (_in(CONTEXT), _in(SOLVER), _in(SIMPLIFIER)))
*/
//go:linkname SolverAddSimplifier C.Z3_solver_add_simplifier
func SolverAddSimplifier(c Context, solver Solver, simplifier Simplifier) Solver

/**
  \brief Return a simplifier that applies \c t1 to a given goal and \c t2
  to every subgoal produced by \c t1.

  def_API('Z3_simplifier_and_then', SIMPLIFIER, (_in(CONTEXT), _in(SIMPLIFIER), _in(SIMPLIFIER)))
*/
//go:linkname SimplifierAndThen C.Z3_simplifier_and_then
func SimplifierAndThen(c Context, t1 Simplifier, t2 Simplifier) Simplifier

/**
  \brief Return a simplifier that applies \c t using the given set of parameters.

  def_API('Z3_simplifier_using_params', SIMPLIFIER, (_in(CONTEXT), _in(SIMPLIFIER), _in(PARAMS)))
*/
//go:linkname SimplifierUsingParams C.Z3_simplifier_using_params
func SimplifierUsingParams(c Context, t Simplifier, p Params) Simplifier

/**
  \brief Return the number of builtin simplifiers available in Z3.

  \sa Z3_get_simplifier_name

  def_API('Z3_get_num_simplifiers', UINT, (_in(CONTEXT),))
*/
//go:linkname GetNumSimplifiers C.Z3_get_num_simplifiers
func GetNumSimplifiers(c Context) c.Uint

/**
  \brief Return the name of the idx simplifier.

  \pre i < Z3_get_num_simplifiers(c)

  \sa Z3_get_num_simplifiers

  def_API('Z3_get_simplifier_name', STRING, (_in(CONTEXT), _in(UINT)))
*/
//go:linkname GetSimplifierName C.Z3_get_simplifier_name
func GetSimplifierName(c Context, i c.Uint) String

/**
  \brief Return a string containing a description of parameters accepted by the given simplifier.

  def_API('Z3_simplifier_get_help', STRING, (_in(CONTEXT), _in(SIMPLIFIER)))
*/
//go:linkname SimplifierGetHelp C.Z3_simplifier_get_help
func SimplifierGetHelp(c Context, t Simplifier) String

/**
  \brief Return the parameter description set for the given simplifier object.

  def_API('Z3_simplifier_get_param_descrs', PARAM_DESCRS, (_in(CONTEXT), _in(SIMPLIFIER)))
*/
//go:linkname SimplifierGetParamDescrs C.Z3_simplifier_get_param_descrs
func SimplifierGetParamDescrs(c Context, t Simplifier) ParamDescrs

/**
  \brief Return a string containing a description of the simplifier with the given name.

  def_API('Z3_simplifier_get_descr', STRING, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname SimplifierGetDescr C.Z3_simplifier_get_descr
func SimplifierGetDescr(c Context, name String) String

/**
  \brief Return a probe that always evaluates to val.

  def_API('Z3_probe_const', PROBE, (_in(CONTEXT), _in(DOUBLE)))
*/
//go:linkname ProbeConst C.Z3_probe_const
func ProbeConst(x Context, val float64) Probe

/**
  \brief Return a probe that evaluates to "true" when the value returned by \c p1 is less than the value returned by \c p2.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_lt', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeLt C.Z3_probe_lt
func ProbeLt(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when the value returned by \c p1 is greater than the value returned by \c p2.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_gt', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeGt C.Z3_probe_gt
func ProbeGt(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when the value returned by \c p1 is less than or equal to the value returned by \c p2.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_le', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeLe C.Z3_probe_le
func ProbeLe(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when the value returned by \c p1 is greater than or equal to the value returned by \c p2.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_ge', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeGe C.Z3_probe_ge
func ProbeGe(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when the value returned by \c p1 is equal to the value returned by \c p2.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_eq', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeEq C.Z3_probe_eq
func ProbeEq(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when \c p1 and \c p2 evaluates to true.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_and', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeAnd C.Z3_probe_and
func ProbeAnd(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when \c p1 or \c p2 evaluates to true.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_or', PROBE, (_in(CONTEXT), _in(PROBE), _in(PROBE)))
*/
//go:linkname ProbeOr C.Z3_probe_or
func ProbeOr(x Context, p1 Probe, p2 Probe) Probe

/**
  \brief Return a probe that evaluates to "true" when \c p does not evaluate to true.

  \remark For probes, "true" is any value different from 0.0.

  def_API('Z3_probe_not', PROBE, (_in(CONTEXT), _in(PROBE)))
*/
//go:linkname ProbeNot C.Z3_probe_not
func ProbeNot(x Context, p Probe) Probe

/**
  \brief Return the number of builtin tactics available in Z3.

  \sa Z3_get_tactic_name

  def_API('Z3_get_num_tactics', UINT, (_in(CONTEXT),))
*/
//go:linkname GetNumTactics C.Z3_get_num_tactics
func GetNumTactics(c Context) c.Uint

/**
  \brief Return the name of the idx tactic.

  \pre i < Z3_get_num_tactics(c)

  \sa Z3_get_num_tactics

  def_API('Z3_get_tactic_name', STRING, (_in(CONTEXT), _in(UINT)))
*/
//go:linkname GetTacticName C.Z3_get_tactic_name
func GetTacticName(c Context, i c.Uint) String

/**
  \brief Return the number of builtin probes available in Z3.

  \sa Z3_get_probe_name

  def_API('Z3_get_num_probes', UINT, (_in(CONTEXT),))
*/
//go:linkname GetNumProbes C.Z3_get_num_probes
func GetNumProbes(c Context) c.Uint

/**
  \brief Return the name of the \c i probe.

  \pre i < Z3_get_num_probes(c)

  \sa Z3_get_num_probes

  def_API('Z3_get_probe_name', STRING, (_in(CONTEXT), _in(UINT)))
*/
//go:linkname GetProbeName C.Z3_get_probe_name
func GetProbeName(c Context, i c.Uint) String

/**
  \brief Return a string containing a description of parameters accepted by the given tactic.

  def_API('Z3_tactic_get_help', STRING, (_in(CONTEXT), _in(TACTIC)))
*/
//go:linkname TacticGetHelp C.Z3_tactic_get_help
func TacticGetHelp(c Context, t Tactic) String

/**
  \brief Return the parameter description set for the given tactic object.

  def_API('Z3_tactic_get_param_descrs', PARAM_DESCRS, (_in(CONTEXT), _in(TACTIC)))
*/
//go:linkname TacticGetParamDescrs C.Z3_tactic_get_param_descrs
func TacticGetParamDescrs(c Context, t Tactic) ParamDescrs

/**
  \brief Return a string containing a description of the tactic with the given name.

  def_API('Z3_tactic_get_descr', STRING, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname TacticGetDescr C.Z3_tactic_get_descr
func TacticGetDescr(c Context, name String) String

/**
  \brief Return a string containing a description of the probe with the given name.

  def_API('Z3_probe_get_descr', STRING, (_in(CONTEXT), _in(STRING)))
*/
//go:linkname ProbeGetDescr C.Z3_probe_get_descr
func ProbeGetDescr(c Context, name String) String

/**
  \brief Execute the probe over the goal. The probe always produce a double value.
  "Boolean" probes return 0.0 for false, and a value different from 0.0 for true.

  def_API('Z3_probe_apply', DOUBLE, (_in(CONTEXT), _in(PROBE), _in(GOAL)))
*/
//go:linkname ProbeApply C.Z3_probe_apply
func ProbeApply(c Context, p Probe, g Goal) float64

/**
  \brief Apply tactic \c t to the goal \c g.

  \sa Z3_tactic_apply_ex

  def_API('Z3_tactic_apply', APPLY_RESULT, (_in(CONTEXT), _in(TACTIC), _in(GOAL)))
*/
//go:linkname TacticApply C.Z3_tactic_apply
func TacticApply(c Context, t Tactic, g Goal) ApplyResult

/**
  \brief Apply tactic \c t to the goal \c g using the parameter set \c p.

  \sa Z3_tactic_apply

  def_API('Z3_tactic_apply_ex', APPLY_RESULT, (_in(CONTEXT), _in(TACTIC), _in(GOAL), _in(PARAMS)))
*/
//go:linkname TacticApplyEx C.Z3_tactic_apply_ex
func TacticApplyEx(c Context, t Tactic, g Goal, p Params) ApplyResult

/**
  \brief Increment the reference counter of the given \c Z3_apply_result object.

  def_API('Z3_apply_result_inc_ref', VOID, (_in(CONTEXT), _in(APPLY_RESULT)))
*/
//go:linkname ApplyResultIncRef C.Z3_apply_result_inc_ref
func ApplyResultIncRef(c Context, r ApplyResult)

/**
  \brief Decrement the reference counter of the given \c Z3_apply_result object.

  def_API('Z3_apply_result_dec_ref', VOID, (_in(CONTEXT), _in(APPLY_RESULT)))
*/
//go:linkname ApplyResultDecRef C.Z3_apply_result_dec_ref
func ApplyResultDecRef(c Context, r ApplyResult)

/**
  \brief Convert the \c Z3_apply_result object returned by #Z3_tactic_apply into a string.

  def_API('Z3_apply_result_to_string', STRING, (_in(CONTEXT), _in(APPLY_RESULT)))
*/
//go:linkname ApplyResultToString C.Z3_apply_result_to_string
func ApplyResultToString(c Context, r ApplyResult) String

/**
  \brief Return the number of subgoals in the \c Z3_apply_result object returned by #Z3_tactic_apply.

  \sa Z3_apply_result_get_subgoal

  def_API('Z3_apply_result_get_num_subgoals', UINT, (_in(CONTEXT), _in(APPLY_RESULT)))
*/
//go:linkname ApplyResultGetNumSubgoals C.Z3_apply_result_get_num_subgoals
func ApplyResultGetNumSubgoals(c Context, r ApplyResult) c.Uint

/**
  \brief Return one of the subgoals in the \c Z3_apply_result object returned by #Z3_tactic_apply.

  \pre i < Z3_apply_result_get_num_subgoals(c, r)

  \sa Z3_apply_result_get_num_subgoals

  def_API('Z3_apply_result_get_subgoal', GOAL, (_in(CONTEXT), _in(APPLY_RESULT), _in(UINT)))
*/
//go:linkname ApplyResultGetSubgoal C.Z3_apply_result_get_subgoal
func ApplyResultGetSubgoal(c Context, r ApplyResult, i c.Uint) Goal

/** @name Solvers*/
/**@{*/
/**
  \brief Create a new solver. This solver is a "combined solver" (see
  combined_solver module) that internally uses a non-incremental (solver1) and an
  incremental solver (solver2). This combined solver changes its behaviour based
  on how it is used and how its parameters are set.

  If the solver is used in a non incremental way (i.e. no calls to
  #Z3_solver_push() or #Z3_solver_pop(), and no calls to
  #Z3_solver_assert() or #Z3_solver_assert_and_track() after checking
  satisfiability without an intervening #Z3_solver_reset()) then solver1
  will be used. This solver will apply Z3's "default" tactic.

  The "default" tactic will attempt to probe the logic used by the
  assertions and will apply a specialized tactic if one is supported.
  Otherwise the general `(and-then simplify smt)` tactic will be used.

  If the solver is used in an incremental way then the combined solver
  will switch to using solver2 (which behaves similarly to the general
  "smt" tactic).

  Note however it is possible to set the `solver2_timeout`,
  `solver2_unknown`, and `ignore_solver1` parameters of the combined
  solver to change its behaviour.

  The function #Z3_solver_get_model retrieves a model if the
  assertions is satisfiable (i.e., the result is \c
  Z3_L_TRUE) and model construction is enabled.
  The function #Z3_solver_get_model can also be used even
  if the result is \c Z3_L_UNDEF, but the returned model
  is not guaranteed to satisfy quantified assertions.

  \remark User must use #Z3_solver_inc_ref and #Z3_solver_dec_ref to manage solver objects.
  Even if the context was created using #Z3_mk_context instead of #Z3_mk_context_rc.

  \sa Z3_mk_simple_solver
  \sa Z3_mk_solver_for_logic
  \sa Z3_mk_solver_from_tactic

  def_API('Z3_mk_solver', SOLVER, (_in(CONTEXT),))
*/
//go:linkname MkSolver C.Z3_mk_solver
func MkSolver(c Context) Solver

/**
  \brief Create a new incremental solver.

  This is equivalent to applying the "smt" tactic.

  Unlike #Z3_mk_solver() this solver
    - Does not attempt to apply any logic specific tactics.
    - Does not change its behaviour based on whether it used
      incrementally/non-incrementally.

  Note that these differences can result in very different performance
  compared to #Z3_mk_solver().

  The function #Z3_solver_get_model retrieves a model if the
  assertions is satisfiable (i.e., the result is \c
  Z3_L_TRUE) and model construction is enabled.
  The function #Z3_solver_get_model can also be used even
  if the result is \c Z3_L_UNDEF, but the returned model
  is not guaranteed to satisfy quantified assertions.

  \remark User must use #Z3_solver_inc_ref and #Z3_solver_dec_ref to manage solver objects.
  Even if the context was created using #Z3_mk_context instead of #Z3_mk_context_rc.

  \sa Z3_mk_solver
  \sa Z3_mk_solver_for_logic
  \sa Z3_mk_solver_from_tactic

  def_API('Z3_mk_simple_solver', SOLVER, (_in(CONTEXT),))
*/
//go:linkname MkSimpleSolver C.Z3_mk_simple_solver
func MkSimpleSolver(c Context) Solver

/**
  \brief Create a new solver customized for the given logic.
  It behaves like #Z3_mk_solver if the logic is unknown or unsupported.

  \remark User must use #Z3_solver_inc_ref and #Z3_solver_dec_ref to manage solver objects.
  Even if the context was created using #Z3_mk_context instead of #Z3_mk_context_rc.

  \sa Z3_mk_solver
  \sa Z3_mk_simple_solver
  \sa Z3_mk_solver_from_tactic

  def_API('Z3_mk_solver_for_logic', SOLVER, (_in(CONTEXT), _in(SYMBOL)))
*/
//go:linkname MkSolverForLogic C.Z3_mk_solver_for_logic
func MkSolverForLogic(c Context, logic Symbol) Solver

/**
  \brief Create a new solver that is implemented using the given tactic.
  The solver supports the commands #Z3_solver_push and #Z3_solver_pop, but it
  will always solve each #Z3_solver_check from scratch.

  \remark User must use #Z3_solver_inc_ref and #Z3_solver_dec_ref to manage solver objects.
  Even if the context was created using #Z3_mk_context instead of #Z3_mk_context_rc.

  \sa Z3_mk_solver
  \sa Z3_mk_simple_solver
  \sa Z3_mk_solver_for_logic

  def_API('Z3_mk_solver_from_tactic', SOLVER, (_in(CONTEXT), _in(TACTIC)))
*/
//go:linkname MkSolverFromTactic C.Z3_mk_solver_from_tactic
func MkSolverFromTactic(c Context, t Tactic) Solver

/**
  \brief Copy a solver \c s from the context \c source to the context \c target.

  def_API('Z3_solver_translate', SOLVER, (_in(CONTEXT), _in(SOLVER), _in(CONTEXT)))
*/
//go:linkname SolverTranslate C.Z3_solver_translate
func SolverTranslate(source Context, s Solver, target Context) Solver

/**
  \brief Ad-hoc method for importing model conversion from solver.

  This method is used for scenarios where \c src has been used to solve a set
  of formulas and was interrupted. The \c dst solver may be a strengthening of \c src
  obtained from cubing (assigning a subset of literals or adding constraints over the
  assertions available in \c src). If \c dst ends up being satisfiable, the model for \c dst
  may not correspond to a model of the original formula due to inprocessing in \c src.
  This method is used to take the side-effect of inprocessing into account when returning
  a model for \c dst.

  def_API('Z3_solver_import_model_converter', VOID, (_in(CONTEXT), _in(SOLVER), _in(SOLVER)))
*/
//go:linkname SolverImportModelConverter C.Z3_solver_import_model_converter
func SolverImportModelConverter(ctx Context, src Solver, dst Solver)

/**
  \brief Return a string describing all solver available parameters.

  \sa Z3_solver_get_param_descrs
  \sa Z3_solver_set_params

  def_API('Z3_solver_get_help', STRING, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetHelp C.Z3_solver_get_help
func SolverGetHelp(c Context, s Solver) String

/**
  \brief Return the parameter description set for the given solver object.

  \sa Z3_solver_get_help
  \sa Z3_solver_set_params

  def_API('Z3_solver_get_param_descrs', PARAM_DESCRS, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetParamDescrs C.Z3_solver_get_param_descrs
func SolverGetParamDescrs(c Context, s Solver) ParamDescrs

/**
  \brief Set the given solver using the given parameters.

  \sa Z3_solver_get_help
  \sa Z3_solver_get_param_descrs

  def_API('Z3_solver_set_params', VOID, (_in(CONTEXT), _in(SOLVER), _in(PARAMS)))
*/
//go:linkname SolverSetParams C.Z3_solver_set_params
func SolverSetParams(c Context, s Solver, p Params)

/**
  \brief Increment the reference counter of the given solver.

  def_API('Z3_solver_inc_ref', VOID, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverIncRef C.Z3_solver_inc_ref
func SolverIncRef(c Context, s Solver)

/**
  \brief Decrement the reference counter of the given solver.

  def_API('Z3_solver_dec_ref', VOID, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverDecRef C.Z3_solver_dec_ref
func SolverDecRef(c Context, s Solver)

/**
  \brief Solver local interrupt.
  Normally you should use Z3_interrupt to cancel solvers because only
  one solver is enabled concurrently per context.
  However, per GitHub issue #1006, there are use cases where
  it is more convenient to cancel a specific solver. Solvers
  that are not selected for interrupts are left alone.

  def_API('Z3_solver_interrupt', VOID, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverInterrupt C.Z3_solver_interrupt
func SolverInterrupt(c Context, s Solver)

/**
  \brief Create a backtracking point.

  The solver contains a stack of assertions.

  \sa Z3_solver_get_num_scopes
  \sa Z3_solver_pop

  def_API('Z3_solver_push', VOID, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverPush C.Z3_solver_push
func SolverPush(c Context, s Solver)

/**
  \brief Backtrack \c n backtracking points.

  \sa Z3_solver_get_num_scopes
  \sa Z3_solver_push

  \pre n <= Z3_solver_get_num_scopes(c, s)

  def_API('Z3_solver_pop', VOID, (_in(CONTEXT), _in(SOLVER), _in(UINT)))
*/
//go:linkname SolverPop C.Z3_solver_pop
func SolverPop(c Context, s Solver, n c.Uint)

/**
  \brief Remove all assertions from the solver.

  \sa Z3_solver_assert
  \sa Z3_solver_assert_and_track

  def_API('Z3_solver_reset', VOID, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverReset C.Z3_solver_reset
func SolverReset(c Context, s Solver)

/**
  \brief Return the number of backtracking points.

  \sa Z3_solver_push
  \sa Z3_solver_pop

  def_API('Z3_solver_get_num_scopes', UINT, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetNumScopes C.Z3_solver_get_num_scopes
func SolverGetNumScopes(c Context, s Solver) c.Uint

/**
  \brief Assert a constraint into the solver.

  The functions #Z3_solver_check and #Z3_solver_check_assumptions should be
  used to check whether the logical context is consistent or not.

  \sa Z3_solver_assert_and_track
  \sa Z3_solver_reset

  def_API('Z3_solver_assert', VOID, (_in(CONTEXT), _in(SOLVER), _in(AST)))
*/
//go:linkname SolverAssert C.Z3_solver_assert
func SolverAssert(c Context, s Solver, a Ast)

/**
  \brief Assert a constraint \c a into the solver, and track it (in the unsat) core using
  the Boolean constant \c p.

  This API is an alternative to #Z3_solver_check_assumptions for extracting unsat cores.
  Both APIs can be used in the same solver. The unsat core will contain a combination
  of the Boolean variables provided using Z3_solver_assert_and_track and the Boolean literals
  provided using #Z3_solver_check_assumptions.

  \pre \c a must be a Boolean expression
  \pre \c p must be a Boolean constant (aka variable).

  \sa Z3_solver_assert
  \sa Z3_solver_reset

  def_API('Z3_solver_assert_and_track', VOID, (_in(CONTEXT), _in(SOLVER), _in(AST), _in(AST)))
*/
//go:linkname SolverAssertAndTrack C.Z3_solver_assert_and_track
func SolverAssertAndTrack(c Context, s Solver, a Ast, p Ast)

/**
  \brief load solver assertions from a file.

  \sa Z3_solver_from_string
  \sa Z3_solver_to_string

  def_API('Z3_solver_from_file', VOID, (_in(CONTEXT), _in(SOLVER), _in(STRING)))
*/
//go:linkname SolverFromFile C.Z3_solver_from_file
func SolverFromFile(c Context, s Solver, file_name String)

/**
  \brief load solver assertions from a string.

  \sa Z3_solver_from_file
  \sa Z3_solver_to_string

  def_API('Z3_solver_from_string', VOID, (_in(CONTEXT), _in(SOLVER), _in(STRING)))
*/
//go:linkname SolverFromString C.Z3_solver_from_string
func SolverFromString(c Context, s Solver, file_name String)

/**
  \brief Return the set of asserted formulas on the solver.

  def_API('Z3_solver_get_assertions', AST_VECTOR, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetAssertions C.Z3_solver_get_assertions
func SolverGetAssertions(c Context, s Solver) AstVector

/**
  \brief Return the set of units modulo model conversion.

  def_API('Z3_solver_get_units', AST_VECTOR, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetUnits C.Z3_solver_get_units
func SolverGetUnits(c Context, s Solver) AstVector

/**
  \brief Return the trail modulo model conversion, in order of decision level
  The decision level can be retrieved using \c Z3_solver_get_level based on the trail.

  def_API('Z3_solver_get_trail', AST_VECTOR, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetTrail C.Z3_solver_get_trail
func SolverGetTrail(c Context, s Solver) AstVector

/**
  \brief Return the set of non units in the solver state.

  def_API('Z3_solver_get_non_units', AST_VECTOR, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetNonUnits C.Z3_solver_get_non_units
func SolverGetNonUnits(c Context, s Solver) AstVector

/**
  \brief retrieve the decision depth of Boolean literals (variables or their negations).
  Assumes a check-sat call and no other calls (to extract models) have been invoked.

  def_API('Z3_solver_get_levels', VOID, (_in(CONTEXT), _in(SOLVER), _in(AST_VECTOR), _in(UINT), _in_array(3, UINT)))
*/
//go:linkname SolverGetLevels C.Z3_solver_get_levels
func SolverGetLevels(c Context, s Solver, literals AstVector, sz c.Uint, levels *c.Uint)

/**
  \brief retrieve the congruence closure root of an expression.
  The root is retrieved relative to the state where the solver was in when it completed.
  If it completed during a set of case splits, the congruence roots are relative to these case splits.
  That is, the congruences are not consequences but they are true under the current state.

  def_API('Z3_solver_congruence_root', AST, (_in(CONTEXT), _in(SOLVER), _in(AST)))
*/
//go:linkname SolverCongruenceRoot C.Z3_solver_congruence_root
func SolverCongruenceRoot(c Context, s Solver, a Ast) Ast

/**
  \brief retrieve the next expression in the congruence class. The set of congruent siblings form a cyclic list.
  Repeated calls on the siblings will result in returning to the original expression.

  def_API('Z3_solver_congruence_next', AST, (_in(CONTEXT), _in(SOLVER), _in(AST)))
*/
//go:linkname SolverCongruenceNext C.Z3_solver_congruence_next
func SolverCongruenceNext(c Context, s Solver, a Ast) Ast

/**
  \brief register a callback to that retrieves assumed, inferred and deleted clauses during search.

  \param c - context.
  \param s - solver object.
  \param user_context - a context used to maintain state for callbacks.
  \param on_clause_eh - a callback that is invoked by when a clause is
                          - asserted to the CDCL engine (corresponding to an input clause after pre-processing)
                          - inferred by CDCL(T) using either a SAT or theory conflict/propagation
                          - deleted by the CDCL(T) engine

  def_API('Z3_solver_register_on_clause', VOID, (_in(CONTEXT), _in(SOLVER), _in(VOID_PTR), _fnptr(Z3_on_clause_eh)))
*/
//go:linkname SolverRegisterOnClause C.Z3_solver_register_on_clause
func SolverRegisterOnClause(c Context, s Solver, user_context unsafe.Pointer, on_clause_eh OnClauseEh)

/**
  \brief register a user-propagator with the solver.

  \param c - context.
  \param s - solver object.
  \param user_context - a context used to maintain state for callbacks.
  \param push_eh - a callback invoked when scopes are pushed
  \param pop_eh - a callback invoked when scopes are popped
  \param fresh_eh - a solver may spawn new solvers internally. This callback is used to produce a fresh user_context to be associated with fresh solvers.

  def_API('Z3_solver_propagate_init', VOID, (_in(CONTEXT), _in(SOLVER), _in(VOID_PTR), _fnptr(Z3_push_eh), _fnptr(Z3_pop_eh), _fnptr(Z3_fresh_eh)))
*/
//go:linkname SolverPropagateInit C.Z3_solver_propagate_init
func SolverPropagateInit(c Context, s Solver, user_context unsafe.Pointer, push_eh PushEh, pop_eh PopEh, fresh_eh FreshEh)

/**
  \brief register a callback for when an expression is bound to a fixed value.
  The supported expression types are
  - Booleans
  - Bit-vectors

  def_API('Z3_solver_propagate_fixed', VOID, (_in(CONTEXT), _in(SOLVER), _fnptr(Z3_fixed_eh)))
*/
//go:linkname SolverPropagateFixed C.Z3_solver_propagate_fixed
func SolverPropagateFixed(c Context, s Solver, fixed_eh FixedEh)

/**
  \brief register a callback on final check.
  This provides freedom to the propagator to delay actions or implement a branch-and bound solver.
  The final check is invoked when all decision variables have been assigned by the solver.

  The \c final_eh callback takes as argument the original user_context that was used
  when calling \c Z3_solver_propagate_init, and it takes a callback context with the
  opaque type \c Z3_solver_callback.
  The callback context is passed as argument to invoke the \c Z3_solver_propagate_consequence function.
  The callback context can only be accessed (for propagation and for dynamically registering expressions) within a callback.
  If the callback context gets used for propagation or conflicts, those propagations take effect and
  may trigger new decision variables to be set.

  def_API('Z3_solver_propagate_final', VOID, (_in(CONTEXT), _in(SOLVER), _fnptr(Z3_final_eh)))
*/
//go:linkname SolverPropagateFinal C.Z3_solver_propagate_final
func SolverPropagateFinal(c Context, s Solver, final_eh FinalEh)

/**
  \brief register a callback on expression equalities.

  def_API('Z3_solver_propagate_eq', VOID, (_in(CONTEXT), _in(SOLVER), _fnptr(Z3_eq_eh)))
*/
//go:linkname SolverPropagateEq C.Z3_solver_propagate_eq
func SolverPropagateEq(c Context, s Solver, eq_eh EqEh)

/**
  \brief register a callback on expression dis-equalities.

  def_API('Z3_solver_propagate_diseq', VOID, (_in(CONTEXT), _in(SOLVER), _fnptr(Z3_eq_eh)))
*/
//go:linkname SolverPropagateDiseq C.Z3_solver_propagate_diseq
func SolverPropagateDiseq(c Context, s Solver, eq_eh EqEh)

/**
  \brief register a callback when a new expression with a registered function is used by the solver
  The registered function appears at the top level and is created using \ref Z3_solver_propagate_declare.

  def_API('Z3_solver_propagate_created', VOID, (_in(CONTEXT), _in(SOLVER), _fnptr(Z3_created_eh)))
*/
//go:linkname SolverPropagateCreated C.Z3_solver_propagate_created
func SolverPropagateCreated(c Context, s Solver, created_eh CreatedEh)

/**
  \brief register a callback when the solver decides to split on a registered expression.
  The callback may change the arguments by providing other values by calling \ref Z3_solver_next_split

  def_API('Z3_solver_propagate_decide', VOID, (_in(CONTEXT), _in(SOLVER), _fnptr(Z3_decide_eh)))
*/
//go:linkname SolverPropagateDecide C.Z3_solver_propagate_decide
func SolverPropagateDecide(c Context, s Solver, decide_eh DecideEh)

/**
    Sets the next (registered) expression to split on.
    The function returns false and ignores the given expression in case the expression is already assigned internally
    (due to relevancy propagation, this assignments might not have been reported yet by the fixed callback).
    In case the function is called in the decide callback, it overrides the currently selected variable and phase.

  def_API('Z3_solver_next_split', BOOL, (_in(CONTEXT), _in(SOLVER_CALLBACK), _in(AST), _in(UINT), _in(LBOOL)))
*/
//go:linkname SolverNextSplit C.Z3_solver_next_split
func SolverNextSplit(c Context, cb SolverCallback, t Ast, idx c.Uint, phase Lbool) bool

/**
    Create uninterpreted function declaration for the user propagator.
    When expressions using the function are created by the solver invoke a callback
    to \ref Z3_solver_propagate_created with arguments
    1. context and callback solve
    2. declared_expr: expression using function that was used as the top-level symbol
    3. declared_id: a unique identifier (unique within the current scope) to track the expression.

  def_API('Z3_solver_propagate_declare', FUNC_DECL, (_in(CONTEXT), _in(SYMBOL), _in(UINT), _in_array(2, SORT), _in(SORT)))
*/
//go:linkname SolverPropagateDeclare C.Z3_solver_propagate_declare
func SolverPropagateDeclare(c Context, name Symbol, n c.Uint, domain *Sort, range_ Sort) FuncDecl

/**
  \brief register an expression to propagate on with the solver.
  Only expressions of type Bool and type Bit-Vector can be registered for propagation.

  def_API('Z3_solver_propagate_register', VOID, (_in(CONTEXT), _in(SOLVER), _in(AST)))
*/
//go:linkname SolverPropagateRegister C.Z3_solver_propagate_register
func SolverPropagateRegister(c Context, s Solver, e Ast)

/**
  \brief register an expression to propagate on with the solver.
  Only expressions of type Bool and type Bit-Vector can be registered for propagation.
  Unlike \ref Z3_solver_propagate_register, this function takes a solver callback context
  as argument. It can be invoked during a callback to register new expressions.

  def_API('Z3_solver_propagate_register_cb', VOID, (_in(CONTEXT), _in(SOLVER_CALLBACK), _in(AST)))
*/
//go:linkname SolverPropagateRegisterCb C.Z3_solver_propagate_register_cb
func SolverPropagateRegisterCb(c Context, cb SolverCallback, e Ast)

/**
  \brief propagate a consequence based on fixed values and equalities.
  A client may invoke it during the \c propagate_fixed, \c propagate_eq, \c propagate_diseq, and \c propagate_final callbacks.
  The callback adds a propagation consequence based on the fixed values passed \c ids and equalities \c eqs based on parameters \c lhs, \c rhs.

  The solver might discard the propagation in case it is true in the current state.
  The function returns false in this case; otw. the function returns true.
  At least one propagation in the final callback has to return true in order to
  prevent the solver from finishing.

  Assume the callback has the signature: \c propagate_consequence_eh(context, solver_cb, num_ids, ids, num_eqs, lhs, rhs, consequence).
  \param c - context
  \param solver_cb - solver callback
  \param num_ids - number of fixed terms used as premise to propagation
  \param ids - array of length \c num_ids containing terms that are fixed in the current scope
  \param num_eqs - number of equalities used as premise to propagation
  \param lhs - left side of equalities
  \param rhs - right side of equalities
  \param consequence - consequence to propagate. It is typically an atomic formula, but it can be an arbitrary formula.

  def_API('Z3_solver_propagate_consequence', BOOL, (_in(CONTEXT), _in(SOLVER_CALLBACK), _in(UINT), _in_array(2, AST), _in(UINT), _in_array(4, AST), _in_array(4, AST), _in(AST)))
*/
//go:linkname SolverPropagateConsequence C.Z3_solver_propagate_consequence
func SolverPropagateConsequence(c Context, cb SolverCallback, num_fixed c.Uint, fixed *Ast, num_eqs c.Uint, eq_lhs *Ast, eq_rhs *Ast, conseq Ast) bool

/**
  \brief Check whether the assertions in a given solver are consistent or not.

  The function #Z3_solver_get_model retrieves a model if the
  assertions is satisfiable (i.e., the result is \c
  Z3_L_TRUE) and model construction is enabled.
  Note that if the call returns \c Z3_L_UNDEF, Z3 does not
  ensure that calls to #Z3_solver_get_model succeed and any models
  produced in this case are not guaranteed to satisfy the assertions.

  The function #Z3_solver_get_proof retrieves a proof if proof
  generation was enabled when the context was created, and the
  assertions are unsatisfiable (i.e., the result is \c Z3_L_FALSE).

  \sa Z3_solver_check_assumptions

  def_API('Z3_solver_check', LBOOL, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverCheck C.Z3_solver_check
func SolverCheck(c Context, s Solver) Lbool

/**
  \brief Check whether the assertions in the given solver and
  optional assumptions are consistent or not.

  The function #Z3_solver_get_unsat_core retrieves the subset of the
  assumptions used in the unsatisfiability proof produced by Z3.

  \sa Z3_solver_check

  def_API('Z3_solver_check_assumptions', LBOOL, (_in(CONTEXT), _in(SOLVER), _in(UINT), _in_array(2, AST)))
*/
//go:linkname SolverCheckAssumptions C.Z3_solver_check_assumptions
func SolverCheckAssumptions(c Context, s Solver, num_assumptions c.Uint, assumptions *Ast) Lbool

/**
  \brief Retrieve congruence class representatives for terms.

  The function can be used for relying on Z3 to identify equal terms under the current
  set of assumptions. The array of terms and array of class identifiers should have
  the same length. The class identifiers are numerals that are assigned to the same
  value for their corresponding terms if the current context forces the terms to be
  equal. You cannot deduce that terms corresponding to different numerals must be all different,
  (especially when using non-convex theories).
  All implied equalities are returned by this call.
  This means that two terms map to the same class identifier if and only if
  the current context implies that they are equal.

  A side-effect of the function is a satisfiability check on the assertions on the solver that is passed in.
  The function return \c Z3_L_FALSE if the current assertions are not satisfiable.

  def_API('Z3_get_implied_equalities', LBOOL, (_in(CONTEXT), _in(SOLVER), _in(UINT), _in_array(2, AST), _out_array(2, UINT)))
*/
//go:linkname GetImpliedEqualities C.Z3_get_implied_equalities
func GetImpliedEqualities(c Context, s Solver, num_terms c.Uint, terms *Ast, class_ids *c.Uint) Lbool

/**
  \brief retrieve consequences from solver that determine values of the supplied function symbols.

  def_API('Z3_solver_get_consequences', LBOOL, (_in(CONTEXT), _in(SOLVER), _in(AST_VECTOR), _in(AST_VECTOR), _in(AST_VECTOR)))
*/
//go:linkname SolverGetConsequences C.Z3_solver_get_consequences
func SolverGetConsequences(c Context, s Solver, assumptions AstVector, variables AstVector, consequences AstVector) Lbool

/**
  \brief extract a next cube for a solver. The last cube is the constant \c true or \c false.
  The number of (non-constant) cubes is by default 1. For the sat solver cubing is controlled
  using parameters sat.lookahead.cube.cutoff and sat.lookahead.cube.fraction.

  The third argument is a vector of variables that may be used for cubing.
  The contents of the vector is only used in the first call. The initial list of variables
  is used in subsequent calls until it returns the unsatisfiable cube.
  The vector is modified to contain a set of Autarky variables that occur in clauses that
  are affected by the (last literal in the) cube. These variables could be used by a different
  cuber (on a different solver object) for further recursive cubing.

  The last argument is a backtracking level. It instructs the cube process to backtrack below
  the indicated level for the next cube.

  def_API('Z3_solver_cube', AST_VECTOR, (_in(CONTEXT), _in(SOLVER), _in(AST_VECTOR), _in(UINT)))
*/
//go:linkname SolverCube C.Z3_solver_cube
func SolverCube(c Context, s Solver, vars AstVector, backtrack_level c.Uint) AstVector

/**
  \brief Retrieve the model for the last #Z3_solver_check or #Z3_solver_check_assumptions

  The error handler is invoked if a model is not available because
  the commands above were not invoked for the given solver, or if the result was \c Z3_L_FALSE.

  def_API('Z3_solver_get_model', MODEL, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetModel C.Z3_solver_get_model
func SolverGetModel(c Context, s Solver) Model

/**
  \brief Retrieve the proof for the last #Z3_solver_check or #Z3_solver_check_assumptions

  The error handler is invoked if proof generation is not enabled,
  or if the commands above were not invoked for the given solver,
  or if the result was different from \c Z3_L_FALSE.

  def_API('Z3_solver_get_proof', AST, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetProof C.Z3_solver_get_proof
func SolverGetProof(c Context, s Solver) Ast

/**
  \brief Retrieve the unsat core for the last #Z3_solver_check_assumptions
  The unsat core is a subset of the assumptions \c a.

  By default, the unsat core will not be minimized. Generation of a minimized
  unsat core can be enabled via the `"sat.core.minimize"` and `"smt.core.minimize"`
  settings for SAT and SMT cores respectively. Generation of minimized unsat cores
  will be more expensive.

  def_API('Z3_solver_get_unsat_core', AST_VECTOR, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetUnsatCore C.Z3_solver_get_unsat_core
func SolverGetUnsatCore(c Context, s Solver) AstVector

/**
  \brief Return a brief justification for an "unknown" result (i.e., \c Z3_L_UNDEF) for
  the commands #Z3_solver_check and #Z3_solver_check_assumptions

  def_API('Z3_solver_get_reason_unknown', STRING, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetReasonUnknown C.Z3_solver_get_reason_unknown
func SolverGetReasonUnknown(c Context, s Solver) String

/**
  \brief Return statistics for the given solver.

  \remark User must use #Z3_stats_inc_ref and #Z3_stats_dec_ref to manage Z3_stats objects.

  def_API('Z3_solver_get_statistics', STATS, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverGetStatistics C.Z3_solver_get_statistics
func SolverGetStatistics(c Context, s Solver) Stats

/**
  \brief Convert a solver into a string.

  \sa Z3_solver_from_file
  \sa Z3_solver_from_string

  def_API('Z3_solver_to_string', STRING, (_in(CONTEXT), _in(SOLVER)))
*/
//go:linkname SolverToString C.Z3_solver_to_string
func SolverToString(c Context, s Solver) String

/**
  \brief Convert a solver into a DIMACS formatted string.
  \sa Z3_goal_to_diamcs_string for requirements.

  def_API('Z3_solver_to_dimacs_string', STRING, (_in(CONTEXT), _in(SOLVER), _in(BOOL)))
*/
//go:linkname SolverToDimacsString C.Z3_solver_to_dimacs_string
func SolverToDimacsString(c Context, s Solver, include_names bool) String

/**
  \brief Convert a statistics into a string.

  def_API('Z3_stats_to_string', STRING, (_in(CONTEXT), _in(STATS)))
*/
//go:linkname StatsToString C.Z3_stats_to_string
func StatsToString(c Context, s Stats) String

/**
  \brief Increment the reference counter of the given statistics object.

  def_API('Z3_stats_inc_ref', VOID, (_in(CONTEXT), _in(STATS)))
*/
//go:linkname StatsIncRef C.Z3_stats_inc_ref
func StatsIncRef(c Context, s Stats)

/**
  \brief Decrement the reference counter of the given statistics object.

  def_API('Z3_stats_dec_ref', VOID, (_in(CONTEXT), _in(STATS)))
*/
//go:linkname StatsDecRef C.Z3_stats_dec_ref
func StatsDecRef(c Context, s Stats)

/**
  \brief Return the number of statistical data in \c s.

  def_API('Z3_stats_size', UINT, (_in(CONTEXT), _in(STATS)))
*/
//go:linkname StatsSize C.Z3_stats_size
func StatsSize(c Context, s Stats) c.Uint

/**
  \brief Return the key (a string) for a particular statistical data.

  \pre idx < Z3_stats_size(c, s)

  def_API('Z3_stats_get_key', STRING, (_in(CONTEXT), _in(STATS), _in(UINT)))
*/
//go:linkname StatsGetKey C.Z3_stats_get_key
func StatsGetKey(c Context, s Stats, idx c.Uint) String

/**
  \brief Return \c true if the given statistical data is a unsigned integer.

  \pre idx < Z3_stats_size(c, s)

  def_API('Z3_stats_is_uint', BOOL, (_in(CONTEXT), _in(STATS), _in(UINT)))
*/
//go:linkname StatsIsUint C.Z3_stats_is_uint
func StatsIsUint(c Context, s Stats, idx c.Uint) bool

/**
  \brief Return \c true if the given statistical data is a double.

  \pre idx < Z3_stats_size(c, s)

  def_API('Z3_stats_is_double', BOOL, (_in(CONTEXT), _in(STATS), _in(UINT)))
*/
//go:linkname StatsIsDouble C.Z3_stats_is_double
func StatsIsDouble(c Context, s Stats, idx c.Uint) bool

/**
  \brief Return the unsigned value of the given statistical data.

  \pre idx < Z3_stats_size(c, s) && Z3_stats_is_uint(c, s)

  def_API('Z3_stats_get_uint_value', UINT, (_in(CONTEXT), _in(STATS), _in(UINT)))
*/
//go:linkname StatsGetUintValue C.Z3_stats_get_uint_value
func StatsGetUintValue(c Context, s Stats, idx c.Uint) c.Uint

/**
  \brief Return the double value of the given statistical data.

  \pre idx < Z3_stats_size(c, s) && Z3_stats_is_double(c, s)

  def_API('Z3_stats_get_double_value', DOUBLE, (_in(CONTEXT), _in(STATS), _in(UINT)))
*/
//go:linkname StatsGetDoubleValue C.Z3_stats_get_double_value
func StatsGetDoubleValue(c Context, s Stats, idx c.Uint) float64

/**
  \brief Return the estimated allocated memory in bytes.

  def_API('Z3_get_estimated_alloc_size', UINT64, ())
*/
//go:linkname GetEstimatedAllocSize C.Z3_get_estimated_alloc_size
func GetEstimatedAllocSize() uint64
