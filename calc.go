package main

import (
	"bytes"
)

const (
	ADD = 0
	SUB = 1
	MUL = 2
	DIV = 3
)

func CalcAllOperation(qs []int, a int) string {
	operations := make([][]int, 4)
	operations = generateOperations(operations, len(qs)-1)

	for _, operation := range operations {
		res, ok := calcOperation(qs, operation)
		if ok && res == a {
			ops := bytes.NewBuffer(make([]byte, 0, 100))
			for _, typ := range operation {
				switch typ {
				case ADD:
					ops.WriteString("+")
				case SUB:
					ops.WriteString("-")
				case MUL:
					ops.WriteString("*")
				case DIV:
					ops.WriteString("/")
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
	return generateOperations(newOps, remainLen-1)
}

func calcOperation(qs []int, op []int) (int, bool) {
	if len(op) <= 0 {
		return qs[0], true
	}

	if len(op) == 1 || op[1] == ADD || op[1] == SUB {
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
	case ADD:
		return n + m, true
	case SUB:
		return n - m, true
	case MUL:
		return n * m, true
	case DIV:
		if n%m == 0 {
			return n / m, true
		}
		return 0, false
	}
	return 0, false
}
