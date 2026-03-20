package combinator

import (
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("Success: transform rune to string", func(t *testing.T) {
		input := "abc"
		ctx := NewParsingContext(input)
		p := Map(OneChar('a'), func(r rune) string {
			return "char: " + string(r)
		})

		res, err := p(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if res.Result != "char: a" {
			t.Errorf("Incorrect result: expected \"char: a\", got %q", res.Result)
		}

		if string(res.Context.Remaining) != "bc" {
			t.Errorf("Incorrect remaining context: expected \"bc\", got %q", string(res.Context.Remaining))
		}
	})

	t.Run("Success: transform int to bool", func(t *testing.T) {
		input := "123"
		ctx := NewParsingContext(input)
		p := Map(Integer(), func(i int) bool {
			return i > 100
		})

		res, err := p(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if res.Result != true {
			t.Errorf("Incorrect result: expected true, got %v", res.Result)
		}
	})

	t.Run("Failure: original parser fails", func(t *testing.T) {
		input := "abc"
		ctx := NewParsingContext(input)
		// Integer() expects digits
		p := Map(Integer(), func(i int) int {
			return i * 2
		})

		_, err := p(ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func TestLeft(t *testing.T) {
	t.Run("Success: parse two characters and return left", func(t *testing.T) {
		input := "ab"
		ctx := NewParsingContext(input)
		p := Left(OneChar('a'), OneChar('b'))

		res, err := p(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if res.Result != 'a' {
			t.Errorf("Incorrect result: expected 'a', got %q", res.Result)
		}

		if !res.Context.AtEnd() {
			t.Errorf("Expected context to be at end")
		}
	})

	t.Run("Success: parse string and int, return string", func(t *testing.T) {
		input := "abc123"
		ctx := NewParsingContext(input)
		p := Left(StringMatch("abc"), Integer())

		res, err := p(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if res.Result != "abc" {
			t.Errorf("Incorrect result: expected \"abc\", got %q", res.Result)
		}

		if !res.Context.AtEnd() {
			t.Errorf("Expected context to be at end")
		}
	})

	t.Run("Failure: left parser fails", func(t *testing.T) {
		input := "bc"
		ctx := NewParsingContext(input)
		p := Left(OneChar('a'), OneChar('b'))

		_, err := p(ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("Failure: right parser fails", func(t *testing.T) {
		input := "ac"
		ctx := NewParsingContext(input)
		p := Left(OneChar('a'), OneChar('b'))

		_, err := p(ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func TestRight(t *testing.T) {
	t.Run("Success: parse two characters and return right", func(t *testing.T) {
		input := "ab"
		ctx := NewParsingContext(input)
		p := Right(OneChar('a'), OneChar('b'))

		res, err := p(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if res.Result != 'b' {
			t.Errorf("Incorrect result: expected 'b', got %q", res.Result)
		}

		if !res.Context.AtEnd() {
			t.Errorf("Expected context to be at end")
		}
	})

	t.Run("Success: parse string and int, return int", func(t *testing.T) {
		input := "abc123"
		ctx := NewParsingContext(input)
		p := Right(StringMatch("abc"), Integer())

		res, err := p(ctx)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if res.Result != 123 {
			t.Errorf("Incorrect result: expected 123, got %v", res.Result)
		}

		if !res.Context.AtEnd() {
			t.Errorf("Expected context to be at end")
		}
	})

	t.Run("Failure: left parser fails", func(t *testing.T) {
		input := "bc"
		ctx := NewParsingContext(input)
		p := Right(OneChar('a'), OneChar('b'))

		_, err := p(ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})

	t.Run("Failure: right parser fails", func(t *testing.T) {
		input := "ac"
		ctx := NewParsingContext(input)
		p := Right(OneChar('a'), OneChar('b'))

		_, err := p(ctx)
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}
