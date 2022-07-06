package postgres

/*
 * Author      : jodyalmaida (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : auth-skm
 */
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	dto "auth-skm/src/app/dtos/layanan"
	repositories "auth-skm/src/domain/repositories"

	_ "github.com/joho/godotenv/autoload"
)

type layananRepository struct {
}

func NewLayananService() repositories.LayananRepository {
	return &layananRepository{}
}

func (s *layananRepository) GetLayananByOpdId(reqDTO *dto.GetLayananReqDTO) ([]*dto.LayananRespDTO, error) {

	ordersListUrl := os.Getenv("GET_LAYANAN_URL")
	var responseData *dto.LayananIntegrateDTO

	req, err := http.NewRequest("GET", ordersListUrl, nil)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to create http request Get Layanan: %v", err)
	}

	q := req.URL.Query()
	q.Add("opd_id", reqDTO.OpdID)

	req.URL.RawQuery = q.Encode()

	client := &http.Client{
		Timeout: time.Second * 50,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	log.Println("Success execute  : ", ordersListUrl)

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, fmt.Errorf("failed decode get orders response: %v", err)
	}

	if len(responseData.Data) < 1 {
		return nil, errors.New("data not found")
	}

	return responseData.Data, nil
}
