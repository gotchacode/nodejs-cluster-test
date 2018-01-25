const http = require('http');

// Workers can share any TCP connection
// In this case it is an HTTP server
http.createServer((req, res) => {
  res.writeHead(200);
  res.end('hello world\n');
}).listen(8000);


console.log('server started at http://127.0.0.1:8000/');
