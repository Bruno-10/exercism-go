package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    count := 0
    for i, v := range cb {
        if i == file {
            for _, v2 := range v {
                if v2 {
                    count++
                }
            }
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    count := 0

    for _, v := range cb {
        if rank - 1 > len(v) || rank -1 < 0 {
            return count
        }
        if v[rank -1] {
            count ++
        }
    }
    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    count := 0
    for _, v := range cb {
        for range v {
            count++
        }
    }
    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    count := 0
    for _, v := range cb {
        for _, v2 := range v {
            if v2 {
                count++
            }
        }
    }
    return count
}
