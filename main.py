# import numpy as np
import  pandas as pd
import difflib
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import cosine_similarity

movies_data = pd.read_csv('./data/movies.csv')
#print(movies_data)

# relevant columns for my ml model
selected_features = ['genres', 'keywords', 'tagline', 'cast', 'director']

# data treatment for error prevention: replace null by '' (empty strings)
for feature in selected_features:
  movies_data[feature] = movies_data[feature].fillna('')

# combine relevant columns for analysis
combined_features = movies_data['genres']+' '+movies_data['keywords']+' '+movies_data[ 'tagline']+' '+ movies_data[ 'cast']+' '+ movies_data[ 'director']

# transform my column into values into a vector (pretty complex stuff)
vectorizer = TfidfVectorizer()
feature_vectors = vectorizer.fit_transform(combined_features)

# getting similarity scores
similarity = cosine_similarity(feature_vectors)

# getting the movie name from the user
movie_name = input('MOVIE NAME BROH: ')
# movie_name = 'iron man'

# creating a list with all the movie names given in the dataset
all_items_titles = movies_data['title'].tolist()

# finding the close match for the movie name given by the user
find_close_match = difflib.get_close_matches(movie_name, all_items_titles)
close_match = find_close_match[0]

# finding movie's index with title
close_match_index = movies_data[movies_data.title == close_match]['index'].values[0]

# getting a list of similar movies
similarity_score = list(enumerate(similarity[close_match_index]))

# sorting the movie based on similarity score
sorted_similar_movies = sorted(similarity_score, key = lambda x:x[1], reverse = True)

# print similar movie names on index
count = 1
print('MOVIE RECOMMENDATIONS:')
for movie in sorted_similar_movies:
  index = movie[0]
  title_from_index = movies_data[movies_data.index == index]['title'].values[0]
  if(count<=30):
    print(count, title_from_index)
    count += 1
