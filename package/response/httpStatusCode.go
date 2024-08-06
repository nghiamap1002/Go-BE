package response

const (
	ErrorCodeSuccess      = 20001
	ErrorCodeParamInvalid = 20003
	ErrorInvalidToken     = 30001
)

var msg = map[int]string{
	ErrorCodeSuccess:      "success",
	ErrorInvalidToken:     "Invalid Token",
	ErrorCodeParamInvalid: "Invalid Param",
}
