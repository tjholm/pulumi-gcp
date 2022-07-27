// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.appengine;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.appengine.ApplicationUrlDispatchRulesArgs;
import com.pulumi.gcp.appengine.inputs.ApplicationUrlDispatchRulesState;
import com.pulumi.gcp.appengine.outputs.ApplicationUrlDispatchRulesDispatchRule;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Rules to match an HTTP request and dispatch that request to a service.
 * 
 * To get more information about ApplicationUrlDispatchRules, see:
 * 
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps#UrlDispatchRule)
 * 
 * ## Example Usage
 * ### App Engine Application Url Dispatch Rules Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.storage.Bucket;
 * import com.pulumi.gcp.storage.BucketArgs;
 * import com.pulumi.gcp.storage.BucketObject;
 * import com.pulumi.gcp.storage.BucketObjectArgs;
 * import com.pulumi.gcp.appengine.StandardAppVersion;
 * import com.pulumi.gcp.appengine.StandardAppVersionArgs;
 * import com.pulumi.gcp.appengine.inputs.StandardAppVersionEntrypointArgs;
 * import com.pulumi.gcp.appengine.inputs.StandardAppVersionDeploymentArgs;
 * import com.pulumi.gcp.appengine.inputs.StandardAppVersionDeploymentZipArgs;
 * import com.pulumi.gcp.appengine.ApplicationUrlDispatchRules;
 * import com.pulumi.gcp.appengine.ApplicationUrlDispatchRulesArgs;
 * import com.pulumi.gcp.appengine.inputs.ApplicationUrlDispatchRulesDispatchRuleArgs;
 * import com.pulumi.asset.FileAsset;
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
 *         var bucket = new Bucket(&#34;bucket&#34;, BucketArgs.builder()        
 *             .location(&#34;US&#34;)
 *             .build());
 * 
 *         var object = new BucketObject(&#34;object&#34;, BucketObjectArgs.builder()        
 *             .bucket(bucket.name())
 *             .source(new FileAsset(&#34;./test-fixtures/appengine/hello-world.zip&#34;))
 *             .build());
 * 
 *         var adminV3 = new StandardAppVersion(&#34;adminV3&#34;, StandardAppVersionArgs.builder()        
 *             .versionId(&#34;v3&#34;)
 *             .service(&#34;admin&#34;)
 *             .runtime(&#34;nodejs10&#34;)
 *             .entrypoint(StandardAppVersionEntrypointArgs.builder()
 *                 .shell(&#34;node ./app.js&#34;)
 *                 .build())
 *             .deployment(StandardAppVersionDeploymentArgs.builder()
 *                 .zip(StandardAppVersionDeploymentZipArgs.builder()
 *                     .sourceUrl(Output.tuple(bucket.name(), object.name()).applyValue(values -&gt; {
 *                         var bucketName = values.t1;
 *                         var objectName = values.t2;
 *                         return String.format(&#34;https://storage.googleapis.com/%s/%s&#34;, bucketName,objectName);
 *                     }))
 *                     .build())
 *                 .build())
 *             .envVariables(Map.of(&#34;port&#34;, &#34;8080&#34;))
 *             .noopOnDestroy(true)
 *             .build());
 * 
 *         var webService = new ApplicationUrlDispatchRules(&#34;webService&#34;, ApplicationUrlDispatchRulesArgs.builder()        
 *             .dispatchRules(            
 *                 ApplicationUrlDispatchRulesDispatchRuleArgs.builder()
 *                     .domain(&#34;*&#34;)
 *                     .path(&#34;/*&#34;)
 *                     .service(&#34;default&#34;)
 *                     .build(),
 *                 ApplicationUrlDispatchRulesDispatchRuleArgs.builder()
 *                     .domain(&#34;*&#34;)
 *                     .path(&#34;/admin/*&#34;)
 *                     .service(adminV3.service())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ApplicationUrlDispatchRules can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules default {{project}}
 * ```
 * 
 */
@ResourceType(type="gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules")
public class ApplicationUrlDispatchRules extends com.pulumi.resources.CustomResource {
    /**
     * Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     * 
     */
    @Export(name="dispatchRules", type=List.class, parameters={ApplicationUrlDispatchRulesDispatchRule.class})
    private Output<List<ApplicationUrlDispatchRulesDispatchRule>> dispatchRules;

    /**
     * @return Rules to match an HTTP request and dispatch that request to a service.
     * Structure is documented below.
     * 
     */
    public Output<List<ApplicationUrlDispatchRulesDispatchRule>> dispatchRules() {
        return this.dispatchRules;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
    public ApplicationUrlDispatchRules(String name) {
        this(name, ApplicationUrlDispatchRulesArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApplicationUrlDispatchRules(String name, ApplicationUrlDispatchRulesArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApplicationUrlDispatchRules(String name, ApplicationUrlDispatchRulesArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, args == null ? ApplicationUrlDispatchRulesArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApplicationUrlDispatchRules(String name, Output<String> id, @Nullable ApplicationUrlDispatchRulesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:appengine/applicationUrlDispatchRules:ApplicationUrlDispatchRules", name, state, makeResourceOptions(options, id));
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
    public static ApplicationUrlDispatchRules get(String name, Output<String> id, @Nullable ApplicationUrlDispatchRulesState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApplicationUrlDispatchRules(name, id, state, options);
    }
}
