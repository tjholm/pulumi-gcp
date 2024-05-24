// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HealthCheckSslHealthCheckArgs extends com.pulumi.resources.ResourceArgs {

    public static final HealthCheckSslHealthCheckArgs Empty = new HealthCheckSslHealthCheckArgs();

    /**
     * The TCP port number for the HTTP2 health check request.
     * The default value is 443.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The TCP port number for the HTTP2 health check request.
     * The default value is 443.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * Port name as defined in InstanceGroup#NamedPort#name. If both port and
     * port_name are defined, port takes precedence.
     * 
     */
    @Import(name="portName")
    private @Nullable Output<String> portName;

    /**
     * @return Port name as defined in InstanceGroup#NamedPort#name. If both port and
     * port_name are defined, port takes precedence.
     * 
     */
    public Optional<Output<String>> portName() {
        return Optional.ofNullable(this.portName);
    }

    /**
     * Specifies how port is selected for health checking, can be one of the
     * following values:
     * * `USE_FIXED_PORT`: The port number in `port` is used for health checking.
     * * `USE_NAMED_PORT`: The `portName` is used for health checking.
     * * `USE_SERVING_PORT`: For NetworkEndpointGroup, the port specified for each
     *   network endpoint is used for health checking. For other backends, the
     *   port or named port specified in the Backend Service is used for health
     *   checking.
     *   If not specified, HTTP2 health check follows behavior specified in `port` and
     *   `portName` fields.
     *   Possible values are: `USE_FIXED_PORT`, `USE_NAMED_PORT`, `USE_SERVING_PORT`.
     * 
     */
    @Import(name="portSpecification")
    private @Nullable Output<String> portSpecification;

    /**
     * @return Specifies how port is selected for health checking, can be one of the
     * following values:
     * * `USE_FIXED_PORT`: The port number in `port` is used for health checking.
     * * `USE_NAMED_PORT`: The `portName` is used for health checking.
     * * `USE_SERVING_PORT`: For NetworkEndpointGroup, the port specified for each
     *   network endpoint is used for health checking. For other backends, the
     *   port or named port specified in the Backend Service is used for health
     *   checking.
     *   If not specified, HTTP2 health check follows behavior specified in `port` and
     *   `portName` fields.
     *   Possible values are: `USE_FIXED_PORT`, `USE_NAMED_PORT`, `USE_SERVING_PORT`.
     * 
     */
    public Optional<Output<String>> portSpecification() {
        return Optional.ofNullable(this.portSpecification);
    }

    /**
     * Specifies the type of proxy header to append before sending data to the
     * backend.
     * Default value is `NONE`.
     * Possible values are: `NONE`, `PROXY_V1`.
     * 
     */
    @Import(name="proxyHeader")
    private @Nullable Output<String> proxyHeader;

    /**
     * @return Specifies the type of proxy header to append before sending data to the
     * backend.
     * Default value is `NONE`.
     * Possible values are: `NONE`, `PROXY_V1`.
     * 
     */
    public Optional<Output<String>> proxyHeader() {
        return Optional.ofNullable(this.proxyHeader);
    }

    /**
     * The application data to send once the SSL connection has been
     * established (default value is empty). If both request and response are
     * empty, the connection establishment alone will indicate health. The request
     * data can only be ASCII.
     * 
     */
    @Import(name="request")
    private @Nullable Output<String> request;

    /**
     * @return The application data to send once the SSL connection has been
     * established (default value is empty). If both request and response are
     * empty, the connection establishment alone will indicate health. The request
     * data can only be ASCII.
     * 
     */
    public Optional<Output<String>> request() {
        return Optional.ofNullable(this.request);
    }

    /**
     * The bytes to match against the beginning of the response data. If left empty
     * (the default value), any response will indicate health. The response data
     * can only be ASCII.
     * 
     */
    @Import(name="response")
    private @Nullable Output<String> response;

    /**
     * @return The bytes to match against the beginning of the response data. If left empty
     * (the default value), any response will indicate health. The response data
     * can only be ASCII.
     * 
     */
    public Optional<Output<String>> response() {
        return Optional.ofNullable(this.response);
    }

    private HealthCheckSslHealthCheckArgs() {}

    private HealthCheckSslHealthCheckArgs(HealthCheckSslHealthCheckArgs $) {
        this.port = $.port;
        this.portName = $.portName;
        this.portSpecification = $.portSpecification;
        this.proxyHeader = $.proxyHeader;
        this.request = $.request;
        this.response = $.response;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HealthCheckSslHealthCheckArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HealthCheckSslHealthCheckArgs $;

        public Builder() {
            $ = new HealthCheckSslHealthCheckArgs();
        }

        public Builder(HealthCheckSslHealthCheckArgs defaults) {
            $ = new HealthCheckSslHealthCheckArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param port The TCP port number for the HTTP2 health check request.
         * The default value is 443.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The TCP port number for the HTTP2 health check request.
         * The default value is 443.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param portName Port name as defined in InstanceGroup#NamedPort#name. If both port and
         * port_name are defined, port takes precedence.
         * 
         * @return builder
         * 
         */
        public Builder portName(@Nullable Output<String> portName) {
            $.portName = portName;
            return this;
        }

        /**
         * @param portName Port name as defined in InstanceGroup#NamedPort#name. If both port and
         * port_name are defined, port takes precedence.
         * 
         * @return builder
         * 
         */
        public Builder portName(String portName) {
            return portName(Output.of(portName));
        }

        /**
         * @param portSpecification Specifies how port is selected for health checking, can be one of the
         * following values:
         * * `USE_FIXED_PORT`: The port number in `port` is used for health checking.
         * * `USE_NAMED_PORT`: The `portName` is used for health checking.
         * * `USE_SERVING_PORT`: For NetworkEndpointGroup, the port specified for each
         *   network endpoint is used for health checking. For other backends, the
         *   port or named port specified in the Backend Service is used for health
         *   checking.
         *   If not specified, HTTP2 health check follows behavior specified in `port` and
         *   `portName` fields.
         *   Possible values are: `USE_FIXED_PORT`, `USE_NAMED_PORT`, `USE_SERVING_PORT`.
         * 
         * @return builder
         * 
         */
        public Builder portSpecification(@Nullable Output<String> portSpecification) {
            $.portSpecification = portSpecification;
            return this;
        }

        /**
         * @param portSpecification Specifies how port is selected for health checking, can be one of the
         * following values:
         * * `USE_FIXED_PORT`: The port number in `port` is used for health checking.
         * * `USE_NAMED_PORT`: The `portName` is used for health checking.
         * * `USE_SERVING_PORT`: For NetworkEndpointGroup, the port specified for each
         *   network endpoint is used for health checking. For other backends, the
         *   port or named port specified in the Backend Service is used for health
         *   checking.
         *   If not specified, HTTP2 health check follows behavior specified in `port` and
         *   `portName` fields.
         *   Possible values are: `USE_FIXED_PORT`, `USE_NAMED_PORT`, `USE_SERVING_PORT`.
         * 
         * @return builder
         * 
         */
        public Builder portSpecification(String portSpecification) {
            return portSpecification(Output.of(portSpecification));
        }

        /**
         * @param proxyHeader Specifies the type of proxy header to append before sending data to the
         * backend.
         * Default value is `NONE`.
         * Possible values are: `NONE`, `PROXY_V1`.
         * 
         * @return builder
         * 
         */
        public Builder proxyHeader(@Nullable Output<String> proxyHeader) {
            $.proxyHeader = proxyHeader;
            return this;
        }

        /**
         * @param proxyHeader Specifies the type of proxy header to append before sending data to the
         * backend.
         * Default value is `NONE`.
         * Possible values are: `NONE`, `PROXY_V1`.
         * 
         * @return builder
         * 
         */
        public Builder proxyHeader(String proxyHeader) {
            return proxyHeader(Output.of(proxyHeader));
        }

        /**
         * @param request The application data to send once the SSL connection has been
         * established (default value is empty). If both request and response are
         * empty, the connection establishment alone will indicate health. The request
         * data can only be ASCII.
         * 
         * @return builder
         * 
         */
        public Builder request(@Nullable Output<String> request) {
            $.request = request;
            return this;
        }

        /**
         * @param request The application data to send once the SSL connection has been
         * established (default value is empty). If both request and response are
         * empty, the connection establishment alone will indicate health. The request
         * data can only be ASCII.
         * 
         * @return builder
         * 
         */
        public Builder request(String request) {
            return request(Output.of(request));
        }

        /**
         * @param response The bytes to match against the beginning of the response data. If left empty
         * (the default value), any response will indicate health. The response data
         * can only be ASCII.
         * 
         * @return builder
         * 
         */
        public Builder response(@Nullable Output<String> response) {
            $.response = response;
            return this;
        }

        /**
         * @param response The bytes to match against the beginning of the response data. If left empty
         * (the default value), any response will indicate health. The response data
         * can only be ASCII.
         * 
         * @return builder
         * 
         */
        public Builder response(String response) {
            return response(Output.of(response));
        }

        public HealthCheckSslHealthCheckArgs build() {
            return $;
        }
    }

}
