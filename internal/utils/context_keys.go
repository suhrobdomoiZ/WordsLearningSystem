package utils

type ContextKey int

const (
	UserKey ContextKey = iota // Value must be (*dto.User) or nil
)
