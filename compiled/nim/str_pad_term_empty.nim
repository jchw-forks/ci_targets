import kaitai_struct_nim_runtime
import encodings

type
  StrPadTermEmpty* = ref StrPadTermEmptyObj
  StrPadTermEmptyObj* = object
    strPad*: string
    strTerm*: string
    strTermAndPad*: string
    strTermInclude*: string
    io*: KaitaiStream
    root*: StrPadTermEmpty
    parent*: ref RootObj

### StrPadTermEmpty ###
proc read*(_: typedesc[StrPadTermEmpty], io: KaitaiStream, root: StrPadTermEmpty, parent: ref RootObj): StrPadTermEmpty =
  result = new(StrPadTermEmpty)
  let root = if root == nil: cast[StrPadTermEmpty](result) else: root
  result.io = io
  result.root = root
  result.parent = parent
  result.strPad = convert(result.io.readBytes(20).bytesStripRight(64), srcEncoding = "UTF-8")
  result.strTerm = convert(result.io.readBytes(20).bytesTerminate(64, false), srcEncoding = "UTF-8")
  result.strTermAndPad = convert(result.io.readBytes(20).bytesStripRight(43).bytesTerminate(64, false), srcEncoding = "UTF-8")
  result.strTermInclude = convert(result.io.readBytes(20).bytesTerminate(64, true), srcEncoding = "UTF-8")

proc fromFile*(_: typedesc[StrPadTermEmpty], filename: string): StrPadTermEmpty =
  StrPadTermEmpty.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var StrPadTermEmptyObj) =
  close(x.io)

