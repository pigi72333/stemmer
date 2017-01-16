# Yet another implementation of the Porter Stemmer algorithm 

## Stemmig:
Stemming is an algorithm of finding a word’s stem for a given source word. Word’s stem may not be equal to morphological word’s root. The algorithm doesn’t use words database, but uses some rules step-wise, cutting off endings and suffixes according to language features. As a result it works fat but not always accurate.

## Porter Stemmer algorithm:
https://tartarus.org/martin/PorterStemmer/def.txt
## Difference from the published algorithm:
https://tartarus.org/martin/PorterStemmer/
## Usage:

```
package main

import (
  "fmt"
  "github.com/pigi72333/stemmer"
)

func main() {
  word := []byte("probate")
  stem := stemmer.Stem(word)
  fmt.Println(stem)
}
```