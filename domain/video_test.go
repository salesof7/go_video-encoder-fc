package domain_test

import (
	"testing"
	"time"

	"github.com/salesof7/go_video-encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "not_uuid"
	video.ResourceID = "any_resource"
	video.FilePath = "any_path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.ResourceID = "any_resource"
	video.FilePath = "any_path"
	video.CreatedAt = time.Now()

	err := video.Validate()
	require.Nil(t, err)
}
