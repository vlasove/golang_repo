import shutil

import requests

url = 'https://cdn-images-1.medium.com/max/2000/1*30aoNxlSnaYrLhBT0O1lzw.png'
response = requests.get(url, stream=True)
with open('GoLogo.png', 'wb') as out_file:
    shutil.copyfileobj(response.raw, out_file)
del response