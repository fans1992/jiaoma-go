package seeders

import (
	"fmt"
	"github.com/fans1992/jiaoma/database/factories"
	"github.com/fans1992/jiaoma/pkg/console"
	"github.com/fans1992/jiaoma/pkg/logger"
	"github.com/fans1992/jiaoma/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedLinksTable", func(db *gorm.DB) {

		links := factories.MakeLinks(5)

		result := db.Table("links").Create(&links)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
