package memory

func (m MemoryStorage) UpdateBalance(id string, value int) error {
	_, err := m.GetById(id)
	if err != nil {
		return err
	}
	accountStored := m.storageAccount[id]
	accountStored.balance = value
	m.storageAccount[id] = accountStored

	return nil
}
