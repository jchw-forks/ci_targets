import ../../runtime/nim/kaitai
import options

{.experimental: "dotOperators".}

type
  EnumToIClassBorder2* = ref EnumToIClassBorder2Obj
  EnumToIClassBorder2Obj* = object
    io: KaitaiStream
    root*: EnumToIClassBorder2
    parent*: ref RootObj
    isDogInst: proc(): bool

# EnumToIClassBorder2
template `.`*(a: EnumToIClassBorder2, b: untyped): untyped =
  (a.`b inst`)()

proc read*(_: typedesc[EnumToIClassBorder2], io: KaitaiStream, root: EnumToIClassBorder2, parent: ref RootObj): owned EnumToIClassBorder2 =
  result = new(EnumToIClassBorder2)
  let root = if root == nil: cast[EnumToIClassBorder2](result) else: root
  result.io = io
  result.root = root
  result.parent = parent


  let shadow = result
  var isDog: Option[bool]
  result.isDogInst = proc(): bool =
    if isNone(isDog):
      isDog = some(bool(ord(shadow.parent.someDog) == 4))
    get(isDog)

proc fromFile*(_: typedesc[EnumToIClassBorder2], filename: string): owned EnumToIClassBorder2 =
  EnumToIClassBorder2.read(newKaitaiStream(filename), nil, nil)

proc `=destroy`(x: var EnumToIClassBorder2Obj) =
  close(x.io)
