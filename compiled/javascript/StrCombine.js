// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.StrCombine = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var StrCombine = (function() {
  function StrCombine(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  StrCombine.prototype._read = function() {
    this.strTerm = KaitaiStream.bytesToStr(this._io.readBytesTerm(124, false, true, true), "ASCII");
    this.strLimit = KaitaiStream.bytesToStr(this._io.readBytes(4), "ASCII");
    this.strEos = KaitaiStream.bytesToStr(this._io.readBytesFull(), "ASCII");
  }
  Object.defineProperty(StrCombine.prototype, 'limitOrCalcBytes', {
    get: function() {
      if (this._m_limitOrCalcBytes !== undefined)
        return this._m_limitOrCalcBytes;
      this._m_limitOrCalcBytes = (true ? this.strLimit : this.strCalcBytes);
      return this._m_limitOrCalcBytes;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'limitOrCalc', {
    get: function() {
      if (this._m_limitOrCalc !== undefined)
        return this._m_limitOrCalc;
      this._m_limitOrCalc = (false ? this.strLimit : this.strCalc);
      return this._m_limitOrCalc;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'termOrLimit', {
    get: function() {
      if (this._m_termOrLimit !== undefined)
        return this._m_termOrLimit;
      this._m_termOrLimit = (true ? this.strTerm : this.strLimit);
      return this._m_termOrLimit;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'limitOrEos', {
    get: function() {
      if (this._m_limitOrEos !== undefined)
        return this._m_limitOrEos;
      this._m_limitOrEos = (true ? this.strLimit : this.strEos);
      return this._m_limitOrEos;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'calcOrCalcBytes', {
    get: function() {
      if (this._m_calcOrCalcBytes !== undefined)
        return this._m_calcOrCalcBytes;
      this._m_calcOrCalcBytes = (false ? this.strCalc : this.strCalcBytes);
      return this._m_calcOrCalcBytes;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'strCalcBytes', {
    get: function() {
      if (this._m_strCalcBytes !== undefined)
        return this._m_strCalcBytes;
      this._m_strCalcBytes = KaitaiStream.bytesToStr(this.calcBytes, "ASCII");
      return this._m_strCalcBytes;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'eosOrCalc', {
    get: function() {
      if (this._m_eosOrCalc !== undefined)
        return this._m_eosOrCalc;
      this._m_eosOrCalc = (false ? this.strEos : this.strCalc);
      return this._m_eosOrCalc;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'termOrCalc', {
    get: function() {
      if (this._m_termOrCalc !== undefined)
        return this._m_termOrCalc;
      this._m_termOrCalc = (true ? this.strTerm : this.strCalc);
      return this._m_termOrCalc;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'termOrCalcBytes', {
    get: function() {
      if (this._m_termOrCalcBytes !== undefined)
        return this._m_termOrCalcBytes;
      this._m_termOrCalcBytes = (false ? this.strTerm : this.strCalcBytes);
      return this._m_termOrCalcBytes;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'termOrEos', {
    get: function() {
      if (this._m_termOrEos !== undefined)
        return this._m_termOrEos;
      this._m_termOrEos = (false ? this.strTerm : this.strEos);
      return this._m_termOrEos;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'strCalc', {
    get: function() {
      if (this._m_strCalc !== undefined)
        return this._m_strCalc;
      this._m_strCalc = "bar";
      return this._m_strCalc;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'eosOrCalcBytes', {
    get: function() {
      if (this._m_eosOrCalcBytes !== undefined)
        return this._m_eosOrCalcBytes;
      this._m_eosOrCalcBytes = (true ? this.strEos : this.strCalcBytes);
      return this._m_eosOrCalcBytes;
    }
  });
  Object.defineProperty(StrCombine.prototype, 'calcBytes', {
    get: function() {
      if (this._m_calcBytes !== undefined)
        return this._m_calcBytes;
      this._m_calcBytes = [98, 97, 122];
      return this._m_calcBytes;
    }
  });

  return StrCombine;
})();
return StrCombine;
}));
