package repository

import (
	"database/sql"
	"net/http"
)

type MentorshipRepository struct {
	db *sql.DB
}

func NewMentorshipRepository(db *sql.DB) *MentorshipRepository {
	return &MentorshipRepository{db: db}
}

func (m *MentorshipRepository) GetUserDataMentorship(dosenID int) ([]User, error) {
	sqlStatement := "SELECT u.id,u.name, ud.company, ud.program,ud.batch,ud.nrp,ud.prodi FROM users u LEFT JOIN mentorship m ON u.id = m.mahasiswa_id LEFT JOIN user_details ud ON u.id = ud.user_id WHERE m.dosen_id = $1"
	var users []User
	rows, err := m.db.Query(sqlStatement, dosenID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Company, &user.Program, &user.Batch, &user.Nrp, &user.Prodi)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (m *MentorshipRepository) InsertMentorship(mahasiswa_id, dosen_id int) (responseCode int, err error) {
	sqlStatement := "INSERT INTO mentorship (mahasiswa_id,dosen_id) VALUES ($1,$2) RETURNING id"
	var id int
	err = m.db.QueryRow(sqlStatement, mahasiswa_id, dosen_id).Scan(&id)

	return http.StatusAccepted, err
}
func (m *MentorshipRepository) GetUserDataMentorships() ([]User, error) {
	sqlStatement := "SELECT u.id,u.name,ud.company FROM users u LEFT JOIN mentorship m ON u.id = m.mahasiswa_id LEFT JOIN user_details ud ON u.id = ud.user_id WHERE u.role = 'mahasiswa' "
	var users []User
	rows, err := m.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Company)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
