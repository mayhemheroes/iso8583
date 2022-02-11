package field

type Field interface {
	// Spec returns the field spec
	Spec() *Spec

	// SetSpec sets the field spec
	SetSpec(spec *Spec)

	// Pack serializes field value into binary representation according
	// to the field spec
	Pack() ([]byte, error)

	// Unpack deserialises the field by reading length prefix and reading
	// corresponding number of bytes from the provided data parameter and
	// then decoding it according to the field spec
	Unpack(data []byte) (int, error)

	// SetBytes sets the field Value using its binary representation
	// provided in the data parameter
	SetBytes(data []byte) error

	// Bytes returns binary representation of the field Value
	Bytes() ([]byte, error)

	// SetData links provided data parameter with the field Value
	SetData(data interface{}) error

	// UnmarshalValue sets field Value into provided v. If v is nil or not
	// a pointer it returns error.
	UnmarshalValue(v interface{}) error

	// MarshalValue sets field Value into provided v. If v is nil or not
	// a pointer it returns error.
	MarshalValue(v interface{}) error

	// String returns a string representation of the field Value
	String() (string, error)
}
