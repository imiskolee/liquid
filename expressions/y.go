//line expressions.y:2
package expressions

import __yyfmt__ "fmt"

//line expressions.y:2
import (
	"fmt"
	"github.com/osteele/liquid/generics"
)

func init() {
	// This allows adding and removing references to fmt in the rules below,
	// without having to edit the import statement to avoid erorrs each time.
	_ = fmt.Sprint("")
}

//line expressions.y:15
type yySymType struct {
	yys           int
	name          string
	val           interface{}
	f             func(Context) interface{}
	loopmods      LoopModifiers
	filter_params []valueFn
}

const LITERAL = 57346
const IDENTIFIER = 57347
const KEYWORD = 57348
const RELATION = 57349
const ASSIGN = 57350
const LOOP = 57351
const EQ = 57352
const NEQ = 57353
const FOR = 57354
const IN = 57355
const AND = 57356
const OR = 57357
const CONTAINS = 57358

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LITERAL",
	"IDENTIFIER",
	"KEYWORD",
	"RELATION",
	"ASSIGN",
	"LOOP",
	"EQ",
	"NEQ",
	"FOR",
	"IN",
	"AND",
	"OR",
	"CONTAINS",
	"'.'",
	"'|'",
	"'<'",
	"'>'",
	"';'",
	"'='",
	"'['",
	"']'",
	"'('",
	"')'",
	"','",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 62

var yyAct = [...]int{

	7, 12, 13, 46, 6, 8, 9, 27, 18, 3,
	4, 17, 18, 37, 19, 43, 8, 9, 19, 28,
	32, 33, 34, 35, 36, 31, 10, 17, 39, 39,
	44, 42, 38, 40, 16, 20, 21, 10, 48, 49,
	2, 14, 18, 51, 22, 23, 5, 50, 19, 12,
	13, 24, 29, 30, 47, 1, 11, 45, 41, 25,
	26, 15,
}
var yyPact = [...]int{

	1, -1000, 35, 36, 29, -1000, -7, 25, -1000, -1000,
	12, -1000, 12, 12, -15, -1000, 6, 47, 20, 12,
	12, 12, 12, 12, -13, -1000, -1000, 12, 12, -1000,
	12, -1000, -9, -5, -5, -5, -5, -1000, 9, -5,
	-7, -24, -5, -1000, -1000, 33, 12, -1000, -1000, 39,
	-5, -1000,
}
var yyPgo = [...]int{

	0, 0, 46, 4, 40, 61, 58, 57, 55,
}
var yyR1 = [...]int{

	0, 8, 8, 8, 5, 7, 7, 7, 1, 1,
	1, 1, 1, 3, 3, 3, 6, 6, 2, 2,
	2, 2, 2, 4, 4, 4,
}
var yyR2 = [...]int{

	0, 2, 5, 2, 5, 0, 2, 3, 1, 1,
	3, 4, 3, 1, 3, 4, 1, 3, 1, 3,
	3, 3, 3, 1, 3, 3,
}
var yyChk = [...]int{

	-1000, -8, -4, 8, 9, -2, -3, -1, 4, 5,
	25, 21, 14, 15, 5, -5, 5, 18, 17, 23,
	10, 11, 19, 20, -4, -2, -2, 22, 13, 5,
	6, 5, -1, -1, -1, -1, -1, 26, -3, -1,
	-3, -6, -1, 24, 21, -7, 27, 21, 5, 6,
	-1, 4,
}
var yyDef = [...]int{

	0, -2, 0, 0, 0, 23, 18, 13, 8, 9,
	0, 1, 0, 0, 0, 3, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 24, 25, 0, 0, 14,
	0, 10, 0, 19, 20, 21, 22, 12, 0, 13,
	5, 15, 16, 11, 2, 0, 0, 4, 6, 0,
	17, 7,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	25, 26, 3, 3, 27, 3, 17, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 21,
	19, 22, 20, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 23, 3, 24, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 18,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line expressions.y:33
		{
			yylex.(*lexer).val = yyDollar[1].f
		}
	case 2:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line expressions.y:34
		{
			name, expr := yyDollar[2].name, yyDollar[4].f
			yylex.(*lexer).val = func(ctx Context) interface{} {
				ctx.Set(name, expr(ctx))
				return nil
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line expressions.y:41
		{
			yylex.(*lexer).val = yyDollar[2].f
		}
	case 4:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line expressions.y:44
		{
			name, expr, mods := yyDollar[1].name, yyDollar[3].f, yyDollar[4].loopmods
			yyVAL.f = func(ctx Context) interface{} {
				return &Loop{name, expr(ctx), mods}
			}
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line expressions.y:52
		{
			yyVAL.loopmods = LoopModifiers{}
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line expressions.y:53
		{
			switch yyDollar[2].name {
			case "reversed":
				yyDollar[1].loopmods.Reversed = true
			default:
				panic(ParseError(fmt.Sprintf("undefined loop modifier: %s", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:62
		{ // TODO can this be a variable?
			switch yyDollar[2].name {
			case "limit":
				limit, ok := yyDollar[3].val.(int)
				if !ok {
					panic(ParseError(fmt.Sprintf("loop limit must an integer")))
				}
				yyDollar[1].loopmods.Limit = &limit
			case "offset":
				offset, ok := yyDollar[3].val.(int)
				if !ok {
					panic(ParseError(fmt.Sprintf("loop offset must an integer")))
				}
				yyDollar[1].loopmods.Offset = offset
			default:
				panic(ParseError(fmt.Sprintf("undefined loop modifier: %s", yyDollar[2].name)))
			}
			yyVAL.loopmods = yyDollar[1].loopmods
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line expressions.y:84
		{
			val := yyDollar[1].val
			yyVAL.f = func(_ Context) interface{} { return val }
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line expressions.y:85
		{
			name := yyDollar[1].name
			yyVAL.f = func(ctx Context) interface{} { return ctx.Get(name) }
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:86
		{
			yyVAL.f = makeObjectPropertyEvaluator(yyDollar[1].f, yyDollar[3].name)
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line expressions.y:87
		{
			yyVAL.f = makeIndexEvaluator(yyDollar[1].f, yyDollar[3].f)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:88
		{
			yyVAL.f = yyDollar[2].f
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:93
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, nil)
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line expressions.y:94
		{
			yyVAL.f = makeFilter(yyDollar[1].f, yyDollar[3].name, yyDollar[4].filter_params)
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line expressions.y:98
		{
			yyVAL.filter_params = []valueFn{yyDollar[1].f}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:100
		{
			yyVAL.filter_params = append(yyDollar[1].filter_params, yyDollar[3].f)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:104
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return generics.Equal(a, b)
			}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:111
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return !generics.Equal(a, b)
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:117
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return generics.Less(a, b)
			}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:124
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				a, b := fa(ctx), fb(ctx)
				return generics.Less(b, a)
			}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:135
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				return generics.IsTrue(fa(ctx)) && generics.IsTrue(fb(ctx))
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line expressions.y:141
		{
			fa, fb := yyDollar[1].f, yyDollar[3].f
			yyVAL.f = func(ctx Context) interface{} {
				return generics.IsTrue(fa(ctx)) || generics.IsTrue(fb(ctx))
			}
		}
	}
	goto yystack /* stack new state and value */
}
