# Go Alfred
It's a command line to make easy some repetable tasks.  

## Summary
### encode
It's possible encode and decode base64 strings  

```
alfred base64 -encode  
```
or  
```
alfred base64 -encode  
```

### HTTP Request
It's possible make a HTTP request using any HTTP verbs.  
The command use a YAML file to read parameter to make the request. There is a example of YAML file on the root repository.  
If you will do a GET request, remove the **body** section fom file.  

```
alfred url -request -file file_example.yml
```
It'll return the status code and the body of HTTP response.