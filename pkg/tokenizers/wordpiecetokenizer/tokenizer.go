// Copyright 2020 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wordpiecetokenizer

import (
	"strings"

	"github.com/rh1nox/cybertron/pkg/tokenizers"
	"github.com/rh1nox/cybertron/pkg/tokenizers/basetokenizer"
	"github.com/rh1nox/cybertron/pkg/vocabulary"
)

const (
	// DefaultClassToken is the default class token value for the WordPiece tokenizer.
	DefaultClassToken = "[CLS]"
	// DefaultSequenceSeparator is the default sequence separator value for the WordPiece tokenizer.
	DefaultSequenceSeparator = "[SEP]"
	// DefaultUnknownToken is the default unknown token value for the WordPiece tokenizer.
	DefaultUnknownToken = "[UNK]"
	// DefaultMaskToken is the default mask token value for the WordPiece tokenizer.
	DefaultMaskToken = "[MASK]"
	// DefaultSplitPrefix is the default split prefix value for the WordPiece tokenizer.
	DefaultSplitPrefix = "##"
	// DefaultMaxWordChars is the default maximum word length for the WordPiece tokenizer.
	DefaultMaxWordChars = 100
)

var defaultNeverSplit = []string{
	DefaultClassToken,
	DefaultSequenceSeparator,
	DefaultUnknownToken,
	DefaultMaskToken,
}

var _ tokenizers.Tokenizer = &WordPieceTokenizer{}

// WordPieceTokenizer is a tokenizer that breaks tokens into sub-word units based on a supplied vocabulary.
// See https://arxiv.org/pdf/1609.08144.pdf Section 4.1 for details.
// WordPieceTokenizers uses BaseTokenizer to preprocess the input text.
type WordPieceTokenizer struct {
	baseTokenizer *basetokenizer.BaseTokenizer
	vocabulary    *vocabulary.Vocabulary
	unkToken      string
	splitPrefix   string
	maxWordChars  int
	neverSplit    []string
}

// New returns a new WordPieceTokenizer.
func New(vocabulary *vocabulary.Vocabulary) *WordPieceTokenizer {
	return &WordPieceTokenizer{
		baseTokenizer: basetokenizer.New(
			basetokenizer.RegisterSpecialWords(DefaultUnknownToken, DefaultClassToken, DefaultSequenceSeparator, DefaultMaskToken)),
		vocabulary:   vocabulary,
		unkToken:     DefaultUnknownToken,
		splitPrefix:  DefaultSplitPrefix,
		maxWordChars: DefaultMaxWordChars,
		neverSplit:   defaultNeverSplit,
	}
}

// Tokenize converts the input text to a slice of words or sub-words token units based on the supplied vocabulary.
// The resulting tokens preserve the alignment with the portion of the original text they belong to.
func (t *WordPieceTokenizer) Tokenize(text string) []tokenizers.StringOffsetsPair {
	return t.WordPieceTokenize(t.baseTokenizer.Tokenize(text))
}

// WordPieceTokenize transforms the input token in a new slice of words or sub-words units based on the supplied vocabulary.
// The resulting tokens preserve the alignment with the portion of the original text they belong to.
func (t *WordPieceTokenizer) WordPieceTokenize(tokens []tokenizers.StringOffsetsPair) []tokenizers.StringOffsetsPair {
	outputTokens := make([]tokenizers.StringOffsetsPair, 0)

	for _, stringOffsetsPair := range tokens {
		token := stringOffsetsPair.String
		initialOffsets := stringOffsetsPair.Offsets
		characters := []rune(token)

		if len(characters) > t.maxWordChars {
			if _, exists := t.vocabulary.ID(t.unkToken); !exists {
				panic("Missing unk-token")
			}
			outputTokens = append(outputTokens, tokenizers.StringOffsetsPair{
				String:  t.unkToken,
				Offsets: initialOffsets,
			})
			continue
		}

		isBad := false
		start := 0
		subTokens := make([]tokenizers.StringOffsetsPair, 0)

		for start < len(characters) {
			end := len(characters)
			var curStrToken tokenizers.StringOffsetsPair
			found := false

			for start < end {
				subStr := string(characters[start:end])
				if start > 0 {
					subStr = t.splitPrefix + subStr
				}

				if _, exists := t.vocabulary.ID(subStr); exists {
					found = true
					curStrToken.String = subStr
					curStrToken.Offsets = tokenizers.OffsetsType{
						Start: initialOffsets.Start + start,
						End:   initialOffsets.Start + end,
					}
					break
				}
				end--
			}
			if !found {
				isBad = true
				break
			}
			subTokens = append(subTokens, curStrToken)
			start = end
		}

		if isBad {
			if _, exists := t.vocabulary.ID(t.unkToken); !exists {
				panic("Missing unk-token")
			}
			outputTokens = append(outputTokens, tokenizers.StringOffsetsPair{
				String:  t.unkToken,
				Offsets: initialOffsets,
			})
		} else {
			outputTokens = append(outputTokens, subTokens...)
		}
	}
	return outputTokens
}

// IsDefaultSpecial return whether the word matches a special token, or not.
func IsDefaultSpecial(word string) bool {
	switch word {
	case DefaultUnknownToken, DefaultClassToken, DefaultSequenceSeparator, DefaultMaskToken:
		return true
	default:
		return false
	}
}

// GroupSubWords returns a list of tokens range each of which represents
// the start and the end index of the tokens that form a complete word.
func GroupSubWords(tokens []tokenizers.StringOffsetsPair) []tokenizers.StringOffsetsPair {
	result := make([]tokenizers.StringOffsetsPair, 0)
	for _, token := range tokens {
		if strings.HasPrefix(token.String, DefaultSplitPrefix) {
			last := &result[len(result)-1]
			last.String += token.String[len(DefaultSplitPrefix):]
			last.Offsets.End = token.Offsets.End
		} else {
			result = append(result, tokenizers.StringOffsetsPair{
				String:  token.String,
				Offsets: token.Offsets,
			})
		}
	}
	return result
}
