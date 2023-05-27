package tokenizer

import (
	"context"

	"github.com/pipebuild/parser/config"
)

const (
	// Comments
	GroovyDoc = iota
	MultiLine
	Shebang
	SingleLine

	// Reserved keywords
	Abstract
	Assert
	Break
	Case
	Catch
	Class
	Const
	Continue
	Def
	Default
	Do
	Else
	Enum
	Extends
	Final
	Finally
	For
	Goto
	If
	Implements
	Import
	Instanceof
	Interface
	Native
	New_
	Null
	NonSealed
	Package
	Public
	Protected
	Private
	Return
	Static
	Strictfp
	Super
	Switch
	Synchronized
	This
	Threadsafe
	Throw
	Throws
	Transient
	Try
	While

	// Contextual keywords
	As
	In
	Permits
	Record
	Sealed
	Trait
	Var
	Yields

	// Other reserved words
	Boolean
	Byte
	Char
	Double
	False
	Float
	Int
	Long
	Short
	True

	// Identifiers
	Normal
	Quoted

	// Strings
	Character
	DollarSlashy
	DoubleQuoted
	SingleQuoted
	Slashy
	TripleDoubleQuoted
	TripleSingleQuoted

	// Normal arithmetic operators
	// +
	Addition
	// -
	Subtraction
	// *
	Multiplication
	// /
	Division
	// %
	Remainder
	// **
	Power

	// Unary operators
	// ++
	Increment
	// --
	Decrement

	// Assignment arithmetic operators
	// +=
	AddAssign
	// -=
	SubAssign
	// *=
	MulAssign
	// /=
	DivAssign
	// %=
	RemAssign
	// **=
	PowAssign

	// Relational operators
	// ==
	Equal
	// !=
	Different
	// <
	LessThan
	// <=
	LessThanOrEqual
	// >
	GreaterThan
	// >=
	GreaterThanOrEqual
	// ===
	Identical
	// !==
	NotIdentical

	// Logical operators
	// &&
	LogicalAnd
	// ||
	LogicalOr
	// !
	LogicalNot

	// Bitwise operators
	// &
	And
	// |
	Or
	// ^
	Caret
	// ~
	Negation

	// Bit shift operators
	// <<
	LShift
	// >>
	RShift
	// >>>
	RShiftUnsigned

	// Not operator
	// !
	Not

	// Ternary operator
	Ternary

	// Elvis operator
	Elvis

	// Elvis assignment operator
	ElvisAssign

	// Object operators
	SafeNavigation
	DirectFieldAccess
	MethodPointer
	MethodReference

	// Regular expression operators
	Pattern
	Find
	Match

	// Other operators
	// =
	Assign
	Spread
	Range
	Spaceship
	Subscript
	SafeIndex
	Membership
	Identity
	Coercion
	Diamond
	Call

	// Boolean
	False_
	True_

	// Delimiters
	// {
	LBrace
	// }
	RBrace
	// [
	LBracket
	// ]
	RBracket
	// (
	LParen
	// )
	RParen

	// Misc
	Eof
)

type Tokenizer interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Config config.Config
}

type tokenizer struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Tokenizer {
	return &tokenizer{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (t *tokenizer) Init(_ context.Context) error {
	return nil
}

func (t *tokenizer) Deinit(_ context.Context) error {
	return nil
}

func (t *tokenizer) Run(ctx context.Context) error {
	return nil
}
