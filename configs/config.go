package configs

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nakiner/guestcovider/internal/database"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ServiceName = "guestcovider"

var options = []option{
	{"config", "string", "", "config file"},

	{"server.http.port", "int", 8080, "server http port"},
	{"server.http.timeout_sec", "int", 86400, "server http connection timeout"},
	{"server.grpc.port", "int", 9194, "server grpc port"},
	{"server.grpc.timeout_sec", "int", 86400, "server grpc connection timeout"},

	{"postgres.host", "string", "localhost", "postgres master host"},
	{"postgres.port", "int", 5432, "postgres master port"},
	{"postgres.user", "string", "guestcovider", "postgres master user"},
	{"postgres.password", "string", "guestcovider", "postgres master password"},
	{"postgres.database_name", "string", "guestcovider", "postgres master database name"},
	{"postgres.secure", "string", "disable", "postgres master SSL support"},

	{"logger.level", "string", "emerg", "Level of logging. A string that correspond to the following levels: emerg, alert, crit, err, warning, notice, info, debug"},
	{"logger.time_format", "string", "2006-01-02T15:04:05.999999999", "Date format in logs"},

	{"sentry.enabled", "bool", false, "Enables or disables sentry"},
	{"sentry.dsn", "string", "https://hasbool@sentry.com/1", "Data source name. Sentry addr"},
	{"sentry.environment", "string", "local", "The environment to be sent with events."},

	{"tracer.enabled", "bool", false, "Enables or disables tracing"},
	{"tracer.host", "string", "127.0.0.1", "The tracer host"},
	{"tracer.port", "int", 5775, "The tracer port"},
	{"tracer.name", "string", "export", "The tracer name"},

	{"metrics.enabled", "bool", false, "Enables or disables metrics"},
	{"metrics.port", "int", 9153, "server http port"},

	{"limiter.enabled", "bool", false, "Enables or disables limiter"},
	{"limiter.limit", "float64", 10000.0, "Limit tokens per second"},
}

type Config struct {
	Server struct {
		GRPC struct {
			Port       int
			TimeoutSec int `mapstructure:"timeout_sec"`
		}
		HTTP struct {
			Port       int
			TimeoutSec int `mapstructure:"timeout_sec"`
		}
	}
	Logger struct {
		Level      string
		TimeFormat string `mapstructure:"time_format"`
	}
	Sentry struct {
		Enabled     bool
		Dsn         string
		Environment string
	}
	Tracer struct {
		Enabled bool
		Host    string
		Port    int
		Name    string
	}
	Metrics struct {
		Enabled bool
		Port    int
	}
	Limiter struct {
		Enabled bool
		Limit   float64
	}
	Postgres database.Config
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}

// NewConfig returns and prints struct with config parameters
func NewConfig() *Config {
	return &Config{}
}

// read gets parameters from environment variables, flags or file.
func (c *Config) Read() error {
	viper.SetEnvPrefix(ServiceName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	for _, o := range options {
		switch o.typing {
		case "string":
			pflag.String(o.name, o.value.(string), o.description)
		case "int":
			pflag.Int(o.name, o.value.(int), o.description)
		case "bool":
			pflag.Bool(o.name, o.value.(bool), o.description)
		case "float64":
			pflag.Float64(o.name, o.value.(float64), o.description)
		default:
			viper.SetDefault(o.name, o.value)
		}
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	if fileName := viper.GetString("config"); fileName != "" {
		viper.SetConfigFile(fileName)
		viper.SetConfigType("toml")

		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to read from file")
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}
	return nil
}

func (c *Config) GenerateMdTable() error {
	var t string

	t += "Command line | Environment | Default |Description"
	t += fmt.Sprintln()
	t += "--- | --- | --- | ---"
	t += fmt.Sprintln()

	for _, o := range options {
		t += o.name
		t += " | "
		t += strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(ServiceName+"_"+o.name, ".", "_"), "-", "_"))
		t += " | "
		t += fmt.Sprintf("%v", o.value)
		t += " | "
		t += o.description
		t += fmt.Sprintln()
	}
	fmt.Fprintln(os.Stdout, t)
	return nil
}

func (c *Config) GenerateEnvironment() error {
	var t string

	for _, o := range options {
		t += strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(ServiceName+"_"+o.name, ".", "_"), "-", "_"))
		t += ": "
		t += fmt.Sprintf("%v", o.value)
		t += fmt.Sprintln()
	}
	fmt.Fprintln(os.Stdout, t)
	return nil
}

func (c *Config) GenerateFromTask() error {
	var t string

	t += "| Command line | Environment |"
	t += fmt.Sprintln()
	for _, o := range options {
		t += "| "
		t += strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(ServiceName+"_"+o.name, ".", "_"), "-", "_"))
		t += " | "
		t += fmt.Sprintf("%v", o.value)
		t += " |"
		t += fmt.Sprintln()
	}

	fmt.Fprintln(os.Stdout, t)
	return nil
}

func (c Config) Print() error {
	c.Postgres.Password = "******"

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, string(b))
	return nil
}
