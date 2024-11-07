package database

import (
	"direct/internal/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
)

const createClientListTable = `
	CREATE TABLE IF NOT EXISTS public.client_list (
		id BIGSERIAL NOT NULL,
		archived VARCHAR NULL,
		client_id INT8 NOT NULL,
		created_at VARCHAR NOT NULL,
		login VARCHAR NOT NULL,
		client_info VARCHAR NOT NULL,
		CONSTRAINT client_list_client_id_key UNIQUE (client_id),
		CONSTRAINT client_list_pkey PRIMARY KEY (id)
	);`

const createBonusesTable = `
CREATE TABLE IF NOT EXISTS public.bonuses (
    id BIGSERIAL NOT NULL,
    awaiting_bonus INT8 NULL,
    awaiting_bonus_without_nds INT8 NULL,
    fk_client_list_client_id BIGINT NOT NULL,
    CONSTRAINT bonuses_fk_client_list_client_id_key UNIQUE (fk_client_list_client_id),
    CONSTRAINT bonuses_id_key UNIQUE (id),
    CONSTRAINT bonuses_pkey PRIMARY KEY (id, fk_client_list_client_id),
    CONSTRAINT bonuses_fk_client_list_client_id_fkey FOREIGN KEY (fk_client_list_client_id) REFERENCES public.client_list(client_id)
);`

const createCampaignDataTable = `
CREATE TABLE IF NOT EXISTS public.campaign_data (
    id BIGSERIAL NOT NULL,
    update_date VARCHAR(50) NOT NULL,
    clicks INT8 NOT NULL,
    avg_traffic_volume FLOAT8 NOT NULL,
    cost FLOAT8 NOT NULL,
    avg_impression_position FLOAT8 NOT NULL,
    avg_cpc FLOAT8 NOT NULL,
    avg_pageviews FLOAT8 NOT NULL,
    bounce_rate FLOAT8 NOT NULL,
    client_login VARCHAR(255) NOT NULL,
    CONSTRAINT campaign_data_pkey PRIMARY KEY (id),
    CONSTRAINT campaign_data_client_login UNIQUE (client_login)
);`

func PostgresConnect() (*sqlx.DB, error) {
	token, err := utils.GetToken()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	pool, err := sqlx.Connect("postgres", token.DirectTable)
	if err != nil {
		log.Fatalln(err)
	}

	queries := []string{
		createClientListTable,
		createBonusesTable,
		createCampaignDataTable,
	}

	for _, query := range queries {
		_, err := pool.Exec(query)
		if err != nil {
			return nil, fmt.Errorf("cannot create table: %w", err)
		}
	}

	fmt.Println("Postgres connected")

	return pool, nil
}
