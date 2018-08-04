# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetRegionsResult(object):
    """
    A collection of values returned by getRegions.
    """
    def __init__(__self__, names=None, project=None, id=None):
        if names and not isinstance(names, list):
            raise TypeError('Expected argument names to be a list')
        __self__.names = names
        """
        A list of regions available in the given project
        """
        if project and not isinstance(project, basestring):
            raise TypeError('Expected argument project to be a basestring')
        __self__.project = project
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_regions(project=None, status=None):
    """
    Provides access to available Google Compute regions for a given project.
    See more about [regions and regions](https://cloud.google.com/compute/docs/regions-zones/) in the upstream docs.
    
    ```
    data "google_compute_regions" "available" {}
    
    resource "google_compute_subnetwork" "cluster" {
      count = "${length(data.google_compute_regions.available.names)}"
      name          = "my-network"
      ip_cidr_range = "10.36.${count.index}.0/24"
      network       = "my-network"
      region        = "${data.google_compute_regions.available.names[count.index]}"
    }
    ```
    """
    __args__ = dict()

    __args__['project'] = project
    __args__['status'] = status
    __ret__ = pulumi.runtime.invoke('gcp:compute/getRegions:getRegions', __args__)

    return GetRegionsResult(
        names=__ret__.get('names'),
        project=__ret__.get('project'),
        id=__ret__.get('id'))
