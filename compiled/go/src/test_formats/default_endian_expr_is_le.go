// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"bytes"
)

type DefaultEndianExprIsLe struct {
	Docs []*DefaultEndianExprIsLe_Doc
	_io *kaitai.Stream
	_root *DefaultEndianExprIsLe
	_parent interface{}
}
func NewDefaultEndianExprIsLe() *DefaultEndianExprIsLe {
	return &DefaultEndianExprIsLe{
	}
}

func (this *DefaultEndianExprIsLe) Read(io *kaitai.Stream, parent interface{}, root *DefaultEndianExprIsLe) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	for i := 1;; i++ {
		tmp1, err := this._io.EOF()
		if err != nil {
			return err
		}
		if tmp1 {
			break
		}
		tmp2 := NewDefaultEndianExprIsLe_Doc()
		err = tmp2.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Docs = append(this.Docs, tmp2)
	}
	return err
}
type DefaultEndianExprIsLe_Doc struct {
	Indicator []byte
	Main *DefaultEndianExprIsLe_Doc_MainObj
	_io *kaitai.Stream
	_root *DefaultEndianExprIsLe
	_parent *DefaultEndianExprIsLe
}
func NewDefaultEndianExprIsLe_Doc() *DefaultEndianExprIsLe_Doc {
	return &DefaultEndianExprIsLe_Doc{
	}
}

func (this *DefaultEndianExprIsLe_Doc) Read(io *kaitai.Stream, parent *DefaultEndianExprIsLe, root *DefaultEndianExprIsLe) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadBytes(int(2))
	if err != nil {
		return err
	}
	tmp3 = tmp3
	this.Indicator = tmp3
	tmp4 := NewDefaultEndianExprIsLe_Doc_MainObj()
	err = tmp4.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Main = tmp4
	return err
}
type DefaultEndianExprIsLe_Doc_MainObj struct {
	SomeInt uint32
	SomeIntBe uint16
	SomeIntLe uint16
	_io *kaitai.Stream
	_root *DefaultEndianExprIsLe
	_parent *DefaultEndianExprIsLe_Doc
	_is_le int
}
func NewDefaultEndianExprIsLe_Doc_MainObj() *DefaultEndianExprIsLe_Doc_MainObj {
	return &DefaultEndianExprIsLe_Doc_MainObj{
	}
}

func (this *DefaultEndianExprIsLe_Doc_MainObj) Read(io *kaitai.Stream, parent *DefaultEndianExprIsLe_Doc, root *DefaultEndianExprIsLe) (err error) {
	this._io = io
	this._parent = parent
	this._root = root
	this._is_le = -1

	switch (true) {
	case bytes.Equal(this._parent.Indicator, []uint8{73, 73}):
		this._is_le = int(1)
	default:
		this._is_le = int(0)
	}

	switch this._is_le {
	case 0:
		err = this._read_be()
	case 1:
		err = this._read_le()
	default:
		err = kaitai.UndecidedEndiannessError{}
	}
	return err
}

func (this *DefaultEndianExprIsLe_Doc_MainObj) _read_le() (err error) {
	tmp5, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.SomeInt = uint32(tmp5)
	tmp6, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.SomeIntBe = uint16(tmp6)
	tmp7, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.SomeIntLe = uint16(tmp7)
	return err
}

func (this *DefaultEndianExprIsLe_Doc_MainObj) _read_be() (err error) {
	tmp8, err := this._io.ReadU4be()
	if err != nil {
		return err
	}
	this.SomeInt = uint32(tmp8)
	tmp9, err := this._io.ReadU2be()
	if err != nil {
		return err
	}
	this.SomeIntBe = uint16(tmp9)
	tmp10, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.SomeIntLe = uint16(tmp10)
	return err
}
