import kaitai_struct_nim_runtime
import options
import encodings

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  ParamsCallShort* = ref object of KaitaiStruct
    buf1*: ParamsCallShort_MyStr1
    buf2*: ParamsCallShort_MyStr2
    parent*: KaitaiStruct
  ParamsCallShort_MyStr1* = ref object of KaitaiStruct
    body*: string
    len*: uint32
    parent*: ParamsCallShort
  ParamsCallShort_MyStr2* = ref object of KaitaiStruct
    body*: string
    trailer*: uint8
    len*: uint32
    hasTrailer*: bool
    parent*: ParamsCallShort

proc read*(_: typedesc[ParamsCallShort], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ParamsCallShort
proc read*(_: typedesc[ParamsCallShort_MyStr1], io: KaitaiStream, root: KaitaiStruct, parent: ParamsCallShort): ParamsCallShort_MyStr1
proc read*(_: typedesc[ParamsCallShort_MyStr2], io: KaitaiStream, root: KaitaiStruct, parent: ParamsCallShort): ParamsCallShort_MyStr2


proc read*(_: typedesc[ParamsCallShort], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ParamsCallShort =
  template this: untyped = result
  this = new(ParamsCallShort)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.buf1 = ParamsCallShort_MyStr1.read(this.io, this.root, this, 5)
  this.buf2 = ParamsCallShort_MyStr2.read(this.io, this.root, this, (2 + 3), true)

proc fromFile*(_: typedesc[ParamsCallShort], filename: string): ParamsCallShort =
  ParamsCallShort.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[ParamsCallShort_MyStr1], io: KaitaiStream, root: KaitaiStruct, parent: ParamsCallShort): ParamsCallShort_MyStr1 =
  template this: untyped = result
  this = new(ParamsCallShort_MyStr1)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.body = convert(this.io.readBytes(int(this.len)), srcEncoding = "UTF-8")

proc fromFile*(_: typedesc[ParamsCallShort_MyStr1], filename: string): ParamsCallShort_MyStr1 =
  ParamsCallShort_MyStr1.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[ParamsCallShort_MyStr2], io: KaitaiStream, root: KaitaiStruct, parent: ParamsCallShort): ParamsCallShort_MyStr2 =
  template this: untyped = result
  this = new(ParamsCallShort_MyStr2)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.body = convert(this.io.readBytes(int(this.len)), srcEncoding = "UTF-8")
  if this.hasTrailer:
    this.trailer = this.io.readU1()

proc fromFile*(_: typedesc[ParamsCallShort_MyStr2], filename: string): ParamsCallShort_MyStr2 =
  ParamsCallShort_MyStr2.read(newKaitaiFileStream(filename), nil, nil)

