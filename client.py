import zmq
context = zmq.Context()
socket = context.socket(zmq.REQ)
#socket.connect("ipc:///tmp/customer")
socket.connect("tcp://0.0.0.0:17563")
 
for i in range(10):
    msg = "msg %s" % i
    socket.send(msg)
    print "Sending", msg
    msg_in = socket.recv()
    print msg_in

