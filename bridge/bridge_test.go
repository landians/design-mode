package bridge

import (
	"bytes"
	"testing"
)

func Test_Bridge(t *testing.T) {
	config := newDBConfig("mysql", "root:pass@tcp(localhost:3306)/test?charset=utf8", "root", "pass")
	fetcher := newMysqlDataFetcher(config)

	fnTestExporter := func(exporter IDataExporter) {
		var writer bytes.Buffer
		if err := exporter.Export("SELECT * FROM `user`", &writer); err != nil {
			t.Error(err)
		}
	}

	fnTestExporter(newCsvExporter(fetcher))
	fnTestExporter(newJsonExporter(fetcher))
}
