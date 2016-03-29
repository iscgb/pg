package types

type ValueAppender interface {
	AppendValue(b []byte, quote int) ([]byte, error)
}

//------------------------------------------------------------------------------

// Q is a ValueAppender that represents safe SQL query.
type Q []byte

var _ ValueAppender = Q(nil)

func (q Q) AppendValue(dst []byte, quote int) ([]byte, error) {
	return append(dst, q...), nil
}

//------------------------------------------------------------------------------

// F is a ValueAppender that represents SQL field, e.g. table or column name.
type F []byte

var _ ValueAppender = F(nil)

func (f F) AppendValue(dst []byte, quote int) ([]byte, error) {
	return AppendFieldBytes(dst, f, quote), nil
}
