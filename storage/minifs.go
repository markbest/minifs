package storage

import "github.com/syndtr/goleveldb/leveldb"

const (
	defaultBasePath = "./data"
	propertyPrefix  = "leveldb."
)

type MiniFs struct {
	Options
}

type Options struct {
	BasePath string
}

// New returns an initialized MiniDB structure, ready to use.
func NewMiniFs(o Options) *MiniFs {
	if o.BasePath == "" {
		o.BasePath = defaultBasePath
	}

	m := &MiniFs{
		Options: o,
	}
	return m
}

// writes the key-value pair to disk, making it immediately
func (m *MiniFs) Write(key string, value string) error {
	//open leveldb
	db, err := leveldb.OpenFile(m.BasePath, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	//writes content to leveldb
	err = db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		return err
	}
	return nil
}

// GetProperty returns value of the given property name.
//
// Property names:
//	num-files-at-level{n}
//		Returns the number of files at level 'n'.
//	stats
//		Returns statistics of the underlying DB.
//	sstables
//		Returns sstables list for each level.
//	blockpool
//		Returns block pool stats.
//	cachedblock
//		Returns size of cached block.
//	openedtables
//		Returns number of opened tables.
//	alivesnaps
//		Returns number of alive snapshots.
//	aliveiters
//		Returns number of alive iterators.
func (m *MiniFs) GetProperty(name string) (string, error) {
	//open leveldb
	db, err := leveldb.OpenFile(m.BasePath, nil)
	if err != nil {
		return "", err
	}
	defer db.Close()

	//get property
	value, err := db.GetProperty(propertyPrefix + name)
	if err != nil {
		return "", err
	}
	return value, nil
}
