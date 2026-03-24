package factory

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/tanmaykulkarni2112/Winterest/backend/internal/auth/model"
)

// DataServiceImpl implements the DataService interface
type DataServiceImpl struct {
	users      map[string]model.Login
	usersMutex *sync.Mutex
	filePath   string
}

// NewDataService creates a new DataService instance
func NewDataService(filePath string) DataService {
	return &DataServiceImpl{
		users:      make(map[string]model.Login),
		usersMutex: &sync.Mutex{},
		filePath:   filePath,
	}
}

// GetUser retrieves a user from the data store
func (ds *DataServiceImpl) GetUser(username string) (model.Login, bool) {
	ds.usersMutex.Lock()
	defer ds.usersMutex.Unlock()
	user, ok := ds.users[username]
	return user, ok
}

// UserExists checks if a user exists
func (ds *DataServiceImpl) UserExists(username string) bool {
	ds.usersMutex.Lock()
	defer ds.usersMutex.Unlock()
	_, ok := ds.users[username]
	return ok
}

// SaveUser saves a user to the data store
func (ds *DataServiceImpl) SaveUser(username string, user model.Login) error {
	ds.usersMutex.Lock()
	defer ds.usersMutex.Unlock()
	ds.users[username] = user
	return ds.saveToFile()
}

// UpdateUser updates an existing user
func (ds *DataServiceImpl) UpdateUser(username string, user model.Login) error {
	return ds.SaveUser(username, user)
}

// LoadUsersFromFile loads users from the JSON file
func (ds *DataServiceImpl) LoadUsersFromFile() error {
	ds.usersMutex.Lock()
	defer ds.usersMutex.Unlock()

	data, err := os.ReadFile(ds.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(data) == 0 {
		return nil
	}

	err = json.Unmarshal(data, &ds.users)
	return err
}

// SaveUsersToFile saves all users to the JSON file
func (ds *DataServiceImpl) SaveUsersToFile() error {
	ds.usersMutex.Lock()
	defer ds.usersMutex.Unlock()
	return ds.saveToFile()
}

// saveToFile is a private method that saves to file without locking
func (ds *DataServiceImpl) saveToFile() error {
	data, err := json.MarshalIndent(ds.users, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(ds.filePath, data, 0644)
	return err
}

// GetAllUsers returns all users (for testing purposes)
func (ds *DataServiceImpl) GetAllUsers() map[string]model.Login {
	ds.usersMutex.Lock()
	defer ds.usersMutex.Unlock()
	return ds.users
}
