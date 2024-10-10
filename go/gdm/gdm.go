package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"
)

const fileContent = "Hey Man, Nice Shot!"

var fileExtensions = []string{
	"md", "nomedia", "png", "jpg", "json", "yml", "css", "pdf",
	"js", "gif", "el", "txt", "py", "sh", "mp3", "ttf",
	"webm", "org", "xml", "svg", "toml", "ini", "zsh",
	"license", "iml", "gitignore", "gitattributes",
}

// Weights for each extension based on their occurrence in the actual dataset
var extensionWeights = map[string]int{
	"md": 50, "nomedia": 20, "png": 10, "jpg": 5,
	"json": 4, "yml": 3, "css": 2, "pdf": 2,
	"js": 2, "gif": 1, "el": 1, "txt": 1, "py": 1,
	"sh": 1, "mp3": 1, "ttf": 1, "webm": 1,
	"org": 1, "xml": 1, "svg": 1, "toml": 1,
	"ini": 1, "zsh": 1, "license": 1, "iml": 1,
	"gitignore": 1, "gitattributes": 1,
}

// Function to pick an extension based on weights
func pickRandomExtension() string {
	totalWeight := 0
	for _, weight := range extensionWeights {
		totalWeight += weight
	}

	randomWeight := rand.Intn(totalWeight)
	cumulativeWeight := 0

	for ext, weight := range extensionWeights {
		cumulativeWeight += weight
		if randomWeight < cumulativeWeight {
			return ext
		}
	}
	return "txt" // Default fallback
}

// Predefined complexity levels and their margins
var levels = map[int][3]int{
	1: {100, 500, 10},    // Low complexity: 100 folders, 500 files, 10% margin
	2: {1000, 5000, 10},  // Medium complexity: 1000 folders, 5000 files, 10% margin
	3: {2000, 10000, 10}, // High complexity: 2000 folders, 10000 files, 10% margin
	4: {5000, 200000, 2}, // Extreme complexity: 5000 folders, 200000 files, 2% margin
}

// Ensure every folder created has at least one file
// createRandomFiles creates numFiles files in basePath and returns their total size
func createRandomFiles(numFiles int, basePath string) int64 {
	var totalSize int64 = 0

	if numFiles == 0 {
		numFiles = 1
	}

	for i := 0; i < numFiles; i++ {
		// Use weighted extension selection
		ext := pickRandomExtension()
		fileName := fmt.Sprintf("file_%d.%s", i, ext)
		filePath := filepath.Join(basePath, fileName)

		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Could not create file:", err)
			continue
		}

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(fileContent)
		if err != nil {
			fmt.Println("Could not write to file:", err)
			writer.Flush()
			file.Close()
			continue
		}

		err = writer.Flush()
		if err != nil {
			fmt.Println("Could not flush writer:", err)
			file.Close()
			continue
		}

		err = file.Close()
		if err != nil {
			fmt.Println("Could not close file:", err)
			continue
		}

		totalSize += int64(len(fileContent)) + int64(len(fileName))
	}
	return totalSize
}

// createStructure creates the directory structure within the specified limits
func createStructure(basePath string, maxFolders int, maxFiles int) (int64, int64, int64) {
	var totalSize int64 = 0
	var totalFiles int64 = 0
	var totalFolders int64 = 0

	var wg sync.WaitGroup

	// Generate directories and files in parallel
	for dirIndex := 0; dirIndex < maxFolders && atomic.LoadInt64(&totalFiles) < int64(maxFiles); dirIndex++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			dirName := fmt.Sprintf("main_dir_%d", index)
			dirPath := filepath.Join(basePath, dirName)

			err := os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				fmt.Println("Could not create directory:", err)
				return
			}
			atomic.AddInt64(&totalFolders, 1)

			numFiles := mini(rand.Intn(10)+1, maxFiles-int(totalFiles))
			dirSize := createRandomFiles(numFiles, dirPath)

			atomic.AddInt64(&totalSize, dirSize)
			atomic.AddInt64(&totalFiles, int64(numFiles))

			numSubdirs := rand.Intn(5) + 1
			for subIndex := 0; subIndex < numSubdirs && atomic.LoadInt64(&totalFolders) < int64(maxFolders) && atomic.LoadInt64(&totalFiles) < int64(maxFiles); subIndex++ {
				subDirName := fmt.Sprintf("subdir_%d_%d", index, subIndex)
				subDirPath := filepath.Join(dirPath, subDirName)

				err := os.MkdirAll(subDirPath, os.ModePerm)
				if err != nil {
					fmt.Println("Could not create subdirectory:", err)
					continue
				}
				atomic.AddInt64(&totalFolders, 1)

				subNumFiles := mini(rand.Intn(10)+1, maxFiles-int(totalFiles))
				subSize := createRandomFiles(subNumFiles, subDirPath)

				atomic.AddInt64(&totalSize, subSize)
				atomic.AddInt64(&totalFiles, int64(subNumFiles))
			}
		}(dirIndex)
	}

	wg.Wait()
	return totalFolders, totalFiles, totalSize
}

// mini returns the minimum of two integers
func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func promptSizeOption() (int, int, int) {
	for {
		fmt.Println("Choose a directory complexity level:")
		fmt.Println("  1. Low (100 folders, 500 files)")
		fmt.Println("  2. Medium (1000 folders, 5000 files)")
		fmt.Println("  3. High (2000 folders, 10000 files)")
		fmt.Println("  4. Extreme (5000 folders, 200000 files)")
		fmt.Println("  5. Custom (Enter your own values)")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if choice >= 1 && choice <= 4 {
			return levels[choice][0], levels[choice][1], levels[choice][2]
		} else if choice == 5 {
			var maxFolders, maxFiles int
			fmt.Println("Enter the number of folders (max 5000):")
			_, err = fmt.Scanln(&maxFolders)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			fmt.Println("Enter the number of files (max 200000):")
			_, err = fmt.Scanln(&maxFiles)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			return mini(maxFolders, 5000), mini(maxFiles, 200000), 10 // Set default margin for custom to 10%
		} else {
			fmt.Println("Invalid choice! Please choose a valid option.")
		}
	}
}

// Function to check and fill empty directories
func fillEmptyDirectories(basePath string) {
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				fmt.Println("Error reading directory:", err)
				return nil
			}

			if len(entries) == 0 {
				fileName := filepath.Join(path, "placeholder.txt")
				file, err := os.Create(fileName)
				if err != nil {
					fmt.Println("Error creating placeholder file:", err)
					return nil
				}
				defer file.Close()
				_, err = file.WriteString("This is a placeholder to ensure the folder is not empty.")
				if err != nil {
					fmt.Println("Error writing to placeholder file:", err)
					return nil
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking through directories:", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<target_directory>")
		os.Exit(1)
	}

	// Initialize the random number generator
	rand.Seed(time.Now().UnixNano())

	// Get the target directory from arguments
	basePath := os.Args[1]

	// Prompt user for folder and file limits and margin
	maxFolders, maxFiles, margin := promptSizeOption()

	// Calculate margins
	folderLimit := int(float64(maxFolders) * (1 + float64(margin)/100))
	fileLimit := int(float64(maxFiles) * (1 + float64(margin)/100))

	// Start timing the generation process
	startTime := time.Now()

	// Create the directory structure until the specified limits are reached
	totalFolders, totalFiles, totalSize := createStructure(basePath, folderLimit, fileLimit)

	// Ensure no directory is empty
	fillEmptyDirectories(basePath)

	// Calculate metrics
	duration := time.Since(startTime).Seconds()
	ops := float64(totalFolders+totalFiles) / duration

	// Print metrics
	fmt.Println("--- Generation Metrics ---")
	fmt.Printf("Directory structure generated in: %s\n", basePath)
	fmt.Printf("Folders Created: %d (Max Limit: %d, Margin: ±%d%%)\n", totalFolders, maxFolders, margin)
	fmt.Printf("Files Created: %d (Max Limit: %d, Margin: ±%d%%)\n", totalFiles, maxFiles, margin)
	fmt.Printf("Total Size of Data Generated: %d bytes\n", totalSize)
	fmt.Printf("Total Time Taken: %.3f seconds\n", duration)
	fmt.Printf("Operations per Second (OPS): %.3f\n", ops)
	fmt.Println("--------------------------")
}
