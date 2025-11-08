package config

import (
	"fmt"
	"os"
	"strconv"
)

type FileConfig struct {
	SavePath         string
	MaxFileSize      int
	MaxProfileHeight int
	MaxProfileWidth  int
	MaxCoverHeight   int
	MaxCoverWidth    int
}

func LoadFileConfig() (*FileConfig, error) {
	savePath := os.Getenv("FILE_SAVE_PATH")
	if savePath == "" {
		return nil, fmt.Errorf("FILE_SAVE_PATH environment variable is required")
	}

	maxFileSize, err := getEnvAsInt("MAX_FILE_SIZE")
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_FILE_SIZE: %w", err)
	}

	maxProfileHeight, err := getEnvAsInt("MAX_PROFILE_HEIGHT")
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_PROFILE_HEIGHT: %w", err)
	}

	maxProfileWidth, err := getEnvAsInt("MAX_PROFILE_WIDTH")
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_PROFILE_WIDTH: %w", err)
	}

	maxCoverHeight, err := getEnvAsInt("MAX_COVER_HEIGHT")
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_COVER_HEIGHT: %w", err)
	}

	maxCoverWidth, err := getEnvAsInt("MAX_COVER_WIDTH")
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_COVER_WIDTH: %w", err)
	}

	return &FileConfig{
		SavePath:         savePath,
		MaxFileSize:      maxFileSize,
		MaxProfileHeight: maxProfileHeight,
		MaxProfileWidth:  maxProfileWidth,
		MaxCoverHeight:   maxCoverHeight,
		MaxCoverWidth:    maxCoverWidth,
	}, nil
}

func getEnvAsInt(key string) (int, error) {
	value := os.Getenv(key)

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("unable to parse %s as int: %w", key, err)
	}
	return parsedValue, nil
}
