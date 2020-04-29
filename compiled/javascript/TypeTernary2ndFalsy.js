// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.TypeTernary2ndFalsy = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var TypeTernary2ndFalsy = (function() {
  function TypeTernary2ndFalsy(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  TypeTernary2ndFalsy.prototype._read = function() {
    this.intTruthy = this._io.readU1();
    this.ut = new Foo(this._io, this, this._root);
    if (false) {
      this.nullUt = new Foo(this._io, this, this._root);
    }
    this.intArray = new Array(2);
    for (var i = 0; i < 2; i++) {
      this.intArray[i] = this._io.readU1();
    }
    this.intArrayEmpty = new Array(0);
    for (var i = 0; i < 0; i++) {
      this.intArrayEmpty[i] = this._io.readU1();
    }
  }

  var Foo = TypeTernary2ndFalsy.Foo = (function() {
    function Foo(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;

      this._read();
    }
    Foo.prototype._read = function() {
      this.m = this._io.readU1();
    }

    return Foo;
  })();
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vFloatZero', {
    get: function() {
      if (this._m_vFloatZero !== undefined)
        return this._m_vFloatZero;
      this._m_vFloatZero = (this.t ? 0.0 : 3.14);
      return this._m_vFloatZero;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 't', {
    get: function() {
      if (this._m_t !== undefined)
        return this._m_t;
      this._m_t = true;
      return this._m_t;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vIntNegZero', {
    get: function() {
      if (this._m_vIntNegZero !== undefined)
        return this._m_vIntNegZero;
      this._m_vIntNegZero = (this.t ? -0 : -20);
      return this._m_vIntNegZero;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vIntZero', {
    get: function() {
      if (this._m_vIntZero !== undefined)
        return this._m_vIntZero;
      this._m_vIntZero = (this.t ? 0 : 10);
      return this._m_vIntZero;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'nullUtInst', {
    get: function() {
      if (this._m_nullUtInst !== undefined)
        return this._m_nullUtInst;
      if (false) {
        this._m_nullUtInst = this.ut;
      }
      return this._m_nullUtInst;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vNullUtInst', {
    get: function() {
      if (this._m_vNullUtInst !== undefined)
        return this._m_vNullUtInst;
      this._m_vNullUtInst = (this.t ? this.nullUtInst : this.ut);
      return this._m_vNullUtInst;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vFalse', {
    get: function() {
      if (this._m_vFalse !== undefined)
        return this._m_vFalse;
      this._m_vFalse = (this.t ? false : true);
      return this._m_vFalse;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vStrEmpty', {
    get: function() {
      if (this._m_vStrEmpty !== undefined)
        return this._m_vStrEmpty;
      this._m_vStrEmpty = (this.t ? "" : "kaitai");
      return this._m_vStrEmpty;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vIntArrayEmpty', {
    get: function() {
      if (this._m_vIntArrayEmpty !== undefined)
        return this._m_vIntArrayEmpty;
      this._m_vIntArrayEmpty = (this.t ? this.intArrayEmpty : this.intArray);
      return this._m_vIntArrayEmpty;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vNullUt', {
    get: function() {
      if (this._m_vNullUt !== undefined)
        return this._m_vNullUt;
      this._m_vNullUt = (this.t ? this.nullUt : this.ut);
      return this._m_vNullUt;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vFloatNegZero', {
    get: function() {
      if (this._m_vFloatNegZero !== undefined)
        return this._m_vFloatNegZero;
      this._m_vFloatNegZero = (this.t ? -0.0 : -2.72);
      return this._m_vFloatNegZero;
    }
  });
  Object.defineProperty(TypeTernary2ndFalsy.prototype, 'vStrWZero', {
    get: function() {
      if (this._m_vStrWZero !== undefined)
        return this._m_vStrWZero;
      this._m_vStrWZero = (this.t ? "0" : "30");
      return this._m_vStrWZero;
    }
  });

  return TypeTernary2ndFalsy;
})();
return TypeTernary2ndFalsy;
}));
