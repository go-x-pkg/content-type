// Package contenttype provides simple content (media) type.
package contenttype

// Kind provides ContentType enum
type Kind uint8

const (
	// JSON / JavaScript Object Notation
	JSON Kind = iota

	// YAML / Yet Another Markup Language
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

// String impl Stringer for Kind
func (k Kind) String() string { return contentTypeText[k] }

// Codec provides short string representation.
func (k Kind) Codec() string {
	return codecText[k]
}

// ContentType returnes MIME type as string.
func (k Kind) ContentType() string {
	return contentTypeText[k]
}
