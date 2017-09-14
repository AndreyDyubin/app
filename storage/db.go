package storage

import (
	"gopkg.in/reform.v1"
	"database/sql"
	"os"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
	"github.com/AndreyDyubin/app/models"
)

var DB *reform.DB

func ConnectDB() error {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123321 dbname=server sslmode=disable")
	if err != nil {
		return err
	}
	logger := log.New(os.Stderr, "SQL: ", log.Flags())
	DB = reform.NewDB(conn, postgresql.Dialect, reform.NewPrintfLogger(logger.Printf))
	return nil
}

func SaveDataFile(uuID, name string) (string, error) {
	f := &models.DataFiles{
		UUID: uuID,
		Name: name,
	}
	err := DB.Save(f)
	return f.UUID, err
}

func DataFile(uuID string) (*models.DataFiles, error) {
	f, err := DB.FindByPrimaryKeyFrom(models.DataFilesTable, uuID)
	if err != nil {
		return nil, err
	}
	return f.(*models.DataFiles), nil
}

func SelectList() ([]models.DataFiles, error) {
	tail := "ORDER BY id DESC LIMIT 20"
	rows, err := DB.SelectRows(models.DataFilesTable, tail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	datafiles := make([]models.DataFiles, 0, 20)
	for {
		var datafile models.DataFiles
		err = DB.NextRow(&datafile, rows)
		if err != nil {
			break
		}
		datafiles = append(datafiles, datafile)
	}
	if err != reform.ErrNoRows {
		return nil, err
	}
	return datafiles, nil
}
