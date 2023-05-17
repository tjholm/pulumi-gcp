// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firebase.AppleAppArgs;
import com.pulumi.gcp.firebase.inputs.AppleAppState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ### Firebase Apple App Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebase.AppleApp;
 * import com.pulumi.gcp.firebase.AppleAppArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var default_ = new AppleApp(&#34;default&#34;, AppleAppArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .displayName(&#34;Display Name Basic&#34;)
 *             .bundleId(&#34;apple.app.12345&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Firebase Apple App Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebase.AppleApp;
 * import com.pulumi.gcp.firebase.AppleAppArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var full = new AppleApp(&#34;full&#34;, AppleAppArgs.builder()        
 *             .project(&#34;my-project-name&#34;)
 *             .displayName(&#34;Display Name Full&#34;)
 *             .bundleId(&#34;apple.app.12345&#34;)
 *             .appStoreId(&#34;12345&#34;)
 *             .teamId(&#34;9987654321&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AppleApp can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:firebase/appleApp:AppleApp default projects/{{project}}/iosApps/{{appId}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firebase/appleApp:AppleApp default {{project}}/{{appId}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firebase/appleApp:AppleApp default iosApps/{{appId}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:firebase/appleApp:AppleApp default {{appId}}
 * ```
 * 
 */
@ResourceType(type="gcp:firebase/appleApp:AppleApp")
public class AppleApp extends com.pulumi.resources.CustomResource {
    /**
     * The globally unique, Firebase-assigned identifier of the App.
     * This identifier should be treated as an opaque token, as the data format is not specified.
     * 
     */
    @Export(name="appId", type=String.class, parameters={})
    private Output<String> appId;

    /**
     * @return The globally unique, Firebase-assigned identifier of the App.
     * This identifier should be treated as an opaque token, as the data format is not specified.
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.
     * 
     */
    @Export(name="appStoreId", type=String.class, parameters={})
    private Output</* @Nullable */ String> appStoreId;

    /**
     * @return The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.
     * 
     */
    public Output<Optional<String>> appStoreId() {
        return Codegen.optional(this.appStoreId);
    }
    /**
     * The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.
     * 
     */
    @Export(name="bundleId", type=String.class, parameters={})
    private Output<String> bundleId;

    /**
     * @return The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.
     * 
     */
    public Output<String> bundleId() {
        return this.bundleId;
    }
    /**
     * (Optional) Set to &#39;ABANDON&#39; to allow the Apple to be untracked from terraform state rather than deleted upon &#39;terraform
     * destroy&#39;. This is useful because the Apple may be serving traffic. Set to &#39;DELETE&#39; to delete the Apple. Defaults to
     * &#39;DELETE&#39;.
     * 
     */
    @Export(name="deletionPolicy", type=String.class, parameters={})
    private Output</* @Nullable */ String> deletionPolicy;

    /**
     * @return (Optional) Set to &#39;ABANDON&#39; to allow the Apple to be untracked from terraform state rather than deleted upon &#39;terraform
     * destroy&#39;. This is useful because the Apple may be serving traffic. Set to &#39;DELETE&#39; to delete the Apple. Defaults to
     * &#39;DELETE&#39;.
     * 
     */
    public Output<Optional<String>> deletionPolicy() {
        return Codegen.optional(this.deletionPolicy);
    }
    /**
     * The user-assigned display name of the App.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return The user-assigned display name of the App.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * The fully qualified resource name of the App, for example:
     * projects/projectId/iosApps/appId
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The fully qualified resource name of the App, for example:
     * projects/projectId/iosApps/appId
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
     * The Apple Developer Team ID associated with the App in the App Store.
     * 
     */
    @Export(name="teamId", type=String.class, parameters={})
    private Output</* @Nullable */ String> teamId;

    /**
     * @return The Apple Developer Team ID associated with the App in the App Store.
     * 
     */
    public Output<Optional<String>> teamId() {
        return Codegen.optional(this.teamId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppleApp(String name) {
        this(name, AppleAppArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppleApp(String name, AppleAppArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppleApp(String name, AppleAppArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/appleApp:AppleApp", name, args == null ? AppleAppArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppleApp(String name, Output<String> id, @Nullable AppleAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/appleApp:AppleApp", name, state, makeResourceOptions(options, id));
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
    public static AppleApp get(String name, Output<String> id, @Nullable AppleAppState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppleApp(name, id, state, options);
    }
}
