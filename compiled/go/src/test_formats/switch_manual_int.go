// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type SwitchManualInt struct {
	Opcodes []*SwitchManualInt_Opcode
	_io *kaitai.Stream
	_root *SwitchManualInt
	_parent interface{}
}
func NewSwitchManualInt() *SwitchManualInt {
	return &SwitchManualInt{
	}
}

func (this *SwitchManualInt) Read(io *kaitai.Stream, parent interface{}, root *SwitchManualInt) (err error) {
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
		tmp2 := NewSwitchManualInt_Opcode()
		err = tmp2.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Opcodes = append(this.Opcodes, tmp2)
	}
	return err
}
type SwitchManualInt_Opcode struct {
	Code uint8
	Body interface{}
	_io *kaitai.Stream
	_root *SwitchManualInt
	_parent *SwitchManualInt
}
func NewSwitchManualInt_Opcode() *SwitchManualInt_Opcode {
	return &SwitchManualInt_Opcode{
	}
}

func (this *SwitchManualInt_Opcode) Read(io *kaitai.Stream, parent *SwitchManualInt, root *SwitchManualInt) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Code = tmp3
	switch (this.Code) {
	case 73:
		tmp4 := NewSwitchManualInt_Opcode_Intval()
		err = tmp4.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp4
	case 83:
		tmp5 := NewSwitchManualInt_Opcode_Strval()
		err = tmp5.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp5
	}
	return err
}
type SwitchManualInt_Opcode_Intval struct {
	Value uint8
	_io *kaitai.Stream
	_root *SwitchManualInt
	_parent *SwitchManualInt_Opcode
}
func NewSwitchManualInt_Opcode_Intval() *SwitchManualInt_Opcode_Intval {
	return &SwitchManualInt_Opcode_Intval{
	}
}

func (this *SwitchManualInt_Opcode_Intval) Read(io *kaitai.Stream, parent *SwitchManualInt_Opcode, root *SwitchManualInt) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp6, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Value = tmp6
	return err
}
type SwitchManualInt_Opcode_Strval struct {
	Value string
	_io *kaitai.Stream
	_root *SwitchManualInt
	_parent *SwitchManualInt_Opcode
}
func NewSwitchManualInt_Opcode_Strval() *SwitchManualInt_Opcode_Strval {
	return &SwitchManualInt_Opcode_Strval{
	}
}

func (this *SwitchManualInt_Opcode_Strval) Read(io *kaitai.Stream, parent *SwitchManualInt_Opcode, root *SwitchManualInt) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp7, err := this._io.ReadBytesTerm(0, false, true, true)
	if err != nil {
		return err
	}
	this.Value = string(tmp7)
	return err
}
