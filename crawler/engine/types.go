package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

type Request struct {
	Url        string
	ParserFunc func([]byte, string) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	PayLoad interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
