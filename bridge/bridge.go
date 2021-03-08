package bridge

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// 数据库连接配置信息
type DBConfig struct {
	DbType string
	URL    string
	UID    string
	PWD    string
}

func newDBConfig(dbType string, url string, uid string, pwd string) *DBConfig {
	return &DBConfig{
		DbType: dbType,
		URL:    url,
		UID:    uid,
		PWD:    pwd,
	}
}

type DataTypes string

const DATA_TYPE_INT = "int"
const DATA_TYPE_FLOAT = "float"
const DATA_TYPE_STRING = "string"
const DATA_TYPE_BOOL = "bool"
const DATA_TYPE_DATETIME = "datetime"

// 导出数据行的某个字段
type DataField struct {
	Name     string
	DataType DataTypes

	IntValue      int
	FloatValue    float64
	StringValue   string
	BoolValue     bool
	DateTimeValue *time.Time
}

func newMockDataField(name string, dataType DataTypes) *DataField {
	it := &DataField{
		Name:     name,
		DataType: dataType,

		IntValue:      0,
		FloatValue:    0,
		StringValue:   "",
		BoolValue:     false,
		DateTimeValue: nil,
	}

	switch dataType {
	case DATA_TYPE_INT:
		it.IntValue = 1
		break

	case DATA_TYPE_FLOAT:
		it.FloatValue = 1.1
		break

	case DATA_TYPE_STRING:
		it.StringValue = "hello"
		break

	case DATA_TYPE_DATETIME:
		t := time.Now()
		it.DateTimeValue = &t
		break

	case DATA_TYPE_BOOL:
		it.BoolValue = false
		break
	}

	return it
}

func (me *DataField) ValueString() string {
	switch me.DataType {
	case DATA_TYPE_INT:
		return fmt.Sprintf("%v", me.IntValue)

	case DATA_TYPE_FLOAT:
		return fmt.Sprintf("%v", me.FloatValue)

	case DATA_TYPE_STRING:
		return fmt.Sprintf("\"%s\"", me.StringValue)

	case DATA_TYPE_DATETIME:
		return fmt.Sprintf("\"%s\"", me.DateTimeValue.Format("2006-01-02T15:04:05"))

	case DATA_TYPE_BOOL:
		return fmt.Sprintf("%v", me.BoolValue)
	}

	return ""
}

// 导出数据行的中间结果
type DataRow struct {
	FieldList []*DataField
}

func newMockDataRow() *DataRow {
	it := &DataRow{FieldList: make([]*DataField, 0)}

	it.FieldList = append(it.FieldList, newMockDataField("int-1", DATA_TYPE_INT))
	it.FieldList = append(it.FieldList, newMockDataField("float-1", DATA_TYPE_FLOAT))
	it.FieldList = append(it.FieldList, newMockDataField("string-1", DATA_TYPE_STRING))
	return it
}

func (d *DataRow) FieldsString() string {
	list := make([]string, 0, len(d.FieldList))
	for _, f := range d.FieldList {
		list = append(list, fmt.Sprintf("%s=%s", f.Name, f.ValueString()))
	}
	return strings.Join(list, ",")
}

// 数据获取接口，执行 SQL 语句，并转换为数据行的集合
type IDataFetcher interface {
	Fetch(sql string) []*DataRow
}

// Mysql 数据获取，实现 IDataFetcher 接口
type mysqlDataFetcher struct {
	Config *DBConfig
}

func newMysqlDataFetcher(config *DBConfig) IDataFetcher {
	return &mysqlDataFetcher{Config: config}
}

func (m *mysqlDataFetcher) Fetch(sql string) []*DataRow {
	rows := make([]*DataRow, 0)
	rows = append(rows, newMockDataRow())
	return rows
}

// Oracle 数据获取接口，实现 IDataFetcher 接口
type oracleDataFetcher struct {
	Config *DBConfig
}

func newOracleDataFetcher(config *DBConfig) IDataFetcher {
	return &oracleDataFetcher{Config: config}
}

func (o *oracleDataFetcher) Fetch(sql string) []*DataRow {
	rows := make([]*DataRow, 0)
	rows = append(rows, newMockDataRow())
	return rows
}

// 数据导出接口
type IDataExporter interface {
	Fetcher(fetcher IDataFetcher)
	Export(sql string, writer io.Writer) error
}

// csv 格式的文件导出
type csvExporter struct {
	fetcher IDataFetcher
}

func newCsvExporter(fetcher IDataFetcher) IDataExporter {
	return &csvExporter{
		fetcher: fetcher,
	}
}

func (c *csvExporter) Fetcher(fetcher IDataFetcher) {
	c.fetcher = fetcher
}

func (c *csvExporter) Export(sql string, writer io.Writer) error {
	rows := c.fetcher.Fetch(sql)
	fmt.Printf("CsvFetcher.Export, got %v rows\n", len(rows))
	for i, it := range rows {
		fmt.Printf(" %v %s\n", i+1, it.FieldsString())
	}
	return nil
}

// Json 格式文件导出
type jsonExporter struct {
	fetcher IDataFetcher
}

func newJsonExporter(fetcher IDataFetcher) IDataExporter {
	return &csvExporter{
		fetcher: fetcher,
	}
}

func (j *jsonExporter) Fetcher(fetcher IDataFetcher) {
	j.fetcher = fetcher
}

func (j *jsonExporter) Export(sql string, writer io.Writer) error {
	rows := j.fetcher.Fetch(sql)
	fmt.Printf("JsonFetcher.Export, got %v rows\n", len(rows))
	for i, it := range rows {
		fmt.Printf(" %v %s\n", i+1, it.FieldsString())
	}
	return nil
}
