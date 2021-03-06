import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  ParamsPassStruct* = ref object of KaitaiStruct
    first*: ParamsPassStruct_Block
    one*: ParamsPassStruct_StructType
    parent*: KaitaiStruct
  ParamsPassStruct_Block* = ref object of KaitaiStruct
    foo*: uint8
    parent*: ParamsPassStruct
  ParamsPassStruct_StructType* = ref object of KaitaiStruct
    bar*: ParamsPassStruct_StructType_Baz
    foo*: KaitaiStruct
    parent*: ParamsPassStruct
  ParamsPassStruct_StructType_Baz* = ref object of KaitaiStruct
    qux*: uint8
    foo*: KaitaiStruct
    parent*: ParamsPassStruct_StructType

proc read*(_: typedesc[ParamsPassStruct], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ParamsPassStruct
proc read*(_: typedesc[ParamsPassStruct_Block], io: KaitaiStream, root: KaitaiStruct, parent: ParamsPassStruct): ParamsPassStruct_Block
proc read*(_: typedesc[ParamsPassStruct_StructType], io: KaitaiStream, root: KaitaiStruct, parent: ParamsPassStruct): ParamsPassStruct_StructType
proc read*(_: typedesc[ParamsPassStruct_StructType_Baz], io: KaitaiStream, root: KaitaiStruct, parent: ParamsPassStruct_StructType): ParamsPassStruct_StructType_Baz


proc read*(_: typedesc[ParamsPassStruct], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ParamsPassStruct =
  template this: untyped = result
  this = new(ParamsPassStruct)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.first = ParamsPassStruct_Block.read(this.io, this.root, this)
  this.one = ParamsPassStruct_StructType.read(this.io, this.root, this, this.first)

proc fromFile*(_: typedesc[ParamsPassStruct], filename: string): ParamsPassStruct =
  ParamsPassStruct.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[ParamsPassStruct_Block], io: KaitaiStream, root: KaitaiStruct, parent: ParamsPassStruct): ParamsPassStruct_Block =
  template this: untyped = result
  this = new(ParamsPassStruct_Block)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.foo = this.io.readU1()

proc fromFile*(_: typedesc[ParamsPassStruct_Block], filename: string): ParamsPassStruct_Block =
  ParamsPassStruct_Block.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[ParamsPassStruct_StructType], io: KaitaiStream, root: KaitaiStruct, parent: ParamsPassStruct): ParamsPassStruct_StructType =
  template this: untyped = result
  this = new(ParamsPassStruct_StructType)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.bar = ParamsPassStruct_StructType_Baz.read(this.io, this.root, this, this.foo)

proc fromFile*(_: typedesc[ParamsPassStruct_StructType], filename: string): ParamsPassStruct_StructType =
  ParamsPassStruct_StructType.read(newKaitaiFileStream(filename), nil, nil)

proc read*(_: typedesc[ParamsPassStruct_StructType_Baz], io: KaitaiStream, root: KaitaiStruct, parent: ParamsPassStruct_StructType): ParamsPassStruct_StructType_Baz =
  template this: untyped = result
  this = new(ParamsPassStruct_StructType_Baz)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.qux = this.io.readU1()

proc fromFile*(_: typedesc[ParamsPassStruct_StructType_Baz], filename: string): ParamsPassStruct_StructType_Baz =
  ParamsPassStruct_StructType_Baz.read(newKaitaiFileStream(filename), nil, nil)

