package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

//var out []UserDataOutput

func GetUsers(context *gin.Context) {
	//out := []UserDataOutput{}
	users := dao.GetAllUsersData()
	//out := convertEntityToValidJSON(users)
	//for _, v := range out{
	//	fmt.Println(v.ToString())
	//}

	//json, _ := json2.Marshal(out)
	//fmt.Println(json)
	//context.JSON(200, string(json))

	context.JSON(200, users)
	//context.JSON(200, out)

	//for _, v := range users{
	//	fmt.Println(v.ToString())
	//}
	//json, _ := json2.Marshal(out[0])
	//fmt.Println(json)
	//context.JSON(200, string(json))

	//outRefObjValue := reflect.ValueOf(out)
	//fmt.Printf("out value is %s\n", outRefObjValue)
	//fmt.Printf("can out value changes %s\n", outRefObjValue.CanSet())
	//outRefObjType := reflect.TypeOf(out)
	//fmt.Printf("out type is %s\n", outRefObjType)
	//outPointRefObjValue := reflect.ValueOf(&out)
	//fmt.Printf("pointer out value is %s\n", outPointRefObjValue)
	//fmt.Printf("can pointer out value changes %s\n", outPointRefObjValue.CanSet())
	//fmt.Printf("can pointer out value changes to elem%s\n", outPointRefObjValue.Elem().CanSet())
	//outPointRefObjType := reflect.TypeOf(&out)
	//fmt.Printf("pointer out type is %s\n\n", outPointRefObjType)
	//
	//usersRefObjValue := reflect.ValueOf(users)
	//fmt.Printf("users value is %s\n", usersRefObjValue)
	//fmt.Printf("can users value changes %s\n", usersRefObjValue.CanSet())
	//usersRefObjType := reflect.TypeOf(users)
	//fmt.Printf("users type is %s\n", usersRefObjType)
	//usersPointRefObjValue := reflect.ValueOf(&users)
	//fmt.Printf("pointer users value is %s\n", usersPointRefObjValue)
	//fmt.Printf("can pointer users value changes %s\n", usersPointRefObjValue.CanSet())
	//fmt.Printf("can pointer users value changes to elem%s\n", usersPointRefObjValue.Elem().CanSet())
	//usersPointRefObjType := reflect.TypeOf(&users)
	//fmt.Printf("pointer users type is %s\n\n", usersPointRefObjType)
}

func GetUser(context *gin.Context) {
	stringID := context.Param("id")

	id, err := strconv.Atoi(stringID)
	if err != nil {
		log.Println(err)
	}
	user := dao.GetUserDataById(id)
	context.JSON(200, user)
}

func PutState(context *gin.Context) {
	//var input UserDataInput
	//fmt.Println("1")
	if err := context.ShouldBindJSON(apiserver.); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("2")
	//id, _ := strconv.Atoi(input.Id)
	//user := entity.UserData{Name: input.Name, Surname: input.Surname,
	//Login: input.Login, Email: input.Email, BirthDate: input.BirthDate}

	fmt.Println(user)

	//dao.UpdateUserDataById(&user)
}

func DeleteUser(context *gin.Context) {
	stringID := context.Param("id")

	id, _ := strconv.Atoi(stringID)
	dao.DropUserData(id)
}
