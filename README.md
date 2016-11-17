# News Clustering

Very experimental project to try and scrape the web and cluster articles into similar events.

At the moment the whole thing doesn't work very well. Seems similarity matrix gives false positives.

## Usage

### 0. Clone repo

```
$ git clone https://github.com/smaxwellstewart/news-clustering
cd news-clustering
```

### 1. Scrape data

```bash
$ cd scrape
$ go run main.go
```

### 2. Compute Similarity Matrix

```bash
$ cd ../similarities
$ python similarities.python
```

### 3. Cluster Results

```bash
$ cd ../cluster
$ go run main.go -threshold 0.9
```

## Improvements

- Add synonyms to find higher similarity between related terms
- Improve clustering:
 - Make time score continuous
 - Add score to text body
