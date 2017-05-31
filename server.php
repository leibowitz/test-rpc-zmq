<?php
$context = new ZMQContext(1);

//  Socket to talk to clients
$responder = new ZMQSocket($context, ZMQ::SOCKET_REP);
//$responder->bind("tcp://*:17563");
$responder->bind("ipc://myprocess");

while (true) {
    //  Wait for next request from client
    $request = $responder->recv();
    //$r = unserialize($request);
    var_dump($request);
    printf ("Received request: \n");
    //printf ("Received request: [%s]\n", $request);

    //  Do some 'work'
    sleep (1);

    //  Send reply back to client
    $responder->send("World");
}
