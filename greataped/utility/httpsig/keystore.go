package httpsig

type MemoryKeyStore struct {
	keys map[string]interface{}
}

func NewMemoryKeyStore() *MemoryKeyStore {
	return &MemoryKeyStore{
		keys: make(map[string]interface{}),
	}
}

func (m *MemoryKeyStore) GetKey(id string) interface{} {
	return m.keys[id]
}

func (m *MemoryKeyStore) SetKey(id string, key interface{}) {
	m.keys[id] = key
}
