package utils

const (
	CreateUserQuery = `INSERT INTO users (name, username, position, salary, password) VALUES ($1, $2, $3, $4, $5)`

	GetAllUsersQuery        = `SELECT * FROM users`
	UpdateUserQuery         = `UPDATE users SET name = $1, position = $2, salary = $3 WHERE id = $4`
	DeleteUserQuery         = `DELETE FROM users WHERE id=$1`
	FindUserByUsernameQuery = `SELECT * FROM users WHERE username = $1`
	UpdatePasswordQuery     = `UPDATE users SET password = $1 WHERE id = $2`

	GetAllMusicQuery  = `SELECT * FROM music`
	GetMusicByIDQuery = `SELECT * FROM music WHERE id = $1`
	CreateMusicQuery  = `INSERT INTO music (title, artist, album, genre, duration) VALUES ($1, $2, $3, $4, $5)`
	UpdateMusicQuery  = `UPDATE music SET title = $1, artist = $2, album = $3, genre = $4, duration = $5 WHERE id = $6`
	DeleteMusicQuery  = `DELETE FROM music WHERE id = $1`

	GetMyPlaylistsQuery = ` SELECT 
        p.id, 
        p.name, 
        p.user_id, 
        COALESCE(json_agg(json_build_object(
            'id', m.id,
            'title', m.title,
            'artist', m.artist,
            'album', m.album,
            'genre', m.genre,
            'duration', m.duration
        )) FILTER (WHERE m.id IS NOT NULL), '[]') AS music
    FROM 
        playlists p
    LEFT JOIN 
        playlist_music pm ON p.id = pm.playlist_id
    LEFT JOIN 
        music m ON pm.music_id = m.id
    WHERE 
        p.user_id = $1
    GROUP BY 
        p.id;`
	CreatePlaylistQuery     = `INSERT INTO playlists (name, user_id) VALUES ($1, $2)`
	DeletePlaylistQuery     = `DELETE FROM playlists WHERE id = $1`
	AddMusicToPlaylistQuery = `INSERT INTO playlist_music (playlist_id, music_id) VALUES ($1, $2)`
)
