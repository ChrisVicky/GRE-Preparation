package synonyms

import "testing"

func TestSynonyms(t *testing.T){

    if s, err := ReadSynonyms(""); err != nil {
        t.Errorf("Error! : %v", err)
    }else{
        t.Logf("%+v", s)
    }
}
