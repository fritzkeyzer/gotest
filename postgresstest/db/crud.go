package db

import (
	"context"
	"github.com/fritzkeyzer/test/postgresstest/types"
	_ "github.com/lib/pq"
	//"github.com/davecgh/go-spew/spew"
)


func (cl *DBClient) DropTableStudents() (err error){
	defer func() {
		if e := recover(); e != nil {
			err = cl.error("DropTableStudents")
		}
	}()

	err = cl.open()
	if err != nil{
		return cl.error("DropTableStudents", err)
	}
	defer cl.close()

	stmt := `drop table if exists students;`
	_, err = cl.db.Exec(context.Background(), stmt)
	if err != nil{
		return cl.error("DropTableStudents", err)
	}

	cl.debug("DropTableStudents success")

	return nil
}

func (cl *DBClient) CreateTableStudents() (err error){
	defer func() {
		if e := recover(); e != nil {
			err = cl.error("CreateTableStudents")
		}
	}()

	err = cl.open()
	if err != nil{
		return cl.error("CreateTableStudents", err)
	}
	defer cl.close()

	stmt := `create table students (name text, age int);`
	_, err = cl.db.Exec(context.Background(), stmt)
	if err != nil{
		return cl.error("CreateTableStudents", err)
	}

	cl.debug("CreateTableStudents success")

	return nil
}

func (cl *DBClient) DeleteAllStudents() (err error){
	defer func() {
		if e := recover(); e != nil {
			err = cl.error("DeleteAllStudents")
		}
	}()

	err = cl.open()
	if err != nil{
		return cl.error("DeleteAllStudents", err)
	}
	defer cl.close()

	stmt := `delete from public.students where name like '%';`
	_, err = cl.db.Exec(context.Background(), stmt)
	if err != nil{
		return cl.error("DeleteAllStudents", err)
	}

	cl.debug("DeleteAllStudents success")

	return nil
}

func (cl *DBClient) InsertStudent(student *types.Student) (err error){
	defer func() {
		if e := recover(); e != nil {
			err = cl.error("InsertStudent")
		}
	}()

	err = cl.open()
	if err != nil{
		return cl.error("InsertStudent", err)
	}
	defer cl.close()

	stmt := `insert into "students"("name", "age") values($1, $2)`
	_, err = cl.db.Exec(context.Background(), stmt, student.Name, student.Age)
	if err != nil{
		return cl.error("InsertStudent", err)
	}

	cl.debug("InsertStudent success", *student)

	return nil
}

func (cl *DBClient) GetAllStudents(students *[]types.Student) (err error){
	defer func() {
		if e := recover(); e != nil {
			err = cl.error("GetAllStudents")
		}
	}()

	err = cl.open()
	if err != nil{
		return cl.error("GetAllStudents", err)
	}
	defer cl.close()

	stmt := `select name, age from students`
	rows, err := cl.db.Query(context.Background(), stmt)
	if err != nil{
		return cl.error("GetAllStudents", err)
	}


	for rows.Next(){
		var name string
		var age int

		err = rows.Scan(&name, &age)
		if err != nil{
			return cl.error("GetAllStudents rows.Scan", err)
		}

		student := types.Student{
			Name: name,
			Age:  age,
		}

		*students = append(*students, student)
	}

	cl.debug("GetAllStudents success", *students)

	return nil
}
