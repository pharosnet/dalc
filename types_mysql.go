package dalc

type Varchar []uint8

func (v Varchar) String() string {
	return string(v)
}
