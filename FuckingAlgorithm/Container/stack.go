package main

type Stack struct {
	S []interface{}
}

func (st *Stack) Empty() bool {
	return len(st.S) <= 0
}

func (st *Stack) Size() int {
	return len(st.S)
}

func (st *Stack) Top() interface{} {
	if st.Empty() {
		return nil
	}
	return st.S[len(st.S)-1]
}

func (st *Stack) Push(val interface{}) {
	st.S = append(st.S, val)
}

func (st *Stack) Pop() {
	nCurLen := len(st.S) - 1
	st.S = st.S[:nCurLen]
}
