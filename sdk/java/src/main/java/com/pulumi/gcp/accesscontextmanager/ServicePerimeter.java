// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.accesscontextmanager;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.accesscontextmanager.ServicePerimeterArgs;
import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterState;
import com.pulumi.gcp.accesscontextmanager.outputs.ServicePerimeterSpec;
import com.pulumi.gcp.accesscontextmanager.outputs.ServicePerimeterStatus;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ServicePerimeter describes a set of GCP resources which can freely import
 * and export data amongst themselves, but not export outside of the
 * ServicePerimeter. If a request with a source within this ServicePerimeter
 * has a target outside of the ServicePerimeter, the request will be blocked.
 * Otherwise the request is allowed. There are two types of Service Perimeter
 * - Regular and Bridge. Regular Service Perimeters cannot overlap, a single
 *   GCP project can only belong to a single regular Service Perimeter. Service
 *   Perimeter Bridges can contain only GCP projects as members, a single GCP
 *   project may belong to multiple Service Perimeter Bridges.
 * 
 * To get more information about ServicePerimeter, see:
 * 
 * * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.servicePerimeters)
 * * How-to Guides
 *     * [Service Perimeter Quickstart](https://cloud.google.com/vpc-service-controls/docs/quickstart)
 * 
 * &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
 * you must specify a `billing_project` and set `user_project_override` to true
 * in the provider configuration. Otherwise the ACM API will return a 403 error.
 * Your account must have the `serviceusage.services.use` permission on the
 * `billing_project` you defined.
 * 
 * ## Example Usage
 * 
 * ### Access Context Manager Service Perimeter Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicy;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicyArgs;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeter;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeterArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterStatusArgs;
 * import com.pulumi.gcp.accesscontextmanager.AccessLevel;
 * import com.pulumi.gcp.accesscontextmanager.AccessLevelArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.AccessLevelBasicArgs;
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
 *         var access_policy = new AccessPolicy("access-policy", AccessPolicyArgs.builder()
 *             .parent("organizations/123456789")
 *             .title("my policy")
 *             .build());
 * 
 *         var service_perimeter = new ServicePerimeter("service-perimeter", ServicePerimeterArgs.builder()
 *             .parent(access_policy.name().applyValue(name -> String.format("accessPolicies/%s", name)))
 *             .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/servicePerimeters/restrict_storage", name)))
 *             .title("restrict_storage")
 *             .status(ServicePerimeterStatusArgs.builder()
 *                 .restrictedServices("storage.googleapis.com")
 *                 .build())
 *             .build());
 * 
 *         var access_level = new AccessLevel("access-level", AccessLevelArgs.builder()
 *             .parent(access_policy.name().applyValue(name -> String.format("accessPolicies/%s", name)))
 *             .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/accessLevels/chromeos_no_lock", name)))
 *             .title("chromeos_no_lock")
 *             .basic(AccessLevelBasicArgs.builder()
 *                 .conditions(AccessLevelBasicConditionArgs.builder()
 *                     .devicePolicy(AccessLevelBasicConditionDevicePolicyArgs.builder()
 *                         .requireScreenLock(false)
 *                         .osConstraints(AccessLevelBasicConditionDevicePolicyOsConstraintArgs.builder()
 *                             .osType("DESKTOP_CHROME_OS")
 *                             .build())
 *                         .build())
 *                     .regions(                    
 *                         "CH",
 *                         "IT",
 *                         "US")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Access Context Manager Service Perimeter Secure Data Exchange
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicy;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicyArgs;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeters;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimetersArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimetersServicePerimeterArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimetersServicePerimeterStatusArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimetersServicePerimeterStatusVpcAccessibleServicesArgs;
 * import com.pulumi.gcp.accesscontextmanager.AccessLevel;
 * import com.pulumi.gcp.accesscontextmanager.AccessLevelArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.AccessLevelBasicArgs;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeter;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeterArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterStatusArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterStatusVpcAccessibleServicesArgs;
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
 *         var access_policy = new AccessPolicy("access-policy", AccessPolicyArgs.builder()
 *             .parent("organizations/123456789")
 *             .title("my policy")
 *             .build());
 * 
 *         var secure_data_exchange = new ServicePerimeters("secure-data-exchange", ServicePerimetersArgs.builder()
 *             .parent(access_policy.name().applyValue(name -> String.format("accessPolicies/%s", name)))
 *             .servicePerimeters(            
 *                 ServicePerimetersServicePerimeterArgs.builder()
 *                     .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/servicePerimeters/", name)))
 *                     .title("")
 *                     .status(ServicePerimetersServicePerimeterStatusArgs.builder()
 *                         .restrictedServices("storage.googleapis.com")
 *                         .build())
 *                     .build(),
 *                 ServicePerimetersServicePerimeterArgs.builder()
 *                     .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/servicePerimeters/", name)))
 *                     .title("")
 *                     .status(ServicePerimetersServicePerimeterStatusArgs.builder()
 *                         .restrictedServices("bigtable.googleapis.com")
 *                         .vpcAccessibleServices(ServicePerimetersServicePerimeterStatusVpcAccessibleServicesArgs.builder()
 *                             .enableRestriction(true)
 *                             .allowedServices("bigquery.googleapis.com")
 *                             .build())
 *                         .build())
 *                     .build())
 *             .build());
 * 
 *         var access_level = new AccessLevel("access-level", AccessLevelArgs.builder()
 *             .parent(access_policy.name().applyValue(name -> String.format("accessPolicies/%s", name)))
 *             .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/accessLevels/secure_data_exchange", name)))
 *             .title("secure_data_exchange")
 *             .basic(AccessLevelBasicArgs.builder()
 *                 .conditions(AccessLevelBasicConditionArgs.builder()
 *                     .devicePolicy(AccessLevelBasicConditionDevicePolicyArgs.builder()
 *                         .requireScreenLock(false)
 *                         .osConstraints(AccessLevelBasicConditionDevicePolicyOsConstraintArgs.builder()
 *                             .osType("DESKTOP_CHROME_OS")
 *                             .build())
 *                         .build())
 *                     .regions(                    
 *                         "CH",
 *                         "IT",
 *                         "US")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var test_access = new ServicePerimeter("test-access", ServicePerimeterArgs.builder()
 *             .parent(String.format("accessPolicies/%s", test_accessGoogleAccessContextManagerAccessPolicy.name()))
 *             .name(String.format("accessPolicies/%s/servicePerimeters/%s", test_accessGoogleAccessContextManagerAccessPolicy.name()))
 *             .title("%s")
 *             .perimeterType("PERIMETER_TYPE_REGULAR")
 *             .status(ServicePerimeterStatusArgs.builder()
 *                 .restrictedServices(                
 *                     "bigquery.googleapis.com",
 *                     "storage.googleapis.com")
 *                 .accessLevels(access_level.name())
 *                 .vpcAccessibleServices(ServicePerimeterStatusVpcAccessibleServicesArgs.builder()
 *                     .enableRestriction(true)
 *                     .allowedServices(                    
 *                         "bigquery.googleapis.com",
 *                         "storage.googleapis.com")
 *                     .build())
 *                 .ingressPolicies(ServicePerimeterStatusIngressPolicyArgs.builder()
 *                     .ingressFrom(ServicePerimeterStatusIngressPolicyIngressFromArgs.builder()
 *                         .sources(ServicePerimeterStatusIngressPolicyIngressFromSourceArgs.builder()
 *                             .accessLevel(test_accessGoogleAccessContextManagerAccessLevel.name())
 *                             .build())
 *                         .identityType("ANY_IDENTITY")
 *                         .build())
 *                     .ingressTo(ServicePerimeterStatusIngressPolicyIngressToArgs.builder()
 *                         .resources("*")
 *                         .operations(                        
 *                             ServicePerimeterStatusIngressPolicyIngressToOperationArgs.builder()
 *                                 .serviceName("bigquery.googleapis.com")
 *                                 .methodSelectors(                                
 *                                     ServicePerimeterStatusIngressPolicyIngressToOperationMethodSelectorArgs.builder()
 *                                         .method("BigQueryStorage.ReadRows")
 *                                         .build(),
 *                                     ServicePerimeterStatusIngressPolicyIngressToOperationMethodSelectorArgs.builder()
 *                                         .method("TableService.ListTables")
 *                                         .build(),
 *                                     ServicePerimeterStatusIngressPolicyIngressToOperationMethodSelectorArgs.builder()
 *                                         .permission("bigquery.jobs.get")
 *                                         .build())
 *                                 .build(),
 *                             ServicePerimeterStatusIngressPolicyIngressToOperationArgs.builder()
 *                                 .serviceName("storage.googleapis.com")
 *                                 .methodSelectors(ServicePerimeterStatusIngressPolicyIngressToOperationMethodSelectorArgs.builder()
 *                                     .method("google.storage.objects.create")
 *                                     .build())
 *                                 .build())
 *                         .build())
 *                     .build())
 *                 .egressPolicies(ServicePerimeterStatusEgressPolicyArgs.builder()
 *                     .egressFrom(ServicePerimeterStatusEgressPolicyEgressFromArgs.builder()
 *                         .identityType("ANY_USER_ACCOUNT")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Access Context Manager Service Perimeter Dry-Run
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicy;
 * import com.pulumi.gcp.accesscontextmanager.AccessPolicyArgs;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeter;
 * import com.pulumi.gcp.accesscontextmanager.ServicePerimeterArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterStatusArgs;
 * import com.pulumi.gcp.accesscontextmanager.inputs.ServicePerimeterSpecArgs;
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
 *         var access_policy = new AccessPolicy("access-policy", AccessPolicyArgs.builder()
 *             .parent("organizations/123456789")
 *             .title("my policy")
 *             .build());
 * 
 *         var service_perimeter = new ServicePerimeter("service-perimeter", ServicePerimeterArgs.builder()
 *             .parent(access_policy.name().applyValue(name -> String.format("accessPolicies/%s", name)))
 *             .name(access_policy.name().applyValue(name -> String.format("accessPolicies/%s/servicePerimeters/restrict_bigquery_dryrun_storage", name)))
 *             .title("restrict_bigquery_dryrun_storage")
 *             .status(ServicePerimeterStatusArgs.builder()
 *                 .restrictedServices("bigquery.googleapis.com")
 *                 .build())
 *             .spec(ServicePerimeterSpecArgs.builder()
 *                 .restrictedServices("storage.googleapis.com")
 *                 .build())
 *             .useExplicitDryRunSpec(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ServicePerimeter can be imported using any of these accepted formats:
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, ServicePerimeter can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:accesscontextmanager/servicePerimeter:ServicePerimeter default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:accesscontextmanager/servicePerimeter:ServicePerimeter")
public class ServicePerimeter extends com.pulumi.resources.CustomResource {
    /**
     * Time the AccessPolicy was created in UTC.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time the AccessPolicy was created in UTC.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Description of the ServicePerimeter and its use. Does not affect
     * behavior.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the ServicePerimeter and its use. Does not affect
     * behavior.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Resource name for the ServicePerimeter. The short_name component must
     * begin with a letter and only include alphanumeric and &#39;_&#39;.
     * Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Resource name for the ServicePerimeter. The short_name component must
     * begin with a letter and only include alphanumeric and &#39;_&#39;.
     * Format: accessPolicies/{policy_id}/servicePerimeters/{short_name}
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The AccessPolicy this ServicePerimeter lives in.
     * Format: accessPolicies/{policy_id}
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * Specifies the type of the Perimeter. There are two types: regular and
     * bridge. Regular Service Perimeter contains resources, access levels,
     * and restricted services. Every resource can be in at most
     * ONE regular Service Perimeter.
     * In addition to being in a regular service perimeter, a resource can also
     * be in zero or more perimeter bridges. A perimeter bridge only contains
     * resources. Cross project operations are permitted if all effected
     * resources share some perimeter (whether bridge or regular). Perimeter
     * Bridge does not contain access levels or services: those are governed
     * entirely by the regular perimeter that resource is in.
     * Perimeter Bridges are typically useful when building more complex
     * topologies with many independent perimeters that need to share some data
     * with a common perimeter, but should not be able to share data among
     * themselves.
     * Default value is `PERIMETER_TYPE_REGULAR`.
     * Possible values are: `PERIMETER_TYPE_REGULAR`, `PERIMETER_TYPE_BRIDGE`.
     * 
     */
    @Export(name="perimeterType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> perimeterType;

    /**
     * @return Specifies the type of the Perimeter. There are two types: regular and
     * bridge. Regular Service Perimeter contains resources, access levels,
     * and restricted services. Every resource can be in at most
     * ONE regular Service Perimeter.
     * In addition to being in a regular service perimeter, a resource can also
     * be in zero or more perimeter bridges. A perimeter bridge only contains
     * resources. Cross project operations are permitted if all effected
     * resources share some perimeter (whether bridge or regular). Perimeter
     * Bridge does not contain access levels or services: those are governed
     * entirely by the regular perimeter that resource is in.
     * Perimeter Bridges are typically useful when building more complex
     * topologies with many independent perimeters that need to share some data
     * with a common perimeter, but should not be able to share data among
     * themselves.
     * Default value is `PERIMETER_TYPE_REGULAR`.
     * Possible values are: `PERIMETER_TYPE_REGULAR`, `PERIMETER_TYPE_BRIDGE`.
     * 
     */
    public Output<Optional<String>> perimeterType() {
        return Codegen.optional(this.perimeterType);
    }
    /**
     * Proposed (or dry run) ServicePerimeter configuration.
     * This configuration allows to specify and test ServicePerimeter configuration
     * without enforcing actual access restrictions. Only allowed to be set when
     * the `useExplicitDryRunSpec` flag is set.
     * Structure is documented below.
     * 
     */
    @Export(name="spec", refs={ServicePerimeterSpec.class}, tree="[0]")
    private Output</* @Nullable */ ServicePerimeterSpec> spec;

    /**
     * @return Proposed (or dry run) ServicePerimeter configuration.
     * This configuration allows to specify and test ServicePerimeter configuration
     * without enforcing actual access restrictions. Only allowed to be set when
     * the `useExplicitDryRunSpec` flag is set.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ServicePerimeterSpec>> spec() {
        return Codegen.optional(this.spec);
    }
    /**
     * ServicePerimeter configuration. Specifies sets of resources,
     * restricted services and access levels that determine
     * perimeter content and boundaries.
     * Structure is documented below.
     * 
     */
    @Export(name="status", refs={ServicePerimeterStatus.class}, tree="[0]")
    private Output</* @Nullable */ ServicePerimeterStatus> status;

    /**
     * @return ServicePerimeter configuration. Specifies sets of resources,
     * restricted services and access levels that determine
     * perimeter content and boundaries.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ServicePerimeterStatus>> status() {
        return Codegen.optional(this.status);
    }
    /**
     * Human readable title. Must be unique within the Policy.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return Human readable title. Must be unique within the Policy.
     * 
     */
    public Output<String> title() {
        return this.title;
    }
    /**
     * Time the AccessPolicy was updated in UTC.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time the AccessPolicy was updated in UTC.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
     * for all Service Perimeters, and that spec is identical to the status for those
     * Service Perimeters. When this flag is set, it inhibits the generation of the
     * implicit spec, thereby allowing the user to explicitly provide a
     * configuration (&#34;spec&#34;) to use in a dry-run version of the Service Perimeter.
     * This allows the user to test changes to the enforced config (&#34;status&#34;) without
     * actually enforcing them. This testing is done through analyzing the differences
     * between currently enforced and suggested restrictions. useExplicitDryRunSpec must
     * bet set to True if any of the fields in the spec are set to non-default values.
     * 
     */
    @Export(name="useExplicitDryRunSpec", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useExplicitDryRunSpec;

    /**
     * @return Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
     * for all Service Perimeters, and that spec is identical to the status for those
     * Service Perimeters. When this flag is set, it inhibits the generation of the
     * implicit spec, thereby allowing the user to explicitly provide a
     * configuration (&#34;spec&#34;) to use in a dry-run version of the Service Perimeter.
     * This allows the user to test changes to the enforced config (&#34;status&#34;) without
     * actually enforcing them. This testing is done through analyzing the differences
     * between currently enforced and suggested restrictions. useExplicitDryRunSpec must
     * bet set to True if any of the fields in the spec are set to non-default values.
     * 
     */
    public Output<Optional<Boolean>> useExplicitDryRunSpec() {
        return Codegen.optional(this.useExplicitDryRunSpec);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServicePerimeter(java.lang.String name) {
        this(name, ServicePerimeterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServicePerimeter(java.lang.String name, ServicePerimeterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServicePerimeter(java.lang.String name, ServicePerimeterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:accesscontextmanager/servicePerimeter:ServicePerimeter", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServicePerimeter(java.lang.String name, Output<java.lang.String> id, @Nullable ServicePerimeterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:accesscontextmanager/servicePerimeter:ServicePerimeter", name, state, makeResourceOptions(options, id), false);
    }

    private static ServicePerimeterArgs makeArgs(ServicePerimeterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServicePerimeterArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static ServicePerimeter get(java.lang.String name, Output<java.lang.String> id, @Nullable ServicePerimeterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServicePerimeter(name, id, state, options);
    }
}
