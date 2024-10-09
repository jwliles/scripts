### **Treedoc Command Flags Tree**

```
treedoc
│
├── --summarize (Generate content summaries)
│   ├── --scope
│   │   ├── file (Generate a summary for each file)
│   │   ├── folder (Generate a summary for each folder)
│   │   ├── overview (Generate a comprehensive summary for the directory)
│   ├── --format
│   │   ├── README (Output as README.md)
│   │   ├── man (Output as a man page)
│   │   ├── txt (Output as a plain text file)
│   ├── --exclude [file pattern] (Exclude specific file types, e.g., *.log)
│   ├── --length
│   │   ├── brief (Generate a brief summary)
│   │   ├── detailed (Generate a detailed summary)
│
├── --search (Search file content)
│   ├── --token
│   │   ├── function (Search within function definitions)
│   │   ├── code_block (Search within code blocks)
│   │   ├── string (Search within string literals)
│   │   ├── comment (Search within comments)
│   ├── --fuzz [0.0 - 1.0] (Set fuzzy search precision level)
│   ├── --undo-tree (Search within the undo tree)
│   ├── --context [token] (Search within a specific token or context)
│
├── --restore (Restore a file or folder)
│   ├── --hash [hash] (Restore to a specific hash)
│   ├── --incremental (Restore incrementally)
│   ├── --verify-restore (Verify the existence of restore points)
│   ├── --commute (Restore with change commutation)
│   ├── --preview (Preview changes before restoring)
│
├── --index (Index files for faster searches)
│   └── --fuzz (Use fuzzy logic for indexing content)
│
├── --log (View logs)
│   ├── --view [file] (View the log for a specific file)
│   ├── --error (Generate error reports)
│
└── --help (Display help documentation)
    ├── --man (View help in man page format)
    ├── --examples (Show usage examples)
    └── --commands (List all available commands)
```

---

### **Key Features and Sub-Options**

1. **`--summarize`**:
   - **Scope**: Summarize by file, folder, or a comprehensive directory overview.
   - **Format**: Choose output formats like README, man pages, or plaintext.
   - **Exclude**: Option to exclude specific file types from the summary.
   - **Length**: Choose between brief or detailed summaries.

2. **`--search`**:
   - **Token**: Limit searches to specific grammar tokens like functions, code blocks, comments, or strings.
   - **Fuzz**: Control the strictness of the fuzzy search (e.g., `--fuzz 0.5` for moderate fuzziness).
   - **Undo Tree Search**: Perform searches within the undo tree using specific keywords or tokens.

3. **`--restore`**:
   - **Hash**: Restore files to a specific version based on a hash.
   - **Incremental**: Restore incrementally, rolling back changes step by step.
   - **Verify**: Check if valid restore points exist before restoring.
   - **Commutation**: Apply or reverse changes using change commutation logic.
   - **Preview**: Preview the result of the restoration before applying it.

4. **`--index`**:
   - Index files to make searches faster and more efficient. Optionally, use fuzzy logic for approximate matches.

5. **`--log`**:
   - Access log files for debugging or error tracking, with options to view specific logs or generate error reports.

6. **`--help`**:
   - Access help documentation, usage examples, and command listings.

---

### **Examples of Usage**:

1. **Generate a summary for each file in README format, excluding `.log` files**:
   ```bash
   treedoc --summarize --scope file --format README --exclude "*.log" /path/to/project
   ```

2. **Search for a function name with moderate fuzziness**:
   ```bash
   treedoc --search "myFunction" --token function --fuzz 0.5
   ```

3. **Restore a file incrementally and preview the changes before applying**:
   ```bash
   treedoc --restore /path/to/file --incremental --preview
   ```

4. **Index files with fuzzy logic for faster future searches**:
   ```bash
   treedoc --index /path/to/project --fuzz
   ```

5. **View error logs for a specific file**:
   ```bash
   treedoc --log --view /path/to/file --error
   ```
