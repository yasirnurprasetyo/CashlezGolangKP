// This package contains utilities for CRUD (Create, Read, Update, Delete) SQL data
package crud

import (
    "fmt"
    "math"
    "strings"
    "reflect"
    "strconv"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    // "strconv"
    log "github.com/Sirupsen/logrus"
)

// Format DSN by given connection parameters to SQL database
func GetFormattedDsn(host_ip string, port int, schema string, username string, 
    password string) (string) {
    var p string
    if port > 1000 {
        p= ":"+strconv.Itoa(port)
    }
    return username+":"+password+"@tcp("+host_ip+p+")/"+schema+"?parseTime=true"
}

// DB Struct
type DB struct {
    connection *sqlx.DB
}
// Using DB, connect to database with given DSN
func (o *DB) Connect(dsn string) {
    var err error
    o.connection, err = sqlx.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
}
// Using DB, disconnect to database
func (o *DB) Disconnect() {
    o.connection.Close()
}
// Using DB, query rows and returns the result in list of interface data maps
func (o *DB) QueryToMap(query string) ([]map[string]interface{}, []string, error) {
    var err error
    rows, err := o.connection.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    columns, err := rows.Columns()
    if err != nil {
        log.Fatal(err)
    }
    //var rowCount int64
    data := []map[string]interface{}{}

    srcValues := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(srcValues))
    for k := range srcValues {
        scanArgs[k] = &srcValues[k]
    }
    for rows.Next() {
        m := make(map[string]interface{})
        err = rows.Scan(scanArgs...)
        if err != nil {
            log.Fatal(err)
        }
        for i, c := range srcValues {
            if c != nil {
                m[columns[i]] = string(c)
            } else {
                m[columns[i]] = nil
            }
        }
        data = append(data, m)
    }
    return data, columns, err
}
// Using DB, query rows and returns the result by raw sql.Rows
func (o *DB) RawQuery(query string) (*sqlx.Rows, error) {
    rows, err := o.connection.Queryx(query)
    if err != nil {
        log.Fatal(err)
    }
    return rows, err
}
// Using DB, query a single row and returns the result by raw sql.Row
func (o *DB) RawQueryOne(query string) (*sqlx.Row) {
    row := o.connection.QueryRowx(query)
    return row
}
// Using DB, execute direct statement. But when it fails, the process moves 
// forward for the next one
func (o *DB) TryExecute(statement string) (result sql.Result, err error) {
    defer func() {
        if r := recover(); r != nil {
            log.Fatal(r)
            err = r.(error)
        }
    }()
    result, err = Execute(statement, o.connection)
    return result, err
}
// Using DB, execute prepated statement. But when it fails, the process moves forward 
// for next the next one
func (o *DB) TryExecutePrepared(statement string, data []interface{}) (
    result sql.Result, err error) {
    defer func() {
        if err := recover(); err != nil {
            log.Fatal(err)
        }
    }()
    result, err = ExecutePrepared(statement, data, o.connection)
    return result, err
}

// Execute the SQL statement
func Execute(statement string, db *sqlx.DB) (sql.Result, error) {
    result, err := db.Exec(statement)
    if err != nil {
        log.Fatal(err)
    }
    return result, err
}

// Execute prepared SQL statement
func ExecutePrepared(statement string, data []interface{}, db *sqlx.DB) (
    sql.Result, error) {
    prepared, err := db.Prepare(statement)
    if err != nil {
        log.Fatal(err)
        log.Debug("Statement SQL: "+statement)
    }
    result, err := prepared.Exec(data...)
    if err != nil {
        log.Fatal("Cannot execute statement", err)
    }
    return result, err
}

func TruncateTable(tableName string, dsn string) (bool) {
    // Create connection to SQL for inserting data
    db := &DB{}
    db.Connect(dsn)
    defer db.Disconnect()

    _, err := db.TryExecute("TRUNCATE TABLE "+tableName)
    if err != nil {
        log.Error(err.Error())
    } else {
        return true
    }
    return false
}

func InsertRows(data [][]interface{}, columns []string, tableName string, 
    dsn string, chunkSize int) (bool) {
    var passed bool

    // Create connection to SQL for inserting data
    db := &DB{}
    db.Connect(dsn)
    defer db.Disconnect()

    // Create basic form of insert placeholder
    p := []string{}
    for h:=0; h<len(columns); h++ {
        p = append(p, "?")
    }

    var loops int
    if chunkSize > 0 {
        // chunkSize := 100
        l := float64(len(data)) / float64(chunkSize)
        loops = int(math.Ceil(l))
    } else {
        chunkSize = len(data)
        loops = 1
    }

    log.Info("There are "+strconv.Itoa(len(data))+
        " rows to copy, splitted into "+
        strconv.Itoa(loops)+" chunk(s), each with size of "+
        strconv.Itoa(chunkSize)+" rows")
    
    for x := 0; x < loops; x++ {
        chunk := []interface{}{}

        y0 := x*chunkSize
        y9 := x*chunkSize+chunkSize
        if y9 > len(data) {
            y9 = len(data)
        }

        // Multiple the placeholder by size of a chunk
        placeholder := []string{}
        for y := y0; y < y9; y++ {
            placeholder = append(placeholder, "("+strings.Join(p, ",")+")")
            if len(columns) == len(data[y]) {
                for _, c := range data[y] {
                    chunk = append(chunk, c)
                }
            } else {
                panic("Numbers of data column does not match the header size")
            }
        }

        // Form the prepared insert statement with prepared placeholder
        insert := "INSERT INTO "+tableName+
            " ("+strings.Join(columns, ",")+")"+
            " VALUES "+strings.Join(placeholder, ",")
        log.Println("Statement: "+insert)
        _, err := db.TryExecutePrepared(insert, chunk) 
        if err != nil {
            log.Error(err.Error())
        } else {
            passed = true
        }
    }
    return passed
}

func StructFields(values interface{}) ([]string) {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fields := []string{}
	if v.Kind() == reflect.Struct {
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field != "" {
				fields = append(fields, field)
			}
		}
		return fields
	}
	if v.Kind() == reflect.Map {
		for _, keyv := range v.MapKeys() {
			fields = append(fields, keyv.String())
		}
		return fields
	}
    panic(fmt.Errorf("DbFields requires a struct or a map, found: %s", 
        v.Kind().String()))
}

func StructValues(values interface{}) ([]interface{}) {
    v := reflect.ValueOf(values).Elem()
    data := []interface{}{}
    
    for i := 0; i < v.NumField(); i++ {
        d := v.Field(i).Interface()
        data = append(data, d)

        // fmt.Println("%v", d)
        // switch value.Kind() {
        // case reflect.String:
        //     v := value.String()
        //     fmt.Print(v, "\n")
        // case reflect.Int:
        //     v := strconv.FormatInt(value.Int(), 10)
        //     fmt.Print(v, "\n")
        // case reflect.Int32:
        //     v := strconv.FormatInt(value.Int(), 10)
        //     fmt.Print(v, "\n")
        // case reflect.Int64:
        //     v := strconv.FormatInt(value.Int(), 10)
        //     fmt.Print(v, "\n")
        // default:
        //     assert.Fail(t, "Not support type of struct")
        // }
    }
    return data
}
