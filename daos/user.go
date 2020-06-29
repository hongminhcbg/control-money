package daos

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/hongminhcbg/control-money/bcrypt"
	"github.com/hongminhcbg/control-money/models"
)

type UserDao interface {
	Login(userName, pass string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	CreateLog(log models.Log) (*models.Log, error)
	GetLog(userID int64, begin *time.Time, end *time.Time) ([]models.Log, error)
	GetByUserName(userName string) (*models.User, error)
}

type userDaoImpl struct {
	db           *gorm.DB
	bcryptClient bcrypt.BcryptClient
}

func NewUserDao(db *gorm.DB) UserDao {
	bcryptClient := bcrypt.NewBcryptClient()
	return &userDaoImpl{db: db, bcryptClient: bcryptClient}
}

func (dao *userDaoImpl) Login(userName, pass string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}

	if !dao.bcryptClient.IsMatchingHashAndPassword([]byte(user.Password), []byte(pass)) {
		return nil, fmt.Errorf("not matching password")
	}
	return &user, nil
}

func (dao *userDaoImpl) Create(user models.User) (*models.User, error) {
	pwBytes, _ := dao.bcryptClient.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(pwBytes)
	if err := dao.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDaoImpl) CreateLog(log models.Log) (*models.Log, error) {
	if err := dao.db.Create(&log).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func (dao *userDaoImpl) GetLog(userID int64, begin *time.Time, end *time.Time) ([]models.Log, error) {
	logs := make([]models.Log, 0)
	if err := dao.db.Where("user_id = ? AND updated_at>=? AND updated_at<?", userID,
		begin.Format("2006-01-02 15:04:05"),
		end.Format("2006-01-02 15:04:05")).Find(&logs).Error; err != nil {
		return nil, err
	}

	return logs, nil
}

func (dao *userDaoImpl) GetByUserName(userName string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("username = ?", userName).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
