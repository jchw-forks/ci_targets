// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type SwitchManualIntElse struct {
	Opcodes []*SwitchManualIntElse_Opcode
	_io *kaitai.Stream
	_root *SwitchManualIntElse
	_parent interface{}
}
func NewSwitchManualIntElse() *SwitchManualIntElse {
	return &SwitchManualIntElse{
	}
}

func (this *SwitchManualIntElse) Read(io *kaitai.Stream, parent interface{}, root *SwitchManualIntElse) (err error) {
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
		tmp2 := NewSwitchManualIntElse_Opcode()
		err = tmp2.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Opcodes = append(this.Opcodes, tmp2)
	}
	return err
}
type SwitchManualIntElse_Opcode struct {
	Code uint8
	Body interface{}
	_io *kaitai.Stream
	_root *SwitchManualIntElse
	_parent *SwitchManualIntElse
}
func NewSwitchManualIntElse_Opcode() *SwitchManualIntElse_Opcode {
	return &SwitchManualIntElse_Opcode{
	}
}

func (this *SwitchManualIntElse_Opcode) Read(io *kaitai.Stream, parent *SwitchManualIntElse, root *SwitchManualIntElse) (err error) {
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
		tmp4 := NewSwitchManualIntElse_Opcode_Intval()
		err = tmp4.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp4
	case 83:
		tmp5 := NewSwitchManualIntElse_Opcode_Strval()
		err = tmp5.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp5
	default:
		tmp6 := NewSwitchManualIntElse_Opcode_Noneval()
		err = tmp6.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp6
	}
	return err
}
type SwitchManualIntElse_Opcode_Intval struct {
	Value uint8
	_io *kaitai.Stream
	_root *SwitchManualIntElse
	_parent *SwitchManualIntElse_Opcode
}
func NewSwitchManualIntElse_Opcode_Intval() *SwitchManualIntElse_Opcode_Intval {
	return &SwitchManualIntElse_Opcode_Intval{
	}
}

func (this *SwitchManualIntElse_Opcode_Intval) Read(io *kaitai.Stream, parent *SwitchManualIntElse_Opcode, root *SwitchManualIntElse) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp7, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Value = tmp7
	return err
}
type SwitchManualIntElse_Opcode_Strval struct {
	Value string
	_io *kaitai.Stream
	_root *SwitchManualIntElse
	_parent *SwitchManualIntElse_Opcode
}
func NewSwitchManualIntElse_Opcode_Strval() *SwitchManualIntElse_Opcode_Strval {
	return &SwitchManualIntElse_Opcode_Strval{
	}
}

func (this *SwitchManualIntElse_Opcode_Strval) Read(io *kaitai.Stream, parent *SwitchManualIntElse_Opcode, root *SwitchManualIntElse) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp8, err := this._io.ReadBytesTerm(0, false, true, true)
	if err != nil {
		return err
	}
	this.Value = string(tmp8)
	return err
}
type SwitchManualIntElse_Opcode_Noneval struct {
	Filler uint32
	_io *kaitai.Stream
	_root *SwitchManualIntElse
	_parent *SwitchManualIntElse_Opcode
}
func NewSwitchManualIntElse_Opcode_Noneval() *SwitchManualIntElse_Opcode_Noneval {
	return &SwitchManualIntElse_Opcode_Noneval{
	}
}

func (this *SwitchManualIntElse_Opcode_Noneval) Read(io *kaitai.Stream, parent *SwitchManualIntElse_Opcode, root *SwitchManualIntElse) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp9, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Filler = uint32(tmp9)
	return err
}
