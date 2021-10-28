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

### Print a Recipe as Plain Text

```sh
$ cuisiner print [RECIPE_FILE]
```

Output is printed to `stdout`.

### Format a Recipe as HTML

```sh
$ cuisiner html [RECIPE_FILE]
```

Output is printed to `stdout`.

## License

This project is licensed under the MIT License.
