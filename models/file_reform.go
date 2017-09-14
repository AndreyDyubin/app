package models

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type dataFilesTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *dataFilesTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("datafiles").
func (v *dataFilesTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *dataFilesTableType) Columns() []string {
	return []string{"uuid", "name"}
}

// NewStruct makes a new struct for that view or table.
func (v *dataFilesTableType) NewStruct() reform.Struct {
	return new(DataFiles)
}

// NewRecord makes a new record for that table.
func (v *dataFilesTableType) NewRecord() reform.Record {
	return new(DataFiles)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *dataFilesTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// DataFilesTable represents datafiles view or table in SQL database.
var DataFilesTable = &dataFilesTableType{
	s: parse.StructInfo{Type: "DataFiles", SQLSchema: "", SQLName: "datafiles", Fields: []parse.FieldInfo{{Name: "UUID", PKType: "string", Column: "uuid"}, {Name: "Name", PKType: "", Column: "name"}}, PKFieldIndex: 0},
	z: new(DataFiles).Values(),
}

// String returns a string representation of this struct or record.
func (s DataFiles) String() string {
	res := make([]string, 2)
	res[0] = "UUID: " + reform.Inspect(s.UUID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *DataFiles) Values() []interface{} {
	return []interface{}{
		s.UUID,
		s.Name,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *DataFiles) Pointers() []interface{} {
	return []interface{}{
		&s.UUID,
		&s.Name,
	}
}

// View returns View object for that struct.
func (s *DataFiles) View() reform.View {
	return DataFilesTable
}

// Table returns Table object for that record.
func (s *DataFiles) Table() reform.Table {
	return DataFilesTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *DataFiles) PKValue() interface{} {
	return s.UUID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *DataFiles) PKPointer() interface{} {
	return &s.UUID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *DataFiles) HasPK() bool {
	return s.UUID != DataFilesTable.z[DataFilesTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *DataFiles) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.UUID = string(i64)
	} else {
		s.UUID = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = DataFilesTable
	_ reform.Struct = new(DataFiles)
	_ reform.Table  = DataFilesTable
	_ reform.Record = new(DataFiles)
	_ fmt.Stringer  = new(DataFiles)
)

func init() {
	parse.AssertUpToDate(&DataFilesTable.s, new(DataFiles))
}
