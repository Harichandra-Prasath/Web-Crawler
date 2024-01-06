### A simple web crawler in Go

For now , scraping of urls is working with vanilla html package   
Using Recursive crawler function to identify the ref tags and scraping the url  
No usage of scraping libraries  

### Concurrency

Used Goroutines and Waigroups for concurrency
Ran worker goroutines as same number as the current queue length 