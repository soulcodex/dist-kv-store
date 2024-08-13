package store

const joiningNodeErrorMessage = "Error joining a node"

type JoiningNodeError struct {
	context map[string]interface{}
}

func NewJoiningNodeErrorWithCtx(context map[string]interface{}) *JoiningNodeError {
	return &JoiningNodeError{context: context}
}

func (jne *JoiningNodeError) Error() string {
	return joiningNodeErrorMessage
}
