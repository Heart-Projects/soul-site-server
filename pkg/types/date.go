package types

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

const (
	timeFormat = "2006-01-02 15:04:05"
)

func (d *DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", d.Format(timeFormat))
	return []byte(formatted), nil
}

// 实现 time.Time 转换
func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	d.Time, err = time.Parse(timeFormat, string(data))
	return
}

// 实现 driver.Valuer 接口，转换为数据库支持的值, 注意这里是 ct DateTime 类型，否则转换会失败
func (ct DateTime) Value() (driver.Value, error) {
	if ct.Time.IsZero() {
		return nil, nil
	}
	v := ct.Format(timeFormat) // 标准 MySQL/SQL 格式
	return v, nil
	//ct.Time, _ = time.Parse(timeFormat, ct.String())
	//return ct.Time, nil
}

// 实现 sql.Scanner 接口，从数据库中扫描值
func (dt *DateTime) Scan(value interface{}) error {
	if value == nil {
		*dt = DateTime{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*dt = DateTime{Time: v}
		return nil
	case string:
		t, err := time.Parse(timeFormat, v)
		if err != nil {
			return err
		}
		*dt = DateTime{Time: t}
		return nil
	default:
		return fmt.Errorf("cannot convert %T to DateTime", value)
	}
}

func (dt *DateTime) String() string {
	return dt.Format(timeFormat)
}
