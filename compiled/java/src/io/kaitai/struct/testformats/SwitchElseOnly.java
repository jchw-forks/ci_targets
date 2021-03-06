// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package io.kaitai.struct.testformats;

import io.kaitai.struct.ByteBufferKaitaiStream;
import io.kaitai.struct.KaitaiStruct;
import io.kaitai.struct.KaitaiStream;
import java.io.IOException;

public class SwitchElseOnly extends KaitaiStruct {
    public static SwitchElseOnly fromFile(String fileName) throws IOException {
        return new SwitchElseOnly(new ByteBufferKaitaiStream(fileName));
    }

    public SwitchElseOnly(KaitaiStream _io) {
        this(_io, null, null);
    }

    public SwitchElseOnly(KaitaiStream _io, KaitaiStruct _parent) {
        this(_io, _parent, null);
    }

    public SwitchElseOnly(KaitaiStream _io, KaitaiStruct _parent, SwitchElseOnly _root) {
        super(_io);
        this._parent = _parent;
        this._root = _root == null ? this : _root;
        _read();
    }
    private void _read() {
        this.opcode = this._io.readS1();
        switch (opcode()) {
        default: {
            this.primByte = this._io.readS1();
            break;
        }
        }
        switch (opcode()) {
        default: {
            this.struct = new Data(this._io, this, _root);
            break;
        }
        }
        switch (opcode()) {
        default: {
            this._raw_structSized = this._io.readBytes(4);
            KaitaiStream _io__raw_structSized = new ByteBufferKaitaiStream(_raw_structSized);
            this.structSized = new Data(_io__raw_structSized, this, _root);
            break;
        }
        }
    }
    public static class Data extends KaitaiStruct {
        public static Data fromFile(String fileName) throws IOException {
            return new Data(new ByteBufferKaitaiStream(fileName));
        }

        public Data(KaitaiStream _io) {
            this(_io, null, null);
        }

        public Data(KaitaiStream _io, SwitchElseOnly _parent) {
            this(_io, _parent, null);
        }

        public Data(KaitaiStream _io, SwitchElseOnly _parent, SwitchElseOnly _root) {
            super(_io);
            this._parent = _parent;
            this._root = _root;
            _read();
        }
        private void _read() {
            this.value = this._io.readBytes(4);
        }
        private byte[] value;
        private SwitchElseOnly _root;
        private SwitchElseOnly _parent;
        public byte[] value() { return value; }
        public SwitchElseOnly _root() { return _root; }
        public SwitchElseOnly _parent() { return _parent; }
    }
    private byte opcode;
    private byte primByte;
    private Data struct;
    private Data structSized;
    private SwitchElseOnly _root;
    private KaitaiStruct _parent;
    private byte[] _raw_structSized;
    public byte opcode() { return opcode; }
    public byte primByte() { return primByte; }
    public Data struct() { return struct; }
    public Data structSized() { return structSized; }
    public SwitchElseOnly _root() { return _root; }
    public KaitaiStruct _parent() { return _parent; }
    public byte[] _raw_structSized() { return _raw_structSized; }
}
