# leveltype
Improve your typing speed by focusing on spacegrams

## How leveltype works

leveltype is a typing trainer that puts its focus on *spacegrams* and forcing you to "unlearn" bad muscle memory. Spacegrams are the final two letters of a word, followed by a space, followed by the beginning two or three letters of the next word.

For instance, in these two words:
```
potato farmer
```
the spacegram is the sequence `to_far`

Spacegrams can be thought of as their own words.  In the lowercase Latin characterset (and accounting for the apostrophe), there are a staggering 234,375,000 possible spacegrams!

Luckily only a small portion of possible spacegrams regularly occur in English.  As you make typing errors, the word that you typed correctly before the typo is brought back for you to retype.   As you perfect your muscle memory on the problem word that you made a mistake, the "entry word" coming into the typo as well as selection of random words are given to you.

Perfecting the spacegrams that lead into your typos is how you gain more "flow" as you type.

## Using Leveltype to "clean" your typing muscle memory and improve accuracy

Oftentimes novice typists accumulate "garbage" muscle memory where a typo, the backspace key, and the correction are memorized when typing a word.

Leveltype deactivates the `Backspace` key and you are not allowed to correct your typing mistakes in a typing session.  This forces you to learn the keystrokes 'cleanly' without the use of the Backspace key.  

This appoximates the experience that manual typewriter users used to gain muscle memory when training and to achieve very high accuracy levels.

## Building Source

Install the [Go SDK][go sdk]

```
leveltype/cmd #  go run main.go
```

## TODOs
- Makefile
- Packaging CI/CD

[go sdk]: https://go.dev/dl/