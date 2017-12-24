Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"

  config.vm.provider :virtualbox do |v|
    v.customize [ "modifyvm", :id, "--cpus", 2 ]
    v.customize [ "modifyvm", :id, "--memory", 1024 ]
  end

  config.vm.synced_folder ".", "/vagrant", type: "nfs"
  config.vm.synced_folder "src", "/home/vagrant/go/src/github.com/bpaulin/hajimete", type: "nfs", owner: "vagrant", group: "vagrant"

  config.vm.network "forwarded_port", guest: 8080, host: 8080
  config.vm.network "forwarded_port", guest: 8091, host: 8091
end
