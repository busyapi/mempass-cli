# mempass-cli

## Introduction

CLI program to generate strong yet memorable passwords.

Based on the [mempass](https://github.com/busyapi/mempass) library.

## Installation

### From source

First be sure you have a valid [Go](https://go.dev/dl/) installation.

Then clone this repo and run `make && make install` (or `make && sudo make install` depending on the permissions on the install directory).

The default install directory is `/usr/local/bin`. You can change it by running `make install INSTALL_DIR=/path/to/install/dir`

### Binary

SOON

## Usage

### Options

```sh
$ mempass -h
Usage: mempass

Flags:
  -h, --help                      Show context-sensitive help.
  -m, --mode="dict"               Generation mode. Possible value:'dict,rand,passphrase'. Read passphrase from stdin. Default is 'dict'
  -c, --word-count=3              Number of words to generate. Using less than 2 is discouraged. Default is 3
  -w, --min-word-length=6         Minimum word length. O = no minimum. Using less than 4 is discouraged. Default is 6
  -W, --max-word-length=8         Maximum word length. O = no maximum. Default is 8
  -d, --digits-after=0            Number of digits to add at the end of each word. Default is 0
  -D, --digits-before=0           Number of digits to add at the begining of each word. Default is 0
  -u, --uppercase-rule="none"     Capitalization rule. Possible value:'none,all,alternate,word_alternate,first_letter,last_letter,all_but_first_letter,all_but_last_letter,random'. Default is none
  -R, --uppercase-ratio=0.2       Uppercase ratio. 0.0 = no uppercase, 1.0 = all uppercase, 0.3 = 1/3 uppercase, etc. Only used if --uppercase-rule is random. Default is 0.2
  -s, --symbols-after=0           Number of symbols to add at the end of each word. Default is 0
  -S, --symbols-before=0          Number of symbols to add at the begining of each word. Default is 0
  -y, --symbol-pool=STRING        Symbols pool. Only used if --symbols-before and/or --symbols-after are set. Default is '@&!-_^$*%,.;:/=+'
  -Y, --symbol=CHAR               Symbol character. Only used if --symbols-before and/or --symbols-after are set. Default is /
  -t, --separator-rule="fixed"    Separator rule. Possible value:'fixed,random'. Default is 'fixed'
  -e, --separator-pool=STRING     Seperators pool. Only used if --separator-rule is random. Default is '@&!-_^$*%,.;:/=+'
  -E, --separator=CHAR            Separator character. Only used if --separator-rule is fixed. Default is '-'
  -a, --padding-rule="fixed"      Padding rule. Possible value:'fixed,random'. Only used if --padding-length is greater than 0
  -A, --padding-symbol=.          Padding symbol. Only used if --padding-rule is fixed. Default is '.'
  -l, --padding-length=UINT       Password length to reach with padding.
  -L, --leet-ratio=0              1337 coding ratio. 0.0 = no 1337, 1.0 = all 1337, 0.3 = 1/3 1337, etc. Default is 0
  -n, --calculate-entropy         Calculate entropy. Default is false
  -T, --password-count=1          Number of passwords to generate. Default is 1
  -o, --output="simple"           Output format (simple, json). Default is simple
  -C, --config=CONFIG-FLAG        Path to config file
```

### Passphrase mode

With this mode you can input a passphrase and the program will transform it into a strong password by adding uppercase letters, numbers and symbols.

You should enter a passphrase of at leat 12 ou 14 characters to get good results.

```bash
$ echo "Generate a strong password" | mempass -mpassphrase -T5
Gen3raTe-a-strong-pA55worD
G3Nera7e-a-sTronG-p4ssword
Gener47E-A-s7rong-passworD
GENerate-4-s7r0ng-paSsword
GEneraT3-a-s7rong-pa5sworD
```

## Configuration file

By default `mempass` will try to read `/etc/mempass/mempass.json` to get its configuration. You can set a different path with the `-C` flag. Parameters that are set on the command line will override those defined in the configuration file.

The configuration file is a simple JSON in which field names are camelCase version of the flags, i.e. `useRand` for `--use-rand`, `wordCount` for `--word-count`, etc.

Here is an example:

```json
{
  "Mode": "rand",
  "minWordLength": 4,
  "maxWordLength": 6,
  "separator": "!"
}
```
