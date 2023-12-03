package errs

type Type string

const (
	InvalidOperationType Type = "invalid_operation"
	NotFoundType         Type = "not_found"
	InternalType         Type = "internal_type"
	InvalidParamsType    Type = "invalid_params"
	ConflictType         Type = "conflict"
)
