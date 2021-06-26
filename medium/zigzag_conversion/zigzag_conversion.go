package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var (
		str       string
		goingDown bool
		row       int
	)

	length := numRows

	if len(s) < numRows {
		length = len(s)
	}

	arr := make([]string, length)

	for _, c := range s {
		arr[row] += string(c)

		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			row += 1
		} else {
			row -= 1
		}
	}

	for _, v := range arr {
		str += v
	}

	return str
}
