## Golang websocket on OpenShift

### How to deploy
```
git clone https://github.com/kevydotvinu/ws && cd ws
oc new-project ws
oc new-app --name=ws --strategy=docker . --allow-missing-images
oc start-build ws --from-dir=./ --follow=true --wait=true
oc expose svc ws
oc patch route ws --type merge --patch '{"spec":{"tls":{"termination":"edge"}}}'
websocat -k wss://<route>/echo
```
