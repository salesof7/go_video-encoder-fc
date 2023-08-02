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

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"
	repoJob.Update(job)

	foundJob, err := repoJob.Find(job.ID)
	require.NotEmpty(t, foundJob.ID)
	require.Nil(t, err)
	require.Equal(t, foundJob.Status, job.Status)
}
