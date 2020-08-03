// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package civo

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LoadBalancer struct {
	pulumi.CustomResourceState

	// a list of backend instances, each containing an instance_id, protocol (http or https) and port
	Backends LoadBalancerBackendArrayOutput `pulumi:"backends"`
	// how long to wait in seconds before determining a backend has failed, defaults to 30
	FailTimeout pulumi.IntOutput `pulumi:"failTimeout"`
	// what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
	HealthCheckPath pulumi.StringPtrOutput `pulumi:"healthCheckPath"`
	// the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
	// blank)
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// should self-signed/invalid certificates be ignored from the backend servers, defaults to true
	IgnoreInvalidBackendTls pulumi.BoolPtrOutput `pulumi:"ignoreInvalidBackendTls"`
	// how many concurrent connections can each backend handle, defaults to 10
	MaxConns pulumi.IntOutput `pulumi:"maxConns"`
	// the size in megabytes of the maximum request content that will be accepted, defaults to 20
	MaxRequestSize pulumi.IntOutput `pulumi:"maxRequestSize"`
	// one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
	// round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
	// same backend), default is random
	Policy pulumi.StringOutput `pulumi:"policy"`
	// you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
	// (commonly 80 for HTTP or 443 for HTTPS)
	Port pulumi.IntOutput `pulumi:"port"`
	// either http or https. If you specify https then you must also provide the next two fields, the default is http
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
	TlsCertificate pulumi.StringPtrOutput `pulumi:"tlsCertificate"`
	// if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
	TlsKey pulumi.StringPtrOutput `pulumi:"tlsKey"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil || args.Backends == nil {
		return nil, errors.New("missing required argument 'Backends'")
	}
	if args == nil || args.FailTimeout == nil {
		return nil, errors.New("missing required argument 'FailTimeout'")
	}
	if args == nil || args.Hostname == nil {
		return nil, errors.New("missing required argument 'Hostname'")
	}
	if args == nil || args.MaxConns == nil {
		return nil, errors.New("missing required argument 'MaxConns'")
	}
	if args == nil || args.MaxRequestSize == nil {
		return nil, errors.New("missing required argument 'MaxRequestSize'")
	}
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.Port == nil {
		return nil, errors.New("missing required argument 'Port'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil {
		args = &LoadBalancerArgs{}
	}
	var resource LoadBalancer
	err := ctx.RegisterResource("civo:index/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("civo:index/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// a list of backend instances, each containing an instance_id, protocol (http or https) and port
	Backends []LoadBalancerBackend `pulumi:"backends"`
	// how long to wait in seconds before determining a backend has failed, defaults to 30
	FailTimeout *int `pulumi:"failTimeout"`
	// what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
	// blank)
	Hostname *string `pulumi:"hostname"`
	// should self-signed/invalid certificates be ignored from the backend servers, defaults to true
	IgnoreInvalidBackendTls *bool `pulumi:"ignoreInvalidBackendTls"`
	// how many concurrent connections can each backend handle, defaults to 10
	MaxConns *int `pulumi:"maxConns"`
	// the size in megabytes of the maximum request content that will be accepted, defaults to 20
	MaxRequestSize *int `pulumi:"maxRequestSize"`
	// one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
	// round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
	// same backend), default is random
	Policy *string `pulumi:"policy"`
	// you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
	// (commonly 80 for HTTP or 443 for HTTPS)
	Port *int `pulumi:"port"`
	// either http or https. If you specify https then you must also provide the next two fields, the default is http
	Protocol *string `pulumi:"protocol"`
	// if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
	TlsCertificate *string `pulumi:"tlsCertificate"`
	// if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
	TlsKey *string `pulumi:"tlsKey"`
}

type LoadBalancerState struct {
	// a list of backend instances, each containing an instance_id, protocol (http or https) and port
	Backends LoadBalancerBackendArrayInput
	// how long to wait in seconds before determining a backend has failed, defaults to 30
	FailTimeout pulumi.IntPtrInput
	// what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
	HealthCheckPath pulumi.StringPtrInput
	// the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
	// blank)
	Hostname pulumi.StringPtrInput
	// should self-signed/invalid certificates be ignored from the backend servers, defaults to true
	IgnoreInvalidBackendTls pulumi.BoolPtrInput
	// how many concurrent connections can each backend handle, defaults to 10
	MaxConns pulumi.IntPtrInput
	// the size in megabytes of the maximum request content that will be accepted, defaults to 20
	MaxRequestSize pulumi.IntPtrInput
	// one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
	// round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
	// same backend), default is random
	Policy pulumi.StringPtrInput
	// you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
	// (commonly 80 for HTTP or 443 for HTTPS)
	Port pulumi.IntPtrInput
	// either http or https. If you specify https then you must also provide the next two fields, the default is http
	Protocol pulumi.StringPtrInput
	// if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
	TlsCertificate pulumi.StringPtrInput
	// if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
	TlsKey pulumi.StringPtrInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// a list of backend instances, each containing an instance_id, protocol (http or https) and port
	Backends []LoadBalancerBackend `pulumi:"backends"`
	// how long to wait in seconds before determining a backend has failed, defaults to 30
	FailTimeout int `pulumi:"failTimeout"`
	// what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
	HealthCheckPath *string `pulumi:"healthCheckPath"`
	// the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
	// blank)
	Hostname string `pulumi:"hostname"`
	// should self-signed/invalid certificates be ignored from the backend servers, defaults to true
	IgnoreInvalidBackendTls *bool `pulumi:"ignoreInvalidBackendTls"`
	// how many concurrent connections can each backend handle, defaults to 10
	MaxConns int `pulumi:"maxConns"`
	// the size in megabytes of the maximum request content that will be accepted, defaults to 20
	MaxRequestSize int `pulumi:"maxRequestSize"`
	// one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
	// round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
	// same backend), default is random
	Policy string `pulumi:"policy"`
	// you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
	// (commonly 80 for HTTP or 443 for HTTPS)
	Port int `pulumi:"port"`
	// either http or https. If you specify https then you must also provide the next two fields, the default is http
	Protocol string `pulumi:"protocol"`
	// if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
	TlsCertificate *string `pulumi:"tlsCertificate"`
	// if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
	TlsKey *string `pulumi:"tlsKey"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// a list of backend instances, each containing an instance_id, protocol (http or https) and port
	Backends LoadBalancerBackendArrayInput
	// how long to wait in seconds before determining a backend has failed, defaults to 30
	FailTimeout pulumi.IntInput
	// what URL should be used on the backends to determine if it's OK (2xx/3xx status), defaults to /
	HealthCheckPath pulumi.StringPtrInput
	// the hostname to receive traffic for, e.g. www.example.com (optional: sets hostname to loadbalancer-uuid.civo.com if
	// blank)
	Hostname pulumi.StringInput
	// should self-signed/invalid certificates be ignored from the backend servers, defaults to true
	IgnoreInvalidBackendTls pulumi.BoolPtrInput
	// how many concurrent connections can each backend handle, defaults to 10
	MaxConns pulumi.IntInput
	// the size in megabytes of the maximum request content that will be accepted, defaults to 20
	MaxRequestSize pulumi.IntInput
	// one of: least_conn (sends new requests to the least busy server), random (sends new requests to a random backend),
	// round_robin (sends new requests to the next backend in order), ip_hash (sends requests from a given IP address to the
	// same backend), default is random
	Policy pulumi.StringInput
	// you can listen on any port, the default is 80 to match the default protocol of http,if not you must specify it here
	// (commonly 80 for HTTP or 443 for HTTPS)
	Port pulumi.IntInput
	// either http or https. If you specify https then you must also provide the next two fields, the default is http
	Protocol pulumi.StringInput
	// if your protocol is https then you should send the TLS certificate in Base64-encoded PEM format
	TlsCertificate pulumi.StringPtrInput
	// if your protocol is https then you should send the TLS private key in Base64-encoded PEM format
	TlsKey pulumi.StringPtrInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}