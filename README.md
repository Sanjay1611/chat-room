# chat-room

## Run the server
`./main`

## use following way to do corresponding tasks

- `/chat_room` [GET] -> Get all chat rooms

- `/chat_room` [POST] -> Create a chatroom

RequestPayload
```
{
name: <name of chat room>
}
```

- `/user` [GET] -> Get all users

- `/user` [POST] -> Create a user

Request Payload
```
{
name: <name of user>
age: <age of user>
}
```

- /`user` [PUT] -> User wants to join a chat room

Request Payload
```
{
user_id: <uid of the user>
chat_room_id: <uid of the chat he wants to join>
}
```

- `/message` [GET] -> Get all messages

- `/message` [POST] -> Create a message

Request Payload
```
{
content: <content of the message>
user_id: <uid of the user>
}
```
