# twitch-streaming

### Time management
For implementing backend services I spend almost 4 hours. For frontend ui - almost 5 hours (my first Vue.js application and js application for frontend)

### Architecture

The project consist of three applications:
- API: GraphQL gateway with implemented subscriptions for realtime
- Server: GRPC server to handle callbacks from Twitch and expose data to API
- UI: Vue.js application for UI

I used additional addons:
- PostgreSQL: to save events 
- Redis: for Pub/Sub implementation between API and Server for realtime events
- GRPC + Protobuf: as default protocol for communication between API and Server
- JWT authorization: for auth checks and session management

### Deployment: 

Unfortunately, I can't deploy existing application to Heroku, as I can't configure private network usage without charging.     
I can create docker compose configuration to run it locally, but in this case embedding and webhooks won't work.

### Answers:
##### How to deploy to AWS
To deploy to AWS, first of all, we should build existing apps to pack to docker images and push to some registry.
Then we can use K8S cluster to deploy existing applications. Also, we need to set up postgres database and Redis instance. 
Along with it we will need to configure a traffic balancer, Ingress can be used. 

#### Bottlenecks 
We have separated services, so we may deploy a service in many replicas (to handle huge traffic). 
Postgres can be a bottleneck, but it depends on final implementation and application needs.
Redis will be good enough for realtime needs.

### Todo:
- Add tests 
- Add metrics