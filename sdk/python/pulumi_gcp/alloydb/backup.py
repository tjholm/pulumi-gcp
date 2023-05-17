# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BackupArgs', 'Backup']

@pulumi.input_type
class BackupArgs:
    def __init__(__self__, *,
                 backup_id: pulumi.Input[str],
                 cluster_name: pulumi.Input[str],
                 location: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_config: Optional[pulumi.Input['BackupEncryptionConfigArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Backup resource.
        :param pulumi.Input[str] backup_id: The ID of the alloydb backup.
        :param pulumi.Input[str] cluster_name: The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        :param pulumi.Input[str] location: The location where the alloydb backup should reside.
        :param pulumi.Input[str] description: User-provided description of the backup.
        :param pulumi.Input['BackupEncryptionConfigArgs'] encryption_config: EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-defined labels for the alloydb backup.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        pulumi.set(__self__, "backup_id", backup_id)
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "location", location)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encryption_config is not None:
            pulumi.set(__self__, "encryption_config", encryption_config)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Input[str]:
        """
        The ID of the alloydb backup.
        """
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        """
        The location where the alloydb backup should reside.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided description of the backup.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptionConfig")
    def encryption_config(self) -> Optional[pulumi.Input['BackupEncryptionConfigArgs']]:
        """
        EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        Structure is documented below.
        """
        return pulumi.get(self, "encryption_config")

    @encryption_config.setter
    def encryption_config(self, value: Optional[pulumi.Input['BackupEncryptionConfigArgs']]):
        pulumi.set(self, "encryption_config", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-defined labels for the alloydb backup.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


@pulumi.input_type
class _BackupState:
    def __init__(__self__, *,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_config: Optional[pulumi.Input['BackupEncryptionConfigArgs']] = None,
                 encryption_infos: Optional[pulumi.Input[Sequence[pulumi.Input['BackupEncryptionInfoArgs']]]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reconciling: Optional[pulumi.Input[bool]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Backup resources.
        :param pulumi.Input[str] backup_id: The ID of the alloydb backup.
        :param pulumi.Input[str] cluster_name: The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        :param pulumi.Input[str] create_time: Time the Backup was created in UTC.
        :param pulumi.Input[str] description: User-provided description of the backup.
        :param pulumi.Input['BackupEncryptionConfigArgs'] encryption_config: EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['BackupEncryptionInfoArgs']]] encryption_infos: EncryptionInfo describes the encryption information of a cluster or a backup.
               Structure is documented below.
        :param pulumi.Input[str] etag: A hash of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-defined labels for the alloydb backup.
        :param pulumi.Input[str] location: The location where the alloydb backup should reside.
        :param pulumi.Input[str] name: Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reconciling: If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
        :param pulumi.Input[str] state: The current state of the backup.
        :param pulumi.Input[str] uid: Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
        :param pulumi.Input[str] update_time: Time the Backup was updated in UTC.
        """
        if backup_id is not None:
            pulumi.set(__self__, "backup_id", backup_id)
        if cluster_name is not None:
            pulumi.set(__self__, "cluster_name", cluster_name)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encryption_config is not None:
            pulumi.set(__self__, "encryption_config", encryption_config)
        if encryption_infos is not None:
            pulumi.set(__self__, "encryption_infos", encryption_infos)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if reconciling is not None:
            pulumi.set(__self__, "reconciling", reconciling)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the alloydb backup.
        """
        return pulumi.get(self, "backup_id")

    @backup_id.setter
    def backup_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_id", value)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> Optional[pulumi.Input[str]]:
        """
        The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the Backup was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided description of the backup.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptionConfig")
    def encryption_config(self) -> Optional[pulumi.Input['BackupEncryptionConfigArgs']]:
        """
        EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        Structure is documented below.
        """
        return pulumi.get(self, "encryption_config")

    @encryption_config.setter
    def encryption_config(self, value: Optional[pulumi.Input['BackupEncryptionConfigArgs']]):
        pulumi.set(self, "encryption_config", value)

    @property
    @pulumi.getter(name="encryptionInfos")
    def encryption_infos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackupEncryptionInfoArgs']]]]:
        """
        EncryptionInfo describes the encryption information of a cluster or a backup.
        Structure is documented below.
        """
        return pulumi.get(self, "encryption_infos")

    @encryption_infos.setter
    def encryption_infos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackupEncryptionInfoArgs']]]]):
        pulumi.set(self, "encryption_infos", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        A hash of the resource.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        User-defined labels for the alloydb backup.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        The location where the alloydb backup should reside.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def reconciling(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
        """
        return pulumi.get(self, "reconciling")

    @reconciling.setter
    def reconciling(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reconciling", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The current state of the backup.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[str]]:
        """
        Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
        """
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uid", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time the Backup was updated in UTC.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Backup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_config: Optional[pulumi.Input[pulumi.InputType['BackupEncryptionConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        An AlloyDB Backup.

        To get more information about Backup, see:

        * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.backups/create)
        * How-to Guides
            * [AlloyDB](https://cloud.google.com/alloydb/docs/)

        ## Example Usage
        ### Alloydb Backup Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.get_network(name="alloydb-network")
        default_cluster = gcp.alloydb.Cluster("defaultCluster",
            cluster_id="alloydb-cluster",
            location="us-central1",
            network=default_network.id)
        private_ip_alloc = gcp.compute.GlobalAddress("privateIpAlloc",
            address_type="INTERNAL",
            purpose="VPC_PEERING",
            prefix_length=16,
            network=default_network.id)
        vpc_connection = gcp.servicenetworking.Connection("vpcConnection",
            network=default_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[private_ip_alloc.name])
        default_instance = gcp.alloydb.Instance("defaultInstance",
            cluster=default_cluster.name,
            instance_id="alloydb-instance",
            instance_type="PRIMARY",
            opts=pulumi.ResourceOptions(depends_on=[vpc_connection]))
        default_backup = gcp.alloydb.Backup("defaultBackup",
            location="us-central1",
            backup_id="alloydb-backup",
            cluster_name=default_cluster.name,
            opts=pulumi.ResourceOptions(depends_on=[default_instance]))
        ```
        ### Alloydb Backup Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.get_network(name="alloydb-network")
        default_cluster = gcp.alloydb.Cluster("defaultCluster",
            cluster_id="alloydb-cluster",
            location="us-central1",
            network=default_network.id)
        private_ip_alloc = gcp.compute.GlobalAddress("privateIpAlloc",
            address_type="INTERNAL",
            purpose="VPC_PEERING",
            prefix_length=16,
            network=default_network.id)
        vpc_connection = gcp.servicenetworking.Connection("vpcConnection",
            network=default_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[private_ip_alloc.name])
        default_instance = gcp.alloydb.Instance("defaultInstance",
            cluster=default_cluster.name,
            instance_id="alloydb-instance",
            instance_type="PRIMARY",
            opts=pulumi.ResourceOptions(depends_on=[vpc_connection]))
        default_backup = gcp.alloydb.Backup("defaultBackup",
            location="us-central1",
            backup_id="alloydb-backup",
            cluster_name=default_cluster.name,
            description="example description",
            labels={
                "label": "key",
            },
            opts=pulumi.ResourceOptions(depends_on=[default_instance]))
        ```

        ## Import

        Backup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:alloydb/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
        ```

        ```sh
         $ pulumi import gcp:alloydb/backup:Backup default {{project}}/{{location}}/{{backup_id}}
        ```

        ```sh
         $ pulumi import gcp:alloydb/backup:Backup default {{location}}/{{backup_id}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_id: The ID of the alloydb backup.
        :param pulumi.Input[str] cluster_name: The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        :param pulumi.Input[str] description: User-provided description of the backup.
        :param pulumi.Input[pulumi.InputType['BackupEncryptionConfigArgs']] encryption_config: EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
               Structure is documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-defined labels for the alloydb backup.
        :param pulumi.Input[str] location: The location where the alloydb backup should reside.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        An AlloyDB Backup.

        To get more information about Backup, see:

        * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.backups/create)
        * How-to Guides
            * [AlloyDB](https://cloud.google.com/alloydb/docs/)

        ## Example Usage
        ### Alloydb Backup Basic

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.get_network(name="alloydb-network")
        default_cluster = gcp.alloydb.Cluster("defaultCluster",
            cluster_id="alloydb-cluster",
            location="us-central1",
            network=default_network.id)
        private_ip_alloc = gcp.compute.GlobalAddress("privateIpAlloc",
            address_type="INTERNAL",
            purpose="VPC_PEERING",
            prefix_length=16,
            network=default_network.id)
        vpc_connection = gcp.servicenetworking.Connection("vpcConnection",
            network=default_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[private_ip_alloc.name])
        default_instance = gcp.alloydb.Instance("defaultInstance",
            cluster=default_cluster.name,
            instance_id="alloydb-instance",
            instance_type="PRIMARY",
            opts=pulumi.ResourceOptions(depends_on=[vpc_connection]))
        default_backup = gcp.alloydb.Backup("defaultBackup",
            location="us-central1",
            backup_id="alloydb-backup",
            cluster_name=default_cluster.name,
            opts=pulumi.ResourceOptions(depends_on=[default_instance]))
        ```
        ### Alloydb Backup Full

        ```python
        import pulumi
        import pulumi_gcp as gcp

        default_network = gcp.compute.get_network(name="alloydb-network")
        default_cluster = gcp.alloydb.Cluster("defaultCluster",
            cluster_id="alloydb-cluster",
            location="us-central1",
            network=default_network.id)
        private_ip_alloc = gcp.compute.GlobalAddress("privateIpAlloc",
            address_type="INTERNAL",
            purpose="VPC_PEERING",
            prefix_length=16,
            network=default_network.id)
        vpc_connection = gcp.servicenetworking.Connection("vpcConnection",
            network=default_network.id,
            service="servicenetworking.googleapis.com",
            reserved_peering_ranges=[private_ip_alloc.name])
        default_instance = gcp.alloydb.Instance("defaultInstance",
            cluster=default_cluster.name,
            instance_id="alloydb-instance",
            instance_type="PRIMARY",
            opts=pulumi.ResourceOptions(depends_on=[vpc_connection]))
        default_backup = gcp.alloydb.Backup("defaultBackup",
            location="us-central1",
            backup_id="alloydb-backup",
            cluster_name=default_cluster.name,
            description="example description",
            labels={
                "label": "key",
            },
            opts=pulumi.ResourceOptions(depends_on=[default_instance]))
        ```

        ## Import

        Backup can be imported using any of these accepted formats

        ```sh
         $ pulumi import gcp:alloydb/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
        ```

        ```sh
         $ pulumi import gcp:alloydb/backup:Backup default {{project}}/{{location}}/{{backup_id}}
        ```

        ```sh
         $ pulumi import gcp:alloydb/backup:Backup default {{location}}/{{backup_id}}
        ```

        :param str resource_name: The name of the resource.
        :param BackupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_id: Optional[pulumi.Input[str]] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_config: Optional[pulumi.Input[pulumi.InputType['BackupEncryptionConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupArgs.__new__(BackupArgs)

            if backup_id is None and not opts.urn:
                raise TypeError("Missing required property 'backup_id'")
            __props__.__dict__["backup_id"] = backup_id
            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["description"] = description
            __props__.__dict__["encryption_config"] = encryption_config
            __props__.__dict__["labels"] = labels
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["encryption_infos"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["reconciling"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        super(Backup, __self__).__init__(
            'gcp:alloydb/backup:Backup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backup_id: Optional[pulumi.Input[str]] = None,
            cluster_name: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encryption_config: Optional[pulumi.Input[pulumi.InputType['BackupEncryptionConfigArgs']]] = None,
            encryption_infos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupEncryptionInfoArgs']]]]] = None,
            etag: Optional[pulumi.Input[str]] = None,
            labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            location: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            reconciling: Optional[pulumi.Input[bool]] = None,
            state: Optional[pulumi.Input[str]] = None,
            uid: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Backup':
        """
        Get an existing Backup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_id: The ID of the alloydb backup.
        :param pulumi.Input[str] cluster_name: The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        :param pulumi.Input[str] create_time: Time the Backup was created in UTC.
        :param pulumi.Input[str] description: User-provided description of the backup.
        :param pulumi.Input[pulumi.InputType['BackupEncryptionConfigArgs']] encryption_config: EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
               Structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackupEncryptionInfoArgs']]]] encryption_infos: EncryptionInfo describes the encryption information of a cluster or a backup.
               Structure is documented below.
        :param pulumi.Input[str] etag: A hash of the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: User-defined labels for the alloydb backup.
        :param pulumi.Input[str] location: The location where the alloydb backup should reside.
        :param pulumi.Input[str] name: Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[bool] reconciling: If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
        :param pulumi.Input[str] state: The current state of the backup.
        :param pulumi.Input[str] uid: Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
        :param pulumi.Input[str] update_time: Time the Backup was updated in UTC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupState.__new__(_BackupState)

        __props__.__dict__["backup_id"] = backup_id
        __props__.__dict__["cluster_name"] = cluster_name
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["encryption_config"] = encryption_config
        __props__.__dict__["encryption_infos"] = encryption_infos
        __props__.__dict__["etag"] = etag
        __props__.__dict__["labels"] = labels
        __props__.__dict__["location"] = location
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["reconciling"] = reconciling
        __props__.__dict__["state"] = state
        __props__.__dict__["uid"] = uid
        __props__.__dict__["update_time"] = update_time
        return Backup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupId")
    def backup_id(self) -> pulumi.Output[str]:
        """
        The ID of the alloydb backup.
        """
        return pulumi.get(self, "backup_id")

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time the Backup was created in UTC.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        User-provided description of the backup.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptionConfig")
    def encryption_config(self) -> pulumi.Output[Optional['outputs.BackupEncryptionConfig']]:
        """
        EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
        Structure is documented below.
        """
        return pulumi.get(self, "encryption_config")

    @property
    @pulumi.getter(name="encryptionInfos")
    def encryption_infos(self) -> pulumi.Output[Sequence['outputs.BackupEncryptionInfo']]:
        """
        EncryptionInfo describes the encryption information of a cluster or a backup.
        Structure is documented below.
        """
        return pulumi.get(self, "encryption_infos")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A hash of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        User-defined labels for the alloydb backup.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The location where the alloydb backup should reside.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The ID of the project in which the resource belongs.
        If it is not provided, the provider project is used.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def reconciling(self) -> pulumi.Output[bool]:
        """
        If true, indicates that the service is actively updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
        """
        return pulumi.get(self, "reconciling")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the backup.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time the Backup was updated in UTC.
        """
        return pulumi.get(self, "update_time")

