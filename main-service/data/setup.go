package data

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Log *zap.Logger

/* make database connection set up */
func makeDbConnection() *gorm.DB{

    time.Sleep(time.Duration(5000 * time.Millisecond))

	dsn := "docker:docker@tcp(mysql_host:3306)/test_database?multiStatements=True&charset=utf8mb4&parseTime=True&loc=Local"

	sqlDB, error := sql.Open("mysql", dsn)
    if error != nil {
		Log.Fatal(error.Error())
	}

    db, err := gorm.Open(mysql.New(mysql.Config{
        Conn: sqlDB,
    }), &gorm.Config{})

	if err != nil {
			Log.Warn("database not yet ready!")
		} else {
			Log.Info("connected to database!")
		}

	if db == nil {
		Log.Panic("can't connect to database!")
	}

  db.AutoMigrate(&LoginUser{},&Messages{},&SendMessage{})
  return db
}

/* execute message list sql file */
func executeSqlFile(db *gorm.DB,filePath string) {

    // Read file from the directory
    file, err := ioutil.ReadFile(filePath)

    if err != nil {
        Log.Fatal(err.Error())
    }

    // Execute sql file 
    result:= db.Exec(string(file))
    if result.Error != nil {
        Log.Fatal(result.Error.Error())
    }
}

func MakeLogger() {

	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "Message",
	    "levelKey": "Level",
	    "levelEncoder": "capital",
		"timeKey": "Time",
		"timeEncoder": "iso8601"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}

	Log, _ = cfg.Build()
	defer Log.Sync()
	Log.Info("logger construction succeeded")
}

