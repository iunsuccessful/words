package dictionary;

import (
    "os/exec"
    "path/filepath"
)

func XmlPath(bookName string) (string, error) {
    // 可执行文件所在目录目录
    path, err :=  exec.LookPath("words");
    if err != nil {
        return "", err;
    }
    dir := filepath.Dir(path);
    return filepath.Join(dir ,"dict", bookName + ".xml"), nil;
}


