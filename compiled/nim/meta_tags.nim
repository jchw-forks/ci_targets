import kaitai_struct_nim_runtime
import options

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  MetaTags* = ref object of KaitaiStruct
    parent*: KaitaiStruct

proc read*(_: typedesc[MetaTags], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): MetaTags


proc read*(_: typedesc[MetaTags], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): MetaTags =
  template this: untyped = result
  this = new(MetaTags)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent


proc fromFile*(_: typedesc[MetaTags], filename: string): MetaTags =
  MetaTags.read(newKaitaiFileStream(filename), nil, nil)

