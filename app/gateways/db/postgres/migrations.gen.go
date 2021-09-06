package postgres

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_create_accounts_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x4c\x4e\xce\x2f\xcd\x2b\x29\xb6\xe6\xe2\x02\x04\x00\x00\xff\xff\x97\xf8\xa4\x75\x20\x00\x00\x00")

func _000001_create_accounts_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_create_accounts_table_down_sql,
		"000001_create_accounts_table.down.sql",
	)
}

var __000001_create_accounts_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcc\xc1\x8a\x83\x30\x14\x46\xe1\xbd\x4f\xf1\x93\x95\xc2\xbc\xc1\xac\x32\x33\x11\x64\x32\x4e\xd1\x08\x75\x55\xae\x31\x6d\x03\x35\x4a\xbc\xb6\xd0\xa7\x2f\x75\xd1\x45\xc1\xf5\x77\x38\xdf\x95\x92\x46\xc1\xc8\x2f\xad\x50\xe4\x28\xff\x0d\xd4\xbe\xa8\x4d\x0d\x41\xd6\x8e\x4b\xe0\x59\x20\x4d\x00\xe1\x7b\x81\x65\xf1\x3d\x76\x55\xf1\x27\xab\x16\xbf\xaa\xfd\x78\x42\xa0\xc1\x09\x5c\x29\xda\x33\xc5\xf5\x50\x36\x5a\xaf\x64\xa7\xe3\x86\xcc\xce\x46\xc7\x1b\xd8\xd1\x85\x82\x75\x02\x9d\x3f\xf9\xc0\x6f\xcf\xe8\x88\x5d\x7f\x20\x16\x60\x3f\xb8\x99\x69\x98\xf8\xfe\x8a\xf0\xa3\x72\xd9\x68\x83\x34\x8c\xb7\x34\xcb\x92\xec\xf3\x11\x00\x00\xff\xff\x47\x10\x60\xf5\xe6\x00\x00\x00")

func _000001_create_accounts_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_create_accounts_table_up_sql,
		"000001_create_accounts_table.up.sql",
	)
}

var __000002_create_transfers_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x29\x4a\xcc\x2b\x4e\x4b\x2d\x2a\xb6\x06\x04\x00\x00\xff\xff\x9f\x5c\x6f\x39\x1f\x00\x00\x00")

func _000002_create_transfers_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000002_create_transfers_table_down_sql,
		"000002_create_transfers_table.down.sql",
	)
}

var __000002_create_transfers_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xce\xc1\x4a\xc4\x30\x10\x06\xe0\x7b\x9f\xe2\x67\x4e\x2d\xf8\x06\x9e\xea\x3a\x85\x62\xac\x92\x66\xc1\x3d\x95\xd8\x44\xc9\xa1\x89\x24\x53\x04\x9f\x5e\x82\x65\x7d\x80\xbd\xce\xff\xcf\x37\x73\xd2\xdc\x1b\x86\xe9\x1f\x14\x63\x1c\x30\xbd\x18\xf0\xdb\x38\x9b\x19\x24\xd9\xc6\xf2\xe1\x73\x21\xb4\x0d\x00\x50\x70\x84\x7d\x0f\x0e\xaf\x7a\x7c\xee\xf5\x05\x4f\x7c\xb9\xfb\x8b\x52\x0e\x9f\x21\x2e\x76\x5d\xd3\x1e\x65\xb9\x36\x2b\x38\x9d\x95\x82\xe6\x81\x35\x4f\x27\x9e\x41\x47\xab\x50\x5b\xc9\xee\x20\x9c\x2f\x12\xa2\x95\x90\x6e\x73\xec\x56\x87\x84\xf7\xfa\x91\x5c\x37\x8f\x74\xcd\xde\x8a\x77\x8b\x15\x82\x84\xcd\x17\xb1\xdb\x97\xfc\xfc\x1f\x78\xe4\xa1\x3f\x2b\x83\x36\xa6\xef\xb6\xeb\x9a\xee\xbe\xf9\x0d\x00\x00\xff\xff\x4a\xdb\x6d\xdb\x26\x01\x00\x00")

func _000002_create_transfers_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000002_create_transfers_table_up_sql,
		"000002_create_transfers_table.up.sql",
	)
}

var __000003_create_indexes_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x50\xca\x4c\x51\xb2\xe6\xc2\x2e\x95\x5c\x90\x86\x53\x2e\xbf\x28\x33\x3d\x33\x2f\x3e\x31\x39\x39\xbf\x34\xaf\x24\x1e\x6c\x0a\x20\x00\x00\xff\xff\xe3\x43\x7e\xfe\x61\x00\x00\x00")

func _000003_create_indexes_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000003_create_indexes_table_down_sql,
		"000003_create_indexes_table.down.sql",
	)
}

var __000003_create_indexes_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x0e\x72\x75\x0c\x71\x55\xf0\xf4\x73\x71\x8d\x50\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\x50\xca\x4c\x51\x52\xf0\xf7\x53\x50\x4a\x4c\x4e\xce\x2f\xcd\x2b\x29\x56\xd2\x00\x09\x69\x2a\x58\x73\xe1\xd3\x95\x5c\x90\x86\xa6\x4d\x41\x03\x2c\xa8\x89\x5f\x5f\x7e\x51\x66\x7a\x66\x5e\x3c\x54\x57\x3c\xcc\xf2\x92\xa2\xc4\xbc\xe2\xb4\xd4\x22\xb0\x31\x98\x6a\x34\xad\x01\x01\x00\x00\xff\xff\xdd\xed\x9a\x1f\xc2\x00\x00\x00")

func _000003_create_indexes_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000003_create_indexes_table_up_sql,
		"000003_create_indexes_table.up.sql",
	)
}

var __000004_create_tokens_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\xc9\xcf\x4e\xcd\x2b\xb6\x06\x04\x00\x00\xff\xff\x2b\xef\x20\x7a\x1c\x00\x00\x00")

func _000004_create_tokens_table_down_sql() ([]byte, error) {
	return bindata_read(
		__000004_create_tokens_table_down_sql,
		"000004_create_tokens_table.down.sql",
	)
}

var __000004_create_tokens_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\xbd\x0a\xc2\x30\x18\x85\xe1\xbd\x57\x71\xc8\xd4\x82\x77\xe0\x14\x35\x85\x62\xac\xd2\xa6\x60\x27\x89\x6d\xc0\x28\xfd\x21\xf9\xa2\xe2\xd5\x0b\x5a\xdc\xec\xfc\x9c\xf7\xac\x0b\xc1\x95\x80\xe2\x2b\x29\x90\xa5\xc8\xf7\x0a\xe2\x98\x95\xaa\x04\xa3\xe1\x66\x7a\xcf\x10\x47\x00\xc0\x6c\xcb\x10\x82\x6d\x71\x28\xb2\x1d\x2f\x6a\x6c\x45\xbd\xf8\x52\xaf\x3b\xc3\x70\xd7\xae\xb9\x68\xf7\xf9\xc8\x2b\x29\x27\xf4\xe1\x7c\x35\x0d\xfd\x75\xeb\x7d\x30\x6e\x9e\xdb\x93\x26\x06\xb2\x9d\xf1\xa4\xbb\x91\x5e\xbf\x15\xb0\x11\x29\xaf\xa4\x42\xdc\x0f\x8f\x38\x49\xa6\xcc\x3c\x47\xeb\x66\xba\x28\x59\xbe\x03\x00\x00\xff\xff\xdb\x45\x73\xe2\xfc\x00\x00\x00")

func _000004_create_tokens_table_up_sql() ([]byte, error) {
	return bindata_read(
		__000004_create_tokens_table_up_sql,
		"000004_create_tokens_table.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"000001_create_accounts_table.down.sql":  _000001_create_accounts_table_down_sql,
	"000001_create_accounts_table.up.sql":    _000001_create_accounts_table_up_sql,
	"000002_create_transfers_table.down.sql": _000002_create_transfers_table_down_sql,
	"000002_create_transfers_table.up.sql":   _000002_create_transfers_table_up_sql,
	"000003_create_indexes_table.down.sql":   _000003_create_indexes_table_down_sql,
	"000003_create_indexes_table.up.sql":     _000003_create_indexes_table_up_sql,
	"000004_create_tokens_table.down.sql":    _000004_create_tokens_table_down_sql,
	"000004_create_tokens_table.up.sql":      _000004_create_tokens_table_up_sql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_create_accounts_table.down.sql":  &_bintree_t{_000001_create_accounts_table_down_sql, map[string]*_bintree_t{}},
	"000001_create_accounts_table.up.sql":    &_bintree_t{_000001_create_accounts_table_up_sql, map[string]*_bintree_t{}},
	"000002_create_transfers_table.down.sql": &_bintree_t{_000002_create_transfers_table_down_sql, map[string]*_bintree_t{}},
	"000002_create_transfers_table.up.sql":   &_bintree_t{_000002_create_transfers_table_up_sql, map[string]*_bintree_t{}},
	"000003_create_indexes_table.down.sql":   &_bintree_t{_000003_create_indexes_table_down_sql, map[string]*_bintree_t{}},
	"000003_create_indexes_table.up.sql":     &_bintree_t{_000003_create_indexes_table_up_sql, map[string]*_bintree_t{}},
	"000004_create_tokens_table.down.sql":    &_bintree_t{_000004_create_tokens_table_down_sql, map[string]*_bintree_t{}},
	"000004_create_tokens_table.up.sql":      &_bintree_t{_000004_create_tokens_table_up_sql, map[string]*_bintree_t{}},
}}
