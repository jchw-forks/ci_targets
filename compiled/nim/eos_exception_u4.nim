import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  EosExceptionU4* = ref object of KaitaiStruct
    envelope*: EosExceptionU4_Data
    parent*: KaitaiStruct
    rawEnvelope*: string
  EosExceptionU4_Data* = ref object of KaitaiStruct
    prebuf*: string
    failInt*: uint32
    parent*: EosExceptionU4

proc read*(_: typedesc[EosExceptionU4], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): EosExceptionU4
proc read*(_: typedesc[EosExceptionU4_Data], io: KaitaiStream, root: KaitaiStruct, parent: EosExceptionU4): EosExceptionU4_Data


proc read*(_: typedesc[EosExceptionU4], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): EosExceptionU4 =
  template this: untyped = result
  this = new(EosExceptionU4)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.rawEnvelope = this.io.readBytes(int(6))
  let rawEnvelopeIo = newKaitaiStringStream(this.rawEnvelope)
  this.envelope = EosExceptionU4_Data.read(rawEnvelopeIo, this.root, this)

proc fromFile*(_: typedesc[EosExceptionU4], filename: string): EosExceptionU4 =
  EosExceptionU4.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[EosExceptionU4_Data], io: KaitaiStream, root: KaitaiStruct, parent: EosExceptionU4): EosExceptionU4_Data =
  template this: untyped = result
  this = new(EosExceptionU4_Data)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.prebuf = this.io.readBytes(int(3))
  this.failInt = this.io.readU4le()

proc fromFile*(_: typedesc[EosExceptionU4_Data], filename: string): EosExceptionU4_Data =
  EosExceptionU4_Data.read(newKaitaiFileStream(filename), nil, nil)

