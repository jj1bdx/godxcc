package godxcc

// WAE prefix to DXCC prefix

var WAEToDXCC = map[string]string{
	// European Turkey = Turkey
	"*TA1": "TA",
	// Vienna Intl Ctr = Austria
	"*4U1V": "OE",
	// Shetland Islands = Scotland
	"*GM/s": "GM",
	// African Italy = Italy
	"*IG9": "I",
	// Sicily = Italy
	"*IT9": "I",
	// Bear Island = Svalbard
	"*JW/b": "JW",
}

// DXCC prefix to DXCC entity code

var ctyToEntitycode = map[string]int{
	"1A":    246,
	"1S":    247,
	"3A":    260,
	"3B6":   4,
	"3B8":   165,
	"3B9":   207,
	"3C":    49,
	"3C0":   195,
	"3D2":   176,
	"3D2/c": 489,
	"3D2/r": 460,
	"3DA":   468,
	"3V":    474,
	"3W":    293,
	"3X":    107,
	"3Y/b":  24,
	"3Y/p":  199,
	"4J":    18,
	"4L":    75,
	"4O":    514,
	"4S":    315,
	"4U1I":  117,
	"4U1U":  289,
	"4W":    511,
	"4X":    336,
	"5A":    436,
	"5B":    215,
	"5H":    470,
	"5N":    450,
	"5R":    438,
	"5T":    444,
	"5U":    187,
	"5V":    483,
	"5W":    190,
	"5X":    286,
	"5Z":    430,
	"6W":    456,
	"6Y":    82,
	"7O":    492,
	"7P":    432,
	"7Q":    440,
	"7X":    400,
	"8P":    62,
	"8Q":    159,
	"8R":    129,
	"9A":    497,
	"9G":    424,
	"9H":    257,
	"9J":    482,
	"9K":    348,
	"9L":    458,
	"9M2":   299,
	"9M6":   46,
	"9N":    369,
	"9Q":    414,
	"9U":    404,
	"9V":    381,
	"9X":    454,
	"9Y":    90,
	"A2":    402,
	"A3":    160,
	"A4":    370,
	"A5":    306,
	"A6":    391,
	"A7":    376,
	"A9":    304,
	"AP":    372,
	"BS7":   506,
	"BV":    386,
	"BV9P":  505,
	"BY":    318,
	"C2":    157,
	"C3":    203,
	"C5":    422,
	"C6":    60,
	"C9":    181,
	"CE":    112,
	"CE0X":  217,
	"CE0Y":  47,
	"CE0Z":  125,
	"CE9":   13,
	"CM":    70,
	"CN":    446,
	"CP":    104,
	"CT":    272,
	"CT3":   256,
	"CU":    149,
	"CX":    144,
	"CY0":   211,
	"CY9":   252,
	"D2":    401,
	"D4":    409,
	"D6":    411,
	"DL":    230,
	"DU":    375,
	"E3":    51,
	"E4":    510,
	"E5/n":  191,
	"E5/s":  234,
	"E6":    188,
	"E7":    501,
	"EA":    281,
	"EA6":   21,
	"EA8":   29,
	"EA9":   32,
	"EI":    245,
	"EK":    14,
	"EL":    434,
	"EP":    330,
	"ER":    179,
	"ES":    52,
	"ET":    53,
	"EU":    27,
	"EX":    135,
	"EY":    262,
	"EZ":    280,
	"F":     227,
	"FG":    79,
	"FH":    169,
	"FJ":    516,
	"FK":    162,
	"FK/c":  512,
	"FM":    84,
	"FO":    175,
	"FO/a":  508,
	"FO/c":  36,
	"FO/m":  509,
	"FP":    277,
	"FR":    453,
	"FS":    213,
	"FT/g":  99,
	"FT/j":  124,
	"FT/t":  276,
	"FT/w":  41,
	"FT/x":  131,
	"FT/z":  10,
	"FW":    298,
	"FY":    63,
	"G":     223,
	"GD":    114,
	"GI":    265,
	"GJ":    122,
	"GM":    279,
	"GU":    106,
	"GW":    294,
	"H4":    185,
	"H40":   507,
	"HA":    239,
	"HB":    287,
	"HB0":   251,
	"HC":    120,
	"HC8":   71,
	"HH":    78,
	"HI":    72,
	"HK":    116,
	"HK0/a": 216,
	"HK0/m": 161,
	"HL":    137,
	"HP":    88,
	"HR":    80,
	"HS":    387,
	"HV":    295,
	"HZ":    378,
	"I":     248,
	"IS":    225,
	"J2":    382,
	"J3":    77,
	"J5":    109,
	"J6":    97,
	"J7":    95,
	"J8":    98,
	"JA":    339,
	"JD/m":  177,
	"JD/o":  192,
	"JT":    363,
	"JW":    259,
	"JX":    118,
	"JY":    342,
	"K":     291,
	"KG4":   105,
	"KH0":   166,
	"KH1":   20,
	"KH2":   103,
	"KH3":   123,
	"KH4":   174,
	"KH5":   197,
	"KH6":   110,
	"KH7K":  138,
	"KH8":   9,
	"KH8/s": 515,
	"KH9":   297,
	"KL":    6,
	"KP1":   182,
	"KP2":   285,
	"KP4":   202,
	"KP5":   43,
	"LA":    266,
	"LU":    100,
	"LX":    254,
	"LY":    146,
	"LZ":    212,
	"OA":    136,
	"OD":    354,
	"OE":    206,
	"OH":    224,
	"OH0":   5,
	"OJ0":   167,
	"OK":    503,
	"OM":    504,
	"ON":    209,
	"OX":    237,
	"OY":    222,
	"OZ":    221,
	"P2":    163,
	"P4":    91,
	"P5":    344,
	"PA":    263,
	"PJ2":   517,
	"PJ4":   520,
	"PJ5":   519,
	"PJ7":   518,
	"PY":    108,
	"PY0F":  56,
	"PY0S":  253,
	"PY0T":  273,
	"PZ":    140,
	"R1FJ":  61,
	"S0":    302,
	"S2":    305,
	"S5":    499,
	"S7":    379,
	"S9":    219,
	"SM":    284,
	"SP":    269,
	"ST":    466,
	"SU":    478,
	"SV":    236,
	"SV/a":  180,
	"SV5":   45,
	"SV9":   40,
	"T2":    282,
	"T30":   301,
	"T31":   31,
	"T32":   48,
	"T33":   490,
	"T5":    232,
	"T7":    278,
	"T8":    22,
	"TA":    390,
	"TF":    242,
	"TG":    76,
	"TI":    308,
	"TI9":   37,
	"TJ":    406,
	"TK":    214,
	"TL":    408,
	"TN":    412,
	"TR":    420,
	"TT":    410,
	"TU":    428,
	"TY":    416,
	"TZ":    442,
	"UA":    54,
	"UA2":   126,
	"UA9":   15,
	"UK":    292,
	"UN":    130,
	"UR":    288,
	"V2":    94,
	"V3":    66,
	"V4":    249,
	"V5":    464,
	"V6":    173,
	"V7":    168,
	"V8":    345,
	"VE":    1,
	"VK":    150,
	"VK0H":  111,
	"VK0M":  153,
	"VK9C":  38,
	"VK9L":  147,
	"VK9M":  171,
	"VK9N":  189,
	"VK9W":  303,
	"VK9X":  35,
	"VP2E":  12,
	"VP2M":  96,
	"VP2V":  65,
	"VP5":   89,
	"VP6":   172,
	"VP6/d": 513,
	"VP8":   141,
	"VP8/g": 235,
	"VP8/h": 241,
	"VP8/o": 238,
	"VP8/s": 240,
	"VP9":   64,
	"VQ9":   33,
	"VR":    321,
	"VU":    324,
	"VU4":   11,
	"VU7":   142,
	"XE":    50,
	"XF4":   204,
	"XT":    480,
	"XU":    312,
	"XW":    143,
	"XX9":   152,
	"XZ":    309,
	"YA":    3,
	"YB":    327,
	"YI":    333,
	"YJ":    158,
	"YK":    384,
	"YL":    145,
	"YN":    86,
	"YO":    275,
	"YS":    74,
	"YU":    296,
	"YV":    148,
	"YV0":   17,
	"Z2":    452,
	"Z3":    502,
	"Z6":    522,
	"Z8":    521,
	"ZA":    7,
	"ZB":    233,
	"ZC4":   283,
	"ZD7":   250,
	"ZD8":   205,
	"ZD9":   274,
	"ZF":    69,
	"ZK3":   270,
	"ZL":    170,
	"ZL7":   34,
	"ZL8":   133,
	"ZL9":   16,
	"ZP":    132,
	"ZS":    462,
	"ZS8":   201,
}

// DXCC entity code to DXCC prefix

var entitycodeToCty = map[int]string{
	246: "1A",
	247: "1S",
	260: "3A",
	4:   "3B6",
	165: "3B8",
	207: "3B9",
	49:  "3C",
	195: "3C0",
	176: "3D2",
	489: "3D2/c",
	460: "3D2/r",
	468: "3DA",
	474: "3V",
	293: "3W",
	107: "3X",
	24:  "3Y/b",
	199: "3Y/p",
	18:  "4J",
	75:  "4L",
	514: "4O",
	315: "4S",
	117: "4U1I",
	289: "4U1U",
	511: "4W",
	336: "4X",
	436: "5A",
	215: "5B",
	470: "5H",
	450: "5N",
	438: "5R",
	444: "5T",
	187: "5U",
	483: "5V",
	190: "5W",
	286: "5X",
	430: "5Z",
	456: "6W",
	82:  "6Y",
	492: "7O",
	432: "7P",
	440: "7Q",
	400: "7X",
	62:  "8P",
	159: "8Q",
	129: "8R",
	497: "9A",
	424: "9G",
	257: "9H",
	482: "9J",
	348: "9K",
	458: "9L",
	299: "9M2",
	46:  "9M6",
	369: "9N",
	414: "9Q",
	404: "9U",
	381: "9V",
	454: "9X",
	90:  "9Y",
	402: "A2",
	160: "A3",
	370: "A4",
	306: "A5",
	391: "A6",
	376: "A7",
	304: "A9",
	372: "AP",
	506: "BS7",
	386: "BV",
	505: "BV9P",
	318: "BY",
	157: "C2",
	203: "C3",
	422: "C5",
	60:  "C6",
	181: "C9",
	112: "CE",
	217: "CE0X",
	47:  "CE0Y",
	125: "CE0Z",
	13:  "CE9",
	70:  "CM",
	446: "CN",
	104: "CP",
	272: "CT",
	256: "CT3",
	149: "CU",
	144: "CX",
	211: "CY0",
	252: "CY9",
	401: "D2",
	409: "D4",
	411: "D6",
	230: "DL",
	375: "DU",
	51:  "E3",
	510: "E4",
	191: "E5/n",
	234: "E5/s",
	188: "E6",
	501: "E7",
	281: "EA",
	21:  "EA6",
	29:  "EA8",
	32:  "EA9",
	245: "EI",
	14:  "EK",
	434: "EL",
	330: "EP",
	179: "ER",
	52:  "ES",
	53:  "ET",
	27:  "EU",
	135: "EX",
	262: "EY",
	280: "EZ",
	227: "F",
	79:  "FG",
	169: "FH",
	516: "FJ",
	162: "FK",
	512: "FK/c",
	84:  "FM",
	175: "FO",
	508: "FO/a",
	36:  "FO/c",
	509: "FO/m",
	277: "FP",
	453: "FR",
	213: "FS",
	99:  "FT/g",
	124: "FT/j",
	276: "FT/t",
	41:  "FT/w",
	131: "FT/x",
	10:  "FT/z",
	298: "FW",
	63:  "FY",
	223: "G",
	114: "GD",
	265: "GI",
	122: "GJ",
	279: "GM",
	106: "GU",
	294: "GW",
	185: "H4",
	507: "H40",
	239: "HA",
	287: "HB",
	251: "HB0",
	120: "HC",
	71:  "HC8",
	78:  "HH",
	72:  "HI",
	116: "HK",
	216: "HK0/a",
	161: "HK0/m",
	137: "HL",
	88:  "HP",
	80:  "HR",
	387: "HS",
	295: "HV",
	378: "HZ",
	248: "I",
	225: "IS",
	382: "J2",
	77:  "J3",
	109: "J5",
	97:  "J6",
	95:  "J7",
	98:  "J8",
	339: "JA",
	177: "JD/m",
	192: "JD/o",
	363: "JT",
	259: "JW",
	118: "JX",
	342: "JY",
	291: "K",
	105: "KG4",
	166: "KH0",
	20:  "KH1",
	103: "KH2",
	123: "KH3",
	174: "KH4",
	197: "KH5",
	110: "KH6",
	138: "KH7K",
	9:   "KH8",
	515: "KH8/s",
	297: "KH9",
	6:   "KL",
	182: "KP1",
	285: "KP2",
	202: "KP4",
	43:  "KP5",
	266: "LA",
	100: "LU",
	254: "LX",
	146: "LY",
	212: "LZ",
	136: "OA",
	354: "OD",
	206: "OE",
	224: "OH",
	5:   "OH0",
	167: "OJ0",
	503: "OK",
	504: "OM",
	209: "ON",
	237: "OX",
	222: "OY",
	221: "OZ",
	163: "P2",
	91:  "P4",
	344: "P5",
	263: "PA",
	517: "PJ2",
	520: "PJ4",
	519: "PJ5",
	518: "PJ7",
	108: "PY",
	56:  "PY0F",
	253: "PY0S",
	273: "PY0T",
	140: "PZ",
	61:  "R1FJ",
	302: "S0",
	305: "S2",
	499: "S5",
	379: "S7",
	219: "S9",
	284: "SM",
	269: "SP",
	466: "ST",
	478: "SU",
	236: "SV",
	180: "SV/a",
	45:  "SV5",
	40:  "SV9",
	282: "T2",
	301: "T30",
	31:  "T31",
	48:  "T32",
	490: "T33",
	232: "T5",
	278: "T7",
	22:  "T8",
	390: "TA",
	242: "TF",
	76:  "TG",
	308: "TI",
	37:  "TI9",
	406: "TJ",
	214: "TK",
	408: "TL",
	412: "TN",
	420: "TR",
	410: "TT",
	428: "TU",
	416: "TY",
	442: "TZ",
	54:  "UA",
	126: "UA2",
	15:  "UA9",
	292: "UK",
	130: "UN",
	288: "UR",
	94:  "V2",
	66:  "V3",
	249: "V4",
	464: "V5",
	173: "V6",
	168: "V7",
	345: "V8",
	1:   "VE",
	150: "VK",
	111: "VK0H",
	153: "VK0M",
	38:  "VK9C",
	147: "VK9L",
	171: "VK9M",
	189: "VK9N",
	303: "VK9W",
	35:  "VK9X",
	12:  "VP2E",
	96:  "VP2M",
	65:  "VP2V",
	89:  "VP5",
	172: "VP6",
	513: "VP6/d",
	141: "VP8",
	235: "VP8/g",
	241: "VP8/h",
	238: "VP8/o",
	240: "VP8/s",
	64:  "VP9",
	33:  "VQ9",
	321: "VR",
	324: "VU",
	11:  "VU4",
	142: "VU7",
	50:  "XE",
	204: "XF4",
	480: "XT",
	312: "XU",
	143: "XW",
	152: "XX9",
	309: "XZ",
	3:   "YA",
	327: "YB",
	333: "YI",
	158: "YJ",
	384: "YK",
	145: "YL",
	86:  "YN",
	275: "YO",
	74:  "YS",
	296: "YU",
	148: "YV",
	17:  "YV0",
	452: "Z2",
	502: "Z3",
	522: "Z6",
	521: "Z8",
	7:   "ZA",
	233: "ZB",
	283: "ZC4",
	250: "ZD7",
	205: "ZD8",
	274: "ZD9",
	69:  "ZF",
	270: "ZK3",
	170: "ZL",
	34:  "ZL7",
	133: "ZL8",
	16:  "ZL9",
	132: "ZP",
	462: "ZS",
	201: "ZS8",
}