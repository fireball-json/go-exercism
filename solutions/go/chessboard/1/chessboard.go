package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File = []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard = map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	arr := cb[file]
	count := 0
	for _, v := range arr {
		if v {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	for _, arr := range cb {
		if rank > 0 && rank-1 < len(arr) && arr[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, arr := range cb {
		count += len(arr)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for k := range cb {
		count += CountInFile(cb, k)
	}
	return count
}
