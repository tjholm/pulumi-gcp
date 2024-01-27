// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.compute.inputs.SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityPolicyAdvancedOptionsConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityPolicyAdvancedOptionsConfigArgs Empty = new SecurityPolicyAdvancedOptionsConfigArgs();

    /**
     * Custom configuration to apply the JSON parsing. Only applicable when
     * `json_parsing` is set to `STANDARD`. Structure is documented below.
     * 
     */
    @Import(name="jsonCustomConfig")
    private @Nullable Output<SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs> jsonCustomConfig;

    /**
     * @return Custom configuration to apply the JSON parsing. Only applicable when
     * `json_parsing` is set to `STANDARD`. Structure is documented below.
     * 
     */
    public Optional<Output<SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs>> jsonCustomConfig() {
        return Optional.ofNullable(this.jsonCustomConfig);
    }

    /**
     * Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
     * 
     */
    @Import(name="jsonParsing")
    private @Nullable Output<String> jsonParsing;

    /**
     * @return Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
     * 
     */
    public Optional<Output<String>> jsonParsing() {
        return Optional.ofNullable(this.jsonParsing);
    }

    /**
     * Log level to use. Defaults to `NORMAL`.
     * 
     */
    @Import(name="logLevel")
    private @Nullable Output<String> logLevel;

    /**
     * @return Log level to use. Defaults to `NORMAL`.
     * 
     */
    public Optional<Output<String>> logLevel() {
        return Optional.ofNullable(this.logLevel);
    }

    /**
     * An optional list of case-insensitive request header names to use for resolving the callers client IP address.
     * 
     */
    @Import(name="userIpRequestHeaders")
    private @Nullable Output<List<String>> userIpRequestHeaders;

    /**
     * @return An optional list of case-insensitive request header names to use for resolving the callers client IP address.
     * 
     */
    public Optional<Output<List<String>>> userIpRequestHeaders() {
        return Optional.ofNullable(this.userIpRequestHeaders);
    }

    private SecurityPolicyAdvancedOptionsConfigArgs() {}

    private SecurityPolicyAdvancedOptionsConfigArgs(SecurityPolicyAdvancedOptionsConfigArgs $) {
        this.jsonCustomConfig = $.jsonCustomConfig;
        this.jsonParsing = $.jsonParsing;
        this.logLevel = $.logLevel;
        this.userIpRequestHeaders = $.userIpRequestHeaders;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityPolicyAdvancedOptionsConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityPolicyAdvancedOptionsConfigArgs $;

        public Builder() {
            $ = new SecurityPolicyAdvancedOptionsConfigArgs();
        }

        public Builder(SecurityPolicyAdvancedOptionsConfigArgs defaults) {
            $ = new SecurityPolicyAdvancedOptionsConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param jsonCustomConfig Custom configuration to apply the JSON parsing. Only applicable when
         * `json_parsing` is set to `STANDARD`. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder jsonCustomConfig(@Nullable Output<SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs> jsonCustomConfig) {
            $.jsonCustomConfig = jsonCustomConfig;
            return this;
        }

        /**
         * @param jsonCustomConfig Custom configuration to apply the JSON parsing. Only applicable when
         * `json_parsing` is set to `STANDARD`. Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder jsonCustomConfig(SecurityPolicyAdvancedOptionsConfigJsonCustomConfigArgs jsonCustomConfig) {
            return jsonCustomConfig(Output.of(jsonCustomConfig));
        }

        /**
         * @param jsonParsing Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
         * 
         * @return builder
         * 
         */
        public Builder jsonParsing(@Nullable Output<String> jsonParsing) {
            $.jsonParsing = jsonParsing;
            return this;
        }

        /**
         * @param jsonParsing Whether or not to JSON parse the payload body. Defaults to `DISABLED`.
         * 
         * @return builder
         * 
         */
        public Builder jsonParsing(String jsonParsing) {
            return jsonParsing(Output.of(jsonParsing));
        }

        /**
         * @param logLevel Log level to use. Defaults to `NORMAL`.
         * 
         * @return builder
         * 
         */
        public Builder logLevel(@Nullable Output<String> logLevel) {
            $.logLevel = logLevel;
            return this;
        }

        /**
         * @param logLevel Log level to use. Defaults to `NORMAL`.
         * 
         * @return builder
         * 
         */
        public Builder logLevel(String logLevel) {
            return logLevel(Output.of(logLevel));
        }

        /**
         * @param userIpRequestHeaders An optional list of case-insensitive request header names to use for resolving the callers client IP address.
         * 
         * @return builder
         * 
         */
        public Builder userIpRequestHeaders(@Nullable Output<List<String>> userIpRequestHeaders) {
            $.userIpRequestHeaders = userIpRequestHeaders;
            return this;
        }

        /**
         * @param userIpRequestHeaders An optional list of case-insensitive request header names to use for resolving the callers client IP address.
         * 
         * @return builder
         * 
         */
        public Builder userIpRequestHeaders(List<String> userIpRequestHeaders) {
            return userIpRequestHeaders(Output.of(userIpRequestHeaders));
        }

        /**
         * @param userIpRequestHeaders An optional list of case-insensitive request header names to use for resolving the callers client IP address.
         * 
         * @return builder
         * 
         */
        public Builder userIpRequestHeaders(String... userIpRequestHeaders) {
            return userIpRequestHeaders(List.of(userIpRequestHeaders));
        }

        public SecurityPolicyAdvancedOptionsConfigArgs build() {
            return $;
        }
    }

}
