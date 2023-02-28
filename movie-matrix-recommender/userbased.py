import pickle
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from sklearn.utils import shuffle
from datetime import datetime
from sortedcontainers import SortedList

import os
if not os.path.exists('bin/movie2user.json') or \
        not os.path.exists('bin/user2movie.json') or \
        not os.path.exists('bin/usermovie2ratings.json') or \
        not os.path.exists('bin/usermovie2ratings_test.json'):
    import c_preprocess_2dict

# with open
