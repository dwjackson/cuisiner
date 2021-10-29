# Cuisiner: Command-Line Recipe Management

Cuisiner is a command-line utility to manage recipes. It uses (a subset of) the
[Cook Language](https://cooklang.org/docs/spec/) to define recipes. The name of
the program is both the French word for "to cook" and also bad-English for
"person who does cuisine."

## Usage

The `cuisiner` command line tool is used as follows:

```sh
$ cuisiner [COMMAND] [ARGS...]
```

### Print a Recipe

```sh
$ cuisiner print [RECIPE_FILE] [...PANTRY_FILE]
```

If a "pantry" file is specified then those ingredients are removed from the
list, on the assumption that they needn't be bought. The pantry file is in
recipe format without any real directions.

Output is printed to `stdout` and is in markdown.

### Print a Shopping List

Cuisiner can be used to create a shopping list from a bunch of recipes. Input
is taken from `stdin`, one file name per line.

```sh
$ cuisiner shopping
recipe1.cook
recipe2.cook
recipe3.cook
^D
```

The output is printed to `stdout` as is in markdown.

## Note: Fractional Amounts

At the moment, fractional amounts can only be written as decimals in recipes.
For example, the amount 1/8 must be written as 0.125.

## License

This project is licensed under the MIT License.
