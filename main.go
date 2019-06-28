package main

import (
	"strings"
	"strconv"
	"bytes"
	"fmt"
)

func main() {
	Init()

	st := Get("3")

	q2 := strings.SplitN(st.Question, " = ", 2)
	a, _ := strconv.Atoi(q2[1])
	q3 := strings.Split(q2[0], " ? ")
	q4 := make([]int, len(q3))
	for i, qStr := range q3 {
		q, _ := strconv.Atoi(qStr)
		q4[i] = q
	}

	ans := calcAllOperation(q4, a)

	fmt.Println("Ans: "+ans)

	Post(st.ID, ans)
}

func calcAllOperation(qs []int, a int) string {
	operations := make([][]int, 4)
	operations = generateOperations(operations, len(qs)-1)

	for _, operation := range operations {
		res, ok := calcOperation(qs, operation)
		if ok && res == a {
			ops := bytes.NewBuffer(make([]byte, 0, 100))
			for _, typ := range operation {
				switch typ {
				case 0: ops.WriteString("+")
				case 1: ops.WriteString("-")
				case 2: ops.WriteString("*")
				case 3: ops.WriteString("/")
				}
			}
			return ops.String()
		}
	}
	panic("What???")
}

func generateOperations(ops [][]int, remainLen int) [][]int {
	if remainLen <= 0 {
		return ops
	}

	newOps := make([][]int, len(ops)*4)
	for i, a := range ops {
		for j := 0; j < 4; j++ {
			op := make([]int, 0, len(a)+1)
			op = append(op, a...)
			op = append(op, j)
			newOps[i*4+j] = op
		}
	}
	return generateOperations(newOps, remainLen - 1)
}

func calcOperation(qs []int, op []int) (int, bool) {
	if (len(op) <= 0) {
		return qs[0], true
	}

	if (len(op) == 1 || op[1] < 2) {
		newQ, ok := calcOne(qs[0], qs[1], op[0])
		if !ok {
			return 0, false
		}
		res := make([]int, 0, len(qs)-1)
		res = append(res, newQ)
		res = append(res, qs[2:]...)
		return calcOperation(res, op[1:])
	}
	newQ, ok := calcOne(qs[1], qs[2], op[1])
	if !ok {
		return 0, false
	}
	res := make([]int, 0, len(qs)-1)
	res = append(res, qs[0])
	res = append(res, newQ)
	res = append(res, qs[3:]...)

	oN := make([]int, 0, len(op)-1)
	oN = append(oN, op[0])
	oN = append(oN, op[2:]...)
	return calcOperation(res, oN)
}

func calcOne(n, m, typ int) (int, bool) {
	switch typ {
	case 0: return n + m, true
	case 1: return n - m, true
	case 2: return n * m, true
	case 3:
		if (n % m == 0) {
			return n / m, true
		}
		return 0, false
	}
	return 0, false
}
