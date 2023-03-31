from flask import Flask
import sys

app = Flask(__name__)

server_name = sys.argv[1]

@app.route("/")
def index():
    return server_name

if  __name__ == '__main__':
    app.run(port=sys.argv[2])