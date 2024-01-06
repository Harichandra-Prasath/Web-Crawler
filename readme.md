### A simple web crawler in Go

For now , scraping of urls is working with vanilla html package   
Using Recursive crawler function to identify the ref tags and scraping the url  
No usage of scraping libraries  


### Concurrency

Used Goroutines and Waigroups for concurrency   
Running worker goroutines as same number as the current queue length   

### Benchmarks

All urls of scrapeme site are scraped as scrapeme.live/shop/ as Initial Url   
All 804 urls are accessed around 1.30 minutes with failures as well   