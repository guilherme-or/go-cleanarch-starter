package database_repository

import (
	"database/sql"

	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/entity"
	"github.com/guilherme-or/go-cleanarch-starter/internal/domain/repository"
)

var (
	queryFindAllArtists = `
		SELECT
			ar.id, ar.name, ar.birth_date, ar.nationality
		FROM artists ar
	`
	queryFindArtistByName = `
		SELECT
			ar.id, ar.name, ar.birth_date, ar.nationality,
			al.id, al.title, al.release_date
		FROM artists ar
		LEFT JOIN albums al
		ON ar.id = al.artist_id
		WHERE ar.name = $1
	`
	queryFindArtistByNationality = `
		SELECT
			ar.id, ar.name, ar.birth_date, ar.nationality
		FROM artists ar
		WHERE ar.nationality = $1
	`
)

type PostgreSQLArtistRepository struct {
	db *sql.DB
}

func NewPostgreSQLArtistRepository(db *sql.DB) repository.ArtistRepository {
	return &PostgreSQLArtistRepository{db: db}
}

// FindAll implements repository.ArtistRepository.
func (p *PostgreSQLArtistRepository) FindAll() ([]*entity.Artist, error) {
	rows, err := p.db.Query(queryFindAllArtists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []*entity.Artist
	for rows.Next() {
		var artist entity.Artist
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.BirthDate, &artist.Nationality); err != nil {
			return nil, err
		}
		artists = append(artists, &artist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}

// FindByName implements repository.ArtistRepository.
func (p *PostgreSQLArtistRepository) FindByName(name string) (*entity.Artist, error) {
	rows, err := p.db.Query(queryFindArtistByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artist entity.Artist
	var albums []*entity.Album
	for rows.Next() {
		var album entity.Album
		if err := rows.Scan(
			&artist.ID,
			&artist.Name,
			&artist.BirthDate,
			&artist.Nationality,
			&album.ID,
			&album.Title,
			&album.ReleaseDate,
		); err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}

	artist.Albums = albums

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &artist, nil
}

// FindByNationality implements repository.ArtistRepository.
func (p *PostgreSQLArtistRepository) FindByNationality(nationality string) ([]*entity.Artist, error) {
	rows, err := p.db.Query(queryFindArtistByNationality, nationality)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []*entity.Artist
	for rows.Next() {
		var artist entity.Artist
		if err := rows.Scan(
			&artist.ID,
			&artist.Name,
			&artist.BirthDate,
			&artist.Nationality,
		); err != nil {
			return nil, err
		}
		artists = append(artists, &artist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artists, nil
}
