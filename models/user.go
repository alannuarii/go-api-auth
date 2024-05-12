package models

type User struct {
    ID          		int    `db:"id_user"`
    Name    			string `db:"name"`
    Email      			string `db:"email"`
    Role      			string `db:"role"`
}