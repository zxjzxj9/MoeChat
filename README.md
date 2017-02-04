# MoeChat
A chatting tool with server implemented by Go, having multiple clients implements including Go, Python etc.

M-O-E is short for multiple-online-exchange

Authored by Victor Xiao-Jie Zhang, 2017/2/4

# Program details
##1.The MoeChat protocal
###1.1 Using JSON as the underlying message exchange data-format
JSON is the basic exchange data-format for MoeChat, which is exampled as follows,

{"status": "q", "status_code": 80, "message": "", ...}

The most important part is status, and the status_code, which will guide the interactions between server and client.

Where status has type String, containing only one character, status_code is a Number ranging from 0 to 100.

Following is a brief description of the status and status_code

### Table 1 The status summary
|Status    |Meaning                     |Accompying status_code    |
|----------|----------------------------|--------------------------|
|"q"       |Querying for information    |Any is OK                 |
|"r"       |Response to query           |10 - 40                   |
|"e"       |Some error happened         |50 - 70                   |
|"m"       |Transferring message        |80 - 90                   |
