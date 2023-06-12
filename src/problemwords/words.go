// Package problemwords creates typing test words
package problemwords

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var (
	PWords   = make(map[string]int)
	PreWords = make(map[string]string)
	rnd      = rand.New(rand.NewSource(time.Now().Local().UnixMicro()))
	ListSize = 20 // E20 is the starting point for the vocab size
)

// Update the inventory of problem words that you need to work on
func UpdateWords(problemwords map[string]int, prewords map[string]string) {
	for word, count := range problemwords {
		PWords[word] = PWords[word] + count

		// If there is a pre-word entry, keep it.
		if w, ok := prewords[word]; ok {
			PreWords[word] = w
		}

		if PWords[word] <= 0 {
			delete(PWords, word)
			delete(PreWords, word)
		}
	}
}

// Returns the top {ListSize} English words
func WordList() string {
	vocab, err := ioutil.ReadFile("e10000.txt")
	if err != nil {
		panic("Missing dictionary in load path.  Where is e10000.txt?")
	}
	filewords := strings.Split(string(vocab), "\n")
	filecount := len(filewords)
	if ListSize > filecount {
		return strings.Join(filewords, "\n")
	}
	return strings.Join(filewords, "\n")[:ListSize]
}

// Looks in PWords for a problem word that the user needs to focus on
func E200BuildFocusLesson() string {
	var topword = ""
	var topcount = 0

	// Select the word performing word that needs focus
	for word, count := range PWords {
		if count > topcount {
			topword = word
			topcount = count
		}
	}

	if topcount > 5 {
		// Time for intense focus
		return E200IntenseFocusLesson(topword)
	}

	// If there are a bunch of 1-count problem words then just drill those
	if topcount == 1 && len(PWords) > 5 {
		return DrillOutProblemWords()
	}

	return E200FocusLesson(topword)
}

// Returns just the problem words, repeated 3 times.  This gives you a chance to clear out the problem words buffer quicker
// and advance faster.
func DrillOutProblemWords() string {
	var selection = ""
	for i := 0; i < 3; i++ {
		for word, _ := range PWords {
			selection += word + " "
		}
	}

	return strings.TrimSpace(selection)
}

// Returns an intense focus lesson concentrating hard on the missed word.  If there is a preword it's repeated with the typo
func E200IntenseFocusLesson(focusword string) string {
	var words = strings.Split(WordList(), " ")
	var selection = fmt.Sprintf("%s %s %s %s ", focusword, focusword, focusword, focusword)

	if w, ok := PreWords[focusword]; ok {
		for i := 0; i < 12; i++ {
			selection += fmt.Sprintf("%s %s ", w, focusword)
		}
	} else {
		for i := 0; i < 12; i++ {
			selection += words[rnd.Intn(len(words)-1)] + " " + focusword
		}
	}

	return strings.TrimSpace(selection)
}

// Returns a focus lesson concentrating on a given problem word.  An emptry string will select random words from E200.
func E200FocusLesson(focusword string) string {
	var selection = ""
	var words = strings.Split(WordList(), " ")

	if focusword == "" {
		for i := 0; i <= 24; i++ {
			selection += words[rnd.Intn(len(words)-1)]
			if i < 24 { // Don't insert a final space
				selection += " "
			}
		}
	} else {
		for i := 0; i <= 12; i++ {
			// if there is a preword then interleave with the problem word
			if pre, ok := PreWords[focusword]; ok {
				if i%2 == 0 {
					selection += pre + " " + focusword + " "
				} else {
					selection += words[rnd.Intn(len(words)-1)] + " " + focusword + " "
				}
			} else {
				selection += words[rnd.Intn(len(words)-1)] + " " + focusword + " "
			}
		}
	}

	return strings.TrimSpace(selection)
}
