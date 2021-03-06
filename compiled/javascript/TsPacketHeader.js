// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.TsPacketHeader = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
/**
 * describes the first 4 header bytes of a TS Packet header
 */

var TsPacketHeader = (function() {
  TsPacketHeader.AdaptationFieldControlEnum = Object.freeze({
    RESERVED: 0,
    PAYLOAD_ONLY: 1,
    ADAPTATION_FIELD_ONLY: 2,
    ADAPTATION_FIELD_AND_PAYLOAD: 3,

    0: "RESERVED",
    1: "PAYLOAD_ONLY",
    2: "ADAPTATION_FIELD_ONLY",
    3: "ADAPTATION_FIELD_AND_PAYLOAD",
  });

  function TsPacketHeader(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  TsPacketHeader.prototype._read = function() {
    this.syncByte = this._io.readU1();
    this.transportErrorIndicator = this._io.readBitsInt(1) != 0;
    this.payloadUnitStartIndicator = this._io.readBitsInt(1) != 0;
    this.transportPriority = this._io.readBitsInt(1) != 0;
    this.pid = this._io.readBitsInt(13);
    this.transportScramblingControl = this._io.readBitsInt(2);
    this.adaptationFieldControl = this._io.readBitsInt(2);
    this.continuityCounter = this._io.readBitsInt(4);
    this._io.alignToByte();
    this.tsPacketRemain = this._io.readBytes(184);
  }

  return TsPacketHeader;
})();
return TsPacketHeader;
}));
