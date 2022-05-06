package sqlstore

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/utils"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

func (s *SQLStore) SaveFileInfo(id, fileName, extension string, size int64) error {
	query := s.getQueryBuilder(s.db).
		Insert(s.tablePrefix+"file_info").
		Columns(
			"id",
			"create_at",
			"name",
			"extension",
			"size",
		).
		Values(
			id,
			utils.GetMillis(),
			fileName,
			extension,
			size,
		)

	if _, err := query.Exec(); err != nil {
		s.logger.Error(
			"failed to save fileinfo",
			mlog.String("file_name", fileName),
			mlog.String("extension", extension),
			mlog.Int64("size", size), mlog.Err(err),
		)
		return err
	}

	return nil
}

func (s *SQLStore) GetFileInfo(id string) (*model.FileInfo, error) {
	query := s.getQueryBuilder(s.db).
		Select(
			"id",
			"create_at",
			"delete_at",
			"name",
			"extension",
			"size",
		).
		From(s.tablePrefix + "file_info").
		Where(sq.Eq{"Id": id})

	row := query.QueryRow()

	fileInfo := model.FileInfo{}

	err := row.Scan(
		&fileInfo.Id,
		&fileInfo.CreateAt,
		&fileInfo.DeleteAt,
		&fileInfo.Name,
		&fileInfo.Extension,
		&fileInfo.Size,
	)

	if err != nil {
		s.logger.Error("error scanning fileinfo row", mlog.String("id", id), mlog.Err(err))
		return nil, err
	}

	return &fileInfo, nil
}
