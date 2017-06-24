## API Documentation

### POST /api/search
Returns a list of video

### GET /api/queue&channel={channel_id}
Returns a list of queue

### POST /api/queue&channel={channel_id}
Add to queue

### POST /api/users
#### Request
```json
{
    "nickname": "tom",
    "start_channel_id": 2,
}
```

#### Response
```json
{
    "id": 1,
    "nickname": "tom"
}
```

### POST /api/add-channel-user
#### Request
```json
{
    "channel_id": 1,
    "user_id": 1
}
```

### POST /api/remove-channel-user
#### Request
```json
{
    "channel_id": 1,
    "user_id": 1
}
```

### POST /channel
Start a new channel
Returns a url with unique id

#### Request
```json
{
    "name": "k-pop rocks",
    "description": "k-pop rocks a lot",
    "created_by": 1
}
```

#### Response
```json
{    
    "id": 1,
    "name": "k-pop rocks",
    "description": "k-pop rocks a lot",
    "channel_url": "chorus.com/asd55cz8xsa",
    "created_by": 2
}
```

## Websocket communication
```json
{
    "cmd": "cmd_string"
}
```

### GET /api/stream&channel={channel_id}
Connect to channel WS Stream
Audio data will be streamed through binary

### WS Pause
```json
{
    "cmd": "PAUSE"
}
```

### WS PLAY
```json
{
    "cmd": "PAUSE"
}
```
