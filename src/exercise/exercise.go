package exercise

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/leveltype/src/config"
	"github.com/leveltype/src/problemwords"
)

var (
	CleanPasses = 0
)

// Exercise tracks the state of the current grinding problem the user is working on.
type Exercise struct {
	Text          string            // The original exercize text
	Words         []string          // The words to be keyed
	RenderedWords []string          // The words color-formatted for display
	WrongWord     []bool            // WrongWord[i] is true if Words[i] was a typo
	WordStart     []time.Time       // WordStart[i] = Time when typing of Words[i] began
	WordEnd       []time.Time       // WordEnd[i] = Time when typing of Words[i] finished
	WordCount     int               // Len(Words)
	WrongWords    map[string]int    // Counts the wrong words entered and # of times they were wrong
	PreWrongWords map[string]string // A word that was detected before a problem word, used to train spacegrams
	CurIdx        int               // Current word index, tracks the user typing
	CurText       string            // The text the user has typed for this word
	LastEntry     string            // The last entry the user made before advancing to the next word
}

// Load preps a batch of text for the user to grind on; and progress to be tracked.
func (e *Exercise) Load(words string) {
	e.Text = words
	e.Words = strings.Split(words, " ")
	e.RenderedWords = strings.Split(words, " ")
	e.WrongWords = make(map[string]int)          // summary of wrong words
	e.PreWrongWords = make(map[string]string)    // (optional) the word that came before the typo
	e.WordCount = len(e.Words)                   // per-word wrong flag on that word
	e.WrongWord = make([]bool, e.WordCount)      // flags a problem word
	e.WordStart = make([]time.Time, e.WordCount) // marks when a word was started
	e.WordEnd = make([]time.Time, e.WordCount)   // marks when a word was completed
	e.LastEntry = ""

	for i := 0; i < len(e.Words); i++ {
		e.Words[i] = strings.TrimSpace(e.Words[i])
	}
}

// WordEntry deals with each typechar and checks to see if a typo was made
func (e *Exercise) WordEntry(char rune) {
	if e.WordStart[e.CurIdx].IsZero() {
		e.WordStart[e.CurIdx] = time.Now()
	}

	if char == ' ' {
		e.LastEntry = e.CurText
		e.Advance()
		return
	}

	if unicode.IsLetter(char) {
		e.CurText += string(char)
	}

}

// Called to close out the exercize in progress and tally-up the wrong words
// and decrement the correct words
func (e *Exercise) CompleteExercize() {
	// Tally the wrong words (the pre-word accumulation is done in Advance() )
	for i, wrong := range e.WrongWord {
		if wrong {
			e.WrongWords[e.Words[i]]++
		} else { // Give credit for fixed words
			e.WrongWords[e.Words[i]]--
		}
	}
	problemwords.UpdateWords(e.WrongWords, e.PreWrongWords)

	// Was this a clean pass?
	if len(problemwords.PWords) == 0 {
		CleanPasses++
		if CleanPasses%3 == 0 {
			problemwords.ListSize = problemwords.ListSize + 20
		}
	}

	// Save vocab progress
	config.SaveVocabularyLevel(problemwords.ListSize)

	// Reset exercize
	var lesson = problemwords.E200BuildFocusLesson()
	e.Load(lesson)

}

// Advance is triggered when the spacebar is pressed to advanced to the next word. It's at this point we
// can analyse the word that was typed.
func (e *Exercise) Advance() {

	if e.CurIdx == e.WordCount-1 {
		// Does the word match?
		e.WrongWord[e.CurIdx] = e.Words[e.CurIdx] != e.CurText

		// If it did not match, can we save the prior word?
		if e.WrongWord[e.CurIdx] && e.CurIdx > 0 {
			e.PreWrongWords[e.Words[e.CurIdx]] = e.Words[e.CurIdx-1]
		}

		// Stop
		e.CurText = ""
		e.CurIdx = 0
		e.CompleteExercize()
	} else {

		// Does the word match?
		e.WrongWord[e.CurIdx] = e.Words[e.CurIdx] != e.CurText

		// If it did not match, can we save the prior word?
		if e.WrongWord[e.CurIdx] && e.CurIdx > 0 {
			e.PreWrongWords[e.Words[e.CurIdx]] = e.Words[e.CurIdx-1]
		}

		// Advance to the next word
		e.CurText = ""
		e.WordEnd[e.CurIdx] = time.Now()
		e.CurIdx++

	}

}

func (e *Exercise) RenderStats() string {
	var words = ""
	for word, count := range problemwords.PWords {
		words += fmt.Sprintf("%s(%d) ", word, count)
	}
	var lastproblem = ""
	if e.CurIdx-1 > 0 {
		if e.WrongWord[e.CurIdx-1] {
			lastproblem = fmt.Sprintf("Word: %s\tYou typed: %s\n", e.Words[e.CurIdx-1], e.LastEntry)
		}
	}
	var stats = fmt.Sprintf("%sVocabulary: Top %s words\nProblem word: %s", lastproblem, strconv.Itoa(problemwords.ListSize), words)

	return stats
}

func (e *Exercise) Render() string {
	for i, word := range e.Words {
		if i < e.CurIdx {
			if e.WrongWord[i] {
				e.RenderedWords[i] = "[orange]" + word + "[white]"
			} else {
				e.RenderedWords[i] = "[green]" + word + "[white]"
			}
		} else if i == e.CurIdx {
			if len(e.CurText) > 0 && len(e.CurText) <= len(word) {
				if e.WrongWord[i] {
					e.RenderedWords[i] = "[pink]" + word[:len(e.CurText)] + "[gray]"
				} else {
					e.RenderedWords[i] = "[greenyellow]" + word[:len(e.CurText)] + "[gray]" + word[len(e.CurText):]
				}
			} else if len(e.CurText) > 0 && len(e.CurText) > len(word) {
				e.WrongWord[i] = true
				e.RenderedWords[i] = "[orange]" + word + "[white]"
			} else if len(e.CurText) == 0 {
				e.RenderedWords[i] = "[greenyellow]" + word + "[gray]"
			}
		} else {
			if e.WrongWord[i] {
				e.RenderedWords[i] = "[pink]" + word + "[white]"
			} else {
				e.RenderedWords[i] = "[gray]" + word + "[white]"
			}
		}
	}

	return strings.Join(e.RenderedWords, " ")
}
