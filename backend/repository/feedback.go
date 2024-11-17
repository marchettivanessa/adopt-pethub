package repository

import (
	"adopt-pethub/backend/database"
	"adopt-pethub/backend/domain"

	"github.com/pkg/errors"
)

type FeedbackRepository struct{}

type FeedbackRepositoryInterface interface {
	GetFeedbacks(db *database.Database) ([]domain.Feedback, error)
}

func (r *FeedbackRepository) GetFeedbacks(db *database.Database) ([]domain.Feedback, error) {
	var feedbacks []domain.Feedback
	err := db.Connection.Find(&feedbacks).Error
	if err != nil {
		return nil, errors.New("failed to get feedbacks from database")
	}
	return feedbacks, nil
}

func (r *FeedbackRepository) InsertFeedback(feedback domain.Feedback, db *database.Database) error {
	if err := db.Connection.Create(&feedback).Error; err != nil {
		return errors.Wrap(err, "failed to insert feedback into database")
	}

	return nil
}
