package postgres

import (
	"context"

	"github.com/ademuanthony/achibiti/accounts/postgres/models"
	go_micro_srv_accounts "github.com/ademuanthony/achibiti/accounts/proto/accounts"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func (pg *PgDb) CreateUser(ctx context.Context, user go_micro_srv_accounts.User, hashedPassword string) error {
	userModel := models.AccountUser{
		ID:           user.Id,
		Name:         user.Name,
		Username:     user.Username,
		PasswordHash: hashedPassword,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		Role:         user.Role,
	}

	return userModel.Insert(ctx, pg.db, boil.Infer())
}

func (pg *PgDb) FindUserByUsername(ctx context.Context, username string) (user *go_micro_srv_accounts.User, err error) {
	accountModel, err := models.AccountUsers(models.AccountUserWhere.Username.EQ(username)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	return &go_micro_srv_accounts.User{
		Id:                   accountModel.ID,
		Username:             accountModel.Username,
		Email:                accountModel.Email,
		PhoneNumber:          accountModel.PhoneNumber,
		Name:                 accountModel.Name,
		Role:                 accountModel.Role,
	}, nil
}

func (pg *PgDb) FindUserByEmail(ctx context.Context, email string) (user *go_micro_srv_accounts.User, err error) {
	accountModel, err := models.AccountUsers(models.AccountUserWhere.Email.EQ(email)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	return &go_micro_srv_accounts.User{
		Id:          accountModel.ID,
		Username:    accountModel.Username,
		Email:       accountModel.Email,
		PhoneNumber: accountModel.PhoneNumber,
		Name:        accountModel.Name,
		Role:        accountModel.Role,
	}, nil
}

func (pg *PgDb) FindUserByPhone(ctx context.Context, phoneNumber string) (user *go_micro_srv_accounts.User, err error) {
	accountModel, err := models.AccountUsers(models.AccountUserWhere.PhoneNumber.EQ(phoneNumber)).One(ctx, pg.db)
	if err != nil {
		return nil, err
	}

	return &go_micro_srv_accounts.User{
		Id:          accountModel.ID,
		Username:    accountModel.Username,
		Email:       accountModel.Email,
		PhoneNumber: accountModel.PhoneNumber,
		Name:        accountModel.Name,
		Role:        accountModel.Role,
	}, nil
}

func (pg *PgDb) GetPasswordHash(ctx context.Context, username string) (string, error) {
	accountModel, err := models.AccountUsers(models.AccountUserWhere.Username.EQ(username)).One(ctx, pg.db)
	if err != nil {
		return "", err
	}

	return accountModel.PasswordHash, nil
}

func (pg *PgDb) Disable(ctx context.Context, username string) error {
	panic("implement me")
}

func (pg *PgDb) GetUsers(ctx context.Context, skipCount int32, maxResultCount int32) ([]*go_micro_srv_accounts.User, int32, error) {
	totalCount, err := models.AccountUsers().Count(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	userSlice, err := models.AccountUsers(qm.Offset(int(skipCount)), qm.Limit(int(maxResultCount))).All(ctx, pg.db)
	if err != nil {
		return nil, 0, err
	}

	var users []*go_micro_srv_accounts.User
	for _, accountModel := range userSlice  {
		users = append(users, &go_micro_srv_accounts.User{
			Id:                   accountModel.ID,
			Username:             accountModel.Username,
			Email:                accountModel.Email,
			PhoneNumber:          accountModel.PhoneNumber,
			Name:                 accountModel.Name,
			Role:                 accountModel.Role,
		})
	}

	return users, int32(totalCount), nil
}
