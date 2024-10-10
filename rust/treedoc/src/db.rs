use log::error;
use rusqlite::{params, Connection, Result, Transaction};

pub fn init_db(db_path: &str) -> Result<Connection> {
    log::info!("Initializing database at path: {}", db_path);

    let conn = match Connection::open(db_path) {
        Ok(c) => {
            log::info!("Database connection opened at: {}", db_path);
            c
        }
        Err(e) => {
            error!("Failed to open database at {}: {}", db_path, e);
            return Err(e);
        }
    };

    if let Err(e) = conn.execute(
        "CREATE TABLE IF NOT EXISTS file_hashes (
            id INTEGER PRIMARY KEY,
            file_path TEXT NOT NULL,
            file_hash TEXT NOT NULL
        )",
        [],
    ) {
        error!("Failed to create table in {}: {}", db_path, e);
        return Err(e);
    }

    log::info!(
        "Table 'file_hashes' created or already exists in: {}",
        db_path
    );
    Ok(conn)
}

pub fn get_file_hash(tx: &Transaction, file_path: &str) -> Option<String> {
    let mut stmt = tx
        .prepare("SELECT file_hash FROM file_hashes WHERE file_path = ?")
        .ok()?;
    let hash_result: Result<String, _> = stmt.query_row([file_path], |row| row.get(0));
    hash_result.ok()
}

pub fn update_file_hash(tx: &Transaction, file_path: &str, file_hash: &str) -> Result<()> {
    tx.execute(
        "INSERT OR REPLACE INTO file_hashes (file_path, file_hash) VALUES (?, ?)",
        params![file_path, file_hash],
    )
    .map_err(|e| {
        error!("Failed to update file hash for {}: {}", file_path, e);
        e
    })?;
    Ok(())
}
