import json
from sklearn.feature_extraction.text import TfidfVectorizer
from pprint import pprint

# open up scrape
with open('../sample.json') as data_file:
    data = json.load(data_file)

# construct array of
titles = []
for document in data:
    titles.append(document['title'])

# construct similarities matrix
vect = TfidfVectorizer(min_df=1)
tfidf = vect.fit_transform(titles)
similarities = (tfidf * tfidf.T).A

# dump matrix to json
with open('../similarities.json', 'w') as outfile:
    json.dump(similarities.tolist(), outfile)
