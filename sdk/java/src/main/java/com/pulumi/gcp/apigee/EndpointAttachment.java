// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigee;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigee.EndpointAttachmentArgs;
import com.pulumi.gcp.apigee.inputs.EndpointAttachmentState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Apigee Endpoint Attachment.
 * 
 * To get more information about EndpointAttachment, see:
 * 
 * * [API documentation](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest/v1/organizations.endpointAttachments/create)
 * * How-to Guides
 *     * [Creating an environment](https://cloud.google.com/apigee/docs/api-platform/get-started/create-environment)
 * 
 * ## Example Usage
 * 
 * ## Import
 * 
 * EndpointAttachment can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/endpointAttachment:EndpointAttachment default {{org_id}}/endpointAttachments/{{endpoint_attachment_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigee/endpointAttachment:EndpointAttachment default {{org_id}}/{{endpoint_attachment_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigee/endpointAttachment:EndpointAttachment")
public class EndpointAttachment extends com.pulumi.resources.CustomResource {
    /**
     * State of the endpoint attachment connection to the service attachment.
     * 
     */
    @Export(name="connectionState", type=String.class, parameters={})
    private Output<String> connectionState;

    /**
     * @return State of the endpoint attachment connection to the service attachment.
     * 
     */
    public Output<String> connectionState() {
        return this.connectionState;
    }
    /**
     * ID of the endpoint attachment.
     * 
     */
    @Export(name="endpointAttachmentId", type=String.class, parameters={})
    private Output<String> endpointAttachmentId;

    /**
     * @return ID of the endpoint attachment.
     * 
     */
    public Output<String> endpointAttachmentId() {
        return this.endpointAttachmentId;
    }
    /**
     * Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
     * 
     */
    @Export(name="host", type=String.class, parameters={})
    private Output<String> host;

    /**
     * @return Host that can be used in either HTTP Target Endpoint directly, or as the host in Target Server.
     * 
     */
    public Output<String> host() {
        return this.host;
    }
    /**
     * Location of the endpoint attachment.
     * 
     */
    @Export(name="location", type=String.class, parameters={})
    private Output<String> location;

    /**
     * @return Location of the endpoint attachment.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Name of the Endpoint Attachment in the following format:
     * organizations/{organization}/endpointAttachments/{endpointAttachment}.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the Endpoint Attachment in the following format:
     * organizations/{organization}/endpointAttachments/{endpointAttachment}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The Apigee Organization associated with the Apigee instance,
     * in the format `organizations/{{org_name}}`.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    /**
     * @return The Apigee Organization associated with the Apigee instance,
     * in the format `organizations/{{org_name}}`.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * Format: projects/*{@literal /}regions/*{@literal /}serviceAttachments/*
     * 
     */
    @Export(name="serviceAttachment", type=String.class, parameters={})
    private Output<String> serviceAttachment;

    /**
     * @return Format: projects/*{@literal /}regions/*{@literal /}serviceAttachments/*
     * 
     */
    public Output<String> serviceAttachment() {
        return this.serviceAttachment;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointAttachment(String name) {
        this(name, EndpointAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointAttachment(String name, EndpointAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointAttachment(String name, EndpointAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/endpointAttachment:EndpointAttachment", name, args == null ? EndpointAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EndpointAttachment(String name, Output<String> id, @Nullable EndpointAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigee/endpointAttachment:EndpointAttachment", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static EndpointAttachment get(String name, Output<String> id, @Nullable EndpointAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointAttachment(name, id, state, options);
    }
}
