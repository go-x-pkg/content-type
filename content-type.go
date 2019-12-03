package contenttype

type Kind uint8

const (
	JSON Kind = iota
	YAML
)

var codecText = map[Kind]string{
	JSON: "json",
	YAML: "yaml",
}

var contentTypeText = map[Kind]string{
	JSON: "application/json",
	YAML: "application/x-yaml",
}

func (k Kind) String() string { return contentTypeText[k] }

func (k Kind) Codec() string {
	return codecText[k]
}

func (k Kind) ContentType() string {
	return contentTypeText[k]
}
