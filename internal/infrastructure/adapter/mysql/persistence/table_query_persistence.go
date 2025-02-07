package persistence

import (
	"database/sql"
	"easydbTools/internal/common/easytool/common/page"
	"easydbTools/internal/common/easytool/convert"
	"easydbTools/internal/common/easytool/util/str_util"
	"easydbTools/internal/domain/mysql/model"
	"easydbTools/internal/domain/mysql/repository"
	"easydbTools/internal/domain/mysql/repository/params"
	"fmt"
	"strings"
)

type TableQueryRepositoryImpl struct {
	datasourceRepository repository.DataSourceRepository
}

func NewTableQueryRepositoryImpl() *TableQueryRepositoryImpl {
	return &TableQueryRepositoryImpl{
		datasourceRepository: NewDataSourceRepositoryImpl(),
	}
}

// ListTableNames is used to get all table names from the database
func (t *TableQueryRepositoryImpl) ListTableNames(database model.Database) ([]string, error) {
	// get the connection
	connect := t.datasourceRepository.ConnectById(database.DataSourceId)
	// build the sql
	executeSql := fmt.Sprintf("show tables from %s;", database.Name)
	rows, err := connect.Query(executeSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// scan the result
	var tableNames []string
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}
	return tableNames, nil
}

// GetTable is used to get the table information
func (t *TableQueryRepositoryImpl) GetTable(params *params.TableQueryParams) (*model.Table, error) {
	columns := t.parseColumns(params)
	indexes := t.parseIndexes(params)
	comment := t.parseTableComment(params)
	createSQL := t.parseCreateTableSQL(params)
	table := &model.Table{
		Name:      params.TableName,
		Desc:      comment,
		CreateSQL: createSQL,
		Database:  &model.Database{Name: params.DatabaseName},
		Columns:   columns,
		Indexes:   indexes,
	}
	for _, index := range *table.Indexes {
		for k := range *index.Columns {
			cols := *index.Columns
			col, _ := table.GetColumnByName(cols[k].Name)
			cols[k] = *col
		}
	}
	return table, nil
}

// parseColumns is query describe table sql  to parse columns
func (t *TableQueryRepositoryImpl) parseColumns(params *params.TableQueryParams) *[]model.Column {
	connect := t.datasourceRepository.ConnectById(params.DataSourceId)
	showColSQL := fmt.Sprintf("SHOW FULL COLUMNS FROM %s.%s", params.DatabaseName, params.TableName)
	cols, err := connect.Query(showColSQL)
	if err != nil {
		fmt.Println(err)
	}
	defer cols.Close()
	var columns []model.Column
	// scans the columns
	for cols.Next() {
		var field, colType, collation, null, key, colDefault, extra, privileges, comment sql.NullString
		if err := cols.Scan(&field, &colType, &collation, &null, &key, &colDefault, &extra, &privileges, &comment); err != nil {
			panic(err)
		}
		colTypeName := strings.Split(colType.String, "(")[0]
		size := 0
		pointStr := ""
		if colTypeName == "decimal" {
			pointStr = str_util.ExtractStringFromBorder(colType.String, "(", ")")
		} else {
			size = convert.ToInt(str_util.ExtractStringFromBorder(colType.String, "(", ")"))
		}
		column := model.Column{
			Name:    field.String,
			Type:    colTypeName,
			Size:    size,
			Points:  pointStr,
			Default: colDefault.String,
			NotNull: null.String == "YES",
			Comment: comment.String,
			Pk:      key.String == "PRI",
		}
		columns = append(columns, column)
	}
	return &columns
}

func (t *TableQueryRepositoryImpl) parseIndexes(params *params.TableQueryParams) *[]model.Index {
	connect := t.datasourceRepository.ConnectById(params.DataSourceId)
	showIndexSQL := fmt.Sprintf("SHOW INDEX FROM %s.%s", params.DatabaseName, params.TableName)

	idxes, err := connect.Query(showIndexSQL)
	if err != nil {
		fmt.Println(idxes)
	}
	defer idxes.Close()
	var indexes []model.Index
	for idxes.Next() {
		var (
			table, nonUnique, indexName, seqInIndex, columnName, collation, cardinality, subPart, packed, null, indexType, comment, indexComment, visibility, expression sql.NullString
		)
		if err := idxes.Scan(&table, &nonUnique, &indexName, &seqInIndex, &columnName, &collation, &cardinality, &subPart, &packed, &null, &indexType, &comment, &indexComment, &visibility, &expression); err != nil {
			panic(err)
		}
		// if the index exist, add the column to the index
		isExist := false
		for _, index := range indexes {
			if indexName.String == index.Name {
				*index.Columns = append(*index.Columns, model.Column{Name: columnName.String})
				isExist = true
				break
			}
		}
		if !isExist && indexName.String != "PRIMARY" {
			index := model.Index{
				Name:   indexName.String,
				Unique: nonUnique.String == "0",
				Columns: &[]model.Column{
					{
						Name: columnName.String,
					},
				},
				Comment: indexComment.String,
			}
			indexes = append(indexes, index)
		}
	}
	return &indexes
}

// parseTableComment is query describe table sql  to parse table comment
func (t *TableQueryRepositoryImpl) parseTableComment(params *params.TableQueryParams) string {
	connect := t.datasourceRepository.ConnectById(params.DataSourceId)
	sqlStr := fmt.Sprintf("SELECT TABLE_COMMENT FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '%s' AND TABLE_NAME = '%s';", params.DatabaseName, params.TableName)
	var tableComment sql.NullString
	rows, _ := connect.Query(sqlStr)
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	if rows.Next() {
		_ = rows.Scan(&tableComment)
	}
	return tableComment.String
}

// parseCreateTableSQL query the table create sql
func (t *TableQueryRepositoryImpl) parseCreateTableSQL(params *params.TableQueryParams) string {
	connect := t.datasourceRepository.ConnectById(params.DataSourceId)
	sqlStr := fmt.Sprintf("SHOW CREATE TABLE %s.%s;", params.DatabaseName, params.TableName)
	var table, createSQL sql.NullString
	rows, _ := connect.Query(sqlStr)
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	if rows.Next() {
		_ = rows.Scan(&table, &createSQL)
	}
	return createSQL.String
}

// PageTableData is used to get table data from the database
func (t *TableQueryRepositoryImpl) PageTableData(param *params.TablePageParams) *page.Page[[]map[string]interface{}] {
	connect := t.datasourceRepository.ConnectById(param.TableQueryParams.DataSourceId)
	totalSQL := fmt.Sprintf("SELECT COUNT(0) FROM %s.%s;", param.TableQueryParams.DatabaseName, param.TableQueryParams.TableName)
	query, err := connect.Query(totalSQL)
	total := 0
	if query.Next() {
		query.Scan(&total)
	}
	fmt.Println(query, err)
	return nil
}
