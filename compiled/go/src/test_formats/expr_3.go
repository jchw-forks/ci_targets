// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type Expr3 struct {
	One uint8
	Two string
	_io *kaitai.Stream
	_root *Expr3
	_parent interface{}
	_f_three bool
	three string
	_f_isStrGe bool
	isStrGe bool
	_f_isStrNe bool
	isStrNe bool
	_f_isStrGt bool
	isStrGt bool
	_f_isStrLe bool
	isStrLe bool
	_f_isStrLt2 bool
	isStrLt2 bool
	_f_testNot bool
	testNot bool
	_f_isStrLt bool
	isStrLt bool
	_f_four bool
	four string
	_f_isStrEq bool
	isStrEq bool
}
func NewExpr3() *Expr3 {
	return &Expr3{
	}
}

func (this *Expr3) Read(io *kaitai.Stream, parent interface{}, root *Expr3) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.One = tmp1
	tmp2, err := this._io.ReadBytes(int(3))
	if err != nil {
		return err
	}
	tmp2 = tmp2
	this.Two = string(tmp2)
	return err
}
func (this *Expr3) Three() (v string, err error) {
	if (this._f_three) {
		return this.three, nil
	}
	this.three = string("@" + this.Two)
	this._f_three = true
	return this.three, nil
}
func (this *Expr3) IsStrGe() (v bool, err error) {
	if (this._f_isStrGe) {
		return this.isStrGe, nil
	}
	this.isStrGe = bool(this.Two >= "ACK2")
	this._f_isStrGe = true
	return this.isStrGe, nil
}
func (this *Expr3) IsStrNe() (v bool, err error) {
	if (this._f_isStrNe) {
		return this.isStrNe, nil
	}
	this.isStrNe = bool(this.Two != "ACK")
	this._f_isStrNe = true
	return this.isStrNe, nil
}
func (this *Expr3) IsStrGt() (v bool, err error) {
	if (this._f_isStrGt) {
		return this.isStrGt, nil
	}
	this.isStrGt = bool(this.Two > "ACK2")
	this._f_isStrGt = true
	return this.isStrGt, nil
}
func (this *Expr3) IsStrLe() (v bool, err error) {
	if (this._f_isStrLe) {
		return this.isStrLe, nil
	}
	this.isStrLe = bool(this.Two <= "ACK2")
	this._f_isStrLe = true
	return this.isStrLe, nil
}
func (this *Expr3) IsStrLt2() (v bool, err error) {
	if (this._f_isStrLt2) {
		return this.isStrLt2, nil
	}
	tmp3, err := this.Three()
	if err != nil {
		return false, err
	}
	this.isStrLt2 = bool(tmp3 < this.Two)
	this._f_isStrLt2 = true
	return this.isStrLt2, nil
}
func (this *Expr3) TestNot() (v bool, err error) {
	if (this._f_testNot) {
		return this.testNot, nil
	}
	this.testNot = bool(!(false))
	this._f_testNot = true
	return this.testNot, nil
}
func (this *Expr3) IsStrLt() (v bool, err error) {
	if (this._f_isStrLt) {
		return this.isStrLt, nil
	}
	this.isStrLt = bool(this.Two < "ACK2")
	this._f_isStrLt = true
	return this.isStrLt, nil
}
func (this *Expr3) Four() (v string, err error) {
	if (this._f_four) {
		return this.four, nil
	}
	this.four = string("_" + this.Two + "_")
	this._f_four = true
	return this.four, nil
}
func (this *Expr3) IsStrEq() (v bool, err error) {
	if (this._f_isStrEq) {
		return this.isStrEq, nil
	}
	this.isStrEq = bool(this.Two == "ACK")
	this._f_isStrEq = true
	return this.isStrEq, nil
}
