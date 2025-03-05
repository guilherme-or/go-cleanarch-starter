package database_repository

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

var (
	queryFindAlbumByTitle = `
		SELECT
			al.id, al.title, al.release_date,
			ar.id, ar.name, ar.birth_date, ar.nationality,
			so.id, so.title, so.duration, so.track_number
		FROM albums al
		INNER JOIN songs so ON so.album_id = al.id
		INNER JOIN artists ar ON al.artist_id = ar.id
		WHERE al.title = $1
	`
	queryFindAlbumsByArtist = `
		SELECT
			al.id, al.title, al.release_date
		FROM albums al
		INNER JOIN artists ar
		ON al.artist_id = ar.id
		WHERE ar.id = $1
	`
	queryFindAlbumsByYear = `
		SELECT
			al.id, al.title, al.release_date
		FROM albums al
		WHERE EXTRACT(YEAR FROM al.release_date) = $1
	`
)

type PostgreSQLAlbumRepository struct {
	db *sql.DB
}

func NewPostgreSQLAlbumRepository(db *sql.DB) repository.AlbumRepository {
	return &PostgreSQLAlbumRepository{db: db}
}

// FindByTitle implements repository.AlbumRepository.
func (p *PostgreSQLAlbumRepository) FindByTitle(title string) (*entity.Album, error) {
	rows, err := p.db.Query(queryFindAlbumByTitle, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var album entity.Album
	var artist entity.Artist
	var songs []*entity.Song
	for rows.Next() {
		var song entity.Song
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.ReleaseDate,
			&artist.ID,
			&artist.Name,
			&artist.BirthDate,
			&artist.Nationality,
			&song.ID,
			&song.Title,
			&song.Duration,
			&song.TrackNumber,
		); err != nil {
			return nil, err
		}
		songs = append(songs, &song)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	album.Artist = &artist
	album.Songs = songs

	return &album, nil
}

// FindByArtist implements repository.AlbumRepository.
func (p *PostgreSQLAlbumRepository) FindByArtist(artistID int) ([]*entity.Album, error) {
	rows, err := p.db.Query(queryFindAlbumsByArtist, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []*entity.Album
	for rows.Next() {
		var album entity.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.ReleaseDate,
		); err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

// FindByYear implements repository.AlbumRepository.
func (p *PostgreSQLAlbumRepository) FindByYear(year int) ([]*entity.Album, error) {
	rows, err := p.db.Query(queryFindAlbumsByYear, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []*entity.Album
	for rows.Next() {
		var album entity.Album
		if err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.ReleaseDate,
		); err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	if err := rows.Err(); err != nil {
		return nil, rows.Err()
	}

	return albums, nil
}
