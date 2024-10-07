# Simple Reverse HTTP Proxy

Allow url requests to the same host be routed to different ports based on their paths  
Not recommended for Use in Production  
Example:
```
https://localhost/survey/home?color=red
Can map to:
http://localhost:3000/home?color=red
```
*Will copy headers and payload back and forth*


### Possible Applications
- One process forwards requests to docker containers running on other ports
- Forward requests to other hosts in the same network
- Require authentication for certain routes but not others
