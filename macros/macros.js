"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
var axios_1 = require("axios");
var fs = require("fs");
var readline = require("readline");
require("dotenv/config");
var rl = readline.createInterface({
    input: process.stdin, output: process.stdout,
});
// API end point as a constant
var FOOD_DATA_API_URL = 'https://api.nal.usda.gov/fdc/v1/foods/search';
var getApiKey = function () {
    var apiKey = process.env.FOOD_DATA_CENTRAL_API_KEY;
    if (!apiKey) {
        console.error('API key is not set.');
        rl.close();
        throw new Error('API key is not set');
    }
    return apiKey;
};
var fetchInitialData = function (query) { return __awaiter(void 0, void 0, void 0, function () {
    var apiKey, response, foods, error_1;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                apiKey = getApiKey();
                return [4 /*yield*/, axios_1.default.get("https://api.nal.usda.gov/fdc/v1/foods/search?query=".concat(query, "&pageSize=10&api_key=").concat(apiKey))];
            case 1:
                response = _a.sent();
                foods = response.data.foods;
                if (foods.length === 0) {
                    console.log("No foods found for the query: ".concat(query));
                    rl.close();
                    return [2 /*return*/];
                }
                return [2 /*return*/, foods[0]];
            case 2:
                error_1 = _a.sent();
                console.error('Error fetching data:', error_1);
                rl.close();
                return [3 /*break*/, 3];
            case 3: return [2 /*return*/];
        }
    });
}); };
var readFromJSONFile = function (filePath) {
    try {
        var fileContents = fs.readFileSync(filePath, 'utf8');
        console.log("JSON File Contents:", fileContents); // Debug log
        var data = JSON.parse(fileContents);
        console.log("Parsed JSON Data:", data); // Debug log
        return data;
    }
    catch (error) {
        console.error('Error reading JSON file:', error);
        return [];
    }
};
var readFromSimpleFile = function (filePath) {
    try {
        var fileContents = fs.readFileSync(filePath, 'utf8');
        return fileContents.split('\n').map(function (line) { return line.trim(); }).filter(function (line) { return line; });
    }
    catch (error) {
        console.error('Error reading file:', error);
        return [];
    }
};
var saveAllFoodsToMarkdown = function (foods, filePath) {
    var markdownContent = '# Food List\n\n';
    foods.forEach(function (food, index) {
        markdownContent += "".concat(index + 1, ". ").concat(food, "\n");
    });
    fs.writeFileSync(filePath, markdownContent);
    console.log("All food items have been written to ".concat(filePath));
};
var readTemplateFile = function (filePath) {
    try {
        return fs.readFileSync(filePath, 'utf8');
    }
    catch (error) {
        console.error("Error reading template file at path: ".concat(filePath), error);
        return '';
    }
};
var readNutrientNamesFromFile = function (filePath) { return __awaiter(void 0, void 0, void 0, function () {
    var fileContent, error_2;
    return __generator(this, function (_a) {
        switch (_a.label) {
            case 0:
                _a.trys.push([0, 2, , 3]);
                return [4 /*yield*/, fs.promises.readFile(filePath, 'utf-8')];
            case 1:
                fileContent = _a.sent();
                return [2 /*return*/, fileContent.split('\n').filter(function (line) { return line.trim().length > 0; })];
            case 2:
                error_2 = _a.sent();
                console.error("Error reading file: ".concat(error_2));
                return [2 /*return*/, []];
            case 3: return [2 /*return*/];
        }
    });
}); };
function fetchAndSaveData(foodItem, templateFilePath, outputFilePath, nutrientNames) {
    return __awaiter(this, void 0, void 0, function () {
        var foodData_1, processedData, formattedData, error_3;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0:
                    _a.trys.push([0, 2, , 3]);
                    return [4 /*yield*/, fetchData(foodItem)];
                case 1:
                    foodData_1 = _a.sent();
                    processedData = nutrientNames.map(function (nutrientName) {
                        return {
                            name: nutrientName, value: findNutrientValue(foodData_1, nutrientName)
                        };
                    });
                    formattedData = formatData(processedData, templateFilePath);
                    saveData(formattedData, outputFilePath);
                    return [3 /*break*/, 3];
                case 2:
                    error_3 = _a.sent();
                    console.error("Error fetching data for ".concat(foodItem, ":"), error_3);
                    return [3 /*break*/, 3];
                case 3: return [2 /*return*/];
            }
        });
    });
}
// Helper functions for fetching, formatting, and saving data (placeholders)
function fetchData(foodItem) {
    return __awaiter(this, void 0, void 0, function () {
        return __generator(this, function (_a) {
            return [2 /*return*/];
        });
    });
}
function formatData(processedData, templateFilePath) {
    // Example: Load the template, and replace placeholders with actual data
    var template = fs.readFileSync(templateFilePath, 'utf-8');
    // Ensure that the function returns a string
    return formattedData;
}
function saveData(data, outputFilePath) {
    // Write the formatted data to the specified file
    fs.writeFileSync(outputFilePath, data, 'utf-8');
}
// Add this function to find the nutrient value by nutrient name
var findNutrientValue = function (foodData, nutrientName) {
    if (foodData && Array.isArray(foodData.nutrients)) {
        var nutrient = foodData.nutrients.find(function (nutrient) { return nutrient.nutrientName.trim().toLowerCase() === nutrientName.trim().toLowerCase(); });
        return nutrient ? nutrient.value || '' : '';
    }
    else {
        console.error("Nutrients not found in food data or food data is invalid: ".concat(JSON.stringify(foodData)));
        return '';
    }
};
var processFile = function (filePath) {
    var fileExtension = filePath.split('.').pop();
    var foodItems = [];
    if (fileExtension === 'csv') {
        // Read and parse the CSV file
        var fileContent = fs.readFileSync(filePath, 'utf-8');
        var lines = fileContent.split('\n');
        for (var _i = 0, lines_1 = lines; _i < lines_1.length; _i++) {
            var line = lines_1[_i];
            var columns = line.split(','); // Simple CSV parsing, adjust as needed
            // Assuming the food item is in the first column
            var foodItem = columns[0].trim();
            if (foodItem) {
                foodItems.push(foodItem);
            }
        }
    }
    else {
        console.error('Unsupported file type. Only CSV files are supported.');
        return [];
    }
    return foodItems;
};
rl.question('Do you want to (1) select an item or (2) save all to a file? Enter 1 or 2: ', function (choice) {
    if (choice === '1') {
        processFoods(foodItems); // Use foodItems
    }
    else if (choice === '2') {
        rl.question('Enter the template file path (e.g., ./template.md): ', function (templateFilePath) {
            rl.question('Enter the output file path (e.g., ./output.md): ', function (outputFilePath) { return __awaiter(void 0, void 0, void 0, function () {
                var _i, foodItems_2, foodItem, nutrientNames_1;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0:
                            _i = 0, foodItems_2 = foodItems;
                            _a.label = 1;
                        case 1:
                            if (!(_i < foodItems_2.length)) return [3 /*break*/, 4];
                            foodItem = foodItems_2[_i];
                            return [4 /*yield*/, readNutrientNamesFromFile(nutrientNamesFilePath)];
                        case 2:
                            nutrientNames_1 = _a.sent();
                            fetchAndSaveData(foodItem, templateFilePath, outputFilePath, nutrientNames_1);
                            _a.label = 3;
                        case 3:
                            _i++;
                            return [3 /*break*/, 1];
                        case 4:
                            // Call saveAllFoodsToMarkdown here
                            saveAllFoodsToMarkdown(foodItems, outputFilePath); // Use foodItems
                            rl.close();
                            return [2 /*return*/];
                    }
                });
            }); });
        });
    }
    else {
        console.log('Invalid choice.');
        rl.close();
    }
});
var saveAllFoodsToMarkdown = function (foods, filePath) {
    var markdownContent = '# Food List\n\n';
    foods.forEach(function (food, index) {
        markdownContent += "".concat(index + 1, ". ").concat(food, "\n");
    });
    fs.writeFileSync(filePath, markdownContent);
    console.log("All food items have been written to ".concat(filePath));
};
var processFoods = function (foods) {
    if (foods.length === 0) {
        console.log('No foods found.');
        rl.close();
        return;
    }
    foods.forEach(function (food, index) {
        console.log("".concat(index + 1, ": ").concat(food)); // Changed from food.description to food
    });
    rl.question('Select a food item number for more details: ', function (number) {
        var selectedFood = foods[parseInt(number) - 1];
        if (!selectedFood) {
            console.log('Invalid selection.');
            rl.close();
            return;
        }
        console.log("You selected: ".concat(selectedFood)); // Changed from selectedFood.description to selectedFood
        rl.question('Enter the file path and name where you want to save the output (e.g., ./output.md): ', function (filePath) {
            saveToFile(selectedFood, filePath, function () {
                rl.close(); // Close readline interface after saving the file
            });
        });
    });
};
var saveToFile = function (selectedFood, filePath, callback) {
    try {
        // The content to be saved. This can be customized as needed.
        var content = "Selected Food: ".concat(selectedFood, "\n");
        // Writing content to the file
        fs.writeFileSync(filePath, content);
        console.log("Details of the selected food item have been saved to ".concat(filePath));
    }
    catch (error) {
        console.error('Error saving to file:', error);
    }
};
var filePathOrQuery = process.argv[2];
var templateFilePath = 'table.md'; // Set the path to your template file
var outputFilePath = 'allitems.md'; // Set the path to the output file
var nutrientNamesFilePath = 'nutrients.csv'; // Set the path to your nutrient names file
var nutrientNames = readNutrientNamesFromFile(nutrientNamesFilePath);
var foodItems = processFile(filePathOrQuery);
if (filePathOrQuery) {
    if (fs.existsSync(filePathOrQuery)) {
        var items = readFromSimpleFile(filePathOrQuery);
        var _loop_1 = function (item) {
            (function () { return __awaiter(void 0, void 0, void 0, function () {
                var nutrientNames;
                return __generator(this, function (_a) {
                    switch (_a.label) {
                        case 0: return [4 /*yield*/, readNutrientNamesFromFile(nutrientNamesFilePath)];
                        case 1:
                            nutrientNames = _a.sent();
                            fetchAndSaveData(item, templateFilePath, outputFilePath, nutrientNames);
                            return [2 /*return*/];
                    }
                });
            }); })(); // Correctly closing the async IIFE
        };
        for (var _i = 0, items_1 = items; _i < items_1.length; _i++) {
            var item = items_1[_i];
            _loop_1(item);
        }
    }
    else {
        // Assuming it's a query
        fetchInitialData(filePathOrQuery);
    }
}
else {
    console.log('No argument provided.');
    rl.close();
}
var _loop_2 = function (item) {
    (function () { return __awaiter(void 0, void 0, void 0, function () {
        var nutrientNames;
        return __generator(this, function (_a) {
            switch (_a.label) {
                case 0: return [4 /*yield*/, readNutrientNamesFromFile(nutrientNamesFilePath)];
                case 1:
                    nutrientNames = _a.sent();
                    fetchAndSaveData(item, templateFilePath, outputFilePath, nutrientNames);
                    return [2 /*return*/];
            }
        });
    }); })();
};
for (var _a = 0, foodItems_1 = foodItems; _a < foodItems_1.length; _a++) {
    var item = foodItems_1[_a];
    _loop_2(item);
}
