# download package
curl -OL https://redirector.gvt1.com/edgedl/go/go1.9.2.linux-amd64.tar.gz
# untar package
tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
# add go binaries to PATH
sudo echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
# go use git to download package
sudo yum install git -y
# vagrant...
chown -R vagrant:vagrant /home/vagrant/go
