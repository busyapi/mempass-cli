# mempass-cli

## Introduction

CLI program to generate strong yet memorable passwords.

Based on the [mempass](https://github.com/busyapi/mempass) library.

## Installation

Download source and then run `make && make install`.

## Usage

```sh
$ mempass -h
Usage: mempass [<password-count>]

Arguments:
  [<password-count>]    Number of passwords to generate. Default is 1

Flags:
  -h, --help                      Show context-sensitive help.
  -r, --use-rand                  Use randomly generated memorable words instead of dictionary words
  -c, --word-count=3              Number of words to generate. Using less than 2 is discouraged. Default is 3
  -m, --min-word-length=6         Minimum word length. O = no minimum. Using less than 4 is discouraged. Default is 6
  -M, --max-word-length=8         Maximum word length. O = no maximum. Default is 8
  -d, --digits-after=0            Number of digits to add at the end of each word. Default is 0
  -D, --digits-before=0           Number of digits to add at the begining of each word. Default is 0
  -u, --uppercase-rule="none"     Capitalization rule. Default is none
  -R, --uppercase-ratio=0.2       Uppercase ratio. 0.0 = no uppercase, 1.0 = all uppercase, 0.3 = 1/3 uppercase, etc.
                                  Only used if --uppercase-rule is random. Default is 0.2
  -s, --symbols-after=0           Number of symbols to add at the end of each word. Default is 0
  -S, --symbols-before=0          Number of symbols to add at the begining of each word. Default is 0
  -y, --symbol-pool=STRING        Symbols pool. Only used if --symbols-before and/or --symbols-after are set. Default
                                  is '@&!-_^$*%,.;:/=+'
  -Y, --symbol=CHAR-AS-BYTE       Symbol character. Only used if --symbols-before and/or --symbols-after are set.
                                  Default is /
  -t, --separator-rule="fixed"    Separator rule. Default is 'fixed'
  -e, --separator-pool=STRING     Seperators pool. Only used if --separator-rule is random. Default is
                                  '@&!-_^$*%,.;:/=+'
  -E, --separator=CHAR-AS-BYTE    Separator character. Only used if --separator-rule is fixed. Default is '-'
  -a, --padding-rule="fixed"      Padding rule. Only used if --padding-length is greater than 0
  -A, --padding-symbol=.          Padding symbol. Only used if --padding-rule is fixed. Default is '.'
  -l, --padding-length=UINT       Password length to reach with padding.
  -L, --leet-ratio=0               1337 coding ratio. 0.0 = no 1337, 1.0 = all 1337, 0.3 = 1/3 1337, etc. Default is 0
  -n, --calculate-entropy         Calculate entropy. Default is false
  -o, --output="simple"           Output format (simple, json). Default is simple
```
