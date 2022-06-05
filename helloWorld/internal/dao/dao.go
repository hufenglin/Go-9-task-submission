package dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v3"
	"helloWorld/internal/model"
	"io/ioutil"
	"os"
)

var D dao

type dao struct {
	db          *gorm.DB
}

type UserDB interface {
	QueryUserList(level int) (err error, userList *[]model.User)
}

//profile variables
type conf struct {
	Data data `yaml:"data"`
}

type data struct {
	Database database `yaml:"database"`
}

type database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

func (c *conf) getConf() (*conf, error) {
	dir, _ := os.Getwd()
	fmt.Println("当前路径：",dir)

	yamlFile, err := ioutil.ReadFile(dir + "/configs/config.yaml")
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (D *dao) QueryUserList(level int32) (userList *[]model.User, count int32, err error) {
	var c conf
	dbConfig, err := c.getConf()
	if err != nil {
		return
	}

	list := make([]model.User, 0)
	db, err := sql.Open(dbConfig.Data.Database.Driver, dbConfig.Data.Database.Source)
	if err != nil {
		return
	}

	var selectSql string
	if level == 0 {
		selectSql = "select id, name, class, level from User where leve = ?"
	} else {
		selectSql = "select id, name, class, level from User where level = ?"
	}

	rows, err := db.Query(selectSql, level)
	if err != nil {
		//sql.ErrNoRows，则返回没有数据
		if errors.Is(err, sql.ErrNoRows) {
			return &list, 0, nil
		} else {
			//其他错误，则包装返回
			return nil, 0, fmt.Errorf("Query User List Failed: %w", err)
		}
	}

	for rows.Next() {
		var name string
		var id, class, level int32
		_ = rows.Scan(&id, &name, &class, &level)
		list = append(list, model.User{ID: id, Name: name, Class: class, Level: level})
	}

	userList = &list
	count = int32(len(list))
	return
}