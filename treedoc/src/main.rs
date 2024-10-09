mod change_detector;
mod db;
mod directory_scanner;
mod logger;
mod metrics;
mod readme_manager;
use crate::directory_scanner::scan_directory;
use crate::metrics::Metrics;
use rusqlite::Connection;
use std::env;
use std::path::PathBuf;
use std::sync::atomic::AtomicU64;
use std::sync::{Arc, Mutex};
use std::time::Instant;
mod summary_generator;

fn main() {
    logger::init_logger();

    let args: Vec<String> = env::args().collect();
    let path = if args.len() > 1 { &args[1] } else { "." };

    log::info!("Using directory: {}", path);

    // Construct the correct database path
    let mut db_path = PathBuf::from(path);
    db_path.push("file_hashes.db"); // Append file_hashes.db to the directory

    let db_path_str = db_path.to_str().unwrap(); // Convert PathBuf to &str

    // Initialize the database at this path
    let _conn = match db::init_db(db_path_str) {
        Ok(c) => {
            log::info!("Database initialized successfully.");
            Arc::new(Mutex::new(c))
        }
        Err(e) => {
            log::error!("Failed to initialize database: {}", e);
            return;
        }
    };

    let metrics = Arc::new(Mutex::new(Metrics::new()));
    let file_count = Arc::new(AtomicU64::new(0));
    let db_path_str = format!("{}/file_hashes.db", path);
    let conn = Arc::new(Mutex::new(Connection::open(db_path_str.clone()).unwrap())); // Clone db_path_str here
    let start_time = Instant::now();

    // Call scan_directory with correct parameters
    scan_directory(
        &path,
        metrics.clone(),
        &db_path_str, // Pass a &str instead of String
        file_count.clone(),
        conn.clone(),
    );

    let metrics_lock = metrics.lock().unwrap();
    metrics_lock.display_metrics();

    let total_time = start_time.elapsed();

    println!(
        "The terminal reports an execution time of {:.3} seconds.",
        total_time.as_secs_f64()
    );
}
