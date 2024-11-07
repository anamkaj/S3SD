package models

type StatStore interface {
	GetStatClient(clientLogin string) (*[]StatisticsClient, error)
}

type ListClientStore interface {
	GetClientList() (*[]List, error)
	InsertClientList(list *ResApiDirect) error
}

type StatisticsClient struct {
	ID                    int64   `json:"id" db:"id"`
	UpdateDate            string  `json:"update_date" db:"update_date"`
	Clicks                int64   `json:"clicks" db:"clicks"`
	Cost                  float64 `json:"cost" db:"cost"`
	AvgImpressionPosition float64 `json:"avg_impression_position" db:"avg_impression_position"`
	AvgCPC                float64 `json:"avg_cpc" db:"avg_cpc"`
	AvgPageviews          float64 `json:"avg_pageviews" db:"avg_pageviews"`
	BounceRate            float64 `json:"bounce_rate,omitempty" db:"bounce_rate"`
	ClientLogin           string  `json:"client_login" db:"client_login"`
	AvgTrafficVolume      float64 `json:"avg_traffic_volume" db:"avg_traffic_volume"`
}

// Получение списка клиентов из db
type List struct {
	Id                      int64  `json:"id" db:"id"`
	Archived                string `json:"archived" db:"archived"`
	ClientID                int64  `json:"client_id" db:"client_id"`
	CreatedAt               string `json:"createdAt" db:"created_at"`
	Login                   string `json:"login" db:"login"`
	ClientInfo              string `json:"clientInfo" db:"client_info"`
	AwaitingBonus           int64  `json:"awaiting_bonus" db:"awaiting_bonus"`
	AwaitingBonusWithoutNds int64  `json:"awaiting_bonus_without_nds" db:"awaiting_bonus_without_nds"`
}

type ResApiDirect struct {
	Result struct {
		Clients []struct {
			Archived  string `json:"Archived"`
			ClientID  int    `json:"ClientId"`
			CreatedAt string `json:"CreatedAt"`
			Bonuses   struct {
				AwaitingBonus           int `json:"AwaitingBonus"`
				AwaitingBonusWithoutNds int `json:"AwaitingBonusWithoutNds"`
			} `json:"Bonuses"`
			Login      string `json:"Login"`
			ClientInfo string `json:"ClientInfo"`
		} `json:"Clients"`
	} `json:"result"`
}
