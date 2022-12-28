package flags

import "github.com/urfave/cli/v2"

func Flags() []cli.Flag {
	var flags = []cli.Flag{}

	flags = append(flags, region...)
	flags = append(flags, env...)
	flags = append(flags, dbconnflags...)
	flags = append(flags, http...)
	flags = append(flags, log...)

	return flags
}

// region
var (
	Region string

	region = []cli.Flag{
		&cli.StringFlag{
			Name:        "region",
			Usage:       "set deployment region",
			EnvVars:     []string{"REGION"},
			Value:       "default",
			Destination: &Region,
		},
	}

	Env string

	env = []cli.Flag{
		&cli.StringFlag{
			Name:        "env",
			Usage:       "set deployment env",
			EnvVars:     []string{"ENV"},
			Value:       "default",
			Destination: &Env,
		},
	}
)

// HTTP Flags
var (
	HTTPHost string
	HTTPPort string

	http = []cli.Flag{
		&cli.StringFlag{
			Name:        "http.host",
			Usage:       "set http host name",
			EnvVars:     []string{"DOKLA_HTTP_HOST"},
			Value:       "0.0.0.0",
			Destination: &HTTPHost,
		},
		&cli.StringFlag{
			Name:        "http.port",
			Usage:       "set http port name",
			EnvVars:     []string{"DOKLA_HTTP_PORT"},
			Value:       "9090",
			Destination: &HTTPPort,
		},
	}
)

// Log Cli Flags
var (
	LogLevel    string
	LogEncoding string
	LogOutput   string

	log = []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Usage:       "logging level (info, error, debug, warn)",
			Value:       "info",
			EnvVars:     []string{"DOKLA_LOG_LEVEL"},
			Destination: &LogLevel,
			Aliases:     []string{"level", "l"},
		},
		&cli.StringFlag{
			Name:        "log.enc",
			Usage:       "logging encoding (json, console)",
			Value:       "json",
			EnvVars:     []string{"DOKLA_LOG_ENC"},
			Destination: &LogEncoding,
		},
		&cli.StringFlag{
			Name:        "log.output",
			Usage:       "logging output (stdout, stderr, <filesystem path>)",
			EnvVars:     []string{"DOKLA_LOG_OUTPUT"},
			Value:       "stdout",
			Destination: &LogOutput,
		},
	}
)

var (
	DbConnTimeout  int
	DbName         string
	DbRootUser     string
	DbRootPassword string
	DbReadTimeout  int
	DbWriteTimeout int
	DbConn         string

	dbconnflags = []cli.Flag{
		&cli.StringFlag{
			Name:        "mongo.conn",
			Value:       "mongodb://mongodb-sharded:27017/?authSource=admin",
			Usage:       "set the connection string for mongodb",
			DefaultText: "--mongo.conn=mongodb://mongodb-sharded:27017?authSource=admin",
			EnvVars:     []string{"DOKLA_MONGO_CONN"},
			Destination: &DbConn,
		},
		&cli.StringFlag{
			Name:        "mongo.dbname",
			Value:       "posts",
			Usage:       "set the database name for mongodb",
			DefaultText: "--mongo.dbname=posts",
			EnvVars:     []string{"DOKLA_MONGO_DB_NAME"},
			Destination: &DbName,
		},
		&cli.IntFlag{
			Name:        "mongo.conn.timeout",
			Value:       30000,
			Usage:       "set the connection timeout for mongo",
			DefaultText: "--mongo.conn.timeout=30000",
			EnvVars:     []string{"DOKLA_MONGO_CONN_TIMEOUT"},
			Destination: &DbConnTimeout,
		},
		&cli.IntFlag{
			Name:        "mongo.read.timeout",
			Value:       200,
			Usage:       "set the read timeout for mongo",
			EnvVars:     []string{"DOKLA_MONGO_READ_TIMEOUT"},
			Destination: &DbReadTimeout,
		},
		&cli.IntFlag{
			Name:        "mongo.write.timeout",
			Value:       300,
			Usage:       "set the write timeout for mongo",
			EnvVars:     []string{"DOKLA_MONGO_WRITE_TIMEOUT"},
			Destination: &DbWriteTimeout,
		},
		&cli.StringFlag{
			Name:        "mongo.root.user",
			Value:       "root",
			Usage:       "set the root user for mongodb",
			DefaultText: "--mongo.root.user=root",
			EnvVars:     []string{"DOKLA_MONGO_ROOT_USER"},
			Destination: &DbRootUser,
		},
		&cli.StringFlag{
			Name:        "mongo.root.password",
			Value:       "rootaa",
			Usage:       "set the root password for mongodb",
			DefaultText: "--mongo.root.password=rootaa",
			EnvVars:     []string{"DOKLA_MONGO_ROOT_PASSWORD"},
			Destination: &DbRootPassword,
		},
	}
)
