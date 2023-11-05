# Dad Joke CLI Tool

I'm a Dad. I like Jokes. It is only natural to have a tool for Dad_Jokes.
[icanhazdadjoke.com](https://icanhazdadjoke.com) provides a great, free, API to a nice library of Dad Jokes.

## Credit To

Credit to [divrhino](https://dev.to/divrhino/building-a-command-line-tool-with-go-and-cobra-3mjd) for the inspiration.

## Command

### dadjoke

`go build -o bin/dadjoke && ./bin/dadjoke --help`

```bash
I'm a dad. I got jokes. Therefore, I support dad jokes. 

This is a simple CLI tool to get dad jokes from icanhazdadjoke.com to your terminal.

Usage:
  dadjoke [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  random      GET a random dad joke

Flags:
  -h, --help     help for dadjoke
  -t, --toggle   Help message for toggle

Use "dadjoke [command] --help" for more information about a command.
```

## Subcommands

### dadjoke random

`go build -o bin/dadjoke && ./bin/dadjoke random --help`

```bash
This command will GET a random dad joke from icanhazdadjoke.com and print it to your terminal.

Dad jokes are essential to life and children's development.

Usage:
  dadjoke random [flags]

Flags:
  -h, --help          help for random
      --term string   A search term to find a Dad joke with.
```

### Example

```bash
go build -o bin/dadjoke && ./bin/dadjoke random --term=foot

“Hold on, I have something in my shoe”  “I’m pretty sure it’s a foot”
```
