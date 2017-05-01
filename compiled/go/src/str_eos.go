// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
import "io"
import "golang.org/x/text/encoding/charmap"
import "golang.org/x/text/encoding/traditionalchinese"
import "golang.org/x/text/encoding/japanese"

var _ = io.SeekStart
var _ = charmap.CodePage437
var _ = japanese.ShiftJIS
var _ = traditionalchinese.Big5

type StrEos struct {
	Str string
	_io *kaitai.Stream
	_root *StrEos
	_parent interface{}
}

func (this *StrEos) Read(io *kaitai.Stream, parent interface{}, root *StrEos) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytesFull()
	if err != nil {
		return err
	}
	this.Str = string(tmp1)
	return err
}
