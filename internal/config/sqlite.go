package config

import (
	"os"

	"github.com/H0wZy/user-api/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "../../internal/db"
	dbFullPath := dbPath + "/user.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Infof("arquivo db não existe, criando...")
		err := os.MkdirAll(dbPath, os.ModePerm)

		if err != nil {
			logger.Errorf("erro na criação do diretório db: %v", err)
			return nil, err
		}

		file, err := os.Create(dbFullPath)

		if err != nil {
			logger.Errorf("erro na criação do arquivo db: %v", err)
			return nil, err
		}

		err_close := file.Close()

		if err_close != nil {
			logger.Errorf("erro ao fechar arquivo db: %v", err_close)
			return nil, err_close
		}

		logger.Infof("arquivo db criado com sucesso!")
	}

	db, err := gorm.Open(sqlite.Open(dbFullPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("erro ao abrir conexão com sqlite: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		logger.Errorf("erro ao fazer migration da entidade User: %v", err)
		return nil, err
	}

	logger.Infof("migration concluída com sucesso!")
	return db, nil
}
