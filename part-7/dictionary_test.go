package dictionary

import "testing"

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertKV(t *testing.T, dict Dictionary, key, value string) {
	t.Helper()

	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if value != got {
		t.Errorf("got error '%s' want '%s'", got, value)
	}
}

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknow")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	key := "test"
	value := "this is just a test"
	t.Run("new word", func(t *testing.T) {
		err := dict.Add(key, value)
		assertError(t, err, nil)
		assertKV(t, dict, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		err := dict.Add(key, "new test")
		assertError(t, err, ErrWordExists)
		assertKV(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		old := "this is just a test"
		new := "new value"
		dict := Dictionary{key: old}

		err := dict.Update(key, new)
		assertError(t, err, nil)
		assertKV(t, dict, key, new)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(key, value)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dict := Dictionary{key: "test value"}

	dict.Delete(key)

	_, err := dict.Search(key)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", key)
	}
}
