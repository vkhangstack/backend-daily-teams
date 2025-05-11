package benchmark

import (
	"fmt"
	"github.com/vkhangstack/dlt/internal/core/domain/model"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/vkhangstack/dlt/internal/adapters/repository"
)

func BenchmarkCreateUser(b *testing.B) {
	db, err := gorm.Open("postgres", "postgres://root:khangdev@localhost:5432/opulentila?sslmode=disable")
	if err != nil {
		panic(err)
	}

	//redisCache, err := cache.NewRedisCache("localhost:6379", "")
	//if err != nil {
	//	panic(err)
	//}

	store := repository.NewDB(db)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		email := fmt.Sprintf("test_user_%d@example.com", i)
		// Delete user if it exists
		user := &model.User{}
		if err := db.Where("email = ?", email).First(&user).Error; err == nil {
			if err := db.Delete(&user).Error; err != nil {
				b.Fatalf("failed to delete user: %v", err)
			}
		}

		_, err := store.CreateUser(user)
		if err != nil {
			b.Fatalf("failed to create test user: %v", err)
		}
	}
}
