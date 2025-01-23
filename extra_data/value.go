package extra_data

type ValueType int

var BoolType ValueType = 0
var StringType ValueType = 1
var U8Type ValueType = 2
var U16Type ValueType = 3
var U32Type ValueType = 4
var U64Type ValueType = 5
var U128Type ValueType = 6
var HashType ValueType = 7
var BlobType ValueType = 8

type Hash [32]byte
type Blob []uint8

type Value interface{}
