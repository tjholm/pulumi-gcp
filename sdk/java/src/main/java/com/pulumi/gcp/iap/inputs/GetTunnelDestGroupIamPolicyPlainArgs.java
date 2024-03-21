// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.iap.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTunnelDestGroupIamPolicyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTunnelDestGroupIamPolicyPlainArgs Empty = new GetTunnelDestGroupIamPolicyPlainArgs();

    @Import(name="destGroup", required=true)
    private String destGroup;

    public String destGroup() {
        return this.destGroup;
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable String project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     * 
     */
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The region of the tunnel group. Must be the same as the network resources in the group.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The region of the tunnel group. Must be the same as the network resources in the group.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    private GetTunnelDestGroupIamPolicyPlainArgs() {}

    private GetTunnelDestGroupIamPolicyPlainArgs(GetTunnelDestGroupIamPolicyPlainArgs $) {
        this.destGroup = $.destGroup;
        this.project = $.project;
        this.region = $.region;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTunnelDestGroupIamPolicyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTunnelDestGroupIamPolicyPlainArgs $;

        public Builder() {
            $ = new GetTunnelDestGroupIamPolicyPlainArgs();
        }

        public Builder(GetTunnelDestGroupIamPolicyPlainArgs defaults) {
            $ = new GetTunnelDestGroupIamPolicyPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder destGroup(String destGroup) {
            $.destGroup = destGroup;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable String project) {
            $.project = project;
            return this;
        }

        /**
         * @param region The region of the tunnel group. Must be the same as the network resources in the group.
         * Used to find the parent resource to bind the IAM policy to. If not specified,
         * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
         * region is specified, it is taken from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        public GetTunnelDestGroupIamPolicyPlainArgs build() {
            if ($.destGroup == null) {
                throw new MissingRequiredPropertyException("GetTunnelDestGroupIamPolicyPlainArgs", "destGroup");
            }
            return $;
        }
    }

}
