package service

import (
	"log"
	 
    "github.com/gin-gonic/gin"

    "github.com/kkjoker/recipe/library/rdbms"
    "github.com/kkjoker/recipe/model"
)

type Service struct{}

// GetAll 全ユーザ取得
func (s Service) GetAll() ([]*model.User, error) {
    
	u := []*model.User{}
	
	db := rdbms.GetDB()
	
	db = db.Find(&u)

	return u, rdbms.GetDBError(db)
}

// CreateModel User作成
func (s Service) CreateModel(c *gin.Context) (*model.User, error) {
    
    u := model.User{}

    if err := c.BindJSON(&u); err != nil {
        return &u, err
    }

	db := rdbms.GetDB()
	
	log.Println(u.Name)
	log.Println("---1---")
	db = db.Create(&u)

    return &u, rdbms.GetDBError(db)
}

// GetByID User取得
func (s Service) GetByID(id string) (*model.User, error) {
	
	u := model.User{}
	
	db := rdbms.GetDB()

	if db == nil {
		log.Println("DBが空です")
	}
	log.Println("---2---")
	log.Println(id)
	db = db.Where("id = ?", id)
	db = db.First(&u)

    return &u, rdbms.GetDBError(db)
}

// UpdateByID User更新
func (s Service) UpdateByID(id string, c *gin.Context) (*model.User, error) {
    
    u := model.User{}

	if err := c.BindJSON(&u); err != nil {
        return &u, err
    }

	db := rdbms.GetDB()
	
	db = db.Where("id = ?", id)
    db = db.Save(&u)

    return &u, rdbms.GetDBError(db)
}

// DeleteByID User削除
func (s Service) DeleteByID(id string) error {
    
    u := model.User{}

	db := rdbms.GetDB()

	db = db.Where("id = ?", id)
	db = db.Delete(&u)

    return rdbms.GetDBError(db)
}