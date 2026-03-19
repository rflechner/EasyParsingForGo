# EasyParsingForGo
Easy parsing for GO / Parser combinator

Port of [EasyParsing](https://github.com/rflechner/EasyParsing) for Golang.

__STATUS__: WIP – not yet ready for production

# Parser Combinators

Some basic parser combinators are implemented.

## Currently implemented

- `OneChar(rune)`: Matches a single specified character.
- `Satisfy(predicate)`: Matches a character that satisfies a given predicate.
- `AnyChar()`: Matches any character.
- `StringMatch(string)`: Matches an exact string as a prefix.
- `Digit()`: Matches a single digit.
- `Integer()`: Matches an integer.
- `Spaces()`: Matches a sequence of spaces, tabs, or newlines.
- `Many(parser)`: Attempts to apply a parser multiple times and collects the results into a slice.
- `Optional(parser)`: Attempts to apply a parser, returns `None` if it fails instead of causing an error.
- `OrElse(parsers...)`: Attempts several parsers in sequence until one succeeds.
- `Combine(left, right)`: Combines two parsers to run one after the other.
- `Map(parser, function)`: Applies a transformation function to a parser's result.
- `Between(before, middle, after)`: Parses content located between two other elements.
- `UntilText(parser, delimiter, includeDelimiter)`: Applies a parser to all text until a specific delimiter is encountered.
- `SeparatedBy(parser, separator, matchTrailing)`: Parses a list of elements separated by a delimiter.
- `LazyParse(factory)`: Allows defining recursive parsers by delaying their creation.

## To do
