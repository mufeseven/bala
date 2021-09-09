package serializable

var (
	instance Serializable
)

type Serializable interface {
	Serialization(value interface{}) ([]byte, error)
	Deserialization(bytes []byte, ref interface{}) error
}

func GetInstance() Serializable {
	return instance
}

func SetInstance(s Serializable) {
	instance = s
}
