// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.DebugSwitchUser = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var DebugSwitchUser = (function() {
  function DebugSwitchUser(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;
    this._debug = {};

  }
  DebugSwitchUser.prototype._read = function() {
    this._debug.code = { start: this._io.pos, ioOffset: this._io._byteOffset };
    this.code = this._io.readU1();
    this._debug.code.end = this._io.pos;
    this._debug.data = { start: this._io.pos, ioOffset: this._io._byteOffset };
    switch (this.code) {
    case 1:
      this.data = new One(this._io, this, this._root);
      this.data._read();
      break;
    case 2:
      this.data = new Two(this._io, this, this._root);
      this.data._read();
      break;
    }
    this._debug.data.end = this._io.pos;
  }

  var One = DebugSwitchUser.One = (function() {
    function One(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;
      this._debug = {};

    }
    One.prototype._read = function() {
      this._debug.val = { start: this._io.pos, ioOffset: this._io._byteOffset };
      this.val = this._io.readS2le();
      this._debug.val.end = this._io.pos;
    }

    return One;
  })();

  var Two = DebugSwitchUser.Two = (function() {
    function Two(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;
      this._debug = {};

    }
    Two.prototype._read = function() {
      this._debug.val = { start: this._io.pos, ioOffset: this._io._byteOffset };
      this.val = this._io.readU2le();
      this._debug.val.end = this._io.pos;
    }

    return Two;
  })();

  return DebugSwitchUser;
})();
return DebugSwitchUser;
}));
