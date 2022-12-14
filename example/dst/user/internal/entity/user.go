// Code generated by skeleton; DO NOT EDIT.

package entity
const (
	UserCreateEvent = "USER_CREATE_EVENT"
	UserUpdateEvent = "USER_UPDATE_EVENT"
	UserDeleteEvent = "USER_DELETE_EVENT"
	)

// User ...
type User struct {
	id int
}

// NewUser ...
func NewUser(id int) *User {
	u := &User{id}
	return u
}

// Id ...
func (m *User) Id() int {
	return m.id
}
