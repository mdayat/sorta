# Sorta

Stands for "Sorting Algorithm," a minimalist terminal-based sorting algorithms cheat sheet to help you remember the common sorting algorithms.

## Purposes

Since I work in a field that rarely uses sorting algorithms, I need to find a way to quickly remember them. I found that building a minimalist terminal-based app would meet my needs.

## Quick Start

Since this app is built with Go, I recommend installing it using Go commands

```bash
go install github.com/mdayat/sorta@latest
```

> I assume that you have already configured your PATH variable to your Go bin's path.

Afterward, type this command and you're done

```bash
sorta bubble -v
```

> To see the available commands, type `sorta -h`.

## Usage

Available commands:

- `bubble`

> Note: Currently, the only available command is `bubble,` but I will gradually add the others.

Available flags:

- `-v` - A detailed description of the selected sorting algorithms
- `-d` - To demonstrate the selected sorting algorithms

## Examples

Print the basic description

```bash
sorta bubble
```

Print the detailed description

```bash
sorta bubble -v
```

Demonstrate the sorting algorithm

```bash
sorta bubble -d
```

> The `-d` flag prints both the unsorted and sorted numbers.
