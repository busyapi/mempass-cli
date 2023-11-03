package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/busyapi/mempass"
)

func main() {
	var CLI struct {
		UseRand          bool            `short:"r" help:"Use randomly generated memorable words instead of dictionary words"`
		WordCount        uint            `short:"c" help:"Number of words to generate. Using less than 2 is discouraged. Default is 3" default:"3"`
		MinWordLength    uint            `short:"m" help:"Minimum word length. O = no minimum. Using less than 4 is discouraged. Default is 6" default:"6"`
		MaxWordLength    uint            `short:"M" help:"Maximum word length. O = no maximum. Default is 8" default:"8"`
		DigitsAfter      uint            `short:"d" help:"Number of digits to add at the end of each word. Default is 0" default:"0"`
		DigitsBefore     uint            `short:"D" help:"Number of digits to add at the begining of each word. Default is 0" default:"0"`
		UppercaseRule    mempass.CapRule `short:"u" help:"Capitalization rule. Default is none" default:"none" enum:"none,all,alternate,word_alternate,first_letter,last_letter,all_but_first_letter,all_but_last_letter,random"`
		UppercaseRatio   float32         `short:"R" help:"Uppercase ratio. 0.0 = no uppercase, 1.0 = all uppercase, 0.3 = 1/3 uppercase, etc. Only used if --uppercase-rule is random. Default is 0.2" default:"0.2"`
		SymbolsAfter     uint            `short:"s" help:"Number of symbols to add at the end of each word. Default is 0" default:"0"`
		SymbolsBefore    uint            `short:"S" help:"Number of symbols to add at the begining of each word. Default is 0" default:"0"`
		SymbolPool       string          `short:"y" help:"Symbols pool. Only used if --symbols-before and/or --symbols-after are set. Default is '@&!-_^$*%,.;:/=+'" xor:"symbol"`
		Symbol           CharAsByte      `short:"Y" help:"Symbol character. Only used if --symbols-before and/or --symbols-after are set. Default is /" xor:"symbol"`
		SeparatorRule    mempass.SepRule `short:"t" help:"Separator rule. Default is 'fixed'" enum:"none,fixed,random" default:"fixed"`
		SeparatorPool    string          `short:"e" help:"Seperators pool. Only used if --separator-rule is random. Default is '@&!-_^$*%,.;:/=+'" xor:"separator"`
		Separator        CharAsByte      `short:"E" help:"Separator character. Only used if --separator-rule is fixed. Default is '-'" xor:"separator"`
		PaddingRule      mempass.PadRule `short:"a" help:"Padding rule. Only used if --padding-length is greater than 0" enum:"fixed,random" default:"fixed"`
		PaddingSymbol    CharAsByte      `short:"A" help:"Padding symbol. Only used if --padding-rule is fixed. Default is '.'" default:"."`
		PaddingLength    uint            `short:"l" help:"Password length to reach with padding."`
		LeetRatio        float32         `short:"L" help:"1337 coding ratio. 0.0 = no 1337, 1.0 = all 1337, 0.3 = 1/3 1337, etc. Default is 0" default:"0"`
		CalculateEntropy bool            `short:"n" help:"Calculate entropy. Default is false"`
		PasswordCount    uint            `arg:"" help:"Number of passwords to generate. Default is 1" default:"1"`
		Output           string          `short:"o" help:"Output format (simple, json). Default is simple" enum:"simple,json" default:"simple"`
	}

	kong.Parse(&CLI)

	opt := mempass.Options{
		UseRand:          CLI.UseRand,
		WordCount:        CLI.WordCount,
		MinWordLength:    CLI.MinWordLength,
		MaxWordLength:    CLI.MaxWordLength,
		DigitsAfter:      CLI.DigitsAfter,
		DigitsBefore:     CLI.DigitsBefore,
		CapRule:          CLI.UppercaseRule,
		CapRatio:         CLI.UppercaseRatio,
		SymbolsAfter:     CLI.SymbolsAfter,
		SymbolsBefore:    CLI.SymbolsBefore,
		SymbolPool:       CLI.SymbolPool,
		Symbol:           byte(CLI.Symbol),
		SepRule:          CLI.SeparatorRule,
		SeparatorPool:    CLI.SeparatorPool,
		Separator:        byte(CLI.Separator),
		PadRule:          CLI.PaddingRule,
		PadSymbol:        byte(CLI.PaddingSymbol),
		PadLength:        CLI.PaddingLength,
		L33tRatio:        CLI.LeetRatio,
		CalculateEntropy: CLI.CalculateEntropy,
	}

	for i := 0; i < int(CLI.PasswordCount); i++ {
		gen := mempass.NewGenerator(&opt)
		pwd, ent, _ := gen.GenPassword()
		fmt.Println(format(pwd, ent, CLI.Output))
	}
}

// CharAsByte is a type that will store the byte value of the provided character.
type CharAsByte byte

// UnmarshalText ensures that only a single character is used for the flag
// and converts it to its byte value.
func (c *CharAsByte) UnmarshalText(text []byte) error {
	if len(text) != 1 {
		return fmt.Errorf("the input must be exactly one character long")
	}
	*c = CharAsByte(text[0])
	return nil
}

func format(pwd string, ent float64, format string) string {
	var output string

	if format == "json" || format == "JSON" {
		output = fmt.Sprintf(`{"password":"%v"`, pwd)

		if ent > 0 {
			output += fmt.Sprintf(`,"entropy":%v`, ent)
		}

		output += "}"
	} else {
		if ent > 0 {
			output = fmt.Sprintf("%v %v", pwd, ent)
		} else {
			output = pwd
		}
	}

	return output
}
