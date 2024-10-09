use crate::change_detector::process_file;
use crate::metrics::Metrics;
use crate::readme_writer::generate_readme; // Delegates README writing to this module
use chrono::{FixedOffset, Utc};
use indicatif::ProgressBar; // Adding the progress bar
use indicatif::ProgressStyle;
use log::{error, info, warn};
use rayon::prelude::*;
use rusqlite::Connection;
use std::collections::HashMap;
use std::fs;
use std::sync::{
    atomic::{AtomicU64, Ordering},
    Arc, Mutex,
};
use walkdir::{DirEntry, WalkDir}; // For handling time zone offsets

const UTC_OFFSET_SECONDS: i32 = 5 * 3600;

fn is_hidden(entry: &DirEntry) -> bool {
    entry
        .file_name()
        .to_str()
        .map(|s| s.starts_with('.'))
        .unwrap_or(false)
}

pub fn scan_directory(
    path: &str,
    metrics: Arc<Mutex<Metrics>>,
    _db_path_str: &str,
    file_count: Arc<AtomicU64>,
    _conn: Arc<Mutex<Connection>>,
) {
    info!("Starting to scan directory: {}", path);

    let offset =
        FixedOffset::west_opt(UTC_OFFSET_SECONDS).expect("Failed to create time zone offset");
    let scan_date = Utc::now()
        .with_timezone(&offset)
        .format("%Y-%m-%d")
        .to_string();

    let directory_count = Arc::new(Mutex::new(0));
    let folder_map = Arc::new(Mutex::new(HashMap::new()));
    let root_files = Arc::new(Mutex::new(Vec::new()));

    let progress_bar = ProgressBar::new_spinner();
    progress_bar.set_style(
        ProgressStyle::default_spinner()
            .tick_strings(&["⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"])
            .template("{spinner:.green} [{elapsed}] Scanning... {pos} files processed")
            .to_owned(),
    );

    WalkDir::new(path)
        .into_iter()
        .filter_entry(|e| !is_hidden(e)) // Skip hidden files and directories
        .filter_map(|e| e.ok())
        .par_bridge()
        .for_each(|entry| {
            let file_path = match entry.path().strip_prefix(path) {
                Ok(relative_path) => relative_path,
                Err(_e) => {
                    error!(
                        "Failed to strip prefix from path: {}",
                        entry.path().display()
                    );
                    return;
                }
            };

            let file_path_str = match file_path.to_str() {
                Some(path) => path,
                None => {
                    error!("Failed to convert path to string: {}", file_path.display());
                    return;
                }
            };

            info!("Processing file: {}", file_path_str);

            if entry.file_type().is_file() {
                if fs::metadata(entry.path()).is_err() {
                    warn!("File not found or inaccessible: {}", file_path_str);
                    return; // Skip this file if it doesn't exist
                }
                file_count.fetch_add(1, Ordering::SeqCst);
                if let Ok(mut metrics_lock) = metrics.lock() {
                    metrics_lock.increment_files_scanned();
                } else {
                    warn!("Failed to lock metrics for updating file count.");
                }
                let file_count_value = file_count.load(Ordering::SeqCst);
                if file_count_value % 100 == 0 {
                    progress_bar.inc(100);
                }
                let mut conn_guard = _conn.lock().expect("Failed to lock database connection");
                let tx = conn_guard
                    .transaction()
                    .expect("Failed to create transaction");
                if let Err(e) = process_file(file_path_str, &tx) {
                    error!("Failed to process file {}: {}", file_path_str, e);
                }
                tx.commit().expect("Failed to commit transaction");
            }
        });

    progress_bar.finish_with_message("Scanning completed");

    let folder_name = path.split('/').last().unwrap_or("notes");
    generate_readme(
        folder_name,
        path,
        &scan_date,
        &root_files.lock().expect("Failed to lock root_files"),
        &folder_map.lock().expect("Failed to lock folder_map"),
        *directory_count
            .lock()
            .expect("Failed to lock directory_count"),
        file_count.load(Ordering::SeqCst),
    );
}
