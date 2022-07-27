// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.organizations;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.organizations.IamAuditConfigArgs;
import com.pulumi.gcp.organizations.inputs.IamAuditConfigState;
import com.pulumi.gcp.organizations.outputs.IamAuditConfigAuditLogConfig;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Allows management of audit logging config for a given service for a Google Cloud Platform Organization.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.organizations.IamAuditConfig;
 * import com.pulumi.gcp.organizations.IamAuditConfigArgs;
 * import com.pulumi.gcp.organizations.inputs.IamAuditConfigAuditLogConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var config = new IamAuditConfig(&#34;config&#34;, IamAuditConfigArgs.builder()        
 *             .auditLogConfigs(IamAuditConfigAuditLogConfigArgs.builder()
 *                 .exemptedMembers(&#34;user:joebloggs@hashicorp.com&#34;)
 *                 .logType(&#34;DATA_READ&#34;)
 *                 .build())
 *             .orgId(&#34;your-organization-id&#34;)
 *             .service(&#34;allServices&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * IAM audit config imports use the identifier of the resource in question and the service, e.g.
 * 
 * ```sh
 *  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig config &#34;your-organization-id foo.googleapis.com&#34;
 * ```
 * 
 */
@ResourceType(type="gcp:organizations/iamAuditConfig:IamAuditConfig")
public class IamAuditConfig extends com.pulumi.resources.CustomResource {
    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     * 
     */
    @Export(name="auditLogConfigs", type=List.class, parameters={IamAuditConfigAuditLogConfig.class})
    private Output<List<IamAuditConfigAuditLogConfig>> auditLogConfigs;

    /**
     * @return The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     * 
     */
    public Output<List<IamAuditConfigAuditLogConfig>> auditLogConfigs() {
        return this.auditLogConfigs;
    }
    /**
     * The etag of iam policy
     * 
     */
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    /**
     * @return The etag of iam policy
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The numeric ID of the organization in which you want to manage the audit logging config.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output<String> orgId;

    /**
     * @return The numeric ID of the organization in which you want to manage the audit logging config.
     * 
     */
    public Output<String> orgId() {
        return this.orgId;
    }
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
     * 
     */
    @Export(name="service", type=String.class, parameters={})
    private Output<String> service;

    /**
     * @return Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `log_types` specified in each `audit_log_config` are enabled, and the `exempted_members` in each `audit_log_config` are exempted.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IamAuditConfig(String name) {
        this(name, IamAuditConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IamAuditConfig(String name, IamAuditConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IamAuditConfig(String name, IamAuditConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:organizations/iamAuditConfig:IamAuditConfig", name, args == null ? IamAuditConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IamAuditConfig(String name, Output<String> id, @Nullable IamAuditConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:organizations/iamAuditConfig:IamAuditConfig", name, state, makeResourceOptions(options, id));
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
    public static IamAuditConfig get(String name, Output<String> id, @Nullable IamAuditConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IamAuditConfig(name, id, state, options);
    }
}
