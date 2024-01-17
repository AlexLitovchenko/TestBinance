package errs

const (
	ErrDefault       = 98
	ErrMarshal       = 99
	ErrUnmarshall    = 100
	ErrGetBinance    = 101
	ErrParsePrice    = 102
	ErrPairsNotFound = 103
)

var listErrors = map[int]string{
	ErrDefault:       "unknown error",
	ErrUnmarshall:    "error to unmarshall",
	ErrGetBinance:    "error to get binance response",
	ErrParsePrice:    "errot to parse price from binance",
	ErrPairsNotFound: "pairs not found",
	ErrMarshal:       "error to marshall",
}

type Error struct {
	Key int    `json:"key" example:"98"`
	Msg string `json:"message" example:"unknown error"`
}

func Err(key int, msg ...string) *Error {
	_, str := GetErr(key)
	e := &Error{Key: key, Msg: str}
	for _, m := range msg {
		e.Msg = e.Msg + ": " + m
	}
	return e
}

func GetErr(num int, str ...string) (int, string) {
	er, ok := listErrors[num]
	if !ok {
		return ErrDefault, listErrors[ErrDefault]
	}
	if len(str) > 0 {
		er = er + " : " + str[0]
	}
	return num, er
}
