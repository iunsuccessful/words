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
    "github.com/iunsuccessful/tools/words/module"
    // "github.com/iunsuccessful/common/array"
)
//
// func SelectCount() int {
//     return 10;
// }
//

var wordbook Wordbook;

func init() {
    // absPath, _ := filepath.Abs("./words.xml");
    absPath, err := XmlPath();
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

func SelectRandom(limit int) *list.List {
    limit = min(limit, len(wordbook.Items));
    result := list.New();
    Shuffle(wordbook.Items);
    for i := 0; i < limit; i++ {
        item := wordbook.Items[i];
        dict := &module.Dictionary{ Spell: item.Word, Interpretation: item.Trans };
        result.PushBack(dict);
    }
    return result;
}

func Shuffle(a []Item) {
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

