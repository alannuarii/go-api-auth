package models

type User struct {
    ID          		int    `db:"id_user"`
    Name    			string `db:"name"`
    Email      			string `db:"email"`
    Password      		string `db:"password"`
    Role      			string `db:"role"`
}