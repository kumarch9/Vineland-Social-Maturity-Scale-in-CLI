package expitems

var ItemExp = map[int]string{
	1:  "Vocalizes inarticulately (other than crying or fretting). Spontaneously gargles or coos. Laughs spontaneously or when simulated.",
	6:  "Attempts to obtain nearby but beyond reach.",
	7:  "Plays with rattle or simple objects for quarter hour or longer without need of attention.",
	14: "Indicates desire to be “Talked” to or beyond mere handling, or care for physical needs.",
	16: "Has established control ofsaliva so that mouth or chin does not require wiping except while eating.",
	17: "Comes when called, points to particular objects in pictures when asked; in general cooperates on verbal request in very simple activities.",
	22: "Pours from one vessel to another without messing; removes, transfers, or replaces objects in somewhat purposeful manner.",
	23: "Opens closed doors; climbs up on chair; uses stool for reaching; removes simple impediments.",
	26: "Walks by pushing a cart on wheels or a walker.",
	27: "Activity is individual rather than cooperative, but he “gets along “ with other children.",
	28: "Eats things like biscuit or bread holding in his own hand or uses spoon to eat from a bow, a cup, or a plate.",
	35: "By actions or speech expresses to go to urinate or care himself; may be assisted at the same.",
	36: "Occupies self at play such as drawing or colouring with pencil, looking at books or pictures.",
	41: "Comes in out of rain. Shows some caution regarding strangers. Is careful as regards falling on stairs.",
	44: "Gives simple accounts of experience or tells stories.",
	46: "Participates in coordinated group activity as kindergarten circle games, cooking or group play.",
	49: "Entertains others by reciting, singing, or dancing.",
	55: "Draws forms like man, house, tree, animal etc.",
	56: "Engages in tag, hide and seek, jumping, rope, tops, skipping, or marbles.",
	57: "Hoops-ring pushed by hand or stick, cycle tyre.",
	59: "Games with othersrequiring taking turns, observing rules without undue dissension; caroms, draft, snake, and ladder, trade etc.",
	60: "Is responsible with small sums ofmoney when sent to make payments of explicit purchases.",
	63: "Writes (not prints) legibly with a pencil a dozen or more simple words with correct spelling.",
	65: "Performs bedtime operation without help.; goes to room alone , changes dress and turns out-light.",
	67: "After the meals is served first, helps himselfmore according to the need.",
	69: "Boys: Games not requiring definite skill and with only less rules such as unorganised hockey, football, khokho and follow the leader. Takes hikes or bicycle rides." +
		"Girls: Engages in dramatic play symbolizing domestic or social situation such as playing house, school, doctor-nurses etc.,(Note: *Sex differentiation in play is noted at this stage and there is a shift in girls play to more sedentary ones." +
		"However, credit item regardless ofsex ifthis differentiation has not yet been established.)",
	71: "Makes practical use of hammer, screwdriver and household articles. Sews. Uses garden tools etc.",
	72: "Helps effectively at simple tasksfor which some continuous responsibility is assumed; dusting; arranging; cleaning; washing dishes, making bed etc.",
	73: "Reads comic strips, movie titles,simple stories, notessimple instructions, elementary newsitem for own entertainment or information.",
	76: "Buys useful articles, exercises some choice or discretion in doing so and is responsible for safety of articles, money and correct change.",
	79: "Able to decide for self, which shop to go for purchasing different articles.",
	80: "Makes articles for self use, e.g. making simple gardens, stitching buttons, preparing tea forself, doing small repairs, talking care of own cabinet, table and room or performs occasional work on own initiative such as odd jobs, housework, helping in care of children, sewing, selling magazines, carrying newspapers for which some money is paid.",
	81: "Writes letters to get information regarding some books, magazine or toys.",
	82: "Makes useful articles; cooks, bakes, raises pets, writes simple stories or poems; produce simple drawings or painting.",
	83: "Is sometimes left along and is successful in looking after own immediate needs or hose of others who may be left in his care.",
	85: "Participates in skilled games and sports as card games, basketball, tennis, hockey, badminton, understands rules and methods ofscoring.",
	86: "Includes washing and drying hair, care of nails, proper selection of clothing according to occasion and weather.",
	87: "Selects and purchases minor articles of personal clothing with regard for appropriateness, such as ribbons, underwear, linen, shoes etc.",
	88: "Is an active member of a cooperative group, athletic team, club, social or literary organisation. Plans or participates in picnic trips, outdoor sports, etc.",
	89: "Such as assisting in house work, caring for garden, cleaning car, washing window, waiting at table, bringing water etc.",
}

var SocialAreas = map[int]string{
	1: "SHG- Self Help General",
	2: "SHE- Self Help Eating",
	3: "SHD- Self Help Dressing",
	4: "SD- Self-Direction",
	5: "OCC- Occupation",
	6: "COM- Communication",
	7: "LOC- Locomotion",
	8: "SOC- Socialisation",
}

//score|SA
var RawScale = map[float32]float32{
	1.0: 0.06,
	1.5: 0.09,
	2.0: 0.12,
	2.5: 0.15,
	3.0: 0.18,
	3.5: 0.21,
	4.0: 0.24,
	5:   0.30,
	6:   0.35,
	7:   0.41,
	8:   0.47,
	9:   0.53,
	10:  0.59,
	11:  0.65,
	12:  0.71,
	13:  0.77,
	14:  0.83,
	15:  0.89,
	16:  0.94,
	17:  1.00,
	18:  1.06,
	19:  1.12,
	20:  1.18,
	21:  1.24,
	22:  1.30,
	23:  1.35,
	24:  1.41,
	25:  1.47,
	26:  1.53,
	28:  1.59,
	29:  1.71,
	30:  1.77,
	31:  1.83,
	32:  1.89,
	33:  1.94,
	34:  2.00,
	35:  2.10,
	36:  2.20,
	37:  2.30,
	38:  2.40,
	39:  2.50,
	40:  2.60,
	41:  2.70,
	42:  2.80,
	43:  2.90,
	44:  3.00,
	45:  3.20,
	46:  3.30,
	47:  3.50,
	48:  3.70,
	49:  3.80,
	50:  4.00,
	51:  4.20,
	52:  4.30,
	53:  4.50,
	54:  4.70,
	55:  4.80,
	56:  5.00,
	57:  5.20,
	58:  5.40,
	59:  5.60,
	60:  5.80,
	61:  6.00,
	62:  6.30,
	63:  6.30,
	64:  6.80,
	65:  7.00,
	66:  7.20,
	67:  7.40,
	68:  7.60,
	69:  7.80,
	70:  8.00,
	71:  8.30,
	72:  8.50,
	73:  8.80,
	74:  9.00,
	75:  9.30,
	76:  9.70,
	77:  10.00,
	78:  10.30,
	79:  10.50,
	80:  10.80,
	81:  11.00,
	82:  11.30,
	83:  11.70,
	84:  12.00,
	85:  12.60,
	86:  13.20,
	87:  13.80,
	88:  14.40,
	89:  15.00,
}

//var MatualityMonth = map[int]float32{}
type ScaleNorm struct {
	Months                                [][]float32
	SHG, SHE, SHD, SD, OCC, COM, LOC, SOC [][]int
}

var NormsAnalysis = ScaleNorm{

	Months: [][]float32{{1.5, 3.0, 4.5, 6.0, 7.5, 9.0, 10.5, 12.0}, {14.4, 10.8, 19.2, 21.6, 24.0}, {28, 32, 36}, {40, 44, 48}, {52, 56, 60}, {64, 68, 72}, {76, 80, 84}, {88, 92, 96}, {100, 104, 108}, {112, 116, 120}, {124, 128, 132}, {136, 140, 144}, {156, 168, 180}},
	SHG:    [][]int{{2, 3, 5, 6, 8, 9, 13, 15}, {23, 26}, {35, 41}, {0}, {51}, {0}, {0}, {66}, {0}, {0}, {0}, {0}, {0}},
	SHE:    [][]int{{11, 16}, {20, 25, 28, 30, 33}, {38, 39}, {0}, {0}, {0}, {62}, {67}, {0}, {75}, {0}, {0}, {0}},
	SHD:    [][]int{{0}, {21}, {37, 40, 42}, {47}, {50, 52, 54}, {0}, {64}, {65}, {70, 70}, {74}, {0}, {0}, {86}},
	SD:     [][]int{{0}, {0}, {0}, {0}, {0}, {60}, {0}, {0}, {0}, {76}, {79}, {83}, {87}},
	OCC:    [][]int{{7}, {19, 22, 24}, {36, 43}, {48}, {55}, {57}, {0}, {0}, {71, 72}, {0}, {80}, {82}, {89}},
	COM:    [][]int{{1, 10}, {17, 31, 34}, {0}, {44}, {0}, {58}, {63}, {0}, {73}, {0}, {81}, {0}, {84}},
	LOC:    [][]int{{12}, {18, 29, 32}, {0}, {45}, {53}, {0}, {61}, {0}, {0}, {77}, {0}, {0}, {0}},
	SOC:    [][]int{{4, 14}, {27}, {0}, {46, 49}, {0}, {56, 59}, {0}, {68, 69}, {0}, {0}, {0}, {0}, {85, 88}},
}