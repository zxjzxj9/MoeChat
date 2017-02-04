# MoeChat
A chatting tool with server implemented by Go, having multiple clients implements including Go, Python etc.

M-O-E is short for multiple-online-exchange

Authored by Victor Xiao-Jie Zhang, 2017/2/4

# Program details
## 1.The MoeChat protocal
### 1.1 Using JSON as the underlying message exchange data-format
JSON is the basic exchange data-format for MoeChat, which is exampled as follows,

{"status": "q", "status_code": 80, "info": {...} }

The most important part is status, and the status_code, which will guide the interactions between server and client.

Where status has type String, containing only one character, status_code is a Number ranging from 0 to 100.

Following is a brief description of the status and status_code

#### Table 1 The status summary                           
|Status    |Meaning                     |Accompying status_code    |Notes                            |
|----------|----------------------------|--------------------------|---------------------------------|
|"q"       |Querying for information    |Any is OK                 |"info" Object for query data     |
|"r"       |Response to query           |10 - 50                   |"info" Object for response data  |
|"e"       |Some error happened         |60 - 70                   |"info" Object for error info     |
|"m"       |Transferring message        |80 - 90                   |"info" Object for message detail |

#### Table 2 The status code summary
|Status Code     |Meaning                |Notes                                  |
|----------------|-----------------------|---------------------------------------|
|1x              |Login success          |Message sending port given in "info".  |
|2x              |Login fail             |Access denied, detailed in "info".     |
|3x              |Query success          |Response given in "info"               |
|4x              |Query fail             |No response, reason given in "info"    |
|5x              |Message success        |Message successfully sent              |
|6x              |Server fail            |Server errror, reason given in "info"  |
|7x              |Message fail           |Cannot send message, details in "info" |
|8x              |Message to one person  |Send message to someone                |
|9x              |Boardcasting message   |Send message to everyone               |

