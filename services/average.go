package services

import (
	"errors"
	"github.com/hongminhcbg/control-money/daos"
	"time"
)

type AverageService interface {
	CalculateByDay(userID int64, begin *time.Time, end *time.Time) (int64, error)
}

func NewAverageService(userDao daos.UserDao) AverageService {
	return &averageServiceImpl{userDao:userDao,}
}

type averageServiceImpl struct {
	userDao daos.UserDao
}

func (avr *averageServiceImpl) CalculateByDay(userID int64, begin *time.Time, end *time.Time) (int64, error){
	delta := end.Sub(*begin)
	days := delta / (24*time.Hour)
	if days <= 0 {
		return -1, errors.New("end time must > begin time")
	}

	logs, err := avr.userDao.GetLog(userID, begin, end)
	if err != nil {
		return -1, err
	}

	sumMoney := int64(0)
	for _, val := range logs {
		sumMoney += val.Money
	}

	result := sumMoney / int64(days)
	return result, nil
}

