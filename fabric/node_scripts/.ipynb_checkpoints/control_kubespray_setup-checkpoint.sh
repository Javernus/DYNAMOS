#!/bin/bash

{
git clone https://github.com/kubernetes-sigs/kubespray.git
cd kubespray
git checkout release-2.27
rm -rf .git

python3 -m venv venv
source ./venv/bin/activate

pip3 install -r requirements.txt

mv inventory/sample inventory/dynamos

# sed -i -e 's/kube_proxy_mode: ipvs/kube_proxy_mode: iptables/g' ./inventory/dynamos/group_vars/k8s_cluster/k8s-cluster.yml

ansible --version
}
