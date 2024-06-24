package tags

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTag(tag *Tag) error {
	result, err := r.db.Exec("INSERT INTO tags (name) VALUES (?)", tag.Name)
	if err != nil {
		return err
	}
	tagID, _ := result.LastInsertId()
	tag.ID = int(tagID)
	return nil
}

// Add other repository methods as needed, such as GetTagByID, UpdateTag, DeleteTag, etc.
