package memory

type MemoryStorage struct {
	storageAccount  map[string]Account
	storageTransfer map[string]Transfer
}

func NewMemoryStorage(storageAccount map[string]Account, storageTransfer map[string]Transfer) MemoryStorage {
	return MemoryStorage{
		storageAccount:  storageAccount,
		storageTransfer: storageTransfer,
	}
}
