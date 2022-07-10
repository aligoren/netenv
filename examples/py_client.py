import socket

HOST = "127.0.0.1"  # The server's hostname or IP address
PORT = 8080  # The port used by the server

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    # s.sendall(b"echo\n")
    # s.sendall(b"auth:admin:admin:echo\n")
    # s.sendall(b"auth:admin:admin:file$project1:env$dev:var$name,TEST,MYSQL_USERNAME\n")
    # s.sendall(b"auth:admin:admin:file$project1:env$dev:var$*\n")
    # s.sendall(b"auth:admin:admin:file$project1:env$dev:var$TEST\n")
    # s.sendall(b"auth:admin:admin:file$project2:env$prod:var$name\n")
    s.sendall(b"file$project2:env$prod:var$name\n")
    #s.sendall(b"echo\n")
    data = s.recv(1024)

print(f"Received: {data.decode('utf-8')}")
