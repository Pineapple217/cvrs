package config

import "fmt"

type Database struct {
	SqliteOptions string `yaml:"sqliteOptions"`
	DataLocation  string `yaml:"dataLocation"`
}

func (c *Database) SetDefault() {
	c.DataLocation = "./data"
	c.SqliteOptions = "file:%s/database.db?_fk=1&_journal_mode=WAL"
}

func (c *Database) Validate() {
	c.SqliteOptions = fmt.Sprintf(c.SqliteOptions, c.DataLocation)
}
