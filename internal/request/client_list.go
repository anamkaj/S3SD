package request

import (
	"bytes"
	"direct/internal/models"
	"direct/internal/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAgencyClients() (*models.ResApiDirect, error) {

	token, err := utils.GetToken()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}

	url := "https://api.direct.yandex.com/json/v501/agencyclients"
	bodyData := map[string]interface{}{
		"method": "get",
		"params": map[string]interface{}{
			"SelectionCriteria": map[string]string{
				"Archived": "NO",
			},
			"FieldNames": []string{
				"ClientId",
				"ClientInfo",
				"Login",
				"Archived",
				"CreatedAt",
				"Bonuses",
			},
		},
	}

	jsonData, err := json.Marshal(bodyData)
	if err != nil {
		fmt.Println("Ошибка при кодировании в JSON:", err)

	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", token.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при отправке запроса:", err)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении:", err)

	}

	var respBody models.ResApiDirect

	if err := json.Unmarshal(body, &respBody); err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
	}

	return &respBody, nil
}
