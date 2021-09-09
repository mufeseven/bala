//go:generate stringer -type ErrorCode -linecomment

package errcode

type ErrorCode int

const (
	None ErrorCode = iota
)
