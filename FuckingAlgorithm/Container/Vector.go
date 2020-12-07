package main

type IVector interface {
	Push_back(v interface{})
	Pop_back()
	Front() interface{}
	Back() interface{}
	Empty() bool
	Size() int
	Insert(v interface{}, index int)
	Erase(index int)
}

// vector
type Vector struct {
	S []interface{} `json:"s"`
}

func (this *Vector) Push_back(v interface{}) {

	this.S = append(this.S, v)
}

func (this *Vector) Pop_back() {

	this.S = this.S[:len(this.S)-1]

}

// e.g. temp := v.Front
func (this *Vector) Front() interface{} {
	if len(this.S) <= 0 {
		return nil
	}
	return this.S[0]
}

func (this *Vector) Back() interface{} {
	if len(this.S) <= 0 {
		return nil
	}
	return this.S[len(this.S)-1]
}

func (this *Vector) Empty() bool {
	return len(this.S) <= 0
}

func (this *Vector) Size() int {
	return len(this.S)
}

func (this *Vector) Insert(value interface{}, nIndex int) {
	temp := this.S[nIndex : len(this.S)-1]
	this.S = this.S[:nIndex]
	this.S = append(this.S, value)
	this.S = append(temp)
}

func (this *Vector) Erase(nIndex int) {
	this.S = append(this.S[:nIndex], this.S[nIndex+1:])
}
