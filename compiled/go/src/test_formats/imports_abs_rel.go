// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ImportsAbsRel struct {
	One uint8
	Two *ImportedAndRel
	_io *kaitai.Stream
	_root *ImportsAbsRel
	_parent interface{}
}

func (this *ImportsAbsRel) Read(io *kaitai.Stream, parent interface{}, root *ImportsAbsRel) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.One = tmp1
	tmp2 := new(ImportedAndRel)
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Two = tmp2
	return err
}
