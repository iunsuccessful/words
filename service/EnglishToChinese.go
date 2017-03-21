package service;

import (
    "container/list"
    "fmt"
    "github.com/iunsuccessful/tools/words/module"
)

func englishToChinese(words *list.List) {
    score := 0;
    errs := list.New();
    for e := words.Front(); e != nil; e = e.Next() {
        if dict, ok := e.Value.(*module.Dictionary); ok {
            fmt.Println("-------------------------------");
            fmt.Println(dict.Spell);
            var answer string;
            // 这个处理空格会出问题
            fmt.Scanln(&answer);
            fmt.Printf("%s\n", dict.Interpretation);
            fmt.Scanln(&answer);
            if len(answer) == 0 {
                score++;
            } else {
                errs.PushBack(dict);
            }
        }
    }
    fmt.Printf("Total: %d score: %d\n", words.Len(), score);
    if errs.Len() > 0 {
        englishToChinese(errs);
    }
}

