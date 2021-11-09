package repository

import (
	"context"
	"strings"
	"tServerOra/internal/models"

	guuid "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ServerRepo struct {
	connStr string
	db      *sqlx.DB
	cancel  context.CancelFunc
}

type UsersRepo struct {
	Data      map[string]int
	CurrentID int
}

// func NewGuuid генерирует Uuid
func NewGuuid() string {
	return strings.ToUpper(strings.Replace(guuid.New().String(), "-", "", 4))
}

// func (s *ServerRepo) SaveCard
func (s *ServerRepo) SaveCard(ctx context.Context, cTC *models.CardTC) error {
	db := s.db
	tx := db.MustBegin()
	defer tx.Rollback()
	//string([]rune(cTC.DriverName)[0:100]),
	//string([]rune(cTC.NumTC)[0:20]),
	//string([]rune(cTC.ModelTC)[0:30]))
	tx.MustExec("INSERT INTO viewblank (datestart, journal_uuid, fio, numtc, markatc, numpricep, plomb1, container1, typetc, sizetc, damage) VALUES (sysdate, :journal_uuid, :fio, :numtc, :markatc, :numpricep, :plomb1, :container1, :typetc, :sizetc, :damage)",
		NewGuuid(),
		strings.TrimSpace(cTC.DriverName),
		strings.TrimSpace(cTC.NumTC),
		strings.TrimSpace(cTC.ModelTC),
		strings.TrimSpace(cTC.NumPric),
		strings.TrimSpace(cTC.NumPlomb),
		strings.TrimSpace(cTC.ContNum),
		strings.TrimSpace(cTC.TypeTC),
		strings.TrimSpace(cTC.SizeTC),
		strings.TrimSpace(cTC.Remark))

	tx.Commit()

	return nil
}

func (s *ServerRepo) CreateUser(ctx context.Context) (string, error) {
	/*db := s.db
	ctx, cancelfunc := context.WithTimeout(ctx, 5*time.Second)
	defer cancelfunc()

	ur := uuid.New()
	urEnc, err := encription.EncriptStr(ur.String())
	if err != nil {
		return "", err
	}
	q := `INSERT INTO users (user_uuid, user_enc_id) VALUES ($1, $2)`

	if _, err := db.ExecContext(ctx, q, ur, urEnc); err != nil {
		return "", err
	}

	return urEnc, nil
	*/
	return " ", nil
}
