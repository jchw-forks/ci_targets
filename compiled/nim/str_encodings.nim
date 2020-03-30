import kaitai_struct_nim_runtime
import options
import encodings

type
  StrEncodings* = ref StrEncodingsObj
  StrEncodingsObj* = object
    lenOf1*: uint16
    str1*: string
    lenOf2*: uint16
    str2*: string
    lenOf3*: uint16
    str3*: string
    lenOf4*: uint16
    str4*: string
    io*: KaitaiStream
    root*: StrEncodings
    parent*: ref RootObj

## StrEncodings
proc read*(_: typedesc[StrEncodings], io: KaitaiStream, root: StrEncodings, parent: ref RootObj): StrEncodings =
  let this = new(StrEncodings)
  let root = if root == nil: cast[StrEncodings](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.lenOf1 = this.io.readU2le()
  this.str1 = convert(this.io.readBytes(int(this.lenOf1)), srcEncoding = "ASCII")
  this.lenOf2 = this.io.readU2le()
  this.str2 = convert(this.io.readBytes(int(this.lenOf2)), srcEncoding = "UTF-8")
  this.lenOf3 = this.io.readU2le()
  this.str3 = convert(this.io.readBytes(int(this.lenOf3)), srcEncoding = "SJIS")
  this.lenOf4 = this.io.readU2le()
  this.str4 = convert(this.io.readBytes(int(this.lenOf4)), srcEncoding = "CP437")
  result = this

proc fromFile*(_: typedesc[StrEncodings], filename: string): StrEncodings =
  StrEncodings.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var StrEncodingsObj) =
  close(x.io)

