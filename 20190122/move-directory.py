import os
import shutil

path = 'D:\CODE\leetcode'
for file in os.listdir(path):
    if str(file).isdigit():
        dis = path + '\\' + file + '\\' + 'python'
        if os.path.exists(dis):
            if not os.path.exists(dis + '\\' + 'Solution.py') and os.path.exists(path + '\\' + file + '\\' + 'Solution.py'):
                shutil.move(path + '\\' + file + '\\' + 'Solution.py', dis + '\\' + 'Solution.py')
        else:
            os.mkdir(dis)
