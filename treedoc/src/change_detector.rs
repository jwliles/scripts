use crate::db::{get_file_hash, update_file_hash};
use crate::readme_manager::create_or_update_readme;
use log::{error, info};
use memmap2::Mmap;
use rusqlite::Transaction;
use sha2::{Digest, Sha256};
use std::fs::File;
use std::io; // Added log macros for better logging

/// Computes the SHA-256 hash of the given file.
///
/// # Arguments
/// * `file_path` - The path to the file.
///
/// # Returns
/// A `Result` containing the hash as a `String` or an `io::Error` if the file can't be accessed or read.
fn compute_hash(file_path: &str) -> Result<String, io::Error> {
    // Open the file
    let file = File::open(file_path)?;

    // Memory-map the file for fast access
    let mmap = unsafe { Mmap::map(&file)? };

    // Create a SHA-256 hasher and update it with the file contents
    let mut hasher = Sha256::new();
    hasher.update(&mmap);

    // Return the computed hash as a hexadecimal string
    Ok(format!("{:x}", hasher.finalize()))
}

/// Processes a file by computing its hash and checking if it has changed.
///
/// # Arguments
/// * `file_path` - The path to the file being processed.
/// * `tx` - The current database transaction.
///
/// # Returns
/// A `Result` indicating success or an `io::Error` if there was an issue.
pub fn process_file(file_path: &str, tx: &Transaction) -> Result<(), io::Error> {
    // Try to compute the file hash
    match compute_hash(file_path) {
        Ok(current_hash) => {
            info!("Processing file: {}", file_path);

            // Try to get the previously stored hash from the database
            match get_file_hash(tx, file_path) {
                Some(stored_hash) if stored_hash == current_hash => {
                    info!("No change detected for file: {}", file_path);
                }
                _ => {
                    info!(
                        "File has changed or is new, updating hash and README for file: {}",
                        file_path
                    );

                    // Update the file hash in the database
                    if let Err(e) = update_file_hash(tx, file_path, &current_hash) {
                        error!("Failed to update file hash for {}: {}", file_path, e);
                        return Err(io::Error::new(
                            io::ErrorKind::Other,
                            format!("Failed to update file hash for {}", file_path),
                        ));
                    }

                    // Get the parent directory of the file
                    if let Some(directory) = std::path::Path::new(file_path)
                        .parent()
                        .and_then(|p| p.to_str())
                    {
                        let readme_content = format!("Updated README for directory: {}", directory);

                        // Update the README for the directory
                        if let Err(e) = create_or_update_readme(directory, &readme_content) {
                            error!("Failed to update README for directory {}: {}", directory, e);
                        }
                    } else {
                        error!(
                            "Could not determine parent directory for file: {}",
                            file_path
                        );
                    }
                }
            }
        }
        Err(e) => {
            // Log and propagate the error if hash computation failed
            error!("Failed to compute hash for file {}: {}", file_path, e);
            return Err(io::Error::new(
                io::ErrorKind::Other,
                format!("Failed to compute hash for file {}", file_path),
            ));
        }
    }
    Ok(())
}
