package dao;

import (
    "log"
    "container/list"
    "github.com/iunsuccessful/words/module"
)

func SelectCount() int {
    db := getDB();
    count := 0;
	err := db.QueryRow("SELECT count(*) FROM dictionary").Scan(&count);
    if err != nil {
        log.Fatal(err);
    }
    return count;
}

func SelectRandom(limit int) *list.List {
    db := getDB();
    rows, err := db.Query("select id, spell, interpretation from dictionary order by random() limit $1", limit);
    if err != nil {
        log.Fatal(err);
    }
    defer rows.Close();
    result := list.New();
    for rows.Next() {
        dict := &module.Dictionary{};
        if err := rows.Scan(&dict.Id, &dict.Spell, &dict.Interpretation); err != nil {
            log.Fatal(err);
        }
        result.PushBack(dict);
    }
    return result;
}

