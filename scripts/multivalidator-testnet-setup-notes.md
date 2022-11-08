# Multivalidator testnet setup notes

It is necessary to wait for the deployment of the initial node before building the image of the following nodes.
Here is a list of things that can only be known after the first image is built/deployed.

- RPC address may be known in advance if that address will be pointed toward the new node evenetually
- the persistent peer node id name (requires /root/.duality/config/node_key.json, i.e. `dualityd init` to be run first)
- the persistent peer IPV4 address (not known until AWS )
- the new genesis file to use that was created in the initial nodes startup script process
- the new mnemonic that was generated by the new gentx actions before the chain started
- the docker image of the followers: can only be made after the chain.json, and genesis.json are updated

## What could be improved

### What could be improved about the Docker setup

Currently, 2 Docker builds+images are needed. This seems unnecessary.
The main reason for this is that the genesis file needs to be updated,
and the persistent peer ID and location need to be update.
- its possible to pre-make the genesis file in an image (so it can become a Docker dependency)
- its possible to determine the node id in an image (so it can become a Docker dependency)
- however it is not always possible to know the IP addresses to the initial node before it is deployed
  - this could be overcome by using AWS Elastic IP addresses and re-assigning them as needed or
    using long-lived load balancer IPs that redirect to new instances after they update (as happens in ECS service ALBs)

Common ------------------> common ----------------------------------> dualityd init --> initial nodes
        |                                                  |     |
        --> dualityd init --> compute configs into files --       --> dualityd init --> follower nodes

### What could be improved about the AWS setup

Currently a process is setup to automatically deploy to Elastic Beanstalk (EB), but Elastic Beanstalk doesn't have OS level metrics.
This is a concern because the Cosmos chains appear to be mostly limited by memory limits, not CPU limits. There is a process
to add CloudWatch memory logging into Elastic Beanstalk but I haven't seen an example of it working with the Docker images
and could not get the Docker images working with OS memory logging.

I've found Elastic Container Service (ECS) to be a much better fit, as it is a container service with OS level metrics already well-handled.
I don't currently have a scripted process to deploy these and have instead just point+clicked these into existence using the AWS web GUI.
It would be a lot quicker, and more reliable to be able to have these set up in a simple script or well defined in one CloudFormation deploy script.

We could then very easily deploy a dev-net in the continuous integration (CI) process when merging/previewing branches in the core repo.