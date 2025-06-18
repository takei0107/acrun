package lang

type FileType int

const (
	Invalid FileType = iota
	C
)

type fileTypeConfig struct {
	jsonString string
}

var ftConfigMap = map[FileType]*fileTypeConfig{
	C: {
		jsonString: CJson,
	},
}

func ResolveFileType(lang string) FileType {
	var ft FileType
	switch lang {
	case "c":
		ft = C
	default:
		ft = Invalid
	}
	return ft
}
