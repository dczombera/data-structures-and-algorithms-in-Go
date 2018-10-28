package separate_chaining_hash_table

import "testing"

type testCase struct {
	data []KV
	want []KV
}

type KV struct {
	key Key
	val Value
}

var testCases = []testCase{
	{[]KV{{"han", 1}, {"luke", 2}, {"r2d2", 3}}, []KV{{"han", 1}, {"luke", 2}, {"r2d2", 3}}},
	{[]KV{{"foo", 42}, {"bar", 0}, {"bar", 42}, {"foo", 0}}, []KV{{"bar", 42}, {"foo", 0}}},
	{[]KV{{"han", -42}, {"shot", 42}, {"first", 0}, {"second", 0}, {"jumped", 42}}, []KV{{"second", 0}, {"han", -42}, {"jumped", 42}}},
}

func containsKeyVal(ht *SeparateChainingHT, k Key, v Value) bool {
	for _, curr := range ht.table {
		for ; curr != nil; curr = curr.Next {
			if curr.Key == k && curr.Val == v {
				return true
			}
		}
	}
	return false
}

func TestSeparateChainingHTPut(t *testing.T) {
	for _, tc := range testCases {
		ht := NewEmptySeparateChainingHT()
		for _, d := range tc.data {
			ht.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			if !containsKeyVal(&ht, w.key, w.val) {
				t.Errorf("Hash table does not contain key %v with value %v", w.key, w.val)
			}
		}
	}
}

func TestSeparateChainingHTGet(t *testing.T) {
	for _, tc := range testCases {
		ht := NewEmptySeparateChainingHT()
		for _, d := range tc.data {
			ht.Put(d.key, d.val)
		}

		for _, w := range tc.want {
			val, err := ht.Get(w.key)
			if err != nil {
				t.Errorf("Got error %v, want key %v with value %v", err, w.key, w.val)
			}
			if w.val != val {
				t.Errorf("Got %v value for key %v, want value %v", val, w.key, w.val)
			}
		}
	}
}

func TestSeparateChainingHTDelete(t *testing.T) {
	testCases := []struct {
		data    []KV
		deleted []Key
		size    int
	}{
		{[]KV{{"han", 1}, {"luke", 2}, {"r2d2", 3}}, []Key{"luke"}, 2},
		{[]KV{{"foo", 42}, {"bar", 0}, {"bar", 42}, {"foo", 0}}, []Key{"bar", "foo"}, 0},
		{[]KV{{"han", -42}, {"shot", 42}, {"first", 0}, {"second", 0}, {"jumped", 42}}, []Key{"han", "jumped", "second"}, 2},
		{[]KV{{"c3po", 1}, {"r2d2", 2}}, []Key{"bb8"}, 2},
	}

	for _, tc := range testCases {
		ht := NewEmptySeparateChainingHT()
		for _, d := range tc.data {
			ht.Put(d.key, d.val)
		}

		for _, toDel := range tc.deleted {
			ht.Delete(toDel)
		}

		for _, deleted := range tc.deleted {
			v, err := ht.Get(deleted)
			if err == nil {
				t.Errorf("Got value %v, want KV pair with key %v to be deleted", v, deleted)
			}
			if ht.Size() != tc.size {
				t.Errorf("Got size of %v, want size of %v for hash table", ht.Size(), tc.size)
			}
		}
	}

}
