package stemmer

import "bytes"

const (
	vowel_state = iota
	consonant_state
)

func measure(word []byte) int {
	m := 0
	if len(word) > 0 {
		var state int = consonant_state
		if vowel(word, 0) {
			state = vowel_state
		}
		for i := 0; i < len(word); i++ {
			if vowel(word, i) && state == consonant_state {
				state = vowel_state
			} else if consonant(word, i) && state == vowel_state {
				state = consonant_state
				m++
			}
		}
	}
	return m
}

//
// A \consonant\ in a word is a letter other than A, E, I, O or U, and other
// than Y preceded by a consonant.
//
func consonant(word []byte, i int) bool {
	switch word[i] {
	case byte('a'), byte('e'), byte('i'), byte('o'), byte('u'):
		return false
	case byte('y'):
		if 0 == i {
			return true
		} else {
			return (i > 0 && !consonant(word, i-1))
		}
	default:
		return true
	}
}

func vowel(word []byte, i int) bool {
	return !consonant(word, i)
}

//
// *v* - the stem contains a vowel
//
func containVowel(word []byte) bool {
	for i, _ := range word {
		if vowel(word, i) {
			return true
		}
	}
	return false
}

//
// Step 1a
//
//    SSES -> SS                         caresses  ->  caress
//    IES  -> I                          ponies    ->  poni
//                                       ties      ->  ti
//    SS   -> SS                         caress    ->  caress
//    S    ->                            cats      ->  cat
//
func firstA(word []byte) []byte {
	if bytes.HasSuffix(word, []byte("s")) {
		if bytes.HasSuffix(word, []byte("sses")) || bytes.HasSuffix(word, []byte("ies")) {
			return word[:len(word)-2]
		} else if bytes.HasSuffix(word, []byte("ss")) {
			return word
		}
		return word[:len(word)-1]
	}
	return word
}

//
// Step 1b
//
//    (m>0) EED -> EE                    feed      ->  feed
//                                       agreed    ->  agree
//    (*v*) ED  ->                       plastered ->  plaster
//                                       bled      ->  bled
//    (*v*) ING ->                       motoring  ->  motor
//                                       sing      ->  sing
//
func firstB(word []byte) []byte {
	l := len(word)

	if bytes.HasSuffix(word, []byte("ed")) {
		if bytes.HasSuffix(word, []byte("eed")) {
			if m := measure(word[:l-3]); m > 0 {
				return bytes.TrimSuffix(word, []byte("d"))
			}
		} else {
			if containVowel(word[:l-2]) {
				return firstB2(bytes.TrimSuffix(word, []byte("ed")))
			}
		}

	} else if bytes.HasSuffix(word, []byte("ing")) {
		if containVowel(word[:l-3]) {
			return firstB2(bytes.TrimSuffix(word, []byte("ing")))
		}
	}
	return word
}

//
//  AT -> ATE                       conflat(ed)  ->  conflate
//  BL -> BLE                       troubl(ed)   ->  trouble
//  IZ -> IZE                       siz(ed)      ->  size
//  (*d and not (*L or *S or *Z))
//     -> single letter
//                                  hopp(ing)    ->  hop
//                                  tann(ed)     ->  tan
//                                  fall(ing)    ->  fall
//                                  hiss(ing)    ->  hiss
//                                  fizz(ed)     ->  fizz
//  (m=1 and *o) -> E               fail(ing)    ->  fail
//                                  fil(ing)     ->  file
//
func firstB2(word []byte) []byte {
	l := len(word)
	if bytes.HasSuffix(word, []byte("at")) || bytes.HasSuffix(word, []byte("iz")) || bytes.HasSuffix(word, []byte("bl")) {
		return append(word, 'e')
		// (*d and not (*L or *S or *Z)) -> single letter
	} else if consonant(word, l-1) && word[l-1] == word[l-2] {
		if !bytes.HasSuffix(word, []byte("l")) && !bytes.HasSuffix(word, []byte("s")) && !bytes.HasSuffix(word, []byte("z")) {
			return word[:l-1]
		}
	} else if m := measure(word); m == 1 {
		//*o  - the stem ends cvc, where the second c is not W, X or Y (e.g. -WIL, -HOP).
		//(m=1 and *o) -> E
		if isCVCSuffix(word) {
			return append(word, 'e')
		}
	}
	return word
}

//
//*o  - the stem ends cvc, where the second c is not W, X or Y (e.g.
//       -WIL, -HOP)
//
func isCVCSuffix(word []byte) bool {
	size := len(word) - 1
	if size >= 2 && consonant(word, size-2) && vowel(word, size-1) && consonant(word, size) && word[size] != 'w' && word[size] != 'x' && word[size] != 'y' {
		return true
	}
	return false
}

//
// Step 1c
//
//    (*v*) Y -> I                    happy        ->  happi
//                                    sky          ->  sky
//
func firstC(word []byte) []byte {
	l := len(word)
	if bytes.HasSuffix(word, []byte("y")) && containVowel(word[:l-1]) {
		word[l-1] = byte('i')
		return word
	}
	return word
}

//
// Step 2
//
//    (m>0) ATIONAL ->  ATE           relational     ->  relate
//    (m>0) TIONAL  ->  TION          conditional    ->  condition
//                                    rational       ->  rational
//    (m>0) ENCI    ->  ENCE          valenci        ->  valence
//    (m>0) ANCI    ->  ANCE          hesitanci      ->  hesitance
//    (m>0) IZER    ->  IZE           digitizer      ->  digitize
//    (m>0) ABLI    ->  ABLE          conformabli    ->  conformable
//    (m>0) ALLI    ->  AL            radicalli      ->  radical
//    (m>0) ENTLI   ->  ENT           differentli    ->  different
//    (m>0) ELI     ->  E             vileli        - >  vile
//    (m>0) OUSLI   ->  OUS           analogousli    ->  analogous
//    (m>0) IZATION ->  IZE           vietnamization ->  vietnamize
//    (m>0) ATION   ->  ATE           predication    ->  predicate
//    (m>0) ATOR    ->  ATE           operator       ->  operate
//    (m>0) ALISM   ->  AL            feudalism      ->  feudal
//    (m>0) IVENESS ->  IVE           decisiveness   ->  decisive
//    (m>0) FULNESS ->  FUL           hopefulness    ->  hopeful
//    (m>0) OUSNESS ->  OUS           callousness    ->  callous
//    (m>0) ALITI   ->  AL            formaliti      ->  formal
//    (m>0) IVITI   ->  IVE           sensitiviti    ->  sensitive
//    (m>0) BILITI  ->  BLE           sensibiliti    ->  sensible
//
func second(word []byte) []byte {
	l := len(word)
	if bytes.HasSuffix(word, []byte("tional")) {
		if bytes.HasSuffix(word, []byte("ational")) {
			if m := measure(word[:l-7]); m > 0 {
				return append(bytes.TrimSuffix(word, []byte("ional")), 'e')
			}
		} else if m := measure(word[:l-2]); m > 0 {
			return bytes.TrimSuffix(word, []byte("al"))
		}
	} else if bytes.HasSuffix(word, []byte("nci")) {
		if bytes.HasSuffix(word, []byte("enci")) || bytes.HasSuffix(word, []byte("anci")) {
			if m := measure(word[:l-4]); m > 0 {
				return append(word[:len(word)-1], 'e')
			}
		}
	} else if bytes.HasSuffix(word, []byte("li")) {
		if bytes.HasSuffix(word, []byte("abli")) {
			if m := measure(word[:l-4]); m > 0 {
				return append(word[:len(word)-1], 'e')
			}
		} else if bytes.HasSuffix(word, []byte("bli")) {
			if m := measure(word[:l-3]); m > 0 {
				return append(word[:l-1], 'e')
			}
		} else if bytes.HasSuffix(word, []byte("alli")) {
			if m := measure(word[:l-4]); m > 0 {
				return word[:l-2]
			}
		} else if bytes.HasSuffix(word, []byte("entli")) || bytes.HasSuffix(word, []byte("ousli")) {
			if m := measure(word[:l-5]); m > 0 {
				return word[:l-2]
			}
		} else if bytes.HasSuffix(word, []byte("eli")) {
			if m := measure(word[:l-3]); m > 0 {
				return word[:l-2]
			}
		}
	} else if bytes.HasSuffix(word, []byte("ti")) {
		if bytes.HasSuffix(word, []byte("iviti")) {
			if m := measure(word[:l-5]); m > 0 {
				return append(word[:l-3], 'e')
			}
		} else if bytes.HasSuffix(word, []byte("aliti")) {
			if m := measure(word[:l-5]); m > 0 {
				return word[:l-3]
			}
		} else if bytes.HasSuffix(word, []byte("biliti")) {
			if m := measure(word[:l-6]); m > 0 {
				return append(word[:l-5], []byte("le")...)
			}
		}
	} else if bytes.HasSuffix(word, []byte("ation")) {
		if bytes.HasSuffix(word, []byte("ization")) {
			if m := measure(word[:l-7]); m > 0 {
				return append(word[:len(word)-5], 'e')
			}
		} else if m := measure(word[:l-5]); m > 0 {
			return append(word[:l-3], 'e')
		}
	} else if bytes.HasSuffix(word, []byte("ness")) {
		if bytes.HasSuffix(word, []byte("iveness")) || bytes.HasSuffix(word, []byte("fulness")) || bytes.HasSuffix(word, []byte("ousness")) {
			if m := measure(word[:l-7]); m > 0 {
				return word[:l-4]
			}
		}
	} else if bytes.HasSuffix(word, []byte("izer")) {
		if m := measure(word[:l-4]); m > 0 {
			return word[:l-1]
		}
	} else if bytes.HasSuffix(word, []byte("alism")) {
		if m := measure(word[:l-5]); m > 0 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("ator")) {
		if m := measure(word[:l-4]); m > 0 {
			return append(word[:len(word)-2], 'e')
		}
	} else if bytes.HasSuffix(word, []byte("logi")) {
		if m := measure(word[:l-4]); m > 0 {
			return append(word[:l-1])
		}
	}
	return word
}

//
// Step 3
//
//    (m>0) ICATE ->  IC              triplicate     ->  triplic
//    (m>0) ATIVE ->                  formative      ->  form
//    (m>0) ALIZE ->  AL              formalize      ->  formal
//    (m>0) ICITI ->  IC              electriciti    ->  electric
//    (m>0) ICAL  ->  IC              electrical     ->  electric
//    (m>0) FUL   ->                  hopeful        ->  hope
//    (m>0) NESS  ->                  goodness       ->  good
//
func third(word []byte) []byte {
	l := len(word)
	if bytes.HasSuffix(word, []byte("e")) {
		if bytes.HasSuffix(word, []byte("icate")) {
			if measure(word[:l-5]) > 0 {
				return word[:l-3]
			}
		} else if bytes.HasSuffix(word, []byte("ative")) {
			if measure(word[:l-5]) > 0 {
				return word[:l-5]
			}
		} else if bytes.HasSuffix(word, []byte("alize")) {
			if measure(word[:l-5]) > 0 {
				return word[:l-3]
			}
		}
	} else if bytes.HasSuffix(word, []byte("iciti")) {
		if measure(word[:l-5]) > 0 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("ical")) {
		if measure(word[:l-4]) > 0 {
			return word[:l-2]
		}
	} else if bytes.HasSuffix(word, []byte("ful")) {
		if measure(word[:l-3]) > 0 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("ness")) {
		if measure(word[:l-4]) > 0 {
			return word[:l-4]
		}
	}
	return word
}

//
// Step 4
//
//    (m>1) AL    ->                  revival        ->  reviv
//    (m>1) ANCE  ->                  allowance      ->  allow
//    (m>1) ENCE  ->                  inference      ->  infer
//    (m>1) ER    ->                  airliner       ->  airlin
//    (m>1) IC    ->                  gyroscopic     ->  gyroscop
//    (m>1) ABLE  ->                  adjustable     ->  adjust
//    (m>1) IBLE  ->                  defensible     ->  defens
//    (m>1) ANT   ->                  irritant       ->  irrit
//    (m>1) EMENT ->                  replacement    ->  replac
//    (m>1) MENT  ->                  adjustment     ->  adjust
//    (m>1) ENT   ->                  dependent      ->  depend
//    (m>1 and (*S or *T)) ION ->     adoption       ->  adopt
//    (m>1) OU    ->                  homologou      ->  homolog
//    (m>1) ISM   ->                  communism      ->  commun
//    (m>1) ATE   ->                  activate       ->  activ
//    (m>1) ITI   ->                  angulariti     ->  angular
//    (m>1) OUS   ->                  homologous     ->  homolog
//    (m>1) IVE   ->                  effective      ->  effect
//    (m>1) IZE   ->                  bowdlerize     ->  bowdler
//
func four(word []byte) []byte {
	l := len(word)
	if bytes.HasSuffix(word, []byte("al")) || bytes.HasSuffix(word, []byte("er")) || bytes.HasSuffix(word, []byte("ic")) {
		if m := measure(word[:l-2]); m > 1 {
			return word[:l-2]
		}
	} else if bytes.HasSuffix(word, []byte("nce")) {
		//if word[l-4] == 'a' || word[l-4] == 'e' {
		if bytes.HasSuffix(word, []byte("ance")) || bytes.HasSuffix(word, []byte("ence")) {
			if m := measure(word[:l-4]); m > 1 {
				return word[:l-4]
			}
		}
	} else if bytes.HasSuffix(word, []byte("ble")) {
		if word[l-4] == 'a' || word[l-4] == 'i' {
			//if bytes.HasSuffix(word, []byte("able")) || bytes.HasSuffix(word, []byte("ible")) {
			if m := measure(word[:l-4]); m > 1 {
				return word[:l-4]
			}
		}
	} else if bytes.HasSuffix(word, []byte("ent")) {
		if bytes.HasSuffix(word, []byte("ement")) {
			if m := measure(word[:l-5]); m > 1 {
				return word[:l-5]
			}
		} else if bytes.HasSuffix(word, []byte("ment")) {
			if m := measure(word[:l-4]); m > 1 {
				return word[:l-4]
			}
		} else if m := measure(word[:l-3]); m > 1 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("ant")) {
		if m := measure(word[:l-3]); m > 1 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("e")) {
		if bytes.HasSuffix(word, []byte("ate")) || bytes.HasSuffix(word, []byte("ive")) || bytes.HasSuffix(word, []byte("ize")) {
			if m := measure(word[:l-3]); m > 1 {
				return word[:l-3]
			}
		}
	} else if bytes.HasSuffix(word, []byte("ism")) {
		if m := measure(word[:l-3]); m > 1 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("ous")) || bytes.HasSuffix(word, []byte("iti")) {
		if m := measure(word[:l-3]); m > 1 {
			return word[:l-3]
		}
	} else if bytes.HasSuffix(word, []byte("ou")) {
		if m := measure(word[:l-2]); m > 1 {
			return word[:l-2]
		}
	} else if bytes.HasSuffix(word, []byte("ion")) {
		// *S  - the stem ends with S (and similarly for the other letters). (m>1 and (*S or *T)) ION ->
		l := len(word)
		if measure(word[:l-3]) > 1 {
			if l > 4 && (word[l-4] == 's' || word[l-4] == 't') {
				return word[:l-3]
			}
		}
	}
	return word
}

//
// Step 5a
//
//    (m>1) E     ->                  probate        ->  probat
//                                    rate           ->  rate
//    (m=1 and not *o) E ->           cease          ->  ceas
//
func fiveA(word []byte) []byte {
	l := len(word)
	if bytes.HasSuffix(word, []byte("e")) {
		if m := measure(word[:l-1]); m > 1 {
			return word[:l-1]
		} else if m := measure(word[:l-1]); m == 1 {
			if !isCVCSuffix(word[:l-1]) {
				return word[:l-1]
			}
		}
	}
	return word
}

//
// Step 5b
//
//    (m > 1 and *d and *L) -> single letter
//                                    controll       ->  control
//                                    roll           ->  roll
//
func fiveB(word []byte) []byte {
	l := len(word)
	if measure(word) > 1 && consonant(word, l-1) && consonant(word, l-2) && word[l-1] == word[l-2] && word[l-1] == 'l' {
		return word[:l-1]
	}
	return word
}

func Stem(word []byte) []byte {
	word = bytes.TrimSpace(bytes.ToLower(word))
	if len(word) < 3 {
		return word
	}
	word = firstA(word)
	word = firstB(word)
	word = firstC(word)
	word = second(word)
	word = third(word)
	word = four(word)
	word = fiveA(word)
	word = fiveB(word)
	return word
}
