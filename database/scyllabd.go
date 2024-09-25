package database

import (
	"log"
	"strings"

	"github.com/SergioVenicio/grpc_gtw/settings"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
	"github.com/scylladb/gocqlx/v3/qb"
	"github.com/scylladb/gocqlx/v3/table"
)

type ScyllaDB[T interface{}] struct {
	session *gocqlx.Session
}

func (s *ScyllaDB[T]) Save(metadata table.Metadata, obj T) error {
	t := table.New(metadata)
	qb.Update(t.Name())
	q := s.session.Query(t.Insert()).BindStruct(obj)
	if err := q.ExecRelease(); err != nil {
		log.Printf("fail to save data on scylladb %v\n", err)
		return err
	}

	return nil
}

func (s *ScyllaDB[T]) Get(metadata table.Metadata, obj T) (*T, error) {
	t := table.New(metadata)
	var row T
	q := s.session.Query(t.Get()).BindStruct(obj)
	if err := q.GetRelease(&row); err != nil {
		log.Printf("fail to save data on scylladb %v\n", err)
		return nil, err
	}

	return &row, nil
}

func (s *ScyllaDB[T]) Update(metadata table.Metadata, obj T) error {
	t := table.New(metadata)
	q := s.session.Query(t.Update("name", "email")).BindStruct(obj)
	if err := q.ExecRelease(); err != nil {
		log.Printf("fail to update data on scylladb %v\n", err)
		return err
	}

	return nil
}

func (s *ScyllaDB[T]) Delete(metadata table.Metadata, id int32) error {
	t := table.New(metadata)
	q := s.session.Query(t.Delete()).BindMap(qb.M{"id": id})
	if err := q.ExecRelease(); err != nil {
		log.Printf("fail to delete data on scylladb %v\n", err)
		return err
	}

	return nil
}

func NewScyllaDB[T interface{}](cluster *gocql.ClusterConfig, s *settings.Settings) *ScyllaDB[T] {
	if cluster == nil {
		hosts := strings.Split(s.ScylladbURI, ",")
		cluster = gocql.NewCluster(hosts...)
	}
	cluster.Keyspace = "users"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatalf("fail to connect to scylladb cluster %v", err)
	}
	return &ScyllaDB[T]{
		session: &session,
	}
}
