package auth

import "os"

type Handler struct {
	logger Logger
}

func NewHandler(logger Logger) Handler {
	return Handler{logger: logger}
}

func (h *Handler) Login(email string, rawPassword string) (string, error) {
	jwt := NewJWTService(os.Getenv("JWT_KEY"))
	hasher := NewHasher()
	uRepo := NewUserRDB()
	uc := NewUseCase(h.logger, jwt, hasher, uRepo)

	return uc.Authenticate(email, rawPassword)
}

func (h *Handler) CreateToken(id string) (string, error) {
	jwt := NewJWTService(os.Getenv("JWT_KEY"))
	hasher := NewHasher()
	uRepo := NewUserRDB()
	uc := NewUseCase(h.logger, jwt, hasher, uRepo)

	return uc.GenerateToken(id)
}

func (h *Handler) Verify(token string) (string, error) {
	jwt := NewJWTService(os.Getenv("JWT_KEY"))
	hasher := NewHasher()
	uRepo := NewUserRDB()
	uc := NewUseCase(h.logger, jwt, hasher, uRepo)

	return uc.VerifyToken(token)
}

func (h *Handler) HashPassword(rawPassword string) (string, error) {
	jwt := NewJWTService(os.Getenv("JWT_KEY"))
	hasher := NewHasher()
	uRepo := NewUserRDB()
	uc := NewUseCase(h.logger, jwt, hasher, uRepo)

	return uc.HashPassword(rawPassword)
}
