package cassandra

import (
	"github.com/gocql/gocql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type CassHosts struct {
	Hosts []string `yaml:"hosts"`
}

const (
	in_env_yml_path = "YML_PATH"
)

var (
	csh       CassHosts
	session   *gocql.Session
	yaml_path = os.Getenv(in_env_yml_path)
)

func init() {
	yamlFile, err := ioutil.ReadFile(yaml_path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &csh)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	cluster := gocql.NewCluster()
	for _, cassandra_host := range csh.Hosts {
		cluster = gocql.NewCluster(cassandra_host)
	}

	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
