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

type TermBytes struct {
	S1 byte[]
	S2 byte[]
	S3 byte[]
	_io *kaitai.Stream
	_root *TermBytes
	_parent interface{}
}

func (this *TermBytes) Read(io *kaitai.Stream, parent interface{}, root *TermBytes) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytesTerm(124, false, true, true)
	if err != nil {
		return err
	}
	this.S1 = tmp1
	tmp2, err := this._io.ReadBytesTerm(124, false, false, true)
	if err != nil {
		return err
	}
	this.S2 = tmp2
	tmp3, err := this._io.ReadBytesTerm(64, true, true, true)
	if err != nil {
		return err
	}
	this.S3 = tmp3
	return err
}
