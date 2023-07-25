package config

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Config struct {
	// API.
	Host        string `split_words:"true" default:"localhost"`
	Port        int    `split_words:"true" default:"3000"`
	Environment string `split_words:"true" required:"true"`
	Debug       bool   `split_words:"true" default:"false"`

	SecretKey string `split_words:"true" default:"false"`

	//DB
	DB DBConfig `split_words:"true" required:"true"`
}

func Read() (*Config, error) {
	cfg := Config{}
	if err := envconfig.Process("MEETING_API", &cfg); err != nil {
		log.Fatal(err.Error())
		return nil, err

	}

	return &cfg, nil
}

type DBConfig struct {
	Addr         string        `required:"true"`
	User         string        `required:"true"`
	Password     string        `required:"true"`
	Name         string        `required:"true"`
	AppName      string        `required:"true"`
	TLSMode      string        `default:"disable" envconfig:"tls_mode"`
	DialTimeout  time.Duration `envconfig:"dial_timeout"`
	ReadTimeout  time.Duration `envconfig:"read_timeout"`
	WriteTimeout time.Duration `envconfig:"write_timeout"`
}

func (c DBConfig) Connect() *sql.DB {

	if c.AppName == "" {
		panic(fmt.Errorf("AppName must not be empty"))
	}

	opts := []pgdriver.Option{
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(c.Addr),
		pgdriver.WithUser(c.User),
		pgdriver.WithPassword(c.Password),
		pgdriver.WithDatabase(c.Name),
		pgdriver.WithApplicationName(c.AppName),
	}

	if c.TLSMode == "disable" {
		opts = append(opts, pgdriver.WithInsecure(true))
	}

	if c.DialTimeout > 0 {
		opts = append(opts, pgdriver.WithDialTimeout(c.DialTimeout))
	}
	if c.ReadTimeout > 0 {
		opts = append(opts, pgdriver.WithReadTimeout(c.ReadTimeout))
	}
	if c.WriteTimeout > 0 {
		opts = append(opts, pgdriver.WithWriteTimeout(c.WriteTimeout))
	}

	pgconn := pgdriver.NewConnector(opts...)
	return sql.OpenDB(pgconn)
}

func (c DBConfig) NewDB(opts ...bun.DBOption) *bun.DB {
	return bun.NewDB(c.Connect(), pgdialect.New(), opts...)
}

func (c DBConfig) MustNewDB(opts ...bun.DBOption) *bun.DB {
	db := c.NewDB(opts...)
	result := 0
	if err := db.NewSelect().ColumnExpr("1+1").Scan(context.Background(), &result); err != nil {
		panic(errors.New("test query 1+1"))
	}
	if result != 2 {
		panic(errors.New("wrong test query 1+1 result"))
	}
	return db
}
