package main

import (
    "testing"
    "reflect"
)

type Person struct {
    Name    string
    Profile Profile
}

type Profile struct {
    Age     int
    City    string
}

func TestWalk(t *testing.T) {
    
    cases := []struct {
        Name            string
        Input           interface{}
        ExpectedCalls   []string
    }{
        {
            "struct with one field",
            struct {
                Name string
            }{"Chris"},
            []string{"Chris"},
        },
        {
            "struct with two fields",
            struct {
                Name string
                City string
            }{"Chris", "London"},
            []string{"Chris", "London"},
        },
        {
            "struct with non string field",
            struct {
                Name string
                Age int 
            }{"Chris", 33},
            []string{"Chris"},
        },
        {
            "struct with nested fields",
            Person{
                "Chris",
                Profile{33, "London"},
            },
            []string{"Chris", "London"},
        },
        {
            "pointers to things",
            &Person{
                "Chris", 
                Profile{33, "London"},
            },
            []string{"Chris", "London"},
        },
        {
            "slices",
            []Profile {
                {33, "London"},
                {34, "Tokyo"},
            }, 
            []string {"London", "Tokyo"},
        },
        {
            "arrays", 
            [2]Profile {
                {33, "London"},
                {43, "Tokyo"},
            }, 
            []string {"London", "Tokyo"},
        },
    }
    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string 
            walk(test.Input, func(input string) {
                got = append(got, input)
            })

            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v, want %v", got, test.ExpectedCalls)
            }
        })
    }
    t.Run("with maps", func(t *testing.T) {
        testMap := map[string]string{
            "Cow": "Moo",
            "Sheep": "Baa",
        }

        var got []string
        walk(testMap, func(input string) {
            got = append(got, input)  
        })

        assertContains(t, got, "Moo")
        assertContains(t, got, "Baa")
    })
    t.Run("with channels", func(t *testing.T) {
        testChannel := make(chan Profile)

        go func() {
            testChannel <- Profile{33, "Berlin"}
            testChannel <- Profile{34, "Katowice"}
            close(testChannel)
        }()

        var got []string 
        want := []string{"Berlin", "Katowice"}

        walk(testChannel, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v, wanted %v", got, want)
        }
    })
}

func assertContains (t testing.TB, haystack []string, needle string) {
    t.Helper()
    contains := false
    for _, x := range haystack {
        if x == needle {
            contains = true
        }
    }
    if !contains {
        t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
    }
}
