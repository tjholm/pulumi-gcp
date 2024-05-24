// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.BackendServiceSignedUrlKeyArgs;
import com.pulumi.gcp.compute.inputs.BackendServiceSignedUrlKeyState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * A key for signing Cloud CDN signed URLs for Backend Services.
 * 
 * To get more information about BackendServiceSignedUrlKey, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
 * * How-to Guides
 *     * [Using Signed URLs](https://cloud.google.com/cdn/docs/using-signed-urls/)
 * 
 * ## Example Usage
 * 
 * ### Backend Service Signed Url Key
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.random.RandomId;
 * import com.pulumi.random.RandomIdArgs;
 * import com.pulumi.gcp.compute.InstanceTemplate;
 * import com.pulumi.gcp.compute.InstanceTemplateArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceTemplateDiskArgs;
 * import com.pulumi.gcp.compute.InstanceGroupManager;
 * import com.pulumi.gcp.compute.InstanceGroupManagerArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceGroupManagerVersionArgs;
 * import com.pulumi.gcp.compute.HttpHealthCheck;
 * import com.pulumi.gcp.compute.HttpHealthCheckArgs;
 * import com.pulumi.gcp.compute.BackendService;
 * import com.pulumi.gcp.compute.BackendServiceArgs;
 * import com.pulumi.gcp.compute.inputs.BackendServiceBackendArgs;
 * import com.pulumi.gcp.compute.BackendServiceSignedUrlKey;
 * import com.pulumi.gcp.compute.BackendServiceSignedUrlKeyArgs;
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
 *         var urlSignature = new RandomId("urlSignature", RandomIdArgs.builder()
 *             .byteLength(16)
 *             .build());
 * 
 *         var webserver = new InstanceTemplate("webserver", InstanceTemplateArgs.builder()
 *             .name("standard-webserver")
 *             .machineType("e2-medium")
 *             .networkInterfaces(InstanceTemplateNetworkInterfaceArgs.builder()
 *                 .network("default")
 *                 .build())
 *             .disks(InstanceTemplateDiskArgs.builder()
 *                 .sourceImage("debian-cloud/debian-11")
 *                 .autoDelete(true)
 *                 .boot(true)
 *                 .build())
 *             .build());
 * 
 *         var webservers = new InstanceGroupManager("webservers", InstanceGroupManagerArgs.builder()
 *             .name("my-webservers")
 *             .versions(InstanceGroupManagerVersionArgs.builder()
 *                 .instanceTemplate(webserver.id())
 *                 .name("primary")
 *                 .build())
 *             .baseInstanceName("webserver")
 *             .zone("us-central1-f")
 *             .targetSize(1)
 *             .build());
 * 
 *         var default_ = new HttpHealthCheck("default", HttpHealthCheckArgs.builder()
 *             .name("test")
 *             .requestPath("/")
 *             .checkIntervalSec(1)
 *             .timeoutSec(1)
 *             .build());
 * 
 *         var exampleBackend = new BackendService("exampleBackend", BackendServiceArgs.builder()
 *             .name("my-backend-service")
 *             .description("Our company website")
 *             .portName("http")
 *             .protocol("HTTP")
 *             .timeoutSec(10)
 *             .enableCdn(true)
 *             .backends(BackendServiceBackendArgs.builder()
 *                 .group(webservers.instanceGroup())
 *                 .build())
 *             .healthChecks(default_.id())
 *             .build());
 * 
 *         var backendKey = new BackendServiceSignedUrlKey("backendKey", BackendServiceSignedUrlKeyArgs.builder()
 *             .name("test-key")
 *             .keyValue(urlSignature.b64Url())
 *             .backendService(exampleBackend.name())
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
 * This resource does not support import.
 * 
 */
@ResourceType(type="gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey")
public class BackendServiceSignedUrlKey extends com.pulumi.resources.CustomResource {
    /**
     * The backend service this signed URL key belongs.
     * 
     * ***
     * 
     */
    @Export(name="backendService", refs={String.class}, tree="[0]")
    private Output<String> backendService;

    /**
     * @return The backend service this signed URL key belongs.
     * 
     * ***
     * 
     */
    public Output<String> backendService() {
        return this.backendService;
    }
    /**
     * 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Export(name="keyValue", refs={String.class}, tree="[0]")
    private Output<String> keyValue;

    /**
     * @return 128-bit key value used for signing the URL. The key value must be a
     * valid RFC 4648 Section 5 base64url encoded string.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<String> keyValue() {
        return this.keyValue;
    }
    /**
     * Name of the signed URL key.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the signed URL key.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackendServiceSignedUrlKey(String name) {
        this(name, BackendServiceSignedUrlKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackendServiceSignedUrlKey(String name, BackendServiceSignedUrlKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackendServiceSignedUrlKey(String name, BackendServiceSignedUrlKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, args == null ? BackendServiceSignedUrlKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackendServiceSignedUrlKey(String name, Output<String> id, @Nullable BackendServiceSignedUrlKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/backendServiceSignedUrlKey:BackendServiceSignedUrlKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "keyValue"
            ))
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
    public static BackendServiceSignedUrlKey get(String name, Output<String> id, @Nullable BackendServiceSignedUrlKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackendServiceSignedUrlKey(name, id, state, options);
    }
}
