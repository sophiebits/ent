package configs

// AccountConfig is the config for test accounts in test land
type AccountConfig struct {
	FirstName   string
	LastName    string
	PhoneNumber string
}

// GetTableName returns the underyling database table the account model's data is stored
func (config *AccountConfig) GetTableName() string {
	return "accounts"
}