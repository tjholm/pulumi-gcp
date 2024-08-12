// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MembershipBindingArgs extends com.pulumi.resources.ResourceArgs {

    public static final MembershipBindingArgs Empty = new MembershipBindingArgs();

    /**
     * Labels for this Membership binding.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return Labels for this Membership binding.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * Location of the membership
     * 
     * ***
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return Location of the membership
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * The client-provided identifier of the membership binding.
     * 
     */
    @Import(name="membershipBindingId", required=true)
    private Output<String> membershipBindingId;

    /**
     * @return The client-provided identifier of the membership binding.
     * 
     */
    public Output<String> membershipBindingId() {
        return this.membershipBindingId;
    }

    /**
     * Id of the membership
     * 
     */
    @Import(name="membershipId", required=true)
    private Output<String> membershipId;

    /**
     * @return Id of the membership
     * 
     */
    public Output<String> membershipId() {
        return this.membershipId;
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * A Workspace resource name in the format
     * `projects/*&#47;locations/*&#47;scopes/*`.
     * 
     */
    @Import(name="scope", required=true)
    private Output<String> scope;

    /**
     * @return A Workspace resource name in the format
     * `projects/*&#47;locations/*&#47;scopes/*`.
     * 
     */
    public Output<String> scope() {
        return this.scope;
    }

    private MembershipBindingArgs() {}

    private MembershipBindingArgs(MembershipBindingArgs $) {
        this.labels = $.labels;
        this.location = $.location;
        this.membershipBindingId = $.membershipBindingId;
        this.membershipId = $.membershipId;
        this.project = $.project;
        this.scope = $.scope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MembershipBindingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MembershipBindingArgs $;

        public Builder() {
            $ = new MembershipBindingArgs();
        }

        public Builder(MembershipBindingArgs defaults) {
            $ = new MembershipBindingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param labels Labels for this Membership binding.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels Labels for this Membership binding.
         * 
         * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
         * Please refer to the field `effective_labels` for all of the labels present on the resource.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location Location of the membership
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Location of the membership
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param membershipBindingId The client-provided identifier of the membership binding.
         * 
         * @return builder
         * 
         */
        public Builder membershipBindingId(Output<String> membershipBindingId) {
            $.membershipBindingId = membershipBindingId;
            return this;
        }

        /**
         * @param membershipBindingId The client-provided identifier of the membership binding.
         * 
         * @return builder
         * 
         */
        public Builder membershipBindingId(String membershipBindingId) {
            return membershipBindingId(Output.of(membershipBindingId));
        }

        /**
         * @param membershipId Id of the membership
         * 
         * @return builder
         * 
         */
        public Builder membershipId(Output<String> membershipId) {
            $.membershipId = membershipId;
            return this;
        }

        /**
         * @param membershipId Id of the membership
         * 
         * @return builder
         * 
         */
        public Builder membershipId(String membershipId) {
            return membershipId(Output.of(membershipId));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
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
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param scope A Workspace resource name in the format
         * `projects/*&#47;locations/*&#47;scopes/*`.
         * 
         * @return builder
         * 
         */
        public Builder scope(Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope A Workspace resource name in the format
         * `projects/*&#47;locations/*&#47;scopes/*`.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public MembershipBindingArgs build() {
            if ($.location == null) {
                throw new MissingRequiredPropertyException("MembershipBindingArgs", "location");
            }
            if ($.membershipBindingId == null) {
                throw new MissingRequiredPropertyException("MembershipBindingArgs", "membershipBindingId");
            }
            if ($.membershipId == null) {
                throw new MissingRequiredPropertyException("MembershipBindingArgs", "membershipId");
            }
            if ($.scope == null) {
                throw new MissingRequiredPropertyException("MembershipBindingArgs", "scope");
            }
            return $;
        }
    }

}
