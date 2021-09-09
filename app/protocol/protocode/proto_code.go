//go:generate stringer -type ProtoCode -linecomment

package protocode

type ProtoCode int

const (
	Error ProtoCode = -1
)

const (
	Account_Create ProtoCode = iota + 1000
	Account_Auth
)
