package models

import (
	"github.com/beego/beego/v2/core/logs"
	"strconv"
)

type Student struct {
	Sno   int    `json:"sno,omitempty"`
	Sname string `json:"sname,omitempty"`
	Ssex  string `json:"ssex,omitempty"`
	Sage  int    `json:"sage,omitempty"`
	Sdept string `json:"sdept,omitempty"`
}

// 查询
func (s *Student) QueryAll() ([]Student, error) {
	query := "SELECT sno, sname, ssex, sage, sdept FROM Student"
	rows, err := DB.Query(query)
	if err != nil {
		logs.Error("Failed to query student: %v", err)
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Sno, &student.Sname, &student.Ssex, &student.Sage, &student.Sdept)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

// QueryByParams 根据参数查询,模糊查询
func (s *Student) QueryByParams(Sno int, Sname string) ([]Student, error) {
	query := `SELECT sno, sname, ssex, sage, sdept FROM Student WHERE sno like ? OR sname like ?`
	rows, err := DB.Query(query, "%"+strconv.Itoa(Sno)+"%", "%"+Sname+"%")
	if err != nil {
		logs.Error("Failed to query student: %v", err)
		return nil, err
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Sno, &student.Sname, &student.Ssex, &student.Sage, &student.Sdept)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

func (s *Student) Add() error {
	query := `INSERT INTO Student (sno, sname, ssex, sage, sdept) VALUES (?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, s.Sno, s.Sname, s.Ssex, s.Sage, s.Sdept)
	if err != nil {
		logs.Error("Failed to insert student: %v", err)
		return err
	}
	return nil
}

func (s *Student) Delete() error {
	query := `DELETE FROM Student WHERE sno = ?`
	_, err := DB.Exec(query, s.Sno)
	if err != nil {
		logs.Error("Failed to delete student: %v", err)
		return err
	}
	return nil
}

func (s *Student) Update() error {
	//	Update 更新学生信息,根据学号更新
	query := `UPDATE Student SET sname = ?, ssex = ?, sage = ?, sdept = ? WHERE sno = ?`
	_, err := DB.Exec(query, s.Sname, s.Ssex, s.Sage, s.Sdept, s.Sno)
	if err != nil {
		logs.Error("Failed to update student: %v", err)
		return err
	}
	return nil
}
