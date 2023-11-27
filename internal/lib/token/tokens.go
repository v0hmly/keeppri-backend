package token

type TokenManager interface {
	GenerateToken() (string, error)
}
