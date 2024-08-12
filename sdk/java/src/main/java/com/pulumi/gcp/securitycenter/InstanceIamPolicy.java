// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.securitycenter;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.securitycenter.InstanceIamPolicyArgs;
import com.pulumi.gcp.securitycenter.inputs.InstanceIamPolicyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Represents a Data Fusion instance.
 * 
 * To get more information about Instance, see:
 * 
 * * [API documentation](https://cloud.google.com/data-fusion/docs/reference/rest/v1beta1/projects.locations.instances)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/data-fusion/docs/)
 * 
 * ## Example Usage
 * 
 * ### Data Fusion Instance Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
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
 *         var basicInstance = new Instance("basicInstance", InstanceArgs.builder()
 *             .name("my-instance")
 *             .region("us-central1")
 *             .type("BASIC")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Fusion Instance Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.appengine.AppengineFunctions;
 * import com.pulumi.gcp.appengine.inputs.GetDefaultServiceAccountArgs;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
 * import com.pulumi.gcp.datafusion.inputs.InstanceNetworkConfigArgs;
 * import com.pulumi.gcp.datafusion.inputs.InstanceAcceleratorArgs;
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
 *         final var default = AppengineFunctions.getDefaultServiceAccount();
 * 
 *         var network = new Network("network", NetworkArgs.builder()
 *             .name("datafusion-full-network")
 *             .build());
 * 
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("datafusion-ip-alloc")
 *             .addressType("INTERNAL")
 *             .purpose("VPC_PEERING")
 *             .prefixLength(22)
 *             .network(network.id())
 *             .build());
 * 
 *         var extendedInstance = new Instance("extendedInstance", InstanceArgs.builder()
 *             .name("my-instance")
 *             .description("My Data Fusion instance")
 *             .displayName("My Data Fusion instance")
 *             .region("us-central1")
 *             .type("BASIC")
 *             .enableStackdriverLogging(true)
 *             .enableStackdriverMonitoring(true)
 *             .privateInstance(true)
 *             .dataprocServiceAccount(default_.email())
 *             .labels(Map.of("example_key", "example_value"))
 *             .networkConfig(InstanceNetworkConfigArgs.builder()
 *                 .network("default")
 *                 .ipAllocation(Output.tuple(privateIpAlloc.address(), privateIpAlloc.prefixLength()).applyValue(values -> {
 *                     var address = values.t1;
 *                     var prefixLength = values.t2;
 *                     return String.format("%s/%s", address,prefixLength);
 *                 }))
 *                 .build())
 *             .accelerators(InstanceAcceleratorArgs.builder()
 *                 .acceleratorType("CDC")
 *                 .state("ENABLED")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Fusion Instance Psc
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.NetworkAttachment;
 * import com.pulumi.gcp.compute.NetworkAttachmentArgs;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
 * import com.pulumi.gcp.datafusion.inputs.InstanceNetworkConfigArgs;
 * import com.pulumi.gcp.datafusion.inputs.InstanceNetworkConfigPrivateServiceConnectConfigArgs;
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
 *         var psc = new Network("psc", NetworkArgs.builder()
 *             .name("datafusion-psc-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var pscSubnetwork = new Subnetwork("pscSubnetwork", SubnetworkArgs.builder()
 *             .name("datafusion-psc-subnet")
 *             .region("us-central1")
 *             .network(psc.id())
 *             .ipCidrRange("10.0.0.0/16")
 *             .build());
 * 
 *         var pscNetworkAttachment = new NetworkAttachment("pscNetworkAttachment", NetworkAttachmentArgs.builder()
 *             .name("datafusion-psc-attachment")
 *             .region("us-central1")
 *             .connectionPreference("ACCEPT_AUTOMATIC")
 *             .subnetworks(pscSubnetwork.selfLink())
 *             .build());
 * 
 *         var pscInstance = new Instance("pscInstance", InstanceArgs.builder()
 *             .name("psc-instance")
 *             .region("us-central1")
 *             .type("BASIC")
 *             .privateInstance(true)
 *             .networkConfig(InstanceNetworkConfigArgs.builder()
 *                 .connectionType("PRIVATE_SERVICE_CONNECT_INTERFACES")
 *                 .privateServiceConnectConfig(InstanceNetworkConfigPrivateServiceConnectConfigArgs.builder()
 *                     .networkAttachment(pscNetworkAttachment.id())
 *                     .unreachableCidrBlock("192.168.0.0/25")
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Fusion Instance Cmek
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.kms.KeyRing;
 * import com.pulumi.gcp.kms.KeyRingArgs;
 * import com.pulumi.gcp.kms.CryptoKey;
 * import com.pulumi.gcp.kms.CryptoKeyArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMember;
 * import com.pulumi.gcp.kms.CryptoKeyIAMMemberArgs;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
 * import com.pulumi.gcp.datafusion.inputs.InstanceCryptoKeyConfigArgs;
 * import com.pulumi.resources.CustomResourceOptions;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App }{{@code
 *     public static void main(String[] args) }{{@code
 *         Pulumi.run(App::stack);
 *     }}{@code
 * 
 *     public static void stack(Context ctx) }{{@code
 *         var keyRing = new KeyRing("keyRing", KeyRingArgs.builder()
 *             .name("my-instance")
 *             .location("us-central1")
 *             .build());
 * 
 *         var cryptoKey = new CryptoKey("cryptoKey", CryptoKeyArgs.builder()
 *             .name("my-instance")
 *             .keyRing(keyRing.id())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var cryptoKeyMember = new CryptoKeyIAMMember("cryptoKeyMember", CryptoKeyIAMMemberArgs.builder()
 *             .cryptoKeyId(cryptoKey.id())
 *             .role("roles/cloudkms.cryptoKeyEncrypterDecrypter")
 *             .member(String.format("serviceAccount:service-%s}{@literal @}{@code gcp-sa-datafusion.iam.gserviceaccount.com", project.applyValue(getProjectResult -> getProjectResult.number())))
 *             .build());
 * 
 *         var cmek = new Instance("cmek", InstanceArgs.builder()
 *             .name("my-instance")
 *             .region("us-central1")
 *             .type("BASIC")
 *             .cryptoKeyConfig(InstanceCryptoKeyConfigArgs.builder()
 *                 .keyReference(cryptoKey.id())
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(cryptoKeyMember)
 *                 .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Fusion Instance Enterprise
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
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
 *         var enterpriseInstance = new Instance("enterpriseInstance", InstanceArgs.builder()
 *             .name("my-instance")
 *             .region("us-central1")
 *             .type("ENTERPRISE")
 *             .enableRbac(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Fusion Instance Event
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.pubsub.Topic;
 * import com.pulumi.gcp.pubsub.TopicArgs;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
 * import com.pulumi.gcp.datafusion.inputs.InstanceEventPublishConfigArgs;
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
 *         var eventTopic = new Topic("eventTopic", TopicArgs.builder()
 *             .name("my-instance")
 *             .build());
 * 
 *         var event = new Instance("event", InstanceArgs.builder()
 *             .name("my-instance")
 *             .region("us-central1")
 *             .type("BASIC")
 *             .eventPublishConfig(InstanceEventPublishConfigArgs.builder()
 *                 .enabled(true)
 *                 .topic(eventTopic.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Data Fusion Instance Zone
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.datafusion.Instance;
 * import com.pulumi.gcp.datafusion.InstanceArgs;
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
 *         var zone = new Instance("zone", InstanceArgs.builder()
 *             .name("my-instance")
 *             .region("us-central1")
 *             .zone("us-central1-a")
 *             .type("DEVELOPER")
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
 * Instance can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{region}}/instances/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Instance can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy default projects/{{project}}/locations/{{region}}/instances/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy")
public class InstanceIamPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The ID of the instance or a fully qualified identifier for the instance.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="policyData", refs={String.class}, tree="[0]")
    private Output<String> policyData;

    public Output<String> policyData() {
        return this.policyData;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The region of the Data Fusion instance.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region of the Data Fusion instance.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceIamPolicy(java.lang.String name) {
        this(name, InstanceIamPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceIamPolicy(java.lang.String name, InstanceIamPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceIamPolicy(java.lang.String name, InstanceIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private InstanceIamPolicy(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securitycenter/instanceIamPolicy:InstanceIamPolicy", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceIamPolicyArgs makeArgs(InstanceIamPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceIamPolicyArgs.Empty : args;
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
    public static InstanceIamPolicy get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceIamPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceIamPolicy(name, id, state, options);
    }
}
