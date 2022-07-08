// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.binaryauthorization;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.binaryauthorization.inputs.AttestorIamMemberConditionArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AttestorIamMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final AttestorIamMemberArgs Empty = new AttestorIamMemberArgs();

    /**
     * Used to find the parent resource to bind the IAM policy to
     * 
     */
    @Import(name="attestor", required=true)
    private Output<String> attestor;

    /**
     * @return Used to find the parent resource to bind the IAM policy to
     * 
     */
    public Output<String> attestor() {
        return this.attestor;
    }

    @Import(name="condition")
    private @Nullable Output<AttestorIamMemberConditionArgs> condition;

    public Optional<Output<AttestorIamMemberConditionArgs>> condition() {
        return Optional.ofNullable(this.condition);
    }

    @Import(name="member", required=true)
    private Output<String> member;

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
     * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The role that should be applied. Only one
     * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
     * `[projects|organizations]/{parent-name}/roles/{role-name}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    private AttestorIamMemberArgs() {}

    private AttestorIamMemberArgs(AttestorIamMemberArgs $) {
        this.attestor = $.attestor;
        this.condition = $.condition;
        this.member = $.member;
        this.project = $.project;
        this.role = $.role;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AttestorIamMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AttestorIamMemberArgs $;

        public Builder() {
            $ = new AttestorIamMemberArgs();
        }

        public Builder(AttestorIamMemberArgs defaults) {
            $ = new AttestorIamMemberArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attestor Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder attestor(Output<String> attestor) {
            $.attestor = attestor;
            return this;
        }

        /**
         * @param attestor Used to find the parent resource to bind the IAM policy to
         * 
         * @return builder
         * 
         */
        public Builder attestor(String attestor) {
            return attestor(Output.of(attestor));
        }

        public Builder condition(@Nullable Output<AttestorIamMemberConditionArgs> condition) {
            $.condition = condition;
            return this;
        }

        public Builder condition(AttestorIamMemberConditionArgs condition) {
            return condition(Output.of(condition));
        }

        public Builder member(Output<String> member) {
            $.member = member;
            return this;
        }

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
         * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
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
         * `gcp.binaryauthorization.AttestorIamBinding` can be used per role. Note that custom roles must be of the format
         * `[projects|organizations]/{parent-name}/roles/{role-name}`.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        public AttestorIamMemberArgs build() {
            $.attestor = Objects.requireNonNull($.attestor, "expected parameter 'attestor' to be non-null");
            $.member = Objects.requireNonNull($.member, "expected parameter 'member' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
