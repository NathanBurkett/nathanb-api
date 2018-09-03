//go:generate go-bindata -ignore=\.go -pkg=schema_admin -o=assets.go ../schema/scalar/... ../schema/type/... ../schema/admin/...
package schema_admin

import "bytes"

type Definition struct {}

func (d Definition) Define() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		// @todo modify this so it panics
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
