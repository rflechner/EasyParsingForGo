package combinator

// Map Transforms the result of a parser using a provided function while preserving the parsing context.
func Map[T any, Output any](p Parser[T], f func(T) Output) Parser[Output] {
	return func(context ParsingContext) (ParseResult[Output], error) {
		result, err := p(context)
		if err != nil {
			return ParseResult[Output]{}, err
		}
		return ParseResult[Output]{
			Result:  f(result.Result),
			Context: result.Context,
		}, nil
	}
}
