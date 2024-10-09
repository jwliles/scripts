use crate::db::{get_file_hash, update_file_hash};
use crate::summary_generator::generate_summary; // New summary generation function
use log::{info};
use memmap2::Mmap;
use rusqlite::Transaction;
use sha2::{Digest, Sha256};
use std::collections::HashMap;
use std::fs::File;
use std::io;
use tree_sitter::{Language, Parser};

// Adding Tree-sitter for Python (add more languages as needed)
extern "C" { fn tree_sitter_python() -> Language; }

const SCAN_DATE: &str = "2024-10-07";  // Placeholder for scan date (replace with actual logic)
const DEFAULT_FOLDER_NAME: &str = "root";
const DEFAULT_FOLDER_PATH: &str = "unknown";

/// Computes the SHA-256 hash of the given file.
///
/// # Arguments
/// * `file_path` - The path to the file.
///
/// # Returns
/// A `Result` containing the hash as a `String` or an `io::Error` if the file can't be accessed or read.
fn calculate_file_hash(file_path: &str) -> Result<String, io::Error> {
    let file = File::open(file_path)?;
    let mmap = unsafe { Mmap::map(&file)? };
    let mut hasher = Sha256::new();
    hasher.update(&mmap);
    Ok(format!("{:x}", hasher.finalize()))
}

/// Processes a file by computing its hash, checking if it has changed, and generating a summary.
///
/// # Arguments
/// * `file_path` - The path to the file being processed.
/// * `tx` - The current database transaction.
///
/// # Returns
/// A `Result` indicating success or an `io::Error` if there was an issue.
pub fn process_file(file_path: &str, tx: &Transaction) -> Result<(), io::Error> {
    if let Err(e) = compute_and_process_hash(file_path, tx) {
        return Err(io::Error::new(io::ErrorKind::Other, e));
    }
    Ok(())
}

fn compute_and_process_hash(file_path: &str, tx: &Transaction) -> Result<(), String> {
    let current_hash = calculate_file_hash(file_path)
        .map_err(|e| format!("Failed to compute hash for file {}: {}", file_path, e))?;

    info!("Processing file: {}", file_path);

    match get_file_hash(tx, file_path) {
        Some(stored_hash) if stored_hash == current_hash => {
            info!("No change detected for file: {}", file_path);
        }
        _ => {
            info!("File has changed or is new, updating hash and generating summary for file: {}", file_path);

            update_file_hash(tx, file_path, &current_hash)
                .map_err(|e| format!("Failed to update file hash for {}: {}", file_path, e))?;

            parse_and_generate_summary(file_path)
                .map_err(|e| format!("Failed to generate summary for {}: {}", file_path, e))?;
        }
    }
    Ok(())
}

/// Parses the file using Tree-sitter and generates a summary
fn parse_and_generate_summary(file_path: &str) -> Result<(), String> {
    let folder_name = std::path::Path::new(file_path)
        .parent()
        .and_then(|p| p.file_name())
        .and_then(|f| f.to_str())
        .unwrap_or(DEFAULT_FOLDER_NAME);

    let folder_path = std::path::Path::new(file_path)
        .parent()
        .and_then(|p| p.to_str())
        .unwrap_or(DEFAULT_FOLDER_PATH);

    // Initialize Tree-sitter parser
    let language = unsafe { tree_sitter_python() };  // You can add other languages similarly
    let mut parser = Parser::new();
    parser.set_language(language).map_err(|e| format!("Error loading Python grammar: {}", e))?;

    // Read the file content
    let content = std::fs::read_to_string(file_path)
        .map_err(|e| format!("Error reading file: {}", e))?;

    // Parse the file content using Tree-sitter
    let tree = parser.parse(&content, None)
        .ok_or("Failed to parse content")?;

    // Log the parsed syntax tree (for debugging purposes)
    println!("Parsed tree: {:?}", tree);

    // Walk through the syntax tree to generate a summary (you'll replace this with actual logic)
    let root_files = vec![];  // Placeholder for root files (replace with actual data)
    let folder_map: HashMap<String, Vec<String>> = HashMap::new();  // Placeholder for folder map
    let directory_count = 1;  // Placeholder for directory count (replace with actual value)
    let file_count = 1;  // Placeholder for file count (replace with actual value)

    // Now call generate_summary with the correct arguments
    generate_summary (
        folder_name,
        folder_path,
        SCAN_DATE,
        &root_files,
        &folder_map,
        directory_count,
        file_count,
    );

    Ok(())
}
