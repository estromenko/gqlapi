package database

import (
	"io/ioutil"
)

// Migrate ...
func (d *Database) Migrate() error {

	var migrationsPath string = "./database/migrations"

	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		d.logger.Error(err.Error())
		return err
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(migrationsPath + "/" + file.Name())

		if err != nil {
			d.logger.Error(err.Error())
			return err
		}

		if _, err := d.db.DB().Exec(string(data)); err != nil {
			d.logger.Error(err.Error())
			return err
		}

		d.logger.Info(file.Name() + ": done.")
	}

	return nil
}
