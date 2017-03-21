package service;

import (
    "github.com/iunsuccessful/tools/words/fao/dictionary"
    "github.com/iunsuccessful/tools/words/module"
)

func Run(config *module.Config) error {
    words := dictionary.SelectRandom(config.Number);
    if config.Mode == 0 {
        chineseToEnglish(words, config.Itration);
    } else {
        englishToChinese(words);
    }
    return nil;
}
