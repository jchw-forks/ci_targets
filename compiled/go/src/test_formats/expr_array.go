// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ExprArray struct {
	Aint []uint32
	Afloat []float64
	Astr []string
	_io *kaitai.Stream
	_root *ExprArray
	_parent interface{}
	_f_aintFirst bool
	aintFirst uint32
	_f_afloatSize bool
	afloatSize int
	_f_astrSize bool
	astrSize int
	_f_aintMin bool
	aintMin uint32
	_f_afloatMin bool
	afloatMin float64
	_f_aintSize bool
	aintSize int
	_f_aintLast bool
	aintLast uint32
	_f_afloatLast bool
	afloatLast float64
	_f_astrFirst bool
	astrFirst string
	_f_astrLast bool
	astrLast string
	_f_aintMax bool
	aintMax uint32
	_f_afloatFirst bool
	afloatFirst float64
	_f_astrMin bool
	astrMin string
	_f_astrMax bool
	astrMax string
	_f_afloatMax bool
	afloatMax float64
}
func NewExprArray() *ExprArray {
	return &ExprArray{
	}
}

func (this *ExprArray) Read(io *kaitai.Stream, parent interface{}, root *ExprArray) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	this.Aint = make([]uint32, 4)
	for i := range this.Aint {
		tmp1, err := this._io.ReadU4le()
		if err != nil {
			return err
		}
		this.Aint[i] = tmp1
	}
	this.Afloat = make([]float64, 3)
	for i := range this.Afloat {
		tmp2, err := this._io.ReadF8le()
		if err != nil {
			return err
		}
		this.Afloat[i] = tmp2
	}
	this.Astr = make([]string, 3)
	for i := range this.Astr {
		tmp3, err := this._io.ReadBytesTerm(0, false, true, true)
		if err != nil {
			return err
		}
		this.Astr[i] = string(tmp3)
	}
	return err
}
func (this *ExprArray) AintFirst() (v uint32, err error) {
	if (this._f_aintFirst) {
		return this.aintFirst, nil
	}
	this.aintFirst = uint32(this.Aint[0])
	this._f_aintFirst = true
	return this.aintFirst, nil
}
func (this *ExprArray) AfloatSize() (v int, err error) {
	if (this._f_afloatSize) {
		return this.afloatSize, nil
	}
	this.afloatSize = int(len(this.Afloat))
	this._f_afloatSize = true
	return this.afloatSize, nil
}
func (this *ExprArray) AstrSize() (v int, err error) {
	if (this._f_astrSize) {
		return this.astrSize, nil
	}
	this.astrSize = int(len(this.Astr))
	this._f_astrSize = true
	return this.astrSize, nil
}
func (this *ExprArray) AintMin() (v uint32, err error) {
	if (this._f_aintMin) {
		return this.aintMin, nil
	}
	tmp4 := this.Aint[0]
	for _, tmp5 := range this.Aint {
		if tmp4 > tmp5 {
			tmp4 = tmp5
		}
	}
	this.aintMin = uint32(tmp4)
	this._f_aintMin = true
	return this.aintMin, nil
}
func (this *ExprArray) AfloatMin() (v float64, err error) {
	if (this._f_afloatMin) {
		return this.afloatMin, nil
	}
	tmp6 := this.Afloat[0]
	for _, tmp7 := range this.Afloat {
		if tmp6 > tmp7 {
			tmp6 = tmp7
		}
	}
	this.afloatMin = float64(tmp6)
	this._f_afloatMin = true
	return this.afloatMin, nil
}
func (this *ExprArray) AintSize() (v int, err error) {
	if (this._f_aintSize) {
		return this.aintSize, nil
	}
	this.aintSize = int(len(this.Aint))
	this._f_aintSize = true
	return this.aintSize, nil
}
func (this *ExprArray) AintLast() (v uint32, err error) {
	if (this._f_aintLast) {
		return this.aintLast, nil
	}
	tmp8 := this.Aint
	this.aintLast = uint32(tmp8[len(tmp8) - 1])
	this._f_aintLast = true
	return this.aintLast, nil
}
func (this *ExprArray) AfloatLast() (v float64, err error) {
	if (this._f_afloatLast) {
		return this.afloatLast, nil
	}
	tmp9 := this.Afloat
	this.afloatLast = float64(tmp9[len(tmp9) - 1])
	this._f_afloatLast = true
	return this.afloatLast, nil
}
func (this *ExprArray) AstrFirst() (v string, err error) {
	if (this._f_astrFirst) {
		return this.astrFirst, nil
	}
	this.astrFirst = string(this.Astr[0])
	this._f_astrFirst = true
	return this.astrFirst, nil
}
func (this *ExprArray) AstrLast() (v string, err error) {
	if (this._f_astrLast) {
		return this.astrLast, nil
	}
	tmp10 := this.Astr
	this.astrLast = string(tmp10[len(tmp10) - 1])
	this._f_astrLast = true
	return this.astrLast, nil
}
func (this *ExprArray) AintMax() (v uint32, err error) {
	if (this._f_aintMax) {
		return this.aintMax, nil
	}
	tmp11 := this.Aint[0]
	for _, tmp12 := range this.Aint {
		if tmp11 < tmp12 {
			tmp11 = tmp12
		}
	}
	this.aintMax = uint32(tmp11)
	this._f_aintMax = true
	return this.aintMax, nil
}
func (this *ExprArray) AfloatFirst() (v float64, err error) {
	if (this._f_afloatFirst) {
		return this.afloatFirst, nil
	}
	this.afloatFirst = float64(this.Afloat[0])
	this._f_afloatFirst = true
	return this.afloatFirst, nil
}
func (this *ExprArray) AstrMin() (v string, err error) {
	if (this._f_astrMin) {
		return this.astrMin, nil
	}
	tmp13 := this.Astr[0]
	for _, tmp14 := range this.Astr {
		if tmp13 > tmp14 {
			tmp13 = tmp14
		}
	}
	this.astrMin = string(tmp13)
	this._f_astrMin = true
	return this.astrMin, nil
}
func (this *ExprArray) AstrMax() (v string, err error) {
	if (this._f_astrMax) {
		return this.astrMax, nil
	}
	tmp15 := this.Astr[0]
	for _, tmp16 := range this.Astr {
		if tmp15 < tmp16 {
			tmp15 = tmp16
		}
	}
	this.astrMax = string(tmp15)
	this._f_astrMax = true
	return this.astrMax, nil
}
func (this *ExprArray) AfloatMax() (v float64, err error) {
	if (this._f_afloatMax) {
		return this.afloatMax, nil
	}
	tmp17 := this.Afloat[0]
	for _, tmp18 := range this.Afloat {
		if tmp17 < tmp18 {
			tmp17 = tmp18
		}
	}
	this.afloatMax = float64(tmp17)
	this._f_afloatMax = true
	return this.afloatMax, nil
}
