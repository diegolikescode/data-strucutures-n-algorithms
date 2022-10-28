import pandas as pd
from scipy.sparse import csr_matrix
from sklearn.neighbors import NearestNeighbors
# import numpy as np

movies_df = pd.read_csv('./data/movies.csv', usecols=['movieId', 'title'], dtype={'movieId': 'int32', 'title': 'str'})

ratings_df = pd.read_csv('./data/ratings.csv',
  usecols=['userId', 'movieId', 'rating'],
  dtype={'userId': 'int32', 'movieId':'int32', 'rating': 'float32'})

df = pd.merge(ratings_df, movies_df, on='movieId')

combine_movie_rating = df.dropna(axis = 0, subset = ['title'])
movie_rating_count = (combine_movie_rating.
  groupby(by = ['title'])['rating'].
  count().
  reset_index().
  rename(columns = {'rating': 'totalRatingCount'})[['title', 'totalRatingCount']]
)

rating_with_total_rating_count = combine_movie_rating.merge(movie_rating_count,
  left_on = 'title', right_on = 'title', how = 'left')

pd.set_option('display.float_format', lambda x: '%.3f' % x)
# print(movie_rating_count['totalRatingCount'].describe())

popularity_threshold = 50
rating_popular_movies = rating_with_total_rating_count.query('totalRatingCount >= @popularity_threshold')
movie_features_df = rating_popular_movies.pivot_table(index='title', columns ='userId', values='rating').fillna(0)
# print(movie_features_df.head())

movie_features_df_matrix = csr_matrix(movie_features_df.values)

model_knn = NearestNeighbors(metric='cosine', algorithm='brute')
model_knn.fit(movie_features_df_matrix)

