import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  ValidFailEqInt* = ref object of KaitaiStruct
    foo*: uint8
    parent*: KaitaiStruct

proc read*(_: typedesc[ValidFailEqInt], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ValidFailEqInt


proc read*(_: typedesc[ValidFailEqInt], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ValidFailEqInt =
  template this: untyped = result
  this = new(ValidFailEqInt)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.foo = this.io.readU1()

proc fromFile*(_: typedesc[ValidFailEqInt], filename: string): ValidFailEqInt =
  ValidFailEqInt.read(newKaitaiFileStream(filename), nil, nil)

