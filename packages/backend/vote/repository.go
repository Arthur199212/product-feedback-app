package vote

import "database/sql"

type VoteRepository interface {
}

type voteRepository struct {
	db *sql.DB
}

func NewVoteRepository(db *sql.DB) *voteRepository {
	return &voteRepository{db}
}
