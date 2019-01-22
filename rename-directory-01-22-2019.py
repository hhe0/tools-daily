import os

path = 'D:\CODE\leetcode'
for file in os.listdir(path):
    if str(file).isdigit():
        if len(file) == 5:
            os.rename(file, file[1:])
