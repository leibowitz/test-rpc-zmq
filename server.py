import zmq
context = zmq.Context()
socket = context.socket(zmq.REP)
#socket.bind("ipc:///tmp/customer")
socket.bind("tcp://0.0.0.0:17563")
 
while True:
    msg = socket.recv()
    print msg
    r = "Got" + msg
    print r
    socket.send(r)

