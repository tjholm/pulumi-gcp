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

__all__ = ['BackupScheduleArgs', 'BackupSchedule']

@pulumi.input_type
class BackupScheduleArgs:
    def __init__(__self__, *,
                 retention: pulumi.Input[str],
                 daily_recurrence: Optional[pulumi.Input['BackupScheduleDailyRecurrenceArgs']] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input['BackupScheduleWeeklyRecurrenceArgs']] = None):
        """
        The set of arguments for constructing a BackupSchedule resource.
        :param pulumi.Input[str] retention: At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
               A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
               For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
               
               
               - - -
        :param pulumi.Input['BackupScheduleDailyRecurrenceArgs'] daily_recurrence: For a schedule that runs daily at a specified time.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input['BackupScheduleWeeklyRecurrenceArgs'] weekly_recurrence: For a schedule that runs weekly on a specific day and time.
               Structure is documented below.
        """
        pulumi.set(__self__, "retention", retention)
        if daily_recurrence is not None:
            pulumi.set(__self__, "daily_recurrence", daily_recurrence)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if weekly_recurrence is not None:
            pulumi.set(__self__, "weekly_recurrence", weekly_recurrence)

    @property
    @pulumi.getter
    def retention(self) -> pulumi.Input[str]:
        """
        At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
        A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.


        - - -
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: pulumi.Input[str]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> Optional[pulumi.Input['BackupScheduleDailyRecurrenceArgs']]:
        """
        For a schedule that runs daily at a specified time.
        """
        return pulumi.get(self, "daily_recurrence")

    @daily_recurrence.setter
    def daily_recurrence(self, value: Optional[pulumi.Input['BackupScheduleDailyRecurrenceArgs']]):
        pulumi.set(self, "daily_recurrence", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

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
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> Optional[pulumi.Input['BackupScheduleWeeklyRecurrenceArgs']]:
        """
        For a schedule that runs weekly on a specific day and time.
        Structure is documented below.
        """
        return pulumi.get(self, "weekly_recurrence")

    @weekly_recurrence.setter
    def weekly_recurrence(self, value: Optional[pulumi.Input['BackupScheduleWeeklyRecurrenceArgs']]):
        pulumi.set(self, "weekly_recurrence", value)


@pulumi.input_type
class _BackupScheduleState:
    def __init__(__self__, *,
                 daily_recurrence: Optional[pulumi.Input['BackupScheduleDailyRecurrenceArgs']] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input['BackupScheduleWeeklyRecurrenceArgs']] = None):
        """
        Input properties used for looking up and filtering BackupSchedule resources.
        :param pulumi.Input['BackupScheduleDailyRecurrenceArgs'] daily_recurrence: For a schedule that runs daily at a specified time.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] name: The unique backup schedule identifier across all locations and databases for the given project. Format:
               `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] retention: At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
               A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
               For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
               
               
               - - -
        :param pulumi.Input['BackupScheduleWeeklyRecurrenceArgs'] weekly_recurrence: For a schedule that runs weekly on a specific day and time.
               Structure is documented below.
        """
        if daily_recurrence is not None:
            pulumi.set(__self__, "daily_recurrence", daily_recurrence)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if retention is not None:
            pulumi.set(__self__, "retention", retention)
        if weekly_recurrence is not None:
            pulumi.set(__self__, "weekly_recurrence", weekly_recurrence)

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> Optional[pulumi.Input['BackupScheduleDailyRecurrenceArgs']]:
        """
        For a schedule that runs daily at a specified time.
        """
        return pulumi.get(self, "daily_recurrence")

    @daily_recurrence.setter
    def daily_recurrence(self, value: Optional[pulumi.Input['BackupScheduleDailyRecurrenceArgs']]):
        pulumi.set(self, "daily_recurrence", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique backup schedule identifier across all locations and databases for the given project. Format:
        `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
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
    def retention(self) -> Optional[pulumi.Input[str]]:
        """
        At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
        A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.


        - - -
        """
        return pulumi.get(self, "retention")

    @retention.setter
    def retention(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "retention", value)

    @property
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> Optional[pulumi.Input['BackupScheduleWeeklyRecurrenceArgs']]:
        """
        For a schedule that runs weekly on a specific day and time.
        Structure is documented below.
        """
        return pulumi.get(self, "weekly_recurrence")

    @weekly_recurrence.setter
    def weekly_recurrence(self, value: Optional[pulumi.Input['BackupScheduleWeeklyRecurrenceArgs']]):
        pulumi.set(self, "weekly_recurrence", value)


class BackupSchedule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_recurrence: Optional[pulumi.Input[pulumi.InputType['BackupScheduleDailyRecurrenceArgs']]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['BackupScheduleWeeklyRecurrenceArgs']]] = None,
                 __props__=None):
        """
        A backup schedule for a Cloud Firestore Database.
        This resource is owned by the database it is backing up, and is deleted along with the database.
        The actual backups are not though.

        To get more information about BackupSchedule, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.backupSchedules)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/backups)

        > **Warning:** This resource creates a Firestore Backup Schedule on a project that already has
        a Firestore database.
        This resource is owned by the database it is backing up, and is deleted along
        with the database. The actual backups are not though.

        ## Example Usage
        ### Firestore Backup Schedule Daily

        ```python
        import pulumi
        import pulumi_gcp as gcp

        database = gcp.firestore.Database("database",
            project="my-project-name",
            location_id="nam5",
            type="FIRESTORE_NATIVE",
            delete_protection_state="DELETE_PROTECTION_ENABLED",
            deletion_policy="DELETE")
        daily_backup = gcp.firestore.BackupSchedule("daily-backup",
            project="my-project-name",
            database=database.name,
            retention="604800s",
            daily_recurrence=gcp.firestore.BackupScheduleDailyRecurrenceArgs())
        ```
        ### Firestore Backup Schedule Weekly

        ```python
        import pulumi
        import pulumi_gcp as gcp

        database = gcp.firestore.Database("database",
            project="my-project-name",
            location_id="nam5",
            type="FIRESTORE_NATIVE",
            delete_protection_state="DELETE_PROTECTION_ENABLED",
            deletion_policy="DELETE")
        weekly_backup = gcp.firestore.BackupSchedule("weekly-backup",
            project="my-project-name",
            database=database.name,
            retention="8467200s",
            weekly_recurrence=gcp.firestore.BackupScheduleWeeklyRecurrenceArgs(
                day="SUNDAY",
            ))
        ```

        ## Import

        BackupSchedule can be imported using any of these accepted formats* `projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}` * `{{project}}/{{database}}/{{name}}` * `{{database}}/{{name}}` When using the `pulumi import` command, BackupSchedule can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}
        ```

        ```sh
         $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{project}}/{{database}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{database}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BackupScheduleDailyRecurrenceArgs']] daily_recurrence: For a schedule that runs daily at a specified time.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] retention: At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
               A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
               For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
               
               
               - - -
        :param pulumi.Input[pulumi.InputType['BackupScheduleWeeklyRecurrenceArgs']] weekly_recurrence: For a schedule that runs weekly on a specific day and time.
               Structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackupScheduleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A backup schedule for a Cloud Firestore Database.
        This resource is owned by the database it is backing up, and is deleted along with the database.
        The actual backups are not though.

        To get more information about BackupSchedule, see:

        * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.backupSchedules)
        * How-to Guides
            * [Official Documentation](https://cloud.google.com/firestore/docs/backups)

        > **Warning:** This resource creates a Firestore Backup Schedule on a project that already has
        a Firestore database.
        This resource is owned by the database it is backing up, and is deleted along
        with the database. The actual backups are not though.

        ## Example Usage
        ### Firestore Backup Schedule Daily

        ```python
        import pulumi
        import pulumi_gcp as gcp

        database = gcp.firestore.Database("database",
            project="my-project-name",
            location_id="nam5",
            type="FIRESTORE_NATIVE",
            delete_protection_state="DELETE_PROTECTION_ENABLED",
            deletion_policy="DELETE")
        daily_backup = gcp.firestore.BackupSchedule("daily-backup",
            project="my-project-name",
            database=database.name,
            retention="604800s",
            daily_recurrence=gcp.firestore.BackupScheduleDailyRecurrenceArgs())
        ```
        ### Firestore Backup Schedule Weekly

        ```python
        import pulumi
        import pulumi_gcp as gcp

        database = gcp.firestore.Database("database",
            project="my-project-name",
            location_id="nam5",
            type="FIRESTORE_NATIVE",
            delete_protection_state="DELETE_PROTECTION_ENABLED",
            deletion_policy="DELETE")
        weekly_backup = gcp.firestore.BackupSchedule("weekly-backup",
            project="my-project-name",
            database=database.name,
            retention="8467200s",
            weekly_recurrence=gcp.firestore.BackupScheduleWeeklyRecurrenceArgs(
                day="SUNDAY",
            ))
        ```

        ## Import

        BackupSchedule can be imported using any of these accepted formats* `projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}` * `{{project}}/{{database}}/{{name}}` * `{{database}}/{{name}}` When using the `pulumi import` command, BackupSchedule can be imported using one of the formats above. For example

        ```sh
         $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default projects/{{project}}/databases/{{database}}/backupSchedules/{{name}}
        ```

        ```sh
         $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{project}}/{{database}}/{{name}}
        ```

        ```sh
         $ pulumi import gcp:firestore/backupSchedule:BackupSchedule default {{database}}/{{name}}
        ```

        :param str resource_name: The name of the resource.
        :param BackupScheduleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackupScheduleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 daily_recurrence: Optional[pulumi.Input[pulumi.InputType['BackupScheduleDailyRecurrenceArgs']]] = None,
                 database: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 retention: Optional[pulumi.Input[str]] = None,
                 weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['BackupScheduleWeeklyRecurrenceArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackupScheduleArgs.__new__(BackupScheduleArgs)

            __props__.__dict__["daily_recurrence"] = daily_recurrence
            __props__.__dict__["database"] = database
            __props__.__dict__["project"] = project
            if retention is None and not opts.urn:
                raise TypeError("Missing required property 'retention'")
            __props__.__dict__["retention"] = retention
            __props__.__dict__["weekly_recurrence"] = weekly_recurrence
            __props__.__dict__["name"] = None
        super(BackupSchedule, __self__).__init__(
            'gcp:firestore/backupSchedule:BackupSchedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            daily_recurrence: Optional[pulumi.Input[pulumi.InputType['BackupScheduleDailyRecurrenceArgs']]] = None,
            database: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None,
            retention: Optional[pulumi.Input[str]] = None,
            weekly_recurrence: Optional[pulumi.Input[pulumi.InputType['BackupScheduleWeeklyRecurrenceArgs']]] = None) -> 'BackupSchedule':
        """
        Get an existing BackupSchedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['BackupScheduleDailyRecurrenceArgs']] daily_recurrence: For a schedule that runs daily at a specified time.
        :param pulumi.Input[str] database: The Firestore database id. Defaults to `"(default)"`.
        :param pulumi.Input[str] name: The unique backup schedule identifier across all locations and databases for the given project. Format:
               `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
        :param pulumi.Input[str] project: The ID of the project in which the resource belongs.
               If it is not provided, the provider project is used.
        :param pulumi.Input[str] retention: At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
               A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
               For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.
               
               
               - - -
        :param pulumi.Input[pulumi.InputType['BackupScheduleWeeklyRecurrenceArgs']] weekly_recurrence: For a schedule that runs weekly on a specific day and time.
               Structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackupScheduleState.__new__(_BackupScheduleState)

        __props__.__dict__["daily_recurrence"] = daily_recurrence
        __props__.__dict__["database"] = database
        __props__.__dict__["name"] = name
        __props__.__dict__["project"] = project
        __props__.__dict__["retention"] = retention
        __props__.__dict__["weekly_recurrence"] = weekly_recurrence
        return BackupSchedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dailyRecurrence")
    def daily_recurrence(self) -> pulumi.Output[Optional['outputs.BackupScheduleDailyRecurrence']]:
        """
        For a schedule that runs daily at a specified time.
        """
        return pulumi.get(self, "daily_recurrence")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[Optional[str]]:
        """
        The Firestore database id. Defaults to `"(default)"`.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique backup schedule identifier across all locations and databases for the given project. Format:
        `projects/{{project}}/databases/{{database}}/backupSchedules/{{backupSchedule}}`
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
    def retention(self) -> pulumi.Output[str]:
        """
        At what relative time in the future, compared to its creation time, the backup should be deleted, e.g. keep backups for 7 days.
        A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
        For a daily backup recurrence, set this to a value up to 7 days. If you set a weekly backup recurrence, set this to a value up to 14 weeks.


        - - -
        """
        return pulumi.get(self, "retention")

    @property
    @pulumi.getter(name="weeklyRecurrence")
    def weekly_recurrence(self) -> pulumi.Output[Optional['outputs.BackupScheduleWeeklyRecurrence']]:
        """
        For a schedule that runs weekly on a specific day and time.
        Structure is documented below.
        """
        return pulumi.get(self, "weekly_recurrence")

