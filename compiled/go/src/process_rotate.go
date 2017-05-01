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

type ProcessRotate struct {
	Buf1 byte[]
	Buf2 byte[]
	Key uint8
	Buf3 byte[]
	_io *kaitai.Stream
	_root *ProcessRotate
	_parent interface{}
}

func (this *ProcessRotate) Read(io *kaitai.Stream, parent interface{}, root *ProcessRotate) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytes(int(5))
	if err != nil {
		return err
	}
	this.Buf1 = tmp1
	tmp2, err := this._io.ReadBytes(int(5))
	if err != nil {
		return err
	}
	this.Buf2 = tmp2
	tmp3, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Key = tmp3
	tmp4, err := this._io.ReadBytes(int(5))
	if err != nil {
		return err
	}
	this.Buf3 = tmp4
	return err
}
