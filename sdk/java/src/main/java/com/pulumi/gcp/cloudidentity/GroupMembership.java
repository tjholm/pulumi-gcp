// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudidentity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudidentity.GroupMembershipArgs;
import com.pulumi.gcp.cloudidentity.inputs.GroupMembershipState;
import com.pulumi.gcp.cloudidentity.outputs.GroupMembershipMemberKey;
import com.pulumi.gcp.cloudidentity.outputs.GroupMembershipPreferredMemberKey;
import com.pulumi.gcp.cloudidentity.outputs.GroupMembershipRole;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * A Membership defines a relationship between a Group and an entity belonging to that Group, referred to as a &#34;member&#34;.
 * 
 * To get more information about GroupMembership, see:
 * 
 * * [API documentation](https://cloud.google.com/identity/docs/reference/rest/v1/groups.memberships)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/identity/docs/how-to/memberships-google-groups)
 * 
 * &gt; **Warning:** If you are using User ADCs (Application Default Credentials) with this resource,
 * you must specify a `billing_project` and set `user_project_override` to true
 * in the provider configuration. Otherwise the Cloud Identity API will return a 403 error.
 * Your account must have the `serviceusage.services.use` permission on the
 * `billing_project` you defined.
 * 
 * ## Example Usage
 * 
 * ### Cloud Identity Group Membership
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.cloudidentity.Group;
 * import com.pulumi.gcp.cloudidentity.GroupArgs;
 * import com.pulumi.gcp.cloudidentity.inputs.GroupGroupKeyArgs;
 * import com.pulumi.gcp.cloudidentity.GroupMembership;
 * import com.pulumi.gcp.cloudidentity.GroupMembershipArgs;
 * import com.pulumi.gcp.cloudidentity.inputs.GroupMembershipPreferredMemberKeyArgs;
 * import com.pulumi.gcp.cloudidentity.inputs.GroupMembershipRoleArgs;
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
 *         var group = new Group("group", GroupArgs.builder()
 *             .displayName("my-identity-group")
 *             .parent("customers/A01b123xz")
 *             .groupKey(GroupGroupKeyArgs.builder()
 *                 .id("my-identity-group}{@literal @}{@code example.com")
 *                 .build())
 *             .labels(Map.of("cloudidentity.googleapis.com/groups.discussion_forum", ""))
 *             .build());
 * 
 *         var child_group = new Group("child-group", GroupArgs.builder()
 *             .displayName("my-identity-group-child")
 *             .parent("customers/A01b123xz")
 *             .groupKey(GroupGroupKeyArgs.builder()
 *                 .id("my-identity-group-child}{@literal @}{@code example.com")
 *                 .build())
 *             .labels(Map.of("cloudidentity.googleapis.com/groups.discussion_forum", ""))
 *             .build());
 * 
 *         var cloudIdentityGroupMembershipBasic = new GroupMembership("cloudIdentityGroupMembershipBasic", GroupMembershipArgs.builder()
 *             .group(group.id())
 *             .preferredMemberKey(GroupMembershipPreferredMemberKeyArgs.builder()
 *                 .id(child_group.groupKey().applyValue(groupKey -> groupKey.id()))
 *                 .build())
 *             .roles(GroupMembershipRoleArgs.builder()
 *                 .name("MEMBER")
 *                 .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Cloud Identity Group Membership User
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.cloudidentity.Group;
 * import com.pulumi.gcp.cloudidentity.GroupArgs;
 * import com.pulumi.gcp.cloudidentity.inputs.GroupGroupKeyArgs;
 * import com.pulumi.gcp.cloudidentity.GroupMembership;
 * import com.pulumi.gcp.cloudidentity.GroupMembershipArgs;
 * import com.pulumi.gcp.cloudidentity.inputs.GroupMembershipPreferredMemberKeyArgs;
 * import com.pulumi.gcp.cloudidentity.inputs.GroupMembershipRoleArgs;
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
 *         var group = new Group("group", GroupArgs.builder()
 *             .displayName("my-identity-group")
 *             .parent("customers/A01b123xz")
 *             .groupKey(GroupGroupKeyArgs.builder()
 *                 .id("my-identity-group}{@literal @}{@code example.com")
 *                 .build())
 *             .labels(Map.of("cloudidentity.googleapis.com/groups.discussion_forum", ""))
 *             .build());
 * 
 *         var cloudIdentityGroupMembershipBasic = new GroupMembership("cloudIdentityGroupMembershipBasic", GroupMembershipArgs.builder()
 *             .group(group.id())
 *             .preferredMemberKey(GroupMembershipPreferredMemberKeyArgs.builder()
 *                 .id("cloud_identity_user}{@literal @}{@code example.com")
 *                 .build())
 *             .roles(            
 *                 GroupMembershipRoleArgs.builder()
 *                     .name("MEMBER")
 *                     .build(),
 *                 GroupMembershipRoleArgs.builder()
 *                     .name("MANAGER")
 *                     .build())
 *             .build());
 * 
 *     }}{@code
 * }}{@code
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * GroupMembership can be imported using any of these accepted formats:
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, GroupMembership can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:cloudidentity/groupMembership:GroupMembership default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudidentity/groupMembership:GroupMembership")
public class GroupMembership extends com.pulumi.resources.CustomResource {
    /**
     * The time when the Membership was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time when the Membership was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The name of the Group to create this membership in.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output<String> group;

    /**
     * @return The name of the Group to create this membership in.
     * 
     */
    public Output<String> group() {
        return this.group;
    }
    /**
     * EntityKey of the member.
     * 
     */
    @Export(name="memberKey", refs={GroupMembershipMemberKey.class}, tree="[0]")
    private Output<GroupMembershipMemberKey> memberKey;

    /**
     * @return EntityKey of the member.
     * 
     */
    public Output<GroupMembershipMemberKey> memberKey() {
        return this.memberKey;
    }
    /**
     * The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the Membership, of the form groups/{group_id}/memberships/{membership_id}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * EntityKey of the member.
     * 
     */
    @Export(name="preferredMemberKey", refs={GroupMembershipPreferredMemberKey.class}, tree="[0]")
    private Output<GroupMembershipPreferredMemberKey> preferredMemberKey;

    /**
     * @return EntityKey of the member.
     * 
     */
    public Output<GroupMembershipPreferredMemberKey> preferredMemberKey() {
        return this.preferredMemberKey;
    }
    /**
     * The MembershipRoles that apply to the Membership.
     * Must not contain duplicate MembershipRoles with the same name.
     * Structure is documented below.
     * 
     */
    @Export(name="roles", refs={List.class,GroupMembershipRole.class}, tree="[0,1]")
    private Output<List<GroupMembershipRole>> roles;

    /**
     * @return The MembershipRoles that apply to the Membership.
     * Must not contain duplicate MembershipRoles with the same name.
     * Structure is documented below.
     * 
     */
    public Output<List<GroupMembershipRole>> roles() {
        return this.roles;
    }
    /**
     * The type of the membership.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the membership.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The time when the Membership was last updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The time when the Membership was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupMembership(java.lang.String name) {
        this(name, GroupMembershipArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupMembership(java.lang.String name, GroupMembershipArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupMembership(java.lang.String name, GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudidentity/groupMembership:GroupMembership", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupMembership(java.lang.String name, Output<java.lang.String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudidentity/groupMembership:GroupMembership", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupMembershipArgs makeArgs(GroupMembershipArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupMembershipArgs.Empty : args;
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
    public static GroupMembership get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupMembershipState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupMembership(name, id, state, options);
    }
}
