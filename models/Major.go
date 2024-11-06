package models

import "github.com/beego/beego/v2/core/logs"

// Major +-----------+--------------+------+-----+---------+----------------+
// | Field     | Type         | Null | Key | Default | Extra          |
// +-----------+--------------+------+-----+---------+----------------+
// | MajorID   | int(11)      | NO   | PRI | NULL    | auto_increment |
// | MajorName | varchar(100) | NO   |     | NULL    |                |
// +-----------+--------------+------+-----+---------+----------------+
type Major struct {
	MajorID   int    `json:"major_id,omitempty"`
	MajorName string `json:"major_name,omitempty"`
}

// QueryAll 查询所有专业信息
func (m *Major) QueryAll() ([]Major, error) {
	query := `SELECT MajorID, MajorName FROM Major`
	rows, err := DB.Query(query)
	if err != nil {
		logs.Error("Failed to query major: %v", err)
		return nil, err
	}
	defer rows.Close()
	var majors []Major
	for rows.Next() {
		var major Major
		err = rows.Scan(&major.MajorID, &major.MajorName)
		if err != nil {
			return nil, err
		}
		majors = append(majors, major)
	}
	return majors, nil
}
