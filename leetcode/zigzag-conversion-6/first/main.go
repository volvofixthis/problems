package first

func convert(s string, numRows int) string {
	rows := make([][]byte, numRows)
	cRow := 0
	b := 0
	if numRows > 1 {
		b = 1
	}
	d := 1
	for i := 0; i < len(s); i++ {
		rows[cRow] = append(rows[cRow], s[i])
		if cRow == numRows-1 {
			d = b * -1
		} else if cRow == 0 {
			d = b * 1
		}
		cRow += d
	}
	result := ""
	for _, row := range rows {
		result = result + string(row)
	}
	return result
}
