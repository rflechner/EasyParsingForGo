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

// Left Combines two parsers into a parser that produces a pair of the results of the two parsers and keep the first result.
func Left[A any, B any](left Parser[A], right Parser[B]) Parser[A] {
	return func(context ParsingContext) (ParseResult[A], error) {
		f := func(pair struct {
			Left  A
			Right B
		}) A {
			return pair.Left
		}
		c := Combine(left, right)
		return Map(c, f)(context)
	}
}

// Right Combines two parsers into a parser that produces a pair of the results of the two parsers and keep the second result.
func Right[A any, B any](left Parser[A], right Parser[B]) Parser[B] {
	return func(context ParsingContext) (ParseResult[B], error) {
		f := func(pair struct {
			Left  A
			Right B
		}) B {
			return pair.Right
		}
		c := Combine(left, right)
		return Map(c, f)(context)
	}
}
