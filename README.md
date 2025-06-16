# Hello, Gophercloud

A minimal example of Gophercloud usage targeting different services.

```
go run main.go
```

This can be run against a DevStack deployment using the following `local.conf` (assuming 2025.2 Flamingo).
This is derived from the configuration used in Gophercloud's CI.

```ini
[[local|localrc]]
ADMIN_PASSWORD=secret
DATABASE_PASSWORD=root
RABBIT_PASSWORD=secret
SERVICE_PASSWORD=secret
SWIFT_HASH=1234123412341234
LOGFILE=/tmp/devstack-logs/devstack.log
INSTALL_TEMPEST=False
GIT_BASE=https://github.com

GLANCE_LIMIT_IMAGE_SIZE_TOTAL=5000
SWIFT_MAX_FILE_SIZE=5368709122
KEYSTONE_ADMIN_ENDPOINT=true

ENABLED_SERVICES+=,-horizon,-dstat,-tempest,openstack-cli-server
ENABLED_SERVICES+=,designate,designate-central,designate-api,designate-worker,designate-producer,designate-mdns
ENABLED_SERVICES+=,barbican-svc,barbican-retry,barbican-keystone-listener
ENABLED_SERVICES+=,h-eng,h-api,h-api-cfn,h-api-cw
ENABLED_SERVICES+=,magnum-api,magnum-cond
ENABLED_SERVICES+=,mistral,mistral-api,mistral-engine,mistral-executor,mistral-event-engine

enable_plugin designate https://github.com/openstack/designate master
enable_plugin barbican https://github.com/openstack/barbican master
enable_plugin heat https://github.com/openstack/heat master
enable_plugin magnum https://github.com/openstack/magnum master
enable_plugin mistral https://github.com/openstack/mistral master

MAGNUMCLIENT_BRANCH=master

# ensure we're using a working version of setuptools
if [ -n "$TOP_DIR" ]; then
  sed -i 's/setuptools\[core\]$/setuptools[core]==79.0.1/g' $TOP_DIR/lib/infra $TOP_DIR/inc/python
  sed -i 's/pip_install "-U" "pbr"/pip_install "-U" "pbr" "setuptools[core]==79.0.1"/g' $TOP_DIR/lib/infra
fi
```

You can use a local copy of `gophercloud/gophercloud` and `gophercloud/utils` with the `replace` directive.

```
replace (
	github.com/gophercloud/gophercloud/v2 => /home/user/code/gophercloud/gophercloud
	github.com/gophercloud/utils/v2 => /home/user/code/gophercloud/utils
)
```
