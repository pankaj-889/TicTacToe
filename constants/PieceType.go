package constants

type PieceType int

const (
	X = iota
	O
)

var stp = map[string]PieceType{
	"X": X,
	"O": O,
}

var pts = map[PieceType]string{
	X: "X",
	O: "O",
}

func StringToPiece(piece string) PieceType {
	return stp[piece]
}

func (p PieceType) String() string {
	return pts[p]
}
