package module;

import (
    "strconv"
    "log"
    "fmt"
    "os"
)

type Config struct {
    Number   int // 本次显示多少个单词
    Itration int // 隐藏多少位 -1 为不提示
    Mode     int // 0. chinese to english(default) 1. english to chinese
    BookName string // 单词书

}

func ParseConfig(args []string) *Config {
    config := &Config{ Number: 10, Itration: -1, Mode: 0 };
    for i := 1; i < len(args); i++ {
        if (args[i][0] == '-') || (args[i][0] == '/') {
            switch args[i][1] {
				case 'h': usage();
				case 'n': config.Number = getToInt(args, i+1);
				case 'i': config.Itration = getToInt(args, i+1);
				case 'm': config.Mode = getToInt(args, i+1);
            }
            i++
        } else if len(config.BookName) == 0 {
			config.BookName = args[i]
		}
    }

	if len(config.BookName) == 0 {
		// print help
		usage()
	}

    return config;
}

func getToInt(args []string, index int) int {
    if index >= len(args) {
        log.Fatalf("Array out of range %d %d", index, len(args));
    }
    val, err := strconv.Atoi(args[index]);
    if err!= nil {
        log.Fatalf("convert %s to int error", args[index]);
    }
    return val;
}

func usage() {
    fmt.Printf("Use Words\n\n");
    fmt.Printf("Usage:\n\t words [-n number] [-i hint] [-m mode] <单词本>\n\n");
    fmt.Printf("Options:\n");
    fmt.Printf("\t-n\tChoose Number to this test.\n");
    fmt.Printf("\t-i\tHint Chat in mode 0.\n");
    fmt.Printf("\t-m\t0. chinese to english(default) 1. english to chinese\n");
    fmt.Println();
    os.Exit(0)
}
