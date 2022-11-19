package dbhelper

import (
	"fmt"
	"github.com/go-eagle/eagle/global"
	"github.com/go-eagle/eagle/infrastructure/logger"
	"strings"
)

// NullType 空字节类型
type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// WhereBuild sql build where
// see: https://github.com/jinzhu/gorm/issues/2055
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}

		fmt.Println(strings.Join(ks, ","))
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				fmt.Println()
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, v)
			}
		}
	}
	return
}

func GetAllData(ctx global.SysContext, md interface{}, result interface{}) error {
	d := ctx.ReadDB().Model(md).Find(result)
	if d.Error != nil {
		logger.Errorf("[DB] GetAllData error:%s", d.Error.Error())
		return d.Error
	}
	return nil
}

func SearchAllData(ctx global.SysContext, md interface{}, result interface{}, cond map[string]interface{}) error {
	condSql, valList, err := WhereBuild(cond)
	if err != nil {
		return err
	}
	d := ctx.ReadDB().Model(md).Where(condSql, valList...).Find(result)
	if d.Error != nil {
		logger.Errorf("[DB] SearchAllData error:%s", d.Error.Error())
		return d.Error
	}
	return nil
}

func UpdateDataWithCond(ctx global.SysContext, md interface{}, cond map[string]interface{}, values interface{}) error {
	condSql, valList, err := WhereBuild(cond)
	if err != nil {
		return err
	}
	d := ctx.WriteDB().Model(md).Where(condSql, valList...).Updates(values)
	if d.Error != nil {
		logger.Errorf("[DB] UpdateDataWithCond error:%s", d.Error.Error())
		return d.Error
	}
	return nil
}

func CreateData(ctx global.SysContext, data interface{}) error {
	d := ctx.WriteDB().Model(data).Create(data)
	if d.Error != nil {
		logger.Errorf("[DB] CreateData error:%s", d.Error.Error())
		return d.Error
	}
	return nil
}

func CreateDataInBatches(ctx global.SysContext, data interface{}, batchSize int) error {
	d := ctx.WriteDB().Model(data).CreateInBatches(data, batchSize)
	if d.Error != nil {
		logger.Errorf("[DB] CreateDataInBatches error:%s", d.Error.Error())
		return d.Error
	}
	return nil
}
