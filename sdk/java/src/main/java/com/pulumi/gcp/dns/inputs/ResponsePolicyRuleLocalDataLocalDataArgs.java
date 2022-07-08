// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ResponsePolicyRuleLocalDataLocalDataArgs extends com.pulumi.resources.ResourceArgs {

    public static final ResponsePolicyRuleLocalDataLocalDataArgs Empty = new ResponsePolicyRuleLocalDataLocalDataArgs();

    /**
     * For example, www.example.com.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return For example, www.example.com.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)
     * 
     */
    @Import(name="rrdatas")
    private @Nullable Output<List<String>> rrdatas;

    /**
     * @return As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)
     * 
     */
    public Optional<Output<List<String>>> rrdatas() {
        return Optional.ofNullable(this.rrdatas);
    }

    /**
     * Number of seconds that this ResourceRecordSet can be cached by
     * resolvers.
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return Number of seconds that this ResourceRecordSet can be cached by
     * resolvers.
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * One of valid DNS resource types.
     * Possible values are `A`, `AAAA`, `CAA`, `CNAME`, `DNSKEY`, `DS`, `HTTPS`, `IPSECVPNKEY`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, and `TXT`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return One of valid DNS resource types.
     * Possible values are `A`, `AAAA`, `CAA`, `CNAME`, `DNSKEY`, `DS`, `HTTPS`, `IPSECVPNKEY`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, and `TXT`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ResponsePolicyRuleLocalDataLocalDataArgs() {}

    private ResponsePolicyRuleLocalDataLocalDataArgs(ResponsePolicyRuleLocalDataLocalDataArgs $) {
        this.name = $.name;
        this.rrdatas = $.rrdatas;
        this.ttl = $.ttl;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ResponsePolicyRuleLocalDataLocalDataArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ResponsePolicyRuleLocalDataLocalDataArgs $;

        public Builder() {
            $ = new ResponsePolicyRuleLocalDataLocalDataArgs();
        }

        public Builder(ResponsePolicyRuleLocalDataLocalDataArgs defaults) {
            $ = new ResponsePolicyRuleLocalDataLocalDataArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name For example, www.example.com.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name For example, www.example.com.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rrdatas As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)
         * 
         * @return builder
         * 
         */
        public Builder rrdatas(@Nullable Output<List<String>> rrdatas) {
            $.rrdatas = rrdatas;
            return this;
        }

        /**
         * @param rrdatas As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)
         * 
         * @return builder
         * 
         */
        public Builder rrdatas(List<String> rrdatas) {
            return rrdatas(Output.of(rrdatas));
        }

        /**
         * @param rrdatas As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1)
         * 
         * @return builder
         * 
         */
        public Builder rrdatas(String... rrdatas) {
            return rrdatas(List.of(rrdatas));
        }

        /**
         * @param ttl Number of seconds that this ResourceRecordSet can be cached by
         * resolvers.
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Number of seconds that this ResourceRecordSet can be cached by
         * resolvers.
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param type One of valid DNS resource types.
         * Possible values are `A`, `AAAA`, `CAA`, `CNAME`, `DNSKEY`, `DS`, `HTTPS`, `IPSECVPNKEY`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, and `TXT`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type One of valid DNS resource types.
         * Possible values are `A`, `AAAA`, `CAA`, `CNAME`, `DNSKEY`, `DS`, `HTTPS`, `IPSECVPNKEY`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV`, `SSHFP`, `SVCB`, `TLSA`, and `TXT`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ResponsePolicyRuleLocalDataLocalDataArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}
