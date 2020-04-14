// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"io"
)


type EnumToIClassBorder1_Animal int
const (
	EnumToIClassBorder1_Animal__Dog EnumToIClassBorder1_Animal = 4
	EnumToIClassBorder1_Animal__Cat EnumToIClassBorder1_Animal = 7
	EnumToIClassBorder1_Animal__Chicken EnumToIClassBorder1_Animal = 12
)
type EnumToIClassBorder1 struct {
	Pet1 EnumToIClassBorder1_Animal
	Pet2 EnumToIClassBorder1_Animal
	_io *kaitai.Stream
	_root *EnumToIClassBorder1
	_parent interface{}
	_f_someDog bool
	someDog EnumToIClassBorder1_Animal
	_f_checker bool
	checker *EnumToIClassBorder2
}
func NewEnumToIClassBorder1() *EnumToIClassBorder1 {
	return &EnumToIClassBorder1{
	}
}

func (this *EnumToIClassBorder1) Read(io *kaitai.Stream, parent interface{}, root *EnumToIClassBorder1) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Pet1 = EnumToIClassBorder1_Animal(tmp1)
	tmp2, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Pet2 = EnumToIClassBorder1_Animal(tmp2)
	return err
}
func (this *EnumToIClassBorder1) SomeDog() (v EnumToIClassBorder1_Animal, err error) {
	if (this._f_someDog) {
		return this.someDog, nil
	}
	this.someDog = EnumToIClassBorder1_Animal(EnumToIClassBorder1_Animal(4))
	this._f_someDog = true
	return this.someDog, nil
}
func (this *EnumToIClassBorder1) Checker() (v *EnumToIClassBorder2, err error) {
	if (this._f_checker) {
		return this.checker, nil
	}
	_pos, err := this._io.Pos()
	if err != nil {
		return nil, err
	}
	_, err = this._io.Seek(int64(0), io.SeekStart)
	if err != nil {
		return nil, err
	}
	tmp3 := NewEnumToIClassBorder2(this._root)
	err = tmp3.Read(this._io, this, nil)
	if err != nil {
		return nil, err
	}
	this.checker = tmp3
	_, err = this._io.Seek(_pos, io.SeekStart)
	if err != nil {
		return nil, err
	}
	this._f_checker = true
	this._f_checker = true
	return this.checker, nil
}
