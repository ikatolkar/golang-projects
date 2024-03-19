# go-text-search-engine

Uses reverse index to find a text query in a large data dump.
All words are stored in index, and queried, after having been analysed
In analysis, the core optimization logic depends on stopword filtering, and STEMMING
for stemming we're using a snowball stemmer

The logic for searching and building the index, especially intersection, depends on the document indices being in a sorted order

## Data dump
1. The data dump is from wiki media https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz
2. The data is in format:
```xml
<title>all cat types</title>
<url>http://wikimedia.com/cat-types</url>
<abstract>The wildcat is a species complex comprising two small wild cat species: the European wildcat (Felis silvestris) and the African wildcat (F. lybica)</abstract>
```

## Logic
1. Load all data from xml
2. Analyze loaded data, build index
    1. Tokenize             : divide string into tokens, delimiter is any character other than letter and number
    2. Lowercase Filter     : convert all characters of tokens into lowercase
    3. Stopword Filter      : skip common stop words like "a", "an", "the"
    4. Stemmer Filter       : reduce the 
3. Search for query in index
    0. Perform the same analysis on the query
    1. Build a list of document indices, this is the return value
    2. For each token in query
        0. Build a list of document entries for the token
        1. if token exists in index
            1. Find all document entries for the token
            2. Find intersection of (final document list, current token document list)
        2. else token doesn't exist in index, return nil
            1. this will return nil if any of the tokens is not found in the list
    3. Return list of document indices
4. For each document in list of documents print the document

## Snowball Stemmer
https://www.geeksforgeeks.org/snowball-stemmer-nlp/


```bash
$ go run main.go
2024/03/19 04:56:15 Full text search is in progress...
2024/03/19 04:56:58 Loaded 677241 documents in 43.654688896s
2024/03/19 04:57:28 Indexed 677241 documents in 29.647654459s
2024/03/19 04:57:28 Search found 1 documents in 131.616µs
2024/03/19 04:57:28 130267      The wildcat is a species complex comprising two small wild cat species: the European wildcat (Felis silvestris) and the African wildcat (F. lybica).

```
