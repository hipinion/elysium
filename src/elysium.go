package elysium

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
)

type xOrganization struct {
}

type Page struct {
	Content interface{}
}

type Config struct {
	Name     string         `json:"name"`
	File     string         `json:"filename"`
	Database ConfigDatabase `json:"database"`
	Dev      ConfigDev      `json:"dev"`
	Stage    ConfigStage    `json:"stage"`
	Prod     ConfigProd     `json:"production"`
}

type ConfigDatabase struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"database"`
}

type ConfigDev struct {
	ConfigEnvironment string `json:"env"`
}

type ConfigStage struct {
	ConfigEnvironment string `json:"env"`
}

type ConfigProd struct {
	ConfigEnvironment string `json:"env"`
}

type ConfigEnvironment struct {
	Port string `json:"port"`
}

var (
	Configuration Config
	DB            *sql.DB
)

func Foo() {

}

func Init(filename string) {
	Configuration.File = filename
	loadConfiguration(filename)
	connectDB()
}

func loadConfiguration(filename string) {
	cf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("uga")
	}
	json.Unmarshal(cf, &Configuration)
	log.Println(Configuration)
}

func getConnectString() string {
	cs := Configuration.Database.User + ":" + Configuration.Database.Pass + "@" + Configuration.Database.Host + "/" + Configuration.Database.Database
	return cs
}

func connectDB() {
	log.Println("Connecting to", getConnectString())
	db, err := sql.Open(Configuration.Database.Type, getConnectString())
	if err != nil {
		log.Fatal("EAAAAAHHHH")
	}
	DB = db
	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Elysium() {
	log.Println("UH")
}
