// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type NavParentFalse2 struct {
	Parentless *NavParentFalse2_Child
	_io *kaitai.Stream
	_root *NavParentFalse2
	_parent interface{}
}
func NewNavParentFalse2() *NavParentFalse2 {
	return &NavParentFalse2{
	}
}

func (this *NavParentFalse2) Read(io *kaitai.Stream, parent interface{}, root *NavParentFalse2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := NewNavParentFalse2_Child()
	err = tmp1.Read(this._io, nil, this._root)
	if err != nil {
		return err
	}
	this.Parentless = tmp1
	return err
}
type NavParentFalse2_Child struct {
	Foo uint8
	_io *kaitai.Stream
	_root *NavParentFalse2
	_parent interface{}
}
func NewNavParentFalse2_Child() *NavParentFalse2_Child {
	return &NavParentFalse2_Child{
	}
}

func (this *NavParentFalse2_Child) Read(io *kaitai.Stream, parent interface{}, root *NavParentFalse2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp2, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Foo = tmp2
	return err
}
