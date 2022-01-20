package repositories

import (
	"database/sql"
	"github.com/mytord/fs/backend/internal/entities"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{
		db: db,
	}
}

func (r *ProfileRepository) ExistsByEmail(email string) (bool, error) {
	var cnt int

	err := r.db.QueryRow(`SELECT COUNT(1) FROM profiles WHERE email = ?`, email).Scan(&cnt)

	if err != nil {
		return false, err
	}

	return cnt > 0, nil
}

func (r *ProfileRepository) Add(profile *entities.Profile) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO profiles (email, password, first_name, last_name, city, age, interests) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(profile.Email, profile.Password, profile.FirstName, profile.LastName, profile.City, profile.Age, profile.Interests)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	profile.Id = int(id)

	return err
}

func (r *ProfileRepository) Find(id int) (entities.Profile, error) {
	var profile entities.Profile

	err := r.db.QueryRow(`
		SELECT * FROM profiles WHERE id = ?
	`, id).Scan(
		&profile.Id,
		&profile.Email,
		&profile.Password,
		&profile.FirstName,
		&profile.LastName,
		&profile.City,
		&profile.Age,
		&profile.Interests,
	)

	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *ProfileRepository) FindAll(search string, limit, offset int) ([]entities.Profile, error) {
	var query string

	query += `SELECT * FROM profiles p`

	if search != "" {
		query += ` WHERE first_name LIKE ? OR last_name LIKE ?`
	}

	query += ` ORDER BY id`
	query += ` LIMIT ? OFFSET ?`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if search != "" {
		search = search + string('%')
		rows, err = stmt.Query(search, search, limit, offset)
	} else {
		rows, err = stmt.Query(limit, offset)
	}

	defer rows.Close()

	var profiles []entities.Profile

	for rows.Next() {
		var profile entities.Profile

		err := rows.Scan(
			&profile.Id,
			&profile.Email,
			&profile.Password,
			&profile.FirstName,
			&profile.LastName,
			&profile.City,
			&profile.Age,
			&profile.Interests,
		)

		if err != nil {
			return nil, err
		}

		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (r *ProfileRepository) FindByEmail(email string) (entities.Profile, error) {
	var profile entities.Profile

	err := r.db.QueryRow(`SELECT * FROM profiles WHERE email = ?`, email).Scan(
		&profile.Id,
		&profile.Email,
		&profile.Password,
		&profile.FirstName,
		&profile.LastName,
		&profile.City,
		&profile.Age,
		&profile.Interests,
	)

	if err != nil {
		return profile, err
	}

	return profile, nil
}
