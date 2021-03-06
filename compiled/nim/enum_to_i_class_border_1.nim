import kaitai_struct_nim_runtime
import options

import "enum_to_i_class_border_2"
template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

defineEnum(EnumToIClassBorder1_animal)
const
  dog* = EnumToIClassBorder1_animal(4)
  cat* = EnumToIClassBorder1_animal(7)
  chicken* = EnumToIClassBorder1_animal(12)

type
  EnumToIClassBorder1* = ref object of KaitaiStruct
    pet1*: EnumToIClassBorder1_Animal
    pet2*: EnumToIClassBorder1_Animal
    parent*: KaitaiStruct
    someDogInst*: Option[EnumToIClassBorder1_Animal]
    checkerInst*: Option[EnumToIClassBorder2]

proc read*(_: typedesc[EnumToIClassBorder1], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): EnumToIClassBorder1

proc someDog*(this: EnumToIClassBorder1): EnumToIClassBorder1_Animal
proc checker*(this: EnumToIClassBorder1): EnumToIClassBorder2

proc read*(_: typedesc[EnumToIClassBorder1], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): EnumToIClassBorder1 =
  template this: untyped = result
  this = new(EnumToIClassBorder1)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  this.pet1 = EnumToIClassBorder1_Animal(this.io.readU4le())
  this.pet2 = EnumToIClassBorder1_Animal(this.io.readU4le())

proc someDog(this: EnumToIClassBorder1): EnumToIClassBorder1_Animal = 
  if isSome(this.someDogInst):
    return get(this.someDogInst)
  this.someDogInst = some(EnumToIClassBorder1_Animal(4))
  if isSome(this.someDogInst):
    return get(this.someDogInst)

proc checker(this: EnumToIClassBorder1): EnumToIClassBorder2 = 
  if isSome(this.checkerInst):
    return get(this.checkerInst)
  let pos = this.io.pos()
  this.io.seek(int(0))
  this.checkerInst = some(EnumToIClassBorder2.read(this.io, this.root, this, this._root))
  this.io.seek(pos)
  if isSome(this.checkerInst):
    return get(this.checkerInst)

proc fromFile*(_: typedesc[EnumToIClassBorder1], filename: string): EnumToIClassBorder1 =
  EnumToIClassBorder1.read(newKaitaiFileStream(filename), nil, nil)

