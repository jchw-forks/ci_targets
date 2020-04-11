// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type NestedSameName2 struct {
	Version uint32
	MainData *NestedSameName2_Main
	Dummy *NestedSameName2_DummyObj
	_io *kaitai.Stream
	_root *NestedSameName2
	_parent interface{}
}
func NewNestedSameName2() *NestedSameName2 {
	return &NestedSameName2{
	}
}

func (this *NestedSameName2) Read(io *kaitai.Stream, parent interface{}, root *NestedSameName2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Version = uint32(tmp1)
	tmp2 := NewNestedSameName2_Main()
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.MainData = tmp2
	tmp3 := NewNestedSameName2_DummyObj()
	err = tmp3.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Dummy = tmp3
	return err
}
type NestedSameName2_Main struct {
	MainSize int32
	Foo *NestedSameName2_Main_FooObj
	_io *kaitai.Stream
	_root *NestedSameName2
	_parent *NestedSameName2
}
func NewNestedSameName2_Main() *NestedSameName2_Main {
	return &NestedSameName2_Main{
	}
}

func (this *NestedSameName2_Main) Read(io *kaitai.Stream, parent *NestedSameName2, root *NestedSameName2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.MainSize = int32(tmp4)
	tmp5 := NewNestedSameName2_Main_FooObj()
	err = tmp5.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Foo = tmp5
	return err
}
type NestedSameName2_Main_FooObj struct {
	Data1 []byte
	_io *kaitai.Stream
	_root *NestedSameName2
	_parent *NestedSameName2_Main
}
func NewNestedSameName2_Main_FooObj() *NestedSameName2_Main_FooObj {
	return &NestedSameName2_Main_FooObj{
	}
}

func (this *NestedSameName2_Main_FooObj) Read(io *kaitai.Stream, parent *NestedSameName2_Main, root *NestedSameName2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp6, err := this._io.ReadBytes(int((this._parent.MainSize * 2)))
	if err != nil {
		return err
	}
	tmp6 = tmp6
	this.Data1 = tmp6
	return err
}
type NestedSameName2_DummyObj struct {
	DummySize int32
	Foo *NestedSameName2_DummyObj_FooObj
	_io *kaitai.Stream
	_root *NestedSameName2
	_parent *NestedSameName2
}
func NewNestedSameName2_DummyObj() *NestedSameName2_DummyObj {
	return &NestedSameName2_DummyObj{
	}
}

func (this *NestedSameName2_DummyObj) Read(io *kaitai.Stream, parent *NestedSameName2, root *NestedSameName2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp7, err := this._io.ReadS4le()
	if err != nil {
		return err
	}
	this.DummySize = int32(tmp7)
	tmp8 := NewNestedSameName2_DummyObj_FooObj()
	err = tmp8.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Foo = tmp8
	return err
}
type NestedSameName2_DummyObj_FooObj struct {
	Data2 []byte
	_io *kaitai.Stream
	_root *NestedSameName2
	_parent *NestedSameName2_DummyObj
}
func NewNestedSameName2_DummyObj_FooObj() *NestedSameName2_DummyObj_FooObj {
	return &NestedSameName2_DummyObj_FooObj{
	}
}

func (this *NestedSameName2_DummyObj_FooObj) Read(io *kaitai.Stream, parent *NestedSameName2_DummyObj, root *NestedSameName2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp9, err := this._io.ReadBytes(int((this._parent.DummySize * 2)))
	if err != nil {
		return err
	}
	tmp9 = tmp9
	this.Data2 = tmp9
	return err
}
