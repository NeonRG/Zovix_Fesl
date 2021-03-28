package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

var (
	General  cfg
	Database MySQL
	Cert     Fixtures
)

type cfg struct {

	// Telemetry
	TelemetryIP    string `envconfig:"TELEMETRY_IP" default:"26.35.94.218"`
	TelemetryPort  int    `envconfig:"TELEMETRY_PORT" default:"13505"`
	TelemetryToken string `envconfig:"TELEMETRY_TOKEN"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"DEBUG"`

	HTTPBind  string `envconfig:"HTTP_BIND" default:"0.0.0.0:8080"`
	HTTPSBind string `envconfig:"HTTPS_BIND" default:"0.0.0.0:443"`


	GameSpyIP string `envconfig:"GAMESPY_IP" default:"0.0.0.0"`

	FeslClientPort int `envconfig:"FESL_CLIENT_PORT" default:"18270"`
	FeslServerPort int `envconfig:"FESL_SERVER_PORT" default:"18051"`

	ThtrClientPort int    `envconfig:"THEATER_CLIENT_PORT" default:"18275"`
	ThtrServerPort int    `envconfig:"THEATER_SERVER_PORT" default:"18056"`
	ThtrAddr       string `envconfig:"THEATER_ADDR" default:"26.35.94.218"`

	MessengerAddr string `envconfig:"MESSENGER_ADDR" default:"26.35.94.218"`
	LevelDBPath   string `envconfig:"LEVEL_DB_PATH" default:"_data/lvl.db"`
}

type MySQL struct {
	UserName string `envconfig:"DATABASE_USERNAME" default:"root"`
	Password string `envconfig:"DATABASE_PASSWORD"`
	Host     string `envconfig:"DATABASE_HOST" default:"26.35.94.218"`
	Port     int    `envconfig:"DATABASE_PORT" default:"3306"`
	Name     string `envconfig:"DATABASE_NAME" default:"naomi"`
}

//this is very important
type Fixtures struct {
	Path       string `envconfig:"CERT_PATH" default:"config/cert.pem"`
	PrivateKey string `envconfig:"PRIVATE_KEY_PATH" default:"config/key.pem"`
}


func Initialize() {
	if err := envconfig.Process("", &General); err != nil {
		logrus.WithError(err).Fatal("config: Initialize values for General")
	}
	if err := envconfig.Process("", &Database); err != nil {
		logrus.WithError(err).Fatal("config: Initialize values for Database")
	}
	if err := envconfig.Process("", &Cert); err != nil {
		logrus.Fatal(err)
	}
}

// LogLevel parses a default log level from a string
func LogLevel() logrus.Level {
	lvl, err := logrus.ParseLevel(General.LogLevel)
	if err != nil {
		logrus.WithError(err).Fatal("config: Parse log level")
	}
	return lvl
}

func bindAddr(addr string, port int) string {
	return fmt.Sprintf("%s:%d", addr, port)
}

func FeslClientAddr() string {
	return bindAddr(General.GameSpyIP, General.FeslClientPort)
}

func FeslServerAddr() string {
	return bindAddr(General.GameSpyIP, General.FeslServerPort)
}

func ThtrClientAddr() string {
	return bindAddr(General.GameSpyIP, General.ThtrClientPort)
}

func ThtrServerAddr() string {
	return bindAddr(General.GameSpyIP, General.ThtrServerPort)
}
