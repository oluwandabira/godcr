package config

import (
	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Network          string `long:"network" description:"Network to use"`
	HomeDir          string `long:"appdata" description:"Directory where the app configuration file and wallet data is stored"`
	ConfigFile       string `long:"configfile" description:"Filename of the config file in the app directory"`
	ShowVersion      bool   `short:"V" long:"version" description:"Display version information and exit"`
	MaxLogZips       int    `long:"max-log-zips" description:"The number of zipped log files created by the log rotator to be retained. Setting to 0 will keep all."`
	LogDir           string `long:"logdir" description:"Directory to log output."`
	DebugLevel       string `short:"d" long:"debuglevel" description:"Logging level {trace, debug, info, warn, error, critical}"`
	Quiet            bool   `short:"q" long:"quiet" description:"Easy way to set debuglevel to error"`
	SpendUnconfirmed bool   `long:"spendunconfirmed" description:"Allow the multiwallet to use transactions that have not been confirmed"`
}

func (c *Config) ParseFile(configfile string) error {
	parser := flags.NewParser(c, flags.None)
	fparser := flags.NewIniParser(parser)
	fparser.ParseFile(configfile)
	return nil
}
