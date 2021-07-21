### Use with Linode

Example:

```
export LINODE_TOKEN=[LINODE_TOKEN]
./terraformer import linode -r instance
```

List of supported Linode resources:

*   `domain`
    * `linode_domain`
    * `linode_domain_record`
*   `firewall`
    * `linode_firewall`    
*   `image`
    * `linode_image`
*   `instance`
    * `linode_instance`
*   `lke`
    * `linode_lke_cluster`
*   `nodebalancer`
    * `linode_nodebalancer`
    * `linode_nodebalancer_config`
    * `linode_nodebalancer_node`
*   `objectstorage`
    * `linode_object_storage_bucket`
*   `rdns`
    * `linode_rdns`
*   `sshkey`
    * `linode_sshkey`
*   `stackscript`
    * `linode_stackscript`
*   `token`
    * `linode_token`
*   `user`
    * `linode_user`
*   `volume`
    * `linode_volume`
