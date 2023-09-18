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
    m2w map[string] []string
    w2m map[string] []string
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

    return
}

func (s *Synonyms) GetLen() int {
    return len(s.S)
}

func (s *Synonyms) MakeMap {
    for _, item := range s.S {
    }
}
