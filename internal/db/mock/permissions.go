package mock

import (
	"github.com/mgeale/homeserver/internal/db"
)

var mockPermission = &db.Permissions{"", ""}

type PermissionModel struct{}

func (m PermissionModel) AddForUser(userID int64, codes string) error {
	return nil
}
func (m PermissionModel) GetAllForUser(userID int64) (db.Permissions, error) {
	return *mockPermission, nil
}
