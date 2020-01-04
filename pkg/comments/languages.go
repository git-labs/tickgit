package comments

import "github.com/augmentable-dev/lege"

// CStyleCommentOptions ...
var CStyleCommentOptions *lege.ParseOptions = &lege.ParseOptions{
	Boundaries: []lege.Boundary{
		{
			Start: "//",
			End:   "\n",
		},
		{
			Start: "/*",
			End:   "*/",
		},
	},
}

// HashStyleCommentOptions ...
var HashStyleCommentOptions *lege.ParseOptions = &lege.ParseOptions{
	Boundaries: []lege.Boundary{
		{
			Start: "#",
			End:   "\n",
		},
	},
}

// LispStyleCommentOptions ..
var LispStyleCommentOptions *lege.ParseOptions = &lege.ParseOptions{
	Boundaries: []lege.Boundary{
		{
			Start: ";",
			End:   "\n",
		},
	},
}

// Language is a source language (i.e. "Go")
type Language string

// LanguageParseOptions keeps track of source languages and their corresponding comment options
var LanguageParseOptions map[Language]*lege.ParseOptions = map[Language]*lege.ParseOptions{
	"Go":           CStyleCommentOptions,
	"Java":         CStyleCommentOptions,
	"C":            CStyleCommentOptions,
	"C++":          CStyleCommentOptions,
	"C#":           CStyleCommentOptions,
	"JavaScript":   CStyleCommentOptions,
	"Python":       HashStyleCommentOptions,
	"Ruby":         HashStyleCommentOptions,
	"PHP":          CStyleCommentOptions,
	"Shell":        HashStyleCommentOptions,
	"Visual Basic": {Boundaries: []lege.Boundary{{Start: "'", End: "\n"}}},
	"TypeScript":   CStyleCommentOptions,
	"Objective-C":  CStyleCommentOptions,
	"Groovy":       CStyleCommentOptions,
	"Swift":        CStyleCommentOptions,
	"Common Lisp":  LispStyleCommentOptions,
	"Emacs Lisp":   LispStyleCommentOptions,
	"R":            HashStyleCommentOptions,

	// TODO Currently, the underlying pkg that does the parsing/plucking (lege) doesn't properly support precedance
	// so lines beginning with /// or //! will be picked up by this start // and include a / or ! preceding the comment
	"Rust": {Boundaries: []lege.Boundary{{Start: "///", End: "\n"}, {Start: "//!", End: "\n"}, {Start: "//", End: "\n"}}},
}
