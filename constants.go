package snapkit

// Language type for better readability
type Language string

const (
	LanguageKey     string   = "lang"
	LanguageDefault Language = "en"
	LanguageEnglish Language = "en"
	LanguageThai    Language = "th"
)

const (
	POSTGRES_MAX_IDLE_CONNS = 25
	POSTGRES_MAX_OPEN_CONNS = 25
)

const (
	Tier0 = 0 // Completely blocked

	Tier1 = 1

	Tier2 = 12

	Tier3 = 32

	Tier4 = 64

	Tier5 = 128

	Tier6 = 256

	Tier7 = 512
)

const (
	MaxFailedAttempts = 5
)

type ErrorCode string

// General Success and Error Codes
const (
	SuccessCode             ErrorCode = "SUCCESS"
	InternalErrorCode       ErrorCode = "ERR_INTERNAL"
	NotFoundCode            ErrorCode = "ERR_NOT_FOUND"
	BadRequestCode          ErrorCode = "ERR_BAD_REQUEST"
	ForbiddenCode           ErrorCode = "ERR_FORBIDDEN"
	EndpointNotFoundCode    ErrorCode = "ERR_ENDPOINT_NOT_FOUND"
	ValidationErrorCode     ErrorCode = "ERR_VALIDATION"
	TooManyRequestsCode     ErrorCode = "ERR_TOO_MANY_REQUESTS"
	UnableToGetTrxCode      ErrorCode = "ERR_UNABLE_TO_GET_TRX"
	UnableToCommitTrxCode   ErrorCode = "ERR_UNABLE_TO_COMMIT_TRX"
	UnableToRollbackTrxCode ErrorCode = "ERR_UNABLE_TO_ROLLBACK_TRX"
	UnableToGetUserCode     ErrorCode = "ERR_UNABLE_TO_GET_USER"
	UnauthorizedCode        ErrorCode = "ERR_UNAUTHORIZED"
)
