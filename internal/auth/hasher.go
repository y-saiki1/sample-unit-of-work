package auth

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	hash []byte
}

func NewHasher() Hasher {
	return Hasher{}
}

func (h *Hasher) Hashed(raw string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	h.hash = hashed
	return nil
}

func (h *Hasher) Verify(rawPassword string) error {
	return bcrypt.CompareHashAndPassword(h.hash, []byte(rawPassword))
}

func (h *Hasher) HashString() string {
	return string(h.hash)
}
