package main

type Dictionary map[string]string

const (
    ErrNotFound = DictionaryErr("Cannot find word because it does not exist in our dictionary")
    ErrWordExists = DictionaryErr("Cannot add word because it already exists in the dictionary")
    ErrWordDoesNotExist = DictionaryErr("Cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

func (dict Dictionary) Search (word string) (string, error) {
    def, found := dict[word]
    if !found {
        return "", ErrNotFound 
    }
    return def, nil 
}

func (dict Dictionary) Add (word, def string) error {
    _, err := dict.Search(word)
    
    switch err{
    case ErrNotFound:
        dict[word] = def
    case nil:
        return ErrWordExists
    default:
        return err
    }

    return nil
}

func (dict Dictionary) Update (word, def string) error {
    _, err := dict.Search(word)

    switch err {
    case ErrNotFound:
        return ErrWordDoesNotExist
    case nil:
        dict[word] = def
    default:
        return err
    }

    return nil
}

func (dict Dictionary) Delete (word string) {
    delete(dict, word)
}
