// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ExprSizeofValue0 struct {
	Block1 *ExprSizeofValue0_Block
	More uint16
	_io *kaitai.Stream
	_root *ExprSizeofValue0
	_parent interface{}
	_f_selfSizeof bool
	selfSizeof int
	_f_sizeofBlock bool
	sizeofBlock int
	_f_sizeofBlockB bool
	sizeofBlockB int
	_f_sizeofBlockA bool
	sizeofBlockA int
	_f_sizeofBlockC bool
	sizeofBlockC int
}
func NewExprSizeofValue0() *ExprSizeofValue0 {
	return &ExprSizeofValue0{
	}
}

func (this *ExprSizeofValue0) Read(io *kaitai.Stream, parent interface{}, root *ExprSizeofValue0) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := NewExprSizeofValue0_Block()
	err = tmp1.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Block1 = tmp1
	tmp2, err := this._io.ReadU2le()
	if err != nil {
		return err
	}
	this.More = uint16(tmp2)
	return err
}
func (this *ExprSizeofValue0) SelfSizeof() (v int, err error) {
	if (this._f_selfSizeof) {
		return this.selfSizeof, nil
	}
	this.selfSizeof = int(9)
	this._f_selfSizeof = true
	return this.selfSizeof, nil
}
func (this *ExprSizeofValue0) SizeofBlock() (v int, err error) {
	if (this._f_sizeofBlock) {
		return this.sizeofBlock, nil
	}
	this.sizeofBlock = int(7)
	this._f_sizeofBlock = true
	return this.sizeofBlock, nil
}
func (this *ExprSizeofValue0) SizeofBlockB() (v int, err error) {
	if (this._f_sizeofBlockB) {
		return this.sizeofBlockB, nil
	}
	this.sizeofBlockB = int(4)
	this._f_sizeofBlockB = true
	return this.sizeofBlockB, nil
}
func (this *ExprSizeofValue0) SizeofBlockA() (v int, err error) {
	if (this._f_sizeofBlockA) {
		return this.sizeofBlockA, nil
	}
	this.sizeofBlockA = int(1)
	this._f_sizeofBlockA = true
	return this.sizeofBlockA, nil
}
func (this *ExprSizeofValue0) SizeofBlockC() (v int, err error) {
	if (this._f_sizeofBlockC) {
		return this.sizeofBlockC, nil
	}
	this.sizeofBlockC = int(2)
	this._f_sizeofBlockC = true
	return this.sizeofBlockC, nil
}
type ExprSizeofValue0_Block struct {
	A uint8
	B uint32
	C []byte
	_io *kaitai.Stream
	_root *ExprSizeofValue0
	_parent *ExprSizeofValue0
}
func NewExprSizeofValue0_Block() *ExprSizeofValue0_Block {
	return &ExprSizeofValue0_Block{
	}
}

func (this *ExprSizeofValue0_Block) Read(io *kaitai.Stream, parent *ExprSizeofValue0, root *ExprSizeofValue0) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.A = tmp3
	tmp4, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.B = uint32(tmp4)
	tmp5, err := this._io.ReadBytes(int(2))
	if err != nil {
		return err
	}
	tmp5 = tmp5
	this.C = tmp5
	return err
}
