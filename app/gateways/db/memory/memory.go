package memory

type MemoryStorage struct {
	storageAccount        map[string]Account
	storageTransfer       map[string][]Transfer
	storageAuthentication map[string]Token
}

func NewMemoryStorage(storageAccount map[string]Account, storageTransfer map[string][]Transfer, storageAuthentication map[string]Token) MemoryStorage {
	return MemoryStorage{
		storageAccount:        storageAccount,
		storageTransfer:       storageTransfer,
		storageAuthentication: storageAuthentication,
	}
}
