CREATE TABLE files (
    id VARCHAR(64) PRIMARY KEY,
    source_id VARCHAR(32) NOT NULL,
    title TEXT NOT NULL,
    extension VARCHAR(10) NOT NULL,
    resolution INTEGER NOT NULL,
    audio_bitrate INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
