package hash

type PasswordHasher interface {
	GenerateHash(password string) (string, error)
	CompareHashAndPassword(hash, password string) (bool, error)
}
