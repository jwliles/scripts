use std::fs;
use std::path::Path;

pub fn create_or_update_readme(directory: &str, content: &str) -> Result<(), std::io::Error> {
    let readme_path = Path::new(directory).join("README.md");
    fs::write(readme_path, content)?;
    println!("Created or updated README.md in {}", directory);
    Ok(())
}
