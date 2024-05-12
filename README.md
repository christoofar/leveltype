# leveltype
Improve your typing speed by focusing on spacegrams

## How leveltype works

leveltype is a typing trainer that puts its focus on *spacegrams* and forcing you to "unlearn" bad muscle memory. Spacegrams are the final two letters of a word, followed by a space, followed by the beginning two or three letters of the next word.

For instance, in these two words:
```
potato farmer
```
the spacegram is the sequence `to_far`

Spacegrams can be thought of as their own words.  In the lowercase Latin characterset (and accounting for the apostrophe), there are a staggering 11,142,400 possible spacegrams!

Luckily only a small portion of possible spacegrams regularly occur in English.  As you make typing errors, the word that you typed correctly before the typo is brought back for you to retype.   As you perfect your muscle memory on a problem word, the "entry word" that came before you made a typo as well as a selection of random words are given to you so that you learn the spacegram as well as the word that you mistyped.

Perfecting the spacegrams that lead into your typos is how you gain more "flow" as you type.

## Using Leveltype to "clean" your typing muscle memory and improve accuracy

Oftentimes novice typists accumulate "garbage" muscle memory where a typo, the backspace key, and the correction are a complete sequence memorized when typing a word or a portion thereof.

Leveltype deactivates the `Backspace` key and you are not allowed to correct your typing mistakes in a typing session.  This forces you to learn the keystrokes 'cleanly', without the use of the Backspace key.  

This appoximates the experience that manual typewriter users used to gain muscle memory when training and to achieve very high accuracy levels.

You can skip a word that you have typo'd simply by pressing the spacebar and moving on to the next word.

## Building Source

Install the [Go SDK][go sdk]

```
leveltype/cmd #  go run main.go
```

## TODOs
- Makefile
- Packaging CI/CD

[go sdk]: https://go.dev/dl/