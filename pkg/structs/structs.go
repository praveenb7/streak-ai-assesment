package structs

type Request struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

type Response struct {
	Solutions [][]int `json:"solutions"`
}

type Pair struct {
	A int
	B int
}

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
