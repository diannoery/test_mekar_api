package utils

const (
	ALL_USER        = "select u.id_users,u.nik,u.nama,u.tgl_lahir,u.pekerjaan,pk.label_pekerjaan,u.pendidikan_terakhir,p.label_pendidikan,u.user_status,u.created_date,u.update_date  from tb_users u inner join m_pekerjaan pk on u.pekerjaan = pk.id_pekerjaan inner join m_pendidikan p on u.pendidikan_terakhir=p.id_pendidikan where u.user_status = 1;"
	ALL_USER_PAGE   = "select u.id_users,u.nik,u.nama,u.tgl_lahir,u.pekerjaan,pk.label_pekerjaan,u.pendidikan_terakhir,p.label_pendidikan,u.user_status,u.created_date,u.update_date  from tb_users u inner join m_pekerjaan pk on u.pekerjaan = pk.id_pekerjaan inner join m_pendidikan p on u.pendidikan_terakhir=p.id_pendidikan where u.user_status = 1 LIMIT ?,?;"
	GET_USER_BY_ID  = "select u.id_users,u.nik,u.nama,u.tgl_lahir,u.pekerjaan,pk.label_pekerjaan,u.pendidikan_terakhir,p.label_pendidikan,u.user_status,u.created_date,u.update_date  from tb_users u inner join m_pekerjaan pk on u.pekerjaan = pk.id_pekerjaan inner join m_pendidikan p on u.pendidikan_terakhir=p.id_pendidikan where u.id_users = ?;"
	CREATE_USERS    = "INSERT INTO tb_users (id_users, nik, nama, tgl_lahir, pekerjaan, pendidikan_terakhir, user_status, created_date, update_date) VALUES (?,?,?,?,?,?,?,?,?);"
	UPDATE_USER     = "UPDATE tb_users SET nik= ?, nama = ?, tgl_lahir = ?, pekerjaan = ?, pendidikan_terakhir = ?, user_status = ?, update_date = ? WHERE id_users = ?;"
	PEKERJAN        = "select * from m_pekerjaan;"
	PENDIDIKAN      = "select * from m_pendidikan;"
	DELETE_USER     = "delete from tb_users where id_users = ?;"
	REGISTER_ADMIN  = "insert into user_login(username,password) values (?,?);"
	LOGIN_ADMIN     = "select * from user_login where username =?;"
	COUNT_DATA_USER = `SELECT count(*) as total_data FROM tb_users`
)
