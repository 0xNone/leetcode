package leetcode

//func isValidSudoku(board [][]byte) bool {
//	for row := 0; row < 3; row++ {
//		for col := 0; col < 3; col++ {
//			tMap := map[byte]struct{}{}
//			for irow := row * 3; irow < (row+1)*3; irow++ {
//				for icol := col * 3; icol < (col+1)*3; icol++ {
//					if board[irow][icol] == byte('.') {
//						continue
//					}
//					if _, ok := tMap[board[irow][icol]]; ok {
//						return false
//					} else {
//						tMap[board[irow][icol]] = struct{}{}
//					}
//				}
//			}
//		}
//	}
//	for row := 0; row < 9; row++ {
//		cMap := map[byte]struct{}{}
//		for col := 0; col < 9; col++ {
//			if board[row][col] == byte('.') {
//				continue
//			}
//			if _, ok := cMap[board[row][col]]; ok {
//				return false
//			} else {
//				cMap[board[row][col]] = struct{}{}
//			}
//		}
//	}
//	for col := 0; col < 9; col++ {
//		rMap := map[byte]struct{}{}
//		for row := 0; row < 9; row++ {
//			if board[row][col] == byte('.') {
//				continue
//			}
//			if _, ok := rMap[board[row][col]]; ok {
//				return false
//			} else {
//				rMap[board[row][col]] = struct{}{}
//			}
//		}
//	}
//	return true
//}

func isValidSudoku(board [][]byte) bool {
	rowBoard := [9][9]bool{}
	colBoard := [9][9]bool{}
	mapBoard := [9][9]bool{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			num := board[row][col] - '1'
			if rowBoard[col][num] || colBoard[row][num] || mapBoard[row/3*3+col/3][num] {
				return false
			} else {
				rowBoard[col][num] = true
				colBoard[row][num] = true
				mapBoard[row/3*3+col/3][num] = true
			}
		}
	}
	return true
}
