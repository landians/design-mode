package template_method

import (
	"database/sql"
	"errors"
)

// DAO 对象查询接口
type IDao interface {
	QueryOne(db *sql.DB, sql string, args ...interface{}) error
	QueryMulti(db *sql.DB, sql string, args ...interface{}) error
}

type FNBeforeQuery func() error
type FNScanRow func(rows *sql.Rows) error

// 基本实现 IDao 接口的对象
type baseDAO struct {
	fnBeforeQuery FNBeforeQuery
	fnScanRow     FNScanRow
}

func newBaseDAO(fq FNBeforeQuery, fs FNScanRow) *baseDAO {
	return &baseDAO{
		fnBeforeQuery: fq,
		fnScanRow:     fs,
	}
}

func (b *baseDAO) QueryOne(db *sql.DB, sql string, args ...interface{}) error {
	if b.fnScanRow == nil {
		return errors.New("baseDAO.fnScanRow is nil")
	}

	if b.fnBeforeQuery != nil {
		e := b.fnBeforeQuery()
		if e != nil {
			return e
		}
	}

	rows, e := db.Query(sql, args...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if e != nil {
		return e
	}

	if rows.Next() {
		return b.fnScanRow(rows)
	} else {
		return errors.New("no rows found")
	}
}

func (b *baseDAO) QueryMulti(db *sql.DB, sql string, args ...interface{}) error {
	if b.fnScanRow == nil {
		return errors.New("tBaseDAO.fnScanRow is nil")
	}

	if b.fnBeforeQuery != nil {
		e := b.fnBeforeQuery()
		if e != nil {
			return e
		}
	}

	rows, e := db.Query(sql, args...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()

	if e != nil {
		return e
	}

	for rows.Next() {
		e = b.fnScanRow(rows)
		if e != nil {
			return e
		}
	}

	return nil
}

// 用户实体信息
type UserInfo struct {
	ID     int
	Name   string
	Pwd    string
	OrgID  int
	RoleID string
}

func newUserInfo() *UserInfo {
	return &UserInfo{}
}

// 用户信息查询接口
type IUserDAO interface {
	GetUserByID(db *sql.DB, id int) (error, *UserInfo)
	GetUsersByOrgID(db *sql.DB, orgID int) (error, []*UserInfo)
}

// 继承 baseDAO，并注入数据行映射函数, 同时实现IUserDAO接口, 提供用户信息查询功能
type UserDAO struct {
	baseDAO
	items []*UserInfo
}

func newUserDAO() IUserDAO {
	it := &UserDAO{
		baseDAO: *newBaseDAO(nil, nil),
		items:   nil,
	}
	it.fnBeforeQuery = it.beforeQuery
	it.fnScanRow = it.ScanRow
	return it
}

func (u *UserDAO) beforeQuery() error {
	u.items = make([]*UserInfo, 0)
	return nil
}

func (u *UserDAO) ScanRow(rows *sql.Rows) error {
	user := newUserInfo()
	if err := rows.Scan(&user.ID, &user.Name, &user.Pwd, &user.OrgID, &user.RoleID); err != nil {
		return err
	}
	u.items = append(u.items, user)
	return nil
}

func (u *UserDAO) GetUserByID(db *sql.DB, id int) (error, *UserInfo) {
	if err := u.QueryOne(db, "select id,name,pwd,org_id,role_id from user_info where id=?", id); err != nil {
		return err, nil
	}
	return nil, u.items[0]
}

func (u *UserDAO) GetUsersByOrgID(db *sql.DB, orgID int) (error, []*UserInfo) {
	if err := u.QueryMulti(db, "select id,name,pwd,org_id,role_id from user_info where org_id=?", orgID); err != nil {
		return err, nil
	}
	return nil, u.items
}
