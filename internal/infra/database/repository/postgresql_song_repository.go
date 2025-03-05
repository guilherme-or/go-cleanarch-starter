package database_repository

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

var (
	queryFindByDuration = `
		SELECT
			so.id, so.title, so.duration, so.track_number,
			al.id, al.title, al.release_date
		FROM songs so
		INNER JOIN albums al ON so.album_id = al.id
		WHERE so.duration >= $1 AND so.duration <= $2
	`
	queryFindByAlbum = `
		SELECT
			so.id, so.title, so.duration, so.track_number
		FROM songs so
		INNER JOIN albums al ON so.album_id = al.id
		WHERE so.album_id = $1
	`
	queryFindByArtist = `
		SELECT
			so.id, so.title, so.duration, so.track_number
		FROM songs so
		INNER JOIN albums al ON so.album_id = al.id
		INNER JOIN artists ar ON al.artist_id = ar.id
		WHERE ar.id = $1
	`
)

type PostgreSQLSongRepository struct {
	db *sql.DB
}

func NewPostgreSQLSongRepository(db *sql.DB) repository.SongRepository {
	return &PostgreSQLSongRepository{db: db}
}

// FindByDuration implements repository.SongRepository.
func (p *PostgreSQLSongRepository) FindByDuration(startDuration, endDuration int) ([]*entity.Song, error) {
	rows, err := p.db.Query(queryFindByDuration, startDuration, endDuration)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*entity.Song
	for rows.Next() {
		var song entity.Song
		var album entity.Album
		if err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.Duration,
			&song.TrackNumber,
			&album.ID,
			&album.Title,
			&album.ReleaseDate,
		); err != nil {
			return nil, err
		}
		song.Album = &album
		songs = append(songs, &song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

// FindByAlbum implements repository.SongRepository.
func (p *PostgreSQLSongRepository) FindByAlbum(albumID int) ([]*entity.Song, error) {
	rows, err := p.db.Query(queryFindByAlbum, albumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*entity.Song
	for rows.Next() {
		var song entity.Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.TrackNumber); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

// FindByArtist implements repository.SongRepository.
func (p *PostgreSQLSongRepository) FindByArtist(artistID int) ([]*entity.Song, error) {
	rows, err := p.db.Query(queryFindByArtist, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*entity.Song
	for rows.Next() {
		var song entity.Song
		if err := rows.Scan(&song.ID, &song.Title, &song.Duration, &song.TrackNumber); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
