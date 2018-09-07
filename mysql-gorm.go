package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

var tableName = "test"

func CreateTable(db *gorm.DB) {
	if db == nil {
		return
	}

	//db.Table("users").CreateTable(&Animal{Name: "shuai", ID: 666, Age: 30})
	r := db.HasTable(tableName)
	if r == true {
		return
	}

	db.Table(tableName).CreateTable(&Animal{})
	//db.SingularTable()
}

func InsertData(i interface{}) {
	if i == nil {
		return
	}
	db := i.(*gorm.DB).Table(tableName)

	user := []Animal{
		{ID: 1, Name: "song", Age: 22},
		{ID: 2, Name: "wei", Age: 23},
		{ID: 3, Name: "shuai", Age: 24},
		{ID: 4, Name: "songweishuai", Age: 30},
		{ID: 5, Name: "shuai", Age: 25},
	}

	for i, u := range user {
		fmt.Println("i=", i, ",u=", u)
		db.Create(&u)
	}
}

func QueryData(i interface{}) {
	db := i.(*gorm.DB).Table(tableName)

	var user Animal
	db.First(&user)
	fmt.Println("first data is:", user)

	var user1 Animal
	db.Take(&user1)
	fmt.Println("take data is:", user1)

	var user2 Animal
	db.Last(&user2)
	fmt.Println("last data is:", user2)

	//slice and pointer
	var userAll []Animal
	db.Find(&userAll)
	fmt.Println("userAll is:", userAll)

	var user3 Animal
	db.First(&user3, 3)
}

func WhereData(db *gorm.DB) {
	var user Animal
	db.Table(tableName).Where("name=?", "shuai").First(&user)
	fmt.Println("where first user is:", user)

	var users []Animal
	db.Table(tableName).Where("name=?", "shuai").Find(&users)
	fmt.Println("where all user is:", users)

	var usersNot []Animal
	db.Table(tableName).Where("name!=?", "shuai").Find(&usersNot)
	fmt.Println("where find name!=shuai", usersNot)

	var usersIn []Animal
	db.Table(tableName).Where("name in (?)", []string{"song", "shuai"}).Find(&usersIn)
	fmt.Println("select * from test where name in (song,shuai);", usersIn)

	var usersLike []Animal
	db.Table(tableName).Where("name like ?", "%shuai%").Find(&usersLike)
	fmt.Println("select * from test where name like \"%shuai%\"", usersLike)

	var usersAnd []Animal
	db.Table(tableName).Where("name =? and age =?", "shuai", 25).Find(&usersAnd)
	fmt.Println("select * from test where name=shuai and age=25", usersAnd)

	var user1 Animal
	db.Table(tableName).Where(&Animal{Name: "shuai"}).First(&user1)
	fmt.Println("&Animal{Name:shuai} first is:", user1)

	//struct use value copy should &struct
	var usersObject []Animal
	db.Table(tableName).Where(&Animal{Name: "shuai"}).Find(&usersObject)
	fmt.Println("&Animal{Name:shuai} find is", usersObject)

	//map and slice use pointer copy
	var usersMap []Animal
	db.Table(tableName).Where(map[string]interface{}{"Name": "shuai", "age": 25}).Find(&usersMap)
	fmt.Println("&map[string]interface{}{Name:shuai,age:25}", users)

	var usersSlice []Animal
	db.Table(tableName).Where([]int{2, 3}).Find(&usersSlice)
	fmt.Println("Where(&[]int{2,3})", usersSlice)
}

func NotData(db *gorm.DB) {
	var user Animal
	db.Table(tableName).Not("name=?", "shuai").First(&user)
	fmt.Println("where first user is:", user)

	var users []Animal
	db.Table(tableName).Not("name=?", "shuai").Find(&users)
	fmt.Println("where all user is:", users)

	//var usersNot []Animal
	//db.Table(tableName).Not("name!=?", "shuai").Find(&usersNot)
	//fmt.Println("where find name!=shuai", usersNot)
	//
	//var usersIn []Animal
	//db.Table(tableName).Not("name in (?)", []string{"song", "shuai"}).Find(&usersIn)
	//fmt.Println("select * from test where name in (song,shuai);", usersIn)
	//
	//var usersLike []Animal
	//db.Table(tableName).Not("name like ?", "%shuai%").Find(&usersLike)
	//fmt.Println("select * from test where name like \"%shuai%\"", usersLike)
	//
	//var usersAnd []Animal
	//db.Table(tableName).Not("name =? and age =?", "shuai", 25).Find(&usersAnd)
	//fmt.Println("select * from test where name=shuai and age=25", usersAnd)
	//
	//var user1 Animal
	//db.Table(tableName).Not(&Animal{Name:"shuai"}).First(&user1)
	//fmt.Println("&Animal{Name:shuai} first is:",user1)
	//
	////struct use value copy should &struct
	//var usersObject []Animal
	//db.Table(tableName).Not(&Animal{Name:"shuai"}).Find(&usersObject)
	//fmt.Println("&Animal{Name:shuai} find is",usersObject)
	//
	////map and slice use pointer copy
	//var usersMap []Animal
	//db.Table(tableName).Not(map[string]interface{}{"Name":"shuai","age":25}).Find(&usersMap)
	//fmt.Println("&map[string]interface{}{Name:shuai,age:25}",users)
	//
	//var usersSlice []Animal
	//db.Table(tableName).Not([]int{2,3}).Find(&usersSlice)
	//fmt.Println("Where(&[]int{2,3})",usersSlice)
}

func FirstOrInitData(i interface{}) {
	var user Animal
	db:=i.(*gorm.DB).Table(tableName)
	db.FirstOrInit(&user,Animal{Name:"shuai"})
	fmt.Println("FirstOrInit user is",user)
}

func AttrsData(i interface{})  {
	var user Animal
	db:=i.(*gorm.DB).Table(tableName)
	db.Where(&Animal{Name:"shuaige"}).Attrs(Animal{ID:8,Age:10}).FirstOrInit(&user)

	fmt.Println("attrs is ",user)
}

func UpdateData(db *gorm.DB) {
	var user Animal
	db.Table(tableName).First(&user)
	fmt.Println(user)

	db.Table(tableName).Model(&user).Update(Animal{Name: "song", Age: 27})
	fmt.Println(user)
}

func DeleteOrmData(db *gorm.DB) {
	var user Animal
	//db.Table(tableName).First(&user)

	//delete all record
	db.Table(tableName).Delete(&user)
}

func DeleteTable(db *gorm.DB) {
	if db == nil {
		return
	}
	db.DropTable(tableName)
}

var query = flag.Int("query", 0, "quer data?")
var insert = flag.Bool("insert", false, "insert data?")
var update = flag.Bool("name", false, "update data?")
var dropData = flag.Bool("dropdata", false, "delete record ?")
var dropTable = flag.String("tablename", "", "which table delete?")

const (
	QUERYDATA = iota
	QUERYDATA_QUERY
	QUERYDATA_WHERE
	QUERYDATA_NOT
	QUERYDATA_OR
	QUERYDATA_INLINE
	QUERYDATA_FIRST_OR_INIT
	QUERYDATA_ATTRS
	QUERYDATA_ASSIGN
	QUERYDATA_FIRST_OR_CREATE
	QUERYDATA_SUB_QUERY
	QUERYDATA_SELECT
	QUERYDATA_ORDER
	QUERYDATA_LIMIT
	QUERYDATA_OFFSET
	QUERYDATA_COUNT
	QUERYDATA_PLUCK
	QUERYDATA_SCAN
)

func main() {
	//parse args
	flag.Parse()

	//print args
	fmt.Printf("query=%v,insert=%v,update=%v,dropData=%v,dropTable=%v\n", *query, *insert, *update, *dropData, *dropTable)

	//connect databases
	db, err := gorm.Open("mysql", "root:123456@/mysql")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	CreateTable(db)

	if *insert {
		InsertData(db)
	}

	switch *query {
	case QUERYDATA_QUERY:
		QueryData(db)
	case QUERYDATA_WHERE:
		WhereData(db)
	case QUERYDATA_NOT:
		NotData(db)
	case QUERYDATA_OR:
	case QUERYDATA_INLINE:
	case QUERYDATA_FIRST_OR_INIT:
		FirstOrInitData(db)
	case QUERYDATA_ATTRS:
		AttrsData(db)
	case QUERYDATA_ASSIGN:
	case QUERYDATA_FIRST_OR_CREATE:
	case QUERYDATA_SUB_QUERY:
	case QUERYDATA_SELECT:
	case QUERYDATA_ORDER:
	case QUERYDATA_LIMIT:
	case QUERYDATA_OFFSET:
	case QUERYDATA_COUNT:
	case QUERYDATA_PLUCK:
	case QUERYDATA_SCAN:
	case QUERYDATA:
		fallthrough
	default:

	}

	if *update {
		UpdateData(db)
	}

	if *dropData {
		DeleteOrmData(db)
	}

	if *dropTable != "" {
		DeleteTable(db)
	}
}
