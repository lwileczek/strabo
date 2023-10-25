import { ObjectValues } from './generic'

type BOOLEAN = ObjectValues<typeof BOOLEANS>
type STRING = ObjectValues<typeof STRINGS>
type NUMBER = ObjectValues<typeof NUMBERS>
type DATE = ObjectValues<typeof DATES>
type STRUCTURE = ObjectValues<typeof STRUCTURES>

export type Primative = BOOLEAN | STRING | NUMBER | DATE | STRUCTURE;

const BOOLEANS = {
  bool: "bool",
} as const

const STRINGS = {
	string: "string",
	uuid: "uuid",
	uri: "uri",
	email: "email",
	byte: "byte", //base64
} as const

const DATES = {
	date: "date",
	datetime: "datetime",
	datetimez: "datetimez",
  timetamp: "timetamp",
  timetampz: "timetampz",
} as const

const STRUCTURES= {
	array: "array",
	object: "object",
	hashmap: "hashmap",
} as const

const NUMBERS = {
  	number: "number",
	  int: "int",
	  i8: "i8",
	  i16: "i16",
	  i32: "i32",
	  i64: "i64",
	  uint: "uint",
	  u8: "u8",
	  u16: "u16",
	  u32: "u32",
	  u64: "u64",
	  f32: "f32",
	  f64: "f64",
} as const
