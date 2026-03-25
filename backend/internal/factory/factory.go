package factory

// Factory is responsible for creating and managing service instances
type Factory struct {
	dataService DataService
	authService AuthService
	userService UserService
	deps        *HandlerDependencies
}

// NewFactory creates a new Factory instance with all services
func NewFactory(dataFilePath string) *Factory {
	dataService := NewDataService(dataFilePath)
	authService := NewAuthService(dataService)
	userService := NewUserService(dataService, authService)

	deps := &HandlerDependencies{
		DataService: dataService,
		AuthService: authService,
		UserService: userService,
	}

	return &Factory{
		dataService: dataService,
		authService: authService,
		userService: userService,
		deps:        deps,
	}
}

// GetDataService returns the DataService instance
func (f *Factory) GetDataService() DataService {
	return f.dataService
}

// GetAuthService returns the AuthService instance
func (f *Factory) GetAuthService() AuthService {
	return f.authService
}

// GetUserService returns the UserService instance
func (f *Factory) GetUserService() UserService {
	return f.userService
}

// GetHandlerDependencies returns all dependencies for handlers
func (f *Factory) GetHandlerDependencies() *HandlerDependencies {
	return f.deps
}

// Initialize loads initial data (like users from file)
func (f *Factory) Initialize() error {
	return f.dataService.LoadUsersFromFile()
}
