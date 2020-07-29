package postgres

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type (
	Config struct {
		Address string
	}

	EventsRepo struct {
		db *sqlx.DB
	}
)

func New(cfg Config) (*EventsRepo, error) {
	db, err := sqlx.Connect("postgres", cfg.Address)
	if err != nil {
		return nil, err
	}

	return &EventsRepo{db: db}, nil
}

func (r *EventsRepo) HealthCheck() error {
	return r.db.Ping()
}

func (r *EventsRepo) Close() {
	if err := r.db.Close(); err != nil {
		log.WithError(err).Error("closing postgress connection")
	}
}
