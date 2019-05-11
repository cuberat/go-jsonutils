

# jsonutils
`import "github.com/cuberat/go-jsonutils/jsonutils"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
The jsonutils package provides various utilities for working with JSON
in Go.

Installation


	go get github.com/cuberat/go-jsonutils/jsonutils




## <a name="pkg-index">Index</a>
* [func NewCSVKeyedRecordWriter(delimiter []byte, w io.Writer, marshal_type interface{}) libutils.KeyedRecordWriter](#NewCSVKeyedRecordWriter)
* [func NewCSVKeyedRecordWriterWithEncoder(w io.Writer, encoder libutils.KeyedRecordEncoder) libutils.KeyedRecordWriter](#NewCSVKeyedRecordWriterWithEncoder)
* [func NewTabKeyedRecordWriter(w io.Writer, marshal_type interface{}) libutils.KeyedRecordWriter](#NewTabKeyedRecordWriter)
* [type CSVKeyedRecordCodec](#CSVKeyedRecordCodec)
  * [func NewCSVKeyedRecordCodec(delimiter []byte, marshal_type interface{}) *CSVKeyedRecordCodec](#NewCSVKeyedRecordCodec)
  * [func NewTabKeyedRecordCodec(marshal_type interface{}) *CSVKeyedRecordCodec](#NewTabKeyedRecordCodec)
  * [func (tkre *CSVKeyedRecordCodec) CodecSame() bool](#CSVKeyedRecordCodec.CodecSame)
  * [func (tkre *CSVKeyedRecordCodec) JoinKV(key, val []byte) ([]byte, error)](#CSVKeyedRecordCodec.JoinKV)
  * [func (tkre *CSVKeyedRecordCodec) MarshalVal(data interface{}) ([]byte, error)](#CSVKeyedRecordCodec.MarshalVal)
  * [func (tkrd *CSVKeyedRecordCodec) SplitKV(wire_data []byte) ([]byte, []byte, error)](#CSVKeyedRecordCodec.SplitKV)
  * [func (tkrd *CSVKeyedRecordCodec) UnmarshalVal(val_bytes []byte) (interface{}, error)](#CSVKeyedRecordCodec.UnmarshalVal)
* [type CSVKeyedRecordScanner](#CSVKeyedRecordScanner)
  * [func NewCSVKeyedRecordScanner(delimiter []byte, r io.Reader, marshal_type interface{}) *CSVKeyedRecordScanner](#NewCSVKeyedRecordScanner)
  * [func NewCSVKeyedRecordScannerWithDecoder(r io.Reader, decoder libutils.KeyedRecordDecoder) *CSVKeyedRecordScanner](#NewCSVKeyedRecordScannerWithDecoder)
  * [func NewTabKeyedRecordScanner(r io.Reader, marshal_type interface{}) *CSVKeyedRecordScanner](#NewTabKeyedRecordScanner)
  * [func (tkrs *CSVKeyedRecordScanner) Err() error](#CSVKeyedRecordScanner.Err)
  * [func (tkrs *CSVKeyedRecordScanner) Record() *libutils.KeyedRecord](#CSVKeyedRecordScanner.Record)
  * [func (tkrs *CSVKeyedRecordScanner) Scan() bool](#CSVKeyedRecordScanner.Scan)
* [type CSVKeyedRecordWriter](#CSVKeyedRecordWriter)
  * [func (krw *CSVKeyedRecordWriter) Write(rec *libutils.KeyedRecord) (int, error)](#CSVKeyedRecordWriter.Write)


#### <a name="pkg-files">Package files</a>
[jsonutils.go](/src/github.com/cuberat/go-jsonutils/jsonutils/jsonutils.go) 





## <a name="NewCSVKeyedRecordWriter">func</a> [NewCSVKeyedRecordWriter](/src/target/jsonutils.go?s=7190:7308#L204)
``` go
func NewCSVKeyedRecordWriter(delimiter []byte, w io.Writer,
    marshal_type interface{}) libutils.KeyedRecordWriter
```
Returns a `libutils.KeyedRecordWriter` that uses the provided delimiter
between the key and the value.



## <a name="NewCSVKeyedRecordWriterWithEncoder">func</a> [NewCSVKeyedRecordWriterWithEncoder](/src/target/jsonutils.go?s=7510:7632#L212)
``` go
func NewCSVKeyedRecordWriterWithEncoder(w io.Writer,
    encoder libutils.KeyedRecordEncoder) libutils.KeyedRecordWriter
```
Returns a `libutils.KeyedRecordWriter` that uses the provided encoder.



## <a name="NewTabKeyedRecordWriter">func</a> [NewTabKeyedRecordWriter](/src/target/jsonutils.go?s=6913:7009#L198)
``` go
func NewTabKeyedRecordWriter(w io.Writer, marshal_type interface{}) libutils.KeyedRecordWriter
```
Returns a `libutils.KeyedRecordWriter` that assumes tab delimters.




## <a name="CSVKeyedRecordCodec">type</a> [CSVKeyedRecordCodec](/src/target/jsonutils.go?s=1944:2030#L39)
``` go
type CSVKeyedRecordCodec struct {
    // contains filtered or unexported fields
}
```
Implements the KeyedRecordEncoder and KeyedRecordDecoder interfaces specified
by `github.com/cuberat/go-libutils/libutils`.

This codec encodes and decodes delimited keyed records where the value is a
JSON object.







### <a name="NewCSVKeyedRecordCodec">func</a> [NewCSVKeyedRecordCodec](/src/target/jsonutils.go?s=2324:2422#L51)
``` go
func NewCSVKeyedRecordCodec(delimiter []byte,
    marshal_type interface{}) *CSVKeyedRecordCodec
```
Returns a codec for general delimted records where the key and value are
delimted by `delimter`.


### <a name="NewTabKeyedRecordCodec">func</a> [NewTabKeyedRecordCodec](/src/target/jsonutils.go?s=2077:2153#L45)
``` go
func NewTabKeyedRecordCodec(marshal_type interface{}) *CSVKeyedRecordCodec
```
Returns a codec for tab-delimted records.





### <a name="CSVKeyedRecordCodec.CodecSame">func</a> (\*CSVKeyedRecordCodec) [CodecSame](/src/target/jsonutils.go?s=4177:4226#L117)
``` go
func (tkre *CSVKeyedRecordCodec) CodecSame() bool
```
Returns true so that if this codec is used for both encoder and decoder,
unnecessary re-serialization can be avoided.

This allows for lazy encoding. That is, if the raw record bytes that were
read in do not need to change, they can be written back out as-is, instead of
actually re-encoding.




### <a name="CSVKeyedRecordCodec.JoinKV">func</a> (\*CSVKeyedRecordCodec) [JoinKV](/src/target/jsonutils.go?s=3560:3632#L101)
``` go
func (tkre *CSVKeyedRecordCodec) JoinKV(key, val []byte) ([]byte, error)
```
Joins the key and value bytes, returning the serialized record.




### <a name="CSVKeyedRecordCodec.MarshalVal">func</a> (\*CSVKeyedRecordCodec) [MarshalVal](/src/target/jsonutils.go?s=3754:3831#L107)
``` go
func (tkre *CSVKeyedRecordCodec) MarshalVal(data interface{}) ([]byte, error)
```
Serializes the value data structure.




### <a name="CSVKeyedRecordCodec.SplitKV">func</a> (\*CSVKeyedRecordCodec) [SplitKV](/src/target/jsonutils.go?s=2877:2963#L69)
``` go
func (tkrd *CSVKeyedRecordCodec) SplitKV(wire_data []byte) ([]byte, []byte,
    error)
```
Splits the record, returning the key and the serialized value data
structure.




### <a name="CSVKeyedRecordCodec.UnmarshalVal">func</a> (\*CSVKeyedRecordCodec) [UnmarshalVal](/src/target/jsonutils.go?s=3259:3347#L91)
``` go
func (tkrd *CSVKeyedRecordCodec) UnmarshalVal(val_bytes []byte) (interface{},
    error)
```
Deserializes the value.




## <a name="CSVKeyedRecordScanner">type</a> [CSVKeyedRecordScanner](/src/target/jsonutils.go?s=4354:4458#L123)
``` go
type CSVKeyedRecordScanner struct {
    // contains filtered or unexported fields
}
```
Implements the KeyedRecordScanner interface specified by
`github.com/cuberat/go-libutils/libutils`.







### <a name="NewCSVKeyedRecordScanner">func</a> [NewCSVKeyedRecordScanner](/src/target/jsonutils.go?s=5008:5123#L138)
``` go
func NewCSVKeyedRecordScanner(delimiter []byte, r io.Reader,
    marshal_type interface{}) *CSVKeyedRecordScanner
```
Returns a keyed record scanner that implements the
`libutils.KeyedRecordScanner` interface and assumes the provided delimiter
separates the key from the value in each record.


### <a name="NewCSVKeyedRecordScannerWithDecoder">func</a> [NewCSVKeyedRecordScannerWithDecoder](/src/target/jsonutils.go?s=5361:5480#L149)
``` go
func NewCSVKeyedRecordScannerWithDecoder(r io.Reader,
    decoder libutils.KeyedRecordDecoder) *CSVKeyedRecordScanner
```
Similar to `NewCSVKeyedRecordScanner()`, but allows you to specify a decoder
to use.


### <a name="NewTabKeyedRecordScanner">func</a> [NewTabKeyedRecordScanner](/src/target/jsonutils.go?s=4658:4751#L131)
``` go
func NewTabKeyedRecordScanner(r io.Reader, marshal_type interface{}) *CSVKeyedRecordScanner
```
Returns a keyed record scanner that uses tabs as a delimiter. This is the
same as created a codec with `NewTabKeyedRecordCodec()` and passing it to
`NewCSVKeyedRecordScannerWithDecoder()`.





### <a name="CSVKeyedRecordScanner.Err">func</a> (\*CSVKeyedRecordScanner) [Err](/src/target/jsonutils.go?s=6277:6323#L177)
``` go
func (tkrs *CSVKeyedRecordScanner) Err() error
```
Returns the first non-EOF error that was encountered by the Scanner.




### <a name="CSVKeyedRecordScanner.Record">func</a> (\*CSVKeyedRecordScanner) [Record](/src/target/jsonutils.go?s=5932:5999#L168)
``` go
func (tkrs *CSVKeyedRecordScanner) Record() *libutils.KeyedRecord
```
Returns the most recent serialized record generated by a call to Scan().




### <a name="CSVKeyedRecordScanner.Scan">func</a> (\*CSVKeyedRecordScanner) [Scan](/src/target/jsonutils.go?s=5749:5795#L163)
``` go
func (tkrs *CSVKeyedRecordScanner) Scan() bool
```
Advances the scanner to the next record. It returns false when the scan
stops, either by reaching the end of the input or an error.




## <a name="CSVKeyedRecordWriter">type</a> [CSVKeyedRecordWriter](/src/target/jsonutils.go?s=6467:6593#L183)
``` go
type CSVKeyedRecordWriter struct {
    // contains filtered or unexported fields
}
```
Implements the `libutils.KeyedRecordWriter` interface from
`github.com/cuberat/go-libutils/libutils`.










### <a name="CSVKeyedRecordWriter.Write">func</a> (\*CSVKeyedRecordWriter) [Write](/src/target/jsonutils.go?s=6595:6673#L189)
``` go
func (krw *CSVKeyedRecordWriter) Write(rec *libutils.KeyedRecord) (int, error)
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
