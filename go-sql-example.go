package foo

func query(s ...string) {
	const q1 = "select * from user where nickname like '%`%'"
	const q2 = `select * from user where nickname like '%"%'`
	const q3 = `select
	*
from
	user
`
	const q4 = /*sql*/ `
	select
		*
	from
		user
	where
		nickname like '%"%'
`
	const q5 = "update user set nickname = 'okok' where id = 1"
	const q6 = "insert into user (null, '\"haha\"')"

	query("select * from user")
	query(`select * from user where nickname like '%"%'`)
	query(`select
	*
from
	user
`)
	query( /*sql*/ `
select
	*
from
	user
where
	nickname like '%"%'
`)
	query("update user set nickname = 'okok' where id = 1")
	query("insert into user (null, '\"haha\"')")
}
