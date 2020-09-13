package database

import (
	"context"
	"github.com/jinzhu/gorm"
	"integration-ginkgo-example/internal/account/model"
	"integration-ginkgo-example/pkg/logging"
)

type AccountDB struct {
	db *gorm.DB
}

func (db *AccountDB) FindAccounts(ctx context.Context) ([]*model.Account, error) {
	logger := logging.FromContext(ctx)
	logger.Debug("account.db.findAccounts")

	var res []*model.Account
	if err := db.db.Find(&res).Error; err != nil {
		logger.Errorw("account.db.findAccounts", "err", err)
		return nil, err
	}
	return res, nil
}

func (db *AccountDB) FindById(ctx context.Context, id uint) (*model.Account, error) {
	logger := logging.FromContext(ctx)
	logger.Debug("account.db.findById", "id", id)

	var res model.Account
	if err := db.db.First(&res, "id = ?", id).Error; err != nil {
		logger.Errorw("account.db.findById", "err", err)
		return nil, err
	}
	return &res, nil
}

func (db *AccountDB) Save(ctx context.Context, acc *model.Account) error {
	logger := logging.FromContext(ctx)
	logger.Debug("account.db.save", "acc", acc)

	if err := db.db.Create(acc).Error; err != nil {
		logger.Errorw("account.db.save", "err", err)
		return err
	}
	return nil
}

func (db *AccountDB) Update(ctx context.Context, acc *model.Account) (int64, error) {
	logger := logging.FromContext(ctx)
	logger.Debug("account.db.update", "acc", acc)

	updated := db.db.Model(acc).Update(model.Account{Username: acc.Username})
	if updated.Error != nil {
		logger.Errorw("account.db.update", "err", updated.Error)
		return 0, updated.Error
	}
	return updated.RowsAffected, nil
}

func (db *AccountDB) Delete(ctx context.Context, id uint) (int64, error) {
	logger := logging.FromContext(ctx)
	logger.Debug("account.db.delete", "id", id)

	deleted := db.db.Delete(&model.Account{}, id)
	if deleted.Error != nil {
		logger.Errorw("account.db.delete", "err", deleted.Error)
		return 0, deleted.Error
	}
	return deleted.RowsAffected, nil
}

func NewAccountDB(db *gorm.DB) *AccountDB {
	return &AccountDB{
		db: db,
	}
}
