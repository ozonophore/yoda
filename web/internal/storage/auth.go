package storage

import (
	"errors"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type UserProfile struct {
	UserId      int32  `gorm:"column:id"`
	Email       string `gorm:"column:email"`
	UserName    string `gorm:"column:name"`
	Password    string `gorm:"column:password"`
	Permissions []string
}

func (s *Storage) GetPermissionByUserId(id int32) (*[]string, error) {
	var roles []string
	tx := s.db.Raw(`select ur.role_code from ml.user_role ur
			where ur.user_id = ?`, id).Scan(&roles)
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return &[]string{}, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	var permissions []string
	if slices.Contains(roles, "ADMIN") {
		tx = s.db.Raw(`select rp.code from ml.permission rp`).Scan(&permissions)
	} else {
		tx = s.db.Raw(`select distinct rp.code from ml.role_permission rp
			where rp.role_code in (?)`, roles).Scan(&permissions)
	}
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &permissions, nil
}

func (s *Storage) GetProleByLogin(login string) (*UserProfile, error) {
	var profile UserProfile
	s.db.Raw(`select id, email, name, password from ml.users where email=?`, login).Scan(&profile)
	permissions, err := s.GetPermissionByUserId(profile.UserId)
	if err != nil {
		return nil, err
	}
	profile.Permissions = *permissions
	return &profile, nil
}

func (s *Storage) UserExists(id int32) bool {
	var count int32
	s.db.Raw(`select count(1) from ml.users where id=?`, id).Scan(&count)
	return count == 1
}
