package dictionary;

import (
    "os/exec"
    "path/filepath"
)

func XmlPath() (string, error) {
    path, err :=  exec.LookPath("words");
    if err != nil {
        return "", err;
    }
    dir := filepath.Dir(path);
    return filepath.Join(dir , "words.xml"), nil;
}


