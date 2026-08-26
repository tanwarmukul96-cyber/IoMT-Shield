from http.server import BaseHTTPRequestHandler, HTTPServer

class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        length = int(self.headers.get("Content-Length", 0))
        data = self.rfile.read(length)

        print(f"[MEDICAL SERVICE] Received: {data.decode()}")

        self.send_response(200)
        self.end_headers()
        self.wfile.write(b"Medical data accepted")

server = HTTPServer(("0.0.0.0", 9000), Handler)

print("Simulated Medical Service listening on :9000")
server.serve_forever()