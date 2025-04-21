package reception

import (
	"avitotes/database"
	"fmt"
)

type ReceptionRepoDeps struct {
	Database *database.Db
}

type ReceptionRepo struct {
	Database *database.Db
}

func NewReceptionRepo(database *database.Db) *ReceptionRepo {
	return &ReceptionRepo{
		Database: database,
	}
}

// создание приемки для клинтов
func (repo *ReceptionRepo) Create(reception *Reception) (*Reception, error) {
	query := `INSERT INTO receptions (date_time, pvzId, status) VALUES ($1, $2, $3) RETURNING id`
	err := repo.Database.MyDb.QueryRow(query, reception.DataTime, reception.PvzID, reception.Status).Scan(&reception.ID)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании приемки: %w", err)
	}
	return reception, nil
}
