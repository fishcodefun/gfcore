package server

import (
	"fmt"
	"reflect"

	"github.com/gogf/gf/frame/g"
)

type TableFilter struct {
	Field string
	Label string
	Width int
}

// Column 结构体
type Column struct {
	Field    string // 字段名（对应 User 的字段）
	Label    string // 显示标签（表头名称）
	Sortable bool   // 是否可排序
}

// ziduan
func Field(field string, label string) Column {
	return Column{
		Field: field,
		Label: label,
	}
}

func (c Column) Sort() Column {
	c.Sortable = true
	return c
}

func (c Column) Using([]string) Column {
	return c
}

func Response(users []any, columns []Column) g.Map {
	var datas [][]string
	for _, row := range users {
		g.Dump(row)
		v := reflect.ValueOf(row)
		var content []string
		for _, field := range columns {
			fieldValue := v.FieldByName(field.Field)
			content = append(content, fmt.Sprintf("%v", fieldValue))
		}
		g.Dump(content)
		datas = append(datas, content)
	}
	response := g.Map{
		"Datas":     datas,
		"Columns":   columns,
		"TableName": "表名",
	}
	return response
}
