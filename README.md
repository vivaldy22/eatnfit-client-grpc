# Endpoint List!
#### All /admin endpoints must use token from /auth/login
 - [POST] /auth/login
 > {<br>
 > "user_email": "{your email}",<br>
 > "user_password": "{your password}" <br>
 > }
 - [POST] /auth/register
 > {<br>
 > "user_email":  "email@gmail.com",<br>
 > "user_password":  "password",<br>
 > "user_f_name":  "f name",<br>
 > "user_l_name":  "l name",<br>
 > "user_gender":  "1",<br>
 > "user_photo":  "admin.jpg",<br>
 > "user_balance":  "10",<br>
 > "user_level":  "1"<br>
 > }
 - [GET] /admin/levels
 - [POST] /admin/level
 > {<br>
 > "level_name": "{new level}"<br>
 > }
 - [GET, PUT, DELETE] /admin/level/{id} 
 > {<br>
 > "level_name": "{new level}"<br>
 > }
 - [GET] /admin/genders
 - [POST] /admin/gender
 > {<br>
 > "gender_name": "{new gender}"<br>
 > }
 - [GET, PUT, DELETE] /admin/gender/{id}
 > {<br>
 > "gender_name": "{new gender}"<br>
 > }
 - [GET] /admin/users
 - [POST] /admin/user
 > {<br>
 > "user_email":  "email@gmail.com",<br>
 > "user_password":  "password",<br>
 > "user_f_name":  "f name",<br>
 > "user_l_name":  "l name",<br>
 > "user_gender":  "1",<br>
 > "user_photo":  "admin.jpg",<br>
 > "user_balance":  "10",<br>
 > "user_level":  "1"<br>
 > }
 - [GET, PUT, DELETE] /admin/user/{id}
 > {<br>
 > "user_email":  "email@gmail.com",<br>
 > "user_password":  "password",<br>
 > "user_f_name":  "f name",<br>
 > "user_l_name":  "l name",<br>
 > "user_gender":  "1",<br>
 > "user_photo":  "admin.jpg",<br>
 > "user_balance":  "10",<br>
 > "user_level":  "1"<br>
 > }
 - [GET] /users
 - [GET] /user/{id}
