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

// operationは4進数 下の桁から上の桁に

// iは0～
func getOperation(ops uint, i uint) uint {
	ops = ops >> (i * 2)
	ops = ops & 3 /* 0b11 */
	return ops
}

func operationToString(op uint, len uint) string {
	operation := bytes.NewBuffer(make([]byte, 0, len))
	for i := uint(0); i < len; i++ {
		switch getOperation(op, i) {
			case ADD:
				operation.WriteString("+")
			case SUB:
				operation.WriteString("-")
			case MUL:
				operation.WriteString("*")
			case DIV:
				operation.WriteString("/")
		}
	}
	return operation.String()
}

func CalcAllOperation(qs []int, a int) string {
	len := uint( len(qs)-1 )
	maxOps := uint( 1 << (len * 2) )

	for op := uint(0); op < maxOps; op++ {
		res, ok := calcOperation(qs, op)
		if ok && res == a {
			return operationToString(op, len)
		}
	}
	panic("What???")
}

func calcOperation(qs []int, op uint) (int, bool) {
	op1 := getOperation(op, 1)
	if len(qs) <= 2 || op1 == ADD || op1 == SUB {
		newQ, ok := calcOne(qs[0], qs[1], getOperation(op, 0))
		if !ok {
			return 0, false
		}
		if (len(qs) <= 2) {
			return newQ, true
		}
		res := make([]int, 0, len(qs)-1)
		res = append(res, newQ)
		res = append(res, qs[2:]...)
		op = op >> (1 * 2)
		return calcOperation(res, op)
	}

	newQ, ok := calcOne(qs[1], qs[2], op1)
	if !ok {
		return 0, false
	}
	res := make([]int, 0, len(qs)-1)
	res = append(res, qs[0])
	res = append(res, newQ)
	res = append(res, qs[3:]...)

	o0 := getOperation(op, 0)
	op = op >> (2 * 2) << (1 * 2)
	op += o0
	return calcOperation(res, op)
}

func calcOne(n, m int, typ uint) (int, bool) {
	switch typ {
	case ADD:
		return n + m, true
	case SUB:
		return n - m, true
	case MUL:
		return n * m, true
	case DIV:
		if n % m == 0 {
			return n / m, true
		}
		return 0, false
	}
	return 0, false
}
