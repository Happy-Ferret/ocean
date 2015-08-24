/*
Copyright 2014 Sean Hickey <sean@headzoo.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package ocean

// RuneClass is the type of a UTF-8 character; a character, quote, space, escape.
type RuneClass string

// RuneClassMap is a map of RuneTokeType values.
type RuneClassMap map[rune]RuneClass

const (
	CLASS_SPACE             string = " \t\r\n"
	CLASS_ESCAPING_QUOTE    string = "\""
	CLASS_NONESCAPING_QUOTE string = "'"
	CLASS_ESCAPE            string = "\\"
	CLASS_PIPE              string = "|"
	CLASS_REDIRECT          string = "><"

	RUNE_UNKNOWN      RuneClass = "UNKNOWN"
	RUNE_CHAR         RuneClass = "CHAR"
	RUNE_SPACE        RuneClass = "SPACE"
	RUNE_QUOTE_DOUBLE RuneClass = "QUOTE_DOUBLE"
	RUNE_QUOTE_SINGLE RuneClass = "QUOTE_SINGLE"
	RUNE_ESCAPE       RuneClass = "ESCAPE"
	RUNE_PIPE         RuneClass = "PIPE"
	RUNE_REDIRECT     RuneClass = "REDIRECT"
	RUNE_EOF          RuneClass = "EOF"
)

// Classifier classifies runes by type. This allows for different sorts of
// classifiers - those accepting extended non-ascii chars, or strict posix
// compatibility, for example.
type Classifier struct {
	typeMap RuneClassMap
}

// Create and returns a new rune classifier.
func NewClassifier() *Classifier {
	classifier := &Classifier{
		typeMap: make(RuneClassMap),
	}

	classifier.AddClassification(CLASS_SPACE, RUNE_SPACE)
	classifier.AddClassification(CLASS_ESCAPING_QUOTE, RUNE_QUOTE_DOUBLE)
	classifier.AddClassification(CLASS_NONESCAPING_QUOTE, RUNE_QUOTE_SINGLE)
	classifier.AddClassification(CLASS_ESCAPE, RUNE_ESCAPE)
	classifier.AddClassification(CLASS_PIPE, RUNE_PIPE)
	classifier.AddClassification(CLASS_REDIRECT, RUNE_REDIRECT)

	return classifier
}

// Classify returns the rune token type.
func (classifier *Classifier) Classify(r rune) RuneClass {
	if v, ok := classifier.typeMap[r]; ok {
		return v
	}
	
	// everything else is a char
	return RUNE_CHAR
}

// addRuneClass registers a rune and it's classification.
func (classifier *Classifier) AddClassification(runes string, runeType RuneClass) {
	for _, r := range runes {
		classifier.typeMap[r] = runeType
	}
}
