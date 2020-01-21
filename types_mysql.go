package dalc

type Varchar []uint8

func (v Varchar) String() string {
	return string(v)
}

func NewVarchar(s string) Varchar {
	v := []uint8(s)
	return v
}