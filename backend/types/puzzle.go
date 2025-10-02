package types

type PuzzleCategory string

const (
	Puzzle3x3      PuzzleCategory = "3x3"
	Puzzle2x2      PuzzleCategory = "2x2"
	Puzzle4x4      PuzzleCategory = "4x4"
	Puzzle5x5      PuzzleCategory = "5x5"
	Puzzle6x6      PuzzleCategory = "6x6"
	Puzzle7x7      PuzzleCategory = "7x7"
	Puzzle3x3oh    PuzzleCategory = "3x3oh"
	Puzzle3x3bl    PuzzleCategory = "3x3bl"
	Puzzle3x3mbl   PuzzleCategory = "3x3mbl"
	Puzzle4x4bl    PuzzleCategory = "4x4bl"
	Puzzle5x5bl    PuzzleCategory = "5x5bl"
	PuzzleMegaminx PuzzleCategory = "Megaminx"
	PuzzlePyraminx PuzzleCategory = "Pyraminx"
	PuzzleSkewb    PuzzleCategory = "Skewb"
	PuzzleSquare_1 PuzzleCategory = "Square-1"
	PuzzleClock    PuzzleCategory = "Clock"
	PuzzleFMC      PuzzleCategory = "FMC"
)

type Penalty string

const (
	PlusTwo Penalty = "+2"
	DNF     Penalty = "DNF"
)
