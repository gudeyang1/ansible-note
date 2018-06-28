
## 0. prepare
```bash
3 nodes
    chef-workstation
    chef-server
    vm-node1
time sync on each node
ip hostname map in /etc/hosts
```
https://learn.chef.io/#/

## 1. install chef-server on chef-server node
```bash
https://docs.chef.io/install_server.html
Download the package from https://downloads.chef.io/chef-server/.
chef-server-ctl reconfigure
chef-server-ctl user-create chefadmin Chef Admin admin@4thcoffee.com insecurepassword --filename /root/chefadmin.pem
chef-server-ctl org-create 4thcoffee "Fourth Coffee, Inc." --association_user chefadmin --filename 4thcoffee-validator.pem
#install chef manage
chef-server-ctl install chef-manage --path /root/packages
//The chef-server-ctl command will install the first chef-manage package found in the /root/packages directory.
```
## 2. install chef workstation on chef-workstation node
### 2.1 download and install
```bash
https://downloads.chef.io/chefdk
verfy installation
chef --version
```

### 2.2 configure
```bash
chef-server 和chef-workstation 通过证书通信
//把chef-server 上创建admin生成的证书scp 到wokrstation节点的repo
scp ubuntu@chef-server-hostname:/root/chefadmin.pem ~/learn-chef/.chef/chefadmin.pem

//添加chef环境文件 repo/.chef/knife.rb
current_dir = File.dirname(__FILE__)
log_level                 :info
log_location              STDOUT
node_name                 "chefadmin"
client_key                "#{current_dir}/chefadmin.pem"
chef_server_url           "https://chef-server-hostname/organizations/4thcoffee"
cookbook_path             ["#{current_dir}/../cookbooks"]
```

### 2.3 test
```bash
cd ~/lean-chef

knife ssl fetch
knife ssl check
//如果chek不过,说明证书不对, 看报错信息, 可以修改chef-server /opt下面的nginx 配置文件, 然后chef-server-ctl reconfigure  再次 knife ssl fetch && check

```

## 3. upload a cookbook to chef chef-server

```bash
mkdir ~/learn-chef/cookbooks
cd  ~/lean-chef/cookbooks
git clone https://github.com/learn-chef/learn_chef_httpd.git
knife cookbook upload learn_chef_httpd  //通过之前配置的证书通信
knife cookbook list
```

## 4. 添加node 节点
```bash
//node节点需要能通过hostname 找到chf-server
knife bootstrap $ADDRESS --ssh-user $USER --ssh-password '$PASSWORD' --sudo --use-sudo-password --node-name node1-centos --run-list 'recipe[learn_chef_httpd]'


//check
knife node list

```

## 5. update node configuration
### 5.1 use template
```bash
[root@node learn_chef_httpd]# cat templates/index.html.erb
<html>
  <body>
    <h1>hello from <%= node['fqdn'] %></h1>
  </body>
</html>
```
### 5.2 update cookbook version in metadata
```
version 'v.01.0'
```
### 5.3 upload cookbook to chef_server
```bash
knife cookbook upload learn_chef_httpd
//run cookbook on your node, name是bootstarp 的时候指定的nodename, attribute 写ip, 参考 https://learn.chef.io/modules/manage-a-node-chef-server/rhel/bring-your-own-system/update-your-nodes-configuration#/
knife ssh 'name:node1-centos' 'sudo chef-client' --ssh-user root --ssh-password "123456" --attribute 192.168.232.133
```
### 5.4 复杂点的recipe
```ruby
package 'httpd'

service 'httpd' do
  action [:enable, :start]
end
//先创建组在创建用户, 执行是按照recipe 的顺序
group 'web_admin'

user 'web_admin' do
  group 'web_admin'
  system true
  shell '/bin/bash'
end

template '/var/www/html/index.html' do # ~FC033
  source 'index.html.erb'
  mode '0644'
  owner 'web_admin'
  group 'web_admin'
end
```
