import axios from 'axios';
import * as fs from 'fs';
import * as readline from 'readline';
import 'dotenv/config';

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

interface Nutrient {
  nutrientName: string;
  value: string;
}

interface FoodData {
  itemName: string;
  nutrients: Nutrient[];
}

// API end point as a constant
const FOOD_DATA_API_URL = 'https://api.nal.usda.gov/fdc/v1/foods/search';

const getApiKey = () => {
  const apiKey = process.env.FOOD_DATA_CENTRAL_API_KEY;
  if (!apiKey) {
    console.error('API key is not set.');
    rl.close();
    throw new Error('API key is not set');
  }
  return apiKey;
};

const fetchInitialData = async (
  query: string
): Promise<FoodData | undefined> => {
  try {
    const apiKey = getApiKey();
    const response = await axios.get(
      `${FOOD_DATA_API_URL}?query=${query}&pageSize=10&api_key=${apiKey}`
    );
    const foods: FoodData[] = response.data.foods;
    if (foods.length === 0) {
      console.log(`No foods found for the query: ${query}`);
      rl.close();
      return;
    }
    return foods[0];
  } catch (error) {
    console.error('Error fetching data:', error);
    rl.close();
  }
};

const readFromJSONFile = (filePath: string): any[] => {
  try {
    const fileContents = fs.readFileSync(filePath, 'utf8');
    console.log('JSON File Contents:', fileContents); // Debug log
    const data = JSON.parse(fileContents);
    console.log('Parsed JSON Data:', data); // Debug log
    return data;
  } catch (error) {
    console.error('Error reading JSON file:', error);
    return [];
  }
};

const readFromSimpleFile = (filePath: string): string[] => {
  try {
    const fileContents = fs.readFileSync(filePath, 'utf8');
    return fileContents
      .split('\n')
      .map((line) => line.trim())
      .filter((line) => line);
  } catch (error) {
    console.error('Error reading file:', error);
    return [];
  }
};

const saveAllFoodsToMarkdown = (foods: string[], filePath: string) => {
  let markdownContent = '# Food List\n\n';
  foods.forEach((food, index) => {
    markdownContent += `${index + 1}. ${food}\n`;
  });

  fs.writeFileSync(filePath, markdownContent);
  console.log(`All food items have been written to ${filePath}`);
};

const readTemplateFile = (filePath: string): string => {
  try {
    return fs.readFileSync(filePath, 'utf8');
  } catch (error) {
    console.error(`Error reading template file at path: ${filePath}`, error);
    return '';
  }
};

const readNutrientNamesFromFile = async (
  filePath: string
): Promise<string[]> => {
  try {
    const fileContent = await fs.promises.readFile(filePath, 'utf-8');
    return fileContent.split('\n').filter((line) => line.trim().length > 0);
  } catch (error) {
    console.error(`Error reading file: ${error}`);
    return [];
  }
};

async function fetchAndSaveData(
  foodItem: string,
  templateFilePath: string,
  outputFilePath: string,
  nutrientNames: string[]
) {
  try {
    const foodData = await fetchInitialData(foodItem);

    if (!foodData) {
      console.error(`No data found for ${foodItem}`);
      return;
    }

    const processedData = nutrientNames.map((nutrientName) => {
      return {
        name: nutrientName,
        value: findNutrientValue(foodData, nutrientName),
      };
    });

    const formattedData = formatData(processedData, templateFilePath);
    saveData(formattedData, outputFilePath);
  } catch (error) {
    console.error(`Error fetching data for ${foodItem}:`, error);
  }
}

function formatData(processedData: any, templateFilePath: string): string {
  let template = fs.readFileSync(templateFilePath, 'utf-8');
  // Example: Replace placeholders in the template with actual data
  processedData.forEach((data: any) => {
    template = template.replace(`{{${data.name}}}`, data.value);
  });
  return template;
}

function saveData(data: string, outputFilePath: string): void {
  fs.writeFileSync(outputFilePath, data, 'utf-8');
}

const findNutrientValue = (foodData: any, nutrientName: string): string => {
  if (foodData && Array.isArray(foodData.nutrients)) {
    const nutrient = foodData.nutrients.find(
      (nutrient: any) =>
        nutrient.nutrientName.trim().toLowerCase() ===
        nutrientName.trim().toLowerCase()
    );
    return nutrient ? nutrient.value || '' : '';
  } else {
    console.error(
      `Nutrients not found in food data or food data is invalid: ${JSON.stringify(
        foodData
      )}`
    );
    return '';
  }
};

const processFile = (filePath: string): string[] => {
  const fileExtension = filePath.split('.').pop();
  let foodItems: string[] = [];

  if (fileExtension === 'csv') {
    const fileContent = fs.readFileSync(filePath, 'utf-8');
    const lines = fileContent.split('\n');
    for (const line of lines) {
      const columns = line.split(',');
      const foodItem = columns[0].trim();
      if (foodItem) {
        foodItems.push(foodItem);
      }
    }
  } else {
    console.error('Unsupported file type. Only CSV files are supported.');
    return [];
  }

  return foodItems;
};

rl.question(
  'Do you want to (1) select an item or (2) save all to a file? Enter 1 or 2: ',
  (choice) => {
    if (choice === '1') {
      processFoods(foodItems); // Use foodItems
    } else if (choice === '2') {
      rl.question(
        'Enter the template file path (e.g., ./template.md): ',
        (templateFilePath) => {
          rl.question(
            'Enter the output file path (e.g., ./output.md): ',
            async (outputFilePath) => {
              for (const foodItem of foodItems) {
                // Use foodItems
                const nutrientNames = await readNutrientNamesFromFile(
                  nutrientNamesFilePath
                );
                await fetchAndSaveData(
                  foodItem,
                  templateFilePath,
                  outputFilePath,
                  nutrientNames
                );
              }

              saveAllFoodsToMarkdown(foodItems, outputFilePath); // Use foodItems
              rl.close();
            }
          );
        }
      );
    } else {
      console.log('Invalid choice.');
      rl.close();
    }
  }
);

const processFoods = (foods: string[]) => {
  if (foods.length === 0) {
    console.log('No foods found.');
    rl.close();
    return;
  }

  foods.forEach((food, index) => {
    console.log(`${index + 1}: ${food}`);
  });

  rl.question('Select a food item number for more details: ', (number) => {
    const selectedFood = foods[parseInt(number) - 1];
    if (!selectedFood) {
      console.log('Invalid selection.');
      rl.close();
      return;
    }
    console.log(`You selected: ${selectedFood}`);
    rl.question(
      'Enter the file path and name where you want to save the output (e.g., ./output.md): ',
      (filePath) => {
        saveToFile(selectedFood, filePath, () => {
          rl.close();
        });
      }
    );
  });
};

const saveToFile = (
  selectedFood: string,
  filePath: string,
  callback: () => void
) => {
  try {
    const content = `Selected Food: ${selectedFood}\n`;
    fs.writeFileSync(filePath, content);
    console.log(
      `Details of the selected food item have been saved to ${filePath}`
    );
    callback();
  } catch (error) {
    console.error('Error saving to file:', error);
  }
};

const filePathOrQuery = process.argv[2];
const templateFilePath = 'table.md';
const outputFilePath = 'allitems.md';
const nutrientNamesFilePath = 'nutrients.csv';

let foodItems: string[] = [];

if (filePathOrQuery) {
  if (fs.existsSync(filePathOrQuery)) {
    foodItems = readFromSimpleFile(filePathOrQuery);
    for (const item of foodItems) {
      (async () => {
        const nutrientNames = await readNutrientNamesFromFile(
          nutrientNamesFilePath
        );
        await fetchAndSaveData(
          item,
          templateFilePath,
          outputFilePath,
          nutrientNames
        );
      })();
    }
  } else {
    (async () => {
      const foodData = await fetchInitialData(filePathOrQuery);
      if (foodData) {
        foodItems.push(foodData.itemName);
      }
    })();
  }
} else {
  console.log('No argument provided.');
  rl.close();
}
