use std::collections::HashMap;
use std::fs::OpenOptions;
use std::io::Write;

/// Generates the README.md file for a given directory.
///
/// # Arguments
/// * `folder_name` - The name of the folder (e.g., "notes", "documentation_system").
/// * `folder_path` - The path to the folder where README.md will be created.
/// * `scan_date` - The current date in ISO 8601 format.
/// * `root_files` - A list of files in the root of the folder.
/// * `folder_map` - A HashMap containing folder names and their associated files.
/// * `directory_count` - The number of subdirectories in the folder.
/// * `file_count` - The number of files in the folder.
pub fn generate_readme(
    folder_name: &str,
    folder_path: &str,
    scan_date: &str,
    root_files: &[String],
    folder_map: &HashMap<String, Vec<String>>,
    directory_count: u64,
    file_count: u64,
) {
    // Format description based on directory and file count
    let description = format!(
        "{} contains {} directories and {} files as of {}.",
        folder_name, directory_count, file_count, scan_date
    );

    // Open or create the README.md for writing
    let readme_path = format!("{}/README.md", folder_path);
    let mut readme_file = OpenOptions::new()
        .write(true)
        .create(true)
        .truncate(true) // Replace the existing content
        .open(&readme_path)
        .expect("Failed to open README.md");

    // Write the header and folder description
    writeln!(readme_file, "# README").expect("Failed to write to README.md");
    writeln!(readme_file, "\n{}", description).expect("Failed to write to README.md");
    writeln!(readme_file, "\n---\n").expect("Failed to write to README.md");

    // Write the categories (folders in the root)
    if !folder_map.is_empty() {
        writeln!(readme_file, "## Categories\n").expect("Failed to write to README.md");
        for folder in folder_map.keys() {
            writeln!(readme_file, "- [{}](#{})", folder, folder)
                .expect("Failed to write to README.md");
        }
    }

    writeln!(readme_file, "\n---\n").expect("Failed to write to README.md");

    // Write the Projects section (files in the root directory)
    if !root_files.is_empty() {
        writeln!(readme_file, "### Projects\n").expect("Failed to write to README.md");
        for file_link in root_files {
            writeln!(readme_file, "{}", file_link).expect("Failed to write to README.md");
        }
    }

    // Write the contents of each folder
    for (folder, files) in folder_map {
        writeln!(readme_file, "\n### {}\n", folder).expect("Failed to write to README.md");
        for file_link in files {
            writeln!(readme_file, "{}", file_link).expect("Failed to write to README.md");
        }
    }

    println!("\nREADME.md has been updated with the scanned files.");
}
