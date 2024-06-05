// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.workstations;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.workstations.inputs.WorkstationIamMemberConditionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkstationIamMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkstationIamMemberArgs Empty = new WorkstationIamMemberArgs();

    @Import(name="condition")
    private @Nullable Output<WorkstationIamMemberConditionArgs> condition;

    public Optional<Output<WorkstationIamMemberConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    /**
     * The location where the workstation parent resources reside.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location where the workstation parent resources reside.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    @Import(name="member", required=true)
    private Output<String> member;

    /**
     * @return Identities that will be granted the privilege in `role`.
     * Each entry can have one of the following values:
     * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
     * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
     * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
     * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
     * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
     * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
     * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
     * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
     * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
     * 
     */
    public Output<String> member() {
        return this.member;
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The role that should be applied. Only one
     * `gcp.workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    @Import(name="workstationClusterId", required=true)
    private Output<String> workstationClusterId;

    public Output<String> workstationClusterId() {
        return this.workstationClusterId;
    }

    @Import(name="workstationConfigId", required=true)
    private Output<String> workstationConfigId;

    public Output<String> workstationConfigId() {
        return this.workstationConfigId;
    }

    @Import(name="workstationId", required=true)
    private Output<String> workstationId;

    public Output<String> workstationId() {
        return this.workstationId;
    }

    private WorkstationIamMemberArgs() {}

    private WorkstationIamMemberArgs(WorkstationIamMemberArgs $) {
        this.condition = $.condition;
        this.location = $.location;
        this.member = $.member;
        this.project = $.project;
        this.role = $.role;
        this.workstationClusterId = $.workstationClusterId;
        this.workstationConfigId = $.workstationConfigId;
        this.workstationId = $.workstationId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkstationIamMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkstationIamMemberArgs $;

        public Builder() {
            $ = new WorkstationIamMemberArgs();
        }

        public Builder(WorkstationIamMemberArgs defaults) {
            $ = new WorkstationIamMemberArgs(Objects.requireNonNull(defaults));
        }

        public Builder condition(@Nullable Output<WorkstationIamMemberConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(WorkstationIamMemberConditionArgs condition) {
            return condition(Output.of(condition));
        }

        /**
         * @param location The location where the workstation parent resources reside.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location where the workstation parent resources reside.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
         * location is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param member Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
         * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
         * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
         * 
         * @return builder
         * 
         */
        public Builder member(Output<String> member) {
            $.member = member;
            return this;
        }

        /**
         * @param member Identities that will be granted the privilege in `role`.
         * Each entry can have one of the following values:
         * * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
         * * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
         * * **user:{emailid}**: An email address that represents a specific Google account. For example, alice{@literal @}gmail.com or joe{@literal @}example.com.
         * * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app{@literal @}appspot.gserviceaccount.com.
         * * **group:{emailid}**: An email address that represents a Google group. For example, admins{@literal @}example.com.
         * * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
         * * **projectOwner:projectid**: Owners of the given project. For example, &#34;projectOwner:my-example-project&#34;
         * * **projectEditor:projectid**: Editors of the given project. For example, &#34;projectEditor:my-example-project&#34;
         * * **projectViewer:projectid**: Viewers of the given project. For example, &#34;projectViewer:my-example-project&#34;
         * 
         * @return builder
         * 
         */
        public Builder member(String member) {
            return member(Output.of(member));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The role that should be applied. Only one
         * `gcp.workstations.WorkstationIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder workstationClusterId(Output<String> workstationClusterId) {
            $.workstationClusterId = workstationClusterId;
            return this;
        }

        public Builder workstationClusterId(String workstationClusterId) {
            return workstationClusterId(Output.of(workstationClusterId));
        }

        public Builder workstationConfigId(Output<String> workstationConfigId) {
            $.workstationConfigId = workstationConfigId;
            return this;
        }

        public Builder workstationConfigId(String workstationConfigId) {
            return workstationConfigId(Output.of(workstationConfigId));
        }

        public Builder workstationId(Output<String> workstationId) {
            $.workstationId = workstationId;
            return this;
        }

        public Builder workstationId(String workstationId) {
            return workstationId(Output.of(workstationId));
        }

        public WorkstationIamMemberArgs build() {
            if ($.member == null) {
                throw new MissingRequiredPropertyException("WorkstationIamMemberArgs", "member");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("WorkstationIamMemberArgs", "role");
            }
            if ($.workstationClusterId == null) {
                throw new MissingRequiredPropertyException("WorkstationIamMemberArgs", "workstationClusterId");
            }
            if ($.workstationConfigId == null) {
                throw new MissingRequiredPropertyException("WorkstationIamMemberArgs", "workstationConfigId");
            }
            if ($.workstationId == null) {
                throw new MissingRequiredPropertyException("WorkstationIamMemberArgs", "workstationId");
            }
            return $;
        }
    }

}
