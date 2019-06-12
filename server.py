from http.server import BaseHTTPRequestHandler
from routes.main import routes

class Server(BaseHTTPRequestHandler):
  def do_HEAD(self):
    return
    
  def do_GET(self):
    self.respond()
    
  def do_POST(self):
    return
    
  def handle_http(self, status, content_type):
    self.send_response(status)
    self.send_header("Content-type", content_type)
    self.end_headers()
    route_content = routes[self.path]
    return bytes(“Hello World”, “UTF-8”)
    
  def respond(self):
    content = self.handle_http(200, "text/html")
    self.wfile.write(content)