package bimap

// Bimap is a bidirectional map structure that ensures constant time complexity.
// It maintains two maps: one for forward mapping and another for reverse mapping.
type Bimap[T1 comparable, T2 comparable] struct {
	forward map[T1]T2
	reverse map[T2]T1
}

// Pair represents a key-value pair used in the Bimap.
type Pair[T1 comparable, T2 comparable] struct {
	Key   T1
	Value T2
}

// New creates and initializes a Bimap with specified types T1 and T2.
// It returns a pointer to the newly created Bimap.
func New[T1 comparable, T2 comparable]() *Bimap[T1, T2] {
	return &Bimap[T1, T2]{
		forward: make(map[T1]T2),
		reverse: make(map[T2]T1),
	}
}

// Set adds or updates a key-value pair in the Bimap.
// If the key or value already exists, the old pair is removed to maintain consistency.
func (b *Bimap[T1, T2]) Set(key T1, value T2) {
	if oldVal, ok := b.forward[key]; ok {
		delete(b.reverse, oldVal)
	}
	if oldKey, ok := b.reverse[value]; ok {
		delete(b.forward, oldKey)
	}

	b.forward[key] = value
	b.reverse[value] = key
}

// GetByKey retrieves the value associated with the given key.
// Returns the value and a boolean indicating whether the key exists.
func (b *Bimap[T1, T2]) GetByKey(key T1) (T2, bool) {
	val, ok := b.forward[key]
	return val, ok
}

// GetByValue retrieves the key associated with the given value.
// Returns the key and a boolean indicating whether the value exists.
func (b *Bimap[T1, T2]) GetByValue(value T2) (T1, bool) {
	key, ok := b.reverse[value]
	return key, ok
}

// DeleteByKey removes the pair associated with the given key.
// Returns true if the pair was removed, and false if the key did not exist.
func (b *Bimap[T1, T2]) DeleteByKey(key T1) bool {
	if val, ok := b.forward[key]; ok {
		delete(b.forward, key)
		delete(b.reverse, val)

		return true
	} else {
		return false
	}
}

// DeleteByValue removes the pair associated with the given value.
// Returns true if the pair was removed, and false if the value did not exist.
func (b *Bimap[T1, T2]) DeleteByValue(value T2) bool {
	if key, ok := b.reverse[value]; ok {
		delete(b.reverse, value)
		delete(b.forward, key)

		return true
	} else {
		return false
	}
}

// ToList returns the Bimap as a slice of simple Key, Value Pair.
func (b Bimap[T1, T2]) ToList() []Pair[T1, T2] {
	result := make([]Pair[T1, T2], len(b.forward))
	i := 0

	for k, v := range b.forward {
		result[i] = Pair[T1, T2]{
			Key:   k,
			Value: v,
		}

		i++
	}

	return result
}
