// Package config provides application configuration for all required environments,
// structured as Go types, read in from YAML files. Absolutely no secrets are to be
// stored in these files, or in source control in general. All secrets should be
// loaded into an environment, via remote manager, securely stored in a hosted
// platform, or simply managed locally outside of source control. Each env may have
// its variables injected differently, based on the environment's specific needs/usage.
package config

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var Conf Config

type GrpcServer struct {
	Port string
}

type HTTPServer struct {
	Port string
}

type Logger struct {
	Level string
}

type ScoreFunction struct {
	SourcePoints      []float32
	DestinationPoints []float32
}

type ModelConfig struct {
	Address    string
	Dimentions int
	ScoreFunction
}

type MetricServer struct {
	Port string
}

type VssQServer struct {
	Address                string
	HnswEf                 uint64
	ShardNumber            uint32
	ReplicationFactor      uint32
	WriteConsistencyFactor uint32
}

// Config structures the environment configuration which is read
// in from a YAML file. The file contents should match the structure
// of this type
type Config struct {
	GrpcServer
	HTTPServer
	Logger
	MetricServer
	JcvFaceModels map[string]ModelConfig
	VssQServer
	KnnSearchEngine string
}

// Setup reads the environment file based on the application env,
// and populates a Config instance. Othwerwise this function kills
// the running process if any errors occur
func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./env")
	// This search path is used for testing.
	viper.AddConfigPath("../env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(errors.Wrap(err, "failed to read config file"))
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatal(errors.Wrap(err, "failed to unmarshal config file"))
	}
}
