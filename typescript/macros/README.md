### Script Explanation

The script fetches nutritional data for food items using the USDA Food Data Central API, processes the data, and writes it into different file formats for easy access. Below, I'll break down its core components:

1. **Dependencies**: The script uses several libraries:

   - `axios` for making HTTP requests.
   - `fs` for reading/writing to files.
   - `readline` for user input/output interaction in the command line.
   - `dotenv/config` for accessing environment variables, specifically the API key.

2. **Interfaces and Type Definitions**:

   - `Nutrient` and `FoodData` define TypeScript interfaces to shape the expected data structure for food items and their nutrients.

3. **Helper Functions**:

   - `getApiKey()`: Retrieves the API key from environment variables (`.env` file).
   - `fetchInitialData(query)`: Fetches food data for the given query using the USDA API.
   - `readFromJSONFile(filePath)`: Reads and parses a JSON file.
   - `readFromSimpleFile(filePath)`: Reads simple files containing food names.
   - `saveAllFoodsToMarkdown(foods, filePath)`: Saves a list of foods into a Markdown file.
   - `readTemplateFile(filePath)`: Reads the template file's content.
   - `readNutrientNamesFromFile(filePath)`: Reads nutrient names from a file for processing.

4. **Process Flow**:

   - The script interacts with the user through the command line.
   - Users can either select a single food item or save all fetched data to a file.
   - Data fetched from the API or read from a CSV file is processed and saved using helper functions.
   - Nutrient data is formatted and written to a Markdown file or a custom template.

5. **Main Workflow**:
   - The user is prompted to choose between selecting an individual item or saving all items to a file.
   - Based on the choice, the script performs various operations to fetch and store food data.
   - The script also includes asynchronous operations, notably when reading nutrient data and fetching detailed food information.

### README

Here’s a simple README file for this project:

````markdown
# Food Nutrient Data Tool

This script fetches nutritional information for food items using the USDA Food Data Central API and processes it to generate Markdown files for easy reference.

## Features

- Fetch nutritional data for food items using the USDA API.
- Save all food items to a Markdown file.
- Read food items and nutrients from CSV/JSON files.
- Process and save food data using custom templates.

## Prerequisites

- Node.js (v12 or higher)
- `dotenv` for managing environment variables.
- USDA Food Data Central API Key. Register [here](https://fdc.nal.usda.gov/api-key-signup.html).

## Setup

1. Clone the repository:

   ```bash
   git clone <repository_url>
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Set up a `.env` file with your USDA API key:

   ```env
   FOOD_DATA_CENTRAL_API_KEY=YOUR_API_KEY_HERE
   ```

## Usage

Run the script using Node.js:

```bash
node index.js <filePathOrQuery>
```
````

- `<filePathOrQuery>`: Either provide a file path containing food items or a search query string.
- You can interactively choose to save all food items to a file or get detailed information on a single item.

### Example Commands

- Search for food items by query:

  ```bash
  node index.js "apple"
  ```

- Read from a file:

  ```bash
  node index.js foodList.txt
  ```

## File Structure

- **`index.js`**: Main script that interacts with the user.
- **`utils/`**: Directory containing utility functions.
- **`data/`**: Directory for storing template and output files.

## Environment Variables

This project uses environment variables for security. Set the `FOOD_DATA_CENTRAL_API_KEY` in a `.env` file at the project root.

## Dependencies

- **axios**: For making HTTP requests.
- **fs**: For reading and writing files.
- **readline**: For user input.
- **dotenv**: For managing environment variables.

## License

This project is licensed under the MIT License.

## Contribution

Feel free to submit issues or pull requests to improve this project.

### Documentation

Here’s more detailed documentation about each function and how it works:

#### Helper Functions

- **`getApiKey()`**

  - **Purpose**: Fetches the API key from environment variables.
  - **Throws**: Error if the API key is not set.

- **`fetchInitialData(query)`**

  - **Input**: `query` (string) - Food name to search.
  - **Returns**: First food item from the search result or prints an error if none is found.

- **`readFromJSONFile(filePath)`**

  - **Input**: `filePath` (string) - Path to the JSON file.
  - **Returns**: Parsed JSON data as an array.

- **`readFromSimpleFile(filePath)`**

  - **Input**: `filePath` (string) - Path to a simple text file.
  - **Returns**: Array of non-empty lines after reading.

- **`saveAllFoodsToMarkdown(foods, filePath)`**

  - **Input**: `foods` (array of strings), `filePath` (string).
  - **Action**: Writes all foods to a Markdown file with a numbered list.

- **`readNutrientNamesFromFile(filePath)`**
  - **Input**: `filePath` (string) - Path to a text file.
  - **Returns**: Array of nutrient names found in the file.

#### Process Functions

- **`fetchAndSaveData(foodItem, templateFilePath, outputFilePath, nutrientNames)`**
  - **Input**:
    - `foodItem` (string): The food item name to search for.
    - `templateFilePath` (string): Template to use for formatting the data.
    - `outputFilePath` (string): File path to save the formatted data.
    - `nutrientNames` (array): List of nutrients to extract.
  - **Action**: Fetches and saves food information using the provided template.

#### User Interaction

- **Interactive Commands**:
  - The script prompts users to decide whether they want to save all data or focus on a particular item.
  - Depending on the choice, the script will either save the data to files or allow the user to pick an item for more details.

#### Notes

- **API Usage**: You need to have a USDA Food Data Central API Key to run this script. Make sure it is set up correctly in your `.env` file.
- **Data Storage**: Output files are saved in Markdown format for easy readability.
