package storageservice

// StorageServer server interface.
type StorageServer struct {
	Store Store
}

// NewStorageService initialize ne
func NewStorageService(store Store) *StorageServer {
	return &StorageServer{Store: store}
}
