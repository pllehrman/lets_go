package main

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {

	cases :=[]struct {
		Description	string
		Arabic 		int
		Want 		string
	}{
		{"1 gets converted I", 1, "I"},
		{"2 gets converted II", 2, "II"},
		{"3 gets converted III", 3, "III"},
		{"4 gets converted III", 4, "IV"},


	}
	
	for _,test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
	
}