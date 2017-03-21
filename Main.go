package main;

import (
    "os"
    "github.com/iunsuccessful/tools/words/service"
    "github.com/iunsuccessful/tools/words/module"
)

func main() {
    // service of words
    service.Run(module.ParseConfig(os.Args));
}
