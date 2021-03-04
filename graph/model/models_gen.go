// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CommodityOps struct {
	Create *Comodity `json:"create"`
	Update *Comodity `json:"update"`
}

type Comodity struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Image       []*string `json:"image"`
	UnitPrice   string    `json:"unit_price"`
	UnitType    string    `json:"unit_type"`
	MinPurchase string    `json:"min_purchase"`
	Description *string   `json:"description"`
	Username    string    `json:"username"`
	User        *User     `json:"user"`
}

type ComodityPagination struct {
	Limit     *int        `json:"limit"`
	Page      *int        `json:"page"`
	TotalItem int         `json:"total_item"`
	Nodes     []*Comodity `json:"nodes"`
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
	Email          string   `json:"email"`
	WhatsappNumber string   `json:"whatsapp_number"`
	ProfileImage   string   `json:"profile_image"`
	LookingFor     []string `json:"looking_for"`
}

type Friend struct {
	Username string `json:"username"`
	User     *User  `json:"user"`
}

type FriendOps struct {
	Add *User `json:"add"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	User        *User  `json:"user"`
}

type NewComodity struct {
	Name        string    `json:"name"`
	MinPurchase string    `json:"min_purchase"`
	UnitType    string    `json:"unit_type"`
	UnitPrice   string    `json:"unit_price"`
	Description string    `json:"description"`
	Images      []*string `json:"images"`
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
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	Role            string  `json:"role"`
	WhatsappNumber  *string `json:"whatsapp_number"`
	Password        string  `json:"password"`
	ConfirmPassword string  `json:"confirm_password"`
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

type User struct {
	ID         int         `json:"id"`
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Role       string      `json:"role"`
	Whatsapp   string      `json:"whatsapp"`
	Avatar     *string     `json:"avatar"`
	Friends    []*User     `json:"friends"`
	LookingFor []string    `json:"looking_for"`
	Comodity   []*Comodity `json:"comodity"`
}

type UserOps struct {
	Register   *LoginResponse `json:"register"`
	Login      *LoginResponse `json:"login"`
	Update     *User          `json:"update"`
	DeleteUser *bool          `json:"deleteUser"`
}
