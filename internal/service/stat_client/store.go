package stat_client

import (
	"direct/internal/models"
	"github.com/jmoiron/sqlx"
	"log"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetStatClient(clientLogin string) (*[]models.StatisticsClient, error) {

	qStat := `SELECT *
        FROM campaign_data
        WHERE client_login=$1;`

	data := []models.StatisticsClient{}

	err := s.db.Select(&data, qStat, clientLogin)
	if err != nil {
		log.Fatalln(err)
	}
	return &data, nil
}
