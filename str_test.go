package perf

import "testing"

const text string = `VS Code includes a set of built-in extensions located in the extensions folder, including grammars and snippets for many languages. Extensions that provide rich language support (code completion, Go to Definition) for a language have the suffix language-features. For example, the json extension provides coloring for JSON and the json-language-features extension provides rich language support for JSON..NOSJ rof troppus egaugnal hcir sedivorp noisnetxe serutaef-egaugnal-nosj eht dna NOSJ rof gniroloc sedivorp noisnetxe nosj eht ,elpmaxe roF .serutaef-egaugnal xiffus eht evah egaugnal a rof )noitinifeD ot oG ,noitelpmoc edoc( troppus egaugnal hcir edivorp taht snoisnetxE .segaugnal ynam rof steppins dna srammarg gnidulcni ,redlof snoisnetxe eht ni detacol snoisnetxe ni-tliub fo tes a sedulcni edoC SV`

func BenchmarkP1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !isPalindromeV1(text) {
			b.FailNow()
		}
	}
}

var benches = []struct {
	name     string
	Reverser func(string) string
}{
	{"RuneArr", ReverseWithRune},
	{"StringBuil", ReverseWithBuilder},
	{"ByteArr", ReverseWithByteArray},
}

func BenchmarkStringReverse(b *testing.B) {
	for _, bt := range benches {
		b.Run(bt.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bt.Reverser(text)
			}
		})
	}
}

func BenchmarkP2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if !isPalindromeV2(text) {
			b.FailNow()
		}
	}
}
