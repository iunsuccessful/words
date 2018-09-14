package dictionary;

import (
    "math/rand"
    "time"
    "log"
    "container/list"
    "encoding/xml"
    "io/ioutil"
    // "path/filepath"
)

import (
    "github.com/iunsuccessful/words/module"
    // "github.com/iunsuccessful/common/array"
)
//
// func SelectCount() int {
//     return 10;
// }
//

var wordbook Wordbook;

func loadWordbook(bookName string) {
    // absPath, _ := filepath.Abs("./words.xml");
    absPath, err := XmlPath(bookName);
    if err != nil {
        log.Fatal(err);
    }
    content, err := ioutil.ReadFile(absPath);
    if err != nil {
        log.Fatal(err);
    }
    err = xml.Unmarshal(content, &wordbook);
    if err != nil {
        log.Fatal(err);
    }
}

func SelectRandom(bookName string, limit int) *list.List {
    loadWordbook(bookName)
    limit = min(limit, len(wordbook.Items));
    result := list.New();
    shuffle(wordbook.Items);
    for i := 0; i < limit; i++ {
        item := wordbook.Items[i];
        dict := &module.Dictionary{ Spell: item.Word, Interpretation: item.Trans };
        result.PushBack(dict);
    }
    return result;
}

func shuffle(a []Item) {
    rand.Seed(time.Now().UnixNano());
    for i := range a {
        j := rand.Intn(i + 1)
        a[i], a[j] = a[j], a[i];
    }
}

func min(a, b int) int {
    if a > b {
        return b;
    }
    return a;
}

