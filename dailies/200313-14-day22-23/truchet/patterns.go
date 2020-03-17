package truchet

// PatternA creates truchet pattern A
var PatternA = NewPattern(
	Row{0},
)

// PatternC creates truchet pattern C
var PatternC = NewPattern(
	Row{0, 2},
	Row{2, 0},
)

// PatternD creates truchet pattern D
var PatternD = NewPattern(
	Row{1, 0},
	Row{2, 3},
)

// PatternE creates truchet pattern E
var PatternE = NewPattern(
	Row{2, 1, 0, 3},
	Row{3, 0, 1, 2},
	Row{0, 3, 2, 1},
	Row{1, 2, 3, 0},
)

// PatternX is something
var PatternX = NewPattern(
	Row{1, 3, 2, 0},
	Row{3, 1, 0, 2},
	Row{0, 2, 3, 1},
	Row{2, 0, 1, 3},
)

// PatternDouat72 creates truchet pattern Douat 72
var PatternDouat72 = NewPattern(
	Row{1, 3, 3, 1, 3, 3, 0, 0, 2, 0, 0, 2},
	Row{3, 1, 3, 3, 1, 3, 0, 2, 0, 0, 2, 0},
	Row{3, 3, 1, 3, 3, 1, 2, 0, 0, 2, 0, 0},
	Row{1, 3, 3, 1, 3, 3, 0, 0, 2, 0, 0, 2},
	Row{3, 1, 3, 3, 1, 3, 0, 2, 0, 0, 2, 0},
	Row{3, 3, 1, 3, 3, 1, 2, 0, 0, 2, 0, 0},

	Row{2, 2, 0, 2, 2, 0, 3, 1, 1, 3, 1, 1},
	Row{2, 0, 2, 2, 0, 2, 1, 3, 1, 1, 3, 1},
	Row{0, 2, 2, 0, 2, 2, 1, 1, 3, 1, 1, 3},
	Row{2, 2, 0, 2, 2, 0, 3, 1, 1, 3, 1, 1},
	Row{2, 0, 2, 2, 0, 2, 1, 3, 1, 1, 3, 1},
	Row{0, 2, 2, 0, 2, 2, 1, 1, 3, 1, 1, 3},
)
