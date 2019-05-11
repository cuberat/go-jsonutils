// BSD 2-Clause License
//
// Copyright (c) 2019 Don Owens <don@regexguy.com>.  All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// The jsonutils package provides various utilities for working with JSON
// in Go.
//
// Installation
//
//   go get github.com/cuberat/go-jsonutils/jsonutils
package jsonutils

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/cuberat/go-libutils/libutils"
    "io"
    "reflect"
)

// Implements the KeyedRecordEncoder and KeyedRecordDecoder interfaces specified
// by `github.com/cuberat/go-libutils/libutils`.
//
// This codec encodes and decodes delimited keyed records where the value is a
// JSON object.
type CSVKeyedRecordCodec struct {
    marshal_type reflect.Type
    delimiter []byte
}

// Returns a codec for tab-delimted records.
func NewTabKeyedRecordCodec(marshal_type interface{}) (*CSVKeyedRecordCodec) {
    return NewCSVKeyedRecordCodec([]byte{'\t'}, marshal_type)
}

// Returns a codec for general delimted records where the key and value are
// delimted by `delimter`.
func NewCSVKeyedRecordCodec(delimiter []byte,
    marshal_type interface{}) (*CSVKeyedRecordCodec) {
    tkrd := new(CSVKeyedRecordCodec)
    tkrd.delimiter = delimiter
    value := reflect.ValueOf(marshal_type)
    kind := value.Kind()
    for kind == reflect.Ptr || kind == reflect.Interface {
        value = value.Elem()
        kind = value.Kind()
    }
    marshal_type = value.Interface()
    tkrd.marshal_type = reflect.TypeOf(marshal_type)

    return tkrd
}

func (tkrd *CSVKeyedRecordCodec) SplitKV(wire_data []byte) ([]byte, []byte,
    error) {
    data := bytes.SplitN(wire_data, tkrd.delimiter, 2)

    var (
        key []byte
        val []byte
    )

    key = data[0]
    if len(data) > 1 {
        val = data[1]
    }

    if len(val) == 0 {
        val = []byte("{}")
    }

    return key, val, nil
}

func (tkrd *CSVKeyedRecordCodec) UnmarshalVal(val_bytes []byte) (interface{},
    error) {
    data_value := reflect.New(tkrd.marshal_type).Interface()

    err := json.Unmarshal(val_bytes, &data_value)

    return data_value, err
}

func (tkre *CSVKeyedRecordCodec) JoinKV(key, val []byte) ([]byte, error) {
    kv := bytes.Join([][]byte{key, val}, tkre.delimiter)
    return kv, nil
}

func (tkre *CSVKeyedRecordCodec) MarshalVal(data interface{}) ([]byte, error) {
    return json.Marshal(data)
}

// Implements the KeyedRecordScanner interface specified by
// `github.com/cuberat/go-libutils/libutils`.
type CSVKeyedRecordScanner struct {
    scanner *bufio.Scanner
    decoder libutils.KeyedRecordDecoder
}

// Returns a keyed record scanner that uses tabs as a delimiter. This is the
// same as created a codec with `NewTabKeyedRecordCodec()` and passing it to
// `NewCSVKeyedRecordScannerWithDecoder()`.
func NewTabKeyedRecordScanner(r io.Reader, marshal_type interface{}) (*CSVKeyedRecordScanner) {
    return NewCSVKeyedRecordScanner([]byte{'\t'}, r, marshal_type)
}

func NewCSVKeyedRecordScanner(delimiter []byte, r io.Reader,
    marshal_type interface{}) (*CSVKeyedRecordScanner) {

    decoder := NewCSVKeyedRecordCodec(delimiter, marshal_type)
    tkrs := NewCSVKeyedRecordScannerWithDecoder(r, decoder)

    return tkrs
}

func NewCSVKeyedRecordScannerWithDecoder(r io.Reader,
    decoder libutils.KeyedRecordDecoder) (*CSVKeyedRecordScanner) {

    tkrs := new(CSVKeyedRecordScanner)

    tkrs.scanner = bufio.NewScanner(r)

    tkrs.decoder = decoder

    return tkrs
}

func (tkrs *CSVKeyedRecordScanner) Scan() bool {
    return tkrs.scanner.Scan() // Default line scanner
}

func (tkrs *CSVKeyedRecordScanner) Record() (*libutils.KeyedRecord) {
    wire_data := tkrs.scanner.Bytes()
    wire_data_copy := make([]byte, len(wire_data))
    copy(wire_data_copy, wire_data)

    return libutils.NewKeyedRecordFromBytes(wire_data_copy, tkrs.decoder)
}

func (tkrs *CSVKeyedRecordScanner) Err() error {
    return tkrs.scanner.Err()
}

type CSVKeyedRecordWriter struct {
    marshal_type interface{}
    encoder libutils.KeyedRecordEncoder
    writer io.Writer
}

func (krw *CSVKeyedRecordWriter) Write(rec *libutils.KeyedRecord) (int, error) {
    rec_out_bytes, err := rec.RecordBytesOut(krw.encoder)
    if err != nil {
        return 0, err
    }
    return fmt.Fprintf(krw.writer, "%s\n", rec_out_bytes)
}

func NewTabKeyedRecordWriter(w io.Writer, marshal_type interface{}) (libutils.KeyedRecordWriter) {
    return NewCSVKeyedRecordWriter([]byte{'\t'}, w, marshal_type)
}

func NewCSVKeyedRecordWriter(delimiter []byte, w io.Writer,
    marshal_type interface{}) (libutils.KeyedRecordWriter) {
    encoder := NewCSVKeyedRecordCodec(delimiter, marshal_type)

    return NewCSVKeyedRecordWriterWithEncoder(w, encoder)
}

func NewCSVKeyedRecordWriterWithEncoder(w io.Writer,
    encoder libutils.KeyedRecordEncoder) (libutils.KeyedRecordWriter) {

    writer := new(CSVKeyedRecordWriter)
    writer.writer = w
    writer.encoder = encoder

    return writer
}
