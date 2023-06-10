package repository

const (
	saveURLQuery = `insert into urls (long,short) values ($1,$2) returning short`

	selectExistURLQuery = `SELECT short FROM urls where long = $1;`

	getURLQuery = "select long from urls where short=$1"
)
