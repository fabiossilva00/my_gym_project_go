package model

type TipoVideo int

const (
	Youtube TipoVideo = iota + 1
	Unknown
)

func (t TipoVideo) String() string {
	return [...]string{"Youtube"}[t-1]
}

func (t TipoVideo) EnumIndex() int {
	return int(t)
}

func GetTipo_VideoByString(name string) TipoVideo {
	switch name {
	case Youtube.String():
		return Youtube
	default:
		return Unknown
	}
}
