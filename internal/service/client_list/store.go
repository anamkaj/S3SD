package list_client

import (
	"direct/internal/models"
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetClientList() (*[]models.List, error) {

	list := []models.List{}

	query := `
	SELECT
		c.id,
		c.client_id,
		c.login,
		c.created_at,
		c.client_info,
		c.archived,
		b.awaiting_bonus,
		b.awaiting_bonus_without_nds
	FROM client_list c
	INNER JOIN 
		bonuses b ON c.client_id = b.fk_client_list_client_id;`

	err := s.db.Select(&list, query)
	if err != nil {
		log.Fatalln(err)
	}

	return &list, nil
}

func (s *Store) InsertClientList(list *models.ResApiDirect) error {

	qList := `INSERT INTO client_list 
        (client_id, login, created_at, client_info, archived)
        VALUES ($1, $2, $3, $4, $5)
        ON CONFLICT (client_id) 
        DO UPDATE SET
        login = EXCLUDED.login,
        created_at = EXCLUDED.created_at,
        client_info = EXCLUDED.client_info,
        archived = EXCLUDED.archived;`

	qBonus := `INSERT INTO bonuses 
            ( awaiting_bonus, 
            awaiting_bonus_without_nds, 
            fk_client_list_client_id
            )
            VALUES ($1, $2, $3)
            ON CONFLICT (fk_client_list_client_id) 
            DO UPDATE SET
            awaiting_bonus = EXCLUDED.awaiting_bonus,
            awaiting_bonus_without_nds = EXCLUDED.awaiting_bonus_without_nds;`

	for _, client := range list.Result.Clients {

		_, err := s.db.Exec(qList,
			client.ClientID,
			client.Login,
			client.CreatedAt,
			client.ClientInfo,
			client.Archived,
		)
		if err != nil {
			return fmt.Errorf("ошибка при вставке клиента: %v", err)
		}

		_, err = s.db.Exec(qBonus,
			client.Bonuses.AwaitingBonus,
			client.Bonuses.AwaitingBonusWithoutNds,
			client.ClientID,
		)
		if err != nil {
			return fmt.Errorf("ошибка при вставке бонусов: %v", err)
		}
	}

	return nil
}
