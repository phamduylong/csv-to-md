package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log/slog"
)

type Align int

const (
	Center Align = 0
	Left   Align = 1
	Right  Align = 2
)

type Config struct {
	Align          Align
	URL            string
	InputFilePath  string
	OutputFilePath string
	AutoCopy       bool
	VerboseLogging bool
	OutputToWindow bool
}

func parseConfig() (Config, error) {
	inputPathPtr := flag.String("inputFile", "", "Path of the CSV file to convert")
	outputPathPtr := flag.String("outputFile", "",
		"Path of the output file where converted data will be stored. ⚠️ Existing content will be overwritten")
	urlPtr := flag.String("url", "", "URL of the CSV file to convert")
	alignPtr := flag.Int("align", int(Center), "How should text be aligned in the table?\n0. Center\n1. Left\n2. Right")
	autoCopyPtr := flag.Bool("autoCopy", false, "Should the converted markdown table be copied to clipboard automatically?")
	verboseLoggingPtr := flag.Bool("verboseLogging", false, "Should detailed diagnostic messages be logged?")
	outputToWindowPtr := flag.Bool("outputToWindow", false, "Whether the converted table should be rendered in this window?")
	flag.Parse()

	var cfg Config

	cfg.Align = Align(*alignPtr)
	cfg.InputFilePath = *inputPathPtr
	cfg.OutputFilePath = *outputPathPtr
	cfg.URL = *urlPtr
	cfg.AutoCopy = *autoCopyPtr
	cfg.VerboseLogging = *verboseLoggingPtr
	cfg.OutputToWindow = *outputToWindowPtr

	cfgErr := validateConfig(cfg)

	if cfgErr != nil {
		return Config{}, cfgErr
	}

	if cfg.VerboseLogging {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		// beautiful json here. 4 spaces for indentation
		jsonCfg, serializationErr := json.MarshalIndent(cfg, "", "    ")
		if serializationErr == nil {
			slog.Debug(fmt.Sprintf("Config:\n%s", jsonCfg))
		}
	}

	return cfg, nil
}

func validateConfig(cfg Config) error {
	if cfg.VerboseLogging {
		slog.Debug("Validating config 🤔")
	}
	if cfg.URL != "" && cfg.InputFilePath != "" {
		return errors.New("both URL and file path are given, please provide only one of them exclusively")
	}
	if cfg.URL == "" && cfg.InputFilePath == "" {
		return errors.New("both URL and file path are missing, please provide either one of them exclusively")
	}
	if cfg.VerboseLogging {
		slog.Debug("Config is valid ✅")
	}
	return nil
}
