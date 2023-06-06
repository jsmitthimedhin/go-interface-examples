package storage

type Storage interface {
	ListValues(prefix string) ([]byte, error)
	GetValue(path string) (byte, error)
	PutValue(path string, value []byte) error
	DeleteValue(path string) error
}

type database struct{}

func (d *database) ListValues(prefix string) ([]byte, error) {
	// unique code for listing values from database here
	return nil, nil
}

func (d *database) GetValue(path string) (byte, error) {
	return 0, nil
}

func (d *database) PutValue(path string, value []byte) error {
	return nil
}

func (d *database) DeleteValue(path string) error {
	return nil
}

func saveToStorage(str Storage, path string, values []byte) error {
	// code for saving things to storage
	return nil
}

func testFunction() {
	db := &database{}
	values := make([]byte, 0)
	saveToStorage(db, "path", values)
}
