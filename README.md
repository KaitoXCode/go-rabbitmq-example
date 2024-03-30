# go-rabbitmq-example

example messaging flow

# structure

![alt text](https://github.com/KaitoXCode/go-rabbitmq-example/blob/master/public/message-flow.jpg?raw=true)

# make cmds

launch consumer one: `make consumer-one` <br /> launch consumer two:
`make consumer-two` <br /> launch consumer three: `make consumer-three` <br />
<br /> send message from producer one to consumer one: <br />
`make producer-one-msg RKEY=11` <br /> send message from producer one to
consumer two: <br /> `make producer-one-msg RKEY=22` <br /> send message from
producer one to consumer three: <br /> `make producer-one-msg RKEY=33` <br />
<br /> send message from producer two to consumer one: <br />
`make producer-two-msg RKEY=11` <br /> send message from producer two to
consumer two: <br /> `make producer-two-msg RKEY=22` <br /> send message from
producer two to consumer three: <br /> `make producer-two-msg RKEY=33` <br />
