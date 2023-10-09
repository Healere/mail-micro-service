package controller

import (
	"authentication/cmd/api/common"
	"authentication/cmd/api/dto"
	"authentication/cmd/api/model"
	"authentication/cmd/api/response"
	"authentication/cmd/api/util"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func Login(c *gin.Context) {
	db := common.GetDB()
	var requestUser = model.User{}
	// json.NewDecoder(c.Request.Body).Decode(&requestUser)
	c.Bind(&requestUser) //和上面1行代码一样
	mail := requestUser.Email
	password := requestUser.Password

	//justify password
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Password should not less than 6 digits")
		return
	}
	//justify mail
	if len(mail) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Email address should not less than 6 digits")
		return
	}

	var user model.User
	db.Where("email=?", mail).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "User does not exist")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "Password error")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "System abnormally")
		log.Printf("token generate error : %v", err)
		return
	}

	response.Success(c, gin.H{"token": token}, "Sign in successful")
}

func Register(c *gin.Context) {
	db := common.GetDB()
	var requestUser = model.User{}
	fmt.Println(c.Request.Body)
	json.NewDecoder(c.Request.Body).Decode(&requestUser)
	c.ShouldBind(&requestUser) //和上面1行代码一样
	name := requestUser.Name   //c.PostForm("name")
	mail := requestUser.Email
	password := requestUser.Password
	// name := c.PostForm("name")
	// mail := c.PostForm("email")
	// password := c.PostForm("password")
	fmt.Println("name: " + name + "email: " + mail + "password" + password)
	//justify name
	if len(name) == 0 {
		name = util.RandomString(6)
	}
	//justify password
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Password should not less than 6 digits")
		return
	}
	//justify mail
	if len(mail) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "Email address should not less than 6 digits")
		return
	}

	log.Println(name, mail, password)
	//is mail valid
	if isMailExist(db, mail) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "The email is used")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "Encryption error")
		return
	}

	//user is not exist
	newUser := model.User{
		Name:     name,
		Email:    mail,
		Password: string(hashedPassword),
	}
	err = db.Create(&newUser).Error
	if err != nil {
		log.Println("插入失败！", err)
	}
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "System abnormally")
		log.Printf("token generate error : %v", err)
		return
	}
	response.Success(c, gin.H{"token": token}, "Sign up successful")
}

// is mail valid
func isMailExist(db *gorm.DB, mail string) bool {
	var user model.User

	err := db.Where("email=?", mail).First(&user).Error
	if err != nil {
		log.Println("查询失败！", err)
	}
	return user.ID != 0
}
