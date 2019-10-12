import ../../runtime/nim/kaitai



type
  ImportedAndRel* = ref ImportedAndRelObj
  ImportedAndRelObj* = object
    io: KaitaiStream
    root*: ImportedAndRel
    parent*: ref RootObj
    one*: uint8
    two*: ImportedRoot

# ImportedAndRel
proc read*(_: typedesc[ImportedAndRel], io: KaitaiStream, root: ImportedAndRel, parent: ref RootObj): owned ImportedAndRel =
  result = new(ImportedAndRel)
  let root = if root == nil: cast[ImportedAndRel](result) else: root
  result.io = io
  result.root = root
  result.parent = parent

  result.one = readU1(io)
  result.two = ImportedRoot.read(io)


proc fromFile*(_: typedesc[ImportedAndRel], filename: string): owned ImportedAndRel =
  ImportedAndRel.read(newKaitaiStream(filename), nil, nil)

proc `=destroy`(x: var ImportedAndRelObj) =
  close(x.io)
