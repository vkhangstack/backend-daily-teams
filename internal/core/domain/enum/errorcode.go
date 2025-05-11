package enum

type ErrorCode int

const (
	BadRequest          ErrorCode = -400
	Unauthorized        ErrorCode = -401
	TokenNotFoundError  ErrorCode = -402
	InvalidEmail        ErrorCode = -102
	UserNotFound        ErrorCode = -103
	UseAlreadyExits     ErrorCode = -104
	UserWrongPassword   ErrorCode = -105
	InternalServerError ErrorCode = -500
)

var StatusMap = map[ErrorCode]string{
	BadRequest:          "bad request",
	Unauthorized:        "unauthorized",
	TokenNotFoundError:  "token not found",
	InvalidEmail:        "invalid email",
	UserNotFound:        "user not found",
	InternalServerError: "server error unknown",
	UserWrongPassword:   "wrong username or password",
}

func MsgErr(s ErrorCode) string {
	if status, exists := StatusMap[s]; exists {
		return status
	}
	return "server error unknown"
}
