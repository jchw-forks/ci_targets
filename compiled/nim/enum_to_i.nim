import kaitai_struct_nim_runtime
import options

type
  EnumToI* = ref object of KaitaiStruct
    pet1*: EnumToI_Animal
    pet2*: EnumToI_Animal
    parent*: KaitaiStruct
    pet1IInst*: Option[int]
    pet1ModInst*: Option[int]
    oneLtTwoInst*: Option[bool]
  EnumToI_Animal* = enum
    dog = 4
    cat = 7
    chicken = 12

proc read*(_: typedesc[EnumToI], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): EnumToI

proc pet1I*(this: EnumToI): int
proc pet1Mod*(this: EnumToI): int
proc oneLtTwo*(this: EnumToI): bool

proc read*(_: typedesc[EnumToI], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): EnumToI =
  template this: untyped = result
  this = new(EnumToI)
  let root = if root == nil: cast[EnumToI](this) else: cast[EnumToI](root)
  this.io = io
  this.root = root
  this.parent = parent

  let pet1Expr = EnumToI_Animal(this.io.readU4le())
  this.pet1 = pet1Expr
  let pet2Expr = EnumToI_Animal(this.io.readU4le())
  this.pet2 = pet2Expr

proc pet1I(this: EnumToI): int = 
  if isSome(this.pet1IInst):
    return get(this.pet1IInst)
  let pet1IInstExpr = int(ord(this.pet1))
  this.pet1IInst = pet1IInstExpr
  if isSome(this.pet1IInst):
    return get(this.pet1IInst)

proc pet1Mod(this: EnumToI): int = 
  if isSome(this.pet1ModInst):
    return get(this.pet1ModInst)
  let pet1ModInstExpr = int((ord(this.pet1) + 32768))
  this.pet1ModInst = pet1ModInstExpr
  if isSome(this.pet1ModInst):
    return get(this.pet1ModInst)

proc oneLtTwo(this: EnumToI): bool = 
  if isSome(this.oneLtTwoInst):
    return get(this.oneLtTwoInst)
  let oneLtTwoInstExpr = bool(ord(this.pet1) < ord(this.pet2))
  this.oneLtTwoInst = oneLtTwoInstExpr
  if isSome(this.oneLtTwoInst):
    return get(this.oneLtTwoInst)

proc fromFile*(_: typedesc[EnumToI], filename: string): EnumToI =
  EnumToI.read(newKaitaiFileStream(filename), nil, nil)

