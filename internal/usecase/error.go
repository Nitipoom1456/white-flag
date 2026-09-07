package usecase

type UsecaseError struct {
	Err     error
	Code    string
	Message string
}

const (
	CODE_INVALID_INPUT  = "IVL"
	CODE_SERVER_ERROR   = "SVR"
	CODE_DUPLICATED_KEY = "DUP"

	MSG_INVALID_INPUT      = "invalid input"
	MSG_SERVER_ERROR       = "server error"
	MSG_APP_DUPLICATED_KEY = "app is already registered"

	MSG_APP_VERSION_DUPLICATED_KEY = "app version is already exist"
)
