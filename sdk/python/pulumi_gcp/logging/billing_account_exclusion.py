# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class BillingAccountExclusion(pulumi.CustomResource):
    """
    Manages a billing account logging exclusion. For more information see
    [the official documentation](https://cloud.google.com/logging/docs/) and
    [Excluding Logs](https://cloud.google.com/logging/docs/exclusions).
    
    Note that you must have the "Logs Configuration Writer" IAM role (`roles/logging.configWriter`)
    granted to the credentials used with Terraform.
    """
    def __init__(__self__, __name__, __opts__=None, billing_account=None, description=None, disabled=None, filter=None, name=None):
        """Create a BillingAccountExclusion resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not billing_account:
            raise TypeError('Missing required property billing_account')
        elif not isinstance(billing_account, basestring):
            raise TypeError('Expected property billing_account to be a basestring')
        __self__.billing_account = billing_account
        """
        The billing account to create the exclusion for.
        """
        __props__['billingAccount'] = billing_account

        if description and not isinstance(description, basestring):
            raise TypeError('Expected property description to be a basestring')
        __self__.description = description
        """
        A human-readable description.
        """
        __props__['description'] = description

        if disabled and not isinstance(disabled, bool):
            raise TypeError('Expected property disabled to be a bool')
        __self__.disabled = disabled
        """
        Whether this exclusion rule should be disabled or not. This defaults to
        false.
        """
        __props__['disabled'] = disabled

        if not filter:
            raise TypeError('Missing required property filter')
        elif not isinstance(filter, basestring):
            raise TypeError('Expected property filter to be a basestring')
        __self__.filter = filter
        """
        The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        write a filter.
        """
        __props__['filter'] = filter

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the logging exclusion.
        """
        __props__['name'] = name

        super(BillingAccountExclusion, __self__).__init__(
            'gcp:logging/billingAccountExclusion:BillingAccountExclusion',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'billingAccount' in outs:
            self.billing_account = outs['billingAccount']
        if 'description' in outs:
            self.description = outs['description']
        if 'disabled' in outs:
            self.disabled = outs['disabled']
        if 'filter' in outs:
            self.filter = outs['filter']
        if 'name' in outs:
            self.name = outs['name']
