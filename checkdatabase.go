package timeserver

func CorrelateUser(trialedUser *User, database []User) bool {
	for _, validUser := range database {
		if validUser.Email == trialedUser.Email && validUser.Name == trialedUser.Name && validUser.Password == trialedUser.Password {
			return true
		}
	}
	return false
}
