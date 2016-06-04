package plus

func chunkKeys(ks keys, numParts int64) []keys {
	parts := make([]keys, numParts)
	for i := int64(0); i < numParts; i++ {
		parts[i] = ks[i*int64(len(ks))/numParts : (i+1)*int64(len(ks))/numParts]
	}
	return parts
}

type mockKey struct {
	value int
}

func (mk *mockKey) Compare(other Key) int {
	key := other.(*mockKey)
	if key.value == mk.value {
		return 0
	}
	if key.value > mk.value {
		return 1
	}

	return -1
}

func newMockKey(value int) *mockKey {
	return &mockKey{value}
}
