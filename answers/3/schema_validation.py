from collections import OrderedDict
from string import Template
import json, os, datetime
from jsonschema import validate

def validate_json(data):
    given_json = json.dumps(data["json-data"])
    schema = { 
        "type": "object", 
        "required": [], 
        "properties": { "colors": { "type": "array", 
        "items": { "type": "object", 
                   "required": [], 
                   "properties": { "color": { "type": "string" }, 
                   "category": { "type": "string" }, 
                   "type": { "type": "string" }, 
                   "code": { "type": "object", 
                   "required": [], 
                   "properties": { "rgba": { "type": "array", 
                   "items": { "type": "number" } }, 
                   "hex": { "type": "string" }
                   }}}}}
                }
            }
    try:
        validate(given_json,schema)
    except Exception as e:
        toReturn = '<div class="alert alert-danger"> <strong>Error!</strong>'+str(e)+' </div>'
        return toReturn
    return '<div class="alert alert-success"> <strong>Success!</strong> Everything looks good </div>'


def application(environ, start_response):
    try:
        request_body_size = int(environ.get('CONTENT_LENGTH', 0))
    except (ValueError):
        request_body_size = 0

    body = environ['wsgi.input'].read(request_body_size)
    print body
    data = json.loads(body)
    print data["json-data"]
    response_headers = [('Content-type', 'text/html')]
    
    start_response('200 OK', response_headers)
    
    request = str(data["request"])
    
    if request == "validate_json":
        return validate_json(data)

    if request == "ping":
        return "pong"

