// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"bytes"
)

type ProcessCoerceSwitch struct {
	BufType uint8
	Flag uint8
	BufUnproc interface{}
	BufProc interface{}
	_io *kaitai.Stream
	_root *ProcessCoerceSwitch
	_parent interface{}
	_raw_BufUnproc []byte
	_raw_BufProc []byte
	_raw__raw_BufProc []byte
	_f_buf bool
	buf interface{}
}
func NewProcessCoerceSwitch() *ProcessCoerceSwitch {
	return &ProcessCoerceSwitch{
	}
}

func (this *ProcessCoerceSwitch) Read(io *kaitai.Stream, parent interface{}, root *ProcessCoerceSwitch) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.BufType = tmp1
	tmp2, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Flag = tmp2
	if (this.Flag == 0) {
		switch (this.BufType) {
		case 0:
			tmp3, err := this._io.ReadBytes(int(4))
			if err != nil {
				return err
			}
			tmp3 = tmp3
			this._raw_BufUnproc = tmp3
			_io__raw_BufUnproc := kaitai.NewStream(bytes.NewReader(this._raw_BufUnproc))
			tmp4 := NewProcessCoerceSwitch_Foo()
			err = tmp4.Read(_io__raw_BufUnproc, this, this._root)
			if err != nil {
				return err
			}
			this.BufUnproc = tmp4
		default:
			tmp5, err := this._io.ReadBytes(int(4))
			if err != nil {
				return err
			}
			tmp5 = tmp5
			this._raw_BufUnproc = tmp5
		}
	}
	if (this.Flag != 0) {
		switch (this.BufType) {
		case 0:
			tmp6, err := this._io.ReadBytes(int(4))
			if err != nil {
				return err
			}
			tmp6 = tmp6
			this._raw__raw_BufProc = tmp6
			this._raw_BufProc = kaitai.ProcessXOR(this._raw__raw_BufProc, []byte{170})
			_io__raw_BufProc := kaitai.NewStream(bytes.NewReader(this._raw_BufProc))
			tmp7 := NewProcessCoerceSwitch_Foo()
			err = tmp7.Read(_io__raw_BufProc, this, this._root)
			if err != nil {
				return err
			}
			this.BufProc = tmp7
		default:
			tmp8, err := this._io.ReadBytes(int(4))
			if err != nil {
				return err
			}
			tmp8 = tmp8
			this._raw__raw_BufProc = tmp8
			this._raw_BufProc = kaitai.ProcessXOR(this._raw__raw_BufProc, []byte{170})
		}
	}
	return err
}
func (this *ProcessCoerceSwitch) Buf() (v interface{}, err error) {
	if (this._f_buf) {
		return this.buf, nil
	}
	var tmp9 interface{};
	if (this.Flag == 0) {
		tmp9 = this.BufUnproc
	} else {
		tmp9 = this.BufProc
	}
	this.buf = interface{}(tmp9)
	this._f_buf = true
	return this.buf, nil
}
type ProcessCoerceSwitch_Foo struct {
	Bar []byte
	_io *kaitai.Stream
	_root *ProcessCoerceSwitch
	_parent *ProcessCoerceSwitch
}
func NewProcessCoerceSwitch_Foo() *ProcessCoerceSwitch_Foo {
	return &ProcessCoerceSwitch_Foo{
	}
}

func (this *ProcessCoerceSwitch_Foo) Read(io *kaitai.Stream, parent *ProcessCoerceSwitch, root *ProcessCoerceSwitch) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp10, err := this._io.ReadBytes(int(4))
	if err != nil {
		return err
	}
	tmp10 = tmp10
	this.Bar = tmp10
	return err
}
