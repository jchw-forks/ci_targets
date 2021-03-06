import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  ValidFailMaxInt* = ref object of KaitaiStruct
    foo*: uint8
    parent*: KaitaiStruct

proc read*(_: typedesc[ValidFailMaxInt], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ValidFailMaxInt


proc read*(_: typedesc[ValidFailMaxInt], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ValidFailMaxInt =
  template this: untyped = result
  this = new(ValidFailMaxInt)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.foo = this.io.readU1()

proc fromFile*(_: typedesc[ValidFailMaxInt], filename: string): ValidFailMaxInt =
  ValidFailMaxInt.read(newKaitaiFileStream(filename), nil, nil)

