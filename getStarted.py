#import requests
#import json
#import sys
#import re   # regex library 
import os.path
#from _datetime import date
#from unittest.mock import _get_target
#from deftables import * #this my code that defines tables and supporting lists
#from utils import *

import sqlite3
print("start Brian")
#os.remove('Chronology.db')
#connection = sqlite3.connect('Chronology.db')
#cursor = connection.cursor()
#cursor.execute('''CREATE TABLE presidents3 (year INTEGER PRIMARY KEY)''')
#cursor.execute('''CREATE TABLE presidents4 (year INTEGER PRIMARY KEY)''')
#connection.commit
#connection.close

connection = sqlite3.connect('Chronology.db')
cursor = connection.cursor()
#cursor.execute('''INSERT INTO presidents2 VALUES (1789)''')
cursor.execute('INSERT INTO presidents (year) VALUES (17900)')
#cursor.execute('''CREATE TABLE presidents7 (year INTEGER PRIMARY KEY)''')
connection.commit
connection.close
print("end brian3")