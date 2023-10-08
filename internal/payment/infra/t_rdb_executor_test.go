package infra

import (
	"os"
	"testing"
	"upsidr-coding-test/internal/rdb"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestNewExecutorRDB(t *testing.T) {
	tests := []struct {
		name    string
		wantLen int
		wantCap int
	}{
		{
			name:    "Newで作成された場合",
			wantLen: 0,
			wantCap: 5,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			executor := NewExecutorRDB()
			assert.Equal(t, tt.wantLen, len(executor.funcList))
			assert.Equal(t, tt.wantCap, cap(executor.funcList))
		})
	}
}

func TestExecutorRDB_Append(t *testing.T) {
	type args struct {
		fnList []ExecFunc
	}

	tests := []struct {
		name    string
		args    args
		wantLen int
		wantCap int
	}{
		{
			name:    "3つSQL実行関数を与えられた場合",
			wantLen: 3,
			wantCap: 5,
			args: args{fnList: []ExecFunc{
				func(tx *gorm.DB) error {
					u := rdb.User{
						UserID:    "user99990",
						CompanyID: "company99990",
						Name:      "user99990",
						Email:     "email1@test.com",
						Password:  "asdasdasd",
					}
					return tx.
						Session(&gorm.Session{DryRun: true}).
						Create(&u).Error
				},
				func(tx *gorm.DB) error {
					u := rdb.User{
						UserID:    "user99991",
						CompanyID: "company99991",
						Name:      "user99991",
						Email:     "email2@test.com",
						Password:  "asdasdasd",
					}
					return tx.
						Session(&gorm.Session{DryRun: true}).
						Create(&u).Error
				},
				func(tx *gorm.DB) error {
					u := rdb.User{
						UserID:    "user99993",
						CompanyID: "company99993",
						Name:      "user99993",
						Email:     "email3@test.com",
						Password:  "asdasdasd",
					}
					return tx.
						Session(&gorm.Session{DryRun: true}).
						Create(&u).Error
				},
			}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			executor := NewExecutorRDB()
			executor.Append(tt.args.fnList...)
			assert.Equal(t, tt.wantLen, len(executor.funcList))
			assert.Equal(t, tt.wantCap, cap(executor.funcList))
		})
	}
}

func TestExecutorRDB_Exec(t *testing.T) {
	type args struct {
		fnList []ExecFunc
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "3つSQL実行関数を与えられた場合",
			wantErr: false,
			args: args{fnList: []ExecFunc{
				func(tx *gorm.DB) error {
					u := rdb.User{
						UserID:    "user99990",
						CompanyID: "company99990",
						Name:      "user99990",
						Email:     "email1@test.com",
						Password:  "asdasdasd",
					}
					return tx.
						Session(&gorm.Session{DryRun: true}).
						Create(&u).Error
				},
				func(tx *gorm.DB) error {
					u := rdb.User{
						UserID:    "user99991",
						CompanyID: "company99991",
						Name:      "user99991",
						Email:     "email2@test.com",
						Password:  "asdasdasd",
					}
					return tx.
						Session(&gorm.Session{DryRun: true}).
						Create(&u).Error
				},
				func(tx *gorm.DB) error {
					u := rdb.User{
						UserID:    "user99993",
						CompanyID: "company99993",
						Name:      "user99993",
						Email:     "email3@test.com",
						Password:  "asdasdasd",
					}
					return tx.
						Session(&gorm.Session{DryRun: true}).
						Create(&u).Error
				},
			}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := initDB(false)
			assert.Equal(t, nil, err)
			defer rdb.Close()

			executor := NewExecutorRDB()
			executor.Append(tt.args.fnList...)
			err = executor.Exec()
			assert.Equal(t, tt.wantErr, (err != nil))
		})
	}
}

func initDB(isDebugMode bool) error {
	return rdb.Init(
		isDebugMode,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
