package service

import(
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"github.com/unrolled/render"
	"time"
	"strconv"
)

type UserInfo struct {
    Id        int64  
    UserName   string
    DepartName string
	CreateAt   *time.Time
}

var engine, _ = xorm.NewEngine("mysql", "root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true");

func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}



func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	
		return func(w http.ResponseWriter, req *http.Request) {
			req.ParseForm()
			if len(req.Form["username"][0]) == 0 {
				formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
				return
			}

			has, _ := engine.Exist(new(UserInfo))
			if !has {
				_ = engine.Sync2(new(UserInfo))
			} 
			u := NewUserInfo(UserInfo{UserName: req.Form["username"][0]})
			u.DepartName = req.Form["departname"][0]
			//entities.UserInfoService.Save(u)
			_, _ = engine.Insert(u)
			formatter.JSON(w, http.StatusOK, u)
		}
	}
	
	func getUserInfoHandler(formatter *render.Render) http.HandlerFunc {
	
		return func(w http.ResponseWriter, req *http.Request) {
			req.ParseForm()
			if len(req.Form["userid"][0]) != 0 {
				has, _ := engine.Exist(new(UserInfo))
				if !has {
					_ = engine.Sync2(new(UserInfo))
				}
				i, _ := strconv.ParseInt(req.Form["userid"][0], 10, 32)
				user := UserInfo{}
				_, _ = engine.Where("Id = ?", i).Get(&user)
				formatter.JSON(w, http.StatusOK, user)	
			}
		}
	

	}