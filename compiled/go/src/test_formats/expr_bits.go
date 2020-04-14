// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"io"
)


type ExprBits_Items int
const (
	ExprBits_Items__Foo ExprBits_Items = 1
	ExprBits_Items__Bar ExprBits_Items = 2
)
type ExprBits struct {
	EnumSeq ExprBits_Items
	A uint64
	ByteSize []byte
	RepeatExpr []int8
	SwitchOnType int8
	SwitchOnEndian *ExprBits_EndianSwitch
	_io *kaitai.Stream
	_root *ExprBits
	_parent interface{}
	_f_enumInst bool
	enumInst ExprBits_Items
	_f_instPos bool
	instPos int8
}
func NewExprBits() *ExprBits {
	return &ExprBits{
	}
}

func (this *ExprBits) Read(io *kaitai.Stream, parent interface{}, root *ExprBits) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBitsInt(2)
	if err != nil {
		return err
	}
	this.EnumSeq = ExprBits_Items(tmp1)
	tmp2, err := this._io.ReadBitsInt(3)
	if err != nil {
		return err
	}
	this.A = tmp2
	this._io.AlignToByte()
	tmp3, err := this._io.ReadBytes(int(this.A))
	if err != nil {
		return err
	}
	tmp3 = tmp3
	this.ByteSize = tmp3
	this.RepeatExpr = make([]int8, this.A)
	for i := range this.RepeatExpr {
		tmp4, err := this._io.ReadS1()
		if err != nil {
			return err
		}
		this.RepeatExpr[i] = tmp4
	}
	switch (this.A) {
	case 2:
		tmp5, err := this._io.ReadS1()
		if err != nil {
			return err
		}
		this.SwitchOnType = tmp5
	}
	tmp6 := NewExprBits_EndianSwitch()
	err = tmp6.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.SwitchOnEndian = tmp6
	return err
}
func (this *ExprBits) EnumInst() (v ExprBits_Items, err error) {
	if (this._f_enumInst) {
		return this.enumInst, nil
	}
	this.enumInst = ExprBits_Items(ExprBits_Items(this.A))
	this._f_enumInst = true
	return this.enumInst, nil
}
func (this *ExprBits) InstPos() (v int8, err error) {
	if (this._f_instPos) {
		return this.instPos, nil
	}
	_pos, err := this._io.Pos()
	if err != nil {
		return 0, err
	}
	_, err = this._io.Seek(int64(this.A), io.SeekStart)
	if err != nil {
		return 0, err
	}
	tmp7, err := this._io.ReadS1()
	if err != nil {
		return 0, err
	}
	this.instPos = tmp7
	_, err = this._io.Seek(_pos, io.SeekStart)
	if err != nil {
		return 0, err
	}
	this._f_instPos = true
	this._f_instPos = true
	return this.instPos, nil
}
type ExprBits_EndianSwitch struct {
	Foo int16
	_io *kaitai.Stream
	_root *ExprBits
	_parent *ExprBits
	_is_le int
}
func NewExprBits_EndianSwitch() *ExprBits_EndianSwitch {
	return &ExprBits_EndianSwitch{
	}
}

func (this *ExprBits_EndianSwitch) Read(io *kaitai.Stream, parent *ExprBits, root *ExprBits) (err error) {
	this._io = io
	this._parent = parent
	this._root = root
	this._is_le = -1

	switch (this._parent.A) {
	case 1:
		this._is_le = int(1)
	case 2:
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

func (this *ExprBits_EndianSwitch) _read_le() (err error) {
	tmp8, err := this._io.ReadS2le()
	if err != nil {
		return err
	}
	this.Foo = int16(tmp8)
	return err
}

func (this *ExprBits_EndianSwitch) _read_be() (err error) {
	tmp9, err := this._io.ReadS2be()
	if err != nil {
		return err
	}
	this.Foo = int16(tmp9)
	return err
}
