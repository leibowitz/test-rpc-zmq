<?php
$context = new ZMQContext();

//  Socket to talk to server
echo "Connecting to hello world server…\n";
$requester = new ZMQSocket($context, ZMQ::SOCKET_REQ);
//$requester->connect("tcp://127.0.0.1:17563");
$requester->connect("ipc://myprocess");

for ($request_nbr = 0; $request_nbr != 10; $request_nbr++) {
    printf ("Sending request %d…\n", $request_nbr);
    $requester->send("Hello");

    $reply = $requester->recv();
    printf ("Received reply %d: [%s]\n", $request_nbr, $reply);
}
