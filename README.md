# hajimete

## haji-what ?

hajimete means first time in japanese.

We'll code an API in go with couchbase, to track every 'first time' of my baby (first smile, first words, etc).

## What we'll need

Obviously we'll need

 * a working installation of couchbase server (entreprise version 5.0.1 for now)
 * a working installation of go (1.9.2)
 * the couchbase go sdk to link them
 * optionnal: a baby

### Couchbase

So let's go, we will download and install couchbase.

```bash
# Download couchbase rpm
curl -O https://packages.couchbase.com/releases/5.0.1/couchbase-server-enterprise-5.0.1-centos7.x86_64.rpm
# Install couchbase
rpm --install couchbase-server-enterprise-5.0.1-centos7.x86_64.rpm
# we want couhbase to start with the VM
systemctl enable couchbase-server
# and we want couchbase to be running right now
systemctl start couchbase-server
```

Now couchbase is waiting for us on http://localhost:8091 and we could configure everything on the UI ... but we'll stay on our dear shell and configure everything here

With that, we'll have a cluster with couchbase/couchbase as credentials (Security is important, you know). No need for full text search service for now, so only the three others will be activated (data, index and query).

We'll set ram quota to 256Mo for data service and 256Mo for index service (Sizing is important, you know....)

```bash
/opt/couchbase/bin/couchbase-cli cluster-init \
  --cluster-username=couchbase \
  --cluster-password=couchbase \
  --services=data,index,query \
  --cluster-ramsize=256 \
  --cluster-index-ramsize=256
```

Our application will need only one bucket (think about a database in mySql) so we can give everything we have on the cluster to this bucket

```bash
/opt/couchbase/bin/couchbase-cli bucket-create \
  -c localhost -u couchbase -p couchbase \
  --bucket hajimete  \
  --bucket-type couchbase  \
  --bucket-ramsize 256
```
