package stemmer

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

type word []byte

func TestFirstA(t *testing.T) {
	fixtures := []word{
		[]byte("caresses"),
		[]byte("ponies"),
		[]byte("ties"),
		[]byte("caress"),
		[]byte("cats"),
	}

	stemmed := []word{
		[]byte("caress"),
		[]byte("poni"),
		[]byte("ti"),
		[]byte("caress"),
		[]byte("cat"),
	}

	for k, value := range fixtures {
		if result := firstA(value); !bytes.Equal(result, stemmed[k]) {
			t.Errorf("firstA() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestFirstB(t *testing.T) {
	fixtures := []word{
		[]byte("feed"),
		[]byte("agreed"),
		[]byte("plastered"),
		[]byte("bled"),
		[]byte("motoring"),
		[]byte("sing"),
	}

	stemmed := []word{
		[]byte("feed"),
		[]byte("agree"),
		[]byte("plaster"),
		[]byte("bled"),
		[]byte("motor"),
		[]byte("sing"),
	}

	for k, value := range fixtures {
		if result := firstB([]byte(value)); !bytes.Equal(result, []byte(stemmed[k])) {
			t.Errorf("firstB() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestFirstC(t *testing.T) {
	fixtures := []word{
		[]byte("happy"),
		[]byte("sky"),
	}

	stemmed := []word{
		[]byte("happi"),
		[]byte("sky"),
	}

	for k, value := range fixtures {
		if result := firstC([]byte(value)); !bytes.Equal(result, []byte(stemmed[k])) {
			t.Errorf("firstC() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestSecond(t *testing.T) {
	fixtures := []word{
		[]byte("relational"),
		[]byte("conditional"),
		[]byte("rational"),
		[]byte("valenci"),
		[]byte("hesitanci"),
		[]byte("digitizer"),
		[]byte("conformabli"),
		[]byte("radicalli"),
		[]byte("differentli"),
		[]byte("vileli"),
		[]byte("analogousli"),
		[]byte("vietnamization"),
		[]byte("predication"),
		[]byte("operator"),
		[]byte("feudalism"),
		[]byte("decisiveness"),
		[]byte("hopefulness"),
		[]byte("callousness"),
		[]byte("formaliti"),
		[]byte("sensitiviti"),
		[]byte("sensibiliti"),
	}

	stemmed := []word{
		[]byte("relate"),
		[]byte("condition"),
		[]byte("rational"),
		[]byte("valence"),
		[]byte("hesitance"),
		[]byte("digitize"),
		[]byte("conformable"),
		[]byte("radical"),
		[]byte("different"),
		[]byte("vile"),
		[]byte("analogous"),
		[]byte("vietnamize"),
		[]byte("predicate"),
		[]byte("operate"),
		[]byte("feudal"),
		[]byte("decisive"),
		[]byte("hopeful"),
		[]byte("callous"),
		[]byte("formal"),
		[]byte("sensitive"),
		[]byte("sensible"),
	}

	for k, value := range fixtures {
		if result := second([]byte(value)); !bytes.Equal(result, []byte(stemmed[k])) {
			t.Errorf("second() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}

}

func TestThird(t *testing.T) {
	fixtures := []word{
		[]byte("triplicate"),
		[]byte("formative"),
		[]byte("formalize"),
		[]byte("electriciti"),
		[]byte("electrical"),
		[]byte("hopeful"),
		[]byte("goodness"),
	}

	stemmed := []word{
		[]byte("triplic"),
		[]byte("form"),
		[]byte("formal"),
		[]byte("electric"),
		[]byte("electric"),
		[]byte("hope"),
		[]byte("good"),
	}

	for k, value := range fixtures {
		if result := third([]byte(value)); !bytes.Equal(result, []byte(stemmed[k])) {
			t.Errorf("third() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestFour(t *testing.T) {
	fixtures := []word{
		[]byte("revival"),
		[]byte("allowance"),
		[]byte("inference"),
		[]byte("airliner"),
		[]byte("gyroscopic"),
		[]byte("adjustable"),
		[]byte("defensible"),
		[]byte("irritant"),
		[]byte("replacement"),
		[]byte("adjustment"),
		[]byte("dependent"),
		[]byte("adoption"),
		[]byte("homologou"),
		[]byte("communism"),
		[]byte("activate"),
		[]byte("angulariti"),
		[]byte("homologous"),
		[]byte("effective"),
		[]byte("bowdlerize"),
	}

	stemmed := []word{
		[]byte("reviv"),
		[]byte("allow"),
		[]byte("infer"),
		[]byte("airlin"),
		[]byte("gyroscop"),
		[]byte("adjust"),
		[]byte("defens"),
		[]byte("irrit"),
		[]byte("replac"),
		[]byte("adjust"),
		[]byte("depend"),
		[]byte("adopt"),
		[]byte("homolog"),
		[]byte("commun"),
		[]byte("activ"),
		[]byte("angular"),
		[]byte("homolog"),
		[]byte("effect"),
		[]byte("bowdler"),
	}

	for k, value := range fixtures {
		if result := four([]byte(value)); !bytes.Equal(result, []byte(stemmed[k])) {
			t.Errorf("four() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestFiveA(t *testing.T) {
	fixtures := []word{
		[]byte("probate"),
		[]byte("rate"),
		[]byte("cease"),
	}

	stemmed := []word{
		[]byte("probat"),
		[]byte("rate"),
		[]byte("ceas"),
	}

	for k, value := range fixtures {
		if result := fiveA(value); !bytes.Equal(result, stemmed[k]) {
			t.Errorf("fiveA() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestFiveB(t *testing.T) {
	fixtures := []word{
		[]byte("controll"),
		[]byte("roll"),
	}

	stemmed := []word{
		[]byte("control"),
		[]byte("roll"),
	}

	for k, value := range fixtures {
		if result := fiveB(value); !bytes.Equal(result, stemmed[k]) {
			t.Errorf("fiveB() return value not what was expected, pass: '%v' return: '%v' expected: '%v'", value, result, stemmed[k])
		}
	}
}

func TestConsonant(t *testing.T) {
	word := []byte("ty")
	if condition := consonant(word, 1); condition != false {
		t.Errorf("consonant() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, condition, false)
	}

	word = []byte("x")
	if condition := consonant(word, 0); condition != true {
		t.Errorf("consonant() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, condition, false)
	}

	word = []byte("ay")
	if condition := consonant(word, 1); condition != true {
		t.Errorf("consonant() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, condition, true)
	}
}

func TestVowel(t *testing.T) {
	word := []byte("ty")
	if condition := vowel(word, 1); condition == false {
		t.Errorf("vowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, condition, true)
	}

	word = []byte("ay")
	if condition := vowel(word, 1); condition != false {
		t.Errorf("vowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, condition, false)
	}
}

func TestMeasure(t *testing.T) {
	m := 0
	word := []byte("tr")
	if m = measure(word); m != 0 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 0)
	}

	word = []byte("ee")
	if m = measure(word); m != 0 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 0)
	}

	word = []byte("y")
	if m = measure(word); m != 0 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 0)
	}

	word = []byte("by")
	if m = measure(word); m != 0 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 0)
	}

	word = []byte("tree")
	if m = measure(word); m != 0 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 0)
	}

	word = []byte("trouble")
	if m = measure(word); m != 1 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 1)
	}

	word = []byte("oats")
	if m = measure(word); m != 1 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 1)
	}

	word = []byte("trees")
	if m = measure(word); m != 1 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 1)
	}

	word = []byte("ivy")
	if m = measure(word); m != 1 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 1)
	}

	word = []byte("troubles")
	if m = measure(word); m != 2 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 2)
	}

	word = []byte("private")
	if m = measure(word); m != 2 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 2)
	}

	word = []byte("oaten")
	if m = measure(word); m != 2 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 2)
	}

	word = []byte("orrery")
	if m = measure(word); m != 2 {
		t.Errorf("measure() return value not what was expected, pass: '%s' return: '%d' expected: '%d'", word, m, 2)
	}
}

func TestContainsVowel(t *testing.T) {
	word := []byte("golang")
	if condition := containVowel(word); condition != true {
		t.Errorf("containVowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, false, true)
	}

	word = []byte("yb")
	if condition := containVowel(word); condition != false {
		t.Errorf("containVowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, true, false)
	}
}

func TestStem(t *testing.T) {
	word := []byte("abbreviated")
	stem := []byte("abbrevi")

	if result := Stem(word); !bytes.Equal(result, stem) {
		t.Errorf("Stem() return value not what was expected, pass: '%s' return: '%s' expected: '%s'", word, result, stem)
	}
}

func TestisCVCSuffix(t *testing.T) {
	word := []byte("tyt")
	if condition := isCVCSuffix(word); condition == false {
		t.Errorf("vowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, false, true)
	}

	word = []byte("wil")
	if condition := isCVCSuffix(word); condition == true {
		t.Errorf("vowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, true, false)
	}

	word = []byte("hop")
	if condition := isCVCSuffix(word); condition == true {
		t.Errorf("vowel() return value not what was expected, pass: '%s' return: '%v' expected: '%v'", word, true, false)
	}
}

func TestVocal(t *testing.T) {

	v, err := os.Open("voc.txt")
	if err != nil {
		panic(err)
	}
	defer v.Close()
	vocScanner := bufio.NewScanner(v)

	o, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	defer o.Close()
	outScanner := bufio.NewScanner(o)

	for vocScanner.Scan() {
		outScanner.Scan()
		word := vocScanner.Bytes()
		stem := outScanner.Bytes()

		if result := Stem(word); !bytes.Equal(result, stem) {
			t.Errorf("Stem() return value not what was expected, pass: '%s' return: '%s' expected: '%s'", word, result, stem)
		}
	}
}

func BenchmarkStem(b *testing.B) {
	word := []byte("troubles")
	for n := 0; n < b.N; n++ {
		Stem(word)
	}
}

func BenchmarkFirstA(b *testing.B) {
	word := []byte("caresses")
	for n := 0; n < b.N; n++ {
		firstB(word)
	}
}

func BenchmarkFirstB(b *testing.B) {
	word := []byte("feed")
	for n := 0; n < b.N; n++ {
		firstB(word)
	}
}

func BenchmarkFirstC(b *testing.B) {
	word := []byte("happy")
	for n := 0; n < b.N; n++ {
		firstC(word)
	}
}

func BenchmarkSecond(b *testing.B) {
	word := []byte("vietnamization")
	for n := 0; n < b.N; n++ {
		second(word)
	}
}

func BenchmarkThird(b *testing.B) {
	word := []byte("electriciti")
	for n := 0; n < b.N; n++ {
		third(word)
	}
}

func BenchmarkFour(b *testing.B) {
	word := []byte("allowance")
	for n := 0; n < b.N; n++ {
		four(word)
	}
}

func BenchmarkFiveA(b *testing.B) {
	word := []byte("probate")
	for n := 0; n < b.N; n++ {
		firstA(word)
	}
}

func BenchmarkFiveB(b *testing.B) {
	word := []byte("controll")
	for n := 0; n < b.N; n++ {
		firstB(word)
	}
}
