package main;

import (
    "os"
    "github.com/iunsuccessful/words/service"
    "github.com/iunsuccessful/words/module"
)

func main() {
    // service of words
    service.Run(module.ParseConfig(os.Args));
}
