package main

import "fmt"

const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
	D = iota // 3
	E = iota // 4
	F = iota // 5
	G = iota // 6
	H = iota // 7
	I = iota // 8
	J = iota // 9
	K = iota // 10
)

const (
	L = iota // 0 => iota will be resetted in another const declaration
	M        // 1 => we can omit iota for the next declaration
	N        // 2
)

const (
	O = iota * 10
	P = iota * 10
	Q = iota * 10
)

const (
	_ = iota // use blank identifier
	R
	S
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
	fmt.Println(J)
	fmt.Println(K)

	fmt.Println("===")

	fmt.Println(L)
	fmt.Println(M)
	fmt.Println(N)

	fmt.Println("===")

	fmt.Println(O)
	fmt.Println(P)
	fmt.Println(Q)

	fmt.Println("===")

	fmt.Println(R)
	fmt.Println(S)
}
