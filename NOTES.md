### TODO

- Weighted Random balancer
- Request / Response struct


### Requirements

- Should implements nginx load balancing behaviour using the server directives http://nginx.org/en/docs/http/ngx\_http\_upstream\_module.html#server, the pririties are:
    - default should offer weighted round-robin
    - max\_fails options
    - fail\_status\_code standards detect fail options using HTTP status code
    - fail\_timeout: time window to observe max\_fails attempts

- non-priorities:
    - DNS resolving functionalities


### Design


There are many strategies for designing interfaces

#### http.Transport wrapper

One possible strategy is to implement as a Transport wrapper, so the method calls graph looks like this

hcl.Client.Do -> http.Client.Do -> hcl.Transport -> balancer.Get() -> hcl.Transport.Base ( http.Transport ) -> host

Pros:

- Minimal design, utilising existing library

Cons:

Uknowns:

- May be restricting in terms of features, define more scope of features in (#Requirements) to better understand the limitation.


#### Behaviours

##### Redirect

How should a HTTP client with balancer process redirect requests? :


- If same host / different endpoint: redirect to new URL using same host / new endpoint
- OR if same host / different endpoint: redirect to new URL using new host in balancing scheme for new request

- If different host: use new host for new request
