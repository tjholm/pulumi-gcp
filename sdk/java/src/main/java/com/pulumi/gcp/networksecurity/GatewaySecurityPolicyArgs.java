// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networksecurity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GatewaySecurityPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final GatewaySecurityPolicyArgs Empty = new GatewaySecurityPolicyArgs();

    /**
     * A free-text description of the resource. Max length 1024 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A free-text description of the resource. Max length 1024 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The location of the gateway security policy.
     * The default value is `global`.
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location of the gateway security policy.
     * The default value is `global`.
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
     * gatewaySecurityPolicy should match the pattern:(^a-z?$).
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
     * gatewaySecurityPolicy should match the pattern:(^a-z?$).
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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
     * Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
     * 
     */
    @Import(name="tlsInspectionPolicy")
    private @Nullable Output<String> tlsInspectionPolicy;

    /**
     * @return Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
     * 
     */
    public Optional<Output<String>> tlsInspectionPolicy() {
        return Optional.ofNullable(this.tlsInspectionPolicy);
    }

    private GatewaySecurityPolicyArgs() {}

    private GatewaySecurityPolicyArgs(GatewaySecurityPolicyArgs $) {
        this.description = $.description;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.tlsInspectionPolicy = $.tlsInspectionPolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GatewaySecurityPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GatewaySecurityPolicyArgs $;

        public Builder() {
            $ = new GatewaySecurityPolicyArgs();
        }

        public Builder(GatewaySecurityPolicyArgs defaults) {
            $ = new GatewaySecurityPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A free-text description of the resource. Max length 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param location The location of the gateway security policy.
         * The default value is `global`.
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location of the gateway security policy.
         * The default value is `global`.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
         * gatewaySecurityPolicy should match the pattern:(^a-z?$).
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
         * gatewaySecurityPolicy should match the pattern:(^a-z?$).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param tlsInspectionPolicy Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
         * 
         * @return builder
         * 
         */
        public Builder tlsInspectionPolicy(@Nullable Output<String> tlsInspectionPolicy) {
            $.tlsInspectionPolicy = tlsInspectionPolicy;
            return this;
        }

        /**
         * @param tlsInspectionPolicy Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
         * 
         * @return builder
         * 
         */
        public Builder tlsInspectionPolicy(String tlsInspectionPolicy) {
            return tlsInspectionPolicy(Output.of(tlsInspectionPolicy));
        }

        public GatewaySecurityPolicyArgs build() {
            return $;
        }
    }

}
