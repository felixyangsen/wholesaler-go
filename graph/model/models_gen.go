// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CommodityOps struct {
	Create *Comodity `json:"create"`
	Update *Comodity `json:"update"`
}

type ComodityImage struct {
	ComodityID int `json:"comodity_id"`
	ImageID    int `json:"image_id"`
}

type ComodityPagination struct {
	Limit     *int        `json:"limit"`
	Page      *int        `json:"page"`
	TotalItem int         `json:"total_item"`
	Nodes     []*Comodity `json:"nodes"`
}

type ComodityWithCategory struct {
	Limit      *int        `json:"limit"`
	Page       *int        `json:"page"`
	CategoryID int         `json:"category_id"`
	Category   *Category   `json:"category"`
	TotalItem  int         `json:"total_item"`
	Nodes      []*Comodity `json:"nodes"`
}

type EditSchedule struct {
	ID            string    `json:"id"`
	ScheduleName  string    `json:"schedule_name"`
	CommodityName string    `json:"commodity_name"`
	DealedUnit    string    `json:"dealed_unit"`
	StartDate     string    `json:"start_date"`
	EndDate       string    `json:"end_date"`
	Day           []*string `json:"day"`
	StartTime     string    `json:"start_time"`
	EndTime       string    `json:"end_time"`
}

type EditUser struct {
	Email      string   `json:"email"`
	Whatsapp   string   `json:"whatsapp"`
	Avatar     string   `json:"avatar"`
	LookingFor []string `json:"looking_for"`
}

type Friend struct {
	Username string `json:"username"`
	User     *User  `json:"user"`
}

type FriendOps struct {
	Add *User `json:"add"`
}

type Image struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
	Link string `json:"link"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}

type LookingFor struct {
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
}

type NewComodity struct {
	Name        string   `json:"name"`
	UnitPrice   float64  `json:"unit_price"`
	UnitType    string   `json:"unit_type"`
	MinPurchase int      `json:"min_purchase"`
	Description *string  `json:"description"`
	CategoryID  int      `json:"category_id"`
	Images      []string `json:"images"`
}

type NewImage struct {
	Path string `json:"path"`
	Link string `json:"link"`
}

type NewSchedule struct {
	ScheduleName          string    `json:"schedule_name"`
	CommodityName         string    `json:"commodity_name"`
	DealedUnit            string    `json:"dealed_unit"`
	StartDate             string    `json:"start_date"`
	EndDate               string    `json:"end_date"`
	Day                   []*string `json:"day"`
	StartTime             string    `json:"start_time"`
	EndTime               string    `json:"end_time"`
	InvolvedUsersUsername []*string `json:"involved_users_username"`
}

type NewUser struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	Whatsapp        string `json:"whatsapp"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Schedule struct {
	ID                    int       `json:"id"`
	ScheduleName          string    `json:"schedule_name"`
	CommodityName         string    `json:"commodity_name"`
	DealedUnit            string    `json:"dealed_unit"`
	StartDate             string    `json:"start_date"`
	EndDate               string    `json:"end_date"`
	Day                   []*string `json:"day"`
	StartTime             string    `json:"start_time"`
	EndTime               string    `json:"end_time"`
	InvolvedUsersUsername []*string `json:"involved_users_username"`
	InvolvedUsers         []*User   `json:"involved_users"`
}

type ScheduleOps struct {
	Create *Schedule `json:"create"`
	Update *Schedule `json:"update"`
	Delete bool      `json:"delete"`
}

type UserOps struct {
	Register   *LoginResponse `json:"register"`
	Login      *LoginResponse `json:"login"`
	Update     *User          `json:"update"`
	DeleteUser *bool          `json:"deleteUser"`
}
