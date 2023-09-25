# DevAgent CLI

The DevAgent CLI tool. Use the power of DevAgent, from the comfort of your CLI!

## Installation

1. Make sure you have [Go](https://go.dev/dl/) installed on your system.

2. Clone this repository:

```sh
git clone https://github.com/a-alhusaini/devagent-cli.git
cd devagent-cli
go build
```

## Usage

There are two ways to use the CLI:

`./devagent <question>`

`./devagent`

## Example Response

`./devagent write a simple flask app`

Result:

Here's a simple Flask app that you can use to get started with Flask:

```python
from flask import Flask

app = Flask(__name__)

@app.route('/')
def hello():
    return 'Hello, world!'

@app.route('/<string:name>')
def goodbye(name):
    return f'Goodbye, {name}!'

if __name__ == '__main__':
    app.run(port=80)
```

This Flask app defines two routes: one for a root URL and one for a URL with a `name` parameter. The root URL will respond with a simple "Hello, world!" message, and the URL with a `name` parameter will respond with a personalized "Goodbye" message. The app is then run on port 80.
