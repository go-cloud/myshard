//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE        = []byte("share")
	MODE         = []byte("mode")
	IF_BYTES     = []byte("if")
	VALUES_BYTES = []byte("values")
)

//line sql.y:31
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const KEYRANGE = 57377
const ID = 57378
const STRING = 57379
const NUMBER = 57380
const VALUE_ARG = 57381
const LIST_ARG = 57382
const COMMENT = 57383
const LE = 57384
const GE = 57385
const NE = 57386
const NULL_SAFE_EQUAL = 57387
const UNION = 57388
const MINUS = 57389
const EXCEPT = 57390
const INTERSECT = 57391
const JOIN = 57392
const STRAIGHT_JOIN = 57393
const LEFT = 57394
const RIGHT = 57395
const INNER = 57396
const OUTER = 57397
const CROSS = 57398
const NATURAL = 57399
const USE = 57400
const FORCE = 57401
const ON = 57402
const OR = 57403
const AND = 57404
const NOT = 57405
const UNARY = 57406
const CASE = 57407
const WHEN = 57408
const THEN = 57409
const ELSE = 57410
const END = 57411
const CREATE = 57412
const ALTER = 57413
const DROP = 57414
const RENAME = 57415
const ANALYZE = 57416
const TABLE = 57417
const INDEX = 57418
const VIEW = 57419
const TO = 57420
const IGNORE = 57421
const IF = 57422
const UNIQUE = 57423
const USING = 57424
const SHOW = 57425
const DESCRIBE = 57426
const EXPLAIN = 57427
const BEGIN = 57428
const COMMIT = 57429
const ROLLBACK = 57430
const NAMES = 57431
const REPLACE = 57432

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"KEYRANGE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	" (",
	" =",
	" <",
	" >",
	" ~",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	" ,",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	" &",
	" |",
	" ^",
	" +",
	" -",
	" *",
	" /",
	" %",
	" .",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 214
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 622

var yyAct = []int{

	105, 311, 172, 381, 71, 96, 348, 263, 102, 175,
	103, 213, 303, 254, 91, 101, 224, 191, 256, 149,
	148, 390, 73, 174, 3, 114, 92, 390, 390, 186,
	143, 245, 361, 276, 277, 278, 279, 280, 58, 281,
	282, 36, 37, 38, 39, 76, 75, 309, 143, 80,
	143, 245, 83, 199, 74, 46, 87, 48, 243, 62,
	359, 49, 86, 113, 328, 330, 119, 97, 59, 60,
	358, 357, 392, 78, 76, 110, 111, 112, 391, 389,
	132, 337, 334, 332, 178, 287, 255, 270, 117, 140,
	51, 82, 52, 79, 329, 145, 244, 136, 308, 297,
	57, 295, 246, 176, 53, 171, 173, 177, 54, 55,
	56, 115, 116, 72, 149, 148, 339, 147, 120, 133,
	129, 365, 135, 185, 75, 131, 124, 75, 189, 341,
	195, 194, 74, 118, 181, 74, 156, 157, 158, 159,
	160, 161, 162, 163, 148, 81, 97, 219, 195, 193,
	255, 68, 301, 223, 221, 222, 231, 232, 354, 235,
	236, 237, 238, 239, 240, 241, 242, 233, 217, 196,
	218, 210, 159, 160, 161, 162, 163, 226, 304, 209,
	126, 247, 97, 97, 220, 205, 149, 148, 75, 75,
	266, 304, 259, 139, 249, 251, 74, 261, 265, 252,
	267, 161, 162, 163, 203, 85, 356, 262, 206, 322,
	258, 234, 75, 320, 323, 355, 272, 326, 321, 122,
	74, 325, 125, 324, 192, 126, 192, 286, 247, 271,
	273, 142, 290, 291, 258, 288, 245, 366, 343, 268,
	298, 217, 141, 18, 128, 289, 216, 376, 294, 36,
	37, 38, 39, 97, 226, 375, 215, 227, 202, 204,
	201, 302, 88, 225, 374, 296, 211, 300, 306, 274,
	310, 126, 307, 108, 216, 178, 183, 143, 113, 188,
	187, 119, 182, 180, 215, 179, 318, 319, 109, 95,
	110, 111, 112, 188, 113, 336, 121, 81, 76, 100,
	217, 217, 333, 117, 340, 123, 110, 111, 112, 331,
	75, 315, 345, 285, 338, 346, 349, 146, 344, 90,
	314, 208, 99, 207, 190, 69, 115, 116, 93, 137,
	284, 134, 130, 120, 81, 127, 89, 84, 360, 276,
	277, 278, 279, 280, 362, 281, 282, 387, 118, 363,
	342, 67, 293, 350, 364, 394, 247, 197, 371, 370,
	373, 18, 138, 372, 228, 388, 229, 230, 378, 349,
	65, 264, 380, 379, 63, 382, 382, 382, 75, 383,
	384, 312, 385, 250, 257, 108, 74, 353, 313, 352,
	113, 395, 317, 119, 192, 396, 70, 397, 393, 377,
	109, 95, 110, 111, 112, 18, 19, 20, 21, 18,
	41, 100, 14, 17, 335, 117, 156, 157, 158, 159,
	160, 161, 162, 163, 156, 157, 158, 159, 160, 161,
	162, 163, 16, 22, 99, 15, 198, 47, 115, 116,
	93, 269, 200, 50, 77, 120, 260, 386, 292, 18,
	156, 157, 158, 159, 160, 161, 162, 163, 367, 347,
	118, 351, 316, 299, 108, 184, 253, 107, 104, 113,
	106, 248, 119, 305, 150, 98, 327, 214, 275, 109,
	76, 110, 111, 112, 212, 23, 24, 26, 25, 27,
	100, 94, 283, 144, 117, 64, 35, 66, 28, 29,
	30, 32, 33, 34, 108, 31, 13, 12, 11, 113,
	10, 9, 119, 99, 8, 18, 7, 115, 116, 109,
	76, 110, 111, 112, 120, 368, 369, 6, 5, 4,
	100, 2, 1, 0, 117, 113, 0, 0, 119, 118,
	0, 0, 0, 0, 0, 0, 76, 110, 111, 112,
	0, 0, 0, 99, 0, 0, 178, 115, 116, 0,
	117, 0, 0, 0, 120, 151, 155, 153, 154, 156,
	157, 158, 159, 160, 161, 162, 163, 40, 0, 118,
	0, 0, 0, 115, 116, 0, 167, 168, 169, 170,
	120, 164, 165, 166, 0, 0, 0, 42, 43, 44,
	45, 0, 0, 0, 0, 118, 0, 0, 0, 61,
	0, 0, 0, 152, 156, 157, 158, 159, 160, 161,
	162, 163,
}
var yyPact = []int{

	400, -1000, -1000, 198, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -35, -2, 14, 18, 10, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 404, 357, -1000, -1000, -1000,
	352, -1000, 322, 289, 387, 9, -22, 2, 261, -1000,
	1, 261, -1000, 301, -33, 261, -33, 300, -1000, -1000,
	-1000, 290, -1000, -1000, 253, -1000, 255, 289, 272, 48,
	289, 170, 299, -1000, 197, -1000, 42, 296, 56, 261,
	-1000, -1000, 295, -1000, 4, 293, 342, 127, 261, -1000,
	289, 222, -1000, -1000, 298, 39, 119, 544, -1000, 484,
	444, -1000, -1000, -1000, 38, 239, 237, -1000, 236, 230,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	38, -1000, 247, 262, 288, 384, 262, -1000, 38, 261,
	-1000, 337, -44, -1000, 172, -1000, 287, -1000, -1000, 285,
	-1000, 233, 210, 253, -1000, -1000, 261, 109, 484, 484,
	38, 217, 343, 38, 38, 142, 38, 38, 38, 38,
	38, 38, 38, 38, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 544, -48, -10, -4, 544, -1000, 510, 365,
	253, -1000, 404, 269, 5, 354, 356, 262, 262, 216,
	-1000, 358, 484, -1000, 354, -1000, -1000, -1000, 124, 261,
	-1000, -6, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	356, 262, 214, 283, 294, 238, 7, -1000, -1000, -1000,
	-1000, -1000, 76, 354, -1000, 510, -1000, -1000, 217, 38,
	38, 354, 380, -1000, 327, 99, 99, 99, 126, 126,
	-1000, -1000, -1000, -1000, -1000, 38, -1000, 354, -1000, -5,
	253, -7, 185, 69, -1000, 484, 112, 229, 198, 125,
	-8, -1000, 358, 366, 374, 119, 284, -1000, -1000, 275,
	-1000, -1000, 170, 381, 210, 210, -1000, -1000, 157, 153,
	167, 165, 161, 0, -1000, 273, -23, 266, -24, -1000,
	354, 346, 38, -1000, 354, -1000, -25, -1000, 269, 32,
	-1000, 38, 47, -1000, 320, 183, -1000, -1000, -1000, 262,
	366, -1000, 38, 38, -1000, -1000, 377, 373, 283, 92,
	-1000, 159, -1000, 150, -1000, -1000, -1000, -1000, -20, -21,
	-31, -1000, -1000, -1000, -1000, 38, 354, -1000, -74, -1000,
	354, 38, 318, 229, -1000, -1000, 66, 182, -1000, 499,
	-1000, 358, 484, 38, 484, -1000, -1000, 218, 209, 201,
	354, -1000, 354, 392, -1000, 38, 38, -1000, -1000, -1000,
	366, 119, 181, 119, 261, 261, 261, 262, 354, -1000,
	331, -27, -1000, -28, -34, 170, -1000, 391, 334, -1000,
	261, -1000, -1000, -1000, 261, -1000, 261, -1000,
}
var yyPgo = []int{

	0, 532, 531, 23, 529, 528, 527, 516, 514, 511,
	510, 508, 507, 506, 577, 497, 496, 495, 14, 26,
	493, 492, 491, 484, 11, 478, 477, 151, 476, 3,
	17, 5, 475, 474, 18, 15, 2, 16, 9, 473,
	10, 470, 25, 468, 8, 467, 466, 13, 465, 463,
	462, 461, 7, 459, 6, 458, 1, 447, 29, 446,
	12, 4, 22, 205, 444, 443, 442, 441, 437, 436,
	0, 38, 435, 432, 413, 412, 410,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3,
	4, 4, 75, 75, 5, 6, 7, 7, 72, 73,
	74, 8, 8, 8, 9, 9, 9, 10, 11, 11,
	11, 12, 13, 13, 13, 76, 14, 15, 15, 16,
	16, 16, 16, 16, 17, 17, 18, 18, 19, 19,
	19, 22, 22, 20, 20, 20, 23, 23, 24, 24,
	24, 24, 21, 21, 21, 25, 25, 25, 25, 25,
	25, 25, 25, 25, 26, 26, 26, 27, 27, 28,
	28, 28, 28, 29, 29, 30, 30, 31, 31, 31,
	31, 31, 32, 32, 32, 32, 32, 32, 32, 32,
	32, 32, 32, 33, 33, 33, 33, 33, 33, 33,
	37, 37, 37, 42, 38, 38, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 41, 41, 43, 43, 43, 45, 48,
	48, 46, 46, 47, 49, 49, 44, 44, 35, 35,
	35, 35, 50, 50, 51, 51, 52, 52, 53, 53,
	54, 55, 55, 55, 56, 56, 56, 57, 57, 57,
	58, 58, 59, 59, 60, 60, 34, 34, 39, 39,
	40, 40, 61, 61, 62, 63, 63, 64, 64, 65,
	65, 66, 66, 66, 66, 66, 67, 67, 68, 68,
	69, 69, 70, 71,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 4, 12, 3,
	7, 7, 6, 6, 8, 7, 3, 4, 1, 1,
	1, 5, 8, 4, 6, 7, 4, 5, 4, 5,
	5, 3, 2, 2, 2, 0, 2, 0, 2, 1,
	2, 1, 1, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 3, 3,
	3, 5, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 1, 3, 1, 1, 3, 0,
	5, 5, 5, 1, 3, 0, 2, 1, 3, 3,
	2, 3, 3, 3, 4, 3, 4, 5, 6, 3,
	4, 2, 6, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 1, 3, 1, 3, 1, 1, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 3, 4,
	5, 4, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 2, 1, 1, 3,
	3, 1, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -12, -13, -75, -72, -73, -74, 5, 6,
	7, 8, 33, 85, 86, 88, 87, 89, 98, 99,
	100, 105, 101, 102, 103, -16, 51, 52, 53, 54,
	-14, -76, -14, -14, -14, -14, 90, -68, 92, 96,
	-65, 92, 94, 90, 90, 91, 92, 90, -71, -71,
	-71, -14, -3, 17, -17, 18, -15, 29, -27, 36,
	9, -61, 104, -62, -44, -70, 36, -64, 95, 91,
	-70, 36, 90, -70, 36, -63, 95, -70, -63, 36,
	29, -18, -19, 75, -22, 36, -31, -36, -32, 69,
	46, -35, -44, -40, -43, -70, -41, -45, 20, 35,
	37, 38, 39, 25, -42, 73, 74, 50, 95, 28,
	80, 41, -27, 33, 78, -27, 55, 36, 47, 78,
	36, 69, -70, -71, 36, -71, 93, 36, 20, 66,
	-70, -27, 9, 55, -20, -70, 19, 78, 68, 67,
	-33, 21, 69, 23, 24, 22, 70, 71, 72, 73,
	74, 75, 76, 77, 47, 48, 49, 42, 43, 44,
	45, -31, -36, -31, -3, -38, -36, -36, 46, 46,
	46, -42, 46, 46, -48, -36, -58, 33, 46, -61,
	36, -30, 10, -62, -36, -70, -71, 20, -69, 97,
	-66, 88, 86, 32, 87, 13, 36, 36, 36, -71,
	-58, 33, -23, -24, -26, 46, 36, -42, -19, -70,
	75, -31, -31, -36, -37, 46, -42, 40, 21, 23,
	24, -36, -36, 25, 69, -36, -36, -36, -36, -36,
	-36, -36, -36, 106, 106, 55, 106, -36, 106, -18,
	18, -18, -35, -46, -47, 81, -34, 28, -3, -61,
	-59, -44, -30, -52, 13, -31, 66, -70, -71, -67,
	93, -34, -61, -30, 55, -25, 56, 57, 58, 59,
	60, 62, 63, -21, 36, 19, -24, 78, -38, -37,
	-36, -36, 68, 25, -36, 106, -18, 106, 55, -49,
	-47, 83, -31, -60, 66, -39, -40, -60, 106, 55,
	-52, -56, 15, 14, 36, 36, -50, 11, -24, -24,
	56, 61, 56, 61, 56, 56, 56, -28, 64, 94,
	65, 36, 106, 36, 106, 68, -36, 106, -35, 84,
	-36, 82, 30, 55, -44, -56, -36, -53, -54, -36,
	-71, -51, 12, 14, 66, 56, 56, 91, 91, 91,
	-36, 106, -36, 31, -40, 55, 55, -55, 26, 27,
	-52, -31, -38, -31, 46, 46, 46, 7, -36, -54,
	-56, -29, -70, -29, -29, -61, -57, 16, 34, 106,
	55, 106, 106, 7, 21, -70, -70, -70,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 45, 45,
	45, 45, 45, 208, 199, 0, 0, 0, 213, 213,
	213, 45, 28, 29, 30, 0, 49, 51, 52, 53,
	54, 47, 0, 0, 0, 0, 197, 0, 0, 209,
	0, 0, 200, 0, 195, 0, 195, 0, 42, 43,
	44, 0, 19, 50, 0, 55, 46, 0, 0, 87,
	0, 26, 0, 192, 0, 156, 212, 0, 0, 0,
	213, 212, 0, 213, 0, 0, 0, 0, 0, 41,
	0, 17, 56, 58, 63, 212, 61, 62, 97, 0,
	0, 126, 127, 128, 0, 156, 0, 142, 0, 0,
	158, 159, 160, 161, 191, 145, 146, 147, 143, 144,
	149, 48, 180, 0, 0, 95, 0, 27, 0, 0,
	213, 0, 210, 33, 0, 36, 0, 38, 196, 0,
	213, 180, 0, 0, 59, 64, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 113, 114, 115, 116, 117, 118,
	119, 100, 0, 0, 0, 0, 124, 137, 0, 0,
	0, 111, 0, 0, 0, 150, 0, 0, 0, 95,
	88, 166, 0, 193, 194, 157, 31, 198, 0, 0,
	213, 206, 201, 202, 203, 204, 205, 37, 39, 40,
	0, 0, 95, 66, 72, 0, 84, 86, 57, 65,
	60, 98, 99, 102, 103, 0, 121, 122, 0, 0,
	0, 105, 0, 109, 0, 129, 130, 131, 132, 133,
	134, 135, 136, 101, 123, 0, 190, 124, 138, 0,
	0, 0, 0, 154, 151, 0, 184, 0, 187, 184,
	0, 182, 166, 174, 0, 96, 0, 211, 34, 0,
	207, 22, 23, 162, 0, 0, 75, 76, 0, 0,
	0, 0, 0, 89, 73, 0, 0, 0, 0, 104,
	106, 0, 0, 110, 125, 139, 0, 141, 0, 0,
	152, 0, 0, 20, 0, 186, 188, 21, 181, 0,
	174, 25, 0, 0, 213, 35, 164, 0, 67, 70,
	77, 0, 79, 0, 81, 82, 83, 68, 0, 0,
	0, 74, 69, 85, 120, 0, 107, 140, 0, 148,
	155, 0, 0, 0, 183, 24, 175, 167, 168, 171,
	32, 166, 0, 0, 0, 78, 80, 0, 0, 0,
	108, 112, 153, 0, 189, 0, 0, 170, 172, 173,
	174, 165, 163, 71, 0, 0, 0, 0, 176, 169,
	177, 0, 93, 0, 0, 185, 18, 0, 0, 90,
	0, 91, 92, 178, 0, 94, 0, 179,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 77, 70, 3,
	46, 106, 75, 73, 55, 74, 78, 76, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	48, 47, 49, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 72, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 71, 3, 50,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 51, 52, 53, 54, 56, 57,
	58, 59, 60, 61, 62, 63, 64, 65, 66, 67,
	68, 69, 79, 80, 81, 82, 83, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line sql.y:163
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:169
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		//line sql.y:189
		{
			yyVAL.selStmt = &SimpleSelect{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 18:
		//line sql.y:193
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 19:
		//line sql.y:197
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 20:
		//line sql.y:203
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 21:
		//line sql.y:207
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 22:
		//line sql.y:219
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 23:
		//line sql.y:223
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 24:
		//line sql.y:235
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 25:
		//line sql.y:241
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 26:
		//line sql.y:247
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-1].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 27:
		//line sql.y:251
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 28:
		//line sql.y:257
		{
			yyVAL.statement = &Begin{}
		}
	case 29:
		//line sql.y:263
		{
			yyVAL.statement = &Commit{}
		}
	case 30:
		//line sql.y:269
		{
			yyVAL.statement = &Rollback{}
		}
	case 31:
		//line sql.y:275
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 32:
		//line sql.y:279
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 33:
		//line sql.y:284
		{
			yyVAL.statement = &DDL{Action: AST_CREATE, NewName: yyS[yypt-1].bytes}
		}
	case 34:
		//line sql.y:290
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 35:
		//line sql.y:294
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 36:
		//line sql.y:299
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 37:
		//line sql.y:305
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 38:
		//line sql.y:311
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-0].bytes}
		}
	case 39:
		//line sql.y:315
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 40:
		//line sql.y:320
		{
			yyVAL.statement = &DDL{Action: AST_DROP, Table: yyS[yypt-1].bytes}
		}
	case 41:
		//line sql.y:326
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 42:
		//line sql.y:332
		{
			yyVAL.statement = &Other{}
		}
	case 43:
		//line sql.y:336
		{
			yyVAL.statement = &Other{}
		}
	case 44:
		//line sql.y:340
		{
			yyVAL.statement = &Other{}
		}
	case 45:
		//line sql.y:345
		{
			SetAllowComments(yylex, true)
		}
	case 46:
		//line sql.y:349
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 47:
		//line sql.y:355
		{
			yyVAL.bytes2 = nil
		}
	case 48:
		//line sql.y:359
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 49:
		//line sql.y:365
		{
			yyVAL.str = AST_UNION
		}
	case 50:
		//line sql.y:369
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 51:
		//line sql.y:373
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 52:
		//line sql.y:377
		{
			yyVAL.str = AST_EXCEPT
		}
	case 53:
		//line sql.y:381
		{
			yyVAL.str = AST_INTERSECT
		}
	case 54:
		//line sql.y:386
		{
			yyVAL.str = ""
		}
	case 55:
		//line sql.y:390
		{
			yyVAL.str = AST_DISTINCT
		}
	case 56:
		//line sql.y:396
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 57:
		//line sql.y:400
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 58:
		//line sql.y:406
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 59:
		//line sql.y:410
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 60:
		//line sql.y:414
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 61:
		//line sql.y:420
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 62:
		//line sql.y:424
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 63:
		//line sql.y:429
		{
			yyVAL.bytes = nil
		}
	case 64:
		//line sql.y:433
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 65:
		//line sql.y:437
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 66:
		//line sql.y:443
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 67:
		//line sql.y:447
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 68:
		//line sql.y:453
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 69:
		//line sql.y:457
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 70:
		//line sql.y:461
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 71:
		//line sql.y:465
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 72:
		//line sql.y:470
		{
			yyVAL.bytes = nil
		}
	case 73:
		//line sql.y:474
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 74:
		//line sql.y:478
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 75:
		//line sql.y:484
		{
			yyVAL.str = AST_JOIN
		}
	case 76:
		//line sql.y:488
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 77:
		//line sql.y:492
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 78:
		//line sql.y:496
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 79:
		//line sql.y:500
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 80:
		//line sql.y:504
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 81:
		//line sql.y:508
		{
			yyVAL.str = AST_JOIN
		}
	case 82:
		//line sql.y:512
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 83:
		//line sql.y:516
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 84:
		//line sql.y:522
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 85:
		//line sql.y:526
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 86:
		//line sql.y:530
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 87:
		//line sql.y:536
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 88:
		//line sql.y:540
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 89:
		//line sql.y:545
		{
			yyVAL.indexHints = nil
		}
	case 90:
		//line sql.y:549
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 91:
		//line sql.y:553
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 92:
		//line sql.y:557
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 93:
		//line sql.y:563
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 94:
		//line sql.y:567
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 95:
		//line sql.y:572
		{
			yyVAL.boolExpr = nil
		}
	case 96:
		//line sql.y:576
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 97:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 98:
		//line sql.y:583
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 99:
		//line sql.y:587
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 100:
		//line sql.y:591
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 101:
		//line sql.y:595
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 102:
		//line sql.y:601
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 103:
		//line sql.y:605
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 104:
		//line sql.y:609
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 105:
		//line sql.y:613
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 106:
		//line sql.y:617
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 107:
		//line sql.y:621
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 108:
		//line sql.y:625
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 109:
		//line sql.y:629
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 110:
		//line sql.y:633
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 111:
		//line sql.y:637
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 112:
		//line sql.y:641
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyS[yypt-3].valExpr, End: yyS[yypt-1].valExpr}
		}
	case 113:
		//line sql.y:647
		{
			yyVAL.str = AST_EQ
		}
	case 114:
		//line sql.y:651
		{
			yyVAL.str = AST_LT
		}
	case 115:
		//line sql.y:655
		{
			yyVAL.str = AST_GT
		}
	case 116:
		//line sql.y:659
		{
			yyVAL.str = AST_LE
		}
	case 117:
		//line sql.y:663
		{
			yyVAL.str = AST_GE
		}
	case 118:
		//line sql.y:667
		{
			yyVAL.str = AST_NE
		}
	case 119:
		//line sql.y:671
		{
			yyVAL.str = AST_NSE
		}
	case 120:
		//line sql.y:677
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 121:
		//line sql.y:681
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 122:
		//line sql.y:685
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 123:
		//line sql.y:691
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 124:
		//line sql.y:697
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 125:
		//line sql.y:701
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 126:
		//line sql.y:707
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 127:
		//line sql.y:711
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 128:
		//line sql.y:715
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 129:
		//line sql.y:719
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 130:
		//line sql.y:723
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 131:
		//line sql.y:727
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 132:
		//line sql.y:731
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 133:
		//line sql.y:735
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 134:
		//line sql.y:739
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 135:
		//line sql.y:743
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 136:
		//line sql.y:747
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 137:
		//line sql.y:751
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 138:
		//line sql.y:766
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 139:
		//line sql.y:770
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 140:
		//line sql.y:774
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 141:
		//line sql.y:778
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 142:
		//line sql.y:782
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 143:
		//line sql.y:788
		{
			yyVAL.bytes = IF_BYTES
		}
	case 144:
		//line sql.y:792
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 145:
		//line sql.y:798
		{
			yyVAL.byt = AST_UPLUS
		}
	case 146:
		//line sql.y:802
		{
			yyVAL.byt = AST_UMINUS
		}
	case 147:
		//line sql.y:806
		{
			yyVAL.byt = AST_TILDA
		}
	case 148:
		//line sql.y:812
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 149:
		//line sql.y:817
		{
			yyVAL.valExpr = nil
		}
	case 150:
		//line sql.y:821
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 151:
		//line sql.y:827
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 152:
		//line sql.y:831
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 153:
		//line sql.y:837
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 154:
		//line sql.y:842
		{
			yyVAL.valExpr = nil
		}
	case 155:
		//line sql.y:846
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 156:
		//line sql.y:852
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 157:
		//line sql.y:856
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 158:
		//line sql.y:862
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 159:
		//line sql.y:866
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 160:
		//line sql.y:870
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 161:
		//line sql.y:874
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 162:
		//line sql.y:879
		{
			yyVAL.valExprs = nil
		}
	case 163:
		//line sql.y:883
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 164:
		//line sql.y:888
		{
			yyVAL.boolExpr = nil
		}
	case 165:
		//line sql.y:892
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 166:
		//line sql.y:897
		{
			yyVAL.orderBy = nil
		}
	case 167:
		//line sql.y:901
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 168:
		//line sql.y:907
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 169:
		//line sql.y:911
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 170:
		//line sql.y:917
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 171:
		//line sql.y:922
		{
			yyVAL.str = AST_ASC
		}
	case 172:
		//line sql.y:926
		{
			yyVAL.str = AST_ASC
		}
	case 173:
		//line sql.y:930
		{
			yyVAL.str = AST_DESC
		}
	case 174:
		//line sql.y:935
		{
			yyVAL.limit = nil
		}
	case 175:
		//line sql.y:939
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 176:
		//line sql.y:943
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 177:
		//line sql.y:948
		{
			yyVAL.str = ""
		}
	case 178:
		//line sql.y:952
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 179:
		//line sql.y:956
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 180:
		//line sql.y:969
		{
			yyVAL.columns = nil
		}
	case 181:
		//line sql.y:973
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 182:
		//line sql.y:979
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 183:
		//line sql.y:983
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 184:
		//line sql.y:988
		{
			yyVAL.updateExprs = nil
		}
	case 185:
		//line sql.y:992
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 186:
		//line sql.y:998
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 187:
		//line sql.y:1002
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 188:
		//line sql.y:1008
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 189:
		//line sql.y:1012
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 190:
		//line sql.y:1018
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 191:
		//line sql.y:1022
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 192:
		//line sql.y:1028
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 193:
		//line sql.y:1032
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 194:
		//line sql.y:1038
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 195:
		//line sql.y:1043
		{
			yyVAL.empty = struct{}{}
		}
	case 196:
		//line sql.y:1045
		{
			yyVAL.empty = struct{}{}
		}
	case 197:
		//line sql.y:1048
		{
			yyVAL.empty = struct{}{}
		}
	case 198:
		//line sql.y:1050
		{
			yyVAL.empty = struct{}{}
		}
	case 199:
		//line sql.y:1053
		{
			yyVAL.empty = struct{}{}
		}
	case 200:
		//line sql.y:1055
		{
			yyVAL.empty = struct{}{}
		}
	case 201:
		//line sql.y:1059
		{
			yyVAL.empty = struct{}{}
		}
	case 202:
		//line sql.y:1061
		{
			yyVAL.empty = struct{}{}
		}
	case 203:
		//line sql.y:1063
		{
			yyVAL.empty = struct{}{}
		}
	case 204:
		//line sql.y:1065
		{
			yyVAL.empty = struct{}{}
		}
	case 205:
		//line sql.y:1067
		{
			yyVAL.empty = struct{}{}
		}
	case 206:
		//line sql.y:1070
		{
			yyVAL.empty = struct{}{}
		}
	case 207:
		//line sql.y:1072
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		//line sql.y:1075
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		//line sql.y:1077
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		//line sql.y:1080
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		//line sql.y:1082
		{
			yyVAL.empty = struct{}{}
		}
	case 212:
		//line sql.y:1086
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 213:
		//line sql.y:1091
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
