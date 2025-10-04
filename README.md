# GoCrawl: Concurrent Webcrawler and Inverted Indexer 🕷️
GoCrawl is a command-line application written in Go (Golang) designed to learn about concurrency and data structures in the language. It functions as a concurrent webcrawler that builds an inverted index from the content it scrapes.

## Features
- **Concurrent Crawling:** Efficiently crawl websites using Go's concurrency model (goroutines and channels).

- **Inverted Index Generation:** Creates an inverted index from the scraped page content, mapping words to the URLs where they appear.

- **Data Persistence:**

   - **Save:** Export the current inverted index and list of crawled URLs to a JSON file.

   - **Load:** Import a previously saved dataset from a JSON file.

- **CLI Interface:** Simple, interactive command-line interface for easy operation.

## Command Reference
The application uses a simple, interactive command-line interface. Use the `help` command for a full list and description.


| Command | Description | Example Usage |
| :--- | :--- | :--- |
| **`help`** | Displays the full list of available commands and their usage. | `$ help` |
| **`status`** | Shows the current state of the datastore. | `$ status` |
| **`run`** | **Starts the concurrent webcrawler** from the predefined base url. | `$ run` |
| **`load`** | **Loads** a previously saved inverted index and URL list from the JSON file `data.json`. | `$ load` |
| **`save <filepath>`** | **Saves** the current inverted index and URL list to the JSON file `data.json`. | `$ save` |
| **`urls`** | Lists all the URLs that have been successfully crawled and indexed. | `$ urls` |
| **`index`** | Reports the entire inverted index. | `$ index` |
| **`exit`** | Quits the application. | `$ exit` |
