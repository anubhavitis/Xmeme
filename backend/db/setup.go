package db

import (
	"database/sql"
	"fmt"

	"github.com/anubhavitis/Xmeme/backend/util"
	_ "github.com/go-sql-driver/mysql"
)

//Mydb function
var Mydb *sql.DB

//InitDb function
func InitDb() (*sql.DB, error) {

	dab, err := sql.Open("mysql", "zeddie:anubhav@db@(db4free.net:3306)/xmeme07")
	if err != nil {
		fmt.Println("Error at opening database")
		return nil, err
	}
	if err := dab.Ping(); err != nil {
		fmt.Println("Error at ping.")
		return nil, err
	}

	if e := CreateMemeTable(dab); e != nil {
		fmt.Println("Error at creating Meme:", e)
	}

	return dab, nil
}

//RemRec function
func RemRec(id int) error {

	q := `DELETE FROM memes WHERE id=?;`
	if _, err := Mydb.Exec(q, id); err != nil {
		return err
	}

	return nil
}

//CreateMemeTable function
func CreateMemeTable(db *sql.DB) error {
	queryUsers := `
	CREATE TABLE IF NOT EXISTS
	memes (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		caption varchar(255),
		url varchar(500),
		time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(queryUsers); err != nil {
		return err
	}

	return nil
}

//AddMeme function
func AddMeme(obj util.Meme) (int, error) {
	q := `
	INSERT INTO memes
	(name, caption, url)
	Values (?,?,?)
	`
	if _, e := Mydb.Exec(q, obj.Name, obj.Caption, obj.URL); e != nil {
		return -1, e
	}

	q = `SELECT LAST_INSERT_ID();`
	res, e := Mydb.Query(q)

	if e != nil {
		return -1, e
	}
	defer res.Close()
	var id int
	for res.Next() {
		if err := res.Scan(&id); err != nil {
			return -1, err
		}
	}
	return id, nil
}

//ShowAllMeme fucntion
func ShowAllMeme() ([]util.Meme, error) {
	var res []util.Meme

	q := `Select * from memes`
	rows, e := Mydb.Query(q)

	if e != nil {
		return res, e
	}
	defer rows.Close()

	for rows.Next() {
		var temp util.Meme
		if err := rows.Scan(&temp.ID, &temp.Name, &temp.Caption, &temp.URL, &temp.Time); err != nil {
			return res, err
		}
		res = append(res, temp)
	}

	return res, nil

}

//EditRec function
func EditRec(id int, url string, cap string) error {
	qURL := `
	UPDATE memes
	SET url=?
	where id=?
	`
	qCaption := `
	UPDATE memes
	SET caption=?
	where id=?
	`

	if url != "" {
		if _, e := Mydb.Exec(qURL, url, id); e != nil {
			return e
		}
	}

	if cap != "" {
		if _, e := Mydb.Exec(qCaption, cap, id); e != nil {
			return e
		}
	}

	return nil
}
