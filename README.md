# Simple Reverse HTTP Proxy

Allow url requests to the same host be routed to different ports based on their paths  
**Not recommended for Use in Production  **
Example:
```
https://localhost/survey/home?color=red
Can map to:
http://localhost:3000/home?color=red
```
*Will copy headers and payload back and forth*  


## Setup
config.ini
```ini
REFRESH_PASSCODE=abc123

[PORTS]
default=9001
survey=9001
```
Once config.ini is setup, it will map routes on the left to ports on the right  
You are ready to run main  

If you want to refresh this config file, without restarting the proxy, navigate to this url using the REFRESH_PASSCODE setup in your ini:
`http://localhost/refresh?passcode=abc123`

## Possible Applications
- One process forwards requests to docker containers running on other ports
- Forward requests to other hosts in the same network
- Require authentication for certain routes but not others
