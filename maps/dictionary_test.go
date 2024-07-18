package main 

import ( 
    "testing"
)

func TestSearch (t *testing.T) {
    
    dictionary := Dictionary{"test": "this is just a test"}

    t.Run("known word", func(t *testing.T) {
        got,_ := dictionary.Search("test")
        expected := "this is just a test"

        assertString(t, got, expected)
    })
    t.Run("unknown word", func(t *testing.T) {
        _,err := dictionary.Search("lasagna")

        assertError(t, err, ErrNotFound)
    })
}

func TestAdd (t *testing.T) {
    t.Run("new word", func(t *testing.T) {
        dictionary := Dictionary{}
        word := "test"
        definition := "this is just a test"
        err := dictionary.Add(word, definition)
       
        assertError(t, err, nil)
        assertDefinition(t, dictionary, word, definition)
    })
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        definition := "this is just a test"
        dict := Dictionary{word: definition}
        err := dict.Add(word, "another definition")

        assertError(t, err, ErrWordExists)
        assertDefinition(t, dict, word, definition)
    })
}

func TestUpdate (t *testing.T) {
    t.Run("existing word", func(t *testing.T) {
        word := "test"
        def := "this is just a test"
        dict := Dictionary{word: def}
        new_def := "new definition"
        err := dict.Update(word, new_def)

        assertError(t, err, nil)
        assertDefinition(t, dict, word, new_def)
    })
    t.Run("new word", func(t *testing.T) {
        word := "test"
        def := "this is just a test"
        dict := Dictionary{}
        err := dict.Update(word, def)

        assertError(t, err, ErrWordDoesNotExist)
    })
}

func TestDelete (t *testing.T) {
    word := "test"
    def := "this is just a test"
    dict := Dictionary{word: def}
    dict.Delete(word)

    _, err := dict.Search(word)
    if err != ErrNotFound {
        t.Errorf("Expected %q to be deleted", word)
    }
}

func assertString (t testing.TB, got, expected string) {
    t.Helper()

    if got != expected {
        t.Errorf("got %q but expected %q", got, expected)
    }
}

func assertError (t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got error %q, but expected %q", got, want)
    }
}

func assertDefinition (t testing.TB, dict Dictionary, word, def string) {
    t.Helper()
    got,err := dict.Search(word)

    if err != nil {
        t.Fatal("word was unsuccessfully added")
    }
    assertString(t, got, def)
}
