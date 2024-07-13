# usage
to start _simplehttp_ web server , map local port 8000 into container and provide the env variable **message**:  
```bash
sudo docker run -e message=Hello -p 8000:8000 simplehttp
```  
