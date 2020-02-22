package model

import (
    "log"
    "time"

    "github.com/lht102/go-chi-mysql-test-server/common"
)

type Object struct {
    Timestamp time.Time `json:"timestamp"`
    Key string `json:"key"`
	Value string `json:"value"`
}

func GetAllObjects() ([]Object, error) {
    var objs []Object
    rows, err := common.DB.Query("select * from Obj order by updated desc;")
    if err != nil {
        return objs, err
    }
    defer rows.Close()
    for rows.Next() {
        var obj Object
        if err := rows.Scan(&obj.Key, &obj.Value, &obj.Timestamp); err != nil {
            log.Println(err)
        } else {
            objs = append(objs, obj)
        }
    }
    return objs, err
}

func CreateObject(obj *Object) error {
    stmt, _ := common.DB.Prepare("insert into Obj(id, val, updated) values (?, ?, NOW())")
    defer stmt.Close()
    _, err := stmt.Exec(obj.Key, obj.Value)
    return err;
}
