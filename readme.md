## WEB-CRAWLER
Web-Crawler built in go using standard libraries  
**Shipped as cli using cobra**  

### Set-up 
- Clone this repository  
- Install the necessary packages using go.mod  
- Run it
```bash
    go run main.go [args] [flags]
```
- Optionally Install it   
```bash
    go install
    Web-Crawler [args] [flags]
```
- Make sure the go/bin is in your path  

### Usage

#### Available Commands
- crawl (Main Subcommand of Crawler... Entrypoint)  
- help  
#### Available Flags
- --version 
- --help  

```bash
Web-Crawler --help
```
#### Available flags for SubCommand "crawl"
- --depth (Defines the depth level of crawling including root) (Needed) (int)  
- --help
- --same-domain (Crawling and scraping from same domains) (Optional) (bool)  
- --url  (Root url to start the crawling) (Needed) (string)  
- --generate (Generate a .txt file with all the links crawled) (Optional) (bool)  

**Example**  
```bash
Web-Crawler crawl --url=https://transform.tools/ --depth 3 --generate --same-domain 
```


