package config

type Database struct {
	SqliteOptions string `yaml:"SqliteOptions"`
}

func (c *Database) SetDefault() {
	c.SqliteOptions = "file:./data/database.db?_fk=1&_journal_mode=WAL"
}
