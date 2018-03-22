from construct import *
from construct.lib import *

nested_types__subtype_a__subtype_c = Struct(
	'value_c' / Int8sb,
)

nested_types__subtype_a = Struct(
	'typed_at_root' / nested_types__subtype_b,
	'typed_here' / nested_types__subtype_a__subtype_c,
)

nested_types__subtype_b = Struct(
	'value_b' / Int8sb,
)

nested_types = Struct(
	'one' / nested_types__subtype_a,
	'two' / nested_types__subtype_b,
)

_schema = nested_types
