// Package contenttype provides simple content (media) type.
package contenttype

// Kind provides ContentType enum
type Kind uint8

const (
	// JSON / JavaScript Object Notation
	JSON Kind = iota

	// YAML / Yet Another Markup Language
	YAML

	// TSV / Tab Separated Values
	TSV

	// CSV / Comma-Separated Values
	CSV

	// XML / eXtensible Markup Language
	XML

	// HTML / HyperText Markup Language
	HTML

	// Protobuf / Protocol Buffers
	// see: https://tools.ietf.org/html/draft-rfernando-protocol-buffers-00
	Protobuf

	// MsgPack / Message Pack
	MsgPack

	// M3U8 / M3U8 files are the basis for the HTTP Live Streaming (HLS)
	/* see:
	https://en.wikipedia.org/wiki/M3U#Internet_media_types
	https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/StreamingMediaGuide/DeployingHTTPLiveStreaming/DeployingHTTPLiveStreaming.html
	*/
	M3U8
)

var codecText = map[Kind]string{
	JSON:     "json",
	YAML:     "yaml",
	TSV:      "tsv",
	CSV:      "csv",
	XML:      "xml",
	HTML:     "html",
	Protobuf: "protobuf",
	MsgPack:  "msgpack",
	M3U8:     "m3u8",
}

var contentTypeText = map[Kind]string{
	JSON:     "application/json",
	YAML:     "application/x-yaml",
	XML:      "application/xml",
	Protobuf: "application/protobuf",
	MsgPack:  "application/x-msgpack",
	M3U8:     "application/vnd.apple.mpegurl",
	TSV:      "text/tab-separated-values",
	CSV:      "text/csv",
	HTML:     "text/html",
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
