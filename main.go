package main

import (
	"fmt"
	"strconv"
	"strings"
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

	ans := CalcAllOperation(q4, a)

	fmt.Println("Ans: " + ans)

	Post(st.ID, ans)
}
