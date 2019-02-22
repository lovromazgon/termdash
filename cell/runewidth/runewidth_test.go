package runewidth

import (
	"testing"

	runewidth "github.com/mattn/go-runewidth"
)

func TestRuneWidth(t *testing.T) {
	tests := []struct {
		desc      string
		runes     []rune
		eastAsian bool
		want      int
	}{
		{
			desc:  "ascii characters",
			runes: []rune{'a', 'f', '#'},
			want:  1,
		},
		{
			desc:  "non-printable characters from mattn/runewidth/runewidth_test",
			runes: []rune{'\x00', '\x01', '\u0300', '\u2028', '\u2029'},
			want:  0,
		},
		{
			desc:  "half-width runes from mattn/runewidth/runewidth_test",
			runes: []rune{'ｾ', 'ｶ', 'ｲ', '☆'},
			want:  1,
		},
		{
			desc:  "full-width runes from mattn/runewidth/runewidth_test",
			runes: []rune{'世', '界'},
			want:  2,
		},
		{
			desc:      "ambiguous so double-width in eastAsian from mattn/runewidth/runewidth_test",
			runes:     []rune{'☆'},
			eastAsian: true,
			want:      2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			runewidth.DefaultCondition.EastAsianWidth = tc.eastAsian
			defer func() {
				runewidth.DefaultCondition.EastAsianWidth = false
			}()

			for _, r := range tc.runes {
				if got := RuneWidth(r); got != tc.want {
					t.Errorf("RuneWidth(%v) => %v, want %v", r, got, tc.want)
				}
			}
		})
	}
}

func TestStringWidth(t *testing.T) {
	tests := []struct {
		desc      string
		str       string
		eastAsian bool
		want      int
	}{
		{
			desc: "ascii characters",
			str:  "hello",
			want: 5,
		},
		{
			desc: "string from mattn/runewidth/runewidth_test",
			str:  "■㈱の世界①",
			want: 10,
		},
		{
			desc:      "string in eastAsian from mattn/runewidth/runewidth_test",
			str:       "■㈱の世界①",
			eastAsian: true,
			want:      12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			runewidth.DefaultCondition.EastAsianWidth = tc.eastAsian
			defer func() {
				runewidth.DefaultCondition.EastAsianWidth = false
			}()

			if got := StringWidth(tc.str); got != tc.want {
				t.Errorf("StringWidth(%q) => %v, want %v", tc.str, got, tc.want)
			}
		})
	}
}
