package contentType

type Kind uint8

const (
	Json Kind = iota
	Yaml
)

var codecText = map[Kind]string{
	Json: "json",
	Yaml: "yaml",
}

var contentTypeText = map[Kind]string{
	Json: "application/json",
	Yaml: "application/x-yaml",
}

func (it Kind) String() string { return codecText[it] }

func (ct Kind) Codec() string {
	return codecText[ct]
}

func (ct Kind) ContentType() string {
	return contentTypeText[ct]
}
