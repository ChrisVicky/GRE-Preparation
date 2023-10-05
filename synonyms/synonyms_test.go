package synonyms

import "testing"

func TestSynonyms(t *testing.T){

    if s, err := ReadSynonyms(""); err != nil {
        t.Errorf("Error! : %v", err)
    }else{
        t.Logf("Total: %v vs %v", s.GetWordsNum(), s.GetLen())
    }
}

func TestQ(t *testing.T){
    
    if s, err := ReadSynonyms(""); err != nil {
        t.Errorf("Error! : %v", err)
    }else{
        s.singalQ(0)
    }
}
