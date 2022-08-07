package customer

type Customer struct {
	ID            uint   `json:"id"`
	Nama          string `json:"nama"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gender        string `json:"gender"`
	ID_Hobi       int    `json:"id_hobi"`
	ID_Gender     int    `json:"id_gender"`
}

type Detail_customer struct {
	ID            uint    `gorm:"primarykey"`
	Nama          string  `json:"nama"`
	Tanggal_lahir string  `json:"tanggal_lahir"`
	Gender        string  `json:"gender"`
	ID_Hobi       int     `json:"id_hobi"`
	Hobi          Hobi    `gorm:"foreignkey:ID;references:ID_Hobi"`
	ID_Jurusan    int     `json:"id_jurusan"`
	Jurusan       Jurusan `gorm:"foreignkey:ID;references:ID_Jurusan"`
}

type Regcustomer struct {
	Nama          string `json:"nama"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Gender        int    `json:"gender"`
}

type RegAccount struct {
	ID        uint   `json:"id"`
	Fullname  string `json:"fullname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	Url_photo string `json:"url_photo"`
}

type Jurusan struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
}

type Hobi struct {
	ID   int    `gorm:"primarykey" json:"id"`
	Nama string `json:"nama"`
}
