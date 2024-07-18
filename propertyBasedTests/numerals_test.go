package main

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

type TestCase struct {
    Arabic  uint16
    Roman   string
}

var cases = []TestCase {
    {Arabic: 1, Roman: "I"},
    {Arabic: 2, Roman: "II"},
    {Arabic: 3, Roman: "III"},
    {Arabic: 5, Roman: "V"},
    {Arabic: 9, Roman: "IX"},
    {Arabic: 10, Roman: "X"},
    {Arabic: 14, Roman: "XIV"},
    {Arabic: 18, Roman: "XVIII"},
    {Arabic: 20, Roman: "XX"},
    {Arabic: 39, Roman: "XXXIX"},
    {Arabic: 40, Roman: "XL"},
    {Arabic: 47, Roman: "XLVII"},
    {Arabic: 49, Roman: "XLIX"},
    {Arabic: 50, Roman: "L"},
    {Arabic: 124, Roman: "CXXIV"},
    {Arabic: 267, Roman: "CCLXVII"},
    {Arabic: 584, Roman: "DLXXXIV"},
    {Arabic: 722, Roman: "DCCXXII"},
    {Arabic: 1966, Roman: "MCMLXVI"},
    {Arabic: 2002, Roman: "MMII"},
}

func TestConvertToRoman(t *testing.T) {
    for _, test := range cases {
        t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
            got := ConvertToRoman(test.Arabic)
            if got != test.Roman {
                t.Errorf("got %q, but wanted %q", got, test.Roman)
            }
        })
    }
}

func TestConvertToArabic(t *testing.T) {
    for _,test := range cases {
        t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
            got := ConvertToArabic(test.Roman)
            if got != test.Arabic {
                t.Errorf("got %d but wanted %d", got, test.Arabic)
            }
        })
    }
}

func TestPropertiesOfConversion(t *testing.T) {
    assertion := func(arabic uint16) bool {
        if arabic < 0 || arabic > 3999 {
            log.Println(arabic)
            return true
        }
        roman := ConvertToRoman(arabic)
        fromRoman := ConvertToArabic(roman)
        return fromRoman == arabic
    }

    if err := quick.Check(assertion, nil); err != nil {
        t.Error("failed Conversion Check", err)
    }
}
