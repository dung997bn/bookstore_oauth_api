package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to the cassandra cluster:
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cassandra connect sussessfully")
	defer session.Close()
}

//GetSession get session from Cassandra
func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}