<p align="right">
بِسْــــــــــــــمِ اللَّهِ الرَّحْمَنِ الرَّحِيم 
</p>

# Golang JWT Restful API MySql

Auth login register sederhana berbasis REST API. 
Password hash menggunakan argon2id.

## Installation
- Silahkan download/clone repository ini
- Buat database terlebih dahulu
- ubah file [.env](.env) sesuai konfigurasi kalian
- Jalankan main.go dengan perintah ```go run main.go ```
- terdapat seed saat menjalankan server, jika tidak ingin mengulanginya lagi, silahkan comment di file [main.go](main.go#L17)

## Input dan Response

Contoh inputan tiap url endpoint APInya
* url ```/api/users``` Method POST untuk registrasi akun baru
    - Input
        ```
        {
            "username":"afrizalsblog",
            "email": "dsi@gmail.com",
            "password": "afrizal"
        }
        ```
    - Response
        ```
        {
            "id": 6,
            "username": "afrizalsblog",
            "email": "dsi@gmail.com",
            "password": "$argon2id$v=19$m=65536,t=1,p=2$YCrzu5xuBoOSVWLV4zgTgQ$e+vOvUta80ye440ZOJKsjUvLDN6LCYKHRYZZQ0aMOWw",
            "created_at": "2020-11-21T14:18:43.232855253+07:00",
            "updated_at": "2020-11-21T14:18:43.232855492+07:00"
        }
        ```
* url ```/auth/login``` Method POST untuk login dan mendapatkan token JWT
    - Input
        ```
        {
            "email": "dsi@gmail.com",
            "password": "afrizal"
        }
        ```
    - Response
        ```
        {
            "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU5NDU1NDgsInVzZXJfaWQiOjN9.-vXH-pIilWBhGrpIiqFFzdr1sA5yKyYdIYAWDFWlixM",
            "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDYwMjgzNDgsInVzZXJfaWQiOjN9.9bNDrjjxtxpXM7JVmasBXauAg0lIja-wsE7MhHuHw4I"
        }
        ```
* url ```/api/users``` Method GET untuk melihat semua data users
    - Response
        ```
        [
            {
                "id": 1,
                "username": "afrizal",
                "email": "asd@gmail.com",
                "password": "$argon2id$v=19$m=65536,t=1,p=2$IVpVaa+Vd960IAQgrx/bug$FCcx1wjPQ+t8fRZG6Ze+adMl3CrX4X8QK3iuBHFf2ao",
                "created_at": "2020-11-21T13:53:55+07:00",
                "updated_at": "2020-11-21T13:53:55+07:00"
            }
        ]
        ```
* url ```/api/users/{id}``` Method GET untuk melihat detail data salah satu users
    - Response
        ```
        {
            "id": 1,
            "username": "afrizal",
            "email": "asd@gmail.com",
            "password": "$argon2id$v=19$m=65536,t=1,p=2$IVpVaa+Vd960IAQgrx/bug$FCcx1wjPQ+t8fRZG6Ze+adMl3CrX4X8QK3iuBHFf2ao",
            "created_at": "2020-11-21T13:53:55+07:00",
            "updated_at": "2020-11-21T13:53:55+07:00"
        }
        ```
* url ```/api/users/{id}``` Method PUT untuk mengupdate data, khusus data yang mempunyai id yang sama ketika login.
<b><i>Jangan lupa untuk header bearer isikan value access_token yang didapat waktu login</i></b>
    - Response
        ```
        {
            "id": 1,
            "username": "rizal",
            "email": "asd@gmail.com",
            "password": "$argon2id$v=19$m=65536,t=1,p=2$IVpVaa+Vd960IAQgrx/bug$FCcx1wjPQ+t8fRZG6Ze+adMl3CrX4X8QK3iuBHFf2ao",
            "created_at": "2020-11-21T13:53:55+07:00",
            "updated_at": "2020-11-21T13:53:55+07:00"
        }
        ```
* url ```/api/users/{id}``` Method DELETE untuk menghapus data, khusus data yang mempunyai id yang sama ketika login.
<b><i>Jangan lupa untuk header bearer isikan value access_token yang didapat waktu login</i></b>
    - Response
        ```
        {
            "data sukses terhapus"
        }
        ```
* url ```/api/refresh``` Method POST untuk refresh token JWT untuk mendapatkan token JWT yang baru, inputan berupa <b>refresh_token</b>.
<b><i>Jangan lupa untuk header bearer isikan value access_token yang didapat waktu login</i></b>
    - Input
        ```
        {
            "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDYwMjgzNDgsInVzZXJfaWQiOjN9.9bNDrjjxtxpXM7JVmasBXauAg0lIja-wsE7MhHuHw4I"
        }
        ```
    - Response
        ```
        {
            "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDU5NDYyMDEsInVzZXJfaWQiOjN9.FvYH3v1sZtVzKaf9rNoSPSYrgOs911-rLb-tCIKvBLw",
            "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MDYwMjkwMDEsInVzZXJfaWQiOjN9.dS8b0nGBuz1hlBCU-yi7yUj271LTNJvrkkMdBLs0z7E"
        }
        ```

## Donation

* Mungkin ada yang mau berdonasi atas ide pembuatan repo ini, siapapun, berapapun, saya ucapkan terimakasih sebanyak-banyaknya. Via Gopay / Dana.

### Gopay<br>
<img src="img/gpy.png" height="400"> <br>

### Dana<br>
<img src="img/dana.png" height="350">
