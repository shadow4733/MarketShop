package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserSeedData struct {
	ID                  uuid.UUID
	Username            string
	Email               string
	PasswordHash        string
	Phone               string
	Balance             float64
	TotalSpent          float64
	TotalEarned         float64
	FirstName           string
	LastName            string
	MiddleName          string
	DateOfBirth         string
	AvatarURL           *string
	Country             string
	City                string
	Address             string
	PostalCode          string
	IsVerified          bool
	IsActive            bool
	IsSeller            bool
	Rating              float64
	ReviewCount         int
	LastPasswordChange  *time.Time
	FailedLoginAttempts int
	IsLocked            bool
	EmailNotifications  bool
	SMSNotifications    bool
	PushNotifications   bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	LastLoginAt         *time.Time
	DeletedAt           *time.Time
}

var Users = []UserSeedData{
	{
		ID: uuid.New(), Username: "alice", Email: "alice@example.com", PasswordHash: "password123", Phone: "+1234567890",
		Balance: 100, TotalSpent: 0, TotalEarned: 0, FirstName: "Alice", LastName: "Smith", MiddleName: "M",
		DateOfBirth: "1990-01-01", AvatarURL: nil, Country: "USA", City: "New York", Address: "123 Main St",
		PostalCode: "10001", IsVerified: true, IsActive: true, IsSeller: false, Rating: 4.5, ReviewCount: 10,
		LastPasswordChange: nil, FailedLoginAttempts: 0, IsLocked: false, EmailNotifications: true, SMSNotifications: false,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "bob", Email: "bob@example.com", PasswordHash: "securePass!23", Phone: "+19876543210",
		Balance: 150, TotalSpent: 50, TotalEarned: 20, FirstName: "Bob", LastName: "Johnson", MiddleName: "K",
		DateOfBirth: "1985-05-15", AvatarURL: nil, Country: "USA", City: "Los Angeles", Address: "456 Sunset Blvd",
		PostalCode: "90001", IsVerified: true, IsActive: true, IsSeller: true, Rating: 4.8, ReviewCount: 25,
		LastPasswordChange: nil, FailedLoginAttempts: 1, IsLocked: false, EmailNotifications: true, SMSNotifications: true,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "charlie", Email: "charlie@example.com", PasswordHash: "charlie123", Phone: "+1122334455",
		Balance: 200, TotalSpent: 75, TotalEarned: 50, FirstName: "Charlie", LastName: "Brown", MiddleName: "L",
		DateOfBirth: "1992-07-20", AvatarURL: nil, Country: "USA", City: "Chicago", Address: "789 Lake Shore Dr",
		PostalCode: "60601", IsVerified: false, IsActive: true, IsSeller: false, Rating: 4.2, ReviewCount: 5,
		LastPasswordChange: nil, FailedLoginAttempts: 0, IsLocked: false, EmailNotifications: true, SMSNotifications: false,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "diana", Email: "diana@example.com", PasswordHash: "dianaPass!45", Phone: "+1098765432",
		Balance: 120, TotalSpent: 20, TotalEarned: 30, FirstName: "Diana", LastName: "Prince", MiddleName: "A",
		DateOfBirth: "1988-03-10", AvatarURL: nil, Country: "USA", City: "San Francisco", Address: "321 Market St",
		PostalCode: "94105", IsVerified: true, IsActive: true, IsSeller: true, Rating: 4.9, ReviewCount: 40,
		LastPasswordChange: nil, FailedLoginAttempts: 0, IsLocked: false, EmailNotifications: true, SMSNotifications: true,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "edward", Email: "edward@example.com", PasswordHash: "edwardPass!78", Phone: "+1222333444",
		Balance: 80, TotalSpent: 10, TotalEarned: 15, FirstName: "Edward", LastName: "Norton", MiddleName: "P",
		DateOfBirth: "1995-12-12", AvatarURL: nil, Country: "USA", City: "Miami", Address: "654 Ocean Dr",
		PostalCode: "33139", IsVerified: false, IsActive: true, IsSeller: false, Rating: 4.0, ReviewCount: 3,
		LastPasswordChange: nil, FailedLoginAttempts: 2, IsLocked: false, EmailNotifications: false, SMSNotifications: false,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "fiona", Email: "fiona@example.com", PasswordHash: "fiona1234", Phone: "+1444555666",
		Balance: 90, TotalSpent: 5, TotalEarned: 10, FirstName: "Fiona", LastName: "Apple", MiddleName: "G",
		DateOfBirth: "1991-09-05", AvatarURL: nil, Country: "USA", City: "Seattle", Address: "987 Pine St",
		PostalCode: "98101", IsVerified: true, IsActive: true, IsSeller: false, Rating: 4.3, ReviewCount: 8,
		LastPasswordChange: nil, FailedLoginAttempts: 0, IsLocked: false, EmailNotifications: true, SMSNotifications: true,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "george", Email: "george@example.com", PasswordHash: "georgePass!90", Phone: "+1555666777",
		Balance: 60, TotalSpent: 0, TotalEarned: 0, FirstName: "George", LastName: "Martin", MiddleName: "H",
		DateOfBirth: "1982-11-11", AvatarURL: nil, Country: "USA", City: "Boston", Address: "111 Beacon St",
		PostalCode: "02108", IsVerified: true, IsActive: false, IsSeller: true, Rating: 4.7, ReviewCount: 20,
		LastPasswordChange: nil, FailedLoginAttempts: 3, IsLocked: false, EmailNotifications: true, SMSNotifications: true,
		PushNotifications: false, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "hannah", Email: "hannah@example.com", PasswordHash: "hannahPass!12", Phone: "+1666777888",
		Balance: 110, TotalSpent: 25, TotalEarned: 40, FirstName: "Hannah", LastName: "Montana", MiddleName: "I",
		DateOfBirth: "1993-06-22", AvatarURL: nil, Country: "USA", City: "Austin", Address: "222 Congress Ave",
		PostalCode: "73301", IsVerified: false, IsActive: true, IsSeller: false, Rating: 4.1, ReviewCount: 7,
		LastPasswordChange: nil, FailedLoginAttempts: 0, IsLocked: false, EmailNotifications: true, SMSNotifications: false,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "ian", Email: "ian@example.com", PasswordHash: "ianPass!34", Phone: "+1777888999",
		Balance: 130, TotalSpent: 40, TotalEarned: 35, FirstName: "Ian", LastName: "Curtis", MiddleName: "J",
		DateOfBirth: "1989-02-28", AvatarURL: nil, Country: "USA", City: "Denver", Address: "333 Colfax Ave",
		PostalCode: "80202", IsVerified: true, IsActive: true, IsSeller: true, Rating: 4.6, ReviewCount: 15,
		LastPasswordChange: nil, FailedLoginAttempts: 1, IsLocked: false, EmailNotifications: true, SMSNotifications: true,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
	{
		ID: uuid.New(), Username: "julia", Email: "julia@example.com", PasswordHash: "juliaPass!56", Phone: "+1888999000",
		Balance: 95, TotalSpent: 10, TotalEarned: 5, FirstName: "Julia", LastName: "Roberts", MiddleName: "K",
		DateOfBirth: "1994-08-30", AvatarURL: nil, Country: "USA", City: "San Diego", Address: "444 Ocean Blvd",
		PostalCode: "92101", IsVerified: true, IsActive: true, IsSeller: false, Rating: 4.4, ReviewCount: 12,
		LastPasswordChange: nil, FailedLoginAttempts: 0, IsLocked: false, EmailNotifications: true, SMSNotifications: true,
		PushNotifications: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), LastLoginAt: nil, DeletedAt: nil,
	},
}

//func SeedUsers(db *gorm.DB) error {
//	for _, u := range Users {
//		newUser :=
//	}
//}
