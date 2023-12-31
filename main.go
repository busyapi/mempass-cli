package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/busyapi/mempass"
)

func main() {
	var cli Cli

	kong.Parse(&cli, kong.Configuration(kong.JSON, "/etc/mempass/mempass.json"))

	opt := mempass.Options{
		Mode:             cli.Mode,
		WordCount:        cli.WordCount,
		MinWordLength:    cli.MinWordLength,
		MaxWordLength:    cli.MaxWordLength,
		DigitsAfter:      cli.DigitsAfter,
		DigitsBefore:     cli.DigitsBefore,
		CapRule:          cli.UppercaseRule,
		CapRatio:         cli.UppercaseRatio,
		SymbolsAfter:     cli.SymbolsAfter,
		SymbolsBefore:    cli.SymbolsBefore,
		SymbolPool:       cli.SymbolPool,
		Symbol:           rune(cli.Symbol),
		SepRule:          cli.SeparatorRule,
		SeparatorPool:    cli.SeparatorPool,
		Separator:        rune(cli.Separator),
		PadRule:          cli.PaddingRule,
		PadSymbol:        rune(cli.PaddingSymbol),
		PadLength:        cli.PaddingLength,
		L33tRatio:        cli.LeetRatio,
		CalculateEntropy: cli.CalculateEntropy,
	}

	if cli.Mode == mempass.ModePassphrase {
		opt.Passphrase = readLine()
	}

	pwds := make([]string, cli.PasswordCount)

	for i := 0; i < int(cli.PasswordCount); i++ {
		gen := mempass.NewGenerator(&opt)
		pwd, ent, _ := gen.GenPassword()
		pwds[i] = format(pwd, ent, cli.Output)
	}

	fmt.Println(formatOutput(pwds, cli.Output))
}

func formatOutput(pwds []string, format string) string {
	if len(pwds) == 1 {
		return pwds[0]
	}

	if format == "json" {
		return fmt.Sprintf("[%v]", strings.Join(pwds, ","))
	}

	return strings.Join(pwds, "\n")
}

type Cli struct {
	Mode             mempass.Mode    `short:"m" help:"Generation mode. Possible value:'dict,rand,passphrase'. Read passphrase from stdin. Default is 'dict'" enum:"dict,rand,passphrase" default:"dict" json:"mode"`
	WordCount        uint            `short:"c" help:"Number of words to generate. Using less than 2 is discouraged. Default is 3" default:"3" json:"word_count"`
	MinWordLength    uint            `short:"w" help:"Minimum word length. O = no minimum. Using less than 4 is discouraged. Default is 6" default:"6" json:"minWordLength"`
	MaxWordLength    uint            `short:"W" help:"Maximum word length. O = no maximum. Default is 8" default:"8" json:"maxWordLength"`
	DigitsAfter      uint            `short:"d" help:"Number of digits to add at the end of each word. Default is 0" default:"0" json:"digitsAfter"`
	DigitsBefore     uint            `short:"D" help:"Number of digits to add at the begining of each word. Default is 0" default:"0" json:"digitsBefore"`
	UppercaseRule    mempass.CapRule `short:"u" help:"Capitalization rule. Possible value:'none,all,alternate,word_alternate,first_letter,last_letter,all_but_first_letter,all_but_last_letter,random'. Default is none" default:"none" enum:"none,all,alternate,word_alternate,first_letter,last_letter,all_but_first_letter,all_but_last_letter,random" json:"uppercaseRule"`
	UppercaseRatio   float32         `short:"R" help:"Uppercase ratio. 0.0 = no uppercase, 1.0 = all uppercase, 0.3 = 1/3 uppercase, etc. Only used if --uppercase-rule is random. Default is 0.2" default:"0.2" json:"uppercaseRatio"`
	SymbolsAfter     uint            `short:"s" help:"Number of symbols to add at the end of each word. Default is 0" default:"0" json:"symbolsAfter"`
	SymbolsBefore    uint            `short:"S" help:"Number of symbols to add at the begining of each word. Default is 0" default:"0" json:"symbolsBefore"`
	SymbolPool       string          `short:"y" help:"Symbols pool. Only used if --symbols-before and/or --symbols-after are set. Default is '@&!-_^$*%,.;:/=+'" xor:"symbol" json:"symbolPool"`
	Symbol           Char            `short:"Y" help:"Symbol character. Only used if --symbols-before and/or --symbols-after are set. Default is /" xor:"symbol" json:"symbol"`
	SeparatorRule    mempass.SepRule `short:"t" help:"Separator rule. Possible value:'fixed,random'. Default is 'fixed'" enum:"none,fixed,random" default:"fixed" json:"separatorRule"`
	SeparatorPool    string          `short:"e" help:"Seperators pool. Only used if --separator-rule is random. Default is '@&!-_^$*%,.;:/=+'" xor:"separator" json:"separatorPool"`
	Separator        Char            `short:"E" help:"Separator character. Only used if --separator-rule is fixed. Default is '-'" xor:"separator" json:"separator"`
	PaddingRule      mempass.PadRule `short:"a" help:"Padding rule. Possible value:'fixed,random'. Only used if --padding-length is greater than 0" enum:"fixed,random" default:"fixed" json:"paddingRule"`
	PaddingSymbol    Char            `short:"A" help:"Padding symbol. Only used if --padding-rule is fixed. Default is '.'" default:"." json:"addingSymbol"`
	PaddingLength    uint            `short:"l" help:"Password length to reach with padding." json:"paddingLength"`
	LeetRatio        float32         `short:"L" help:"1337 coding ratio. 0.0 = no 1337, 1.0 = all 1337, 0.3 = 1/3 1337, etc. Default is 0" default:"0" json:"leetRatio"`
	CalculateEntropy bool            `short:"n" help:"Calculate entropy. Default is false" json:"calculateEntropy"`
	PasswordCount    uint            `short:"T" help:"Number of passwords to generate. Default is 1" default:"1" json:"passwordCount"`
	Output           string          `short:"o" help:"Output format (simple, json). Default is simple" enum:"simple,json" default:"simple" json:"output"`
	Config           kong.ConfigFlag `short:"C" help:"Path to config file"`
}

// Char is a type that will store the byte value of the provided character.
type Char byte

// UnmarshalText ensures that only a single character is used for the flag
// and converts it to its byte value.
func (c *Char) UnmarshalText(text []byte) error {
	if len(text) != 1 {
		return fmt.Errorf("the input must be exactly one character long")
	}
	*c = Char(text[0])
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

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Error reading from stdin:", err)
	}

	// Remove the newline character from the end of input if needed
	return strings.TrimSpace(text)
}
