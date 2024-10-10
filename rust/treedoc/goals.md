# **Treedoc: Content Summarizer and File Tracking Tool**

## **Overview**
**Treedoc** is a powerful file summarizer and version tracker that scans directories, parses file content, and generates concise summaries. It uses **Tree-sitter** to perform syntax-aware analysis, allowing targeted searches within specific code blocks, quotes, or any token from a grammar. Designed for flexibility and performance, Treedoc offers features such as incremental backup, restore, change tracking, and customizable search options with tunable fuzziness.

Treedoc’s focus is on content analysis while supporting optional features for version control, file tracking, and undo operations through a change commutation model. Users can seamlessly navigate and roll back changes using a structured **undo tree**.

---

## **Features**

### **1. Content Summarization**
- **Automatic Summaries**: Generate summaries for text-based files (e.g., Markdown, source code) by parsing the file content and capturing key sections, comments, or code blocks.
- **Customizable Summary Length**: Users can choose between brief one-line summaries or detailed multi-paragraph overviews.

   **Example Command**:
   ```bash
   treedoc --summarize --length brief /path/to/file.txt
   ```
   - **Output**: "This file contains a brief overview of project goals, function definitions, and key features."

### **2. Targeted Syntax-Aware Search**
- **Aimed Searches with Tree-sitter**: Perform searches limited to specific syntax elements (code blocks, quotes, comments, or any grammar token). Tree-sitter parses files and directs searches to relevant nodes.
- **Token-Based Search**: Search for specific elements like function names, string literals, or comments.

   **Example Command**:
   ```bash
   treedoc --search "myFunction" --token function --fuzz 0.5
   ```
   - **Output**: Finds "myFunction" only in function definitions, with a moderate fuzziness level for near matches.

- **Tunable Fuzzy Search**: Control the strictness or flexibility of searches with a tunable "fuzzy knob." Lower values yield exact matches, while higher values allow for looser matches.

   **Example Command**:
   ```bash
   treedoc --search "variabl" --fuzz 0.8
   ```
   - **Output**: Locates "variable" even with a typo in the search term.

### **3. Incremental Parsing and Indexing**
- **Incremental Syntax Tree Updates**: Tree-sitter updates only the modified parts of a file’s syntax tree, making searches and indexing highly efficient.
- **Fuzzy Indexing**: Tree-sitter helps maintain a fuzzy index of keywords, code elements, and content for fast, context-aware searches.

   **Example Command**:
   ```bash
   treedoc --index /path/to/project
   ```
   - **Output**: Indexes functions, classes, comments, and other code elements for faster, targeted searches.

### **4. Undo Tree and Change Tracking**
- **Undo Tree Visualization**: Tracks all changes in a structured undo tree, allowing users to review and selectively roll back changes. Supports change commutation for flexible reordering of changes.
- **Search in Undo Tree**: Search for specific changes in the undo tree by content, token, or keyword.

   **Example Command**:
   ```bash
   treedoc --search "oldName" --undo-tree /path/to/file
   ```
   - **Output**: Highlights all instances of "oldName" within the undo tree, showing changes and commutations.

- **Change Commutation**: Allows for independent changes to be applied in any order without changing the final result.

   **Example Command**:
   ```bash
   treedoc --restore --commute /path/to/file --hash abc123
   ```
   - **Output**: Rolls back a specific change while leaving other changes intact.

### **5. Backup and Restore Functionality**
- **Incremental Backups**: Treedoc supports incremental backups, tracking only the changes since the last operation.
- **Restore from Backup**: Restore a file or entire directory to a specific point in time using a timestamp or hash.

   **Example Command**:
   ```bash
   treedoc --restore /path/to/file --hash abc123
   ```

- **Verify Restore Points**: Before restoring, Treedoc verifies that valid restore points exist, preventing errors during the restoration process.

   **Example Command**:
   ```bash
   treedoc --verify-restore /path/to/file
   ```

### **6. Detailed Logging and Reports**
- **Comprehensive Logs**: All operations are logged with detailed reports for easy debugging and analysis.
- **Error Reporting**: Automatically generates detailed error reports, including relevant logs, to help users troubleshoot issues.

   **Example Command**:
   ```bash
   treedoc --log /path/to/file
   ```

---

## **Development Timeline**

### **Phase 1: Core Functionality (Months 1-2)**
1. **File Scanning and Content Summarization**:
   - Implement directory traversal, file scanning, and Tree-sitter integration for content parsing.
   - Generate brief and detailed summaries based on user preferences.
   - Initial CLI setup with commands like `--summarize` and `--scan`.

2. **Targeted Syntax-Aware Search**:
   - Develop token-based search capabilities using Tree-sitter to aim searches at specific code blocks, comments, and other nodes.
   - Add tunable fuzzy search functionality.

### **Phase 2: Database and Backup Integration (Months 3-4)**
1. **Database Setup**:
   - Integrate SQLite for storing syntax trees, timestamps, and metadata.
   - Enable incremental backups for file changes.

2. **Undo Tree and Change Commutation**:
   - Build undo tree support for tracking and visualizing changes.
   - Implement change commutation for flexible undo/redo operations.

3. **Backup and Restore Features**:
   - Add commands to create and verify restore points.
   - Implement rollback capabilities, allowing users to restore files to previous states.

### **Phase 3: Advanced Search and Fuzzy Indexing (Months 5-6)**
1. **Fuzzy Indexing**:
   - Implement fuzzy indexing for faster, context-aware search results based on Tree-sitter’s syntax tree.

2. **Undo Tree Search**:
   - Add search capabilities within the undo tree, allowing users to target specific changes based on content.

3. **Backup and Restore UI**:
   - Develop a TUI or basic GUI for users to visualize and interact with backups, changes, and searches.

### **Phase 4: Final Polishing and Documentation (Months 7-8)**
1. **Comprehensive Documentation**:
   - Write and update the help documentation, man pages, and error reporting system.
   - Create user guides and tutorials for all major features.

2. **Performance Optimizations**:
   - Profile the tool’s performance and optimize indexing, fuzzy search, and backup mechanisms.
   - Add parallel processing to improve the performance of large-scale operations.

---

## **Examples of Usage**

### **1. Generate Summaries for All Files in a Directory**
```bash
treedoc --summarize --recursive /path/to/project
```
- **Output**: Summarizes the contents of all files within `/path/to/project`, generating brief overviews of key sections and comments.

### **2. Search for a Function Name Within Code Blocks**
```bash
treedoc --search "myFunction" --token function
```
- **Output**: Searches only within function definitions for the specified term.

### **3. Rollback a Change Using the Undo Tree**
```bash
treedoc --restore /path/to/file --hash abc123
```
- **Output**: Rolls back to the specified change while keeping independent changes intact through change commutation.

### **4. Search for Changes in the Undo Tree**
```bash
treedoc --search "oldVar" --undo-tree /path/to/file
```
- **Output**: Highlights changes to "oldVar" in the undo tree, showing when and how it was renamed.

---

## **Future Enhancements**
- **Graphical User Interface (GUI)**: Develop a GUI for advanced visualization and easier interaction with the undo tree and backup options.
- **Parallel Processing**: Optimize indexing and summarization for large directories through parallel processing.
- **Non-Text File Support**: Extend parsing and summarization to non-text files.
