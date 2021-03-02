**pengaturan**
- buka file config.json atur sesuai kebutuhan
ex

**langkah**
1. Import .sql ->
2. jalankan main.go

**API ROUTES**
- **User**
1. /user [POST] -> tambah data
2. /user [GET] -> mengambil semua data
3. /user/{id} [GET] mengambil data sesuai id
4. /user [PUT] -> mengedit data sesuai id
5. /user/{id} [DELETE] -> delete data

Output Program
    UNTUK USER
        1.CREATE USER
        URL => /users
        Method : POST
            Body :
                Contoh Body :{
                                "nik": "12123123",
                                "nama": "dian",
                                "tanggal_lahir": "1998-11-01",
                                "pekerjaan": {
                                    "pekerjaan_id": "8219c7cc-30bb-11eb-b405-c85b766bafe8"
                                },
                                "pendidikan_terakhir": {
                                    "id_pendidikan": "5df1331d-30bb-11eb-b405-c85b766bafe8"
                                }
                            }

        2.Update user
            URL -> /users
            Menthod : PUT
            Body : {           
                    "user_id":"ee01f821-1e9e-47db-933c-0122e20e1a90",
                    "nik": "12123123",
                    "nama": "bbb",
                    "tanggal_lahir": "1998-11-01",
                    "pekerjaan": {
                        "pekerjaan_id": "8219c7cc-30bb-11eb-b405-c85b766bafe8"
                    },
                    "pendidikan_terakhir": {
                        "id_pendidikan": "5df1331d-30bb-11eb-b405-c85b766bafe8"
                    }
                }
            
        3.Get User By Id
         URL ->  /users/36d0035f-cd53-4b55-a50c-0ce77d7b8ee4
            Menthod : Get
           
        4.Delete user
            URL ->  /users/36d0035f-cd53-4b55-a50c-0ce77d7b8ee4
            Menthod : DELETE


        5.login admin
            Method : POST
            Body :
                Contoh Body :{
                                "username": "dian",
                                "password": "123",
                            }

         5.register admin
            Method : POST
            Body :
                Contoh Body :{
                                "username": "dian",
                                "password": "123",
                            }