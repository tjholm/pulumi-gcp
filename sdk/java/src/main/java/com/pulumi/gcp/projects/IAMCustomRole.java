// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.projects;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.projects.IAMCustomRoleArgs;
import com.pulumi.gcp.projects.inputs.IAMCustomRoleState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows management of a customized Cloud IAM project role. For more information see
 * [the official documentation](https://cloud.google.com/iam/docs/understanding-custom-roles)
 * and
 * [API](https://cloud.google.com/iam/reference/rest/v1/projects.roles).
 * 
 * &gt; **Warning:** Note that custom roles in GCP have the concept of a soft-delete. There are two issues that may arise
 *  from this and how roles are propagated. 1) creating a role may involve undeleting and then updating a role with the
 *  same name, possibly causing confusing behavior between undelete and update. 2) A deleted role is permanently deleted
 *  after 7 days, but it can take up to 30 more days (i.e. between 7 and 37 days after deletion) before the role name is
 *  made available again. This means a deleted role that has been deleted for more than 7 days cannot be changed at all
 *  by the provider, and new roles cannot share that name.
 * 
 * ## Example Usage
 * 
 * This snippet creates a customized IAM role.
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.projects.IAMCustomRole;
 * import com.pulumi.gcp.projects.IAMCustomRoleArgs;
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
 *         var my_custom_role = new IAMCustomRole(&#34;my-custom-role&#34;, IAMCustomRoleArgs.builder()        
 *             .description(&#34;A description&#34;)
 *             .permissions(            
 *                 &#34;iam.roles.list&#34;,
 *                 &#34;iam.roles.create&#34;,
 *                 &#34;iam.roles.delete&#34;)
 *             .roleId(&#34;myCustomRole&#34;)
 *             .title(&#34;My Custom Role&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Custom Roles can be imported using any of these accepted formats* `projects/{{project}}/roles/{{role_id}}` * `{{project}}/{{role_id}}` * `{{role_id}}` When using the `pulumi import` command, Custom Roles can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMCustomRole:IAMCustomRole default projects/{{project}}/roles/{{role_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMCustomRole:IAMCustomRole default {{project}}/{{role_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:projects/iAMCustomRole:IAMCustomRole default {{role_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:projects/iAMCustomRole:IAMCustomRole")
public class IAMCustomRole extends com.pulumi.resources.CustomResource {
    /**
     * (Optional) The current deleted state of the role.
     * 
     */
    @Export(name="deleted", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> deleted;

    /**
     * @return (Optional) The current deleted state of the role.
     * 
     */
    public Output<Boolean> deleted() {
        return this.deleted;
    }
    /**
     * A human-readable description for the role.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-readable description for the role.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the role in the format `projects/{{project}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
     * 
     */
    @Export(name="permissions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> permissions;

    /**
     * @return The names of the permissions this role grants when bound in an IAM policy. At least one permission must be specified.
     * 
     */
    public Output<List<String>> permissions() {
        return this.permissions;
    }
    /**
     * The project that the custom role will be created in.
     * Defaults to the provider project configuration.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project that the custom role will be created in.
     * Defaults to the provider project configuration.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The camel case role id to use for this role. Cannot contain `-` characters.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The camel case role id to use for this role. Cannot contain `-` characters.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * The current launch stage of the role.
     * Defaults to `GA`.
     * List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
     * 
     */
    @Export(name="stage", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stage;

    /**
     * @return The current launch stage of the role.
     * Defaults to `GA`.
     * List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).
     * 
     */
    public Output<Optional<String>> stage() {
        return Codegen.optional(this.stage);
    }
    /**
     * A human-readable title for the role.
     * 
     */
    @Export(name="title", refs={String.class}, tree="[0]")
    private Output<String> title;

    /**
     * @return A human-readable title for the role.
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IAMCustomRole(String name) {
        this(name, IAMCustomRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IAMCustomRole(String name, IAMCustomRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IAMCustomRole(String name, IAMCustomRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:projects/iAMCustomRole:IAMCustomRole", name, args == null ? IAMCustomRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IAMCustomRole(String name, Output<String> id, @Nullable IAMCustomRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:projects/iAMCustomRole:IAMCustomRole", name, state, makeResourceOptions(options, id));
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
    public static IAMCustomRole get(String name, Output<String> id, @Nullable IAMCustomRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IAMCustomRole(name, id, state, options);
    }
}
