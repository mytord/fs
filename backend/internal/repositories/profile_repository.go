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

func (r *ProfileRepository) FindAll(filterFirstName, filterLastName string, limit, offset int) ([]entities.Profile, error) {
	var query string
	var params []interface{}

	query += `SELECT * FROM profiles p WHERE 1=1`

	if filterFirstName != "" {
		query += ` AND first_name LIKE ?`
		params = append(params, filterFirstName+"%")
	}

	if filterLastName != "" {
		query += ` AND last_name LIKE ?`
		params = append(params, filterLastName+"%")
	}

	query += ` ORDER BY id`
	query += ` LIMIT ? OFFSET ?`
	params = append(params, limit, offset)

	rows, err := r.db.Query(query, params...)

	if err != nil {
		return nil, err
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
