from construct import *
from construct.lib import *

imports_abs_abs = Struct(
	'one' / Int8ub,
	'two' / imported_and_abs,
)

_schema = imports_abs_abs
