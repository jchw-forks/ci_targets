import kaitai_struct_nim_runtime
import options
import encodings

type
  EnumIf_Operation* = ref EnumIf_OperationObj
  EnumIf_OperationObj* = object
    opcode*: EnumIf_Opcodes
    argTuple*: EnumIf_ArgTuple
    argStr*: EnumIf_ArgStr
    io*: KaitaiStream
    root*: EnumIf
    parent*: EnumIf
  EnumIf_ArgTuple* = ref EnumIf_ArgTupleObj
  EnumIf_ArgTupleObj* = object
    num1*: uint8
    num2*: uint8
    io*: KaitaiStream
    root*: EnumIf
    parent*: EnumIf_Operation
  EnumIf_ArgStr* = ref EnumIf_ArgStrObj
  EnumIf_ArgStrObj* = object
    len*: uint8
    str*: string
    io*: KaitaiStream
    root*: EnumIf
    parent*: EnumIf_Operation
  EnumIf* = ref EnumIfObj
  EnumIfObj* = object
    op1*: EnumIf_Operation
    op2*: EnumIf_Operation
    op3*: EnumIf_Operation
    io*: KaitaiStream
    root*: EnumIf
    parent*: ref RootObj
  EnumIf_opcodes* = enum
    a_string = 83
    a_tuple = 84

## EnumIf_Operation
proc read*(_: typedesc[EnumIf_Operation], io: KaitaiStream, root: EnumIf, parent: EnumIf): EnumIf_Operation =
  let this = new(EnumIf_Operation)
  let root = if root == nil: cast[EnumIf](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.opcode = EnumIf_Opcodes(this.io.readU1())
  if this.opcode == EnumIf_Opcodes.a_tuple:
    this.argTuple = EnumIf_ArgTuple.read(this.io, this.root, this)
  if this.opcode == EnumIf_Opcodes.a_string:
    this.argStr = EnumIf_ArgStr.read(this.io, this.root, this)
  result = this

proc fromFile*(_: typedesc[EnumIf_Operation], filename: string): EnumIf_Operation =
  EnumIf_Operation.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var EnumIf_OperationObj) =
  close(x.io)

## EnumIf_ArgTuple
proc read*(_: typedesc[EnumIf_ArgTuple], io: KaitaiStream, root: EnumIf, parent: EnumIf_Operation): EnumIf_ArgTuple =
  let this = new(EnumIf_ArgTuple)
  let root = if root == nil: cast[EnumIf](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.num1 = this.io.readU1()
  this.num2 = this.io.readU1()
  result = this

proc fromFile*(_: typedesc[EnumIf_ArgTuple], filename: string): EnumIf_ArgTuple =
  EnumIf_ArgTuple.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var EnumIf_ArgTupleObj) =
  close(x.io)

## EnumIf_ArgStr
proc read*(_: typedesc[EnumIf_ArgStr], io: KaitaiStream, root: EnumIf, parent: EnumIf_Operation): EnumIf_ArgStr =
  let this = new(EnumIf_ArgStr)
  let root = if root == nil: cast[EnumIf](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.len = this.io.readU1()
  this.str = convert(this.io.readBytes(int(this.len)), srcEncoding = "UTF-8")
  result = this

proc fromFile*(_: typedesc[EnumIf_ArgStr], filename: string): EnumIf_ArgStr =
  EnumIf_ArgStr.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var EnumIf_ArgStrObj) =
  close(x.io)

## EnumIf
proc read*(_: typedesc[EnumIf], io: KaitaiStream, root: EnumIf, parent: ref RootObj): EnumIf =
  let this = new(EnumIf)
  let root = if root == nil: cast[EnumIf](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.op1 = EnumIf_Operation.read(this.io, this.root, this)
  this.op2 = EnumIf_Operation.read(this.io, this.root, this)
  this.op3 = EnumIf_Operation.read(this.io, this.root, this)
  result = this

proc fromFile*(_: typedesc[EnumIf], filename: string): EnumIf =
  EnumIf.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var EnumIfObj) =
  close(x.io)

