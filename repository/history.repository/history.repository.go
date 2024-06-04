package historyrepository

import "github.com/ArdiSasongko/app_ticketing/db/model/domain"

type HistoryRepoInterface interface {
	Create(history *domain.History) error
}
