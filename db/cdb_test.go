package db

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

var db *sql.DB

func init() {
	db = GetDb()
}
