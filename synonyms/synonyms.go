package synonyms

import (
    "log"
    "os"
    "github.com/BurntSushi/toml"
)

type Synonym struct{
    Meaning     string      `toml:"meaning"`
    Words       []string    `toml:"words"` 
}

type Synonyms struct {
    S []Synonym  `toml:"synonyms"`
    M2w map[string] []string
    W2m map[string] []string
    Words []string
}

var DefaultFile = "words.toml"

func ReadSynonyms(fn string) (s *Synonyms, err error){
    s = &Synonyms{}
    var (
        cBytes  []byte
    )

    if _, err := os.Stat(fn); err != nil {
        log.Printf("File '%v' cannot be stat, Use Default File: '%v'", fn, DefaultFile)
        fn = DefaultFile
    }

    if cBytes, err = os.ReadFile(fn); err != nil {
        return
    }

    if err = toml.Unmarshal(cBytes, s); err != nil {
        return
    }

    s.prepare()
    
    return
}

func (s *Synonyms) prepare(){
    s.W2m = make(map[string][]string)
    s.M2w = make(map[string][]string)
    for _, mw := range s.S {
        for _, w := range mw.Words {
            s.W2m[w] = append(s.W2m[w], mw.Meaning)
        }
        s.M2w[mw.Meaning] = mw.Words
    }
    for k := range s.W2m {
        s.Words = append(s.Words, k)
    }
}

func (s *Synonyms) GetLen() int {
    return len(s.S)
}

func (s *Synonyms) GetWordsNum() int {
    return len(s.Words)
}
