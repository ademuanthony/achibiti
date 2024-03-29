package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	acl "github.com/ademuanthony/achibiti/acl/proto/acl"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/micro/go-micro/util/log"
	"golang.org/x/crypto/bcrypt"
)


type Claims struct {
	Username string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type DataStore interface {
	CreateUser(ctx context.Context, user acl.User, hashedPassword string) error
	FindUserByUsername(ctx context.Context, username string) (user *acl.User, err error)
	FindUserByEmail(ctx context.Context, email string) (user *acl.User, err error)
	FindUserByPhone(ctx context.Context, phoneNumber string) (user *acl.User, err error)
	GetPasswordHash(ctx context.Context, username string) (string, error)
	Disable(ctx context.Context, username string) error
	GetUsers(ctx context.Context, skipCount int32, maxResultCount int32) ([]*acl.User, int32, error)
	DeleteUser(ctx context.Context, username string) error
}

type aclHandler struct{
	store DataStore
}

func NewAclHandler(store DataStore) *aclHandler {
	return &aclHandler{
		store:store,
	}
}

func (a aclHandler) CreateUser(ctx context.Context, req *acl.CreateUserRequest, resp *acl.CreateUserResponse) error {
	if u, _ := a.store.FindUserByUsername(ctx, req.Username); u != nil {
		return fmt.Errorf("the username, %s has been taken", req.Username)
	}

	if u, _ := a.store.FindUserByEmail(ctx, req.Email); u != nil {
		return fmt.Errorf("the email, %s has been taken", req.Email)
	}

	if u, _ := a.store.FindUserByPhone(ctx, req.PhoneNumber); u != nil {
		return fmt.Errorf("the phone number, %s has been taken", req.PhoneNumber)
	}

	id, err := uuid.NewV4()
	if err != nil {
		return fmt.Errorf("cannot generate uuid, %s", err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return fmt.Errorf("error in hashinf password, %s", err.Error())
	}

	user := acl.User{
		Id:                   id.String(),
		Username:             req.Username,
		Email:                req.Email,
		PhoneNumber:          req.PhoneNumber,
		Name:                 req.Name,
		Role:                 req.Role,
	}

	if err = a.store.CreateUser(ctx, user, string(hash)); err != nil {
		return err
	}

	resp.Id = id.String()
	return nil
}

func (a aclHandler) Login(ctx context.Context, req *acl.LoginRequest, resp *acl.LoginResponse) error {
	if req.Username == "" || req.Password == "" {
		return errors.New("username and password are required")
	}
	user, err := a.store.FindUserByUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid credentials")
		}

		return fmt.Errorf("internal error occured while trying to log you into the system, %s", err.Error())
	}

	password, err := a.store.GetPasswordHash(ctx, req.Username)
	if err != nil {
		return fmt.Errorf("internal error occured while trying to log you into the system, %s", err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		return errors.New("invalid credentials")
	}

	tokenString, err := generateJwtString(user)
	if err != nil {
		return err
	}

	resp.Token = tokenString
	resp.Email = user.Email
	resp.Name = user.Name
	resp.PhoneNumber = user.PhoneNumber
	resp.Role = user.Role
	resp.Username = user.Username

	return nil
}



func (a aclHandler) RefreshToken(ctx context.Context, req *acl.RefreshTokenRequest, resp *acl.LoginResponse) error {
	user, err := a.store.FindUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid username")
		}

		return err
	}

	tokenString, err := generateJwtString(user)
	if err != nil {
		return err
	}

	resp.Token = tokenString
	resp.Email = user.Email
	resp.Name = user.Name
	resp.PhoneNumber = user.PhoneNumber
	resp.Role = user.Role
	resp.Username = user.Username
	return nil
}

func generateJwtString(user *acl.User) (string, error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(59 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user.Username,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string // todo user a signed key
	tokenString, err := token.SignedString([]byte("234dfsgk593jffjdh9ekjdsfjk43089432kjkfjfadj4390fdjk3490dgskljgdsk2390gshgfddfhjk2398-glsjl"))
	if err != nil {
		log.Log("Error in generating JWT token, %s", err.Error())
		return "", fmt.Errorf("internal error during login")
	}

	return tokenString, nil
}

func (a aclHandler) UpdateUser(ctx context.Context, req *acl.UpdateUserRequest, resp *acl.EmptyMessage) error {
	panic("implement me")
}

func (a aclHandler) DeleteUser(ctx context.Context, req *acl.DeleteUserRequest, resp *acl.EmptyMessage) error {
	return a.store.DeleteUser(ctx, req.GetUsername())
}

func (a aclHandler) DisableUser(ctx context.Context, req *acl.DisableUserRequest, resp *acl.EmptyMessage) error {
	if err := a.store.Disable(ctx, req.Username); err != nil {
		return fmt.Errorf("error in disabling account, %s", err.Error())
	}
	return nil
}

func (a aclHandler) Users(ctx context.Context, req *acl.UsersRequest, resp *acl.UsersResponse) error {
	users, totalCount, err := a.store.GetUsers(ctx, req.SkipCount, req.MaxResultCount)
	resp.Users = users
	resp.TotalCount = totalCount

	return err
}

func (a aclHandler) UserDetails(ctx context.Context, req *acl.UserDetailsRequest, resp *acl.UserDetailsResponse) error {
	user, err := a.store.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return err
	}

	resp.User = user
	return nil
}

func (a aclHandler) PasswordResetToken(ctx context.Context, req *acl.PasswordResetTokenRequest, resp *acl.PasswordResetTokenResponse) error {
	panic("implement me")
}

func (a aclHandler) ResetPassword(ctx context.Context, req *acl.ResetPasswordRequest, resp *acl.EmptyMessage) error {
	panic("implement me")
}

func (a aclHandler) ChangePassword(ctx context.Context, req *acl.ChangePasswordRequest, resp *acl.EmptyMessage) error {
	panic("implement me")
}

func (a aclHandler) AddRole(ctx context.Context, req *acl.AddRoleRequest, resp *acl.AddRoleRequest) error {
	panic("implement me")
}

func (a aclHandler) GetRoles(ctx context.Context, req *acl.EmptyMessage, resp *acl.GetRolesResponse) error {
	panic("implement me")
}

func (a aclHandler) ChangeRole(ctx context.Context, req *acl.ChangeRoleRequest, resp *acl.EmptyMessage) error {
	panic("implement me")
}

