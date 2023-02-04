package models

type Outlet struct {
	IdOutlet   int    `gorm:"primaryKey;autoIncrement" json:"idoutlet"`
	Code       string `json:"code" binding:"required,startswith=LD,len=5"`
	NameOutlet string `json:"nameoutlet" binding:"required,min=3"`
	Status     string `json:"status" binding:"required"`
	Address    string `json:"address" binding:"required,max=100"`
	Phone      string `json:"phone" binding:"required"`
	User       []User `gorm:"foreignKey:outlet_id"`
}

type Customer struct {
	IdCustomer      int    `gorm:"primaryKey;autoIncrement" json:"idcustomer"`
	Nik             string `json:"nik" binding:"required,max=16"`
	NameCustomer    string `json:"namecustomer" binding:"required,min=3"`
	AddressCustomer string `json:"addresscustomer" binding:"required,max=100"`
	PhoneCustomer   string `json:"phonecustomer" binding:"required,startswith=+62"`
}

type LaundryType struct {
	IdType        int             `gorm:"primaryKey;autoIncrement" json:"idtype"`
	NameType      string          `json:"nametype" binding:"required,min=6"`
	LaundryPrices []LaundryPrices `gorm:"foreignKey:type_id"`
}

type User struct {
	IdUser        int             `gorm:"primaryKey;autoIncrement" json:"iduser"`
	NameUser      string          `json:"nameuser" binding:"required,min=3"`
	Email         string          `json:"email" binding:"required,email"`
	Password      string          `json:"password" binding:"required,min=6"`
	Role          string          `json:"role" binding:"required"`
	OutletId      int             `json:"outletid" binding:"required"`
	LaundryPrices []LaundryPrices `gorm:"foreignKey:user_id"`
	Transaction   []Transaction   `gorm:"foreignKey:user_id"`
}

type LaundryPrices struct {
	IdPrice   int    `gorm:"primaryKey;autoIncrement" json:"idprice"`
	NamePrice string `json:"nameprice" binding:"required,min=3"`
	UnitType  int    `json:"unittype" binding:"required,min=1"`
	Price     uint   `json:"price" binding:"required,gte=10000,lte=1000000"`
	TypeId    int    `json:"typeid" binding:"required" `
	UserId    int    `json:"userid" binding:"required" `
}

type Transaction struct {
	IdTransaction     int    `gorm:"primaryKey;autoIncrement" json:"idtransaction"`
	CustomerId        int    `json:"customerid"`
	UserID            int    `json:"userID" `
	Amount            int    `json:"amount" binding:"required"`
	StartDate         string `json:"startdate" binding:"required"`
	EndDate           string `json:"enddate" binding:"required"`
	StatusTransaction string `json:"statustransaction"`
}
