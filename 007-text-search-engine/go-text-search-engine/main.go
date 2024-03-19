package main

import (
    "time"
    "flag"
    "log"
    utils "github.com/ikatolkar/go-text-search-engine/utils"
)

func main() {
    // take command line arguments using flag package
    var dumpPath, query string
    flag.StringVar(&dumpPath, "p", "enwiki-latest-abstract.xml.gz", "wiki abstract dump path")
    flag.StringVar(&query, "q" , "Small wild cat", "search query")
    flag.Parse()

    // start process here, keep track of time between different steps
    log.Println("Full text search is in progress...")
    start := time.Now()

    // load documents
    docs, err := utils.LoadDocuments(dumpPath)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
    
    // build index
    start = time.Now()
    idx := make(utils.Index)
    idx.Add(docs)
    log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

    // search for query in index, get all the matched indices
    start = time.Now()
    matchedIDs := idx.Search(query)
    log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

    // for each index, print document
    for _, id := range matchedIDs {
        doc := docs[id]
        log.Printf("%d\t%s\n", id, doc.Text)
    }
}
