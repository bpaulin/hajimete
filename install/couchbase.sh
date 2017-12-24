## Install
# Download couchbase rpm
curl -O https://packages.couchbase.com/releases/5.0.1/couchbase-server-enterprise-5.0.1-centos7.x86_64.rpm
# Install couchbase
rpm --install couchbase-server-enterprise-5.0.1-centos7.x86_64.rpm
# we want couhbase to start with the VM
systemctl enable couchbase-server
# and we want couchbase to be running right now
systemctl start couchbase-server
## Setup
# Init cluster
/opt/couchbase/bin/couchbase-cli cluster-init \
  --cluster-username=couchbase \
  --cluster-password=couchbase \
  --services=data,index,query \
  --cluster-ramsize=256 \
  --cluster-index-ramsize=256
# Create bucket
/opt/couchbase/bin/couchbase-cli bucket-create \
  -c localhost \
  -u couchbase \
  -p couchbase \
  --bucket hajimete \
  --bucket-type couchbase \
  --bucket-ramsize 256
