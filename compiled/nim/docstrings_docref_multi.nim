import kaitai_struct_nim_runtime
import options

type
  DocstringsDocrefMulti* = ref DocstringsDocrefMultiObj
  DocstringsDocrefMultiObj* = object
    io*: KaitaiStream
    root*: DocstringsDocrefMulti
    parent*: ref RootObj

## DocstringsDocrefMulti
proc read*(_: typedesc[DocstringsDocrefMulti], io: KaitaiStream, root: DocstringsDocrefMulti, parent: ref RootObj): DocstringsDocrefMulti =
  let this = new(DocstringsDocrefMulti)
  let root = if root == nil: cast[DocstringsDocrefMulti](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  result = this

proc fromFile*(_: typedesc[DocstringsDocrefMulti], filename: string): DocstringsDocrefMulti =
  DocstringsDocrefMulti.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var DocstringsDocrefMultiObj) =
  close(x.io)

