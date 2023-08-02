package repositories_test

import (
	"testing"
	"time"

	"github.com/salesof7/go_video-encoder/app/repositories"
	"github.com/salesof7/go_video-encoder/domain"
	"github.com/salesof7/go_video-encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	videoFound, err := repo.Find(video.ID)
	require.NotEmpty(t, videoFound.ID)
	require.Nil(t, err)
	require.Equal(t, videoFound.ID, video.ID)
}
