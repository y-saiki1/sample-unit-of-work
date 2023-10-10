package auth

const (
	PASSWORD_MIN = 8
)

type TokenGenerator interface {
	GenerateToken(id string) (token string, err error)
	VerifyToken(tokenString string) (string, error)
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
}

type Logger interface {
	Error(i ...interface{})
}

type UseCase struct {
	logger         Logger
	tokenGenerator TokenGenerator
	hasher         Hasher
	userRepo       UserRepository
}

func NewUseCase(
	logger Logger,
	tokenGenerator TokenGenerator,
	hasher Hasher,
	userRepo UserRepository,
) UseCase {
	return UseCase{
		logger:         logger,
		tokenGenerator: tokenGenerator,
		hasher:         hasher,
		userRepo:       userRepo,
	}
}

func (s *UseCase) Authenticate(email string, rawPassword string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		s.logger.Error(err)
		return "", ErrorUserAuthentication
	}
	if err := s.hasher.Hashed(rawPassword); err != nil {
		return "", ErrorUserAuthentication
	}
	if err := s.hasher.Verify(rawPassword); err != nil {
		return "", ErrorUserAuthentication
	}

	tk, err := s.GenerateToken(user.Id)
	if err != nil {
		s.logger.Error(err)
		return "", ErrorTokenGenerate
	}

	return tk, nil
}

func (s *UseCase) HashPassword(rawPassword string) (string, error) {
	if len(rawPassword) < PASSWORD_MIN {
		return "", ErrorPasswordTooShort
	}
	if err := s.hasher.Hashed(rawPassword); err != nil {
		return "", ErrorTokenVerification
	}
	return s.hasher.HashString(), nil
}

func (s *UseCase) VerifyToken(token string) (string, error) {
	id, err := s.tokenGenerator.VerifyToken(token)
	if err != nil {
		s.logger.Error(err)
		return "", ErrorTokenVerification
	}

	return id, nil
}

func (s *UseCase) GenerateToken(id string) (string, error) {
	return s.tokenGenerator.GenerateToken(id)
}
