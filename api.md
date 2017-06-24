## API Documentation

### POST /api/search
Returns a list of youtube search results

## Channel
### GET /api/channel
Return channel details

### POST /api/channel
Start a new channel

#### Request
```json
{
    "name": "k-pop rocks",
    "description": "k-pop rocks a lot",
    "creator_name": "harry"
}
```

#### Response
```json
{    
    "id": 1,
    "name": "k-pop rocks",
    "description": "k-pop rocks a lot",
    "channel_url": "chorus.com/asd55cz8xsa"
}
```

### GET /api/channel/:id/queue
Returns a list of queue

### POST /api/channel/:id/queue
Add to queue

**NOT IN USE**
### GET /api/channel/:id/users
Return all users in channel

### POST /api/channel/:id/users
For non-creators, need to send nickname to server.
#### Request
```json
{
    "nickname": "harry"
}
```

## Websocket communication
```json
{
    "cmd": "cmd_string"
}
```

### GET /api/channel/:id/stream
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
