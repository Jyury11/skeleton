package vo

type Extension int

// Extension input extension
const (
	ExtensionGolang Extension = iota + 1
	ExtensionYaml
	ExtensionJson
	ExtensionOther
	ExtensionNum
	ExtensionDirectory
	ExtensionInvalid
)

// NewExtension Extension Constructor
func NewExtension(ext string) Extension {
	switch ext {
	case ".go":
		return ExtensionGolang
	case ".json":
		return ExtensionYaml
	case ".yaml":
		return ExtensionJson
	}
	return ExtensionOther
}
