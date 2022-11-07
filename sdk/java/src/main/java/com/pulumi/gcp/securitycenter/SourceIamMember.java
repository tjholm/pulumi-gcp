// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.securitycenter;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.securitycenter.SourceIamMemberArgs;
import com.pulumi.gcp.securitycenter.inputs.SourceIamMemberState;
import com.pulumi.gcp.securitycenter.outputs.SourceIamMemberCondition;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Cloud Security Command Center&#39;s (Cloud SCC) finding source. A finding
 * source is an entity or a mechanism that can produce a finding. A source is
 * like a container of findings that come from the same scanner, logger,
 * monitor, etc.
 * 
 * To get more information about Source, see:
 * 
 * * [API documentation](https://cloud.google.com/security-command-center/docs/reference/rest/v1/organizations.sources)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/security-command-center/docs)
 * 
 * ## Example Usage
 * ### Scc Source Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.securitycenter.Source;
 * import com.pulumi.gcp.securitycenter.SourceArgs;
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
 *         var customSource = new Source(&#34;customSource&#34;, SourceArgs.builder()        
 *             .description(&#34;My custom Cloud Security Command Center Finding Source&#34;)
 *             .displayName(&#34;My Source&#34;)
 *             .organization(&#34;123456789&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Source can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:securitycenter/sourceIamMember:SourceIamMember default organizations/{{organization}}/sources/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:securitycenter/sourceIamMember:SourceIamMember default {{organization}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:securitycenter/sourceIamMember:SourceIamMember")
public class SourceIamMember extends com.pulumi.resources.CustomResource {
    @Export(name="condition", type=SourceIamMemberCondition.class, parameters={})
    private Output</* @Nullable */ SourceIamMemberCondition> condition;

    public Output<Optional<SourceIamMemberCondition>> condition() {
        return Codegen.optional(this.condition);
    }
    @Export(name="etag", type=String.class, parameters={})
    private Output<String> etag;

    public Output<String> etag() {
        return this.etag;
    }
    @Export(name="member", type=String.class, parameters={})
    private Output<String> member;

    public Output<String> member() {
        return this.member;
    }
    /**
     * The organization whose Cloud Security Command Center the Source
     * lives in.
     * 
     */
    @Export(name="organization", type=String.class, parameters={})
    private Output<String> organization;

    /**
     * @return The organization whose Cloud Security Command Center the Source
     * lives in.
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    @Export(name="role", type=String.class, parameters={})
    private Output<String> role;

    public Output<String> role() {
        return this.role;
    }
    @Export(name="source", type=String.class, parameters={})
    private Output<String> source;

    public Output<String> source() {
        return this.source;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SourceIamMember(String name) {
        this(name, SourceIamMemberArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SourceIamMember(String name, SourceIamMemberArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SourceIamMember(String name, SourceIamMemberArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securitycenter/sourceIamMember:SourceIamMember", name, args == null ? SourceIamMemberArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SourceIamMember(String name, Output<String> id, @Nullable SourceIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:securitycenter/sourceIamMember:SourceIamMember", name, state, makeResourceOptions(options, id));
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
    public static SourceIamMember get(String name, Output<String> id, @Nullable SourceIamMemberState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SourceIamMember(name, id, state, options);
    }
}
