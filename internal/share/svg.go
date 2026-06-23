package share

import "encoding/base64"

const SVGDataURLPrefix = "data:image/svg+xml;base64,"

func SVGDataURLFromBytes(data []byte) string {
	return SVGDataURLPrefix + base64.StdEncoding.EncodeToString(data)
}
