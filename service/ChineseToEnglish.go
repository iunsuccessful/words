package service;

import (
    "fmt"
    "github.com/iunsuccessful/tools/words/module"
    "math/rand"
    "strings"
    "bufio"
    "os"
    "container/list"
)

/**
 * 20161117 增加注释中字母隐藏
 */
func chineseToEnglish(words *list.List, itration int) {
    score := 0;
    errs := list.New();
    for e := words.Front(); e != nil; e = e.Next() {
        if dict, ok := e.Value.(*module.Dictionary); ok {
            hintWord := hintSpell(dict.Spell, itration);
            hintInterpretation := strings.Replace(strings.ToLower(dict.Interpretation), dict.Spell, hintWord, -1);
            fmt.Println(hintInterpretation);
            fmt.Printf("提示：%s\n", hintWord);
            var one = answer(dict.Spell);
            score += one;
            if one <= 0 {
                errs.PushBack(dict);
            }
        }
    }
    fmt.Printf("Total: %d score: %d\n", words.Len(), score);
    if errs.Len() > 0 {
        chineseToEnglish(errs, itration);
    }
}

func hintSpell(spell string, itration int) string {
    for i := 0; i < itration; i++ {
        spell = hint(spell);
    }
    return spell;
}

func hint(spell string) string {
    i := rand.Intn(len(spell));
    return strings.Replace(spell, spell[i:i+1], "_", 1);
}

var inputReader *bufio.Reader

func answer(standrad string) int {
    var answer string;
    score := 2; // 每题一分
    for strings.TrimSpace(answer) != standrad {
        // score -= 1; // 答一次，扣一分
        // 如果 0 分了，就给个提示吧
        if score -= 1; score <= 0 {
            fmt.Printf("Error, Hint: %s\n", standrad);
        }
        inputReader = bufio.NewReader(os.Stdin);
        answer, _ = inputReader.ReadString('\n');
        // fmt.Scanln(&answer); // 这个处理空格会出问题
    }
    return score;
}
