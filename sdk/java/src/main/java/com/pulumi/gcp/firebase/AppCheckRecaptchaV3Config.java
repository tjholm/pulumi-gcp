// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.firebase.AppCheckRecaptchaV3ConfigArgs;
import com.pulumi.gcp.firebase.inputs.AppCheckRecaptchaV3ConfigState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * An app&#39;s reCAPTCHA V3 configuration object.
 * 
 * To get more information about RecaptchaV3Config, see:
 * 
 * * [API documentation](https://firebase.google.com/docs/reference/appcheck/rest/v1/projects.apps.recaptchaV3Config)
 * * How-to Guides
 *     * [Official Documentation](https://firebase.google.com/docs/app-check)
 * 
 * ## Example Usage
 * 
 * ### Firebase App Check Recaptcha V3 Config Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.firebase.WebApp;
 * import com.pulumi.gcp.firebase.WebAppArgs;
 * import com.pulumi.time.sleep;
 * import com.pulumi.time.SleepArgs;
 * import com.pulumi.gcp.firebase.AppCheckRecaptchaV3Config;
 * import com.pulumi.gcp.firebase.AppCheckRecaptchaV3ConfigArgs;
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
 *         var default_ = new WebApp("default", WebAppArgs.builder()
 *             .project("my-project-name")
 *             .displayName("Web App for reCAPTCHA V3")
 *             .build());
 * 
 *         // It takes a while for App Check to recognize the new app
 *         // If your app already exists, you don't have to wait 30 seconds.
 *         var wait30s = new Sleep("wait30s", SleepArgs.builder()
 *             .createDuration("30s")
 *             .build());
 * 
 *         var defaultAppCheckRecaptchaV3Config = new AppCheckRecaptchaV3Config("defaultAppCheckRecaptchaV3Config", AppCheckRecaptchaV3ConfigArgs.builder()
 *             .project("my-project-name")
 *             .appId(default_.appId())
 *             .siteSecret("6Lf9YnQpAAAAAC3-MHmdAllTbPwTZxpUw5d34YzX")
 *             .tokenTtl("7200s")
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
 * RecaptchaV3Config can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/apps/{{app_id}}/recaptchaV3Config`
 * 
 * * `{{project}}/{{app_id}}`
 * 
 * * `{{app_id}}`
 * 
 * When using the `pulumi import` command, RecaptchaV3Config can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config default projects/{{project}}/apps/{{app_id}}/recaptchaV3Config
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config default {{project}}/{{app_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config default {{app_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config")
public class AppCheckRecaptchaV3Config extends com.pulumi.resources.CustomResource {
    /**
     * The ID of an
     * [Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
     * 
     * ***
     * 
     */
    @Export(name="appId", refs={String.class}, tree="[0]")
    private Output<String> appId;

    /**
     * @return The ID of an
     * [Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
     * 
     * ***
     * 
     */
    public Output<String> appId() {
        return this.appId;
    }
    /**
     * The relative resource name of the reCAPTCHA V3 configuration object
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The relative resource name of the reCAPTCHA V3 configuration object
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
     * The site secret used to identify your service for reCAPTCHA v3 verification.
     * For security reasons, this field will never be populated in any response.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    @Export(name="siteSecret", refs={String.class}, tree="[0]")
    private Output<String> siteSecret;

    /**
     * @return The site secret used to identify your service for reCAPTCHA v3 verification.
     * For security reasons, this field will never be populated in any response.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     */
    public Output<String> siteSecret() {
        return this.siteSecret;
    }
    /**
     * Whether the siteSecret was previously set. Since we will never return the siteSecret field, this field is the only way to find out whether it was previously set.
     * 
     */
    @Export(name="siteSecretSet", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> siteSecretSet;

    /**
     * @return Whether the siteSecret was previously set. Since we will never return the siteSecret field, this field is the only way to find out whether it was previously set.
     * 
     */
    public Output<Boolean> siteSecretSet() {
        return this.siteSecretSet;
    }
    /**
     * Specifies the duration for which App Check tokens exchanged from reCAPTCHA V3 artifacts will be valid.
     * If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
     * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    @Export(name="tokenTtl", refs={String.class}, tree="[0]")
    private Output<String> tokenTtl;

    /**
     * @return Specifies the duration for which App Check tokens exchanged from reCAPTCHA V3 artifacts will be valid.
     * If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
     * A duration in seconds with up to nine fractional digits, ending with &#39;s&#39;. Example: &#34;3.5s&#34;.
     * 
     */
    public Output<String> tokenTtl() {
        return this.tokenTtl;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppCheckRecaptchaV3Config(String name) {
        this(name, AppCheckRecaptchaV3ConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppCheckRecaptchaV3Config(String name, AppCheckRecaptchaV3ConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppCheckRecaptchaV3Config(String name, AppCheckRecaptchaV3ConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config", name, args == null ? AppCheckRecaptchaV3ConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppCheckRecaptchaV3Config(String name, Output<String> id, @Nullable AppCheckRecaptchaV3ConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "siteSecret"
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
    public static AppCheckRecaptchaV3Config get(String name, Output<String> id, @Nullable AppCheckRecaptchaV3ConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppCheckRecaptchaV3Config(name, id, state, options);
    }
}
