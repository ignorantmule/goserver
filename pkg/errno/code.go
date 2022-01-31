package errno

var (
	// OK represents a successful request.
	OK = &Errno{Code: 0, Message: "OK"}

	// InternalServerError represents all unknown server-side errors.
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}

	// ErrBind represents a failed parameter binding.
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// ErrPageNotFound represents a route not matched error.
	ErrPageNotFound = &Errno{Code: 10003, Message: "Page not found."}

	// ErrValidation represents all validation failed errors.
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}

	// ErrDatabase represents a database error.
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}

	// ErrToken represents a error when signing JWT token.
	ErrToken = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// ErrEncrypt represents a encrypting error.
	ErrEncrypt = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}

	// ErrUserNotFound represents the user not found.
	ErrUserNotFound = &Errno{Code: 20102, Message: "User was not found."}

	// ErrUserAlreadyExist represents the user already exist.
	ErrUserAlreadyExist = &Errno{Code: 20103, Message: "User already exist."}

	// ErrTokenInvalid represents the token format is wrong.
	ErrTokenInvalid = &Errno{Code: 20104, Message: "Token was invalid."}

	// ErrPasswordIncorrect represents the password is incorrect.
	ErrPasswordIncorrect = &Errno{Code: 20105, Message: "Password was incorrect."}

	// ErrPostNotFound represents  the post not found.
	ErrPostNotFound = &Errno{Code: 20201, Message: "Post was not found."}

	// ErrPostAlreadyExist represents the post already exist.
	ErrPostAlreadyExist = &Errno{Code: 20202, Message: "Post already exist."}
)
