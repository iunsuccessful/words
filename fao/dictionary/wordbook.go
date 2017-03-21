package dictionary;

type Wordbook struct {
    Items []Item `xml:"item"`
}

type Item struct {
    Word      string `xml:"word"`
    Trans     string `xml:"trans"`
    Phonetic  string `xml:"phonetic"`
    Tags      string `xml:"tags"`
    Progress  int `xml:"progress"`
}

