CREATE TABLE IF NOT EXISTS mentorship (
    id serial PRIMARY KEY,
    mahasiswa_id INTEGER NOT NULL,
    dosen_id INTEGER NOT NULL,
    FOREIGN KEY (mahasiswa_id) REFERENCES users(id),
    FOREIGN KEY (dosen_id) REFERENCES users(id)
    );
